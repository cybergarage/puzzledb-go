# Tracing

OpenTelemetry is an open-source project that provides a standardized framework for collecting, processing, and exporting telemetry data, including traces, metrics, and logs. It offers developers a unified API to instrument their applications and send the data to different backends like Jaeger, Prometheus, or Elasticsearch for analysis and visualization purposes.

PuzzleDB supports OpenTelemetry integration. This means that developers can utilize OpenTelemetry to instrument their PuzzleDB instances and gather telemetry data from the database. Subsequently, the collected data can be transmitted to various observability tools, providing developers with a comprehensive overview of their PuzzleDB instances and facilitating efficient troubleshooting. For configuration details, please refer to the following documentation.

-   [Configuring PullzeDB](doc/configuring.md)

## References

-   [OpenTelemetry](https://opentelemetry.io)

-   [OpenTracing specification](https://opentracing.io/specification/n)

    -   [Migrating from OpenTracing | OpenTelemetry](https://opentelemetry.io/docs/migration/opentracing/)
