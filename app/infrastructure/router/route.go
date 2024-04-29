package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"storys-lab-fishing-api/app/adapter/api/action"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/presenter"
	"storys-lab-fishing-api/app/adapter/repository"
	"storys-lab-fishing-api/app/usecase"
)

type ginEngine struct {
	router *gin.Engine
	log    logger.Logger
	db     repository.SQL
	// validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGinServer(
	log logger.Logger,
	db repository.SQL,
	// validator validator.Validator,
	port Port,
	t time.Duration,
) *ginEngine {
	return &ginEngine{
		// router:     gin.Default(),
		router: gin.New(),
		log:    log,
		db:     db,
		// validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	fmt.Println("listen...")

	// サーバーをデバッグモードに設定
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	// サーバーがpanicによるクラッシュから復旧するためのミドルウェアを設定
	gin.Recovery()

	// カスタム関数を用いて、HTTP通信に対するハンドラーをルーターに設定
	g.router.Use(SetupCORS())
	g.setAppHandlers(g.router)

	// HTTPサーバーの設定を定義
	server := &http.Server{
		ReadTimeout:  5 * time.Second,            // タイムアウト設定（読み込み）
		WriteTimeout: 15 * time.Second,           // タイムアウト設定（書き込み）
		Addr:         fmt.Sprintf(":%d", g.port), // サーバーのリスニングアドレス
		Handler:      g.router,                   // HTTPリクエストを処理するハンドラー
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

func (g ginEngine) setAppHandlers(r *gin.Engine) {

	r.GET("/fishes", g.buildFindAllFishAction())

	r.GET("/health", g.healthCheck())
}

func (g ginEngine) buildFindAllFishAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewFindAllFishInteractor(
				repository.NewFishSQL(g.db),
				presenter.NewFindAllFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllFishAction(uc, g.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func SetupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
