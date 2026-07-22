## Usage

Sources:

- [Official README](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/README.md)
- [Official version 1.0 extension SQL](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/tcle--1.0.sql)
- [Official table-access-method implementation](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/tcleheap.c)

`tcle` is an experimental table access method for transparent cell-level encryption. It extends PostgreSQL's heap access method and encrypts columns declared with its special types using AES-256-CBC when tuples cross the storage boundary. The upstream project explicitly says not to use it in production, and the catalog marks it abandoned.

### Setup and Core Workflow

PostgreSQL must be built with OpenSSL. Add `tcle` to `shared_preload_libraries`, restart the cluster, create the extension, and set the database master passphrase before accessing encrypted tables:

```sql
CREATE EXTENSION tcle;
SELECT tcle_set_passphrase('replace-with-a-strong-passphrase');

CREATE TABLE private_records (
    id encrypt_numeric,
    label encrypt_text,
    observed_at encrypt_timestamptz
) USING tcleam;

INSERT INTO private_records
VALUES (1, 'confidential', now());
SELECT * FROM private_records;
```

`encrypt_text` is limited to about 2 KB because it uses plain storage rather than TOAST. `encrypt_numeric` and `encrypt_timestamptz` model numeric and timestamp-with-time-zone values. Each table receives a random table key, stored encrypted by a database master key. The master key is derived from the supplied passphrase and kept in shared memory rather than on disk.

Rotate the master passphrase outside a transaction block:

```sql
SELECT tcle_change_passphrase(
    'replace-with-a-strong-passphrase',
    'replace-with-the-new-passphrase'
);
```

### Critical Limitations

The reviewed upstream version supports only PostgreSQL 12 and 13 and describes itself as experimental. It does not support indexes on encrypted data. Plaintext can still appear in temporary files, SQL statements and logs, client traffic, dumps, and backups; the project also warns that the master passphrase can leak to logs. Table-key rotation requires rebuilding/copying the table and recreating its constraints, indexes, and triggers.

Loss of the master passphrase/key makes data unrecoverable. Test restart, backup/restore, crash recovery, passphrase rotation, `VACUUM FULL`, replication, and key-loss procedures on disposable data. Prefer a maintained encryption design and an external key-management system for real workloads.
