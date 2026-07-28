## Usage

Sources:

- [Official upstream README](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/README.md)
- [Official extension control file](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/tsearch/dict_maxlen/dict_maxlen.control)
- [Official extension SQL](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/tsearch/dict_maxlen/dict_maxlen--1.0.sql)

`dict_maxlen` — A repository of custom PostgreSQL functions and extensions. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dict_maxlen;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dictmaxlen_init(internal)` is an extension function and returns `internal`.
- `dictmaxlen_lexize(internal, internal, internal, internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
