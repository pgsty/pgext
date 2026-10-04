---
title: "sslutils"
linkTitle: "sslutils"
description: "A Postgres extension for managing SSL certificates through SQL"
weight: 7410
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**sslutils**](https://github.com/EnterpriseDB/sslutils) : A Postgres extension for managing SSL certificates through SQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7410** | {{< badge content="sslutils" link="https://github.com/EnterpriseDB/sslutils" >}} | {{< ext "sslutils" >}} | `1.4.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "sslinfo" >}} {{< ext "pg_oidc_validator" >}} {{< ext "oidc_validator" >}} {{< ext "pguecc" >}} {{< ext "pg_session_jwt" >}} {{< ext "pgjwt" >}} {{< ext "pgsodium" >}} {{< ext "login_hook" >}} {{< ext "libx509pq" >}} |

> [!Note] PGDG packages remain absent for PG18 on EL8 x86_64.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `sslutils` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "sslutils_18" "green" >}} {{< bg "17" "sslutils_17" "green" >}} {{< bg "16" "sslutils_16" "green" >}} {{< bg "15" "sslutils_15" "green" >}} {{< bg "14" "sslutils_14" "green" >}} | `sslutils_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "postgresql-18-sslutils" "green" >}} {{< bg "17" "postgresql-17-sslutils" "green" >}} {{< bg "16" "postgresql-16-sslutils" "green" >}} {{< bg "15" "postgresql-15-sslutils" "green" >}} {{< bg "14" "postgresql-14-sslutils" "green" >}} | `postgresql-$v-sslutils` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4" "sslutils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4" "sslutils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "sslutils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4" "sslutils_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4" "sslutils_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4" "sslutils_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4" "sslutils_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "sslutils_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.4" "postgresql-18-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-17-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-16-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-15-sslutils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "postgresql-14-sslutils : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sslutils_18` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.6 KiB | [sslutils_18-1.4-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_18-1.4-3PIGSTY.el8.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.7 KiB | [sslutils_18-1.4-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_18-1.4-3PIGSTY.el8.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [sslutils_18-1.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sslutils_18-1.4-4PGDG.rhel9.8.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [sslutils_18-1.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_18-1.4-2PIGSTY.el9.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [sslutils_18-1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sslutils_18-1.4-2PGDG.rhel9.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [sslutils_18-1.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sslutils_18-1.4-4PGDG.rhel9.8.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.8 KiB | [sslutils_18-1.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_18-1.4-2PIGSTY.el9.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [sslutils_18-1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sslutils_18-1.4-2PGDG.rhel9.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [sslutils_18-1.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/sslutils_18-1.4-4PGDG.rhel10.2.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [sslutils_18-1.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sslutils_18-1.4-2PIGSTY.el10.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.6 KiB | [sslutils_18-1.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/sslutils_18-1.4-2PGDG.rhel10.x86_64.rpm) |
| `sslutils_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [sslutils_18-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sslutils_18-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.7 KiB | [sslutils_18-1.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sslutils_18-1.4-2PIGSTY.el10.aarch64.rpm) |
| `sslutils_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [sslutils_18-1.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sslutils_18-1.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-sslutils` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.2 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-sslutils` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.6 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-sslutils` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-sslutils` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.1 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.1 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.9 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.4 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.1 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 40.0 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-sslutils` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-18-sslutils_1.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-18-sslutils_1.4-2PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sslutils_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [sslutils_17-1.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_17-1.4-2PIGSTY.el8.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [sslutils_17-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/sslutils_17-1.4-1PGDG.rhel8.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.6 KiB | [sslutils_17-1.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_17-1.4-2PIGSTY.el8.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [sslutils_17-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/sslutils_17-1.4-1PGDG.rhel8.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [sslutils_17-1.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sslutils_17-1.4-4PGDG.rhel9.8.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [sslutils_17-1.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_17-1.4-2PIGSTY.el9.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [sslutils_17-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sslutils_17-1.4-1PGDG.rhel9.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [sslutils_17-1.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sslutils_17-1.4-4PGDG.rhel9.8.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.8 KiB | [sslutils_17-1.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_17-1.4-2PIGSTY.el9.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [sslutils_17-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sslutils_17-1.4-1PGDG.rhel9.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [sslutils_17-1.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/sslutils_17-1.4-4PGDG.rhel10.2.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [sslutils_17-1.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sslutils_17-1.4-2PIGSTY.el10.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.6 KiB | [sslutils_17-1.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/sslutils_17-1.4-2PGDG.rhel10.x86_64.rpm) |
| `sslutils_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [sslutils_17-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sslutils_17-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.7 KiB | [sslutils_17-1.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sslutils_17-1.4-2PIGSTY.el10.aarch64.rpm) |
| `sslutils_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [sslutils_17-1.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sslutils_17-1.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-sslutils` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.9 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-sslutils` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.5 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-sslutils` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.5 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-sslutils` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.1 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.8 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.7 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.4 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.1 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 40.0 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-sslutils` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-17-sslutils_1.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-17-sslutils_1.4-2PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sslutils_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [sslutils_16-1.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_16-1.4-2PIGSTY.el8.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [sslutils_16-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/sslutils_16-1.4-1PGDG.rhel8.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.6 KiB | [sslutils_16-1.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_16-1.4-2PIGSTY.el8.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [sslutils_16-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/sslutils_16-1.4-1PGDG.rhel8.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [sslutils_16-1.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sslutils_16-1.4-4PGDG.rhel9.8.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [sslutils_16-1.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_16-1.4-2PIGSTY.el9.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [sslutils_16-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sslutils_16-1.4-1PGDG.rhel9.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [sslutils_16-1.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sslutils_16-1.4-4PGDG.rhel9.8.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.8 KiB | [sslutils_16-1.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_16-1.4-2PIGSTY.el9.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [sslutils_16-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sslutils_16-1.4-1PGDG.rhel9.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [sslutils_16-1.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/sslutils_16-1.4-4PGDG.rhel10.2.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [sslutils_16-1.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sslutils_16-1.4-2PIGSTY.el10.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [sslutils_16-1.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/sslutils_16-1.4-2PGDG.rhel10.x86_64.rpm) |
| `sslutils_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [sslutils_16-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sslutils_16-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.7 KiB | [sslutils_16-1.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sslutils_16-1.4-2PIGSTY.el10.aarch64.rpm) |
| `sslutils_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [sslutils_16-1.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sslutils_16-1.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-sslutils` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.1 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-sslutils` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.5 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-sslutils` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.5 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-sslutils` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.1 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.8 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.7 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.4 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.1 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 40.0 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-sslutils` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-16-sslutils_1.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-16-sslutils_1.4-2PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sslutils_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.6 KiB | [sslutils_15-1.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_15-1.4-2PIGSTY.el8.x86_64.rpm) |
| `sslutils_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [sslutils_15-1.3-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sslutils_15-1.3-4.rhel8.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.6 KiB | [sslutils_15-1.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_15-1.4-2PIGSTY.el8.aarch64.rpm) |
| `sslutils_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [sslutils_15-1.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sslutils_15-1.4-4PGDG.rhel9.8.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [sslutils_15-1.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_15-1.4-2PIGSTY.el9.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [sslutils_15-1.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sslutils_15-1.4-4PGDG.rhel9.8.aarch64.rpm) |
| `sslutils_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.9 KiB | [sslutils_15-1.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_15-1.4-2PIGSTY.el9.aarch64.rpm) |
| `sslutils_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.4 KiB | [sslutils_15-1.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/sslutils_15-1.4-4PGDG.rhel10.2.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [sslutils_15-1.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sslutils_15-1.4-2PIGSTY.el10.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [sslutils_15-1.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/sslutils_15-1.4-2PGDG.rhel10.x86_64.rpm) |
| `sslutils_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [sslutils_15-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sslutils_15-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `sslutils_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.8 KiB | [sslutils_15-1.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sslutils_15-1.4-2PIGSTY.el10.aarch64.rpm) |
| `sslutils_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [sslutils_15-1.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sslutils_15-1.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-sslutils` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.0 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-sslutils` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.6 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-sslutils` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.5 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-sslutils` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.1 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.9 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.7 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.2 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 40.0 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-sslutils` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-15-sslutils_1.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-15-sslutils_1.4-2PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sslutils_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [sslutils_14-1.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_14-1.4-2PIGSTY.el8.x86_64.rpm) |
| `sslutils_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.9 KiB | [sslutils_14-1.3-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sslutils_14-1.3-4.rhel8.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.6 KiB | [sslutils_14-1.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_14-1.4-2PIGSTY.el8.aarch64.rpm) |
| `sslutils_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [sslutils_14-1.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sslutils_14-1.4-4PGDG.rhel9.8.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [sslutils_14-1.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_14-1.4-2PIGSTY.el9.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [sslutils_14-1.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sslutils_14-1.4-4PGDG.rhel9.8.aarch64.rpm) |
| `sslutils_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.8 KiB | [sslutils_14-1.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_14-1.4-2PIGSTY.el9.aarch64.rpm) |
| `sslutils_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.4 KiB | [sslutils_14-1.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/sslutils_14-1.4-4PGDG.rhel10.2.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [sslutils_14-1.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sslutils_14-1.4-2PIGSTY.el10.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.8 KiB | [sslutils_14-1.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/sslutils_14-1.4-2PGDG.rhel10.x86_64.rpm) |
| `sslutils_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [sslutils_14-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/sslutils_14-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `sslutils_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.8 KiB | [sslutils_14-1.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sslutils_14-1.4-2PIGSTY.el10.aarch64.rpm) |
| `sslutils_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [sslutils_14-1.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/sslutils_14-1.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-sslutils` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.0 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-sslutils` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.5 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-sslutils` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.5 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-sslutils` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.0 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.8 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.6 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.2 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 40.0 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-sslutils` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-14-sslutils_1.4-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sslutils/postgresql-14-sslutils_1.4-2PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/sslutils" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/sslutils" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="sslutils-1.4.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg sslutils;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install sslutils;		# install via package name, for the active PG version

pig install sslutils -v 18;   # install for PG 18
pig install sslutils -v 17;   # install for PG 17
pig install sslutils -v 16;   # install for PG 16
pig install sslutils -v 15;   # install for PG 15
pig install sslutils -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION sslutils;
```

## Usage

Sources:

- [1.4.1 README](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/README.sslutils)
- [1.4.1 SQL API](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils--1.4.1.sql)
- [1.4.1 control file](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.control)
- [File-access checks](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.c)

`sslutils` provides SSL certificate-generation and revocation helpers for Postgres Enterprise Manager and other controlled administration workflows. Some functions read or modify files on the database server.

### Install and Inspect

```sql
CREATE EXTENSION sslutils;
SELECT sslutils_version();
SELECT openssl_get_crt_expiry_date('server.crt');
```

A superuser must install this non-relocatable extension. No preload or PostgreSQL restart is required to call the functions; deploying generated certificates is a separate server configuration operation.

### Important Functions

| Function | Purpose |
|---|---|
| `openssl_rsa_generate_key(bits)` | Return an RSA private key as text |
| `openssl_rsa_key_to_csr(key, cn, country, state, location, unit, email)` | Return a CSR |
| `openssl_csr_to_crt(csr, ca_cert_path, private_key_path, days)` | Sign a CSR using server-side certificate and key files |
| `openssl_rsa_generate_crl(ca_cert_path, ca_key_path, days)` | Return a certificate revocation list |
| `openssl_is_crt_expire_on(cert_path, at_time)` | Check certificate expiry; returns 1, -1 or 0 |
| `openssl_get_crt_expiry_date(cert_path)` | Return the expiry timestamp |
| `openssl_revoke_certificate(cert_pem, crl_path, days)` | Revoke a certificate and regenerate the CRL |

For self-signed generation, the certificate-path argument to the signing function is NULL and the key path selects the signing key. The optional validity defaults to 3650 days. File paths refer to the PostgreSQL server, not the SQL client's machine.

### Protect Keys and Server Files

Certificate-inspection functions require membership in `pg_read_server_files` on PostgreSQL 11 and later; revocation checks both that role and `pg_write_server_files`. These roles provide broad server-file access, so grant them only to trusted administrators and restrict function execution as needed.

Private keys returned as SQL text can enter query results, logs, shell history or application traces. Protect their handling and storage. This extension does not provide SSL session-inspection getters or automatically configure PostgreSQL's TLS settings.

### File-Access Changes in 1.4.1

Certificate inspection and CA signing paths are checked against the server data directory; place the input files there and use server-relative paths where suitable. Revocation requires `sslutils.revoke_certificate_crl_paths`, a comma-separated allowlist of CRL output path prefixes configured in postgresql.conf and activated by reload. The revocation check compares lexical string prefixes without resolving canonical directory containment; configure narrow prefixes and protect the output location.

The revocation implementation accepts PEM certificate text as its first argument, despite older README and SQL comments describing a certificate path. It reads the server-side CA files `ca_certificate.crt` and `ca_key.key`, and creates or appends to `revoke_cert.db`. Prepare the CA files before using it. Do not infer that a broad server-file role bypasses these extension-specific checks.
