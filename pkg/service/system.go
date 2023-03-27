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
	ctxGetter func() context.Context
}

func NewSystem(ctxGetter func() context.Context) *System {
	return &System{
		ctxGetter: ctxGetter,
	}
}

func (sys *System) DownloadCert() error {
	opts := runtime.SaveDialogOptions{
		DefaultFilename: "relingo-desktop.crt",
	}
	v, err := runtime.SaveFileDialog(sys.ctxGetter(), opts)
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
