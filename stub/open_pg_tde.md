## Usage

Sources:

- [Official documentation](https://commandprompt.github.io/open_pg_tde/)
- [Official 2.4.0 README](https://github.com/commandprompt/open_pg_tde/blob/2.4.0/README.md)
- [Official extension control file](https://github.com/commandprompt/open_pg_tde/blob/2.4.0/open_pg_tde.control)

`open_pg_tde` 2.4.0 provides transparent encryption at rest for upstream PostgreSQL through the `tde_heap` table access method. It encrypts table data, indexes, TOAST, WAL, and temporary spill files while allowing encrypted and plain tables to coexist.

### Core Workflow

Build PostgreSQL with the matching `open_pg_tde` core patch, install the extension, add `open_pg_tde` to `shared_preload_libraries`, and restart the server. Then configure a key provider and principal key before creating encrypted tables:

```sql
CREATE EXTENSION open_pg_tde;

SELECT open_pg_tde_add_database_key_provider_file(
  'file-keyring',
  '/var/lib/postgresql/open_pg_tde.keyring'
);
SELECT open_pg_tde_create_key_using_database_key_provider(
  'primary-key',
  'file-keyring'
);
SELECT open_pg_tde_set_key_using_database_key_provider(
  'primary-key',
  'file-keyring'
);

CREATE TABLE secret_data (
  id bigint PRIMARY KEY,
  payload text
) USING tde_heap;

SELECT open_pg_tde_is_encrypted('secret_data');
```

Use a production key-management provider rather than the file-keyring example when the threat model requires separation from the database host.

### Important Objects and Settings

- `tde_heap` is the encrypted table access method.
- `open_pg_tde.data_cipher` selects AES-128-XTS or AES-256-XTS for relation data.
- database- and server-key-provider functions register keyring, KMIP-compatible, or OpenBao providers and select principal keys.
- `open_pg_tde_is_encrypted(regclass)` checks whether a relation uses encrypted storage.

### Requirements and Caveats

- Version 2.4.0 supports upstream PostgreSQL 16, 17, and 18 according to its tagged README.
- The extension requires a matching patched PostgreSQL source tree; an unpatched stock server does not provide the required storage-manager and WAL hooks.
- Changing `shared_preload_libraries` requires a server restart.
- System catalogs and statistics are outside the stated encryption coverage. Review the official threat model for backups, keys, memory, logs, and host compromise.
- Back up key material separately and test restore, rotation, failover, WAL replay, and loss-of-provider behavior before production use.
