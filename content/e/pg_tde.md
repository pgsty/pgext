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
| **7500** | {{< badge content="pg_tde" link="https://github.com/percona/pg_tde" >}} | {{< ext "pg_tde" >}} | `2.1` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.1` | {{< bg "18" "percona-postgresql18" "green" >}} {{< bg "17" "percona-postgresql17" "green" >}} {{< bg "16" "percona-postgresql16" "red" >}} {{< bg "15" "percona-postgresql15" "red" >}} {{< bg "14" "percona-postgresql14" "red" >}} | `percona-postgresql$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.1` | {{< bg "18" "percona-postgresql-18" "green" >}} {{< bg "17" "percona-postgresql-17" "green" >}} {{< bg "16" "percona-postgresql-16" "red" >}} {{< bg "15" "percona-postgresql-15" "red" >}} {{< bg "14" "percona-postgresql-14" "red" >}} | `percona-postgresql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : FORK 0" "red" >}}      |


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
