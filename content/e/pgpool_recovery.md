---
title: "pgpool_recovery"
linkTitle: "pgpool_recovery"
description: "recovery functions for pgpool-II for V4.3"
weight: 5910
categories: ["ADMIN"]
width: full
---

recovery functions for pgpool-II for V4.3


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5910** | {{< badge content="pgpool_recovery" link="https://pgpool.net/" >}} | {{< ext "pgpool_recovery" "pgpool" >}} | `4.6.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgautofailover" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "repmgr" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |
|    **Siblings**   | {{< ext "pgpool_adm" >}} {{< ext "pgpool_regclass" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgpool_adm" >}} | `4.6.3` | {{< bg "18" "pgpool-II-pg18-extensions" "green" >}} {{< bg "17" "pgpool-II-pg17-extensions" "green" >}} {{< bg "16" "pgpool-II-pg16-extensions" "green" >}} {{< bg "15" "pgpool-II-pg15-extensions" "green" >}} {{< bg "14" "pgpool-II-pg14-extensions" "green" >}} {{< bg "13" "pgpool-II-pg13-extensions" "green" >}} | `pgpool-II-pg$v-extensions` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgpool_adm" >}} | `4.6.3` | {{< bg "18" "postgresql-18-pgpool2" "green" >}} {{< bg "17" "postgresql-17-pgpool2" "green" >}} {{< bg "16" "postgresql-16-pgpool2" "green" >}} {{< bg "15" "postgresql-15-pgpool2" "green" >}} {{< bg "14" "postgresql-14-pgpool2" "green" >}} {{< bg "13" "postgresql-13-pgpool2" "green" >}} | `postgresql-$v-pgpool2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 15" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 15" >}}   |
|    `el8.aarch64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 11" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 11" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 11" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 14" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 14" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 6" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 9" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 12" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 12" >}}   |
|    `el10.x86_64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 4" >}}   |
|    `el10.aarch64`    |  {{< bg "PGDG 4.6.3" "pgpool-II-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg17-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg16-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg15-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg14-extensions : HIDE 4" >}}   |  {{< bg "PGDG 4.6.3" "pgpool-II-pg13-extensions : HIDE 4" >}}   |
|    `d12.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `d13.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `d13.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "PGDG 4.6.3" "postgresql-18-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-17-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-16-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-15-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-14-pgpool2 : HIDE 1" >}}   |  {{< bg "PGDG 4.6.3" "postgresql-13-pgpool2 : HIDE 1" >}}   |


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
pig ext install pgpool_recovery; # install by extension name, for the current active PG version
pig ext install pgpool; # install via package alias, for the active PG version
pig ext install pgpool_recovery -v 18;   # install for PG 18
pig ext install pgpool_recovery -v 17;   # install for PG 17
pig ext install pgpool_recovery -v 16;   # install for PG 16
pig ext install pgpool_recovery -v 15;   # install for PG 15
pig ext install pgpool_recovery -v 14;   # install for PG 14
pig ext install pgpool_recovery -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgpool_recovery;
```

