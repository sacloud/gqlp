// Copyright 2021 The gqlp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trace

import (
	"log"

	"github.com/sacloud/gqlp/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var Tracer = otel.GetTracerProvider().Tracer(
	"github.com/sacloud/gqlp",
	trace.WithInstrumentationVersion(version.Version),
)

func InitTracer() func() {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(jaeger.CollectorEndpointFromEnv()),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}), // TODO Sampler should be configurable
	)
	if err != nil {
		log.Fatal(err)
	}
	return flush
}
