package router

import (
	"fmt"
	"strconv"
	"time"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/middleware"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/infrastructure/database"
	"storys-lab-fishing-api/app/infrastructure/log"
)

type appConfig struct {
	appName 		string
	logger  		logger.Logger
	validator     	*validator.Validator
	ctxTimeout    	time.Duration
	dbSQL         	*gorm.DB
	webServerPort 	Port
	webServer     	Server
}

func NewAppConfig() *appConfig {
	return &appConfig{}
}

func newGinServer(
	log 		logger.Logger,
	db 			*gorm.DB,
	validator 	*validator.Validator,
	port 		Port,
	t 			time.Duration,
) *ginEngine {
	return &ginEngine{
		router: 	gin.New(),
		log:   		log,
		db:     	db,
		validator:  *validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	fmt.Println("listen...")

	gin.SetMode(gin.ReleaseMode)

	// サーバーがpanicによるクラッシュから復旧するためのミドルウェアを設定
	gin.Recovery()

	// カスタム関数を用いて、HTTP通信に対するハンドラーをルーターに設定
	g.router.Use(middleware.SetupCORS())
	g.setAppHandlers(g.router)

	// HTTPサーバーの設定を定義
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	// 終了シグナルを受け取るためのチャネルを作成
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// go routineでサーバーを起動
	go func() {
		g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		// ListenAndServeでサーバー起動 それ以外はログ出力
		if err := server.ListenAndServe(); err != nil {
			g.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	// シグナルが受け取られるまでプログラムをブロック
	<-stop

	// シャットダウンプロセスの初期化
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	// サーバーのシャットダウン
	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	g.log.Infof("Service down")
}

func (c *appConfig) ContextTimeout(t time.Duration) *appConfig {
	c.ctxTimeout = t
	return c
}

func (c *appConfig) Name(name string) *appConfig {
	c.appName = name
	return c
}

func (c *appConfig) Logger(instance int) *appConfig {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully appConfigured log")
	return c
}

// DB接続
func (c *appConfig) DbSQL() *appConfig {
	db, err := database.DBConnect()
	if err != nil {
		c.logger.Fatalln(err)
	}
	// fmt.Println(db)

	c.logger.Infof("Successfully connected to the SQL database")
	c.dbSQL = db
	return c
}

func (c *appConfig) Validator() *appConfig {
	v := validator.NewValidator()

	c.logger.Infof("Successfully appConfigured validator")

	c.validator = v
	return c
}

func (c *appConfig) WebServer(instance int) *appConfig {
	s, err := NewWebServerFactory(
		instance,
		c.logger,
		c.dbSQL,
		c.validator,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully appConfigured router server")

	c.webServer = s
	return c
}

func (c *appConfig) WebServerPort(port string) *appConfig {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}
	c.webServerPort = Port(p)
	return c
}

func (c *appConfig) Start() {
	fmt.Println("start listen...")
	c.webServer.Listen()
}

