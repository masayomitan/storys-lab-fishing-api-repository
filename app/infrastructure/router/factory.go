package router

import (
	"errors"
	"time"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/repository"
	"storys-lab-fishing-api/app/adapter/logger"
	// "storys-lab-fishing-api/adapter/validator"
)

type Server interface {
	Listen()
}

type Port int

const (
	InstanceGorillaMux int = iota
	InstanceGin
)

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	SQL repository.SQL,
	// validator validator.Validator,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	return newGinServer(
		gin.New(),
		log, 
		SQL, 
		// validator, 
		port, 
		ctxTimeout,
	), nil
}
