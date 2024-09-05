package http

import (
	"context"

	"github.com/nbonair/GoIDM/internal/dataaccess/db/migrations"
	"github.com/nbonair/GoIDM/internal/handlers/consumer"
	"github.com/nbonair/GoIDM/internal/logic"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type Server interface {
	Start() error
}

type server struct {
	pocketbaseApp *pocketbase.PocketBase
}

func NewServer() Server {
	pocketbaseApp := pocketbase.New()
	pocketbaseApp.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		dbMigrator := migrations.NewMigrator(e.App)
		if err := dbMigrator.Migrate(); err != nil {
			return err
		}

		downloadTaskLogic := logic.NewDownloadTask()
		downloadTaskCreatedConsumer := consumer.NewDownloadTaskCreated(downloadTaskLogic)
		e.App.OnRecordAfterCreateRequest("download_task").Add(func(e *core.RecordCreateEvent) error {
			return downloadTaskCreatedConsumer.Handle(context.Background(), e)
		})
		return nil
	})
	return &server{
		pocketbaseApp: pocketbaseApp,
	}
}

func (s server) Start() error {
	return s.pocketbaseApp.Start()
}
