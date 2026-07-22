## Usage

Sources:

- [Official README](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/README.md)
- [Official control file](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/pg_audit.control)
- [Official version 0.1.3 SQL](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/sql/pg_audit--0.1.3.sql)

`pg_audit` from Disqus is a SQL generator for per-table audit tables and row triggers. It is not the similarly named pgaudit server-audit extension: it records row images in ordinary PostgreSQL tables rather than producing statement audit logs.

### Core Workflow

After installation, `create_audit()` returns DDL as text. Capture that text, review it, and execute it as a separate transaction:

```sh
psql -AtX -o actually_create_audit.sql \
  -c 'SELECT create_audit();' -U app_owner appdb

psql -X -1f actually_create_audit.sql -U app_owner appdb
```

Calling `create_audit(true)` additionally generates triggers enabled for `replica`; the no-argument form uses the normal trigger setting.

### Generated Objects

For eligible application schemas, the generator emits corresponding `<schema>_audit` schemas and audit tables. Each audit row contains pairs of old/new source values plus metadata such as the operation, timestamp, current user, and session user. It also emits row triggers, supporting indexes, and initial `INSERT` snapshots of existing rows.

### API

- `create_audit() RETURNS SETOF text`: generate the default audit DDL.
- `create_audit(enable_replica boolean) RETURNS SETOF text`: control whether generated triggers also fire in replica mode.

### Caveats

Always inspect the generated SQL before execution, as the official workflow requires. The generator excludes PostgreSQL catalog schemas and audit schemas, but it operates broadly across other tables visible to the installing role. The extension requires PL/pgSQL and enough ownership privileges to create schemas, tables, indexes, and triggers.

Version `0.1.3` does not provide a complete schema-evolution workflow; its SQL explicitly leaves source columns added after audit setup unhandled. Audit rows are ordinary mutable data, so define access control, retention, and tamper-evidence separately.
