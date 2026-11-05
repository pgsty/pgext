## Usage

Sources:

- [Upstream source 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [Upstream source 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion_citext.control)

`pg_lion_citext` 0.3.0 supplies case-insensitive equality support for citext values in lion indexes.

### Core Workflow

Install `citext` and `pg_lion` first. The companion registers the operator class; applications continue using normal equality predicates on the column.

```sql
CREATE EXTENSION citext;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_citext;
CREATE TABLE lion_contacts (email citext);
INSERT INTO lion_contacts VALUES ('Ada@example.com');
CREATE INDEX ON lion_contacts USING lion (email);
SELECT * FROM lion_contacts WHERE email = 'ada@example.com';
```

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Use a matching pg_lion distribution. Case folding follows citext semantics and collation, not arbitrary Unicode equivalence. The parent index's optional preload/custom-WAL requirement and VACUUM behavior still apply.
