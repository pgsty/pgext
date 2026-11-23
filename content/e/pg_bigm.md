---
title: "pg_bigm"
linkTitle: "pg_bigm"
description: "create 2-gram (bigram) index for faster full text search."
weight: 2120
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**pg_bigm**](https://github.com/pgbigm/pg_bigm) : create 2-gram (bigram) index for faster full text search.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2120** | {{< badge content="pg_bigm" link="https://github.com/pgbigm/pg_bigm" >}} | {{< ext "pg_bigm" >}} | `1.2` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pg_cjk_parser" >}} {{< ext "pg_jieba" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_pinyin" >}} {{< ext "unaccent" >}} {{< ext "pg_trgm" >}} {{< ext "dict_xsyn" >}} {{< ext "biscuit" >}} {{< ext "vgram" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.2` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_bigm` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "19" "pg_bigm_19" "green" >}} {{< bg "18" "pg_bigm_18" "green" >}} {{< bg "17" "pg_bigm_17" "green" >}} {{< bg "16" "pg_bigm_16" "green" >}} {{< bg "15" "pg_bigm_15" "green" >}} | `pg_bigm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "19" "postgresql-19-pg-bigm" "green" >}} {{< bg "18" "postgresql-18-pg-bigm" "green" >}} {{< bg "17" "postgresql-17-pg-bigm" "green" >}} {{< bg "16" "postgresql-16-pg-bigm" "green" >}} {{< bg "15" "postgresql-15-pg-bigm" "green" >}} | `postgresql-$v-pg-bigm` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_bigm_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_bigm_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 8" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-19-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_19` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pg_bigm_19-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_bigm_19-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_19` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_bigm_19-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_bigm_19-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_19` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.3 KiB | [pg_bigm_19-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_bigm_19-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_19` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.1 KiB | [pg_bigm_19-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_bigm_19-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.6 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.2 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.8 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.5 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.0 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 29.6 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 29.8 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 29.7 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 29.6 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb) |
| `postgresql-19-pg-bigm` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 29.8 KiB | [postgresql-19-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-19-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.0 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.9 KiB | [pg_bigm_18-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_18-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_bigm_18-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_18-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.7 KiB | [pg_bigm_18-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_18-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.5 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.2 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.7 KiB | [pg_bigm_18-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_18-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bigm_18-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.9 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bigm_18-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.0 KiB | [pg_bigm_18-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_18-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.1 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.1 KiB | [pg_bigm_18-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.2 KiB | [pg_bigm_18-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.0 KiB | [pg_bigm_18-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_18-1.2-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.1 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.6 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.2 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.9 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.5 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.4 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.3 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.2 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.3 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 30.0 KiB | [postgresql-18-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.8 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.5 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_bigm_17-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_17-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.5 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.1 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pg_bigm_17-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_17-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.7 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.4 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_bigm_17-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_17-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.1 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [pg_bigm_17-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_17-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.2 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_17-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_17-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.9 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.0 KiB | [pg_bigm_17-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.4 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.5 KiB | [pg_bigm_17-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.0 KiB | [pg_bigm_17-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.9 KiB | [pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_17-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_17-1.2-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.7 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.4 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.8 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.7 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.4 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.9 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.8 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.7 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.4 KiB | [postgresql-17-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.8 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.4 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_bigm_16-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_16-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.4 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.1 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.2 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pg_bigm_16-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_16-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.0 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.7 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.4 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_bigm_16-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_16-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [pg_bigm_16-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_16-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.2 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_16-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_16-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.9 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.0 KiB | [pg_bigm_16-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.5 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.5 KiB | [pg_bigm_16-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.0 KiB | [pg_bigm_16-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.2 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.0 KiB | [pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_16-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_16-1.2-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.7 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.4 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.8 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.7 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.3 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.9 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.8 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.7 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.4 KiB | [postgresql-16-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20261118-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.5 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_15-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.4 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.6 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20261118-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_15-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20261118-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20250903-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.6 KiB | [pg_bigm_15-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_15-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.5 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20261118-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.2 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20250903-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.6 KiB | [pg_bigm_15-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_15-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20261118-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.0 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20250903-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.0 KiB | [pg_bigm_15-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_15-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.2 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20261118-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.2 KiB | [pg_bigm_15-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20261118-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_bigm_15-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.3 KiB | [pg_bigm_15-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.4 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.1 KiB | [pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.1 KiB | [pg_bigm_15-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_15-1.2-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.9 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.5 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.8 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.4 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 29.1 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.9 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.9 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.6 KiB | [postgresql-15-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20261118PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgbigm/pg_bigm" title="Repository" icon="github" subtitle="github.com/pgbigm/pg_bigm" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bigm-1.2-20261118.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_bigm;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_bigm;		# install via package name, for the active PG version

pig install pg_bigm -v 19;   # install for PG 19
pig install pg_bigm -v 18;   # install for PG 18
pig install pg_bigm -v 17;   # install for PG 17
pig install pg_bigm -v 16;   # install for PG 16
pig install pg_bigm -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_bigm';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_bigm;
```

## Usage

Sources:

- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/docs/pg_bigm_en.md)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/pg_bigm.control)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/README.md)
- [Official source](https://github.com/pgbigm/pg_bigm/releases/tag/v1.2-20261118)

`pg_bigm` builds GIN bigram indexes for LIKE and similarity searches, including multibyte text. Release 1.2-20261118 keeps SQL version 1.2 and adds PostgreSQL 19 support.

### Search

```sql
CREATE EXTENSION pg_bigm;
CREATE TABLE documents (id bigint PRIMARY KEY, content text);
INSERT INTO documents VALUES (1, 'PostgreSQL database search');
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);
SELECT * FROM documents WHERE content LIKE likequery('database');
SELECT show_bigm('PostgreSQL');
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

`likequery(text)` escapes a literal search keyword for LIKE. `show_bigm(text)` displays its bigrams, and `bigm_similarity(text,text)` returns a similarity value. The `=%` operator applies `pg_bigm.similarity_limit`, normally 0.3. An index is useful only when the planner selects it.

### Settings and maintenance

`pg_bigm.enable_recheck` normally remains on: disabling rechecks can return false positives. `pg_bigm.gin_key_limit` defaults to 0 (no limit); a positive value bounds keys used in a search. `pg_gin_pending_stats(regclass)` reports GIN pending-list pages and tuples. `pg_bigm.last_update` identifies the library revision.

The 20261118 build fixes incorrect multibyte truncation, LIKE underscore handling and missing similarity matches. The truncation fix requires compilation with headers from PostgreSQL 14.21, 15.16, 16.12, 17.8, 18.2 or later, or PostgreSQL 19; the same extension SQL version does not prove the library contains the correction. Deploy the matching library to all servers and replace processes using the old library. Base use requires no preload.
