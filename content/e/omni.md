---
title: "omni"
linkTitle: "omni"
description: "Advanced adapter for Postgres extensions"
weight: 2940
categories: ["FEAT"]
width: full
---

[**omnigres**](https://github.com/omnigres/omnigres) : Advanced adapter for Postgres extensions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2940** | {{< badge content="omni" link="https://github.com/omnigres/omnigres" >}} | {{< ext "omni" "omnigres" >}} | `0.2.14` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `omni` |
|    **Siblings**   | {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_cloudevents" >}} {{< ext "omni_containers" >}} {{< ext "omni_credentials" >}} {{< ext "omni_email" >}} {{< ext "omni_http" >}} {{< ext "omni_httpc" >}} {{< ext "omni_httpd" >}} {{< ext "omni_id" >}} {{< ext "omni_json" >}} {{< ext "omni_kube" >}} {{< ext "omni_ledger" >}} {{< ext "omni_manifest" >}} {{< ext "omni_mimetypes" >}} {{< ext "omni_os" >}} {{< ext "omni_polyfill" >}} {{< ext "omni_python" >}} {{< ext "omni_regex" >}} {{< ext "omni_rest" >}} {{< ext "omni_schema" >}} {{< ext "omni_seq" >}} {{< ext "omni_service" >}} {{< ext "omni_session" >}} {{< ext "omni_sql" >}} {{< ext "omni_sqlite" >}} {{< ext "omni_test" >}} {{< ext "omni_txn" >}} {{< ext "omni_types" >}} {{< ext "omni_var" >}} {{< ext "omni_vfs" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "omni_web" >}} {{< ext "omni_worker" >}} {{< ext "omni_xml" >}} {{< ext "omni_yaml" >}} |

> [!Note] shared lib name is different from ext name!


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.14` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `omnigres` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.14` | {{< bg "18" "omnigres_18" "green" >}} {{< bg "17" "omnigres_17" "green" >}} {{< bg "16" "omnigres_16" "green" >}} {{< bg "15" "omnigres_15" "green" >}} {{< bg "14" "omnigres_14" "green" >}} {{< bg "13" "omnigres_13" "green" >}} | `omnigres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.14` | {{< bg "18" "postgresql-18-omnigres" "green" >}} {{< bg "17" "postgresql-17-omnigres" "green" >}} {{< bg "16" "postgresql-16-omnigres" "green" >}} {{< bg "15" "postgresql-15-omnigres" "green" >}} {{< bg "14" "postgresql-14-omnigres" "green" >}} {{< bg "13" "postgresql-13-omnigres" "green" >}} | `postgresql-$v-omnigres` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-13-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-13-omnigres : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-18-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-18-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-18-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-18-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-18-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-18-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-18-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-18-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnigres_17` | `20250120` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [omnigres_17-20250120-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnigres_17-20250120-1PIGSTY.el8.x86_64.rpm) |
| `omnigres_17` | `20250120` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [omnigres_17-20250120-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnigres_17-20250120-1PIGSTY.el8.aarch64.rpm) |
| `omnigres_17` | `20250507` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [omnigres_17-20250507-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_17-20250507-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_17` | `20250120` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.4 MiB | [omnigres_17-20250120-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_17-20250120-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_17` | `20250507` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [omnigres_17-20250507-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_17-20250507-1PIGSTY.el9.aarch64.rpm) |
| `omnigres_17` | `20250120` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [omnigres_17-20250120-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_17-20250120-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-omnigres` | `20250120` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.6 MiB | [postgresql-17-omnigres_20250120-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-17-omnigres_20250120-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-omnigres` | `20250120` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-17-omnigres_20250120-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-17-omnigres_20250120-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-17-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-17-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-17-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-17-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-omnigres` | `20250120` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-17-omnigres_20250120-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-17-omnigres_20250120-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-omnigres` | `20250120` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-17-omnigres_20250120-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-17-omnigres_20250120-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-17-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-17-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-17-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-17-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnigres_16` | `20250120` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [omnigres_16-20250120-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnigres_16-20250120-1PIGSTY.el8.x86_64.rpm) |
| `omnigres_16` | `20250120` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [omnigres_16-20250120-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnigres_16-20250120-1PIGSTY.el8.aarch64.rpm) |
| `omnigres_16` | `20250507` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [omnigres_16-20250507-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_16-20250507-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_16` | `20250120` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.4 MiB | [omnigres_16-20250120-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_16-20250120-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_16` | `20250507` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [omnigres_16-20250507-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_16-20250507-1PIGSTY.el9.aarch64.rpm) |
| `omnigres_16` | `20250120` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [omnigres_16-20250120-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_16-20250120-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-omnigres` | `20250120` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.6 MiB | [postgresql-16-omnigres_20250120-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-16-omnigres_20250120-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-omnigres` | `20250120` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-16-omnigres_20250120-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-16-omnigres_20250120-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-16-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-16-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-16-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-16-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-omnigres` | `20250120` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-16-omnigres_20250120-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-16-omnigres_20250120-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-omnigres` | `20250120` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-16-omnigres_20250120-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-16-omnigres_20250120-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-16-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-16-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-16-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-16-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnigres_15` | `20250120` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [omnigres_15-20250120-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnigres_15-20250120-1PIGSTY.el8.x86_64.rpm) |
| `omnigres_15` | `20250120` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [omnigres_15-20250120-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnigres_15-20250120-1PIGSTY.el8.aarch64.rpm) |
| `omnigres_15` | `20250507` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [omnigres_15-20250507-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_15-20250507-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_15` | `20250120` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.4 MiB | [omnigres_15-20250120-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_15-20250120-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_15` | `20250507` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [omnigres_15-20250507-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_15-20250507-1PIGSTY.el9.aarch64.rpm) |
| `omnigres_15` | `20250120` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [omnigres_15-20250120-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_15-20250120-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-omnigres` | `20250120` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.6 MiB | [postgresql-15-omnigres_20250120-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-15-omnigres_20250120-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-omnigres` | `20250120` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-15-omnigres_20250120-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-15-omnigres_20250120-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-15-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-15-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-15-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-15-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-omnigres` | `20250120` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-15-omnigres_20250120-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-15-omnigres_20250120-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-omnigres` | `20250120` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-15-omnigres_20250120-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-15-omnigres_20250120-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-15-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-15-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-15-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-15-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnigres_14` | `20250120` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [omnigres_14-20250120-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnigres_14-20250120-1PIGSTY.el8.x86_64.rpm) |
| `omnigres_14` | `20250120` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [omnigres_14-20250120-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnigres_14-20250120-1PIGSTY.el8.aarch64.rpm) |
| `omnigres_14` | `20250507` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [omnigres_14-20250507-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_14-20250507-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_14` | `20250120` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.4 MiB | [omnigres_14-20250120-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_14-20250120-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_14` | `20250507` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [omnigres_14-20250507-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_14-20250507-1PIGSTY.el9.aarch64.rpm) |
| `omnigres_14` | `20250120` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [omnigres_14-20250120-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_14-20250120-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-omnigres` | `20250120` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.6 MiB | [postgresql-14-omnigres_20250120-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-14-omnigres_20250120-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-omnigres` | `20250120` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-14-omnigres_20250120-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-14-omnigres_20250120-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-14-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-14-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-14-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-14-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-omnigres` | `20250120` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-14-omnigres_20250120-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-14-omnigres_20250120-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-omnigres` | `20250120` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-14-omnigres_20250120-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-14-omnigres_20250120-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-14-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-14-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-14-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-14-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnigres_13` | `20250120` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [omnigres_13-20250120-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnigres_13-20250120-1PIGSTY.el8.x86_64.rpm) |
| `omnigres_13` | `20250120` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [omnigres_13-20250120-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnigres_13-20250120-1PIGSTY.el8.aarch64.rpm) |
| `omnigres_13` | `20250507` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [omnigres_13-20250507-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_13-20250507-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_13` | `20250120` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.4 MiB | [omnigres_13-20250120-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnigres_13-20250120-1PIGSTY.el9.x86_64.rpm) |
| `omnigres_13` | `20250507` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [omnigres_13-20250507-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_13-20250507-1PIGSTY.el9.aarch64.rpm) |
| `omnigres_13` | `20250120` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [omnigres_13-20250120-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnigres_13-20250120-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-omnigres` | `20250120` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.6 MiB | [postgresql-13-omnigres_20250120-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-13-omnigres_20250120-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-omnigres` | `20250120` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-13-omnigres_20250120-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnigres/postgresql-13-omnigres_20250120-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-omnigres` | `20251108` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-13-omnigres_20251108-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-13-omnigres_20251108-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-omnigres` | `20251108` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-13-omnigres_20251108-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/omnigres/postgresql-13-omnigres_20251108-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-omnigres` | `20250120` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-13-omnigres_20250120-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-13-omnigres_20250120-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-omnigres` | `20250120` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-13-omnigres_20250120-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnigres/postgresql-13-omnigres_20250120-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-omnigres` | `20251108` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-13-omnigres_20251108-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-13-omnigres_20251108-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-omnigres` | `20251108` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-13-omnigres_20251108-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnigres/postgresql-13-omnigres_20251108-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/omnigres/omnigres" title="Repository" icon="github" subtitle="github.com/omnigres/omnigres" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="omnigres-20251108.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg omnigres;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install omnigres;		# install via package name, for the active PG version
pig install omni;		# install by extension name, for the current active PG version

pig install omni -v 18;   # install for PG 18
pig install omni -v 17;   # install for PG 17
pig install omni -v 16;   # install for PG 16
pig install omni -v 15;   # install for PG 15
pig install omni -v 14;   # install for PG 14
pig install omni -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'omni--0.2.14.so';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION omni;
```


## Usage

Check https://docs.omnigres.org/ for details



## Build

This extension require gcc >= 14 and CMake > 3.25 to build

So actually it is only available on Debian 13, Ubuntu 24.0