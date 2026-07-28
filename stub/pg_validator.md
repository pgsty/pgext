## Usage

Sources:

- [Official upstream README](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/README.md)
- [Official extension control file (pg_validator.control)](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/dist/pg_validator.control)
- [Official extension SQL (pg_validator--1.0.sql)](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/dist/pg_validator--1.0.sql)

`pg_validator` — PostgreSQL extension to validate data with trigger. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_validator;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `alpha_rule("value" ANYELEMENT)` is an extension function and returns `BOOLEAN`.
- `array_is_unique("arr" ANYARRAY)` is an extension function and returns `BOOLEAN`.
- `array_overlap_count("a" ANYARRAY, "b" ANYARRAY)` is an extension function and returns `INT`.
- `array_unique("arr" ANYARRAY)` is an extension function and returns `ANYARRAY`.
- `constraint_def_contained("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` is an extension function and returns `BOOLEAN`.
- `constraint_def_contains("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` is an extension function and returns `BOOLEAN`.
- `constraint_def_eq("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` is an extension function and returns `BOOLEAN`.
- `constraint_def_neq("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` is an extension function and returns `BOOLEAN`.
- `email_rule("value" ANYELEMENT)` is an extension function and returns `BOOLEAN`.
- `exists_rule("relid" REGCLASS, "table_columns" TEXT[], "record" JSONB, "record_columns" TEXT[], "mode" FK_MODE = 'full', "where" TEXT = 'TRUE')` is an extension function and returns `BOOLEAN`.
- `is_distinct_from("a" ANYELEMENT, "b" ANYELEMENT)` is an extension function and returns `BOOLEAN`.
- `is_not_distinct_from("a" ANYELEMENT, "b" ANYELEMENT)` is an extension function and returns `BOOLEAN`.
- `jsonb_array_append("json" JSONB, "path" TEXT[], "value" JSONB)` is an extension function and returns `JSONB`.
- `jsonb_except("a" JSONB, "b" JSONB)` is an extension function and returns `JSONB`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
