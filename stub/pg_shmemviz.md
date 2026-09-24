## Usage

Sources:

- [README.md](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/README.md)
- [extension/pg_shmemviz.control](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/extension/pg_shmemviz.control)
- [extension/pg_shmemviz--0.1.sql](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/extension/pg_shmemviz--0.1.sql)
- [LICENSE](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/LICENSE)

`pg_shmemviz` captures PostgreSQL shared memory for an offline browser viewer. Its extension version is 0.1; the README identifies the diagnostic release as 0.1.0-beta.1 and validates PostgreSQL 20devel. Use only a disposable or isolated development instance.

### Capture and View

Use the exact server development files and an unstripped server executable containing DWARF debug information. The Python CLI needs LLDB with Python support on macOS or GDB with Python support elsewhere, plus local superuser database access.

```sql
CREATE EXTENSION pg_shmemviz;
```

```sh
pg_shmemviz capture --pg-config /opt/postgresql/bin/pg_config --dbname postgres /tmp/pg-memory-snapshot
pg_shmemviz serve /tmp/pg-memory-snapshot
```

### Objects and Boundaries

`pg_shmemviz_snapshot(text)` and `pg_shmemviz_snapshot(text, text)` write snapshots to a server-side path; public execution is revoked. The CLI combines capture with executable layout metadata. `--numa-scope` selects NUMA inspection scope; `pg_buffercache` and `--buffercache-details` add optional buffer information. The viewer binds to loopback by default.

No preload is needed. Capture copies live memory sequentially without locking: related fields can represent different instants and multi-byte values can be torn. Snapshots may contain sensitive data and occupy substantial disk space. NUMA inspection can fault pages and affect placement. The debugger reads the executable’s metadata rather than attaching to the server. The viewer does not modify captured server memory.
