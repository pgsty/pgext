---
title: "login_hook"
linkTitle: "login_hook"
description: "login_hook - hook to execute login_hook.login() at login time"
weight: 7360
categories: ["SEC"]
width: full
---

[**login_hook**](https://github.com/splendiddata/login_hook) : login_hook - hook to execute login_hook.login() at login time


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7360** | {{< badge content="login_hook" link="https://github.com/splendiddata/login_hook" >}} | {{< ext "login_hook" >}} | `1.7` | {{< category "SEC" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `login_hook` |
|   **See Also**    | {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "set_user" >}} {{< ext "pg_permissions" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "pgaudit" >}} {{< ext "auth_delay" >}} {{< ext "passwordcheck" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `login_hook` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.7` | {{< bg "18" "login_hook_18" "green" >}} {{< bg "17" "login_hook_17" "green" >}} {{< bg "16" "login_hook_16" "green" >}} {{< bg "15" "login_hook_15" "green" >}} {{< bg "14" "login_hook_14" "green" >}} | `login_hook_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "postgresql-18-login-hook" "green" >}} {{< bg "17" "postgresql-17-login-hook" "green" >}} {{< bg "16" "postgresql-16-login-hook" "green" >}} {{< bg "15" "postgresql-15-login-hook" "green" >}} {{< bg "14" "postgresql-14-login-hook" "green" >}} | `postgresql-$v-login-hook` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7" "login_hook_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7" "login_hook_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "login_hook_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-login-hook : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-login-hook : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `login_hook_18` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.7 KiB | [login_hook_18-1.7-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/login_hook_18-1.7-3PIGSTY.el8.x86_64.rpm) |
| `login_hook_18` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.9 KiB | [login_hook_18-1.7-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/login_hook_18-1.7-3PIGSTY.el8.aarch64.rpm) |
| `login_hook_18` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.7 KiB | [login_hook_18-1.7-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/login_hook_18-1.7-3PIGSTY.el9.x86_64.rpm) |
| `login_hook_18` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.7 KiB | [login_hook_18-1.7-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/login_hook_18-1.7-3PIGSTY.el9.aarch64.rpm) |
| `login_hook_18` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.7 KiB | [login_hook_18-1.7-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/login_hook_18-1.7-3PIGSTY.el10.x86_64.rpm) |
| `login_hook_18` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [login_hook_18-1.7-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/login_hook_18-1.7-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-login-hook` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.7 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-login-hook` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.6 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-login-hook` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.7 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-login-hook` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.6 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.2 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 29.1 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.6 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-login-hook` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.4 KiB | [postgresql-18-login-hook_1.7-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-18-login-hook_1.7-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `login_hook_17` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/login_hook_17-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.6 KiB | [login_hook_17-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/login_hook_17-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_17` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/login_hook_17-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.5 KiB | [login_hook_17-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/login_hook_17-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_17` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.3 KiB | [login_hook_17-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/login_hook_17-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.8 KiB | [login_hook_17-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/login_hook_17-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_17` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/login_hook_17-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.5 KiB | [login_hook_17-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/login_hook_17-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_17` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.6 KiB | [login_hook_17-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/login_hook_17-1.7-1PGDG.rhel10.x86_64.rpm) |
| `login_hook_17` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.6 KiB | [login_hook_17-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/login_hook_17-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-login-hook` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.7 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-login-hook` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.6 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-login-hook` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.6 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-login-hook` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.6 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.6 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 29.5 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.6 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.5 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-login-hook` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.4 KiB | [postgresql-17-login-hook_1.7-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-17-login-hook_1.7-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `login_hook_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.7 KiB | [login_hook_16-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [login_hook_16-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.1 KiB | [login_hook_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.6 KiB | [login_hook_16-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.7 KiB | [login_hook_16-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.5-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.4 KiB | [login_hook_16-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.5 KiB | [login_hook_16-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/login_hook_16-1.7-1PGDG.rhel10.x86_64.rpm) |
| `login_hook_16` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.5 KiB | [login_hook_16-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/login_hook_16-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-login-hook` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-login-hook` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-login-hook` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.2 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-login-hook` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.2 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.0 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.9 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.2 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.0 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-login-hook` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-16-login-hook_1.7-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-16-login-hook_1.7-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `login_hook_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.7 KiB | [login_hook_15-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [login_hook_15-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.1 KiB | [login_hook_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.6 KiB | [login_hook_15-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.7 KiB | [login_hook_15-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.5-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.4 KiB | [login_hook_15-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.5 KiB | [login_hook_15-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/login_hook_15-1.7-1PGDG.rhel10.x86_64.rpm) |
| `login_hook_15` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.5 KiB | [login_hook_15-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/login_hook_15-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-login-hook` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-login-hook` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-login-hook` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.2 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-login-hook` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.0 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.9 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.2 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.0 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-login-hook` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-15-login-hook_1.7-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-15-login-hook_1.7-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `login_hook_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.0 KiB | [login_hook_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.5 KiB | [login_hook_14-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.7 KiB | [login_hook_14-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.9 KiB | [login_hook_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [login_hook_14-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.7 KiB | [login_hook_14-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.2 KiB | [login_hook_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.7 KiB | [login_hook_14-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [login_hook_14-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.5-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.0 KiB | [login_hook_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [login_hook_14-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.4 KiB | [login_hook_14-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.5 KiB | [login_hook_14-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/login_hook_14-1.7-1PGDG.rhel10.x86_64.rpm) |
| `login_hook_14` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.5 KiB | [login_hook_14-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/login_hook_14-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-login-hook` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-login-hook` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.2 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-login-hook` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.4 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-login-hook` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.2 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.3 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 29.1 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.4 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.1 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.2 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-login-hook` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.0 KiB | [postgresql-14-login-hook_1.7-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/login-hook/postgresql-14-login-hook_1.7-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/splendiddata/login_hook" title="Repository" icon="github" subtitle="github.com/splendiddata/login_hook" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="login_hook-1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg login_hook;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install login_hook;		# install via package name, for the active PG version

pig install login_hook -v 18;   # install for PG 18
pig install login_hook -v 17;   # install for PG 17
pig install login_hook -v 16;   # install for PG 16
pig install login_hook -v 15;   # install for PG 15
pig install login_hook -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION login_hook;
```




## Usage

> [login_hook: Execute code on user login, comparable to Oracle's after logon trigger](https://github.com/splendiddata/login_hook)

`login_hook` allows executing custom PL/pgSQL code whenever a user logs into the database.

```sql
CREATE EXTENSION login_hook;
```

### Configuration

Add to `postgresql.conf`:

```ini
session_preload_libraries = 'login_hook'
```

### Creating the Login Function

Define a `login_hook.login()` function that will execute on every login:

```sql
CREATE OR REPLACE FUNCTION login_hook.login() RETURNS VOID LANGUAGE PLPGSQL AS $$
BEGIN
    IF NOT login_hook.is_executing_login_hook() THEN
        RAISE EXCEPTION 'Should only be invoked by login_hook';
    END IF;

    -- Your login logic here:
    RAISE NOTICE 'Hello %', current_user;

EXCEPTION
    WHEN OTHERS THEN
        RAISE LOG 'Error in login_hook.login(): %', SQLERRM;
END
$$;
GRANT EXECUTE ON FUNCTION login_hook.login() TO PUBLIC;
```

The `PUBLIC` grant is required because the function runs for every connecting user.

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `login_hook.is_executing_login_hook()` | `boolean` | Returns true only when called during login hook execution |
| `login_hook.get_login_hook_version()` | `text` | Returns compiled version of login_hook |
| `login_hook.login()` | `void` | User-provided function executed at login |

### Important Notes

- The function is NOT invoked for background processes or during recovery mode
- Handle all exceptions within the function -- failures will prevent normal users from logging in
- Superusers get a warning but can still log in when the function fails
- For PostgreSQL 17+, consider using the native login event trigger instead
