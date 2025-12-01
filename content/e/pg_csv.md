---
title: "pg_csv"
linkTitle: "pg_csv"
description: "Flexible CSV processing for Postgres"
weight: 4760
categories: ["FUNC"]
width: full
---

[**pg_csv**](https://github.com/PostgREST/pg_csv) : Flexible CSV processing for Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4760** | {{< badge content="pg_csv" link="https://github.com/PostgREST/pg_csv" >}} | {{< ext "pg_csv" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_vecs" >}} {{< ext "first_last_agg" >}} {{< ext "arraymath" >}} {{< ext "intarray" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_csv` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.1` | {{< bg "18" "pg_csv_18*" "green" >}} {{< bg "17" "pg_csv_17*" "green" >}} {{< bg "16" "pg_csv_16*" "green" >}} {{< bg "15" "pg_csv_15*" "green" >}} {{< bg "14" "pg_csv_14*" "green" >}} {{< bg "13" "pg_csv_13*" "red" >}} | `pg_csv_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-pg-csv" "green" >}} {{< bg "17" "postgresql-17-pg-csv" "green" >}} {{< bg "16" "postgresql-16-pg-csv" "green" >}} {{< bg "15" "postgresql-15-pg-csv" "green" >}} {{< bg "14" "postgresql-14-pg-csv" "green" >}} {{< bg "13" "postgresql-13-pg-csv" "red" >}} | `postgresql-$v-pg-csv` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.1" "pg_csv_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.1" "pg_csv_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_csv_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-csv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-csv : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-csv : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_csv_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.3 KiB | [pg_csv_18-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_csv_18-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_csv_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.2 KiB | [pg_csv_18-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_csv_18-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_csv_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_csv_18-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_csv_18-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_csv_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.9 KiB | [pg_csv_18-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_csv_18-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_csv_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.6 KiB | [pg_csv_18-1.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_csv_18-1.0.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_csv_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.6 KiB | [pg_csv_18-1.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_csv_18-1.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-csv` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.0 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.3 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.7 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-csv` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.2 KiB | [postgresql-18-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-18-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_csv_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.3 KiB | [pg_csv_17-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_csv_17-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_csv_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.2 KiB | [pg_csv_17-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_csv_17-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_csv_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_csv_17-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_csv_17-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_csv_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.9 KiB | [pg_csv_17-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_csv_17-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_csv_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.6 KiB | [pg_csv_17-1.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_csv_17-1.0.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_csv_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.6 KiB | [pg_csv_17-1.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_csv_17-1.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-csv` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.3 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.3 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.0 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.6 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.2 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-csv` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.2 KiB | [postgresql-17-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-17-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_csv_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.3 KiB | [pg_csv_16-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_csv_16-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_csv_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.2 KiB | [pg_csv_16-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_csv_16-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_csv_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_csv_16-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_csv_16-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_csv_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.9 KiB | [pg_csv_16-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_csv_16-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_csv_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.5 KiB | [pg_csv_16-1.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_csv_16-1.0.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_csv_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.6 KiB | [pg_csv_16-1.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_csv_16-1.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-csv` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.3 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.2 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.0 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.6 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.2 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-csv` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.2 KiB | [postgresql-16-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-16-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_csv_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.5 KiB | [pg_csv_15-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_csv_15-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_csv_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.4 KiB | [pg_csv_15-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_csv_15-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_csv_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_csv_15-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_csv_15-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_csv_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.3 KiB | [pg_csv_15-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_csv_15-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_csv_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.9 KiB | [pg_csv_15-1.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_csv_15-1.0.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_csv_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.0 KiB | [pg_csv_15-1.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_csv_15-1.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-csv` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.0 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.5 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.9 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-csv` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-15-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-15-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_csv_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.5 KiB | [pg_csv_14-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_csv_14-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_csv_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.4 KiB | [pg_csv_14-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_csv_14-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_csv_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.5 KiB | [pg_csv_14-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_csv_14-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_csv_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.3 KiB | [pg_csv_14-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_csv_14-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_csv_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.9 KiB | [pg_csv_14-1.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_csv_14-1.0.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_csv_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.9 KiB | [pg_csv_14-1.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_csv_14-1.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-csv` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.3 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.0 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.3 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.0 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.8 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.5 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.8 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-csv` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.4 KiB | [postgresql-14-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-csv/postgresql-14-pg-csv_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/PostgREST/pg_csv" title="Repository" icon="github" subtitle="github.com/PostgREST/pg_csv" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_csv-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_csv;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_csv;		# install via package name, for the active PG version

pig install pg_csv -v 18;   # install for PG 18
pig install pg_csv -v 17;   # install for PG 17
pig install pg_csv -v 16;   # install for PG 16
pig install pg_csv -v 15;   # install for PG 15
pig install pg_csv -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_csv;
```
