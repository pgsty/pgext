---
title: "login_hook"
linkTitle: "login_hook"
description: "login_hook - hook to execute login_hook.login() at login time"
weight: 7150
categories: ["Sec"]
width: full
---

login_hook - hook to execute login_hook.login() at login time

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7150** | {{< badge content="login_hook" link="https://github.com/splendiddata/login_hook" >}} | {{< ext "login_hook" "login_hook" >}} | `1.7` | {{< category "SEC" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "set_user" >}} {{< ext "pg_permissions" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "pgaudit" >}} {{< ext "auth_delay" >}} {{< ext "passwordcheck" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/login_hook" >}} | `1.7` | {{< badge content="18" color="red" alt="login_hook_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `login_hook_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/login_hook" >}} | `1.6` | {{< badge content="18" color="red" alt="postgresql-18-login-hook" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-login-hook` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "login_hook_18" >}}     | {{< pkg "login_hook_17" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/login_hook_17-1.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "login_hook_16" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "login_hook_15" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "login_hook_14" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.7-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "login_hook_18" >}}     | {{< pkg "login_hook_17" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/login_hook_17-1.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "login_hook_16" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "login_hook_15" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "login_hook_14" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.7-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "login_hook_18" >}}     | {{< pkg "login_hook_17" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/login_hook_17-1.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "login_hook_16" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "login_hook_15" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "login_hook_14" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.7-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "login_hook_18" >}}     | {{< pkg "login_hook_17" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/login_hook_17-1.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "login_hook_16" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "login_hook_15" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "login_hook_14" "1.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.7-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-login-hook" >}}     | {{< pkg "postgresql-17-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-login-hook" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `login_hook_17` | 1.7 | `el8.aarch64` | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/login_hook_17-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_17` | 1.7 | `el8.x86_64` | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/login_hook_17-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_17` | 1.6 | `el8.aarch64` | pgdg | 17.5 KiB | [login_hook_17-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/login_hook_17-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_17` | 1.6 | `el8.x86_64` | pgdg | 17.6 KiB | [login_hook_17-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/login_hook_17-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_17` | 1.7 | `el9.aarch64` | pgdg | 18.1 KiB | [login_hook_17-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/login_hook_17-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_17` | 1.7 | `el9.x86_64` | pgdg | 18.3 KiB | [login_hook_17-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/login_hook_17-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_17` | 1.6 | `el9.x86_64` | pgdg | 17.8 KiB | [login_hook_17-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/login_hook_17-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_17` | 1.6 | `el9.aarch64` | pgdg | 17.5 KiB | [login_hook_17-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/login_hook_17-1.6-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-login-hook` | 1.6 | `d12.aarch64` | pigsty | 27.6 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-login-hook` | 1.6 | `d12.x86_64` | pigsty | 28.1 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-login-hook` | 1.6 | `u22.x86_64` | pigsty | 29.3 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-login-hook` | 1.6 | `u22.aarch64` | pigsty | 29.2 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-login-hook` | 1.6 | `u24.x86_64` | pigsty | 28.7 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-login-hook` | 1.6 | `u24.aarch64` | pigsty | 28.1 KiB | [postgresql-17-login-hook_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-17-login-hook_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `login_hook_16` | 1.7 | `el8.aarch64` | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | 1.7 | `el8.x86_64` | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | 1.6 | `el8.aarch64` | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | 1.6 | `el8.x86_64` | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | 1.5 | `el8.x86_64` | pgdg | 16.7 KiB | [login_hook_16-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/login_hook_16-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_16` | 1.5 | `el8.aarch64` | pgdg | 16.6 KiB | [login_hook_16-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/login_hook_16-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_16` | 1.7 | `el9.aarch64` | pgdg | 17.9 KiB | [login_hook_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | 1.7 | `el9.x86_64` | pgdg | 18.1 KiB | [login_hook_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_16` | 1.6 | `el9.aarch64` | pgdg | 17.4 KiB | [login_hook_16-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | 1.6 | `el9.x86_64` | pgdg | 17.6 KiB | [login_hook_16-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_16` | 1.5 | `el9.aarch64` | pgdg | 16.4 KiB | [login_hook_16-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/login_hook_16-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_16` | 1.5 | `el9.x86_64` | pgdg | 16.7 KiB | [login_hook_16-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/login_hook_16-1.5-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-login-hook` | 1.6 | `d12.x86_64` | pigsty | 27.5 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-login-hook` | 1.6 | `d12.aarch64` | pigsty | 27.1 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-login-hook` | 1.6 | `u22.aarch64` | pigsty | 28.6 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-login-hook` | 1.6 | `u22.x86_64` | pigsty | 28.7 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-login-hook` | 1.6 | `u24.aarch64` | pigsty | 27.5 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-login-hook` | 1.6 | `u24.x86_64` | pigsty | 28.1 KiB | [postgresql-16-login-hook_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-16-login-hook_1.6-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `login_hook_15` | 1.7 | `el8.aarch64` | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | 1.7 | `el8.x86_64` | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | 1.6 | `el8.aarch64` | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | 1.6 | `el8.x86_64` | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | 1.5 | `el8.aarch64` | pgdg | 16.6 KiB | [login_hook_15-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/login_hook_15-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_15` | 1.5 | `el8.x86_64` | pgdg | 16.7 KiB | [login_hook_15-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/login_hook_15-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_15` | 1.7 | `el9.x86_64` | pgdg | 18.1 KiB | [login_hook_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_15` | 1.7 | `el9.aarch64` | pgdg | 17.9 KiB | [login_hook_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | 1.6 | `el9.x86_64` | pgdg | 17.6 KiB | [login_hook_15-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_15` | 1.6 | `el9.aarch64` | pgdg | 17.4 KiB | [login_hook_15-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | 1.5 | `el9.aarch64` | pgdg | 16.4 KiB | [login_hook_15-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/login_hook_15-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_15` | 1.5 | `el9.x86_64` | pgdg | 16.7 KiB | [login_hook_15-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/login_hook_15-1.5-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-15-login-hook` | 1.6 | `d12.aarch64` | pigsty | 27.1 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-login-hook` | 1.6 | `d12.x86_64` | pigsty | 27.5 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-login-hook` | 1.6 | `u22.aarch64` | pigsty | 28.6 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-login-hook` | 1.6 | `u22.x86_64` | pigsty | 28.7 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-login-hook` | 1.6 | `u24.aarch64` | pigsty | 27.5 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-login-hook` | 1.6 | `u24.x86_64` | pigsty | 28.1 KiB | [postgresql-15-login-hook_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-15-login-hook_1.6-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `login_hook_14` | 1.7 | `el8.x86_64` | pgdg | 18.0 KiB | [login_hook_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | 1.7 | `el8.aarch64` | pgdg | 17.9 KiB | [login_hook_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | 1.6 | `el8.x86_64` | pgdg | 17.5 KiB | [login_hook_14-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | 1.6 | `el8.aarch64` | pgdg | 17.4 KiB | [login_hook_14-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | 1.5 | `el8.aarch64` | pgdg | 16.7 KiB | [login_hook_14-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/login_hook_14-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_14` | 1.5 | `el8.x86_64` | pgdg | 16.7 KiB | [login_hook_14-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/login_hook_14-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_14` | 1.7 | `el9.aarch64` | pgdg | 18.0 KiB | [login_hook_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | 1.7 | `el9.x86_64` | pgdg | 18.2 KiB | [login_hook_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_14` | 1.6 | `el9.x86_64` | pgdg | 17.7 KiB | [login_hook_14-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_14` | 1.6 | `el9.aarch64` | pgdg | 17.4 KiB | [login_hook_14-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | 1.5 | `el9.aarch64` | pgdg | 16.4 KiB | [login_hook_14-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/login_hook_14-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_14` | 1.5 | `el9.x86_64` | pgdg | 16.8 KiB | [login_hook_14-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/login_hook_14-1.5-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-login-hook` | 1.6 | `d12.x86_64` | pigsty | 27.8 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-login-hook` | 1.6 | `d12.aarch64` | pigsty | 27.3 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-login-hook` | 1.6 | `u22.x86_64` | pigsty | 29.0 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-login-hook` | 1.6 | `u22.aarch64` | pigsty | 28.8 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-login-hook` | 1.6 | `u24.aarch64` | pigsty | 27.7 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-login-hook` | 1.6 | `u24.x86_64` | pigsty | 28.3 KiB | [postgresql-14-login-hook_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-14-login-hook_1.6-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `login_hook_13` | 1.7 | `el8.x86_64` | pgdg | 17.9 KiB | [login_hook_13-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/login_hook_13-1.7-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_13` | 1.7 | `el8.aarch64` | pgdg | 17.9 KiB | [login_hook_13-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/login_hook_13-1.7-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_13` | 1.6 | `el8.x86_64` | pgdg | 17.4 KiB | [login_hook_13-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/login_hook_13-1.6-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_13` | 1.6 | `el8.aarch64` | pgdg | 17.4 KiB | [login_hook_13-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/login_hook_13-1.6-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_13` | 1.5 | `el8.aarch64` | pgdg | 16.6 KiB | [login_hook_13-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/login_hook_13-1.5-1PGDG.rhel8.aarch64.rpm) |
| `login_hook_13` | 1.5 | `el8.x86_64` | pgdg | 16.7 KiB | [login_hook_13-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/login_hook_13-1.5-1PGDG.rhel8.x86_64.rpm) |
| `login_hook_13` | 1.7 | `el9.aarch64` | pgdg | 18.0 KiB | [login_hook_13-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/login_hook_13-1.7-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_13` | 1.7 | `el9.x86_64` | pgdg | 18.1 KiB | [login_hook_13-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/login_hook_13-1.7-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_13` | 1.6 | `el9.x86_64` | pgdg | 17.6 KiB | [login_hook_13-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/login_hook_13-1.6-1PGDG.rhel9.x86_64.rpm) |
| `login_hook_13` | 1.6 | `el9.aarch64` | pgdg | 17.4 KiB | [login_hook_13-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/login_hook_13-1.6-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_13` | 1.5 | `el9.aarch64` | pgdg | 16.4 KiB | [login_hook_13-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/login_hook_13-1.5-1PGDG.rhel9.aarch64.rpm) |
| `login_hook_13` | 1.5 | `el9.x86_64` | pgdg | 16.7 KiB | [login_hook_13-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/login_hook_13-1.5-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-login-hook` | 1.6 | `d12.aarch64` | pigsty | 27.4 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-login-hook` | 1.6 | `d12.x86_64` | pigsty | 27.4 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-login-hook` | 1.6 | `u22.aarch64` | pigsty | 28.6 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-login-hook` | 1.6 | `u22.x86_64` | pigsty | 28.9 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-login-hook` | 1.6 | `u24.aarch64` | pigsty | 27.8 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-login-hook` | 1.6 | `u24.x86_64` | pigsty | 28.0 KiB | [postgresql-13-login-hook_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/login-hook/postgresql-13-login-hook_1.6-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/splendiddata/login_hook" title="Repository" icon="github" subtitle="github.com/splendiddata/login_hook" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="login_hook-1.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build get login_hook; # get login_hook source code
pig build dep login_hook; # install build dependencies
pig build pkg login_hook; # build extension rpm or deb
pig build ext login_hook; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install login_hook; # install by extension name, for the current active PG version
pig ext install login_hook; # install via package alias, for the active PG version
pig ext install login_hook -v 17;   # install for PG 17
pig ext install login_hook -v 16;   # install for PG 16
pig ext install login_hook -v 15;   # install for PG 15
pig ext install login_hook -v 14;   # install for PG 14
pig ext install login_hook -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION login_hook CASCADE SCHEMA login_hook;
```

