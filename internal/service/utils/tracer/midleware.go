package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	zerologger "github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
	"io"
)

func SetJaegerTracer(traceHeader string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		zerologger.Error().Msg(fmt.Sprintf("Could not parse Jaeger env vars: %s", err.Error()))
		return nil, nil, err
	}
	customHeaders := &jaeger.HeadersConfig{
		TraceContextHeaderName: traceHeader,
	}
	cfg.Headers = customHeaders
	jLogger := jaegerlog.NullLogger
	jMetricsFactory := metrics.NullFactory
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		jaegercfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		jaegercfg.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		zerologger.Error().Msg(fmt.Sprintf("Could not create Jaeger traces: %s", err.Error()))
	}
	return tracer, closer, err
}
