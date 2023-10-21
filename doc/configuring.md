# Configuring PuzzleDB

PuzzleDB is configured using a configuration file. The configuration file is a YAML file, and you can override the configuration by setting environment variables.

## Configuration File (puzzledb.yaml)

The configuration file is divided into sections. Each section is a YAML map. PuzzleDB will activate a default configuration if a configuration file is not specified or if there is no puzzledb.yaml in the local directory. The following is the default configuration file:

    logger:
      enabled: true
      level: info
    pprof:
      enabled: true
    plugins:
      system:
        grpc:
          enabled: true
          port: 50053
        actor:
          enabled: true
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
          cluster_file: null
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
          fdb:
            enabled: true
            cluster_file: null
        kvcache:
          default: ristretto
          ristretto:
            enabled: true
            num_counters: 1000
            max_cost: 1000000
            buffer_items: 64 
            metrics: false
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
        postgresql:
          enabled: true
          port: 5432
      tracer:
        default: opentelemetry
        opentelemetry:
          enabled: true
          endpoint: "http://localhost:14268/api/traces"
        opentracing:
          enabled: false
          endpoint: "localhost:6831"
      metrics:
        prometheus:
          enabled: true
          port: 9181

## Environment Variables

You can override the configuration file location by setting the PUZZLEDB environment variable. PuzzleDB assumes that the environment variable matches the following format: PUZZLEDB + "\_" + the key name in ALL CAPS.

For example, if the environment variable `PUZZLEDB_LOGGING_ENABLED` is set, then PuzzleDB will override the `logging:enabled` setting.
