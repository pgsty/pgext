---
title: "pg_tde"
linkTitle: "pg_tde"
description: "Percona pg_tde access method"
weight: 7500
categories: ["SEC"]
width: full
---

[**pg_tde**](https://github.com/percona/pg_tde) : Percona pg_tde access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7500** | {{< badge content="pg_tde" link="https://github.com/percona/pg_tde" >}} | {{< ext "pg_tde" >}} | `2.2` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgcrypto" >}} {{< ext "anon" >}} {{< ext "pgcryptokey" >}} {{< ext "faker" >}} {{< ext "sslutils" >}} {{< ext "uuid-ossp" >}} |

> [!Note] works on percona postgres tde fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "percona-postgresql18" "green" >}} {{< bg "17" "percona-postgresql17" "green" >}} {{< bg "16" "percona-postgresql16" "red" >}} {{< bg "15" "percona-postgresql15" "red" >}} {{< bg "14" "percona-postgresql14" "red" >}} | `percona-postgresql$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "percona-postgresql-18" "green" >}} {{< bg "17" "percona-postgresql-17" "green" >}} {{< bg "16" "percona-postgresql-16" "red" >}} {{< bg "15" "percona-postgresql-15" "red" >}} {{< bg "14" "percona-postgresql-14" "red" >}} | `percona-postgresql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "percona-postgresql18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "percona-postgresql-18 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-17 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-16 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-15 : N/A 0" "gray" >}} | {{< bg "N/A" "percona-postgresql-14 : N/A 0" "gray" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_tde" title="Repository" icon="github" subtitle="github.com/percona/pg_tde" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tde;		# install via package name, for the active PG version

pig install pg_tde -v 18;   # install for PG 18
pig install pg_tde -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_tde';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tde;
```

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
