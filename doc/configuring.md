# Configuring PuzzleDB

PuzzleDB is configured using a YAML file; settings can be overridden through environment variables.

## Configuration File (puzzledb.yaml)

The configuration file is divided into sections (YAML maps). A default configuration is activated if no file is specified and `puzzledb.yaml` is absent in the working directory. The default is shown below:

    logger:
      enabled: true
      level: info
    pprof:
      enabled: false
      port: 6060
    tls:
      enabled: false
      key_file: key.pem
      cert_file: cert.pem
      ca_files: [ca.pem]
    plugins:
      auth:
        enabled: false
        plain:
          -
            username: admin
            password: password
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
          enabled: true
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
          tls_port: 0
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
        enabled: false
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

Define environment variables matching `PUZZLEDB_<UPPERCASE_KEY>` to override individual settings. (The configuration file path itself can also be overridden.)

Example: `PUZZLEDB_LOGGER_ENABLED=true` overrides `logging.enabled`.
