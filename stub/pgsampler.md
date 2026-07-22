## Usage

Sources:

- [Official upstream documentation](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/README.md)
- [Official worker and GUC implementation](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/pgsampler.c)
- [Official output loop implementation](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/lib/control_loop.c)

`pgsampler` is an experimental PostgreSQL background worker that periodically reads cluster activity and statistics and writes samples to CSV files or sends them to a network receiver. It exposes no meaningful SQL extension objects; activation is entirely through preload configuration and a server restart.

### Configure and Start

Choose an output mode before adding the worker to `shared_preload_libraries`:

```conf
shared_preload_libraries = 'pgsampler'
pgsampler.target_db = 'postgres'
pgsampler.output_mode = 'csv'
pgsampler.output_csv_dir = '/var/lib/postgresql/pgsampler'
pgsampler.cycle_db_seconds = 500
```

Restart PostgreSQL after installing the matching shared library. For network output, set `pgsampler.output_mode` to `network`, configure `pgsampler.output_network_host`, and set `pgsampler.token`; the source connects to TCP port 24831. There is no required `CREATE EXTENSION` workflow because the versioned SQL file does not create user-facing objects.

### Collected Data and Intervals

The worker reads activity, `pg_stat_statements`, database, table, index, column, background-writer, filesystem, replication, system, and GUC information. Main interval settings include `pgsampler.heartbeat_seconds`, `pgsampler.system_seconds`, `pgsampler.relation_seconds`, `pgsampler.bgwriter_seconds`, `pgsampler.guc_seconds`, `pgsampler.activity_seconds`, `pgsampler.replication_seconds`, and `pgsampler.statements_seconds`. CSV mode appends files below the configured directory; network mode sends the sampled records to the receiver.

### Security and Reliability Boundary

Samples can contain query text, schema details, configuration, replication state, and host information. Protect the CSV directory with PostgreSQL-owner permissions and provide rotation, retention, and disk-space monitoring. The network protocol in the reviewed source has no built-in TLS and uses a token in its handshake, so keep it on a trusted network or place it behind an authenticated encrypted tunnel; restrict outbound connectivity to the intended receiver.

Upstream calls the worker experimental and recommends it only for casual clusters. It runs with background-worker access to shared memory and all databases and was developed against PostgreSQL 9.3-era interfaces. Validate the exact server ABI, statistics overhead, restart behavior, receiver outages, file errors, and removal procedure in a disposable cluster. Removing it from `shared_preload_libraries` also requires a restart.
