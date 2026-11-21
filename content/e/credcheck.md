---
title: "credcheck"
linkTitle: "credcheck"
description: "credcheck - postgresql plain text credential checker"
weight: 7310
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**credcheck**](https://github.com/HexaCluster/credcheck) : credcheck - postgresql plain text credential checker


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7310** | {{< badge content="credcheck" link="https://github.com/HexaCluster/credcheck" >}} | {{< ext "credcheck" >}} | `6.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_pwhash" >}} {{< ext "passwordcheck" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "passwordpolicy" >}} {{< ext "chkpass" >}} {{< ext "pg_enigma" >}} {{< ext "column_encrypt" >}} |

> [!Note] pkg 6.0; SQL 6.0.0; restart then update extension; migrates 5.x history; PgBouncer needs setup.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `6.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `credcheck` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `6.0` | {{< bg "19" "credcheck_19" "green" >}} {{< bg "18" "credcheck_18" "green" >}} {{< bg "17" "credcheck_17" "green" >}} {{< bg "16" "credcheck_16" "green" >}} {{< bg "15" "credcheck_15" "green" >}} | `credcheck_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `6.0` | {{< bg "19" "postgresql-19-credcheck" "green" >}} {{< bg "18" "postgresql-18-credcheck" "green" >}} {{< bg "17" "postgresql-17-credcheck" "green" >}} {{< bg "16" "postgresql-16-credcheck" "green" >}} {{< bg "15" "postgresql-15-credcheck" "green" >}} | `postgresql-$v-credcheck` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "credcheck_19 : N/A 0" "gray" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 11" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 15" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 20" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "credcheck_19 : N/A 0" "gray" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 11" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 15" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 20" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 6.0" "credcheck_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 6.0" "credcheck_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 6.0" "credcheck_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 6.0" "credcheck_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_18 : AVAIL 17" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_17 : AVAIL 17" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_16 : AVAIL 17" "blue" >}} | {{< bg "PGDG 6.0" "credcheck_15 : AVAIL 17" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 6.0" "postgresql-19-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 6.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_19` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.5 KiB | [credcheck_19-6.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/credcheck_19-6.0-2PGDG.rhel9.9.x86_64.rpm) |
| `credcheck_19` | `6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.8 KiB | [credcheck_19-6.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/credcheck_19-6.0-2PGDG.rhel9.9.aarch64.rpm) |
| `credcheck_19` | `6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.1 KiB | [credcheck_19-6.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/credcheck_19-6.0-2PGDG.rhel10.3.x86_64.rpm) |
| `credcheck_19` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.1 KiB | [credcheck_19-6.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/credcheck_19-6.0-2PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.8 KiB | [postgresql-19-credcheck_6.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg12+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.6 KiB | [postgresql-19-credcheck_6.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg12+1_amd64.deb) |
| `postgresql-19-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-19-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.7 KiB | [postgresql-19-credcheck_6.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg12+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.7 KiB | [postgresql-19-credcheck_6.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg12+1_arm64.deb) |
| `postgresql-19-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.7 KiB | [postgresql-19-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.7 KiB | [postgresql-19-credcheck_6.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg13+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.7 KiB | [postgresql-19-credcheck_6.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg13+1_amd64.deb) |
| `postgresql-19-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.9 KiB | [postgresql-19-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 105.1 KiB | [postgresql-19-credcheck_6.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg13+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.8 KiB | [postgresql-19-credcheck_6.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg13+1_arm64.deb) |
| `postgresql-19-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.5 KiB | [postgresql-19-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 99.4 KiB | [postgresql-19-credcheck_6.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 99.2 KiB | [postgresql-19-credcheck_6.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.5 KiB | [postgresql-19-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.3 KiB | [postgresql-19-credcheck_6.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.1 KiB | [postgresql-19-credcheck_6.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.0 KiB | [postgresql-19-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.3 KiB | [postgresql-19-credcheck_6.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.2 KiB | [postgresql-19-credcheck_6.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.2 KiB | [postgresql-19-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.6 KiB | [postgresql-19-credcheck_6.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.7 KiB | [postgresql-19-credcheck_6.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-19-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.2 KiB | [postgresql-19-credcheck_6.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 96.9 KiB | [postgresql-19-credcheck_6.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-19-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.9 KiB | [postgresql-19-credcheck_6.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.9 KiB | [postgresql-19-credcheck_6.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_6.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-19-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-19-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-19-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_18` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [credcheck_18-6.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-6.0-2PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.2 KiB | [credcheck_18-6.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/credcheck_18-6.0-1PGSTY.el8.x86_64.rpm) |
| `credcheck_18` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [credcheck_18-6.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-6.0-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.3 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.5 KiB | [credcheck_18-6.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-6.0-2PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.6 KiB | [credcheck_18-6.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/credcheck_18-6.0-1PGSTY.el8.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.2 KiB | [credcheck_18-6.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-6.0-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.5 KiB | [credcheck_18-6.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-6.0-2PGDG.rhel9.9.x86_64.rpm) |
| `credcheck_18` | `6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.8 KiB | [credcheck_18-6.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-6.0-2PGDG.rhel9.9.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.2 KiB | [credcheck_18-6.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-6.0-2PGDG.rhel10.3.x86_64.rpm) |
| `credcheck_18` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.1 KiB | [credcheck_18-6.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-6.0-2PGDG.rhel10.3.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.1 KiB | [credcheck_18-6.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-6.0-2PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_18` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [credcheck_18-6.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-6.0-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.9 KiB | [postgresql-18-credcheck_6.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg12+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.7 KiB | [postgresql-18-credcheck_6.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.8 KiB | [postgresql-18-credcheck_6.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg12+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.6 KiB | [postgresql-18-credcheck_6.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.7 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.7 KiB | [postgresql-18-credcheck_6.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg13+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.6 KiB | [postgresql-18-credcheck_6.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 105.1 KiB | [postgresql-18-credcheck_6.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg13+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.9 KiB | [postgresql-18-credcheck_6.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 99.4 KiB | [postgresql-18-credcheck_6.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 99.2 KiB | [postgresql-18-credcheck_6.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.6 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.1 KiB | [postgresql-18-credcheck_6.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.1 KiB | [postgresql-18-credcheck_6.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.2 KiB | [postgresql-18-credcheck_6.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.3 KiB | [postgresql-18-credcheck_6.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.7 KiB | [postgresql-18-credcheck_6.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.4 KiB | [postgresql-18-credcheck_6.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.3 KiB | [postgresql-18-credcheck_6.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.0 KiB | [postgresql-18-credcheck_6.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.5 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.1 KiB | [postgresql-18-credcheck_6.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.0 KiB | [postgresql-18-credcheck_6.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_6.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.0 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_17` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [credcheck_17-6.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-6.0-2PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.3 KiB | [credcheck_17-6.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/credcheck_17-6.0-1PGSTY.el8.x86_64.rpm) |
| `credcheck_17` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.5 KiB | [credcheck_17-6.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-6.0-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.1 KiB | [credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.6 KiB | [credcheck_17-6.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-6.0-2PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.7 KiB | [credcheck_17-6.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/credcheck_17-6.0-1PGSTY.el8.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [credcheck_17-6.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-6.0-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.7 KiB | [credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.5 KiB | [credcheck_17-6.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-6.0-2PGDG.rhel9.9.x86_64.rpm) |
| `credcheck_17` | `6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.9 KiB | [credcheck_17-6.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-6.0-2PGDG.rhel9.9.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.2 KiB | [credcheck_17-6.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-6.0-2PGDG.rhel10.3.x86_64.rpm) |
| `credcheck_17` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [credcheck_17-6.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-6.0-2PGDG.rhel10.3.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [credcheck_17-6.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-6.0-2PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_17` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.9 KiB | [credcheck_17-6.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-6.0-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.7 KiB | [postgresql-17-credcheck_6.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg12+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.6 KiB | [postgresql-17-credcheck_6.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.8 KiB | [postgresql-17-credcheck_6.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg12+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.7 KiB | [postgresql-17-credcheck_6.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.5 KiB | [postgresql-17-credcheck_6.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg13+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.6 KiB | [postgresql-17-credcheck_6.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.8 KiB | [postgresql-17-credcheck_6.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg13+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.9 KiB | [postgresql-17-credcheck_6.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.8 KiB | [postgresql-17-credcheck_6.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.6 KiB | [postgresql-17-credcheck_6.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.9 KiB | [postgresql-17-credcheck_6.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.7 KiB | [postgresql-17-credcheck_6.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.7 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.2 KiB | [postgresql-17-credcheck_6.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.0 KiB | [postgresql-17-credcheck_6.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.6 KiB | [postgresql-17-credcheck_6.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.5 KiB | [postgresql-17-credcheck_6.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.2 KiB | [postgresql-17-credcheck_6.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.0 KiB | [postgresql-17-credcheck_6.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.9 KiB | [postgresql-17-credcheck_6.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.8 KiB | [postgresql-17-credcheck_6.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_6.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_16` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [credcheck_16-6.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-6.0-2PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.3 KiB | [credcheck_16-6.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/credcheck_16-6.0-1PGSTY.el8.x86_64.rpm) |
| `credcheck_16` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.5 KiB | [credcheck_16-6.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-6.0-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.6 KiB | [credcheck_16-6.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-6.0-2PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.7 KiB | [credcheck_16-6.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/credcheck_16-6.0-1PGSTY.el8.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [credcheck_16-6.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-6.0-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.5 KiB | [credcheck_16-6.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-6.0-2PGDG.rhel9.9.x86_64.rpm) |
| `credcheck_16` | `6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.9 KiB | [credcheck_16-6.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-6.0-2PGDG.rhel9.9.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.2 KiB | [credcheck_16-6.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-6.0-2PGDG.rhel10.3.x86_64.rpm) |
| `credcheck_16` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [credcheck_16-6.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-6.0-2PGDG.rhel10.3.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [credcheck_16-6.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-6.0-2PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_16` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.9 KiB | [credcheck_16-6.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-6.0-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.8 KiB | [postgresql-16-credcheck_6.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg12+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.6 KiB | [postgresql-16-credcheck_6.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.9 KiB | [postgresql-16-credcheck_6.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg12+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.6 KiB | [postgresql-16-credcheck_6.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.7 KiB | [postgresql-16-credcheck_6.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg13+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.6 KiB | [postgresql-16-credcheck_6.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.7 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 105.0 KiB | [postgresql-16-credcheck_6.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg13+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.9 KiB | [postgresql-16-credcheck_6.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.6 KiB | [postgresql-16-credcheck_6.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.6 KiB | [postgresql-16-credcheck_6.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.7 KiB | [postgresql-16-credcheck_6.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.5 KiB | [postgresql-16-credcheck_6.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.2 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.2 KiB | [postgresql-16-credcheck_6.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.1 KiB | [postgresql-16-credcheck_6.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.8 KiB | [postgresql-16-credcheck_6.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.5 KiB | [postgresql-16-credcheck_6.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.5 KiB | [postgresql-16-credcheck_6.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.3 KiB | [postgresql-16-credcheck_6.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.0 KiB | [postgresql-16-credcheck_6.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.0 KiB | [postgresql-16-credcheck_6.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_6.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_15` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.9 KiB | [credcheck_15-6.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-6.0-2PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.1 KiB | [credcheck_15-6.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/credcheck_15-6.0-1PGSTY.el8.x86_64.rpm) |
| `credcheck_15` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.6 KiB | [credcheck_15-6.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-6.0-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.4 KiB | [credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [credcheck_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [credcheck_15-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [credcheck_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [credcheck_15-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [credcheck_15-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.6 KiB | [credcheck_15-6.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-6.0-2PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.2 KiB | [credcheck_15-6.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/credcheck_15-6.0-1PGSTY.el8.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.2 KiB | [credcheck_15-6.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-6.0-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.5 KiB | [credcheck_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [credcheck_15-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [credcheck_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [credcheck_15-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.9 KiB | [credcheck_15-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.6 KiB | [credcheck_15-6.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-6.0-2PGDG.rhel9.9.x86_64.rpm) |
| `credcheck_15` | `6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.9 KiB | [credcheck_15-6.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-6.0-2PGDG.rhel9.9.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.1 KiB | [credcheck_15-6.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-6.0-2PGDG.rhel10.3.x86_64.rpm) |
| `credcheck_15` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.1 KiB | [credcheck_15-6.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-6.0-2PGDG.rhel10.3.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [credcheck_15-6.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-6.0-2PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_15` | `6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [credcheck_15-6.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-6.0-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.6 KiB | [postgresql-15-credcheck_6.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg12+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 106.4 KiB | [postgresql-15-credcheck_6.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.4 KiB | [postgresql-15-credcheck_6.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg12+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 104.5 KiB | [postgresql-15-credcheck_6.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.7 KiB | [postgresql-15-credcheck_6.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg13+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 106.5 KiB | [postgresql-15-credcheck_6.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.8 KiB | [postgresql-15-credcheck_6.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg13+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 104.5 KiB | [postgresql-15-credcheck_6.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.1 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.5 KiB | [postgresql-15-credcheck_6.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.2 KiB | [postgresql-15-credcheck_6.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.2 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.4 KiB | [postgresql-15-credcheck_6.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.2 KiB | [postgresql-15-credcheck_6.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.0 KiB | [postgresql-15-credcheck_6.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 97.8 KiB | [postgresql-15-credcheck_6.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.2 KiB | [postgresql-15-credcheck_6.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.1 KiB | [postgresql-15-credcheck_6.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.5 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.0 KiB | [postgresql-15-credcheck_6.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.0 KiB | [postgresql-15-credcheck_6.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.7 KiB | [postgresql-15-credcheck_6.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.7 KiB | [postgresql-15-credcheck_6.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_6.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.7 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/credcheck" title="Repository" icon="github" subtitle="github.com/HexaCluster/credcheck" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install credcheck;		# install via package name, for the active PG version

pig install credcheck -v 19;   # install for PG 19
pig install credcheck -v 18;   # install for PG 18
pig install credcheck -v 17;   # install for PG 17
pig install credcheck -v 16;   # install for PG 16
pig install credcheck -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'credcheck';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION credcheck;
```

## Usage

Sources:

- [v6.0 README](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/README.md)
- [v6.0 changelog](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/ChangeLog)
- [SQL objects 6.0.0](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/sql/credcheck--6.0.0.sql)
- [Password history WAL implementation](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/credcheck.c)
- [Login-event setup](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/event_trigger.sql)

`credcheck` enforces username and plaintext-password rules during role creation, password changes and role renames. It also tracks password reuse, bans repeated authentication failures and can require a password change at first login. Configure policy as a superuser; the defaults do not enforce a comprehensive password-strength policy.

### Enable and Set Policy

Add the library to the existing preload list and restart PostgreSQL. Install SQL objects in each database where administrators need its views and reset functions:

```ini
shared_preload_libraries = 'credcheck'
credcheck.password_min_length = 12
credcheck.password_contain_username = on
credcheck.password_reuse_history = 2
credcheck.password_reuse_interval = 365
```

```sql
CREATE EXTENSION credcheck;
CREATE ROLE app_user LOGIN PASSWORD 'example-Strong-Pass#123';
SELECT rolename, password_date FROM pg_password_history;
```

The interval is in days. Installing the SQL extension is separate from loading the server-wide hooks. Upstream release 6.0 uses SQL extension version 6.0.0; installing upgraded files requires a restart to reload the library.

### Policy Index

| Settings | Purpose |
|---|---|
| `credcheck.username_min_length`, `credcheck.username_min_special`, `credcheck.username_min_digit`, `credcheck.username_min_upper`, `credcheck.username_min_lower` | Username length and character requirements |
| `credcheck.password_min_length`, `credcheck.password_min_special`, `credcheck.password_min_digit`, `credcheck.password_min_upper`, `credcheck.password_min_lower` | Password length and character requirements |
| `credcheck.username_min_repeat`, `credcheck.password_min_repeat` | Maximum adjacent repetitions, despite the parameter names |
| `credcheck.username_contain`, `credcheck.username_not_contain`, `credcheck.password_contain`, `credcheck.password_not_contain` | Required or forbidden content |
| `credcheck.username_contain_password`, `credcheck.password_contain_username` | Reject credentials containing one another |
| `credcheck.username_ignore_case`, `credcheck.password_ignore_case` | Case handling |
| `credcheck.password_min_length_su`, `credcheck.password_valid_until_su` | Separate superuser requirements |
| `credcheck.password_valid_until`, `credcheck.password_valid_max` | Minimum and maximum password lifetime; the minimum also supplies an omitted expiry when a password is changed |
| `credcheck.whitelist`, `credcheck.superuser_nocheck` | Explicit policy exemptions |
| `credcheck.no_password_logging` | Suppress passwords in policy-error logs; enabled by default |
{.wide .expand}

CrackLib strength checking is available only when the library was built with that support and its dictionary is available.

### Password History and Replication

History contains SHA-256 password hashes, is shared across databases, and is persisted in `$PGDATA/pg_password_history`. Include that file in backup planning and protect access to the SQL history view, which in 6.0 is revoked from PUBLIC and granted to pg_read_all_stats. `credcheck.history_max_size` changes the history capacity and requires a restart.

Version 5.0 replicates history changes through custom WAL resource manager ID 150 on PostgreSQL 15 and later. Keep the matching library preloaded on replicas and recovery servers that replay its WAL. Earlier PostgreSQL versions retain the file-backed history path without this replication support. History-reset and timestamp-test functions reject execution during recovery.

```sql
SELECT pg_password_history_reset('app_user');
```

Resetting history removes reuse protection for those records; reserve it for administrators.

### Authentication and Password Changes

```ini
credcheck.max_auth_failure = 3
credcheck.auth_delay_ms = 1000
credcheck.whitelist_auth_failure = 'service_user'
credcheck.password_change_first_login = true
```

```sql
SELECT * FROM pg_banned_role;
SELECT pg_banned_role_reset('app_user');
ALTER ROLE app_user SET credcheck_internal.force_change_password = true;
```

Bans remain until reset and their cache is lost at restart. `credcheck.reset_superuser` provides the documented superuser recovery path; `credcheck.auth_failure_cache_size` requires a restart.

Use the actual parameter `credcheck.disallow_change_password` to prohibit password changes. Even superusers are affected unless they enable `credcheck.superuser_nocheck` in their session. This exemption bypasses all corresponding role checks and must be controlled.

`credcheck.password_valid_warning` needs PostgreSQL 17 or later and the official login event trigger installed separately in every relevant database; SQL extension creation does not install that trigger.

### Plaintext Boundary

Strength and reuse checks need plaintext at password-change time. Already hashed passwords are rejected by default, including passwords sent by psql's `\password`. Setting `credcheck.encrypted_password_allowed` accepts them without providing equivalent plaintext checks. Protect the password-change connection and do not assume existing credentials are scanned retroactively. Username checks are skipped when creating a role without a password or renaming a role with no password.

### Upgrade and New Interfaces in 6.0

6.0 moves password history into an integer-indexed hash table in an mmap file. The first upgraded startup imports 5.x history and renames the old file to pg_password_history.legacy-migrated; this happens on both primary and standby. Preserve and verify the history backup before upgrading. After installing matching libraries and restarting, run `ALTER EXTENSION credcheck UPDATE TO '6.0.0'` in databases using its SQL objects. Current SQL revokes PUBLIC access to the history function/view and grants view access to pg_read_all_stats; review existing custom grants.

`credcheck.password_valid_min` sets a short initial password lifetime for new roles; the default 0 disables it. Existing roles changing their password retain the previous lifetime policy. New `pg_check_password(username, password)` checks candidate credentials, raising an error on a violation and returning true on acceptance; it does not evaluate password reuse history.

`credcheck.lastlog` defaults off. When enabled, pg_lastlog exposes sessions and server startup/shutdown events; optional SQL tracking records query text, so define access and retention deliberately. PgBouncer integration for failed-attempt limits and first-login password changes needs the documented extra SQL and pooler hooks; preload alone does not configure it. `credcheck.no_password_logging` now masks password literals while retaining the error statement structure. This release also fixes self-ALTER ROLE privilege escalation and memory defects.
