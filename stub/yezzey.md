## Usage

Sources:

- [README](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/README.md)
- [Control file](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey.control)
- [yezzey.c](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey.c)
- [yezzey--1.8.8.sql](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey--1.8.8.sql)
- [yezzey--1.8.8--1.8.11.sql](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey--1.8.8--1.8.11.sql)

`yezzey` moves append-only table data to S3 while keeping it queryable through Greenplum or Apache Cloudberry. Version 1.8.11 requires a compatible patched kernel and YProxy; it is not a stock PostgreSQL extension.

### Enablement

Deploy the matching kernel and YProxy/S3 configuration on the cluster, add `yezzey` to `shared_preload_libraries`, and restart. The initialization code rejects loading outside server startup. The `yezzey.yproxy_socket` setting selects the proxy socket. Install extension SQL in the intended database.

```sql
CREATE EXTENSION yezzey;
CREATE TABLE archive_events (id integer, payload text)
  WITH (appendonly=true, orientation=column) DISTRIBUTED RANDOMLY;
INSERT INTO archive_events VALUES (1, 'example');
SELECT yezzey_define_offload_policy('public', 'archive_events');
SELECT * FROM yezzey_offload_relation_status('archive_events');
SELECT * FROM archive_events;
```

### Objects and Maintenance

`yezzey_define_offload_policy` configures offloading; `yezzey_load_relation` restores data to local storage. `yezzey_offload_relation_status` reports external size, and `yezzey_relation_describe_external_storage_structure` exposes external file layout. Only AO/AOCO tables are in scope.

Offload and reload operations may acquire strong relation locks. Reads depend on YProxy and object-store availability. Backups and recovery must include the referenced S3 objects as well as database metadata. Review garbage-collection behavior and perform a dry run before deleting external objects; destructive cleanup modes require elevated privileges. The control version is reached through the shipped 1.8.8 installation and 1.8.11 upgrade SQL.
