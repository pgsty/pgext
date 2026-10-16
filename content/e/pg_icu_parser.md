---
title: "pg_icu_parser"
linkTitle: "pg_icu_parser"
description: "Full text search parser using ICU boundary analysis"
weight: 2300
categories: ["FTS"]
languages: ["C"]
licenses: ["MPL-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_icu_parser**](https://github.com/IgKh/pg_icu_parser) : Full text search parser using ICU boundary analysis


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2300** | {{< badge content="pg_icu_parser" link="https://github.com/IgKh/pg_icu_parser" >}} | {{< ext "pg_icu_parser" >}} | `1.0` | {{< category "FTS" >}} | {{< license "MPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jieba" >}} {{< ext "pg_cjk_parser" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "zhparser" >}} `thai_parser` {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} |

> [!Note] SQL/control version 1.0; upstream release 1.0.0. Local RPM/DEB builds support PG14-18 with a PG18 encoding-API compatibility patch and explicit ICU linkage. ts_headline is unsupported. PIGSTY PG14-18 packages verified across EL8/9/10, Debian 12/13 and Ubuntu 22.04/24.04/26.04 on x86_64/amd64 and aarch64/arm64; all 16 platforms imported into the local PIGSTY repository.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_icu_parser` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_icu_parser_18" "green" >}} {{< bg "17" "pg_icu_parser_17" "green" >}} {{< bg "16" "pg_icu_parser_16" "green" >}} {{< bg "15" "pg_icu_parser_15" "green" >}} {{< bg "14" "pg_icu_parser_14" "green" >}} | `pg_icu_parser_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-icu-parser" "green" >}} {{< bg "17" "postgresql-17-pg-icu-parser" "green" >}} {{< bg "16" "postgresql-16-pg-icu-parser" "green" >}} {{< bg "15" "postgresql-15-pg-icu-parser" "green" >}} {{< bg "14" "postgresql-14-pg-icu-parser" "green" >}} | `postgresql-$v-pg-icu-parser` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "pg_icu_parser_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_17 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_16 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_15 : MISS 0" "red" >}} | {{< bg "MISS" "pg_icu_parser_14 : MISS 0" "red" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-18-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-pg-icu-parser : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-pg-icu-parser : MISS 0" "red" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/IgKh/pg_icu_parser" title="Repository" icon="github" subtitle="github.com/IgKh/pg_icu_parser" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_icu_parser-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_icu_parser;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_icu_parser;		# install via package name, for the active PG version

pig install pg_icu_parser -v 18;   # install for PG 18
pig install pg_icu_parser -v 17;   # install for PG 17
pig install pg_icu_parser -v 16;   # install for PG 16
pig install pg_icu_parser -v 15;   # install for PG 15
pig install pg_icu_parser -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_icu_parser;
```

## Usage

Sources:

- [Official PGXN documentation](https://pgxn.org/dist/pg_icu_parser/1.0.0/)

`pg_icu_parser` supplies the `icu_parser` full-text-search parser, which uses ICU Unicode word-boundary analysis. It is useful for languages such as Chinese, Japanese, Korean, and Hebrew where PostgreSQL's whitespace- and pattern-oriented default parser may not produce suitable word boundaries.

### Core Workflow

The server must be PostgreSQL 10 or later and must have been built with ICU support. Create the extension, then create a text-search configuration that selects its parser and maps the token types you need:

```sql
CREATE EXTENSION pg_icu_parser;

CREATE TEXT SEARCH CONFIGURATION sample (PARSER = icu_parser);
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR word WITH english_stem;
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR number, ideo, kana WITH simple;

SELECT to_tsvector('sample', '東京でPostgreSQLを使う');
```

Choose dictionaries for the actual language; `english_stem` above only demonstrates the mapping syntax.

### Parser Surface

- `word`, `number`, `kana`, `ideo`, and `blank` are the token types emitted by `icu_parser`.
- `pg_icu_parser.locale` selects the ICU locale and defaults to `en`. Upstream notes that word-boundary rules are usually not locale-sensitive, so change it only after testing a concrete need.
- Any PostgreSQL text-search dictionary can be mapped to the emitted token types.

### Trade-offs

- The parser requires ICU in the PostgreSQL build; creating the extension cannot add ICU to a server built without it.
- It emits fewer and coarser token types than the default parser and does not recognize special patterns such as URLs or email addresses as distinct categories.
- The reviewed release does not support `ts_headline`.
- Upstream reports that it is slower than the default parser, with a lower conversion overhead when the database encoding is UTF-8. Benchmark representative documents and query shapes on the target system.
- When parser mappings or locale behavior change, rebuild stored `tsvector` values and dependent indexes so old and new rows do not use different tokenization.
