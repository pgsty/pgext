## Usage

Sources:

- [README.md](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/README.md)
- [extension/skiplist--1.0.sql](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/extension/skiplist--1.0.sql)
- [db/init_skiplist.sql](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/db/init_skiplist.sql)
- [extension/skiplist.control](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/extension/skiplist.control)

`skiplist` 1.0 is the C index extension in the Knoo PostgreSQL 18 course project. It is an educational implementation with documented correctness and durability limitations.

### Core Workflow

```sql
CREATE EXTENSION skiplist;
CREATE TABLE skiplist_example (k integer);
INSERT INTO skiplist_example VALUES (1), (2);
CREATE INDEX skiplist_example_idx ON skiplist_example USING skiplist (k);
SELECT skiplist_dump('skiplist_example_idx'::regclass);
```

### Operational Boundaries

Use only disposable demonstrations with unique `int4` keys. The upstream documents single-column equality lookup, at most one result per duplicated key, no range/ordering support, one 8 KB page per key, and only four levels. It lacks its own WAL, complete concurrency handling and specialized VACUUM. The SQL operator-class declaration includes range operators, but this does not establish working range scans. `skiplist_dump(regclass)` prints the structure for debugging. Installation defaults to superuser-only; no license is declared.
