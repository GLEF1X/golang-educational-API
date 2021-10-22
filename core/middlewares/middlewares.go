package middlewares

import (
	"fmt"
	log "github.com/GLEF1X/golang-educational-API/core/logging"
	realip "github.com/GLEF1X/golang-educational-API/utils"
	"github.com/lestrrat-go/strftime"
	"github.com/valyala/fasthttp"
	"mime"
	"time"
)

type measureTimeUtil struct {
	timeFormatter *strftime.Strftime
}

func NewMeasureTimeUtil() *measureTimeUtil {
	f, err := strftime.New("%F %T")
	if err != nil {
		log.GetLogger().Fatal(err)
	}
	return &measureTimeUtil{timeFormatter: f}
}

func EnforceJSONMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		contentType := string(ctx.Request.Header.Peek("Content-Type"))

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				ctx.Error("Malformed Content-Type header", fasthttp.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				ctx.Error("Content-Type header must be application/json", fasthttp.StatusUnsupportedMediaType)
				return
			}
		}

		next(ctx)
	})
}

func (u *measureTimeUtil) MeasureResponseTimeMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		next(ctx)
		elapsed := time.Since(start)
		log.GetLogger().Info(
			fmt.Sprintf(

				"%v Request from %v handled for %v",
				u.timeFormatter.FormatString(time.Now()),
				realip.FromRequest(ctx),
				elapsed,
			),
		)
	})
}
