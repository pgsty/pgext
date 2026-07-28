## Usage

Sources:

- [Official upstream README](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/README.md)
- [Official extension control file (aethel_smgr.control)](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/extension/aethel_smgr/aethel_smgr.control)
- [Official extension SQL (aethel_smgr--1.0.sql)](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/extension/aethel_smgr/aethel_smgr--1.0.sql)

`aethel_smgr` — AethelDB - A decoupled, serverless PostgreSQL platform. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION aethel_smgr;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `aethel_smgr_status()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
