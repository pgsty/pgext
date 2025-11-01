---
title: "pgroonga_database"
linkTitle: "pgroonga_database"
description: "PGroonga database management module"
weight: 2111
categories: ["FTS"]
width: full
---

PGroonga database management module


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2111** | {{< badge content="pgroonga_database" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga_database" "pgroonga" >}} | `4.0.4` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} {{< ext "unaccent" >}} |
|    **Siblings**   | {{< ext "pgroonga" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.4` | {{< bg "18" "pgroonga_18*" "red" >}} {{< bg "17" "pgroonga_17*" "green" >}} {{< bg "16" "pgroonga_16*" "green" >}} {{< bg "15" "pgroonga_15*" "green" >}} {{< bg "14" "pgroonga_14*" "green" >}} {{< bg "13" "pgroonga_13*" "green" >}} | `pgroonga_$v*` | `groonga-libs` |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.4` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} {{< bg "13" "postgresql-13-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgroonga : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgroonga : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgroonga_database; # get pgroonga_database source code
pig build dep pgroonga_database; # install build dependencies
pig build pkg pgroonga_database; # build extension rpm or deb
pig build ext pgroonga_database; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgroonga_database; # install by extension name, for the current active PG version
pig ext install pgroonga; # install via package alias, for the active PG version
pig ext install pgroonga_database -v 18;   # install for PG 18
pig ext install pgroonga_database -v 17;   # install for PG 17
pig ext install pgroonga_database -v 16;   # install for PG 16
pig ext install pgroonga_database -v 15;   # install for PG 15
pig ext install pgroonga_database -v 14;   # install for PG 14
pig ext install pgroonga_database -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgroonga_database;
```

