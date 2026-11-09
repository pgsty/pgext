## Usage

Sources:

- [Official README](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/README.md)
- [Extension implementation](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/src/lib.rs)
- [Extension control file](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/pg_webrtc.control)
- [Package metadata](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/Cargo.toml)

`pg_webrtc` 0.1.0 is an early PostgreSQL 18 preload extension that links pgrx 0.19.2 with `webrtc-rs` and establishes a background-worker process boundary. Its only SQL surface is build information:

```sql
CREATE EXTENSION pg_webrtc;
SELECT webrtc.version();
SELECT webrtc.engine();
```

The worker is registered only when PostgreSQL loads the library during server startup. To use it, add the library to `shared_preload_libraries` and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_webrtc'
```

### Current Boundary

The worker currently owns a Tokio runtime and remains idle; it does not yet expose peer-connection, offer/answer, ICE-candidate, or data-channel SQL commands. The extension is intentionally superuser-only, non-relocatable, and may eventually open network sockets. Treat it as a runtime/process-boundary prototype, not a usable WebRTC API. The upstream repository declares no license, which must be resolved before redistribution.
