


## Usage

> [orioledb: A cloud-native storage engine for PostgreSQL](https://github.com/orioledb/orioledb)

OrioleDB is a new storage engine for PostgreSQL that provides modern approaches to database capacity, capabilities, and performance. It uses undo log-based MVCC, copy-on-write checkpoints, and row-level WAL to eliminate bloat and the need for VACUUM.

### Configuration

Add to `postgresql.conf` (requires restart):

```ini
shared_preload_libraries = 'orioledb.so'
```

Then enable the extension:

```sql
CREATE EXTENSION orioledb;
```

### Creating Tables

Use the `USING orioledb` clause to create tables with the OrioleDB storage engine:

```sql
CREATE TABLE my_table (
    id serial PRIMARY KEY,
    name text,
    value numeric
) USING orioledb;
```

All standard PostgreSQL operations work on OrioleDB tables:

```sql
INSERT INTO my_table (name, value) VALUES ('test', 42);
SELECT * FROM my_table WHERE id = 1;
UPDATE my_table SET value = 100 WHERE id = 1;
DELETE FROM my_table WHERE id = 1;
```

### Collation Requirements

OrioleDB tables support only **ICU**, **C**, and **POSIX** collations. To avoid specifying COLLATE on every text field, create the database with an appropriate default:

```sql
CREATE DATABASE mydb LOCALE 'C' TEMPLATE template0;
-- OR
CREATE DATABASE mydb LOCALE_PROVIDER icu ICU_LOCALE 'en' TEMPLATE template0;
```

### Key Benefits

- **No bloat**: Undo log-based MVCC means old tuple versions do not bloat main storage
- **No VACUUM needed**: Page-merging and undo log reclaim space automatically
- **No wraparound problem**: Native 64-bit transaction identifiers
- **Lock-less page reading**: In-memory pages linked directly to storage pages
- **Row-level WAL**: Compact write-ahead logging suitable for parallel apply

### Limitations

- Public beta status -- recommended for testing, not production
- Requires a patched PostgreSQL build from [orioledb/postgres](https://github.com/orioledb/postgres)
- Only ICU, C, and POSIX collations are supported
