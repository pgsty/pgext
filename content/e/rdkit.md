---
title: "rdkit"
linkTitle: "rdkit"
description: "Cheminformatics functionality for PostgreSQL."
weight: 2930
categories: ["FEAT"]
width: full
---

[**rdkit**](https://github.com/rdkit/rdkit) : Cheminformatics functionality for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2930** | {{< badge content="rdkit" link="https://github.com/rdkit/rdkit" >}} | {{< ext "rdkit" >}} | `202503.1` | {{< category "FEAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] u24 has rdkit for pg17


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `202503.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `rdkit` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `202503.1` | {{< bg "18" "postgresql-18-rdkit" "green" >}} {{< bg "17" "postgresql-17-rdkit" "green" >}} {{< bg "16" "postgresql-16-rdkit" "green" >}} {{< bg "15" "postgresql-15-rdkit" "green" >}} {{< bg "14" "postgresql-14-rdkit" "green" >}} | `postgresql-$v-rdkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.6 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 393.5 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 384.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.1 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 394.5 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 385.2 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.8 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.0 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 394.1 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 385.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.5 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 242.9 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rdkit/rdkit" title="Repository" icon="github" subtitle="github.com/rdkit/rdkit" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install rdkit;		# install via package name, for the active PG version

pig install rdkit -v 18;   # install for PG 18
pig install rdkit -v 17;   # install for PG 17
pig install rdkit -v 16;   # install for PG 16
pig install rdkit -v 15;   # install for PG 15
pig install rdkit -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION rdkit;
```




## Usage

> [rdkit: Cheminformatics and molecule toolkit PostgreSQL cartridge](https://github.com/rdkit/rdkit)

The RDKit PostgreSQL cartridge provides the `mol` datatype for molecules, `fp` datatype for fingerprints, substructure and similarity search operations, and GiST index support.

```sql
CREATE EXTENSION rdkit;
```

### Data Types

| Type | Description |
|------|-------------|
| `mol` | Molecular structure (from SMILES, SMARTS, etc.) |
| `bfp` | Bit vector fingerprint |
| `sfp` | Sparse (count) fingerprint |

### Molecule Input/Output

```sql
-- Create molecule from SMILES
SELECT 'c1ccccc1'::mol;

-- Check if SMILES is valid
SELECT is_valid_smiles('c1ccccc1');

-- Convert molecule to SMILES
SELECT mol_to_smiles('c1ccccc1'::mol);
```

### Substructure Search

```sql
-- Substructure match operator
SELECT 'c1ccccc1O'::mol @> 'c1ccccc1'::mol;   -- true (phenol contains benzene)
SELECT 'c1ccccc1'::mol <@ 'c1ccccc1O'::mol;    -- true

-- Using SMARTS patterns
SELECT 'c1ccccc1O'::mol @> 'c1ccc(O)cc1'::mol;
```

### Similarity Search

```sql
-- Tanimoto similarity (returns value between 0 and 1)
SELECT tanimoto_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));

-- Dice similarity
SELECT dice_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));
```

### Fingerprint Functions

```sql
-- Morgan fingerprint (bit vector)
SELECT morganbv_fp('c1ccccc1'::mol);
SELECT morganbv_fp('c1ccccc1'::mol, 2);  -- radius=2

-- RDKit fingerprint
SELECT rdkit_fp('c1ccccc1'::mol);

-- Topological torsion fingerprint
SELECT torsionbv_fp('c1ccccc1'::mol);

-- Atom pair fingerprint
SELECT atompairbv_fp('c1ccccc1'::mol);
```

### Descriptor Calculations

```sql
SELECT mol_amw('c1ccccc1'::mol);          -- average molecular weight
SELECT mol_logp('c1ccccc1'::mol);         -- LogP
SELECT mol_hba('c1ccccc1O'::mol);         -- H-bond acceptors
SELECT mol_hbd('c1ccccc1O'::mol);         -- H-bond donors
SELECT mol_numrotatablebonds('c1ccccc1'::mol); -- rotatable bonds
SELECT mol_numatoms('c1ccccc1'::mol);     -- number of atoms
SELECT mol_numheavyatoms('c1ccccc1'::mol);    -- heavy atoms
SELECT mol_numrings('c1ccccc1'::mol);     -- number of rings
```

### GiST Index Support

Create indexes for fast substructure and similarity searches:

```sql
-- Substructure search index
CREATE INDEX idx_mol ON molecules USING gist(molecule);

-- Fingerprint similarity index
CREATE INDEX idx_fp ON molecules USING gist(morganbv_fp(molecule));
```

Query with index support:

```sql
-- Substructure search
SELECT * FROM molecules WHERE molecule @> 'c1ccccc1'::mol;

-- Similarity search (with threshold)
SET rdkit.dice_threshold = 0.5;
SELECT * FROM molecules WHERE morganbv_fp(molecule) % morganbv_fp('c1ccccc1O'::mol);
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `rdkit.tanimoto_threshold` | 0.5 | Threshold for Tanimoto similarity operator `<%>` |
| `rdkit.dice_threshold` | 0.5 | Threshold for Dice similarity operator `%` |
