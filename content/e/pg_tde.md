---
title: "pg_tde"
linkTitle: "pg_tde"
description: "Percona pg_tde access method"
weight: 7060
categories: ["SEC"]
width: full
---

Percona pg_tde access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7060** | {{< badge content="pg_tde" link="https://github.com/Percona-Lab/pg_tde" >}} | {{< ext "pg_tde" >}} | `1.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgcrypto" >}} {{< ext "anon" >}} {{< ext "pgcryptokey" >}} {{< ext "faker" >}} {{< ext "sslutils" >}} {{< ext "uuid-ossp" >}} |

> [!Note] works on percona postgres fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_tde" >}} | `1.0` | {{< bg "18" "percona-postgresql18" "red" >}} {{< bg "17" "percona-postgresql17" "green" >}} {{< bg "16" "percona-postgresql16" "red" >}} {{< bg "15" "percona-postgresql15" "red" >}} {{< bg "14" "percona-postgresql14" "red" >}} {{< bg "13" "percona-postgresql13" "red" >}} | `percona-postgresql$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tde" >}} | `1.0` | {{< bg "18" "percona-postgresql-18" "red" >}} {{< bg "17" "percona-postgresql-17" "green" >}} {{< bg "16" "percona-postgresql-16" "red" >}} {{< bg "15" "percona-postgresql-15" "red" >}} {{< bg "14" "percona-postgresql-14" "red" >}} {{< bg "13" "percona-postgresql-13" "red" >}} | `percona-postgresql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "purple" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "purple" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Percona-Lab/pg_tde" title="Repository" icon="github" subtitle="github.com/Percona-Lab/pg_tde" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tde-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_tde; # get pg_tde source code
pig build dep pg_tde; # install build dependencies
pig build pkg pg_tde; # build extension rpm or deb
pig build ext pg_tde; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_tde; # install by extension name, for the current active PG version
pig ext install pg_tde; # install via package alias, for the active PG version
pig ext install pg_tde -v 17;   # install for PG 17

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_tde;
```

