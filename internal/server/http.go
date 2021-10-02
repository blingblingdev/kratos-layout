package server

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdHttp "net/http"
	v1 "github.com/ffy/kratos-layout/api/manage/v1"
	"github.com/ffy/kratos-layout/internal/conf"
	"github.com/ffy/kratos-layout/internal/service"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Code    int64       `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Ts      string      `json:"ts" form:"ts"`
	Data    interface{} `json:"data" form:"data"`
}

const (
	baseContentType = "application"
)

func newResponse() *Response {
	return &Response{
		Ts: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}
}

// contentType returns the content-type with base prefix.
func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) (err error) {
	reply := newResponse()
	reply.Code = 200
	reply.Data = v
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(stdHttp.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	se := errors.FromError(err)
	reply := newResponse()
	reply.Code = int64(se.Code)
	reply.Data = nil
	reply.Message = se.Message

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(stdHttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(stdHttp.StatusOK)
	_, _ = w.Write(body)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, manage *service.ManageService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
		http.ResponseEncoder(responseEncoder),
		http.ErrorEncoder(errorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterManageHTTPServer(srv, manage)
	return srv
}
