## Usage

Sources:

- [Official upstream README](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/README.md)
- [Official extension control file (pg_bert.control)](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/pg_bert.control)
- [Official implementation source](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/src/lib.rs)

`pg_bert` — BERT tokenizer implemented as a PostgreSQL full-text-search parser. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_bert;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
