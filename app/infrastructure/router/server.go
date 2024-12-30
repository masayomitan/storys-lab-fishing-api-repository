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

type config struct {
	appName 		string
	logger  		logger.Logger
	validator     	*validator.Validator
	ctxTimeout    	time.Duration
	dbSQL         	*gorm.DB
	webServerPort 	Port
	webServer     	Server
}

func NewConfig() *config {
	return &config{}
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

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")
	return c
}

// DB接続
func (c *config) DbSQL() *config {
	db, err := database.DBConnect()
	if err != nil {
		c.logger.Fatalln(err)
	}
	// fmt.Println(db)

	c.logger.Infof("Successfully connected to the SQL database")
	c.dbSQL = db
	return c
}

func (c *config) Validator() *config {
	v := validator.NewValidator()

	c.logger.Infof("Successfully configured validator")

	c.validator = v
	return c
}

func (c *config) WebServer(instance int) *config {
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

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}
	c.webServerPort = Port(p)
	return c
}

func (c *config) Start() {
	fmt.Println("start listen...")
	c.webServer.Listen()
}

