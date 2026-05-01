---
title: "pg_readme_test_extension"
linkTitle: "pg_readme_test_extension"
description: "Test generating a README.md document for extension or schema"
weight: 4301
categories: ["UTIL"]
width: full
---

[**pg_readme**](https://github.com/bigsmoke/pg_readme) : Test generating a README.md document for extension or schema


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4301** | {{< badge content="pg_readme_test_extension" link="https://github.com/bigsmoke/pg_readme" >}} | {{< ext "pg_readme_test_extension" "pg_readme" >}} | `0.7.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} |
|   **See Also**    | {{< ext "schedoc" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} |
|    **Siblings**   | {{< ext "pg_readme" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_readme` | `hstore` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.7.0` | {{< bg "18" "pg_readme_18" "green" >}} {{< bg "17" "pg_readme_17" "green" >}} {{< bg "16" "pg_readme_16" "green" >}} {{< bg "15" "pg_readme_15" "green" >}} {{< bg "14" "pg_readme_14" "green" >}} | `pg_readme_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.0` | {{< bg "18" "postgresql-18-pg-readme" "green" >}} {{< bg "17" "postgresql-17-pg-readme" "green" >}} {{< bg "16" "postgresql-16-pg-readme" "green" >}} {{< bg "15" "postgresql-15-pg-readme" "green" >}} {{< bg "14" "postgresql-14-pg-readme" "green" >}} | `postgresql-$v-pg-readme` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_readme" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_readme" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_readme-0.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_readme;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_readme;		# install via package name, for the active PG version
pig install pg_readme_test_extension;		# install by extension name, for the current active PG version

pig install pg_readme_test_extension -v 18;   # install for PG 18
pig install pg_readme_test_extension -v 17;   # install for PG 17
pig install pg_readme_test_extension -v 16;   # install for PG 16
pig install pg_readme_test_extension -v 15;   # install for PG 15
pig install pg_readme_test_extension -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_readme_test_extension CASCADE; -- requires hstore
```




## Usage

> [pg_readme_test_extension: Test extension for pg_readme](https://github.com/bigsmoke/pg_readme)

This is a companion test extension for the `pg_readme` extension. It exists solely to test `pg_readme`'s functionality in generating documentation from PostgreSQL `COMMENT` objects.

It is not intended for direct use in production. It serves as a reference implementation demonstrating how to structure extension comments and processing instructions for `pg_readme` to generate proper README documentation.

### Related

See the [`pg_readme`](https://github.com/bigsmoke/pg_readme) extension for the actual documentation generation functionality.
