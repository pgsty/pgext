## Usage

Sources:

- [PostgreSQL 9.6 tsearch2 control file](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2.control)
- [PostgreSQL 9.6 compatibility SQL](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2--1.0.sql)
- [PostgreSQL 9.6 compatibility implementation](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2.c)
- [Amazon RDS for PostgreSQL extension history](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

`tsearch2` version `1.0` is a backward-compatibility package for the full-text-search API that existed before PostgreSQL 8.3. It creates compatibility domains for `tsvector` and `tsquery`, old parser and dictionary entry points, legacy rank/headline wrappers, and the historical `tsearch2()` update trigger. It is intended only to help migrate old applications.

### Legacy Compatibility

On an old server or Amazon RDS engine that still exposes the package:

```sql
CREATE EXTENSION tsearch2;

SELECT to_tsvector('english', 'fat cats ate rats');
SELECT to_tsquery('english', 'cat & rat');
```

These text-argument overloads adapt old calls to PostgreSQL's built-in text-search implementation. New code should use the current `regconfig`-based full-text-search functions directly.

### Caveats

- PostgreSQL describes this extension as compatibility for the pre-8.3 API, not as a separate search engine. It is non-relocatable and introduces many old names that can shadow or confuse modern built-in functions.
- Amazon RDS marks `tsearch2` deprecated for PostgreSQL 10 and says it was removed beginning with RDS for PostgreSQL 11.1. Current RDS releases must use built-in full-text search.
- Do not add new dependencies on the compatibility domains, mutable current-parser/current-dictionary state, or the `tsearch2()` trigger. Migrate schemas and SQL to modern text-search configurations and `tsvector_update_trigger` alternatives.
- Version `1.0` reflects the frozen compatibility package. Availability on a managed service is determined by the exact historical engine version, not by this catalog entry.
