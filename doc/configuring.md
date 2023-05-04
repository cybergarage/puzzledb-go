# Configuring PuzzleDB

PuzzleDB is configured using a configuration file. The default configuration file is located at `conf/puzzledb.yaml`. You can override the configuration file location by setting the `PUZZLEDB` environment variable.

## puzzledb.yaml

    logging:
      enabled: true
      level: info
    grpc:
      enabled: true
      port: 50053
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
