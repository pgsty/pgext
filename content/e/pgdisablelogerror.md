---
title: "pgdisablelogerror"
linkTitle: "pgdisablelogerror"
description: "Disable selected SQLSTATE error codes from PostgreSQL server logging."
weight: 5260
categories: ["ADMIN"]
width: full
---

[**pgdisablelogerror**](https://github.com/fmbiete/pgdisablelogerror) : Disable selected SQLSTATE error codes from PostgreSQL server logging.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5260** | {{< badge content="pgdisablelogerror" link="https://github.com/fmbiete/pgdisablelogerror" >}} | {{< ext "pgdisablelogerror" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "BSD" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "logerrors" >}} {{< ext "pgauditlogtofile" >}} |

> [!Note] PGDG RPM and Pigsty DEB package fmbiete/pgdisablelogerror 1.0; control is relocatable=true and requires shared_preload_libraries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgdisablelogerror` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "pgdisablelogerror_18" "green" >}} {{< bg "17" "pgdisablelogerror_17" "green" >}} {{< bg "16" "pgdisablelogerror_16" "green" >}} {{< bg "15" "pgdisablelogerror_15" "green" >}} {{< bg "14" "pgdisablelogerror_14" "green" >}} | `pgdisablelogerror_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pgdisablelogerror" "green" >}} {{< bg "17" "postgresql-17-pgdisablelogerror" "green" >}} {{< bg "16" "postgresql-16-pgdisablelogerror" "green" >}} {{< bg "15" "postgresql-15-pgdisablelogerror" "green" >}} {{< bg "14" "postgresql-14-pgdisablelogerror" "green" >}} | `postgresql-$v-pgdisablelogerror` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0" "pgdisablelogerror_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgdisablelogerror : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgdisablelogerror : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdisablelogerror_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.4 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgdisablelogerror_18-1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgdisablelogerror_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_18-1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgdisablelogerror_18-1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.9 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgdisablelogerror` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-18-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdisablelogerror_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgdisablelogerror_17-1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgdisablelogerror_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_17-1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgdisablelogerror_17-1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.9 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.9 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgdisablelogerror` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 KiB | [postgresql-17-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-17-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdisablelogerror_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.2 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgdisablelogerror_16-1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgdisablelogerror_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_16-1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgdisablelogerror_16-1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.8 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.9 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgdisablelogerror` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 KiB | [postgresql-16-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-16-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdisablelogerror_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.2 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgdisablelogerror_15-1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgdisablelogerror_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_15-1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgdisablelogerror_15-1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.9 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.9 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgdisablelogerror` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 KiB | [postgresql-15-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-15-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdisablelogerror_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.2 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.0 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgdisablelogerror_14-1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgdisablelogerror_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [pgdisablelogerror_14-1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgdisablelogerror_14-1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.8 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.9 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgdisablelogerror` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 KiB | [postgresql-14-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdisablelogerror/postgresql-14-pgdisablelogerror_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fmbiete/pgdisablelogerror" title="Repository" icon="github" subtitle="github.com/fmbiete/pgdisablelogerror" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgdisablelogerror-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgdisablelogerror;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgdisablelogerror;		# install via package name, for the active PG version

pig install pgdisablelogerror -v 18;   # install for PG 18
pig install pgdisablelogerror -v 17;   # install for PG 17
pig install pgdisablelogerror -v 16;   # install for PG 16
pig install pgdisablelogerror -v 15;   # install for PG 15
pig install pgdisablelogerror -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/pgdisablelogerror';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgdisablelogerror;
```




## Usage

Sources: [README](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/README.md), [v1.0 release](https://github.com/fmbiete/pgdisablelogerror/releases/tag/v1.0), [control file](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/pgdisablelogerror.control)

`pgdisablelogerror` suppresses PostgreSQL server log entries for configured SQLSTATE error codes. It is useful when expected application errors, such as duplicate-key violations, are too noisy in the server log.

### Enable The Hook

Load the module and restart PostgreSQL:

```conf
shared_preload_libraries = 'pgdisablelogerror'
```

Create the extension once in the `postgres` database:

```sql
CREATE EXTENSION pgdisablelogerror;
```

### Configure SQLSTATE Codes

Set `pgdisablelogerror.sqlerrcode` to a comma-separated list of SQLSTATE codes:

```conf
pgdisablelogerror.sqlerrcode = '23505,23503'
```

An empty or NULL value disables suppression:

```conf
pgdisablelogerror.sqlerrcode = ''
```

To identify SQLSTATE values in normal PostgreSQL logs, add `%e` to `log_line_prefix`.

### Caveats

- Version 1.0 supports PostgreSQL 14-18.
- This extension affects logging, not error behavior. Clients still receive the original error.
- Use narrow SQLSTATE lists. Suppressing broad error classes can hide real operational problems.
