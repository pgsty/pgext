## Usage

Sources:

- [Official README](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/README.md)
- [Official extension SQL](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw--1.0.sql)
- [Official FDW implementation](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw.c)

`orc_fdw` 1.0 is a read-only foreign data wrapper for querying local files in Apache ORC format. A foreign table maps PostgreSQL columns to the ORC file and requires a server-side `filename` option.

### Core Workflow

```sql
CREATE EXTENSION orc_fdw;

CREATE SERVER orc_files
  FOREIGN DATA WRAPPER orc_fdw;

CREATE FOREIGN TABLE orc_events (
  event_id bigint,
  event_name text
)
SERVER orc_files
OPTIONS (filename '/srv/postgresql/orc/events.orc');

SELECT event_id, event_name
FROM orc_events
WHERE event_id >= 1000;
```

The path is resolved on the database server, not on the client. PostgreSQL's operating-system account must be able to open and stat the file.

### Important Objects

- `orc_fdw_handler()` supplies the PostgreSQL FDW callback table.
- `orc_fdw_validator(text[], oid)` validates object options and requires `filename` on each foreign table.
- `orc_fdw` is the installed foreign data wrapper.
- `filename` is the only implemented foreign-table option in the pinned source.

### Operational Notes

The implementation registers scan and explain callbacks but no insert, update, or delete callbacks, so use it only for reads. Files are local to each server process; failover nodes need the same path and content if queries must behave identically. The upstream build instructions target protobuf/protobuf-c and ORC data associated with Hive 0.12, indicating an old and narrow format implementation. Validate column-type mapping, compression, predicate behavior, and files produced by the exact ORC writer before adoption; do not assume current ORC ecosystem compatibility from the format name alone.
