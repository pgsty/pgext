## Usage

Sources:

- [Official upstream README](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/README.md)
- [Official extension control file (accel_ext.control)](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/transparent-runtime/apps/pg_accel/accel_ext.control)
- [Official extension SQL (accel_ext--1.0.sql)](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/transparent-runtime/apps/pg_accel/accel_ext--1.0.sql)

`accel_ext` — Fine-grained accelerator offload runtime that overlaps GPU, HSM, or inference work with other PostgreSQL requests. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION accel_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `accel_async(integer)` is an extension function and returns `integer`.
- `accel_sync(integer)` is an extension function and returns `integer`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
