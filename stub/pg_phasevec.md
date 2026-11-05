## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/README.md)
- [pg_phasevec.control](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/pg_phasevec.control)
- [sql/pg_phasevec--0.1.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.1.0.sql)
- [sql/pg_phasevec--0.16.0--0.17.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.16.0--0.17.0.sql)
- [sql/pg_phasevec--0.17.0--0.18.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.17.0--0.18.0.sql)
- [src/phasevec_ivf.c](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/src/phasevec_ivf.c)

`pg_phasevec` 0.18.0 stores provenance-aware vector readings beside PostgreSQL rows. Each reading identifies its sealed reader, exact text hash and time; stale or foreign readings are withheld from search. The reader runs locally and is supplied by the operator.

### Enablement and Compatibility

```sql
CREATE EXTENSION pg_phasevec;
SELECT extversion FROM pg_extension WHERE extname = 'pg_phasevec';
```

Installation requires a superuser; the schema is fixed to `phasevec`. No other extension is required. Upstream documents PostgreSQL 18–19 testing. Version 0.17.0 fixes the PostgreSQL 18 build failure present in 0.10.0–0.16.0: PostgreSQL 18 uses a per-backend IVF cache, with `phasevec.shared_ivf_cache` off and unavailable; shared IVF requires PostgreSQL 19. This source review does not establish a locally rebuilt binary.

### Readers and Search

Register a reader with its model digest and dimensionality. A watch names the source table, key and text expression. `phasevec.next` supplies work to the reader, and `phasevec.record` stores resulting codes against the text hash. `phasevec.search`, `find` and `alike` search compatible readings; `coverage` reports gaps, while `reader_switch` marks old-reader results as foreign. The extension supplies neither an embedding model nor a durable reader service.

`ivf_build`, `ivf_maintain` and `ivf_absorb` perform explicit approximate-index maintenance. Editing a source row marks its reading stale in the same transaction. Coverage thresholds can refuse queries over an incompletely filled watch; permitting partial results is a caller decision.

### Recall and Brief Output

`phasevec.recall_approve` records an owner-approved baseline and `phasevec.recall_check` compares current recall with it. Checks use the saved search parameters and sampled keys, so a large sample can be expensive; above 100000 live readings the reference search is itself prefiltered. An unavailable baseline or unusable sample is reported as unmeasurable. These functions and their baseline table are revoked from `PUBLIC`; grant monitoring access deliberately.

Version 0.18.0 adds `phasevec.brief(watch, budget)` with a default character budget of 8000. It summarizes still-valid readings and obtains source text with the caller's privileges. It does not grant additional access to the underlying table.

### Upgrade and Trust Boundary

Install matching library and SQL files and follow the supplied upgrade chain. When crossing 0.17.0 on a server using shared IVF, restart to release the old shared-memory segment after the cache format change. The new cache also distinguishes extension installations, preventing reuse of stale state after drop/recreate.

Maintenance functions are revoked from `PUBLIC`. Reader sealing and model digests establish provenance, not model quality; keep reader manifests, maintenance privileges and source-table permissions under the database's operational controls.
