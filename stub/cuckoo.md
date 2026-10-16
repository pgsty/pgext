## Usage

Sources:

- [README.md](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/README.md)
- [cuckoo--1.0.sql](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/cuckoo--1.0.sql)
- [LICENSE.md](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/LICENSE.md)
- [cuckoo.control](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/cuckoo.control)

`cuckoo` 1.0 adds a probabilistic index access method for equality predicates. The project is named pg_cuckoo, but the SQL extension and access method are both `cuckoo`.

### Core Workflow

```sql
CREATE EXTENSION cuckoo;
CREATE TABLE cuckoo_example (label text);
CREATE INDEX ON cuckoo_example USING cuckoo (label);
INSERT INTO cuckoo_example VALUES ('alpha');
SELECT * FROM cuckoo_example WHERE label = 'alpha';
```

### Operational Boundaries

Upstream requires PostgreSQL 16 or later; the control defaults require superuser installation. The index supports equality, not ranges, and heap rechecks remove false positives. Single-column indexes are the simplest use: a multicolumn index combines fingerprints and needs predicates on all indexed columns. Options `bits_per_tag`, `tags_per_bucket` and `max_kicks` default to 12, 4 and 500. More fingerprint bits trade space for a lower false-positive rate; benchmark claims are not workload guarantees.
