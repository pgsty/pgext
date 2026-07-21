## Usage

Sources:

- [pg_tde 2.2 setup](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/setup.md)
- [Key-provider and key-management functions](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/functions.md)
- [pg_tde 2.2.0 release notes](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/release-notes/release-notes-v2.2.0.md)
- [Transparent data encryption limitations](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/index/tde-limitations.md)
- [TDE table access method](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/index/table-access-method.md)
- [WAL encryption](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/wal-encryption.md)

pg_tde provides transparent data encryption for Percona Server for PostgreSQL. It encrypts table data through the tde_heap access method and can encrypt WAL, with keys managed by file, HashiCorp Vault, or KMIP providers. It is not a drop-in extension for community PostgreSQL.

### Preload and Create the Extension

Add the library and restart the server:

    shared_preload_libraries = 'pg_tde'

Then enable pg_tde in every database that will use encrypted tables:

    CREATE EXTENSION pg_tde;

Run setup as a superuser or suitably privileged database owner. Upstream pg_tde 2.2 is tied to compatible Percona Server for PostgreSQL 17 or 18 builds; the 2.2.0 release notes warn that it is incompatible with Percona Distribution releases older than 17.10 and 18.4.

### Configure a Key Provider

Register a provider, then set a principal key. A local file provider is useful for evaluation:

    SELECT pg_tde_add_database_key_provider_file(
      'local-file',
      '/secure/path/pg_tde_keys'
    );

    SELECT pg_tde_set_principal_key(
      'app-principal-key',
      'local-file'
    );

For production, upstream recommends an external provider such as Vault or KMIP rather than the local-file provider. Protect provider credentials, key files, backups, and recovery procedures independently of the database files.

Provider management includes database- and server-global variants for file, Vault, and KMIP providers, plus functions to list, change, and delete providers and to inspect or rotate the principal key.

### Create and Convert Encrypted Tables

Create a table with the encrypted access method:

    CREATE TABLE customer_secrets (
      id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      payload jsonb NOT NULL
    ) USING tde_heap;

Convert an existing table only after testing lock, rewrite, disk-space, and backup implications:

    ALTER TABLE customer_secrets SET ACCESS METHOD tde_heap;

Changing a table access method rewrites the table. Plan maintenance time and confirm indexes, replicas, backups, and restores on a staging copy.

### Enable WAL Encryption

WAL encryption is a separate server setting:

    pg_tde.wal_encrypt = on

Changing it requires a restart. Confirm that every primary, standby, backup, archive, and recovery host has the required provider configuration and key access before enabling it.

### Object Index

- tde_heap: encrypted table access method.
- pg_tde_add_database_key_provider_file/vault/kmip: database-scoped provider registration.
- pg_tde_add_global_key_provider_file/vault/kmip: server-global provider registration.
- pg_tde_set_principal_key and pg_tde_set_server_principal_key: select the key used to protect data-encryption keys.
- pg_tde_list_all_key_providers: inspect registered providers.
- pg_tde_change_key_provider_* and pg_tde_delete_key_provider: manage provider definitions.
- pg_tde.wal_encrypt: enable encryption of write-ahead log records.
- pg_tde_upgrade: upgrade helper introduced in the 2.2 line.

### Security and Recovery Boundaries

- pg_tde encrypts supported user-table storage, not every PostgreSQL artifact. System catalogs, planner statistics, and temporary spill files are among the documented exclusions.
- Upstream warns that pg_rewind and pg_tde_rewind between diverged nodes can corrupt a cluster. Follow the documented rebuild/recovery path instead of assuming ordinary rewind is safe.
- Starting recovery without pg_tde preloaded can corrupt encrypted data. Validate disaster-recovery automation with the library and keys present.
- Percona documents incompatibilities with Citus and TimescaleDB in Percona Server and limitations for several WAL-inspection and recovery tools.
- Encryption does not replace SQL privileges, TLS, host hardening, audit logging, or tested backups. Loss of keys can make otherwise intact backups unrecoverable.
