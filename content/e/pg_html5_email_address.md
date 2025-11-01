---
title: "pg_html5_email_address"
linkTitle: "pg_html5_email_address"
description: "PostgreSQL email validation that is consistent with the HTML5 spec"
weight: 4180
categories: ["UTIL"]
width: full
---

PostgreSQL email validation that is consistent with the HTML5 spec


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4180** | {{< badge content="pg_html5_email_address" link="https://github.com/bigsmoke/pg_html5_email_address" >}} | {{< ext "pg_html5_email_address" >}} | `1.2.3` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_smtp_client" >}} {{< ext "url_encode" >}} {{< ext "pg_render" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_html5_email_address" >}} | `1.2.3` | {{< bg "18" "pg_html5_email_address_18" "green" >}} {{< bg "17" "pg_html5_email_address_17" "green" >}} {{< bg "16" "pg_html5_email_address_16" "green" >}} {{< bg "15" "pg_html5_email_address_15" "green" >}} {{< bg "14" "pg_html5_email_address_14" "green" >}} {{< bg "13" "pg_html5_email_address_13" "green" >}} | `pg_html5_email_address_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_html5_email_address" >}} | `1.2.3` | {{< bg "18" "postgresql-18-pg-html5-email-address" "green" >}} {{< bg "17" "postgresql-17-pg-html5-email-address" "green" >}} {{< bg "16" "postgresql-16-pg-html5-email-address" "green" >}} {{< bg "15" "postgresql-15-pg-html5-email-address" "green" >}} {{< bg "14" "postgresql-14-pg-html5-email-address" "green" >}} {{< bg "13" "postgresql-13-pg-html5-email-address" "green" >}} | `postgresql-$v-pg-html5-email-address` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "pg_html5_email_address_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-html5-email-address : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-html5-email-address : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-html5-email-address : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-html5-email-address : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.3" "postgresql-17-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-pg-html5-email-address : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-13-pg-html5-email-address : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_18` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_18-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_18` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_18-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_18` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_18-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_18` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_18-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_18` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_18-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_18` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_18-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_18-1.2.3-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_17` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_17-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_17` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_17-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_17` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_17-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_17` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_17-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_17` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_17-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_17` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_17-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_17-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `d12.aarch64` | pigsty | 12.3 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-html5-email-address` | 1.2.3 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-17-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_16` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_16-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_16` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_16-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_16` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_16-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_16` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_16-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_16` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_16-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_16` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_16-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_16-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `d12.aarch64` | pigsty | 12.3 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-html5-email-address` | 1.2.3 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-16-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_15` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_15-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_15` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_15-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_15` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_15-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_15` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_15-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_15` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_15-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_15` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_15-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_15-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `d12.aarch64` | pigsty | 12.3 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-html5-email-address` | 1.2.3 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-15-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_14` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_14-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_14` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_14-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_14` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_14-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_14` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_14-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_14` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_14-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_14` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_14-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_14-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `d12.aarch64` | pigsty | 12.3 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-html5-email-address` | 1.2.3 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-14-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_html5_email_address_13` | 1.2.3 | `el8.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_html5_email_address_13-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_html5_email_address_13` | 1.2.3 | `el8.aarch64` | pigsty | 15.4 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_html5_email_address_13-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_html5_email_address_13` | 1.2.3 | `el9.x86_64` | pigsty | 15.4 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_html5_email_address_13-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_html5_email_address_13` | 1.2.3 | `el9.aarch64` | pigsty | 15.3 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_html5_email_address_13-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_html5_email_address_13` | 1.2.3 | `el10.x86_64` | pigsty | 15.5 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_html5_email_address_13-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_html5_email_address_13` | 1.2.3 | `el10.aarch64` | pigsty | 15.5 KiB | [pg_html5_email_address_13-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_html5_email_address_13-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `d12.aarch64` | pigsty | 12.3 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-html5-email-address` | 1.2.3 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-html5-email-address/postgresql-13-pg-html5-email-address_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_html5_email_address" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_html5_email_address" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_html5_email_address-1.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_html5_email_address; # get pg_html5_email_address source code
pig build dep pg_html5_email_address; # install build dependencies
pig build pkg pg_html5_email_address; # build extension rpm or deb
pig build ext pg_html5_email_address; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_html5_email_address; # install by extension name, for the current active PG version
pig ext install pg_html5_email_address; # install via package alias, for the active PG version
pig ext install pg_html5_email_address -v 18;   # install for PG 18
pig ext install pg_html5_email_address -v 17;   # install for PG 17
pig ext install pg_html5_email_address -v 16;   # install for PG 16
pig ext install pg_html5_email_address -v 15;   # install for PG 15
pig ext install pg_html5_email_address -v 14;   # install for PG 14
pig ext install pg_html5_email_address -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_html5_email_address;
```

