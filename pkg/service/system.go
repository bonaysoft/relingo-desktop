package service

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/bonaysoft/relingo-desktop/pkg/proxy"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type System struct {
	ctx context.Context
}

func NewSystem(ctx context.Context) *System {
	return &System{
		ctx: ctx,
	}
}

func (a *System) DownloadCert() error {
	opts := runtime.SaveDialogOptions{
		DefaultFilename: "relingo-desktop.crt",
	}
	v, err := runtime.SaveFileDialog(a.ctx, opts)
	if err != nil {
		return err
	}

	dist, err := os.Create(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(dist, bytes.NewReader(proxy.GetCert()))
	return err
}
