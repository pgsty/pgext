## Usage

Sources:

- [packages/db-utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/README.md)
- [packages/db-utils/sql/db-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/sql/db-utils--0.47.0.sql)
- [packages/db-utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/Makefile)
- [packages/db-utils/db-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/db-utils.control)

`db-utils` provides recursive JSONB utilities in a fixed schema. `db_utils.jsonb_deep_merge` recursively merges objects with the patch winning other conflicts, while `db_utils.jsonb_set_deep` creates intermediate objects when setting a nested path.

### Core Workflow

```sql
CREATE EXTENSION "db-utils";
SELECT db_utils.jsonb_deep_merge('{"a":{"x":1}}', '{"a":{"y":2}}');
SELECT db_utils.jsonb_set_deep('{}', ARRAY['a','b'], '3'::jsonb);
```

### Operational Boundaries

`db_utils.get_column_smart_comment` reads a column comment and runs as `SECURITY DEFINER`. `db_utils.timestamps` alters the named table: it adds missing timestamp columns, replaces its timestamp trigger and creates missing indexes. That helper requires DDL privileges and takes the normal DDL locks; it is not a read-only timestamp function. Dependencies are `inflection-db`, `pgpm-inflection`, `pgpm-stamps`, `pgpm-verify`, `plpgsql`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
