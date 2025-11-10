---
title: "pg_track_settings"
linkTitle: "pg_track_settings"
description: "Track settings changes"
weight: 6260
categories: ["STAT"]
width: full
---

[**pg_track_settings**](https://github.com/rjuju/pg_track_settings)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6260** | {{< badge content="pg_track_settings" link="https://github.com/rjuju/pg_track_settings" >}} | {{< ext "pg_track_settings" >}} | `2.1.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_profile" >}} {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_qualstats" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_track_settings" >}} | `2.1.2` | {{< bg "18" "pg_track_settings_18" "green" >}} {{< bg "17" "pg_track_settings_17" "green" >}} {{< bg "16" "pg_track_settings_16" "green" >}} {{< bg "15" "pg_track_settings_15" "green" >}} {{< bg "14" "pg_track_settings_14" "green" >}} {{< bg "13" "pg_track_settings_13" "green" >}} | `pg_track_settings_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_track_settings" >}} | `2.1.2` | {{< bg "18" "postgresql-18-pg-track-settings" "green" >}} {{< bg "17" "postgresql-17-pg-track-settings" "green" >}} {{< bg "16" "postgresql-16-pg-track-settings" "green" >}} {{< bg "15" "postgresql-15-pg-track-settings" "green" >}} {{< bg "14" "postgresql-14-pg-track-settings" "green" >}} {{< bg "13" "postgresql-13-pg-track-settings" "green" >}} | `postgresql-$v-pg-track-settings` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "pg_track_settings_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.1.2" "postgresql-18-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-17-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-16-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-15-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-14-pg-track-settings : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.2" "postgresql-13-pg-track-settings : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_18` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.5 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_track_settings_18-2.1.2-3PGDG.rhel8.noarch.rpm) |
| `pg_track_settings_18` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.5 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_track_settings_18-2.1.2-3PGDG.rhel8.noarch.rpm) |
| `pg_track_settings_18` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.9 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_track_settings_18-2.1.2-3PGDG.rhel9.noarch.rpm) |
| `pg_track_settings_18` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.8 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_track_settings_18-2.1.2-3PGDG.rhel9.noarch.rpm) |
| `pg_track_settings_18` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_track_settings_18-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_18` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_18-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_track_settings_18-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.7 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.7 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-18-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-18-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_17` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.6 KiB | [pg_track_settings_17-2.1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_track_settings_17-2.1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_track_settings_17` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.5 KiB | [pg_track_settings_17-2.1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_track_settings_17-2.1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_track_settings_17` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.1 KiB | [pg_track_settings_17-2.1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_track_settings_17-2.1.2-2PGDG.rhel9.x86_64.rpm) |
| `pg_track_settings_17` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.0 KiB | [pg_track_settings_17-2.1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_track_settings_17-2.1.2-2PGDG.rhel9.aarch64.rpm) |
| `pg_track_settings_17` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_17-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_track_settings_17-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_17` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_17-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_track_settings_17-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.7 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.7 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-17-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-17-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_16` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.6 KiB | [pg_track_settings_16-2.1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_track_settings_16-2.1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_track_settings_16` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.5 KiB | [pg_track_settings_16-2.1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_track_settings_16-2.1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_track_settings_16` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.1 KiB | [pg_track_settings_16-2.1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_track_settings_16-2.1.2-2PGDG.rhel9.x86_64.rpm) |
| `pg_track_settings_16` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.9 KiB | [pg_track_settings_16-2.1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_track_settings_16-2.1.2-2PGDG.rhel9.aarch64.rpm) |
| `pg_track_settings_16` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_16-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_track_settings_16-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_16` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_16-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_track_settings_16-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.8 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.8 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-16-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-16-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_15` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [pg_track_settings_15-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_track_settings_15-2.1.2-1.rhel8.x86_64.rpm) |
| `pg_track_settings_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_track_settings_15-2.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_track_settings_15-2.1.0-1.rhel8.x86_64.rpm) |
| `pg_track_settings_15` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [pg_track_settings_15-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_track_settings_15-2.1.2-1.rhel8.aarch64.rpm) |
| `pg_track_settings_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_track_settings_15-2.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_track_settings_15-2.1.0-1.rhel8.aarch64.rpm) |
| `pg_track_settings_15` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.0 KiB | [pg_track_settings_15-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_track_settings_15-2.1.2-1.rhel9.x86_64.rpm) |
| `pg_track_settings_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.6 KiB | [pg_track_settings_15-2.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_track_settings_15-2.1.0-1.rhel9.x86_64.rpm) |
| `pg_track_settings_15` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.9 KiB | [pg_track_settings_15-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_track_settings_15-2.1.2-1.rhel9.aarch64.rpm) |
| `pg_track_settings_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.5 KiB | [pg_track_settings_15-2.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_track_settings_15-2.1.0-1.rhel9.aarch64.rpm) |
| `pg_track_settings_15` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_15-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_track_settings_15-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_15` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_15-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_track_settings_15-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.7 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.7 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-15-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-15-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_14` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [pg_track_settings_14-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_track_settings_14-2.1.2-1.rhel8.x86_64.rpm) |
| `pg_track_settings_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_track_settings_14-2.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_track_settings_14-2.1.0-1.rhel8.x86_64.rpm) |
| `pg_track_settings_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.6 KiB | [pg_track_settings_14-2.0.1-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_track_settings_14-2.0.1-3.rhel8.x86_64.rpm) |
| `pg_track_settings_14` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [pg_track_settings_14-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_track_settings_14-2.1.2-1.rhel8.aarch64.rpm) |
| `pg_track_settings_14` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_track_settings_14-2.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_track_settings_14-2.1.0-1.rhel8.aarch64.rpm) |
| `pg_track_settings_14` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.0 KiB | [pg_track_settings_14-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_track_settings_14-2.1.2-1.rhel9.x86_64.rpm) |
| `pg_track_settings_14` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.6 KiB | [pg_track_settings_14-2.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_track_settings_14-2.1.0-1.rhel9.x86_64.rpm) |
| `pg_track_settings_14` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.9 KiB | [pg_track_settings_14-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_track_settings_14-2.1.2-1.rhel9.aarch64.rpm) |
| `pg_track_settings_14` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.5 KiB | [pg_track_settings_14-2.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_track_settings_14-2.1.0-1.rhel9.aarch64.rpm) |
| `pg_track_settings_14` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_14-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_track_settings_14-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_14` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_14-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_track_settings_14-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.7 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.7 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-14-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-14-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-14-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_settings_13` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [pg_track_settings_13-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_track_settings_13-2.1.2-1.rhel8.x86_64.rpm) |
| `pg_track_settings_13` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_track_settings_13-2.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_track_settings_13-2.1.0-1.rhel8.x86_64.rpm) |
| `pg_track_settings_13` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.4 KiB | [pg_track_settings_13-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_track_settings_13-2.1.2-1.rhel8.aarch64.rpm) |
| `pg_track_settings_13` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_track_settings_13-2.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_track_settings_13-2.1.0-1.rhel8.aarch64.rpm) |
| `pg_track_settings_13` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.1 KiB | [pg_track_settings_13-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_track_settings_13-2.1.2-1.rhel9.x86_64.rpm) |
| `pg_track_settings_13` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.6 KiB | [pg_track_settings_13-2.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_track_settings_13-2.1.0-1.rhel9.x86_64.rpm) |
| `pg_track_settings_13` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.9 KiB | [pg_track_settings_13-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_track_settings_13-2.1.2-1.rhel9.aarch64.rpm) |
| `pg_track_settings_13` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.5 KiB | [pg_track_settings_13-2.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_track_settings_13-2.1.0-1.rhel9.aarch64.rpm) |
| `pg_track_settings_13` | `2.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.4 KiB | [pg_track_settings_13-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_track_settings_13-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `pg_track_settings_13` | `2.1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [pg_track_settings_13-2.1.2-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_track_settings_13-2.1.2-3PGDG.rhel10.noarch.rpm) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 9.7 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 9.7 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg12+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 9.8 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 9.8 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg13+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 9.1 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 9.1 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg22.04+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 9.1 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |
| `postgresql-13-pg-track-settings` | `2.1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 9.1 KiB | [postgresql-13-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-track-settings/postgresql-13-pg-track-settings_2.1.2-5.pgdg24.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rjuju/pg_track_settings" title="Repository" icon="github" subtitle="github.com/rjuju/pg_track_settings" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_track_settings; # install by extension name, for the current active PG version
pig ext install pg_track_settings; # install via package alias, for the active PG version
pig ext install pg_track_settings -v 18;   # install for PG 18
pig ext install pg_track_settings -v 17;   # install for PG 17
pig ext install pg_track_settings -v 16;   # install for PG 16
pig ext install pg_track_settings -v 15;   # install for PG 15
pig ext install pg_track_settings -v 14;   # install for PG 14
pig ext install pg_track_settings -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_track_settings;
```

