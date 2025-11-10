---
title: "aws_s3"
linkTitle: "aws_s3"
description: "aws_s3 postgres extension to import/export data from/to s3"
weight: 8800
categories: ["FDW"]
width: full
---

[**aws_s3**](https://github.com/chimpler/postgres-aws-s3) : aws_s3 postgres extension to import/export data from/to s3


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8800** | {{< badge content="aws_s3" link="https://github.com/chimpler/postgres-aws-s3" >}} | {{< ext "aws_s3" >}} | `0.0.1` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_parquet" >}} {{< ext "hdfs_fdw" >}} {{< ext "file_fdw" >}} {{< ext "duckdb_fdw" >}} {{< ext "wrappers" >}} {{< ext "pg_bulkload" >}} {{< ext "columnar" >}} {{< ext "pg_analytics" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `aws_s3` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "aws_s3_18" "green" >}} {{< bg "17" "aws_s3_17" "green" >}} {{< bg "16" "aws_s3_16" "green" >}} {{< bg "15" "aws_s3_15" "green" >}} {{< bg "14" "aws_s3_14" "green" >}} {{< bg "13" "aws_s3_13" "green" >}} | `aws_s3_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-aws-s3" "green" >}} {{< bg "17" "postgresql-17-aws-s3" "green" >}} {{< bg "16" "postgresql-16-aws-s3" "green" >}} {{< bg "15" "postgresql-15-aws-s3" "green" >}} {{< bg "14" "postgresql-14-aws-s3" "green" >}} {{< bg "13" "postgresql-13-aws-s3" "green" >}} | `postgresql-$v-aws-s3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "aws_s3_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-aws-s3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-aws-s3 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-18-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-18-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-17-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-17-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-16-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-16-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-15-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-15-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-14-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-14-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aws_s3_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.3 KiB | [aws_s3_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aws_s3_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `aws_s3_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [aws_s3_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aws_s3_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `aws_s3_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.3 KiB | [aws_s3_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aws_s3_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `aws_s3_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [aws_s3_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aws_s3_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `aws_s3_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.4 KiB | [aws_s3_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aws_s3_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `aws_s3_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.3 KiB | [aws_s3_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aws_s3_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-aws-s3` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.4 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.4 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.4 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.3 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.3 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.3 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-aws-s3` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-13-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aws-s3/postgresql-13-aws-s3_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/chimpler/postgres-aws-s3" title="Repository" icon="github" subtitle="github.com/chimpler/postgres-aws-s3" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="aws_s3-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg aws_s3;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install aws_s3;		# install via package name, for the active PG version

pig install aws_s3 -v 18;   # install for PG 18
pig install aws_s3 -v 17;   # install for PG 17
pig install aws_s3 -v 16;   # install for PG 16
pig install aws_s3 -v 15;   # install for PG 15
pig install aws_s3 -v 14;   # install for PG 14
pig install aws_s3 -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION aws_s3;
```
