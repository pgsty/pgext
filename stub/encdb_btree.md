## Usage

Sources:

- [Alibaba Cloud ciphertext-index guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-encdb-btree-to-facilitate-operations-on-ciphertext-indexes)

`encdb_btree` 1.0.0 is an Alibaba Cloud ApsaraDB RDS for PostgreSQL provider extension. It adds the `enc_btree` index access method for equality and range operations on ciphertext columns in an Always confidential database. It is not a portable community extension.

### Prerequisites and installation

Use an Always confidential database whose minor engine version is 20230830 or later. Install the provider's base `encdb` extension first:

```sql
CREATE EXTENSION IF NOT EXISTS encdb;
CREATE EXTENSION encdb_btree;
```

Create single-column, unique, or multicolumn ciphertext indexes with the new access method:

```sql
CREATE INDEX ON test USING enc_btree (t1);
CREATE UNIQUE INDEX ON test USING enc_btree (t2);
CREATE INDEX ON test USING enc_btree (t1, t2, t3);
```

Existing SQL can use these indexes when the provider planner produces a compatible plan. Confirm index selection with `EXPLAIN` against representative encrypted data.

### Operational limits

An `enc_btree` unique index cannot support `INSERT ... ON CONFLICT` because speculative insertion is unsupported. It is also incompatible with foreign-key constraints. These restrictions affect schema and application design, so test migrations before enabling encryption in production.

The extension cannot be dropped while dependent ciphertext indexes remain. `DROP EXTENSION encdb_btree CASCADE` removes those indexes, although it does not remove their tables or table data. Inventory dependent indexes and retain a rollback plan before any drop or engine change. Availability, privileges, upgrades, and backups follow the managed-service contract rather than upstream PostgreSQL packaging.
