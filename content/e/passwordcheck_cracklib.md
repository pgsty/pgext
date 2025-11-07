---
title: "passwordcheck_cracklib"
linkTitle: "passwordcheck_cracklib"
description: "Strengthen PostgreSQL user password checks with cracklib"
weight: 7000
categories: ["SEC"]
width: full
---

Strengthen PostgreSQL user password checks with cracklib


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7000** | {{< badge content="passwordcheck_cracklib" link="https://github.com/devrimgunduz/passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" >}} | `3.1.0` | {{< category "SEC" >}} | {{< license "LGPL-2.1" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "pgaudit" >}} {{< ext "login_hook" >}} {{< ext "auth_delay" >}} {{< ext "set_user" >}} {{< ext "sepgsql" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/passwordcheck_cracklib" >}} | `3.1.0` | {{< bg "18" "passwordcheck_cracklib_18*" "green" >}} {{< bg "17" "passwordcheck_cracklib_17*" "green" >}} {{< bg "16" "passwordcheck_cracklib_16*" "green" >}} {{< bg "15" "passwordcheck_cracklib_15*" "green" >}} {{< bg "14" "passwordcheck_cracklib_14*" "green" >}} {{< bg "13" "passwordcheck_cracklib_13*" "green" >}} | `passwordcheck_cracklib_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/passwordcheck_cracklib" >}} | `3.1.0` | {{< bg "18" "postgresql-18-passwordcheck-cracklib" "red" >}} {{< bg "17" "postgresql-17-passwordcheck-cracklib" "green" >}} {{< bg "16" "postgresql-16-passwordcheck-cracklib" "green" >}} {{< bg "15" "postgresql-15-passwordcheck-cracklib" "green" >}} {{< bg "14" "postgresql-14-passwordcheck-cracklib" "green" >}} {{< bg "13" "postgresql-13-passwordcheck-cracklib" "green" >}} | `postgresql-$v-passwordcheck-cracklib` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0.0" "passwordcheck_cracklib_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.1.0" "passwordcheck_cracklib_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-passwordcheck-cracklib : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-passwordcheck-cracklib : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-passwordcheck-cracklib : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.0" "postgresql-13-passwordcheck-cracklib : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_18` | 3.1.0 | `el8.x86_64` | pgdg | 12.3 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_18` | 3.1.0 | `el8.aarch64` | pgdg | 12.3 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_18` | 3.1.0 | `el9.x86_64` | pgdg | 11.7 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_18` | 3.1.0 | `el9.aarch64` | pgdg | 11.4 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_18` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_18` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.6 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.8 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.5 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.1 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.6 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.3 KiB | [postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_17` | 3.1.0 | `el8.x86_64` | pgdg | 12.2 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_17` | 3.1.0 | `el8.aarch64` | pgdg | 12.2 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_17` | 3.1.0 | `el9.x86_64` | pgdg | 11.6 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_17` | 3.1.0 | `el9.aarch64` | pgdg | 11.4 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_17` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_17` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.5 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.7 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.6 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.2 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.5 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.2 KiB | [postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_16` | 3.0.0 | `el8.x86_64` | pgdg | 11.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/passwordcheck_cracklib_16-3.0.0-1.rhel8.1.x86_64.rpm) |
| `passwordcheck_cracklib_16` | 3.0.0 | `el8.aarch64` | pgdg | 11.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/passwordcheck_cracklib_16-3.0.0-1.rhel8.1.aarch64.rpm) |
| `passwordcheck_cracklib_16` | 3.0.0 | `el9.x86_64` | pgdg | 11.2 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/passwordcheck_cracklib_16-3.0.0-1.rhel9.1.x86_64.rpm) |
| `passwordcheck_cracklib_16` | 3.0.0 | `el9.aarch64` | pgdg | 10.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/passwordcheck_cracklib_16-3.0.0-1.rhel9.1.aarch64.rpm) |
| `passwordcheck_cracklib_16` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_16` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.5 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.7 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.6 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.2 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.5 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.3 KiB | [postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_15` | 3.0.0 | `el8.x86_64` | pgdg | 11.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/passwordcheck_cracklib_15-3.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_15` | 3.0.0 | `el8.aarch64` | pgdg | 11.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/passwordcheck_cracklib_15-3.0.0-1.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_15` | 3.0.0 | `el9.x86_64` | pgdg | 11.1 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/passwordcheck_cracklib_15-3.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_15` | 3.0.0 | `el9.aarch64` | pgdg | 10.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/passwordcheck_cracklib_15-3.0.0-1.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_15` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_15` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.5 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.7 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.6 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.2 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.5 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.2 KiB | [postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_14` | 3.0.0 | `el8.x86_64` | pgdg | 11.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/passwordcheck_cracklib_14-3.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | 2.0.0 | `el8.x86_64` | pgdg | 17.4 KiB | [passwordcheck_cracklib_14-2.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/passwordcheck_cracklib_14-2.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | 3.0.0 | `el8.aarch64` | pgdg | 11.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/passwordcheck_cracklib_14-3.0.0-1.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_14` | 3.0.0 | `el9.x86_64` | pgdg | 11.1 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordcheck_cracklib_14-3.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_14` | 2.0.0 | `el9.x86_64` | pgdg | 16.7 KiB | [passwordcheck_cracklib_14-2.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordcheck_cracklib_14-2.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_14` | 3.0.0 | `el9.aarch64` | pgdg | 10.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/passwordcheck_cracklib_14-3.0.0-1.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_14` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_14` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.6 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.7 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.6 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.3 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.6 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.3 KiB | [postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_13` | 3.0.0 | `el8.x86_64` | pgdg | 11.7 KiB | [passwordcheck_cracklib_13-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/passwordcheck_cracklib_13-3.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_13` | 2.0.0 | `el8.x86_64` | pgdg | 17.1 KiB | [passwordcheck_cracklib_13-2.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/passwordcheck_cracklib_13-2.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_13` | 3.0.0 | `el8.aarch64` | pgdg | 11.8 KiB | [passwordcheck_cracklib_13-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/passwordcheck_cracklib_13-3.0.0-1.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_13` | 3.0.0 | `el9.x86_64` | pgdg | 11.1 KiB | [passwordcheck_cracklib_13-3.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/passwordcheck_cracklib_13-3.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_13` | 2.0.0 | `el9.x86_64` | pgdg | 16.6 KiB | [passwordcheck_cracklib_13-2.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/passwordcheck_cracklib_13-2.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_13` | 3.0.0 | `el9.aarch64` | pgdg | 10.8 KiB | [passwordcheck_cracklib_13-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/passwordcheck_cracklib_13-3.0.0-1.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_13` | 3.1.0 | `el10.x86_64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_13-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/passwordcheck_cracklib_13-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_13` | 3.1.0 | `el10.aarch64` | pgdg | 12.1 KiB | [passwordcheck_cracklib_13-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/passwordcheck_cracklib_13-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `d12.x86_64` | pigsty | 17.5 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `d12.aarch64` | pigsty | 17.6 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `u22.x86_64` | pigsty | 18.4 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `u22.aarch64` | pigsty | 18.2 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `u24.x86_64` | pigsty | 18.5 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-passwordcheck-cracklib` | 3.1.0 | `u24.aarch64` | pigsty | 18.1 KiB | [postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-13-passwordcheck-cracklib_3.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/devrimgunduz/passwordcheck_cracklib" title="Repository" icon="github" subtitle="github.com/devrimgunduz/passwordcheck_cracklib" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="passwordcheck_cracklib-3.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get passwordcheck_cracklib; # get passwordcheck_cracklib source code
pig build dep passwordcheck_cracklib; # install build dependencies
pig build pkg passwordcheck_cracklib; # build extension rpm or deb
pig build ext passwordcheck_cracklib; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install passwordcheck_cracklib; # install by extension name, for the current active PG version
pig ext install passwordcheck_cracklib; # install via package alias, for the active PG version
pig ext install passwordcheck_cracklib -v 18;   # install for PG 18
pig ext install passwordcheck_cracklib -v 17;   # install for PG 17
pig ext install passwordcheck_cracklib -v 16;   # install for PG 16
pig ext install passwordcheck_cracklib -v 15;   # install for PG 15
pig ext install passwordcheck_cracklib -v 14;   # install for PG 14
pig ext install passwordcheck_cracklib -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION passwordcheck_cracklib;
```

