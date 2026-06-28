---
title: "plpgsql_wrap"
linkTitle: "plpgsql_wrap"
description: "Oracle WRAP-equivalent PL/pgSQL language handler storing AES-256-GCM encrypted procedure source."
weight: 9210
categories: ["SIM"]
width: full
---

[**plpgsql_wrap**](https://github.com/hexacluster/plpgsql_wrap/) : Oracle WRAP-equivalent PL/pgSQL language handler storing AES-256-GCM encrypted procedure source.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9210** | {{< badge content="plpgsql_wrap" link="https://github.com/hexacluster/plpgsql_wrap/" >}} | {{< ext "plpgsql_wrap" >}} | `1.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "orafce" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pgaudit" >}} |

> [!Note] PGDG RPM and Pigsty DEB package hexacluster/plpgsql_wrap 1.0; control requires plpgsql and superuser=true; links OpenSSL.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plpgsql_wrap` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "plpgsql_wrap_18" "green" >}} {{< bg "17" "plpgsql_wrap_17" "green" >}} {{< bg "16" "plpgsql_wrap_16" "green" >}} {{< bg "15" "plpgsql_wrap_15" "green" >}} {{< bg "14" "plpgsql_wrap_14" "green" >}} | `plpgsql_wrap_$v` | `openssl-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-plpgsql-wrap" "green" >}} {{< bg "17" "postgresql-17-plpgsql-wrap" "green" >}} {{< bg "16" "postgresql-16-plpgsql-wrap" "green" >}} {{< bg "15" "postgresql-15-plpgsql-wrap" "green" >}} {{< bg "14" "postgresql-14-plpgsql-wrap" "green" >}} | `postgresql-$v-plpgsql-wrap` | `libssl3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "plpgsql_wrap_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-plpgsql-wrap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-plpgsql-wrap : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_wrap_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_wrap_18-1.0-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_wrap_18-1.0-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_wrap_18-1.0-2PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_wrap_18-1.0-2PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_wrap_18-1.0-2PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.5 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_wrap_18-1.0-2PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_wrap_18-1.0-2PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_wrap_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.8 KiB | [plpgsql_wrap_18-1.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_wrap_18-1.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.7 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.8 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.8 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.4 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.1 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-plpgsql-wrap` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.9 KiB | [postgresql-18-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-18-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_wrap_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_wrap_17-1.0-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_wrap_17-1.0-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_wrap_17-1.0-2PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_wrap_17-1.0-2PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_wrap_17-1.0-2PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_wrap_17-1.0-2PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_wrap_17-1.0-2PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_wrap_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.8 KiB | [plpgsql_wrap_17-1.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_wrap_17-1.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.6 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.7 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.9 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.0 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.0 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-plpgsql-wrap` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.0 KiB | [postgresql-17-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-17-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_wrap_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_wrap_16-1.0-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_wrap_16-1.0-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_wrap_16-1.0-2PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_wrap_16-1.0-2PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_wrap_16-1.0-2PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_wrap_16-1.0-2PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_wrap_16-1.0-2PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_wrap_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.8 KiB | [plpgsql_wrap_16-1.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_wrap_16-1.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.6 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.7 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.5 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.6 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.0 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-plpgsql-wrap` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.9 KiB | [postgresql-16-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-16-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_wrap_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_wrap_15-1.0-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.1 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_wrap_15-1.0-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_wrap_15-1.0-2PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_wrap_15-1.0-2PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_wrap_15-1.0-2PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_wrap_15-1.0-2PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_wrap_15-1.0-2PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_wrap_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.8 KiB | [plpgsql_wrap_15-1.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_wrap_15-1.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.7 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.8 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.5 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.6 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.1 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-plpgsql-wrap` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.0 KiB | [postgresql-15-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-15-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_wrap_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_wrap_14-1.0-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_wrap_14-1.0-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_wrap_14-1.0-2PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_wrap_14-1.0-2PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_wrap_14-1.0-2PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_wrap_14-1.0-2PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.5 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_wrap_14-1.0-2PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.7 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_wrap_14-1.0-2PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_wrap_14-1.0-2PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_wrap_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.8 KiB | [plpgsql_wrap_14-1.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_wrap_14-1.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.9 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.6 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.7 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.4 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.5 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.1 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-plpgsql-wrap` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.9 KiB | [postgresql-14-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-wrap/postgresql-14-plpgsql-wrap_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hexacluster/plpgsql_wrap/" title="Repository" icon="github" subtitle="github.com/hexacluster/plpgsql_wrap/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plpgsql_wrap-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plpgsql_wrap;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plpgsql_wrap;		# install via package name, for the active PG version

pig install plpgsql_wrap -v 18;   # install for PG 18
pig install plpgsql_wrap -v 17;   # install for PG 17
pig install plpgsql_wrap -v 16;   # install for PG 16
pig install plpgsql_wrap -v 15;   # install for PG 15
pig install plpgsql_wrap -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plpgsql_wrap CASCADE; -- requires plpgsql
```
