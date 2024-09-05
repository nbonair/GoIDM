package consumer

import (
	"context"

	"github.com/nbonair/GoIDM/internal/logic"
	"github.com/pocketbase/pocketbase/core"
)

type DownloadTaskCreated interface {
	Handle(ctx context.Context, e *core.RecordCreateEvent) error
}

type downloadTaskCreated struct {
	downloadTaskLogic logic.DownloadTask
}

func NewDownloadTaskCreated(
	downloadTaskLogic logic.DownloadTask,
) DownloadTaskCreated {
	return &downloadTaskCreated{
		downloadTaskLogic: downloadTaskLogic,
	}
}

// Handle implements DownloadTaskCreated.
func (d *downloadTaskCreated) Handle(ctx context.Context, e *core.RecordCreateEvent) error {
	id := e.Record.Id
	downloadType := e.Record.GetInt("download_type")
	url := e.Record.GetString("url")
	return d.downloadTaskLogic.ExecuteDownloadTask(ctx, logic.ExecuteDownloadTaskParams{
		ID:           id,
		URL:          url,
		DownloadType: downloadType,
	})
}
