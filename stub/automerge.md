## Usage

Sources:

- [Official upstream README](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/README.md)
- [Official extension control file (automerge.control)](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/automerge.control)
- [Official implementation source](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/src/automerge/automerge.c)

`automerge` — pg_crdt is an experimental extension adding support for conflict-free replicated data types (CRDTs) in Postgres. Use it when application data needs this type, domain, or its operators. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION automerge;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
