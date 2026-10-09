## Usage

Sources:

- [packages/inflection-db/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/README.md)
- [packages/inflection-db/sql/inflection-db--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/sql/inflection-db--0.47.0.sql)
- [packages/inflection-db/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/Makefile)
- [packages/inflection-db/inflection-db.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/inflection-db.control)

`inflection-db` builds database naming helpers on PGPM inflection. `inflection_db.get_table_name` pluralizes names, and helpers generate field, index, constraint, schema and namespace names.

### Core Workflow

```sql
CREATE EXTENSION "inflection-db";
SELECT inflection_db.get_table_name('BlogPost');
SELECT inflection_db.get_index_name('BlogPost', ARRAY['created_at']);
```

### Operational Boundaries

`inflection_db.get_identifier` truncates to 63 characters in the pinned SQL. This is not a uniqueness guarantee, and multibyte identifiers still need PostgreSQL byte-length checks. Generated names are text, not executed DDL; use identifier quoting when constructing SQL. `inflection_db.get_index_name` and `inflection_db.get_unique_index_name` accept arrays of fields. Dependencies are `pgpm-inflection`, `pgpm-verify`, `plpgsql`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
