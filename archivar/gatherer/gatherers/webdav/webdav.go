package webdav

import (
	"github.com/rwese/archivar/archivar/archiver/archivers"
	"github.com/rwese/archivar/archivar/gatherer/gatherers"
	"github.com/rwese/archivar/internal/utils/config"
	webdavClient "github.com/rwese/archivar/internal/webdav"
	"github.com/sirupsen/logrus"
)

// Webdav allows to upload files to a remote webdav server
type Webdav struct {
	storage          archivers.Archiver
	logger           *logrus.Logger
	client           *webdavClient.Webdav
	directory        string
	deleteDownloaded bool
}

type WebdavConfig struct {
	Directory        string
	DeleteDownloaded bool
}

func init() {
	gatherers.Register(New)
}

// New will return a new webdav downloader
func New(c interface{}, storage archivers.Archiver, logger *logrus.Logger) gatherers.Gatherer {
	wc := &WebdavConfig{}
	config.ConfigFromStruct(c, &wc)

	webdav := &Webdav{
		storage:          storage,
		logger:           logger,
		client:           webdavClient.New(c, logger),
		directory:        wc.Directory,
		deleteDownloaded: wc.DeleteDownloaded,
	}
	return webdav
}

func (w Webdav) Download() (err error) {
	if err = w.Connect(); err != nil {
		return
	}

	var downloadedFiles []string
	if downloadedFiles, err = w.client.DownloadFiles(w.directory, w.storage.Upload); err != nil {
		return
	}

	if w.deleteDownloaded {
		err = w.client.DeleteFiles(downloadedFiles)
		if err != nil {
			return err
		}
	}

	return
}

func (w *Webdav) Connect() (err error) {

	if err = w.client.Connect(); err != nil {
		return
	}

	if !w.client.DirExists(w.directory) {
		w.logger.Fatalf("failed to access upload directory, which will not be automatically created: %s", err.Error())
	}
	return
}
