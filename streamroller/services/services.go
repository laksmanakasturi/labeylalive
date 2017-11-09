package services

import (
	"github.com/dustinblackman/joy4/format/rtmp"
	"github.com/dustinblackman/streamroller/logger"
)

// Service is the interface for Services
type Service interface {
	CanConnect() bool
	ConnectChat() error
	ConnectRTMP() *rtmp.Conn
	Name() string
}

// Services is an accessible export to list all supported services
var Services []Service

func connectRTMP(url string) *rtmp.Conn {
	conn, err := rtmp.Dial(url)
	if err != nil {
		logger.Log.Error(err)
	}

	return conn
}

// RegisterService is called on init for all services
func RegisterService(service Service) {
	Services = append(Services, service)
}
