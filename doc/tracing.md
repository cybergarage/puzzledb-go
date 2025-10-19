# Distributed Tracing

Distributed tracing analyzes end‑to‑end request flows across services, identifying latency, bottlenecks, and failure points. Unique trace IDs and spans depict each step, enabling visualization, optimization, and faster troubleshooting.

## OpenTelemetry Integration

OpenTelemetry provides a unified API and SDK for collecting, processing, and exporting traces, metrics, and logs to backends like Jaeger, Prometheus, or Elasticsearch.

PuzzleDB supports OpenTelemetry instrumentation for collecting database telemetry and exporting it to observability backends. For configuration details see:

- [Configuring PuzzleDB](configuring.md)

## References

- [OpenTelemetry](https://opentelemetry.io)

- [OpenTracing specification](https://opentracing.io/specification/)

  - [Migrating from OpenTracing](https://opentelemetry.io/docs/migration/opentracing/)
