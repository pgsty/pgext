## Usage

Sources:

- [README](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/pg_phasevec.control)
- [Base SQL](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/sql/pg_phasevec--0.1.0.sql)
- [0.16.0 SQL upgrade](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/sql/pg_phasevec--0.15.0--0.16.0.sql)
- [PGXN release](https://pgxn.org/dist/pg_phasevec/0.16.0/)
- [IVF implementation](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/src/phasevec_ivf.c)
- [PostgreSQL 18 DSM interface](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/include/storage/dsm_registry.h)
- [PostgreSQL 18 LWLock interface](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/include/storage/lwlock.h)

`pg_phasevec` 0.16.0 stores provenance-aware vector readings beside PostgreSQL rows. A reading records the sealed reader identity, the exact text hash and the time; stale or foreign readings are withheld from search. Upstream declares PostgreSQL 18 or newer and makes no network calls; the 0.16.0 source has the PostgreSQL 18 compatibility problem described below.

### PostgreSQL 18 Compatibility

Release 0.16.0 does not compile against PostgreSQL 18: its IVF implementation calls `LWLockNewTrancheId` and `GetNamedDSMSegment` with incompatible signatures and calls `GetNamedDSA`, which PostgreSQL 18 does not provide. Both native and LLVM builds fail. The declared minimum version therefore does not establish working PostgreSQL 18 support; use an upstream compatibility fix before following the installation workflow. Other PostgreSQL versions have not been verified here.

### Seal a Reader and Search

```sql
CREATE EXTENSION pg_phasevec;
SELECT extversion FROM pg_extension WHERE extname = 'pg_phasevec';

```

A reader is registered with its model digest and dimensionality. A watch names the source table, key and text expression; `phasevec.next` supplies rows to your local reader, and `phasevec.record` stores the resulting codes against the text hash. `phasevec.search`, `find` and `alike` query only compatible readings. `coverage` reports gaps, while `reader_switch` marks readings from the previous reader as foreign.

### Index and Boundaries

IVF maintenance is explicit through `ivf_build`, `ivf_maintain` and `ivf_absorb`. The index is approximate and coverage thresholds can refuse to answer from a partially filled watch. Editing a watched row makes its reading stale in the same transaction. The extension does not supply an embedding model, reader service, network transport or durable cross-system ordering; callers own reader sealing, privileges and model provenance. The schema is fixed to `phasevec` and the extension is not relocatable.

### Upgrade and Security

Install as a superuser. The source includes a base installation and versioned SQL upgrades through 0.16.0. Install matching library/SQL files and follow the supplied upgrade chain; 0.16.0 hardens allocation bounds, cursor NULL handling and maintenance privileges. Keep model digests, reader manifests and source-table access under the same trust boundary as the database.
