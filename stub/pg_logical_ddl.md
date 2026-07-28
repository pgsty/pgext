## Usage

Sources:

- [Official upstream README](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/README.md)
- [Official extension control file (pg_logical_ddl.control)](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/pg_logical_ddl.control)
- [Official implementation source](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/pg_logical_ddl.c)

`pg_logical_ddl` — Experimental DDL replication implementation for PostgreSQL. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Upstream explicitly says it is not production-ready.

### Core Workflow

Build PostgreSQL with the upstream-linked, in-progress apply-worker message patch. Install the library, preload it on both publisher and subscriber, and restart both servers:

```ini
shared_preload_libraries = 'pg_logical_ddl'
```

On the publisher, select the DDL command tags to capture:

```ini
pg_logical_ddl.log_command_tags = 'CREATE TABLE, ALTER TABLE, DROP TABLE'
```

Create normal logical replication objects, enabling message delivery on the subscription:

```sql
-- publisher
CREATE PUBLICATION ddl_pub FOR ALL TABLES;

-- subscriber
CREATE SUBSCRIPTION ddl_sub
  CONNECTION 'host=publisher dbname=app user=replicator'
  PUBLICATION ddl_pub
  WITH (message = true);
```

DDL executed on the publisher is written as a logical message and re-executed by the subscriber apply worker with the captured role and `search_path`.

### Important Setting

- `pg_logical_ddl.log_command_tags` is a case-insensitive comma-separated list. Its empty default disables DDL capture.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Upstream labels part or all of the project experimental.
- Upstream describes the project as a proof of concept.
- The required core patch and `message = true` subscription option are not available in a released PostgreSQL version.
- There is no loop prevention. Keep capture disabled on the subscriber and use one-directional replication.
- DDL text is replayed at statement level. Environment-dependent names, functions, permissions, and `search_path` contents can diverge.
- Only plain `CREATE TABLE` is automatically registered with the subscription in the reviewed implementation.
