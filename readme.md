## todo: make file to support cross-os builds.

### todo: primary features
- jsonify message pair.
- connection pools for remotes?
- protocol parser: http?
- protocol parser: mysql?
- protocol parser: redis?
- web sink?
- file sink?
- inmemory sink
- replay support
- sink server (kafka, pulsar, db, etc).
-

### todo: mode -> record traffic
- sinks - file, redis, inmemory, kafka, pulsar, database etc.

### todo: mode -> replay traffic
- request matcher (bytestream)?
- request matcher (parsed request)?

### todo: explore alternate approaches
- port mirroring
- via libpcap
- via ha-proxy plugin
- via envoy-proxy plugin
