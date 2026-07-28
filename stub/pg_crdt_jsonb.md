## Usage

Sources:

- [Official upstream README](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/README.md)
- [Official extension control file (pg_crdt_jsonb.control)](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/pg_crdt_jsonb.control)
- [Official extension SQL (pg_crdt_jsonb--1.0.sql)](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/pg_crdt_jsonb--1.0.sql)

`pg_crdt_jsonb` — For each element in the top-level array, its own Timestamp is created. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_crdt_jsonb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `crdt_jsonb_append(crdt_jsonb, jsonb)` is an extension function and returns `crdt_jsonb`.
- `crdt_jsonb_in(cstring)` is an extension function and returns `crdt_jsonb`.
- `crdt_jsonb_out(crdt_jsonb)` is an extension function and returns `cstring`.
- `crdt_jsonb_recv(internal)` is an extension function and returns `crdt_jsonb`.
- `crdt_jsonb_send(crdt_jsonb)` is an extension function and returns `bytea`.
- `get_jsonb(crdt_jsonb)` is an extension function and returns `jsonb`.
- `crdt_jsonb` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
