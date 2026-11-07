## Usage

Sources:

- [1.10.2 README](https://github.com/orioledb/orioledb/blob/v1.10.2/README.md)
- [1.10.2 release](https://github.com/orioledb/orioledb/releases/tag/v1.10.2)
- [Control file](https://github.com/orioledb/orioledb/blob/v1.10.2/orioledb.control)
- [Patched PostgreSQL](https://github.com/orioledb/postgres)
- [Patched server versions](https://github.com/orioledb/orioledb/blob/v1.10.2/.pgtags)

`orioledb` provides an alternative table storage engine using undo-based MVCC, copy-on-write checkpoints and row-level WAL. It requires a specifically patched PostgreSQL server and remains a public beta.

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

Use ordinary DML on OrioleDB tables, subject to the documented compatibility limits:

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

### Compatibility and 1.10.2 Maintenance

The release's `.pgtags` pins PostgreSQL 16, 17 and 18 to `patches16_50`, `patches17_23` and `patches18_4`, respectively. An ordinary PostgreSQL installation is not sufficient. Only ICU, C and POSIX collations are supported. The undo and page-merging design reduces native-table vacuum requirements; other PostgreSQL tables still need normal maintenance.

Release 1.10.2 keeps SQL version `1.10` and introduces no storage or WAL format change. It fixes the runtime version reported by `ORIOLEDB_VERSION`, row-version handling under concurrency, capture of updates during concurrent index builds, background-writer crashes with range/array/composite keys, NULL handling in diagnostic functions, expression-index memory leaks and standby tablespace moves.

Use matching server patches and library files on every node and follow the version-specific upgrade procedure with a backup. Do not request SQL version 1.10.2: it does not exist. The public-beta limitation remains; the upstream project does not recommend general production use.
