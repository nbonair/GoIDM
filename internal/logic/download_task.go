package logic

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
)

type ExecuteDownloadTaskParams struct {
	ID           string
	DownloadType int
	URL          string
}

type DownloadTask interface {
	ExecuteDownloadTask(ctx context.Context, params ExecuteDownloadTaskParams) error
}

type downloadTask struct{}

func NewDownloadTask() DownloadTask {
	return &downloadTask{}
}

func (d *downloadTask) ExecuteDownloadTask(ctx context.Context, params ExecuteDownloadTaskParams) error {
	switch params.DownloadType {
	case 1: //HTTP
		response, err := http.Get(params.URL)
		if err != nil {
			return err
		}

		defer response.Body.Close()

		responseBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		file, err := os.Create(params.ID)
		if err != nil {
			return err
		}

		_, err = file.Write(responseBytes)
		if err != nil {
			return err
		}
		return nil

	default:
		return errors.New("unsupported download type")
	}
}
