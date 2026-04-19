## Usage

Sources: [official repo](https://codeberg.org/kop/pg_isok), [official docs home](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/index.html), [official reference source](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/isok.xml)

`pg_isok` is a query-based data integrity and monitoring extension. Instead of only reporting rows that currently look questionable, it stores prior results and focuses later runs on unresolved or undeferred changes.

```sql
CREATE SCHEMA isok;
CREATE EXTENSION pg_isok SCHEMA isok;

SELECT *
FROM isok.run_isok_queries()
AS problems;
```

### Core Objects

- `ISOK_QUERIES` stores the monitoring queries and their execution settings.
- `ISOK_RESULTS` stores the reported rows, including whether they were resolved or deferred.
- `run_isok_queries()` runs every active check.
- `run_isok_queries($$VALUES ('check_name')$$)` runs only selected checks.

### Typical Workflow

Run one named check:

```sql
SELECT *
FROM isok.run_isok_queries($$VALUES ('new_countries')$$)
AS problems;
```

Accept or postpone a known warning by updating `ISOK_RESULTS`:

```sql
UPDATE isok.isok_results
SET deferred_to = 'infinity'
WHERE iqname = 'new_countries';
```

Use `resolved` when the condition is no longer a concern, or `deferred_to` when it should stay hidden until a later date.

### Where It Fits

- data cleanup after imports
- monitoring unusual but sometimes acceptable patterns
- "soft trigger" style review workflows where hard constraints are too strict

### Caveats

- Upstream recommends installing it in a dedicated schema and qualifying calls accordingly.
- The docs describe it as pure SQL, which is useful on managed PostgreSQL services where C extensions may be restricted.
- The package metadata in this repo says `superuser=false`, but this is not documented upstream as a trusted extension; treat installation privileges conservatively.
