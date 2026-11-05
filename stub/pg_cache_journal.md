## Usage

Sources:

- [Upstream source 1](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/README.md)
- [Upstream source 2](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/crates/pg-cache/pg_cache_journal.control)
- [Upstream source 3](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/Makefile)

`pg_cache_journal` 0.1.0 adds the optional WAL-logged journal used by Flaro's write-behind persistence prototype.

### Core Workflow

Use artifacts built with the `journal-prototype` feature, then create both extensions in the cache database. The parent library needs preload and restart. Batches enter `cache_journal.segment0` and `cache_journal.segment1`; a separate worker materializes them into `cache_journal.entries`.

```ini
pg_cache.journal = on
```

```sql
CREATE EXTENSION pg_cache;
CREATE EXTENSION pg_cache_journal;
SELECT * FROM cache_journal.state;
```

### Operational Boundaries

Requires PostgreSQL 17 and pg_cache. Journal commit and materialization are different stages; reading the materialized table alone can lag committed journal writes. RESP acknowledgments remain asynchronous with respect to persistence. Capacity is bounded and sustained write pressure can exhaust it; do not treat this experimental mode as synchronous Redis durability.
