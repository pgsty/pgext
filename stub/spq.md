## Usage

Sources:

- [Official upstream README](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/README.md)
- [Official extension control file (spq.control)](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/src/backend/distributed/spq.control)
- [Official extension SQL (spq--2.0.0.sql)](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/src/backend/distributed/sql/spq--2.0.0.sql)

`spq` — Spq Share-nothing parallel query. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION spq;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `2.0.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
