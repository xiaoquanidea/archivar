package filesystem

import (
	"github.com/rwese/archivar/archivar/archiver/archivers"
	"github.com/rwese/archivar/archivar/gatherer/gatherers"
	filesystemClient "github.com/rwese/archivar/internal/filesystem"
	"github.com/rwese/archivar/internal/utils/config"
	"github.com/sirupsen/logrus"
)

//
type filesystem struct {
	storage          archivers.Archiver
	logger           *logrus.Logger
	client           *filesystemClient.FileSystem
	directory        string
	deleteDownloaded bool
}

type filesystemConfig struct {
	Directory        string
	DeleteDownloaded bool
}

func init() {
	gatherers.Register(New)
}

// New will return a new fs downloader
func New(c interface{}, storage archivers.Archiver, logger *logrus.Logger) gatherers.Gatherer {
	wc := &filesystemConfig{}
	config.ConfigFromStruct(c, &wc)

	filesystem := &filesystem{
		storage:          storage,
		logger:           logger,
		client:           filesystemClient.New(logger),
		directory:        wc.Directory,
		deleteDownloaded: wc.DeleteDownloaded,
	}
	return filesystem
}

func (w filesystem) Download() (err error) {
	if err = w.Connect(); err != nil {
		return
	}

	var downloadedFiles []string
	if err = w.client.DownloadFiles(w.directory, w.storage.Upload); err != nil {
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

func (w *filesystem) Connect() (err error) {
	if !w.client.DirExists(w.directory) {
		w.logger.Fatalf("failed to access upload directory, which will not be automatically created: %s", err.Error())
	}
	return
}
