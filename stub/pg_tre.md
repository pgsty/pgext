## Usage

Sources:

- [pg_tre v3.0.2 README](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/README.md)
- [pg_tre v3.0.2 control file](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/pg_tre.control)
- [pg_tre v3.0.2 extension SQL](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/sql/pg_tre--3.0.2.sql)
- [pg_tre v3.0.2 changelog](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/CHANGELOG.md)
- [Production sizing and limitations](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/LIMITATIONS.md)
- [pg_tre v3.0.2 PGXN metadata](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/META.json)
- [PG17 and PG18 CI matrix](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/.github/workflows/ci.yml)
- [Official regression tests](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/test/sql)

`pg_tre` 3.0.2 provides the `tre` native PostgreSQL index access method for approximate regular-expression matching over `text`. It combines codepoint trigrams and compressed postings to narrow candidates, then uses the TRE regex engine for an authoritative heap recheck. Use it for typo-tolerant identifiers, log messages, error codes, or other lexical patterns that need true insertion, deletion, and substitution costs; it is not a replacement for linguistic full-text search or semantic-vector search.

### Enable the Extension

The index write path uses a custom WAL resource manager, so preload the library and perform a full PostgreSQL restart before creating or changing a `tre` index. A configuration reload is not sufficient.

```ini
shared_preload_libraries = 'pg_tre'
```

Then install the extension in each required database. The v3.0.2 control file sets `superuser=true`, `trusted=false`, and `relocatable=false`, so creation requires a superuser and the extension cannot be relocated.

```sql
CREATE EXTENSION pg_tre;

CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body text NOT NULL
);

INSERT INTO documents (body) VALUES
  ('The PostgreSQL database system'),
  ('The Postgres databse system');

CREATE INDEX documents_body_tre ON documents USING tre (body);

SELECT id, body,
       body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

Loading on demand without preload leaves legacy functions such as `tre_amatch` and `tre_version` usable, but index mutations fail because the custom resource manager was not registered. Add `pg_tre` to any existing comma-separated preload list instead of replacing the other libraries, then restart.

The tagged v3.0.2 build metadata requires PostgreSQL 17 or newer, the Nix flake exports separate `pg17` and `pg18` builds, and the regression CI matrix runs on PostgreSQL 17 and 18. Those two majors are the evidenced support range for this release. The README and user-guide headlines still say PostgreSQL 18+, which is stale relative to the tagged build and test configuration.

### Patterns, Operators, and Functions

`tre_pattern(text)` creates a pattern using `pg_tre.default_max_cost`; `tre_pattern(text, int4)` supplies an explicit maximum cost; and `tre_pattern(text, int4, int4, int4, int4)` additionally sets insertion, deletion, and substitution costs. The pattern language is POSIX extended regular expressions plus TRE's local approximate budget syntax such as `{~1}`.

- `text %~~ tre_pattern` returns a boolean approximate-regex match and is the main indexable predicate.
- `text <@> tre_pattern` returns the best edit cost, or `NULL` when no alignment fits the pattern budget. The `tre` operator family supports ascending index ordering for top-N queries; still include a selective `%~~` predicate when appropriate and use a deterministic tie breaker.
- `tre_distance` and `tre_similarity` expose distance and normalized similarity as functions. `tre_amatch`, `tre_amatch_cost`, and `tre_amatch_detail` are compatibility functions that call TRE directly and do not use the index.
- `LIKE`, `ILIKE`, `~`, `~*`, and `=` are bound to the v3.0.2 operator family as lossy candidate filters. PostgreSQL always rechecks their native semantics against the heap.
- `tre_trgm_similarity`, `tre_trgm_distance`, `tre_word_similarity`, and `tre_strict_word_similarity` provide pg_trgm-like calculations under distinct names. Since 3.0.0, `pg_tre` deliberately does not create the bare `%`, `<->`, `<%`, `<<->`, `<<%`, or `<<<->` operators, so it can coexist with `pg_trgm`.

`tre` indexes are single-column, accept the `text` opclass `tre_text_ops`, and are lossy: they do not provide index-only scans, and a TRE recheck decides every final result. Expression and partial indexes remain available through normal PostgreSQL facilities.

### Selectivity, Costs, and Safety Limits

Index acceleration depends on useful literal trigrams. A short pattern, a pattern dominated by non-literal regex constructs, a large edit budget, or extraction fanout beyond `pg_tre.max_extraction_fanout=4096` can fall back to a fully lossy candidate bitmap. Correctness is preserved, but the planner may reasonably prefer a sequential scan. Use `EXPLAIN (ANALYZE, BUFFERS)` on representative data; production guidance favors edit budgets of 0-2, requires workload testing at 3, and discourages values above 3 for normal text because TRE recheck cost rises quickly.

Important safety defaults are `pg_tre.max_nfa_states=10000`, `pg_tre.compile_timeout_ms=1000`, and `pg_tre.match_timeout_ms=1000`. Keep them bounded for untrusted patterns. Version 3.0.1 added cooperative interrupt checks throughout `%~~` candidate construction and recheck, so ordinary `statement_timeout` and client cancellation can stop broad scans in v3.0.2.

The supported per-index storage parameters are `fastupdate`, `pending_list_limit`, `range_size_blocks`, and `q`; `q` must remain 3. Their defaults are `fastupdate=true`, `pending_list_limit=4096` KiB, and `range_size_blocks=128`. With fast update enabled, inserts enter a pending list that maintenance later merges; monitor and vacuum write-heavy indexes instead of allowing an indefinitely growing pending list.

Version 3.0.0 removed the per-tuple positional bloom payload and the `tuple_bloom_enable` and `bloom_tuple_bits` GUCs and index options. Old index definitions that still specify either option fail and must be updated. The range bloom is still built, but the v3.0.2 changelog records that scans do not yet probe it; current filtering should therefore be understood as posting-based candidate reduction followed by mandatory TRE recheck, not the three active tiers still shown in older README prose.

### Build, Maintenance, and Upgrade Boundaries

Index builds use PostgreSQL tuplesort, with memory bounded principally by `maintenance_work_mem`; temporary-disk consumption still grows with emitted trigram tuples. Estimate a large build first and prefer concurrent operations on live tables:

```sql
SELECT *
FROM tre_estimate_index_build('documents'::regclass, 2);

CREATE INDEX CONCURRENTLY documents_body_tre_live
ON documents USING tre (body);

REINDEX INDEX CONCURRENTLY documents_body_tre_live;
```

`tre_estimate_index_build` samples at most 2000 rows and reports estimated rows, trigram tuples, temporary disk, and final index size. `pg_tre.build_max_entries_mb=0` leaves the temporary-disk guard disabled; set a measured nonzero ceiling when the temp tablespace is constrained. `pg_tre.min_trigram_freq=1` preserves every posting by default; raising it can reduce index size but increases lossy fallback work, so validate query plans and recall-equivalent correctness against a sequential-scan baseline.

Version 3.0.2 is a packaging and documentation release: it changes neither C behavior, SQL surface, WAL, nor on-disk format. Upgrading from 3.0.1 is a version-only `ALTER EXTENSION pg_tre UPDATE TO '3.0.2'` and does not require `REINDEX`. For older jumps, read every intervening upgrade note before changing the extension; format boundaries before 1.6.0 can require rebuilding indexes, while the 3.0.0 v9 format remains backward-readable for v6-v8 indexes.
