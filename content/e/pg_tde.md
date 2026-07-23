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
| **7500** | {{< badge content="pg_tde" link="https://github.com/percona/pg_tde" >}} | {{< ext "pg_tde" >}} | `2.2.1` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.4` | {{< bg "18" "pgtde-18" "green" >}} {{< bg "17" "pgtde-17" "red" >}} {{< bg "16" "pgtde-16" "red" >}} {{< bg "15" "pgtde-15" "red" >}} {{< bg "14" "pgtde-14" "red" >}} | `pgtde-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.4` | {{< bg "18" "pgtde-18" "green" >}} {{< bg "17" "pgtde-17" "red" >}} {{< bg "16" "pgtde-16" "red" >}} {{< bg "15" "pgtde-15" "red" >}} {{< bg "14" "pgtde-14" "red" >}} | `pgtde-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgtde-18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgtde-17 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgtde-14 : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtde-18` | `18.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 MiB | [pgtde-18-18.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtde-18-18.4-2PIGSTY.el8.x86_64.rpm) |
| `pgtde-18` | `18.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.6 MiB | [pgtde-18-18.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtde-18-18.4-2PIGSTY.el8.aarch64.rpm) |
| `pgtde-18` | `18.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 MiB | [pgtde-18-18.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtde-18-18.4-2PIGSTY.el9.x86_64.rpm) |
| `pgtde-18` | `18.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.3 MiB | [pgtde-18-18.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtde-18-18.4-2PIGSTY.el9.aarch64.rpm) |
| `pgtde-18` | `18.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 MiB | [pgtde-18-18.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtde-18-18.4-2PIGSTY.el10.x86_64.rpm) |
| `pgtde-18` | `18.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.4 MiB | [pgtde-18-18.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtde-18-18.4-2PIGSTY.el10.aarch64.rpm) |
| `pgtde-18` | `18.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.8 MiB | [pgtde-18_18.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~bookworm_amd64.deb) |
| `pgtde-18` | `18.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.3 MiB | [pgtde-18_18.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~bookworm_arm64.deb) |
| `pgtde-18` | `18.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [pgtde-18_18.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~trixie_amd64.deb) |
| `pgtde-18` | `18.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.4 MiB | [pgtde-18_18.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~trixie_arm64.deb) |
| `pgtde-18` | `18.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.1 MiB | [pgtde-18_18.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~jammy_amd64.deb) |
| `pgtde-18` | `18.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.9 MiB | [pgtde-18_18.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~jammy_arm64.deb) |
| `pgtde-18` | `18.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.9 MiB | [pgtde-18_18.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~noble_amd64.deb) |
| `pgtde-18` | `18.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 MiB | [pgtde-18_18.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~noble_arm64.deb) |
| `pgtde-18` | `18.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.0 MiB | [pgtde-18_18.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~resolute_amd64.deb) |
| `pgtde-18` | `18.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.7 MiB | [pgtde-18_18.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtde-18/pgtde-18_18.4-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_tde" title="Repository" icon="github" subtitle="github.com/percona/pg_tde" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="percona-pg_tde18-2.2.1.tar.gz" >}}
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
