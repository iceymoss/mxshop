package middlewares

import (
	"fmt"

	"mxshop-api/order-web/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

//Trace 链路追踪
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		//初始化配置
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d", global.ServerConfig.Tracing.Host, global.ServerConfig.Tracing.Port),
			},
			ServiceName: global.ServerConfig.Tracing.Name,
		}

		//初始化跟踪连
		Tracer, Closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}

		opentracing.SetGlobalTracer(Tracer)
		defer Closer.Close()

		startSpan := Tracer.StartSpan(c.Request.URL.Path)
		defer startSpan.Finish()

		c.Set("tracer", Tracer)
		c.Set("parentSpan", startSpan)

		c.Next()
	}
}
