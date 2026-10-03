## Usage

Sources:

- [pgiceberg.control](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/pgiceberg.control)
- [README.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/README.md)
- [sql/pgiceberg--0.1.0.sql](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/sql/pgiceberg--0.1.0.sql)
- [docs/design/postgres-extension-surfaces.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/docs/design/postgres-extension-surfaces.md)
- [docs/design/postgres-iceberg-commit.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/docs/design/postgres-iceberg-commit.md)

`pgiceberg` offers three separate Iceberg interfaces: foreign tables, native Iceberg-backed tables and logical-decoding mirrors. It uses Apache Iceberg C++ for catalog and Parquet operations; the extension version is 0.1.0, not the bundled library version.

### Core Workflow

```sql
CREATE EXTENSION pgiceberg;
SELECT pgiceberg.add_catalog('demo', 'sqlite',
  '/tmp/pgiceberg_demo.db', '/tmp/pgiceberg_demo_warehouse');
CREATE SERVER iceberg_demo FOREIGN DATA WRAPPER pgiceberg OPTIONS (catalog 'demo');
SELECT * FROM pgiceberg.commit_recovery_log();
```

### Operational Boundaries

Superuser installation and a matching C++ runtime are required. Catalog paths and credentials grant server-side access to external storage. `IMPORT FOREIGN SCHEMA`, schema-diff/refresh helpers, snapshot reads and unpartitioned UPDATE/DELETE are available. Historical snapshots reject writes; logical mirrors require their own logical-decoding setup. Iceberg publishes occur before the PostgreSQL commit and are not a two-phase transaction: crashes or multi-table failures may leave partial external commits. Inspect recovery logs and follow the documented reconciliation procedure. PostgreSQL isolation levels do not isolate other Iceberg writers.
