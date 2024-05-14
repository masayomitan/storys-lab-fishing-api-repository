package infrastructure

import (
	"fmt"
	"storys-lab-fishing-api/app/adapter/repository"
	"strconv"
	"time"
	// "github.com/gsabadini/go-clean-architecture/adapter/validator"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/infrastructure/database"
	"storys-lab-fishing-api/app/infrastructure/log"
	"storys-lab-fishing-api/app/infrastructure/router"
)

type config struct {
	appName string
	logger  logger.Logger
	// validator     validator.Validator
	ctxTimeout    time.Duration
	dbSQL         repository.DBMethods
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
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
	c.dbSQL = &repository.GormAdapter{DB: db}
	return c
}

// func (c *config) Validator(instance int) *config {
// 	v, err := validation.NewValidatorFactory(instance)
// 	if err != nil {
// 		c.logger.Fatalln(err)
// 	}

// 	c.logger.Infof("Successfully configured validator")

// 	c.validator = v
// 	return c
// }

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.dbSQL,
		// c.validator,
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
	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	fmt.Println("start listen...")
	c.webServer.Listen()
}

