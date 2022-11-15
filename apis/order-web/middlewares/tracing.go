package middlewares

import (
	"apis/order-web/global"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func Trace() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d", global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port),
			},
			ServiceName: global.ServerConfig.JaegerInfo.Name,
		}

		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}
		opentracing.SetGlobalTracer(tracer)
		defer closer.Close()

		startSpan := tracer.StartSpan(string(c.Request.URI().Path()))
		defer startSpan.Finish()

		c.Set("tracer", tracer)
		c.Set("parentSpan", startSpan)
		c.Next(ctx)
	}
}
