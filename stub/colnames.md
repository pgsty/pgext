## Usage

Sources:

- [Release documentation](https://github.com/theory/colnames/blob/v1.7.1/doc/colnames.md)
- [SQL function](https://github.com/theory/colnames/blob/v1.7.1/sql/colnames.sql)
- [C implementation](https://github.com/theory/colnames/blob/v1.7.1/src/colnames.c)
- [Extension control](https://github.com/theory/colnames/blob/v1.7.1/colnames.control)
- [Release notes](https://github.com/theory/colnames/blob/v1.7.1/Changes)

`colnames` provides `colnames(record)`, returning the input record's field names as a `name[]`. Use it in generic triggers or dynamic SQL helpers that need a composite value's structure without converting its values to JSON. Distribution version `1.7.1` retains extension version `1.7.0`.

### Core Workflow

```sql
CREATE EXTENSION colnames;

SELECT colnames(ROW(1, 'foo', 458.0));
SELECT colnames(t)
FROM (SELECT 1 AS id, 'Ada'::text AS name) AS t;

CREATE TYPE contact AS (id integer, name text);
SELECT colnames(NULL::contact);
```

The three calls return `{f1,f2,f3}`, `{id,name}`, and `{id,name}` respectively. Anonymous rows receive generated names; named rows preserve their attribute names. A typed null works because its composite type supplies the descriptor. An untyped null does not identify a composite type.

### Function Behavior

`colnames(record)` returns `name[]`, is `STABLE`, and is deliberately not `STRICT`. It preserves column order and quoted identifier spelling, omits dropped columns, and returns an empty array for an empty composite type. It reads the row descriptor rather than table data.

### Operational Notes

The extension has no runtime extension dependencies and needs no preload or restart. Its C library loads when the function is called. Creating the extension requires a superuser because the control file does not mark it trusted. It is relocatable and can be installed in a chosen schema or moved with `ALTER EXTENSION colnames SET SCHEMA`.

Release `1.7.1` documents PostgreSQL 8.2–17 compatibility. The PGSTY package targets PostgreSQL 14–18; PostgreSQL 18 acceptance comes from package and regression testing rather than that release's compatibility statement.

PostgreSQL 9.3 and later also offer `row_to_json` with `json_object_keys` as a pure SQL alternative. If discovered names are used in dynamic SQL, quote identifiers and keep values parameterized.
