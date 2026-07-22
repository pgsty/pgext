---
title: "pgfr_record"
linkTitle: "pgfr_record"
description: "Server-side PostgreSQL performance flight recorder"
weight: 6060
categories: ["STAT"]
width: full
---

[**pg_flight_recorder**](https://github.com/dventimisupabase/pg_flight_recorder) : Server-side PostgreSQL performance flight recorder


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6060** | {{< badge content="pgfr_record" link="https://github.com/dventimisupabase/pg_flight_recorder" >}} | {{< ext "pgfr_record" "pg_flight_recorder" >}} | `2.29.2` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgfr_record` |
|   **Requires**    | {{< ext "pg_cron" >}} |
|    **Need By**    | {{< ext "pgfr_analyze" >}} |
|   **See Also**    | {{< ext "pg_stat_statements" >}} {{< ext "pg_cron" >}} {{< ext "pg_profile" >}} {{< ext "pgmonitor" >}} |
|    **Siblings**   | {{< ext "pgfr_analyze" >}} |

> [!Note] Package normalizes the upstream 0.0.0 control version to 2.29.2; run SELECT pgfr_record.enable() after CREATE EXTENSION. The downstream install patch defers scheduling until the CREATE transaction commits and guards optional pg_stat_statements.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_flight_recorder` | `pg_cron` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "pg_flight_recorder_18" "green" >}} {{< bg "17" "pg_flight_recorder_17" "green" >}} {{< bg "16" "pg_flight_recorder_16" "green" >}} {{< bg "15" "pg_flight_recorder_15" "green" >}} {{< bg "14" "pg_flight_recorder_14" "red" >}} | `pg_flight_recorder_$v` | `pg_cron_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "postgresql-18-pg-flight-recorder" "green" >}} {{< bg "17" "postgresql-17-pg-flight-recorder" "green" >}} {{< bg "16" "postgresql-16-pg-flight-recorder" "green" >}} {{< bg "15" "postgresql-15-pg-flight-recorder" "green" >}} {{< bg "14" "postgresql-14-pg-flight-recorder" "red" >}} | `postgresql-$v-pg-flight-recorder` | `postgresql-$v-cron` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_flight_recorder_18` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.0 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_flight_recorder_18-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_18` | `2.29.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.0 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_flight_recorder_18-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_18` | `2.29.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 86.4 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_flight_recorder_18-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_18` | `2.29.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 86.3 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_flight_recorder_18-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_18` | `2.29.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 86.6 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_flight_recorder_18-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `pg_flight_recorder_18` | `2.29.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 86.5 KiB | [pg_flight_recorder_18-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_flight_recorder_18-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 79.2 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.2 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 79.2 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 79.2 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.0 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.0 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.7 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.7 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.7 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-flight-recorder` | `2.29.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.7 KiB | [postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-18-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_flight_recorder_17` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.0 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_flight_recorder_17-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_17` | `2.29.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.0 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_flight_recorder_17-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_17` | `2.29.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 86.4 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_flight_recorder_17-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_17` | `2.29.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 86.3 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_flight_recorder_17-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_17` | `2.29.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 86.5 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_flight_recorder_17-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `pg_flight_recorder_17` | `2.29.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 86.5 KiB | [pg_flight_recorder_17-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_flight_recorder_17-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 79.2 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.2 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 79.2 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 79.2 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.0 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.0 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.7 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.7 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.7 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-flight-recorder` | `2.29.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.7 KiB | [postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-17-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_flight_recorder_16` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.0 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_flight_recorder_16-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_16` | `2.29.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.0 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_flight_recorder_16-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_16` | `2.29.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 86.4 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_flight_recorder_16-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_16` | `2.29.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 86.3 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_flight_recorder_16-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_16` | `2.29.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 86.5 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_flight_recorder_16-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `pg_flight_recorder_16` | `2.29.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 86.5 KiB | [pg_flight_recorder_16-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_flight_recorder_16-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 79.2 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.2 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 79.2 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 79.2 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.0 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.0 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.7 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.7 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.7 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-flight-recorder` | `2.29.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.7 KiB | [postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-16-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_flight_recorder_15` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.0 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_flight_recorder_15-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_15` | `2.29.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.0 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_flight_recorder_15-2.29.2-1PIGSTY.el8.noarch.rpm) |
| `pg_flight_recorder_15` | `2.29.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 86.3 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_flight_recorder_15-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_15` | `2.29.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 86.3 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_flight_recorder_15-2.29.2-1PIGSTY.el9.noarch.rpm) |
| `pg_flight_recorder_15` | `2.29.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 86.5 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_flight_recorder_15-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `pg_flight_recorder_15` | `2.29.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 86.5 KiB | [pg_flight_recorder_15-2.29.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_flight_recorder_15-2.29.2-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 79.2 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.2 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 79.2 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 79.2 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.0 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.0 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.7 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.7 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.7 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-flight-recorder` | `2.29.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.7 KiB | [postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-flight-recorder/postgresql-15-pg-flight-recorder_2.29.2-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dventimisupabase/pg_flight_recorder" title="Repository" icon="github" subtitle="github.com/dventimisupabase/pg_flight_recorder" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_flight_recorder-2.29.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_flight_recorder;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_flight_recorder;		# install via package name, for the active PG version
pig install pgfr_record;		# install by extension name, for the current active PG version

pig install pgfr_record -v 18;   # install for PG 18
pig install pgfr_record -v 17;   # install for PG 17
pig install pgfr_record -v 16;   # install for PG 16
pig install pgfr_record -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgfr_record CASCADE; -- requires pg_cron
```

## Usage

Sources:

- [pgfr_record v2.29.2 README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/README.md)
- [pgfr_record control file](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/extension.control)
- [pg_flight_recorder v2.29.2 reference](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 release notes](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_record is the collection half of PostgreSQL Flight Recorder. It periodically samples PostgreSQL activity, waits, locks, replication, vacuum, and related health data into bounded in-database buffers, then retains snapshots for incident analysis. Use it when short-lived database conditions must survive long enough for later diagnosis.

### Install and Enable Recording

pgfr_record requires pg_cron:

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    SELECT pgfr_record.enable();

enable() installs and schedules the collector jobs. It also reports configuration warnings; review them rather than treating a successful call as proof that every metric is being collected.

### Inspect Recorder Health

    SELECT * FROM pgfr_record.health_check();
    SELECT * FROM pgfr_record.ring_buffer_health();
    SELECT * FROM pgfr_record.list_profiles();

Use set_mode or apply_profile to select the intended collection profile before an incident:

    SELECT pgfr_record.set_mode('normal');

The available collection modes are normal, light, emergency, and kill. Profile names and sampling intervals can evolve, so list the installed profiles rather than hard-coding an undocumented name.

### Recorded Data Index

- deltas: interval changes for cumulative PostgreSQL counters.
- recent_activity and recent_waits: sampled sessions and wait events.
- recent_locks and recent_idle: lock and idle-session observations.
- recent_replication and recent_vacuum: replication and maintenance state.
- archiver_status: WAL archive health.
- snapshot and ring-buffer tables: retained history used by pgfr_analyze.

Many working buffers are UNLOGGED to reduce write amplification. They are not crash-durable and are not replicated like ordinary logged tables; durable snapshots provide the longer-lived analysis surface.

### Administration Functions

- pgfr_record.enable(): create or activate scheduled collectors.
- pgfr_record.disable(): stop scheduled collection.
- pgfr_record.health_check(): report collector and configuration health.
- pgfr_record.set_mode(...): change collection mode.
- pgfr_record.apply_profile(...): apply a predefined profile.
- pgfr_record.list_profiles(): enumerate available profiles.
- pgfr_record.ring_buffer_health(): inspect capacity and retention pressure.
- pgfr_record.cleanup(...): remove retained history according to the API.

### Retention and Overhead

The default design keeps short ring-buffer history and longer durable snapshots, commonly around 7 and 30 days depending on the installed profile. Verify actual table sizes, job schedules, and retention settings in the installed version.

The recorder creates roughly ten pg_cron jobs. pg_cron.log_run can generate thousands of rows per day; disable that logging or purge cron history when the extra audit trail is unnecessary. Sampling also adds SQL, storage, and catalog traffic, so measure overhead on the target workload.

Version 2.29.2 handles managed-service roles that cannot UPDATE cron.job: jobs can still be scheduled, while the optional nodename normalization is skipped with a warning.

### Caveats

- pg_stat_statements enriches several analyses but is optional; enable and size it separately when needed.
- Collection cannot reconstruct time periods that were never sampled. Enable and validate the recorder before an incident.
- UNLOGGED buffers can be truncated after crash recovery.
- Recorder tables can contain query text, role names, client data, and operational details. Apply appropriate privileges and retention controls.
