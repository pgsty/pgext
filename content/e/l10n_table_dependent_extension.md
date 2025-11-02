---
title: "l10n_table_dependent_extension"
linkTitle: "l10n_table_dependent_extension"
description: "PostgreSQL l10n toolbox"
weight: 3611
categories: ["TYPE"]
width: full
---

PostgreSQL l10n toolbox


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3611** | {{< badge content="l10n_table_dependent_extension" link="https://github.com/bigsmoke/pg_xenophile" >}} | {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | `0.8.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_xenophile" >}} |
|   **See Also**    | {{< ext "country" >}} {{< ext "currency" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |
|    **Siblings**   | {{< ext "pg_xenophile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_xenophile" >}} | `0.8.3` | {{< bg "18" "pg_xenophile_18" "green" >}} {{< bg "17" "pg_xenophile_17" "green" >}} {{< bg "16" "pg_xenophile_16" "green" >}} {{< bg "15" "pg_xenophile_15" "green" >}} {{< bg "14" "pg_xenophile_14" "green" >}} {{< bg "13" "pg_xenophile_13" "green" >}} | `pg_xenophile_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_xenophile" >}} | `0.8.3` | {{< bg "18" "postgresql-18-pg-xenophile" "green" >}} {{< bg "17" "postgresql-17-pg-xenophile" "green" >}} {{< bg "16" "postgresql-16-pg-xenophile" "green" >}} {{< bg "15" "postgresql-15-pg-xenophile" "green" >}} {{< bg "14" "postgresql-14-pg-xenophile" "green" >}} {{< bg "13" "postgresql-13-pg-xenophile" "green" >}} | `postgresql-$v-pg-xenophile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-xenophile : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-xenophile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-xenophile : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-xenophile : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_xenophile" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_xenophile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_xenophile-0.8.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get l10n_table_dependent_extension; # get l10n_table_dependent_extension source code
pig build dep l10n_table_dependent_extension; # install build dependencies
pig build pkg l10n_table_dependent_extension; # build extension rpm or deb
pig build ext l10n_table_dependent_extension; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install l10n_table_dependent_extension; # install by extension name, for the current active PG version
pig ext install pg_xenophile; # install via package alias, for the active PG version
pig ext install l10n_table_dependent_extension -v 18;   # install for PG 18
pig ext install l10n_table_dependent_extension -v 17;   # install for PG 17
pig ext install l10n_table_dependent_extension -v 16;   # install for PG 16
pig ext install l10n_table_dependent_extension -v 15;   # install for PG 15
pig ext install l10n_table_dependent_extension -v 14;   # install for PG 14
pig ext install l10n_table_dependent_extension -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION l10n_table_dependent_extension;
```

