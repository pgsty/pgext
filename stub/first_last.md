## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/first_last/first_last-0.1.1/README)
- [Official extension control file (first_last.control)](https://api.pgxn.org/src/first_last/first_last-0.1.1/first_last.control)
- [Official extension SQL (first_last--0.1.0--0.1.1.sql)](https://api.pgxn.org/src/first_last/first_last-0.1.1/first_last--0.1.0--0.1.1.sql)

`first_last` — This extension provides four aggregates: first(anyelement) first(anyelement, int4) last(anyelement) last(anyelement, int4). Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION first_last;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `agg_first(IN p_state anyarray, IN p_new_element anyelement, IN p_limit int4)` is an extension function and returns `anyarray`.
- `agg_first(IN p_state anyelement, IN p_new_element anyelement)` is an extension function and returns `anyelement`.
- `agg_last(IN p_state anyarray, IN p_new_element anyelement, IN p_limit int4)` is an extension function and returns `anyarray`.
- `agg_last(IN p_state anyelement, IN p_new_element anyelement)` is an extension function and returns `anyelement`.
- `first` is an aggregate exposed by the extension.
- `last` is an aggregate exposed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
