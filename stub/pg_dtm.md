## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/README.md)
- [PostgreSQL patch](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/xtm.patch)
- [External transaction-manager protocol](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/dtmd/README.md)

`pg_dtm` is a historical distributed transaction manager based on snapshot sharing. It is not a standalone extension: upstream requires patching and rebuilding PostgreSQL, running the external `dtmd` service, loading `pg_dtm` through `shared_preload_libraries`, and deploying at least two database instances.

```sql
CREATE EXTENSION pg_dtm;
SELECT dtm_begin_transaction();
-- Pass the returned global transaction ID to another node.
SELECT dtm_join_transaction(42);
```

The participating sessions then begin, modify, and commit local transactions; one commit can block until the others complete. Correctness depends on the patched server, the coordinator, every participant, and the application's global-ID handoff.

Do not load this extension into stock or current PostgreSQL. The reviewed design is research-era software with invasive server patches, an external coordinator, and no current failure or compatibility contract. Prefer maintained distributed transaction facilities. Any archival evaluation must prove atomicity under coordinator crash, backend crash, network partition, timeout, duplicate join, participant loss, failover, and recovery, and must define how in-doubt transactions and coordinator state are backed up and repaired.
