---
title: "pgpool_regclass"
linkTitle: "pgpool_regclass"
description: "replacement for regclass"
weight: 5920
categories: ["ADMIN"]
width: full
---

replacement for regclass


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5920** | {{< badge content="pgpool_regclass" link="https://pgpool.net/" >}} | {{< ext "pgpool_regclass" "pgpool" >}} | `4.6.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} |
|    **Siblings**   | {{< ext "pgpool_adm" >}} {{< ext "pgpool_recovery" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgpool_adm" >}} | `4.6.3` | {{< bg "18" "pgpool-II-pg18-extensions" "red" >}} {{< bg "17" "pgpool-II-pg17-extensions" "green" >}} {{< bg "16" "pgpool-II-pg16-extensions" "green" >}} {{< bg "15" "pgpool-II-pg15-extensions" "green" >}} {{< bg "14" "pgpool-II-pg14-extensions" "green" >}} | `pgpool-II-pg$v-extensions` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgpool_adm" >}} | `4.6.3` | {{< bg "18" "postgresql-18-pgpool2" "green" >}} {{< bg "17" "postgresql-17-pgpool2" "green" >}} {{< bg "16" "postgresql-16-pgpool2" "green" >}} {{< bg "15" "postgresql-15-pgpool2" "green" >}} {{< bg "14" "postgresql-14-pgpool2" "green" >}} | `postgresql-$v-pgpool2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 15" >}}   |
|    `el8.aarch64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 11" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 11" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 14" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 12" >}}   |
|    `d12.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |


## Source

{{< cards cols=3 >}}
{{< card link="https://pgpool.net/" title="Repository" icon="link" subtitle="pgpool.net/" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgpool_regclass; # install by extension name, for the current active PG version
pig ext install pgpool; # install via package alias, for the active PG version
pig ext install pgpool_regclass -v 18;   # install for PG 18
pig ext install pgpool_regclass -v 17;   # install for PG 17
pig ext install pgpool_regclass -v 16;   # install for PG 16
pig ext install pgpool_regclass -v 15;   # install for PG 15
pig ext install pgpool_regclass -v 14;   # install for PG 14
pig ext install pgpool_regclass -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgpool_regclass;
```

