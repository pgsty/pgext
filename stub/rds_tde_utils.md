## Usage

Sources:

- [Official Alibaba Cloud rds_tde_utils guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-tde-utils-extension-to-encrypt-and-decrypt-multiple-data-records-at-a-time)

`rds_tde_utils` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension for encrypting or decrypting one table and its indexes, or every table and index in a database, after Transparent Data Encryption is enabled. Version `1.1` is provider-only and performs potentially large synchronous rewrites.

### Prerequisites and Setup

The provider requires PostgreSQL 10 or later, an RDS minor engine of `20221030` or later, enabled TDE, and a privileged account:

```sql
CREATE EXTENSION rds_tde_utils;
```

Confirm the exact supported engine and take a recoverable backup before changing encryption state.

### Core Workflow

Use the lazy method to reduce impact, or the full method during an off-peak maintenance window:

```sql
SELECT rds_tde_lazy_encrypt_table('public.orders'::regclass);
SELECT rds_tde_encrypt_table('public.orders'::regclass);

SELECT rds_tde_lazy_decrypt_table('public.orders'::regclass);
SELECT rds_tde_decrypt_table('public.orders'::regclass);
```

Database-wide variants process all tables and indexes:

```sql
SELECT rds_tde_lazy_encrypt_database();
SELECT rds_tde_encrypt_database();

SELECT rds_tde_lazy_decrypt_database();
SELECT rds_tde_decrypt_database();
```

### Method Selection

- The four `rds_tde_lazy_*` functions use rewrite behavior described as similar to lazy VACUUM.
- The four full functions use behavior described as similar to `VACUUM FULL`; the provider warns not to run them at peak time.
- Table functions accept `regclass` and include the table's indexes; database functions cover all tables and indexes in the current database.

### Operational Caveats

All eight calls are synchronous and return only after processing completes. Large databases can take a long time and may experience locking, I/O, extra storage demand, replica lag, and latency spikes. Inventory object sizes, exclude peak traffic, monitor free space and replication, and test cancellation and restart behavior on a staging instance. TDE enablement, KMS/key policy, snapshots, backups, replicas, and disaster recovery are managed-service concerns outside these helper functions. Verify encryption state through Alibaba Cloud's supported controls rather than assuming a successful SQL return proves end-to-end key protection.
