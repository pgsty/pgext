---
title: "pg_tde"
linkTitle: "pg_tde"
description: "Percona pg_tde access method"
weight: 7060
categories: ["Sec"]
width: full
---

Percona pg_tde access method

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7060** | {{< badge content="pg_tde" link="https://github.com/Percona-Lab/pg_tde" >}} | {{< ext "pg_tde" "pg_tde" >}} | `1.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgcrypto" >}} {{< ext "anon" >}} {{< ext "pgcryptokey" >}} {{< ext "faker" >}} {{< ext "sslutils" >}} {{< ext "uuid-ossp" >}} |

> [!Note] works on percona postgres fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_tde" >}} | `1.0` | {{< badge content="18" color="red" alt="percona-postgresql18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="percona-postgresql16" >}} {{< badge content="15" color="red" alt="percona-postgresql15" >}} {{< badge content="14" color="red" alt="percona-postgresql14" >}} | `percona-postgresql$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tde" >}} | `1.0` | {{< badge content="18" color="red" alt="percona-postgresql-18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="percona-postgresql-16" >}} {{< badge content="15" color="red" alt="percona-postgresql-15" >}} {{< badge content="14" color="red" alt="percona-postgresql-14" >}} | `percona-postgresql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "percona-postgresql18" >}}     |    {{< pkg "percona-postgresql17" >}}     |    {{< pkg "percona-postgresql16" >}}     |    {{< pkg "percona-postgresql15" >}}     |    {{< pkg "percona-postgresql14" >}}     |
|    `el8.aarch64`    |    {{< pkg "percona-postgresql18" >}}     |    {{< pkg "percona-postgresql17" >}}     |    {{< pkg "percona-postgresql16" >}}     |    {{< pkg "percona-postgresql15" >}}     |    {{< pkg "percona-postgresql14" >}}     |
|    `el9.x86_64`    |    {{< pkg "percona-postgresql18" >}}     |    {{< pkg "percona-postgresql17" >}}     |    {{< pkg "percona-postgresql16" >}}     |    {{< pkg "percona-postgresql15" >}}     |    {{< pkg "percona-postgresql14" >}}     |
|    `el9.aarch64`    |    {{< pkg "percona-postgresql18" >}}     |    {{< pkg "percona-postgresql17" >}}     |    {{< pkg "percona-postgresql16" >}}     |    {{< pkg "percona-postgresql15" >}}     |    {{< pkg "percona-postgresql14" >}}     |
|    `d12.x86_64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |
|    `d12.aarch64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |
|    `u22.x86_64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |
|    `u22.aarch64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |
|    `u24.x86_64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |
|    `u24.aarch64`    |    {{< pkg "percona-postgresql-18" >}}     |    {{< pkg "percona-postgresql-17" >}}     |    {{< pkg "percona-postgresql-16" >}}     |    {{< pkg "percona-postgresql-15" >}}     |    {{< pkg "percona-postgresql-14" >}}     |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Percona-Lab/pg_tde" title="Repository" icon="github" subtitle="github.com/Percona-Lab/pg_tde" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_tde-1.0.0.tar.gz" >}}
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

