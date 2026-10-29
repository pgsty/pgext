## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/README.md)
- [pg_tre.control](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/pg_tre.control)
- [sql/pg_tre--4.5.3.sql](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/sql/pg_tre--4.5.3.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/CHANGELOG.md)
- [LIMITATIONS.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/LIMITATIONS.md)
- [META.json](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/META.json)
- [.github/workflows/ci.yml](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/.github/workflows/ci.yml)
- [Library initialization](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/src/module.c)

`pg_tre` 4.5.3 provides the `tre` native index access method for approximate regular-expression matching over text. Upstream supports PostgreSQL 17–19; Pigsty's 4.5.2 packages target PostgreSQL 17–18. It narrows candidates with codepoint trigrams and postings, then checks matches against the heap. Use it for typo-tolerant identifiers, log messages and lexical patterns with explicit edit costs.

### Enable and Search

The index write path uses PostgreSQL's generic WAL facility. Since 4.0.0, `CREATE EXTENSION` is sufficient: `shared_preload_libraries` and a restart are not required to enable a new installation. Extension creation requires a superuser, and the control file is untrusted and non-relocatable.

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
SELECT id, body, body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

PostgreSQL 19 uses the core regular-expression character classifier; ICU character-class edge cases on PostgreSQL 17–18 may differ from native PostgreSQL regex semantics, as described upstream.

### Patterns and Objects

- `tre_pattern(text)` uses `pg_tre.default_max_cost`; `tre_pattern(text, int4)` supplies a maximum cost; the five-argument form also sets insertion, deletion and substitution costs. Patterns use POSIX extended regular expressions and TRE approximate budgets.
- `text %~~ tre_pattern` is the indexable approximate-match predicate; `text <@> tre_pattern` returns the edit cost or NULL when no match fits the budget.
- `tre_distance`, `tre_similarity`, `tre_amatch`, `tre_amatch_cost` and `tre_amatch_detail` expose matching and compatibility functions.
- The index operator family also supports LIKE, ILIKE, regular-expression and equality predicates with heap rechecks. The separately named trigram functions coexist with `pg_trgm`; `pg_tre` does not install the colliding bare trigram operators removed in 3.0.

Indexes accept one text column with `tre_text_ops`. Expression and partial indexes are available, but index-only scans are not. Short patterns, high edit budgets and non-literal regexes can produce broad candidate sets; inspect real plans and keep cancellation and query-time limits enabled.

### Maintenance and Limits

`fastupdate` must be true; setting it false is now rejected because the direct path could not preserve the index format. Existing off settings can still be opened, but inserts use the pending list. `pending_list_limit` is in KiB and defaults to 0, meaning unlimited. A positive value triggers a synchronous merge when the threshold is crossed; that inserting backend pays the work, and if another merge is active the limit can temporarily be exceeded.

`pg_tre.min_trigram_freq` must remain 1. Values above 1 formerly omitted rare postings and could lose results; they are now rejected. Keep regex state and compilation/matching limits bounded for untrusted patterns. Index construction uses tuplesort and temporary disk, with memory governed by maintenance_work_mem; estimate large builds before scheduling them.

```sql
SELECT * FROM tre_estimate_index_build('documents'::regclass, 2);
SELECT * FROM pg_tre_index_pending('documents_body_tre'::regclass);
SELECT * FROM pg_tre_index_check('documents_body_tre'::regclass, 2);
```

`pg_tre_index_check` compares index answers with sequential scans. Each sample performs multiple sequential scans, so bound the sample count on large tables. Version 4.5.1 fixes its search-path handling, pins trusted lookup paths and qualifies calls; use the fixed version before running this helper as a privileged user.

### Upgrade Boundaries

Install matching files, ensure sessions use the new library, and update each database:

```sql
ALTER EXTENSION pg_tre UPDATE TO '4.5.3';
```

Crossing from 3.x to 4.x changes WAL record types even though the index page format is unchanged. Shut down cleanly before replacing the library; a physical standby cannot run a different major extension version from its primary and must be rebuilt from a fresh base backup after the primary upgrade. There is no rolling upgrade across this boundary. Remove any old `wal_consistency_checking = 'pg_tre'` setting; it prevents startup in 4.x. Upstream recommends `wal_consistency_checking = 'all'` when consistency checking is needed.

Version 4.4 writes format 11 and reads formats 6–10. Upgrading from those formats does not require a routine rebuild; optional REINDEX obtains the newer dense-posting representation. Downgrading to 4.3 after a 4.4+ write requires rebuilding the affected indexes for the older format. Older upgrades may cross additional boundaries described in the changelog.

REINDEX every index built with min_trigram_freq greater than 1: installing the fix cannot restore omitted postings. Version 4.5 also disables stale SuRF acceleration on indexes from earlier affected versions, restoring correctness immediately; a rebuild is optional to regain that acceleration. Upgrading 4.5.0 to 4.5.1 changes neither format nor ordinary query behavior and requires no additional REINDEX. Old definitions using the removed tuple bloom options must still be corrected before rebuilding.

### 4.5.2 Regex Correction

Version 4.5.2 fixes newline-mode `(?n)` anchors and negated character classes. It retains the index format and needs no additional REINDEX when upgrading from 4.5.1. The older WAL-transition and omitted-posting rebuild requirements above still apply when crossing those affected versions.

```sql
SELECT E'a\nfoo' %~~ tre_pattern('(?n)^foo', 0) AS line_start,
       E'foo\n' %~~ tre_pattern('(?n)foo$', 0) AS line_end,
       E'\n' %~~ tre_pattern('(?n)[^x]', 0) AS negated_class;
```

The expected results are true, true and false. The fix affects matching and index rechecks alike; patterns without newline mode retain their previous behavior.

### 4.5.3 Security Fix

An approximate pattern with many reachable states could overflow TRE's fixed work list and crash a backend, causing the postmaster to restart the cluster's other backends. Version 4.5.3 grows that work list within the existing pattern limits. Install the matching library, replace sessions using the old library, and update extension SQL to 4.5.3. The patch changes no on-disk format and adds no REINDEX requirement; the older migration and damaged-index boundaries above remain.
