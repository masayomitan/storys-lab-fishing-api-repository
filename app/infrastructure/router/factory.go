package router

import (
	// "errors"
	"time"
	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/validator"
)

type Server interface {
	Listen()
}

type Port int

const (
	InstanceGorillaMux int = iota
	InstanceGin
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	SQL *gorm.DB,
	validator *validator.Validator,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	return newGinServer(
		log, 
		SQL, 
		validator, 
		port, 
		ctxTimeout,
	), nil
}
