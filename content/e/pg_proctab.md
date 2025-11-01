---
title: "pg_proctab"
linkTitle: "pg_proctab"
description: "PostgreSQL extension to access the OS process table"
weight: 6320
categories: ["STAT"]
width: full
---

PostgreSQL extension to access the OS process table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6320** | {{< badge content="pg_proctab" link="https://github.com/markwkm/pg_proctab" >}} | {{< ext "pg_proctab" "pgnodemx" >}} | `1.7` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prioritize" >}} {{< ext "system_stats" >}} {{< ext "pg_background" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgmeminfo" >}} {{< ext "pgsentinel" >}} {{< ext "pg_profile" >}} |
|    **Siblings**   | {{< ext "pgnodemx" >}} |

> [!Note] from pgnodemx


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgnodemx" >}} | `1.7` | {{< bg "18" "pgnodemx_18" "green" >}} {{< bg "17" "pgnodemx_17" "green" >}} {{< bg "16" "pgnodemx_16" "green" >}} {{< bg "15" "pgnodemx_15" "green" >}} {{< bg "14" "pgnodemx_14" "green" >}} {{< bg "13" "pgnodemx_13" "green" >}} | `pgnodemx_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgnodemx" >}} | `1.7` | {{< bg "18" "postgresql-18-pgnodemx" "green" >}} {{< bg "17" "postgresql-17-pgnodemx" "green" >}} {{< bg "16" "postgresql-16-pgnodemx" "green" >}} {{< bg "15" "postgresql-15-pgnodemx" "green" >}} {{< bg "14" "postgresql-14-pgnodemx" "green" >}} {{< bg "13" "postgresql-13-pgnodemx" "green" >}} | `postgresql-$v-pgnodemx` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/markwkm/pg_proctab" title="Repository" icon="github" subtitle="github.com/markwkm/pg_proctab" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgnodemx-1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_proctab; # get pg_proctab source code
pig build dep pg_proctab; # install build dependencies
pig build pkg pg_proctab; # build extension rpm or deb
pig build ext pg_proctab; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_proctab; # install by extension name, for the current active PG version
pig ext install pgnodemx; # install via package alias, for the active PG version
pig ext install pg_proctab -v 18;   # install for PG 18
pig ext install pg_proctab -v 17;   # install for PG 17
pig ext install pg_proctab -v 16;   # install for PG 16
pig ext install pg_proctab -v 15;   # install for PG 15
pig ext install pg_proctab -v 14;   # install for PG 14
pig ext install pg_proctab -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_proctab;
```

