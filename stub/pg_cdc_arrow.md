## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow.c)

`pg_cdc_arrow` `1.0` provides arrow ipc logical-decoding output plugin. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
SELECT * FROM pg_create_logical_replication_slot('arrow_demo', 'pg_cdc_arrow');
SELECT * FROM pg_logical_slot_get_binary_changes('arrow_demo', NULL, NULL);
```

Creating and consuming replication slots requires the corresponding replication privileges.

### Operational boundaries

The installation SQL defines no SQL objects; decoding loads the output library directly. Configure logical WAL and available replication slots, then use an authorized replication role. Binary changes are Arrow IPC, not JSON or text. Monitor retained WAL and remove unused slots after consumers stop; inspect replica identity for update/delete payloads.

This module corresponds to the former upstream name `alohadb_cdc_arrow`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
