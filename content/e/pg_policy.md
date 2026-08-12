---
title: "pg_policy"
linkTitle: "pg_policy"
description: "Agentic policy language for PostgreSQL with guardrails, guidance, and session-aware controls"
weight: 7440
categories: ["SEC"]
width: full
---

[**pg_policy**](https://github.com/rahiakil/pg-policy) : Agentic policy language for PostgreSQL with guardrails, guidance, and session-aware controls


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7440** | {{< badge content="pg_policy" link="https://github.com/rahiakil/pg-policy" >}} | {{< ext "pg_policy" >}} | `0.1.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `policy` |
|   **See Also**    | {{< ext "pg_command_fw" >}} {{< ext "pgextwlist" >}} {{< ext "set_user" >}} {{< ext "noset" >}} {{< ext "block_copy_command" >}} {{< ext "supautils" >}} {{< ext "anon" >}} {{< ext "pgaudit" >}} |

> [!Note] PIGSTY patches the reserved upstream schema pg_policy to policy and quotes the reserved check function, so the packaged API is policy.check() rather than pg_policy.check(); pure SQL and PL/pgSQL, no preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_policy` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_policy_18" "green" >}} {{< bg "17" "pg_policy_17" "green" >}} {{< bg "16" "pg_policy_16" "green" >}} {{< bg "15" "pg_policy_15" "green" >}} {{< bg "14" "pg_policy_14" "green" >}} | `pg_policy_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-policy" "green" >}} {{< bg "17" "postgresql-17-pg-policy" "green" >}} {{< bg "16" "postgresql-16-pg-policy" "green" >}} {{< bg "15" "postgresql-15-pg-policy" "green" >}} {{< bg "14" "postgresql-14-pg-policy" "green" >}} | `postgresql-$v-pg-policy` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_policy_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-policy : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-policy : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_policy_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.9 KiB | [pg_policy_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_policy_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_policy_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_policy_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 KiB | [pg_policy_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_policy_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_policy_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_policy_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_policy_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_policy_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_policy_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [pg_policy_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_policy_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-policy` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-policy` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.3 KiB | [postgresql-18-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-18-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_policy_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.9 KiB | [pg_policy_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_policy_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_policy_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_policy_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 KiB | [pg_policy_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_policy_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_policy_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_policy_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_policy_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_policy_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_policy_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [pg_policy_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_policy_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-policy` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-policy` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.3 KiB | [postgresql-17-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-17-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_policy_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.9 KiB | [pg_policy_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_policy_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_policy_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_policy_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 KiB | [pg_policy_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_policy_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_policy_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_policy_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_policy_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_policy_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_policy_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [pg_policy_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_policy_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-policy` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-policy` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.3 KiB | [postgresql-16-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-16-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_policy_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.9 KiB | [pg_policy_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_policy_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_policy_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_policy_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 KiB | [pg_policy_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_policy_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_policy_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_policy_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_policy_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_policy_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_policy_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [pg_policy_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_policy_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-policy` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-policy` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.3 KiB | [postgresql-15-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-15-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_policy_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.9 KiB | [pg_policy_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_policy_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_policy_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_policy_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_policy_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 KiB | [pg_policy_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_policy_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_policy_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_policy_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_policy_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_policy_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_policy_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_policy_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [pg_policy_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_policy_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-policy` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-policy` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.3 KiB | [postgresql-14-pg-policy_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-policy/postgresql-14-pg-policy_0.1.0-1PGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rahiakil/pg-policy" title="Repository" icon="github" subtitle="github.com/rahiakil/pg-policy" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_policy-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_policy;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_policy;		# install via package name, for the active PG version

pig install pg_policy -v 18;   # install for PG 18
pig install pg_policy -v 17;   # install for PG 17
pig install pg_policy -v 16;   # install for PG 16
pig install pg_policy -v 15;   # install for PG 15
pig install pg_policy -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_policy;
```

## Usage

Sources:

- [pg_policy 0.1.0 on PGXN](https://pgxn.org/dist/pg_policy/0.1.0/)
- [pg_policy 0.1.0 README](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/README.md)
- [Agent Policy Language reference](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/doc/language.md)
- [pg_policy 0.1.0 security policy](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/SECURITY.md)
- [pg_policy 0.1.0 control file](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/pg_policy.control)
- [pg_policy 0.1.0 extension SQL](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/sql/pg_policy--0.1.0.sql)
- [Pigsty pg_policy package page](https://pgext.cloud/ext/pg_policy)

`pg_policy` 0.1.0 is an experimental SQL and PL/pgSQL policy evaluator for agent and tool actions. It stores Agent Policy Language rules, evaluates context and session history, records every decision, and returns obligations for a gateway to enforce. It complements PostgreSQL roles and row-level security; it does not intercept SQL or tool calls by itself.

### Pigsty Schema Compatibility

Upstream 0.1.0 declares the reserved schema name `pg_policy` and defines an unquoted function named `check`. Pigsty packages patch the installed schema to `policy`, quote the reserved function name as `policy."check"()`, and fix function search paths. The upstream examples therefore cannot be copied verbatim into a Pigsty installation.

```sql
CREATE EXTENSION pg_policy;

SELECT policy.set_setting('enforcement_mode', 'log_only');
```

The extension is not relocatable, requires PostgreSQL 14 or later, and does not require `shared_preload_libraries` or a PostgreSQL restart. Current Pigsty packages cover PostgreSQL 14–18.

### Define and Evaluate a Guardrail

```sql
SELECT policy.upsert_policy('block_ddl', $apl$
forbid
  principal agent "research_bot"
  action tool "execute_sql"
  when { context.statement_type in ["DROP", "TRUNCATE", "ALTER", "CREATE"] }
  reason "Research agents may not run DDL"
$apl$);

SELECT policy.set_setting('enforcement_mode', 'enforce');

SELECT policy.evaluate(
  'agent', 'research_bot',
  'tool', 'execute_sql',
  '*', '*',
  '{"statement_type":"DROP"}'::jsonb,
  NULL
);

SELECT policy."check"(
  'research_bot',
  'execute_sql',
  '{"statement_type":"DROP"}'::jsonb
);
```

`policy.evaluate(...)` returns JSON containing `decision`, `allowed`, `matched_policies`, `obligations`, `reasons`, and `mode`. The convenience wrapper `policy."check"()` returns only a boolean. `policy.enforce()` requests exception-on-deny behavior when the mode is `enforce`.

### APL Surface

An APL document begins with one effect: `permit`, `forbid`, or `guide`. It can match principal, action, and resource types and identifiers. In 0.1.0, context conditions support only `==`, `in [...]`, and `and`. A temporal clause can count matching session events inside an interval when evaluation receives a session identifier.

`forbid` overrides matching `permit` rules. `guide` allows the action and can return `advice`, `prefer_tool`, or `max_rows` obligations. The caller—not the extension—must interpret and apply those obligations.

### Sessions, Temporal Limits, and Audit

```sql
SELECT policy.open_session(
  'sess-1',
  'agent',
  'research_bot'
);

SELECT policy.upsert_policy('export_budget', $apl$
forbid
  principal agent "research_bot"
  action tool "export_csv"
  when temporal {
    count(action == "export_csv") within interval '1 hour' >= 3
  }
  reason "Export budget exceeded"
$apl$);

SELECT policy.evaluate(
  'agent', 'research_bot',
  'tool', 'export_csv',
  '*', '*',
  '{}'::jsonb,
  'sess-1'
);
```

`policy.open_session()` creates or updates a session. Evaluations with a session identifier append an event and can satisfy temporal predicates. Every evaluation writes `policy.decision_log`; other important relations are `policy.policies`, `policy.sessions`, `policy.events`, and `policy.settings`.

### Enforcement and Security Boundaries

- The default `enforcement_mode` is `log_only` and the default decision is `permit`. A matched deny becomes an allow with a `shadow_deny` obligation.
- In `guide` mode, a matched deny becomes an allow with `would_deny`. Only `enforce` preserves a deny and allows `policy.enforce()` to raise an error.
- A gateway must call the evaluator before the protected action and hard-fail on deny. Calling `policy.evaluate(...)` after executing a tool is only auditing.
- Keep PostgreSQL `GRANT` and `REVOKE`, row-level security, network controls, and least-privilege credentials as the authoritative data-plane controls. Superusers and roles with `BYPASSRLS` can bypass row-level controls.
- The 0.1 line is explicitly an experimental MVP, not a hardened production security boundary. Shadow-test policies, restrict who can change `policy.settings` or `policy.policies`, and monitor `policy.decision_log` before switching to `enforce`.
