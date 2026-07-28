## Usage

Sources:

- [Official upstream README](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/README.md)
- [Official extension control file (pg_extism.control)](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/pg_extism.control)
- [Official implementation source](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/src/lib.rs)

`pg_extism` — An Extism sample showing how you can run Extism plugins from PostgreSQl using pgrx! Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_extism;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `extism_call` is an extension function.
- `extism_define` is an extension function.
- `to_lowercase` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
