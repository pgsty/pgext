## Usage

Sources:

- [Official upstream README](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/README.md)
- [Official extension control file (simple.control)](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/simple.control)
- [Official extension SQL (simple--1.0.sql)](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/sql/simple--1.0.sql)

`simple` — This is example of small PostgreSQL extension used for Postgres development trainings. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION simple;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `int_func(int)` is an extension function and returns `int`.
- `text_func(text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
