---
title: "pgroonga_database"
linkTitle: "pgroonga_database"
description: "PGroonga database management module"
weight: 2111
categories: ["FTS"]
width: full
---

[**pgroonga**](https://github.com/pgroonga/pgroonga) : PGroonga database management module


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2111** | {{< badge content="pgroonga_database" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga_database" "pgroonga" >}} | `4.0.4` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} {{< ext "unaccent" >}} |
|    **Siblings**   | {{< ext "pgroonga" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgroonga` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "pgroonga_18" "green" >}} {{< bg "17" "pgroonga_17" "green" >}} {{< bg "16" "pgroonga_16" "green" >}} {{< bg "15" "pgroonga_15" "green" >}} {{< bg "14" "pgroonga_14" "green" >}} | `pgroonga_$v` | `groonga-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgroonga;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgroonga;		# install via package name, for the active PG version
pig install pgroonga_database;		# install by extension name, for the current active PG version

pig install pgroonga_database -v 18;   # install for PG 18
pig install pgroonga_database -v 17;   # install for PG 17
pig install pgroonga_database -v 16;   # install for PG 16
pig install pgroonga_database -v 15;   # install for PG 15
pig install pgroonga_database -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgroonga_database;
```



## Usage

> [PGroonga Documentation](https://pgroonga.github.io/) | [GitHub: pgroonga/pgroonga](https://github.com/pgroonga/pgroonga)

`pgroonga_database` is a sub-extension of the [PGroonga](https://pgroonga.github.io/) project. It provides database management functionality for PGroonga, which makes PostgreSQL a fast full text search platform for all languages.

PGroonga is a comprehensive full text search solution that uses [Groonga](https://groonga.org/) as a backend. It supports all languages including CJK (Chinese, Japanese, Korean) out of the box, and provides rich features such as:

- Fast full text search with all language support
- Rich query syntax (query language, script syntax)
- JSON search
- HTML/XML tag aware highlighting
- Similar search
- Synonym expansion
- Autocomplete
- Query log analysis

The PGroonga documentation is extensive and spans hundreds of pages. For detailed usage, API reference, operators, functions, and tuning guides, please refer to the official documentation site:

- [PGroonga Official Documentation](https://pgroonga.github.io/)
- [Getting Started](https://pgroonga.github.io/install/)
- [Tutorial](https://pgroonga.github.io/tutorial/)
- [How-to Guides](https://pgroonga.github.io/how-to/)
- [Reference](https://pgroonga.github.io/reference/)

## Quick Start

```sql
CREATE EXTENSION pgroonga_database;
CREATE EXTENSION pgroonga;

-- Create a table with text content
CREATE TABLE memos (
  id integer,
  content text
);

-- Create a PGroonga index
CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

-- Insert data
INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as its backend.');

-- Full text search
SELECT * FROM memos WHERE content &@~ 'PostgreSQL OR Groonga';
```
