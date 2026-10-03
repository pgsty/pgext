---
title: "kafgres"
linkTitle: "kafgres"
description: "Kafka protocol broker embedded in PostgreSQL"
weight: 9440
categories: ["SIM"]
languages: ["Rust"]
licenses: ["Elastic-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**kafgres**](https://github.com/RayElg/kafgres) : Kafka protocol broker embedded in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9440** | {{< badge content="kafgres" link="https://github.com/RayElg/kafgres" >}} | {{< ext "kafgres" >}} | `0.3.0` | {{< category "SIM" >}} | {{< license "Elastic-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] PGSTY targets PG16; requires preload and broker readiness before topic SQL; segment logs need separate replication and archiving.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `kafgres` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "kafgres_18" "red" >}} {{< bg "17" "kafgres_17" "red" >}} {{< bg "16" "kafgres_16" "green" >}} {{< bg "15" "kafgres_15" "red" >}} {{< bg "14" "kafgres_14" "red" >}} | `kafgres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-kafgres" "red" >}} {{< bg "17" "postgresql-17-kafgres" "red" >}} {{< bg "16" "postgresql-16-kafgres" "green" >}} {{< bg "15" "postgresql-15-kafgres" "red" >}} {{< bg "14" "postgresql-14-kafgres" "red" >}} | `postgresql-$v-kafgres` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "kafgres_18 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_17 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "kafgres_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "kafgres_15 : N/A 0" "gray" >}} | {{< bg "N/A" "kafgres_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-18-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-17-kafgres : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-kafgres : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-kafgres : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-kafgres : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafgres_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.4 MiB | [kafgres_16-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafgres_16-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `kafgres_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [kafgres_16-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafgres_16-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `kafgres_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.4 MiB | [kafgres_16-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafgres_16-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `kafgres_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [kafgres_16-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafgres_16-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `kafgres_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.4 MiB | [kafgres_16-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafgres_16-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `kafgres_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [kafgres_16-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafgres_16-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-kafgres` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.7 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.7 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.3 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.0 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.0 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-kafgres` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.0 MiB | [postgresql-16-kafgres_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/k/kafgres/postgresql-16-kafgres_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RayElg/kafgres" title="Repository" icon="github" subtitle="github.com/RayElg/kafgres" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="kafgres-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg kafgres;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install kafgres;		# install via package name, for the active PG version

pig install kafgres -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'kafgres';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION kafgres;
```

## Usage

Sources:

- [README 0.3.0](https://github.com/RayElg/kafgres/blob/0.3.0/README.md)
- [Configuration 0.3.0](https://github.com/RayElg/kafgres/blob/0.3.0/docs/configuration.md)

- [0.3.0 release](https://github.com/RayElg/kafgres/releases/tag/0.3.0)

`kafgres` embeds a Kafka protocol broker in PostgreSQL. Upstream release 0.3.0 targets PostgreSQL 16 using pgrx 0.16.1. It needs superuser installation, shared preload and a restart. Its license is Elastic License 2.0.

### Enable the broker

```conf
shared_preload_libraries = 'kafgres'
kafgres.database = 'postgres'
kafgres.bind_host = '127.0.0.1'
kafgres.advertised_host = '127.0.0.1'
kafgres.port = 9092
```

```sql
CREATE EXTENSION kafgres;
SELECT kafgres_create_topic('demo', 1);
BEGIN;
SELECT kafgres_produce('demo', 'key', 'value');
COMMIT;
SELECT * FROM kafgres_partition_offsets('demo');
```

Kafka clients connect to the configured broker port. SQL production participates in the caller's transaction. Configure TLS, authentication and ACLs before exposing the listener beyond a trusted local environment.

### Storage and CDC

`kafgres.storage_engine` defaults to segment; its log uses separate files and requires the extension's replication and archive procedures. Configure `kafgres.segment_archive_command` and monitor `kafgres_archive_status()` before relying on segment retention and recovery. Ordinary PostgreSQL WAL/PITR does not cover the entire segment log. The table engine keeps its log in PostgreSQL tables; changing engines does not migrate existing records.

CDC additionally requires `wal_level = logical`; some PostgreSQL builds also require an `output_plugin_libraries` allowlist. Version 0.3.0 supports SQL CDC mappings with projection and filtering. Review the mapping and recovery procedures before deployment; the release artifacts target PostgreSQL 16, and Cargo feature names alone do not prove support for other majors.

### Durability Settings

Version 0.3.0 defaults `kafgres.fsync_before_ack` to on and `kafgres.relaxed_produce_commit` to off. Relaxing the first can lose acknowledged segment records on a power failure; relaxing the second can lose the newest idempotent-producer state after a crash and permit duplicates after retries. These settings have narrower scope than transactional SQL production and do not apply uniformly to the table engine. Preserve the strict defaults until the durability tradeoff is deliberate.
