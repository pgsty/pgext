## Usage

Sources:

- [README.md](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/README.md)
- [yezzey.control](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey.control)
- [yezzey--2.0.sql](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey--2.0.sql)
- [yezzey.c](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey.c)
- [docs/README.cleanup.md](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/docs/README.cleanup.md)

`yezzey` 2.0 offloads append-only row/column tables to S3 while preserving SQL access. It requires a matching patched OpenGPDB (Greenplum 6) or Apache Cloudberry build and YProxy; it is not a stock PostgreSQL extension. The project release is 2.0.0, while the control version is 2.0.

### Core Workflow

```ini
shared_preload_libraries = 'yezzey'
```

```sql
CREATE EXTENSION yezzey;
CREATE TABLE offload_demo (id integer)
  WITH (appendonly=true, orientation=column) DISTRIBUTED RANDOMLY;
INSERT INTO offload_demo VALUES (1);
SELECT yezzey.offload_relation('offload_demo'::regclass);
SELECT * FROM yezzey.offload_relation_status('offload_demo'::regclass);
SELECT yezzey.load_relation('offload_demo'::regclass);
```

### Operational Boundaries

Preload `yezzey` on the database cluster and restart before creating the extension; configure compatible YProxy endpoints and object storage first. The OpenGPDB branch is OPENGPDB_STABLE and the Cloudberry implementation is on master. SQL installation is marked trusted, but cluster configuration, object-storage access and table ownership remain administrative boundaries.

`yezzey.offload_relation` takes an exclusive relation lock while uploading data, and `yezzey.load_relation` restores local storage. `yezzey.offload_relation_status` reports per-segment byte counts; `yezzey.relation_describe_external_storage_structure` lists external files. Version 2.0 moves and renames the earlier unqualified interfaces: update callers against its SQL definitions and rehearse migration with a restorable backup.

Cleanup uses `yezzey.vacuum`, `yezzey.vacuum_tablespace` or `yezzey.vacuum_relation`. Keep `confirm` false to inspect requests before deletion; `crazyDrop` is an aggressive superuser-only mode. External objects may still be required by retained backups. Preserve their retention boundary and access from recovery/standby systems; SQL VACUUM alone is not permission to delete arbitrary bucket objects.
