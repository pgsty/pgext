## Usage

Sources:

- [Official upstream README](https://github.com/nick-ivanov-edb/pg_describe_resultset/blob/0a0778b671706961226ca06161944dfde2b902b6/README.md)
- [Official extension control file (describe_resultset.control)](https://github.com/nick-ivanov-edb/pg_describe_resultset/blob/0a0778b671706961226ca06161944dfde2b902b6/describe_resultset.control)
- [Official extension SQL (describe_resultset--1.0.0.sql)](https://github.com/nick-ivanov-edb/pg_describe_resultset/blob/0a0778b671706961226ca06161944dfde2b902b6/describe_resultset--1.0.0.sql)

`describe_resultset` — This Postgres extension implements the "describe a result set" functionality, which, given a query or another PostgreSQL command that returns a result set, shows a list of columns in that result set, along with their respective data types and, possibly, other attributes. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION describe_resultset;

SELECT * FROM describe_resultset('select 1 as foo, 1.1 as bar')
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `describe_resultset(cmd TEXT)` is an extension function and returns `TABLE`.
- `describe_resultset_internal(cmd TEXT)` is an extension function and returns `SETOF`.
- `describe_resultset_data` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
