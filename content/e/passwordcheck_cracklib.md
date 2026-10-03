---
title: "passwordcheck_cracklib"
linkTitle: "passwordcheck_cracklib"
description: "Strengthen PostgreSQL user password checks with cracklib"
weight: 7000
categories: ["SEC"]
languages: ["C"]
licenses: ["LGPL-2.1"]
repos: ["PIGSTY"]
page_width: full
---

[**passwordcheck_cracklib**](https://github.com/devrimgunduz/passwordcheck_cracklib) : Strengthen PostgreSQL user password checks with cracklib


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7000** | {{< badge content="passwordcheck_cracklib" link="https://github.com/devrimgunduz/passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" >}} | `3.2.1` | {{< category "SEC" >}} | {{< license "LGPL-2.1" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_pwhash" >}} {{< ext "passwordcheck" >}} {{< ext "credcheck" >}} {{< ext "passwordpolicy" >}} {{< ext "chkpass" >}} {{< ext "pg_enigma" >}} {{< ext "column_encrypt" >}} |

> [!Note] Preload-only; requires cracklib dictionaries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `passwordcheck_cracklib` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "passwordcheck_cracklib_18" "green" >}} {{< bg "17" "passwordcheck_cracklib_17" "green" >}} {{< bg "16" "passwordcheck_cracklib_16" "green" >}} {{< bg "15" "passwordcheck_cracklib_15" "green" >}} {{< bg "14" "passwordcheck_cracklib_14" "green" >}} | `passwordcheck_cracklib_$v` | `cracklib-dicts` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "postgresql-18-passwordcheck-cracklib" "green" >}} {{< bg "17" "postgresql-17-passwordcheck-cracklib" "green" >}} {{< bg "16" "postgresql-16-passwordcheck-cracklib" "green" >}} {{< bg "15" "postgresql-15-passwordcheck-cracklib" "green" >}} {{< bg "14" "postgresql-14-passwordcheck-cracklib" "green" >}} | `postgresql-$v-passwordcheck-cracklib` | `cracklib-runtime`, `libcrack2` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 4" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.2.1" "passwordcheck_cracklib_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-passwordcheck-cracklib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-passwordcheck-cracklib : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_18` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.8 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el8.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.3 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el8.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.3 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el9.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [passwordcheck_cracklib_18-3.1.0-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/passwordcheck_cracklib_18-3.1.0-5PGDG.rhel9.8.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.7 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el9.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.3 KiB | [passwordcheck_cracklib_18-3.1.0-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/passwordcheck_cracklib_18-3.1.0-5PGDG.rhel9.8.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.4 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.8 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el10.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_18-3.1.0-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/passwordcheck_cracklib_18-3.1.0-5PGDG.rhel10.2.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_18` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_18-3.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/passwordcheck_cracklib_18-3.2.1-1PGSTY.el10.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 11.7 KiB | [passwordcheck_cracklib_18-3.1.0-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/passwordcheck_cracklib_18-3.1.0-5PGDG.rhel10.2.aarch64.rpm) |
| `passwordcheck_cracklib_18` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/passwordcheck_cracklib_18-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.0 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.8 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.4 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.6 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.4 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-passwordcheck-cracklib` | `3.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-18-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_17` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el8.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.2 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el8.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.2 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el9.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [passwordcheck_cracklib_17-3.1.0-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/passwordcheck_cracklib_17-3.1.0-5PGDG.rhel9.8.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el9.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.4 KiB | [passwordcheck_cracklib_17-3.1.0-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/passwordcheck_cracklib_17-3.1.0-5PGDG.rhel9.8.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.4 KiB | [passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/passwordcheck_cracklib_17-3.1.0-2PGDG.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el10.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_17-3.1.0-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/passwordcheck_cracklib_17-3.1.0-5PGDG.rhel10.2.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_17` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.9 KiB | [passwordcheck_cracklib_17-3.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/passwordcheck_cracklib_17-3.2.1-1PGSTY.el10.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 11.7 KiB | [passwordcheck_cracklib_17-3.1.0-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/passwordcheck_cracklib_17-3.1.0-5PGDG.rhel10.2.aarch64.rpm) |
| `passwordcheck_cracklib_17` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/passwordcheck_cracklib_17-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.8 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.5 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.5 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.4 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-passwordcheck-cracklib` | `3.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-17-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_16` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el8.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/passwordcheck_cracklib_16-3.0.0-1.rhel8.1.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.9 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el8.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/passwordcheck_cracklib_16-3.0.0-1.rhel8.1.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el9.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [passwordcheck_cracklib_16-3.1.0-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/passwordcheck_cracklib_16-3.1.0-5PGDG.rhel9.8.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.2 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/passwordcheck_cracklib_16-3.0.0-1.rhel9.1.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el9.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.4 KiB | [passwordcheck_cracklib_16-3.1.0-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/passwordcheck_cracklib_16-3.1.0-5PGDG.rhel9.8.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 10.9 KiB | [passwordcheck_cracklib_16-3.0.0-1.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/passwordcheck_cracklib_16-3.0.0-1.rhel9.1.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el10.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_16-3.1.0-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/passwordcheck_cracklib_16-3.1.0-5PGDG.rhel10.2.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_16` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.9 KiB | [passwordcheck_cracklib_16-3.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/passwordcheck_cracklib_16-3.2.1-1PGSTY.el10.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 11.7 KiB | [passwordcheck_cracklib_16-3.1.0-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/passwordcheck_cracklib_16-3.1.0-5PGDG.rhel10.2.aarch64.rpm) |
| `passwordcheck_cracklib_16` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/passwordcheck_cracklib_16-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.8 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.5 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.5 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.4 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-passwordcheck-cracklib` | `3.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-16-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_15` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.8 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el8.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/passwordcheck_cracklib_15-3.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el8.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/passwordcheck_cracklib_15-3.0.0-1.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el9.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [passwordcheck_cracklib_15-3.1.0-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/passwordcheck_cracklib_15-3.1.0-5PGDG.rhel9.8.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.1 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/passwordcheck_cracklib_15-3.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el9.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.3 KiB | [passwordcheck_cracklib_15-3.1.0-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/passwordcheck_cracklib_15-3.1.0-5PGDG.rhel9.8.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 10.8 KiB | [passwordcheck_cracklib_15-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/passwordcheck_cracklib_15-3.0.0-1.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el10.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_15-3.1.0-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/passwordcheck_cracklib_15-3.1.0-5PGDG.rhel10.2.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_15` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_15-3.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/passwordcheck_cracklib_15-3.2.1-1PGSTY.el10.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_15-3.1.0-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/passwordcheck_cracklib_15-3.1.0-5PGDG.rhel10.2.aarch64.rpm) |
| `passwordcheck_cracklib_15` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/passwordcheck_cracklib_15-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.8 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.5 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.5 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.4 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-passwordcheck-cracklib` | `3.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-15-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `passwordcheck_cracklib_14` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.8 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/passwordcheck_cracklib_14-3.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [passwordcheck_cracklib_14-2.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/passwordcheck_cracklib_14-2.0.0-1.rhel8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el8.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/passwordcheck_cracklib_14-3.0.0-1.rhel8.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el9.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [passwordcheck_cracklib_14-3.1.0-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordcheck_cracklib_14-3.1.0-5PGDG.rhel9.8.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.1 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordcheck_cracklib_14-3.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.7 KiB | [passwordcheck_cracklib_14-2.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/passwordcheck_cracklib_14-2.0.0-1.rhel9.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el9.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.4 KiB | [passwordcheck_cracklib_14-3.1.0-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/passwordcheck_cracklib_14-3.1.0-5PGDG.rhel9.8.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 10.8 KiB | [passwordcheck_cracklib_14-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/passwordcheck_cracklib_14-3.0.0-1.rhel9.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.8 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el10.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 11.6 KiB | [passwordcheck_cracklib_14-3.1.0-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/passwordcheck_cracklib_14-3.1.0-5PGDG.rhel10.2.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.x86_64.rpm) |
| `passwordcheck_cracklib_14` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.0 KiB | [passwordcheck_cracklib_14-3.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/passwordcheck_cracklib_14-3.2.1-1PGSTY.el10.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 11.7 KiB | [passwordcheck_cracklib_14-3.1.0-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/passwordcheck_cracklib_14-3.1.0-5PGDG.rhel10.2.aarch64.rpm) |
| `passwordcheck_cracklib_14` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.1 KiB | [passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/passwordcheck_cracklib_14-3.1.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.0 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.1 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.9 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.5 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.6 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.5 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-passwordcheck-cracklib` | `3.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/passwordcheck-cracklib/postgresql-14-passwordcheck-cracklib_3.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/devrimgunduz/passwordcheck_cracklib" title="Repository" icon="github" subtitle="github.com/devrimgunduz/passwordcheck_cracklib" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="passwordcheck_cracklib-3.2.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg passwordcheck_cracklib;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install passwordcheck_cracklib;		# install via package name, for the active PG version

pig install passwordcheck_cracklib -v 18;   # install for PG 18
pig install passwordcheck_cracklib -v 17;   # install for PG 17
pig install passwordcheck_cracklib -v 16;   # install for PG 16
pig install passwordcheck_cracklib -v 15;   # install for PG 15
pig install passwordcheck_cracklib -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [3.2.1 README](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/README.md)
- [3.2.1 password hook](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/passwordcheck_cracklib.c)
- [PostgreSQL passwordcheck manual](https://www.postgresql.org/docs/18/passwordcheck.html)

`passwordcheck_cracklib` checks passwords supplied through `CREATE ROLE` and `ALTER ROLE` with CrackLib. It is a server hook library, with no SQL extension objects.

### Enable the Hook

Add it to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib'
```

Do not run `CREATE EXTENSION passwordcheck_cracklib`. CrackLib's library and dictionary must be available to the PostgreSQL operating-system account.

### Password Checks

```sql
CREATE ROLE app_user LOGIN PASSWORD 'password123';
```

Weak plaintext passwords cause an error. The hook checks length, the relationship to the username, character composition, and the CrackLib dictionary. A password passing those checks is not a guarantee of resistance to every attack.

### Security Boundary

Dictionary checks require the plaintext password at password-change time. When a client supplies an already hashed password, the module cannot perform full strength checks; its remaining check is whether the password equals the username. Enforce the intended password-change path and protect the connection carrying plaintext passwords.

Existing passwords are not scanned retroactively. The hook also chains to a previously installed password-check hook; review other credential-policy libraries before loading them together.
