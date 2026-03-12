


## Usage

> [pg_tde: Transparent Data Encryption for PostgreSQL](https://github.com/percona/pg_tde)

`pg_tde` provides Transparent Data Encryption (TDE) at the file level, encrypting tuples, WAL, and indexes. It works with the `tde_heap` access method and supports keyringfile and external Key Management Systems (KMS).

```sql
CREATE EXTENSION pg_tde;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'pg_tde'
```

### Setting Up a Key Provider

```sql
-- File-based key provider (database-level)
SELECT pg_tde_add_database_key_provider_file('file_keyring', '/path/to/keyring');

-- Or global-level key provider
SELECT pg_tde_add_global_key_provider_file('file_keyring', '/path/to/keyring');

-- Set the encryption key using a database key provider
SELECT pg_tde_set_key_using_database_key_provider('my_key', 'file_keyring');

-- Or using a global key provider
SELECT pg_tde_set_key_using_global_key_provider('my_key', 'file_keyring');
```

### Creating Encrypted Tables

```sql
CREATE TABLE sensitive_data (
    id serial PRIMARY KEY,
    secret text
) USING tde_heap;
```

All data in tables created with `USING tde_heap` is transparently encrypted on disk.

### Checking Encryption Status

```sql
SELECT pg_tde_is_encrypted('sensitive_data');
```

### Additional Functions

| Function | Description |
|----------|-------------|
| `pg_tde_add_database_key_provider_file(name, path)` | Add a file-based database key provider |
| `pg_tde_add_global_key_provider_file(name, path)` | Add a file-based global key provider |
| `pg_tde_add_database_key_provider_vault_v2(...)` | Add a HashiCorp Vault database key provider |
| `pg_tde_add_global_key_provider_vault_v2(...)` | Add a HashiCorp Vault global key provider |
| `pg_tde_set_key_using_database_key_provider(key, provider)` | Set encryption key via database provider |
| `pg_tde_set_key_using_global_key_provider(key, provider)` | Set encryption key via global provider |
| `pg_tde_is_encrypted(table)` | Check if a table is encrypted |

### Notes

- Works only with Percona Server for PostgreSQL 17+
- Encrypts tuples, WAL, and indexes
- Does not yet encrypt temporary files and statistics
