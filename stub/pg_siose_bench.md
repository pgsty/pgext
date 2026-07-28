## Usage

Sources:

- [Official upstream README](https://github.com/siose-innova/pg_siose_bench/blob/e28280ad8b97a19568586056b0d0259e35ad1d1b/README.md)
- [Official extension control file (pg_siose_bench.control)](https://github.com/siose-innova/pg_siose_bench/blob/e28280ad8b97a19568586056b0d0259e35ad1d1b/pg_siose_bench.control)

`pg_siose_bench` — A PostgreSQL extension with tools for benchmarking different SIOSE database configurations (pure relational, indexed, json, jsonb, xml, etc). Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_siose_bench;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `postgis`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
