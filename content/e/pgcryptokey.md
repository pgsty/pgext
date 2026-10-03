---
title: "pgcryptokey"
linkTitle: "pgcryptokey"
description: "cryptographic key management"
weight: 7320
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**pgcryptokey**](https://momjian.us/download/pgcryptokey/) : cryptographic key management


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7320** | {{< badge content="pgcryptokey" link="https://momjian.us/download/pgcryptokey/" >}} | {{< ext "pgcryptokey" >}} | `0.85` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcrypto" >}} |
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "column_encrypt" >}} {{< ext "supabase_vault" >}} {{< ext "pg_enigma" >}} {{< ext "pg_tde" >}} {{< ext "pgcrypto" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pguecc" >}} {{< ext "pgsmcrypto" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgcryptokey` | `pgcrypto` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "pgcryptokey_18" "green" >}} {{< bg "17" "pgcryptokey_17" "green" >}} {{< bg "16" "pgcryptokey_16" "green" >}} {{< bg "15" "pgcryptokey_15" "green" >}} {{< bg "14" "pgcryptokey_14" "green" >}} | `pgcryptokey_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "postgresql-18-pgcryptokey" "green" >}} {{< bg "17" "postgresql-17-pgcryptokey" "green" >}} {{< bg "16" "postgresql-16-pgcryptokey" "green" >}} {{< bg "15" "postgresql-15-pgcryptokey" "green" >}} {{< bg "14" "postgresql-14-pgcryptokey" "green" >}} | `postgresql-$v-pgcryptokey` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_18` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_18` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_18` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_18-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_18-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_18-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.5 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_17` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.1 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.2 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_17-0.85-10PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgcryptokey_17-0.85-10PGDG.rhel9.8.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_17-0.85-10PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgcryptokey_17-0.85-10PGDG.rhel9.8.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pgcryptokey_17-0.85-10PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgcryptokey_17-0.85-10PGDG.rhel10.2.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_17-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pgcryptokey_17-0.85-10PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgcryptokey_17-0.85-10PGDG.rhel10.2.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_17-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_17-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_16` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_16-0.85-10PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgcryptokey_16-0.85-10PGDG.rhel9.8.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_16-0.85-10PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgcryptokey_16-0.85-10PGDG.rhel9.8.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pgcryptokey_16-0.85-10PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgcryptokey_16-0.85-10PGDG.rhel10.2.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_16-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pgcryptokey_16-0.85-10PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgcryptokey_16-0.85-10PGDG.rhel10.2.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_16-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_16-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_15` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgcryptokey_15-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgcryptokey_15-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_15-0.85-10PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgcryptokey_15-0.85-10PGDG.rhel9.8.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [pgcryptokey_15-0.85-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgcryptokey_15-0.85-3.rhel9.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_15-0.85-10PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgcryptokey_15-0.85-10PGDG.rhel9.8.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgcryptokey_15-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pgcryptokey_15-0.85-10PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgcryptokey_15-0.85-10PGDG.rhel10.2.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_15-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pgcryptokey_15-0.85-10PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgcryptokey_15-0.85-10PGDG.rhel10.2.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_15-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_15-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_14` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [pgcryptokey_14-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgcryptokey_14-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_14-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgcryptokey_14-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_14-0.85-10PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgcryptokey_14-0.85-10PGDG.rhel9.8.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_14-0.85-10PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgcryptokey_14-0.85-10PGDG.rhel9.8.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.3 KiB | [pgcryptokey_14-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgcryptokey_14-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pgcryptokey_14-0.85-10PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgcryptokey_14-0.85-10PGDG.rhel10.2.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_14-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pgcryptokey_14-0.85-10PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgcryptokey_14-0.85-10PGDG.rhel10.2.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_14-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_14-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.7 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://momjian.us/download/pgcryptokey/" title="Repository" icon="link" subtitle="momjian.us/download/pgcryptokey/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcryptokey-0.85.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgcryptokey;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgcryptokey;		# install via package name, for the active PG version

pig install pgcryptokey -v 18;   # install for PG 18
pig install pgcryptokey -v 17;   # install for PG 17
pig install pgcryptokey -v 16;   # install for PG 16
pig install pgcryptokey -v 15;   # install for PG 15
pig install pgcryptokey -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcryptokey CASCADE; -- requires pgcrypto
```

## Usage

Sources:

- [Official 0.85 source archive](https://momjian.us/download/pgcryptokey/pgcryptokey-0.85.tar.gz)
- [Official project directory](https://momjian.us/download/pgcryptokey/)

`pgcryptokey` manages encryption data keys wrapped by an access password. It stores wrapped keys in a database table and integrates with `pgcrypto` for encryption, key rotation and re-encryption.

### Install and Unlock

```sql
CREATE EXTENSION pgcryptokey CASCADE;
```

The dependency is `pgcrypto`. Source release 0.85 uses SQL extension version 1.0 and needs superuser installation. In client mode, first establish the documented session access password using `get_shared_key()` and `set_session_access_password(encrypted_password)`. The shared-key exchange supports SSL or Unix-domain socket connections only; the encrypted password argument is hex-encoded.

Boot mode instead preloads `pgcryptokey_acpass`, runs the protected server-side password acquisition script and requires a restart. It makes the access password server-wide and read-only. Choose one mode; boot and client modes cannot be combined while the server is running.

### Create and Use a Key

After unlocking access to the keys:

```sql
SELECT create_cryptokey('app-key', 32);
SELECT set_cryptokey('app-key');

CREATE TEMP TABLE secrets(ciphertext bytea);
INSERT INTO secrets VALUES
  (pgp_sym_encrypt('example', get_cryptokey('app-key')));
SELECT pgp_sym_decrypt(ciphertext, get_cryptokey('app-key'))
FROM secrets;
```

The key length is in bytes. Keys may be selected by name or integer key ID; name lookup refers to an active, non-superseded key.

### Rotate and Re-encrypt

`supersede_cryptokey(name, byte_len)` or its key-ID overload creates a replacement and returns its ID. The old and new keys initially use the same access password. Use the integer key-ID overload `change_key_access_password(key_id, new_encrypted_password)` to change that wrapping password. The session must already have its shared key and current access password set; the new password must be encrypted with the shared key and hex-encoded.

In source release 0.85, the name overload calls an undefined `change_access_password` function and cannot complete the password change. Use the integer overload above.

`reencrypt_data(data, old_key_id, new_key_id)` and `reencrypt_data_bytea(data, old_key_id, new_key_id)` migrate encrypted values. Preserve old key IDs alongside ciphertext and verify re-encryption before calling `drop_cryptokey(name)` or its key-ID overload; removing a key can make remaining ciphertext unreadable.

### Security Boundary

Protect the key table, function grants, access-password acquisition script and backups. Upstream warns that all users can view the boot-time `pgcryptokey.access_password`; table privileges are still required to use wrapped keys. `get_cryptokey(name)` returns raw key material, so do not expose its result through ordinary queries, logs or application traces. This design does not isolate keys from a trusted database administrator.
