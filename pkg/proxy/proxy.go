package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"

	"gopkg.in/elazarl/goproxy.v1"
)

type Proxy struct {
	srv   *http.Server
	query *query.Query
	rc    *relingo.Client
}

func NewProxy(query *query.Query, rc *relingo.Client) *Proxy {
	return &Proxy{
		srv:   &http.Server{Addr: ":8119"},
		rc:    rc,
		query: query,
	}
}

func (p *Proxy) Run() error {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest(goproxy.ReqHostIs("api.relingo.net:443")).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		authzToken := req.Header.Get("x-relingo-token")
		p.rc.SetToken(authzToken)
		return req, nil
	})
	proxy.OnResponse(goproxy.UrlIs("/api/parseContent2")).DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp
		}

		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		go p.hookWords(bodyBytes)
		return resp
	})
	p.srv.Handler = proxy
	return p.srv.ListenAndServe()
}

func (p *Proxy) Stop(ctx context.Context) (err error) {
	fmt.Println("Proxy stopped11")
	return p.srv.Shutdown(ctx)
}

func (p *Proxy) hookWords(body []byte) {
	var data relingo.Response[relingo.RespParseContent2]
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return
	}

	for _, word := range data.Data.Words {
		w, err := p.query.Word.Where(p.query.Word.Source.Eq(word.Source)).Take()
		if err != nil {
			// return
			w = &model.Word{Source: word.Source}
		}

		w.Exposures++
		w.Mastered = word.Mastered
		// fmt.Println(word.Source, w.Exposures)
		if err := p.query.Word.Where(p.query.Word.Id.Eq(w.Id)).Save(w); err != nil {
			log.Println(err)
			return
		}
	}
}
