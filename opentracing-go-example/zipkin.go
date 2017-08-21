package main

import (
	"context"
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"

	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

// Based from https://github.com/openzipkin/zipkin-go-opentracing/tree/master/examples/cli_with_2_services

const (
	// Endpoint to send Zipkin span to
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"
	// Debug mode
	debug = false
	// Our service name
	serviceName = "cli"
	// Host + Port of our service
	hostPort = "0.0.0.0:0"
)

func service(ctx context.Context) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("func", "service")
}

func main() {
	// Create HTTP collector.
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("Unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}
	// Create recorder
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)
	// Create tracer
	tracer, err := zipkin.NewTracer(recorder, zipkin.ClientServerSameSpan(true), zipkin.TraceID128Bit(true))
	if err != nil {
		fmt.Printf("Unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}
	// Explicitly set this tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)
	// Create root span for duration of interaction
	span := opentracing.StartSpan("Main")

	// Put root span in context so it will be used
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	span.LogEvent("Call service")
	service(ctx)

	span.Finish()
	collector.Close()
}
