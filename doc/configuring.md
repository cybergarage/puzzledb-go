# Configuring PuzzleDB

PuzzleDB is configured using a configuration file. The default configuration file is located at `conf/puzzledb.yaml`.

## Configuration File (puzzledb.yaml)

The configuration file is a YAML file. The configuration file is divided into sections. Each section is a YAML map. The following is an example configuration file:

    logging:
      enabled: true
      level: info
    api:
      grpc:
        enabled: true
        port: 50053
    metrics:
      prometheus:
        enabled: true
        port: 9181
    tracing:
      enabled: true
      default: opentelemetry
      tracer: 
        opentelemetry:
          endpoint: "http://localhost:14268/api/traces"
        opentracing:
          endpoint: "localhost:6831"
    plugins:
      coder:
        document:
          default: cbor
          cbor:
            enabled: true
        key: 
          default: tuple
          tuple:
            enabled: true
      coordinator: 
        default: memdb
        memdb:
          enabled: true
        etcd:
          enabled: true
          endpoints: null
        fdb:
          enabled: true
          cluseterFile: null
      store:
        document:
          default: dockv
          kv:
            enabled: true
            store: memdb
        kv:
          default: memdb
          memdb:
            enabled: true
            endpoints: null
          fdb:
            enabled: true
            cluseterFile: null
      query:
        redis:
          enabled: true
          port: 6379
        mongo:
          enabled: true
          port: 27017
        mysql:
          enabled: true
          port: 3306

## Environment Variables

You can override the configuration file location by setting the PUZZLEDB environment variable. PuzzleDB assumes that the environment variable matches the following format: PUZZLEDB + "\_" + the key name in ALL CAPS.

For example, if the configuration parameter is "logging:enabled", PuzzleDB will look for the environment variable "PUZZLEDB\_LOGGING\_ENABLED".
