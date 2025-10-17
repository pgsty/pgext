---
title: "pgpdf"
linkTitle: "pgpdf"
description: "PDF type with meta admin & Full-Text Search"
weight: 3530
categories: ["Type"]
width: full
---

PDF type with meta admin & Full-Text Search

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3530** | {{< badge content="pgpdf" link="https://github.com/Florents-Tselai/pgpdf" >}} | {{< ext "pgpdf" "pgpdf" >}} | `0.1.0` | {{< category "TYPE" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLdtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgpdf" >}} | `0.1.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgpdf_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgpdf" >}} | `0.1.0` | {{< badge content="18" color="red" alt="postgresql-18-pgpdf" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgpdf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgpdf_18" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgpdf_18-0.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpdf_17" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgpdf_17-0.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpdf_16" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgpdf_16-0.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpdf_15" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgpdf_15-0.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpdf_14" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgpdf_14-0.1.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgpdf_18" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgpdf_18-0.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpdf_17" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgpdf_17-0.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpdf_16" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgpdf_16-0.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpdf_15" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgpdf_15-0.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpdf_14" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgpdf_14-0.1.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgpdf_18" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgpdf_18-0.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpdf_17" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgpdf_17-0.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpdf_16" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgpdf_16-0.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpdf_15" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgpdf_15-0.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpdf_14" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgpdf_14-0.1.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgpdf_18" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgpdf_18-0.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpdf_17" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgpdf_17-0.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpdf_16" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgpdf_16-0.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpdf_15" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgpdf_15-0.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpdf_14" "0.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgpdf_14-0.1.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgpdf" >}}     | {{< pkg "postgresql-17-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgpdf" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_18` | 0.1.0 | `el8.x86_64` | pgdg | 15.4 KiB | [pgpdf_18-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgpdf_18-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_18` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_18-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgpdf_18-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_18` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_18-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgpdf_18-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgpdf_18` | 0.1.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgpdf_18-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgpdf_18-0.1.0-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_17` | 0.1.0 | `el8.x86_64` | pigsty | 17.9 KiB | [pgpdf_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpdf_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgpdf_17` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_17-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgpdf_17-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_17` | 0.1.0 | `el8.aarch64` | pigsty | 17.4 KiB | [pgpdf_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpdf_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgpdf_17` | 0.1.0 | `el8.x86_64` | pgdg | 15.4 KiB | [pgpdf_17-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgpdf_17-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_17` | 0.1.0 | `el9.x86_64` | pigsty | 18.3 KiB | [pgpdf_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpdf_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgpdf_17` | 0.1.0 | `el9.aarch64` | pigsty | 17.6 KiB | [pgpdf_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpdf_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgpdf_17` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_17-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgpdf_17-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgpdf_17` | 0.1.0 | `el9.aarch64` | pgdg | 15.0 KiB | [pgpdf_17-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgpdf_17-0.1.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgpdf` | 0.1.0 | `d12.x86_64` | pigsty | 26.0 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgpdf` | 0.1.0 | `d12.aarch64` | pigsty | 25.6 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgpdf` | 0.1.0 | `u22.x86_64` | pigsty | 27.4 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgpdf` | 0.1.0 | `u22.aarch64` | pigsty | 27.0 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgpdf` | 0.1.0 | `u24.x86_64` | pigsty | 26.2 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgpdf` | 0.1.0 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-17-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_16` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_16-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgpdf_16-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_16` | 0.1.0 | `el8.aarch64` | pigsty | 17.4 KiB | [pgpdf_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpdf_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgpdf_16` | 0.1.0 | `el8.x86_64` | pgdg | 15.4 KiB | [pgpdf_16-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgpdf_16-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_16` | 0.1.0 | `el8.x86_64` | pigsty | 17.9 KiB | [pgpdf_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpdf_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgpdf_16` | 0.1.0 | `el9.aarch64` | pigsty | 17.6 KiB | [pgpdf_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpdf_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgpdf_16` | 0.1.0 | `el9.x86_64` | pigsty | 18.3 KiB | [pgpdf_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpdf_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgpdf_16` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_16-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgpdf_16-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgpdf_16` | 0.1.0 | `el9.aarch64` | pgdg | 15.0 KiB | [pgpdf_16-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgpdf_16-0.1.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pgpdf` | 0.1.0 | `d12.aarch64` | pigsty | 25.6 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgpdf` | 0.1.0 | `d12.x86_64` | pigsty | 26.0 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgpdf` | 0.1.0 | `u22.aarch64` | pigsty | 27.0 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgpdf` | 0.1.0 | `u22.x86_64` | pigsty | 27.4 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgpdf` | 0.1.0 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgpdf` | 0.1.0 | `u24.x86_64` | pigsty | 26.2 KiB | [postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-16-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_15` | 0.1.0 | `el8.x86_64` | pgdg | 15.4 KiB | [pgpdf_15-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgpdf_15-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_15` | 0.1.0 | `el8.x86_64` | pigsty | 17.9 KiB | [pgpdf_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpdf_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgpdf_15` | 0.1.0 | `el8.aarch64` | pigsty | 17.3 KiB | [pgpdf_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpdf_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgpdf_15` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_15-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgpdf_15-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_15` | 0.1.0 | `el9.aarch64` | pgdg | 15.0 KiB | [pgpdf_15-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgpdf_15-0.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgpdf_15` | 0.1.0 | `el9.aarch64` | pigsty | 17.6 KiB | [pgpdf_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpdf_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgpdf_15` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_15-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgpdf_15-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgpdf_15` | 0.1.0 | `el9.x86_64` | pigsty | 18.3 KiB | [pgpdf_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpdf_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pgpdf` | 0.1.0 | `d12.aarch64` | pigsty | 25.7 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgpdf` | 0.1.0 | `d12.x86_64` | pigsty | 26.0 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgpdf` | 0.1.0 | `u22.aarch64` | pigsty | 27.0 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgpdf` | 0.1.0 | `u22.x86_64` | pigsty | 27.5 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgpdf` | 0.1.0 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgpdf` | 0.1.0 | `u24.x86_64` | pigsty | 26.2 KiB | [postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-15-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_14` | 0.1.0 | `el8.x86_64` | pgdg | 15.4 KiB | [pgpdf_14-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgpdf_14-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_14` | 0.1.0 | `el8.aarch64` | pigsty | 17.3 KiB | [pgpdf_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpdf_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgpdf_14` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_14-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgpdf_14-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_14` | 0.1.0 | `el8.x86_64` | pigsty | 17.9 KiB | [pgpdf_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpdf_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgpdf_14` | 0.1.0 | `el9.aarch64` | pgdg | 15.0 KiB | [pgpdf_14-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgpdf_14-0.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgpdf_14` | 0.1.0 | `el9.x86_64` | pigsty | 18.3 KiB | [pgpdf_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpdf_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgpdf_14` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_14-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgpdf_14-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgpdf_14` | 0.1.0 | `el9.aarch64` | pigsty | 17.6 KiB | [pgpdf_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpdf_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgpdf` | 0.1.0 | `d12.aarch64` | pigsty | 25.6 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgpdf` | 0.1.0 | `d12.x86_64` | pigsty | 25.9 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgpdf` | 0.1.0 | `u22.x86_64` | pigsty | 27.4 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgpdf` | 0.1.0 | `u22.aarch64` | pigsty | 26.9 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgpdf` | 0.1.0 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgpdf` | 0.1.0 | `u24.x86_64` | pigsty | 26.1 KiB | [postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-14-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpdf_13` | 0.1.0 | `el8.aarch64` | pgdg | 14.8 KiB | [pgpdf_13-0.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgpdf_13-0.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgpdf_13` | 0.1.0 | `el8.x86_64` | pgdg | 15.1 KiB | [pgpdf_13-0.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgpdf_13-0.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgpdf_13` | 0.1.0 | `el8.aarch64` | pigsty | 17.3 KiB | [pgpdf_13-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpdf_13-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgpdf_13` | 0.1.0 | `el8.x86_64` | pigsty | 17.6 KiB | [pgpdf_13-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpdf_13-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgpdf_13` | 0.1.0 | `el9.x86_64` | pigsty | 18.2 KiB | [pgpdf_13-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpdf_13-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgpdf_13` | 0.1.0 | `el9.aarch64` | pigsty | 17.5 KiB | [pgpdf_13-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpdf_13-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgpdf_13` | 0.1.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgpdf_13-0.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgpdf_13-0.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgpdf_13` | 0.1.0 | `el9.x86_64` | pgdg | 15.6 KiB | [pgpdf_13-0.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgpdf_13-0.1.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-pgpdf` | 0.1.0 | `d12.aarch64` | pigsty | 25.5 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgpdf` | 0.1.0 | `d12.x86_64` | pigsty | 25.7 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgpdf` | 0.1.0 | `u22.aarch64` | pigsty | 26.8 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgpdf` | 0.1.0 | `u22.x86_64` | pigsty | 27.2 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgpdf` | 0.1.0 | `u24.x86_64` | pigsty | 25.9 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgpdf` | 0.1.0 | `u24.aarch64` | pigsty | 25.7 KiB | [postgresql-13-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgpdf/postgresql-13-pgpdf_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/pgpdf" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/pgpdf" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgpdf-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgpdf; # get pgpdf source code
pig build dep pgpdf; # install build dependencies
pig build pkg pgpdf; # build extension rpm or deb
pig build ext pgpdf; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgpdf; # install by extension name, for the current active PG version
pig ext install pgpdf; # install via package alias, for the active PG version
pig ext install pgpdf -v 18;   # install for PG 18
pig ext install pgpdf -v 17;   # install for PG 17
pig ext install pgpdf -v 16;   # install for PG 16
pig ext install pgpdf -v 15;   # install for PG 15
pig ext install pgpdf -v 14;   # install for PG 14
pig ext install pgpdf -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgpdf;
```



--------

## Usage

The actual PDF parsing is done by [poppler](https://poppler.freedesktop.org).

This allows you to work with PDFs in an ACID-compliant way.
The usual alternative relies on external scripts or services which can easily
make your data ingestion pipeline brittle and leave your raw data out-of-sync.

- [Full Text Search on PDFs With Postgres](https://tselai.com/full-text-search-pdf-postgres)
- [pgpdf: pdf type for Postgres](https://tselai.com/pgpdf-pdf-type-postgres)

Download some PDFs.

```sh
wget https://wiki.postgresql.org/images/e/ea/PostgreSQL_Introduction.pdf -O /tmp/pgintro.pdf
wget https://pdfobject.com/pdf/sample.pdf -O /tmp/sample.pdf
```

You can create a `pdf` type, by casting either a `text` filepath or `bytea` column.

```sql
CREATE EXTENSION pgpdf;
SELECT '/tmp/pgintro.pdf'::pdf;
```

```sql
                                       pdf                                        
----------------------------------------------------------------------------------
 PostgreSQL Introduction                                                         +
 Digoal.Zhou                                                                     +
 7/20/2011Catalog                                                                +
  PostgreSQL Origin 
```

If you don’t have the PDF file in your filesystem, but have already stored its content in a `bytea` column, you can just cast it to `pdf`.

```sql
SELECT pg_read_binary_file('/tmp/pgintro.pdf')::bytea::pdf;
```

--------

## Examples

Create a table with a `pdf` column:

```sql
CREATE TABLE pdfs(name text primary key, doc pdf);

INSERT INTO pdfs VALUES ('pgintro', '/tmp/pgintro.pdf');
INSERT INTO pdfs VALUES ('pgintro', '/tmp/sample.pdf');
```

Parsing and validation should happen automatically.
The files will be read from the disk only once!

> [!NOTE]
> The filepath should be accessible by the `postgres` process / user!
> That's different than the user running psql.
> If you don't understand what this means, as your DBA!

### String Functions and Operators

Standard Postgres [String Functions and Operators](https://www.postgresql.org/docs/17/functions-string.html)
should work as usual:

```sql
SELECT 'Below is the PDF we received ' || '/tmp/pgintro.pdf'::pdf;
```

```sql
SELECT upper('/tmp/pgintro.pdf'::pdf::text);
```

```sql
SELECT name
FROM pdfs
WHERE doc::text LIKE '%Postgres%';
```

### Full-Text Search (FTS)

You can also perform full-text search (FTS), since you can work on a `pdf` file like normal text.

```sql
SELECT '/tmp/pgintro.pdf'::pdf::text @@ to_tsquery('postgres');
```

```sql
 ?column? 
----------
 t
(1 row)
```

```sql
SELECT '/tmp/pgintro.pdf'::pdf::text @@ to_tsquery('oracle');
```

```sql
 ?column? 
----------
 f
(1 row)
```

### Document similarity with `pg_trgm`

You can use [pg_trgm](https://postgresql.org/docs/17/interactive/pgtrgm.html)
to get the similarity between two documents:

```sql
CREATE EXTENSION pg_trgm;

SELECT similarity('/tmp/pgintro.pdf'::pdf::text, '/tmp/sample.pdf'::pdf::text);
```

### Metadata

The following functions are available:

- `pdf_title(pdf) → text`
- `pdf_author(pdf) → text`
- `pdf_num_pages(pdf) → integer`

  Total number of pages in the document
- `pdf_page(pdf, integer) → text`

  Get the i-th page as text
- `pdf_creator(pdf) → text`
- `pdf_keywords(pdf) → text`
- `pdf_metadata(pdf) → text`
- `pdf_version(pdf) → text`
- `pdf_subject(pdf) → text`
- `pdf_creation(pdf) → timestamp`
- `pdf_modification(pdf) → timestamp`

```sql
SELECT pdf_title('/tmp/pgintro.pdf');
```

```sql
        pdf_title        
-------------------------
 PostgreSQL Introduction
(1 row)
```

```sql
SELECT pdf_author('/tmp/pgintro.pdf');
```

```sql
 pdf_author 
------------
 周正中
(1 row)
```

Getting a subset of pages

```sql
SELECT pdf_num_pages('/tmp/pgintro.pdf');
```

```sql
 pdf_num_pages 
---------------
            24
(1 row)
```

```sql
SELECT pdf_page('/tmp/pgintro.pdf', 1);
```

```sql
           pdf_page           
------------------------------
 Catalog                     +
  PostgreSQL Origin         +
  Layout                    +
  Features                  +
  Enterprise Class Attribute+
  Case
(1 row)
```

```sql
SELECT pdf_subject('/tmp/pgintro.pdf');
```

```sql
 pdf_subject 
-------------
 
(1 row)
```

```sql
SELECT pdf_creation('/tmp/pgintro.pdf');
```

```sql
       pdf_creation       
--------------------------
 Wed Jul 20 11:13:37 2011
(1 row)
```

```sql
SELECT pdf_modification('/tmp/pgintro.pdf');
```

```sql
     pdf_modification     
--------------------------
 Wed Jul 20 11:13:37 2011
(1 row)
```

```sql
SELECT pdf_creator('/tmp/pgintro.pdf');
```

```sql
            pdf_creator             
------------------------------------
 Microsoft® Office PowerPoint® 2007
(1 row)
```

```sql
SELECT pdf_metadata('/tmp/pgintro.pdf');
```

```sql
 pdf_metadata 
--------------
 
(1 row)
```

```sql
SELECT pdf_version('/tmp/pgintro.pdf');
```

```sql
 pdf_version 
-------------
 PDF-1.5
(1 row)
```

## Installation

Install [poppler](https://poppler.freedesktop.org) dependencies

**Linux**
```
sudo apt install -y libpoppler-glib-dev pkg-config
```

**Homebrew/MacOS**

```
brew install poppler pkgconf
```

```
cd /tmp
git clone https://github.com/Florents-Tselai/pgpdf.git
cd pgpdf
make
make install # may need sudo
```

After the installation, in a session:

```sql
CREATE EXTENSION pgpdf;
```

### Docker

Get the [Docker image](https://hub.docker.com/r/florents/pgpdf) with:

```sh
docker pull florents/pgpdf:pg17
```

This adds pgpdf to the [Postgres image](https://hub.docker.com/_/postgres) (replace `17` with your Postgres server version, and run it the same way).

Run the image in a container.

```sh
docker run --name pgpdf -p 5432:5432 -e POSTGRES_PASSWORD=pass florents/pgpdf:pg17
```

Through another terminal, connect to the running server (container).

```sh
PGPASSWORD=pass psql -h localhost -p 5432 -U postgres
```

> [!WARNING]
> Reading arbitrary binary data (PDF) into your database can pose security risks.
> Only use this for files you trust.
