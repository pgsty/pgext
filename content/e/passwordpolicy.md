---
title: "passwordpolicy"
linkTitle: "passwordpolicy"
description: "Dynamically configurable PostgreSQL password complexity checks."
weight: 7040
categories: ["SEC"]
width: full
---

[**passwordpolicy**](https://github.com/fmbiete/passwordpolicy) : Dynamically configurable PostgreSQL password complexity checks.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7040** | {{< badge content="passwordpolicy" link="https://github.com/fmbiete/passwordpolicy" >}} | {{< ext "passwordpolicy" >}} | `2.0.5` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "credcheck" >}} |

> [!Note] PGDG RPM and Pigsty DEB package fmbiete/passwordpolicy 2.0.5; requires shared_preload_libraries and cracklib runtime.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `passwordpolicy` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.5` | {{< bg "18" "passwordpolicy_18" "green" >}} {{< bg "17" "passwordpolicy_17" "green" >}} {{< bg "16" "passwordpolicy_16" "green" >}} {{< bg "15" "passwordpolicy_15" "green" >}} {{< bg "14" "passwordpolicy_14" "green" >}} | `passwordpolicy_$v` | `cracklib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.5` | {{< bg "18" "postgresql-18-passwordpolicy" "green" >}} {{< bg "17" "postgresql-17-passwordpolicy" "green" >}} {{< bg "16" "postgresql-16-passwordpolicy" "green" >}} {{< bg "15" "postgresql-15-passwordpolicy" "green" >}} {{< bg "14" "postgresql-14-passwordpolicy" "green" >}} | `postgresql-$v-passwordpolicy` | `cracklib-runtime`, `libcrack2` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "passwordpolicy_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-18-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-17-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-16-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-15-passwordpolicy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.5" "postgresql-14-passwordpolicy : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordpolicy_18` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.2 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/passwordpolicy_18-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `passwordpolicy_18` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.1 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/passwordpolicy_18-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `passwordpolicy_18` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/passwordpolicy_18-2.0.5-1PGDG.rhel9.8.x86_64.rpm) |
| `passwordpolicy_18` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.5 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/passwordpolicy_18-2.0.5-1PGDG.rhel9.8.aarch64.rpm) |
| `passwordpolicy_18` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.6 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/passwordpolicy_18-2.0.5-1PGDG.rhel10.2.x86_64.rpm) |
| `passwordpolicy_18` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [passwordpolicy_18-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/passwordpolicy_18-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.9 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.3 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.8 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.4 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 55.6 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 55.0 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.2 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.1 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.9 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-passwordpolicy` | `2.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.8 KiB | [postgresql-18-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-18-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordpolicy_17` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.2 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/passwordpolicy_17-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `passwordpolicy_17` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.1 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/passwordpolicy_17-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `passwordpolicy_17` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.6 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/passwordpolicy_17-2.0.5-1PGDG.rhel9.8.x86_64.rpm) |
| `passwordpolicy_17` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.6 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/passwordpolicy_17-2.0.5-1PGDG.rhel9.8.aarch64.rpm) |
| `passwordpolicy_17` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/passwordpolicy_17-2.0.5-1PGDG.rhel10.2.x86_64.rpm) |
| `passwordpolicy_17` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.8 KiB | [passwordpolicy_17-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/passwordpolicy_17-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.8 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.2 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.8 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.2 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.1 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.6 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.3 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.0 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.9 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-passwordpolicy` | `2.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.8 KiB | [postgresql-17-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-17-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordpolicy_16` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.2 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/passwordpolicy_16-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `passwordpolicy_16` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.1 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/passwordpolicy_16-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `passwordpolicy_16` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.6 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/passwordpolicy_16-2.0.5-1PGDG.rhel9.8.x86_64.rpm) |
| `passwordpolicy_16` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.6 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/passwordpolicy_16-2.0.5-1PGDG.rhel9.8.aarch64.rpm) |
| `passwordpolicy_16` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/passwordpolicy_16-2.0.5-1PGDG.rhel10.2.x86_64.rpm) |
| `passwordpolicy_16` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.8 KiB | [passwordpolicy_16-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/passwordpolicy_16-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.8 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.2 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.8 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.2 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.0 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.6 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.2 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.1 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.9 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-passwordpolicy` | `2.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.8 KiB | [postgresql-16-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-16-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordpolicy_15` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/passwordpolicy_15-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `passwordpolicy_15` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.4 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/passwordpolicy_15-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `passwordpolicy_15` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.2 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/passwordpolicy_15-2.0.5-1PGDG.rhel9.8.x86_64.rpm) |
| `passwordpolicy_15` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.1 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/passwordpolicy_15-2.0.5-1PGDG.rhel9.8.aarch64.rpm) |
| `passwordpolicy_15` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.4 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/passwordpolicy_15-2.0.5-1PGDG.rhel10.2.x86_64.rpm) |
| `passwordpolicy_15` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.4 KiB | [passwordpolicy_15-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/passwordpolicy_15-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.2 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.5 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.3 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.6 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.7 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 62.3 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.8 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.6 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.6 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-passwordpolicy` | `2.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 54.2 KiB | [postgresql-15-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-15-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordpolicy_14` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/passwordpolicy_14-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `passwordpolicy_14` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.6 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/passwordpolicy_14-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `passwordpolicy_14` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.2 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordpolicy_14-2.0.5-1PGDG.rhel9.8.x86_64.rpm) |
| `passwordpolicy_14` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.3 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/passwordpolicy_14-2.0.5-1PGDG.rhel9.8.aarch64.rpm) |
| `passwordpolicy_14` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/passwordpolicy_14-2.0.5-1PGDG.rhel10.2.x86_64.rpm) |
| `passwordpolicy_14` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [passwordpolicy_14-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/passwordpolicy_14-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.0 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.7 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.1 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.7 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.4 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 62.3 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.6 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.7 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.0 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-passwordpolicy` | `2.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 54.4 KiB | [postgresql-14-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordpolicy/postgresql-14-passwordpolicy_2.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fmbiete/passwordpolicy" title="Repository" icon="github" subtitle="github.com/fmbiete/passwordpolicy" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="passwordpolicy-2.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg passwordpolicy;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install passwordpolicy;		# install via package name, for the active PG version

pig install passwordpolicy -v 18;   # install for PG 18
pig install passwordpolicy -v 17;   # install for PG 17
pig install passwordpolicy -v 16;   # install for PG 16
pig install passwordpolicy -v 15;   # install for PG 15
pig install passwordpolicy -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/passwordpolicy';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION passwordpolicy;
```
