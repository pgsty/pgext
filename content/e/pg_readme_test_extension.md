---
title: "pg_readme_test_extension"
linkTitle: "pg_readme_test_extension"
description: "Test generating a README.md document for extension or schema"
weight: 4301
categories: ["Util"]
width: full
---

Test generating a README.md document for extension or schema

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4301** | {{< badge content="pg_readme_test_extension" link="https://github.com/bigsmoke/pg_readme" >}} | {{< ext "pg_readme_test_extension" "pg_readme" >}} | `0.7.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} |
|   **See Also**    | {{< ext "schedoc" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} |
|    **Siblings**   | {{< ext "pg_readme" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_readme" >}} | `0.7.0` | {{< badge content="18" color="red" alt="pg_readme_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_readme_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_readme" >}} | `0.7.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-readme" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-readme` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_readme_18" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_17" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_16" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_15" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_14" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_readme_18" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_17" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_16" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_15" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_readme_14" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_readme_18" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_17" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_16" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_15" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_14" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_readme_18" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_17" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_16" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_15" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_readme_14" "0.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-readme" >}}     | {{< pkg "postgresql-17-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readme" "0.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_18` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_18` | 0.7.0 | `el9.aarch64` | pgdg | 30.7 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_17` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | 0.7.0 | `el9.aarch64` | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_17` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pg-readme` | 0.7.0 | `d12.aarch64` | pigsty | 18.4 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-readme` | 0.7.0 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-readme` | 0.7.0 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-readme` | 0.7.0 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-readme` | 0.7.0 | `u24.x86_64` | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-readme` | 0.7.0 | `u24.aarch64` | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_16` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_16` | 0.7.0 | `el9.aarch64` | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-16-pg-readme` | 0.7.0 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-readme` | 0.7.0 | `d12.aarch64` | pigsty | 18.4 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-readme` | 0.7.0 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-readme` | 0.7.0 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-readme` | 0.7.0 | `u24.aarch64` | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-readme` | 0.7.0 | `u24.x86_64` | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_15` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_15` | 0.7.0 | `el9.aarch64` | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-15-pg-readme` | 0.7.0 | `d12.x86_64` | pigsty | 18.4 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-readme` | 0.7.0 | `d12.aarch64` | pigsty | 18.4 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-readme` | 0.7.0 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-readme` | 0.7.0 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-readme` | 0.7.0 | `u24.aarch64` | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-readme` | 0.7.0 | `u24.x86_64` | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_14` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | 0.7.0 | `el9.aarch64` | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_14` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-14-pg-readme` | 0.7.0 | `d12.x86_64` | pigsty | 18.4 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-readme` | 0.7.0 | `d12.aarch64` | pigsty | 18.4 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-readme` | 0.7.0 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-readme` | 0.7.0 | `u22.aarch64` | pigsty | 18.7 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-readme` | 0.7.0 | `u24.aarch64` | pigsty | 18.8 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-readme` | 0.7.0 | `u24.x86_64` | pigsty | 18.8 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readme_13` | 0.7.0 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_readme_13-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_readme_13-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_13` | 0.7.0 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_readme_13-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_readme_13-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_13` | 0.7.0 | `el9.aarch64` | pgdg | 30.8 KiB | [pg_readme_13-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_readme_13-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_13` | 0.7.0 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_readme_13-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_readme_13-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-13-pg-readme` | 0.7.0 | `d12.x86_64` | pigsty | 18.4 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-readme` | 0.7.0 | `d12.aarch64` | pigsty | 18.3 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-readme` | 0.7.0 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-readme` | 0.7.0 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-readme` | 0.7.0 | `u24.x86_64` | pigsty | 18.8 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-readme` | 0.7.0 | `u24.aarch64` | pigsty | 18.8 KiB | [postgresql-13-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-13-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_readme" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_readme" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_readme-0.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_readme_test_extension; # get pg_readme_test_extension source code
pig build dep pg_readme_test_extension; # install build dependencies
pig build pkg pg_readme_test_extension; # build extension rpm or deb
pig build ext pg_readme_test_extension; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_readme_test_extension; # install by extension name, for the current active PG version
pig ext install pg_readme; # install via package alias, for the active PG version
pig ext install pg_readme_test_extension -v 17;   # install for PG 17
pig ext install pg_readme_test_extension -v 16;   # install for PG 16
pig ext install pg_readme_test_extension -v 15;   # install for PG 15
pig ext install pg_readme_test_extension -v 14;   # install for PG 14
pig ext install pg_readme_test_extension -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_readme_test_extension;
```

