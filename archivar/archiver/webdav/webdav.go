package webdav

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/studio-b12/gowebdav"
)

// Webdav allows to upload files to a remote webdav server
type Webdav struct {
	server               string
	userName             string
	userPassword         string
	keepUploaded         bool
	uploadDirectory      string
	logger               *logrus.Logger
	clientConnectedSince time.Time
	client               *gowebdav.Client
}

// New will return a new webdav uploader
func New(server string, userName string, userPassword string, uploadDirectory string, logger *logrus.Logger) *Webdav {
	return &Webdav{
		server:          server,
		userName:        userName,
		userPassword:    userPassword,
		keepUploaded:    false,
		uploadDirectory: uploadDirectory,
		logger:          logger,
	}
}