---
title: "tin"
linkTitle: "tin"
description: "TIN-compatible text search for development and testing"
weight: 2310
categories: ["FTS"]
languages: ["Rust"]
licenses: ["AGPL-3.0-or-later"]
repos: ["PIGSTY"]
page_width: full
---

[**lead**](https://github.com/planetscale/lead) : TIN-compatible text search for development and testing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2310** | {{< badge content="tin" link="https://github.com/planetscale/lead" >}} | {{< ext "tin" "lead" >}} | `1.0.6` | {{< category "FTS" >}} | {{< license "AGPL-3.0-or-later" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tin` |
|   **See Also**    | {{< ext "stannum" >}} {{< ext "pg_search" >}} |

> [!Note] Untagged 1.0.6 source snapshot for development/CI, not production PlanetScale TIN. Conflicts with stannum in the same database.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `lead` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "lead_18" "green" >}} {{< bg "17" "lead_17" "green" >}} {{< bg "16" "lead_16" "red" >}} {{< bg "15" "lead_15" "red" >}} {{< bg "14" "lead_14" "red" >}} | `lead_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "postgresql-18-lead" "green" >}} {{< bg "17" "postgresql-17-lead" "green" >}} {{< bg "16" "postgresql-16-lead" "red" >}} {{< bg "15" "postgresql-15-lead" "red" >}} {{< bg "14" "postgresql-14-lead" "red" >}} | `postgresql-$v-lead` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "lead_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "lead_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "lead_16 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_15 : N/A 0" "gray" >}} | {{< bg "N/A" "lead_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-18-lead : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.6" "postgresql-17-lead : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-lead : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-lead : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lead_18` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [lead_18-1.0.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lead_18-1.0.6-1PGSTY.el8.x86_64.rpm) |
| `lead_18` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [lead_18-1.0.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lead_18-1.0.6-1PGSTY.el8.aarch64.rpm) |
| `lead_18` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [lead_18-1.0.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lead_18-1.0.6-1PGSTY.el9.x86_64.rpm) |
| `lead_18` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.2 MiB | [lead_18-1.0.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lead_18-1.0.6-1PGSTY.el9.aarch64.rpm) |
| `lead_18` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.3 MiB | [lead_18-1.0.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lead_18-1.0.6-1PGSTY.el10.x86_64.rpm) |
| `lead_18` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.2 MiB | [lead_18-1.0.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lead_18-1.0.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-lead` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-18-lead_1.0.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-lead` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.7 MiB | [postgresql-18-lead_1.0.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-lead` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-18-lead_1.0.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-lead` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.7 MiB | [postgresql-18-lead_1.0.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.2 MiB | [postgresql-18-lead_1.0.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-18-lead_1.0.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.2 MiB | [postgresql-18-lead_1.0.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~noble_amd64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-18-lead_1.0.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~noble_arm64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-18-lead_1.0.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-lead` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.0 MiB | [postgresql-18-lead_1.0.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/lead/postgresql-18-lead_1.0.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lead_17` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [lead_17-1.0.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lead_17-1.0.6-1PGSTY.el8.x86_64.rpm) |
| `lead_17` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [lead_17-1.0.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lead_17-1.0.6-1PGSTY.el8.aarch64.rpm) |
| `lead_17` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [lead_17-1.0.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lead_17-1.0.6-1PGSTY.el9.x86_64.rpm) |
| `lead_17` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.2 MiB | [lead_17-1.0.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lead_17-1.0.6-1PGSTY.el9.aarch64.rpm) |
| `lead_17` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.3 MiB | [lead_17-1.0.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lead_17-1.0.6-1PGSTY.el10.x86_64.rpm) |
| `lead_17` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.2 MiB | [lead_17-1.0.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lead_17-1.0.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-lead` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-17-lead_1.0.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-lead` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.7 MiB | [postgresql-17-lead_1.0.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-lead` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-17-lead_1.0.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-lead` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.7 MiB | [postgresql-17-lead_1.0.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.2 MiB | [postgresql-17-lead_1.0.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.0 MiB | [postgresql-17-lead_1.0.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.2 MiB | [postgresql-17-lead_1.0.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~noble_amd64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.0 MiB | [postgresql-17-lead_1.0.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~noble_arm64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-17-lead_1.0.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-lead` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.0 MiB | [postgresql-17-lead_1.0.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/lead/postgresql-17-lead_1.0.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/planetscale/lead" title="Repository" icon="github" subtitle="github.com/planetscale/lead" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lead-1.0.6.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg lead;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install lead;		# install via package name, for the active PG version
pig install tin;		# install by extension name, for the current active PG version

pig install tin -v 18;   # install for PG 18
pig install tin -v 17;   # install for PG 17

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION tin;
```

## Usage

Sources:

- [README](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/README.md)
- [Control file](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/postgres/tin.control)
- [Cargo.toml](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/Cargo.toml)
- [postgres/src/lib.rs](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/postgres/src/lib.rs)
- [PlanetScale TIN and Lead](https://planetscale.com/docs/postgres/search)

`tin` is the extension name of Lead, a TIN-compatible text-search implementation for development, CI, and staging. The reviewed 1.0.6 source snapshot deliberately trades performance for a small reference implementation and is unsuitable for production search.

### Core Workflow

```sql
CREATE EXTENSION tin;
CREATE TABLE documents (id integer, body text);
INSERT INTO documents VALUES (1, 'craft beer'), (2, 'wine'), (3, 'beer festival');
CREATE INDEX documents_search ON documents USING tin (body);
SELECT id, tin.full_score(ctid) AS score
FROM documents WHERE body ==> 'beer' ORDER BY score DESC;
```

### Objects and Behavior

The `tin` access method and `==>` operator evaluate TINQL predicates. `tin.score`, `tin.full_score`, `tin.max_score`, and `tin.score_inspect` expose scoring; `tin.highlight` and `tin.highlight_ansi` mark matches. Scoring calls bind to the matching search expression and index.

Every index scan returns all heap pages as lossy candidates. PostgreSQL rechecks the visible tuples, expressions, and partial-index predicates. The index stores no search data; scoring may rescan visible heap rows. This preserves query semantics without providing a production inverted index.

### Requirements

The reviewed source targets PostgreSQL 17 and 18. Installation requires a superuser; objects use the fixed `tin` schema. The library loads on demand and needs no preload or restart. Keep realistic data volumes in mind when using it in tests.

PlanetScale documents TIN as its production search implementation and Lead as the local development substitute. Both use the canonical extension name `tin`; this catalog entry tracks the open-source Lead implementation. Its source-snapshot version is not a claim about the version deployed by PlanetScale.

The current source adds a trailing `stemmer` argument to `tin.tokenize`, `tin.ql_parse`, `tin.highlight`, and `tin.highlight_ansi`, while keeping the earlier signatures under names ending in `_v1_0_3`. Changing an index's `stemmer` requires `REINDEX`. A score call must be at the same query level as its matching `==>` predicate.
