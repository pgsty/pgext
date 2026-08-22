---
title: "pg_readme_test_extension"
linkTitle: "pg_readme_test_extension"
description: "Fixture extension used to test pg_readme document generation"
weight: 4301
categories: ["UTIL"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_readme**](https://github.com/bigsmoke/pg_readme/tree/master/pg_readme_test_extension) : Fixture extension used to test pg_readme document generation


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4301** | {{< badge content="pg_readme_test_extension" link="https://github.com/bigsmoke/pg_readme/tree/master/pg_readme_test_extension" >}} | {{< ext "pg_readme_test_extension" "pg_readme" >}} | `0.7.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readme" >}} {{< ext "schedoc" >}} {{< ext "ddlx" >}} {{< ext "pgpdf" >}} {{< ext "pg_render" >}} {{< ext "pgdd" >}} {{< ext "meta" >}} |
|    **Siblings**   | {{< ext "pg_readme" >}} |

> [!Note] Bundled test fixture; its control default_version is forever and it does not require hstore; package ownership follows pg_readme: PGDG RPM 0.7.0 and PIGSTY DEB 0.7.1.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_readme` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.7.0` | {{< bg "18" "pg_readme_18" "green" >}} {{< bg "17" "pg_readme_17" "green" >}} {{< bg "16" "pg_readme_16" "green" >}} {{< bg "15" "pg_readme_15" "green" >}} {{< bg "14" "pg_readme_14" "green" >}} | `pg_readme_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.1` | {{< bg "18" "postgresql-18-pg-readme" "green" >}} {{< bg "17" "postgresql-17-pg-readme" "green" >}} {{< bg "16" "postgresql-16-pg-readme" "green" >}} {{< bg "15" "postgresql-15-pg-readme" "green" >}} {{< bg "14" "postgresql-14-pg-readme" "green" >}} | `postgresql-$v-pg-readme` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_readme/tree/master/pg_readme_test_extension" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_readme/tree/master/pg_readme_test_extension" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_readme-0.7.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_readme;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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
CREATE EXTENSION pg_readme_test_extension;
```

## Usage

Sources:

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [Test-extension control file](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension.control)
- [Test-extension SQL fixture](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension--forever.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_readme_test_extension)

`pg_readme_test_extension` is the bundled integration-test fixture for `pg_readme`. It installs commented domains, types, tables, views, routines, triggers, and processing instructions so upstream can verify `pg_extension_readme()`. It is not an application feature or a production dependency.

### Exercise the Fixture

```sql
CREATE EXTENSION pg_readme CASCADE;
CREATE EXTENSION pg_readme_test_extension;

SELECT pg_extension_readme('pg_readme_test_extension'::name);
```

Use the output to test or demonstrate the generator, then remove the fixture from a disposable database:

```sql
DROP EXTENSION pg_readme_test_extension;
```

### Boundary and Caveats

- The upstream distribution version is 0.7.1, but this fixture's control version is deliberately the literal `forever`.
- The fixture ships with `pg_readme`; current Pigsty DEB packages are 0.7.1 and RPM packages remain 0.7.0. The extension version stays `forever` on both sides.
- It is relocatable, does not require `hstore` itself, and creates sample objects with generic names. Install it only where those objects cannot collide with real application schemas.
- Its SQL surface exists to cover generator behavior and can change as tests evolve. Do not build application code against the fixture objects.
