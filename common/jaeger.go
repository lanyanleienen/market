package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracer(serviceName string, addr string)(opentracing.Tracer, io.Closer, error){
	cfg := &jaegercfg.Configuration{
		ServiceName:         serviceName,
		Sampler:             &jaegercfg.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
		},
		Reporter:            &jaegercfg.ReporterConfig{
			BufferFlushInterval:        1 * time.Second,
			LogSpans:                   true,
			LocalAgentHostPort:         addr,
		},
	}
	tracer,closer,err := cfg.NewTracer()
	return tracer,closer,err
}