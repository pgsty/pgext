---
title: "pg_lion_citext"
linkTitle: "pg_lion_citext"
description: "Case-insensitive citext operator class for lion indexes"
weight: 2716
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_lion**](https://github.com/cgkanchi/pg_lion) : Case-insensitive citext operator class for lion indexes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2716** | {{< badge content="pg_lion_citext" link="https://github.com/cgkanchi/pg_lion" >}} | {{< ext "pg_lion_citext" "pg_lion" >}} | `0.3.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_lion" >}} {{< ext "citext" >}} |
|    **Siblings**   | {{< ext "pg_lion" >}} {{< ext "pg_lion_vector" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lion` | `pg_lion`, `citext` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_lion_18" "green" >}} {{< bg "17" "pg_lion_17" "green" >}} {{< bg "16" "pg_lion_16" "green" >}} {{< bg "15" "pg_lion_15" "red" >}} {{< bg "14" "pg_lion_14" "red" >}} | `pg_lion_$v` | `postgresql$v-contrib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-lion" "green" >}} {{< bg "17" "postgresql-17-pg-lion" "green" >}} {{< bg "16" "postgresql-16-pg-lion" "green" >}} {{< bg "15" "postgresql-15-pg-lion" "red" >}} {{< bg "14" "postgresql-14-pg-lion" "red" >}} | `postgresql-$v-pg-lion` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
{.matrix .wide .expand}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cgkanchi/pg_lion" title="Repository" icon="github" subtitle="github.com/cgkanchi/pg_lion" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lion-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_lion;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_lion;		# install via package name, for the active PG version
pig install pg_lion_citext;		# install by extension name, for the current active PG version

pig install pg_lion_citext -v 18;   # install for PG 18
pig install pg_lion_citext -v 17;   # install for PG 17
pig install pg_lion_citext -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_lion_citext CASCADE; -- requires pg_lion, citext
```

## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext.control)
- [Extension SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext--0.3.0.sql)

`pg_lion_citext` 0.3.0 registers the `citext_ops` operator class for case-insensitive text keys in `lion` indexes. It is a SQL companion shipped with `pg_lion`, with no separate shared library.

### Core Workflow

Create `citext` and `pg_lion` first, then enable the companion. Its default operator class uses `citext` hashing and comparison functions for equality, inequality and ranges (`=`, `<>`, `<`, `<=`, `>=`, `>`).

```sql
CREATE EXTENSION citext;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_citext;
CREATE TABLE lion_contacts (email citext);
INSERT INTO lion_contacts VALUES ('Ada@example.com'), ('ADA@example.com');
CREATE INDEX lion_email ON lion_contacts USING lion (email);
SELECT count(*) FROM lion_contacts WHERE email = 'ada@example.com';
```

The example returns two rows in the count because both spellings compare equal. Applications keep using ordinary `citext` predicates.

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Case folding and ordering follow `citext` semantics. Use matching files from the parent distribution. The parent indexes retain their visibility, `VACUUM`, and optional custom-WAL preload requirements. This companion does not enable preloading or change existing indexes by itself.
