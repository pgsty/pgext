## Usage

Sources:

- [PGXN biscuit 2.4.3 distribution](https://pgxn.org/dist/biscuit/2.4.3/)
- [PGXN 2.4.3 metadata](https://api.pgxn.org/dist/biscuit/2.4.3/META.json)
- [PGXN 2.4.3 control file](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/biscuit.control)
- [PGXN 2.4.3 changelog](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/CHANGELOG.md)
- [Official documentation](https://biscuit.readthedocs.io/)

`biscuit` is an experimental PostgreSQL index access method optimized for pattern filters on text. It targets selective `LIKE`, `ILIKE`, `NOT LIKE`, and `NOT ILIKE` predicates, including multi-column and expression indexes, while trading additional memory and write work for faster filtering.

### Core Workflow

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body);

SELECT id, body
FROM message
WHERE body ILIKE '%timeout%';
```

Expression indexes work when the query uses the same expression:

```sql
CREATE INDEX customer_email_biscuit_idx
ON customer USING biscuit (lower(email));

SELECT *
FROM customer
WHERE lower(email) LIKE '%@example.com';
```

For predicates spanning several indexed text columns, use a multi-column index and confirm the chosen plan with `EXPLAIN (ANALYZE, BUFFERS)` on representative data.

### Important Objects

- `biscuit` is the index access method used in `CREATE INDEX ... USING biscuit`.
- `biscuit_operators` reports the supported operators.
- `biscuit_version` and `biscuit_build_info` expose build information; `biscuit_build_info_json` returns it as JSON.
- `biscuit_status` reports the installed build and bitmap configuration.
- `biscuit_index_stats` and `biscuit_index_memory_size` inspect an index and its memory footprint.
- `biscuit_memory_usage` is a view of extension memory use.
- `biscuit_has_roaring` and `biscuit_roaring_version` report optional Roaring bitmap support.

### Limits and Operations

`biscuit` is for filtering, not ordered index scans. It does not provide regular-expression or similarity search. Indexes can be larger and more expensive to maintain than a B-tree; benchmark read selectivity, ingest cost, memory use, and vacuum behavior before production use. Keep a conventional index where ordering, equality, uniqueness, or another access method is still required.

The upstream project labels Biscuit as actively developed. PGXN publishes `2.4.3` as a stable distribution, but that archive's changelog stops at `2.4.2`, and its metadata and control file expose SQL extension version `2.4.1`. Treat `2.4.3` as a distribution/package refresh: no additional SQL API delta is claimed. The material `2.4.2` change fixes a use-after-free in the index cache plus compiler warnings.
