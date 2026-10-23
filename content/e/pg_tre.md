---
title: "pg_tre"
linkTitle: "pg_tre"
description: "Native PostgreSQL index access method for approximate regular-expression matching"
weight: 2330
categories: ["FTS"]
languages: ["C"]
licenses: ["MIT AND BSD-2-Clause"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_tre**](https://codeberg.org/gregburd/pg_tre) : Native PostgreSQL index access method for approximate regular-expression matching


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2330** | {{< badge content="pg_tre" link="https://codeberg.org/gregburd/pg_tre" >}} | {{< ext "pg_tre" >}} | `4.5.1` | {{< category "FTS" >}} | {{< license "MIT AND BSD-2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_fts" >}} {{< ext "vector" >}} `pg_infer` {{< ext "pg_mentat" >}} |

> [!Note] 4.5.1 hardens index_check search_path. PG17-19 upstream; RPM targets PG17-18. Generic WAL since 4.0 needs no preload. The 3.x to 4.x WAL transition requires clean shutdown and physical standby rebuild. REINDEX indexes built with min_trigram_freq>1; fastupdate=false is refused; pending_list_limit defaults to 0. Formats 6-10 remain readable; 4.4+ writes format 11.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_tre` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.1` | {{< bg "18" "pg_tre_18" "green" >}} {{< bg "17" "pg_tre_17" "green" >}} {{< bg "16" "pg_tre_16" "red" >}} {{< bg "15" "pg_tre_15" "red" >}} {{< bg "14" "pg_tre_14" "red" >}} | `pg_tre_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.1" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tre_18` | `4.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 590.3 KiB | [pg_tre_18-4.5.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tre_18-4.5.1-1PGSTY.el8.x86_64.rpm) |
| `pg_tre_18` | `4.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 574.0 KiB | [pg_tre_18-4.5.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tre_18-4.5.1-1PGSTY.el8.aarch64.rpm) |
| `pg_tre_18` | `4.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 560.7 KiB | [pg_tre_18-4.5.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tre_18-4.5.1-1PGSTY.el9.x86_64.rpm) |
| `pg_tre_18` | `4.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 551.8 KiB | [pg_tre_18-4.5.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tre_18-4.5.1-1PGSTY.el9.aarch64.rpm) |
| `pg_tre_18` | `4.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 567.7 KiB | [pg_tre_18-4.5.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tre_18-4.5.1-1PGSTY.el10.x86_64.rpm) |
| `pg_tre_18` | `4.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 556.2 KiB | [pg_tre_18-4.5.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tre_18-4.5.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tre_17` | `4.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 590.9 KiB | [pg_tre_17-4.5.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tre_17-4.5.1-1PGSTY.el8.x86_64.rpm) |
| `pg_tre_17` | `4.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 574.9 KiB | [pg_tre_17-4.5.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tre_17-4.5.1-1PGSTY.el8.aarch64.rpm) |
| `pg_tre_17` | `4.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 561.5 KiB | [pg_tre_17-4.5.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tre_17-4.5.1-1PGSTY.el9.x86_64.rpm) |
| `pg_tre_17` | `4.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 552.4 KiB | [pg_tre_17-4.5.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tre_17-4.5.1-1PGSTY.el9.aarch64.rpm) |
| `pg_tre_17` | `4.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 568.4 KiB | [pg_tre_17-4.5.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tre_17-4.5.1-1PGSTY.el10.x86_64.rpm) |
| `pg_tre_17` | `4.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 557.0 KiB | [pg_tre_17-4.5.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tre_17-4.5.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_tre" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_tre" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tre-4.5.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_tre;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_tre;		# install via package name, for the active PG version

pig install pg_tre -v 18;   # install for PG 18
pig install pg_tre -v 17;   # install for PG 17

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_tre;
```

## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/README.md)
- [pg_tre.control](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/pg_tre.control)
- [sql/pg_tre--4.5.1.sql](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/sql/pg_tre--4.5.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/CHANGELOG.md)
- [LIMITATIONS.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/LIMITATIONS.md)
- [META.json](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/META.json)
- [.github/workflows/ci.yml](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.1/.github/workflows/ci.yml)
- [Library initialization](https://codeberg.org/gregburd/pg_tre/src/tag/v4.5.1/src/module.c)

`pg_tre` 4.5.1 provides the `tre` native index access method for approximate regular-expression matching over text on PostgreSQL 17–19. It narrows candidates with codepoint trigrams and postings, then checks matches against the heap. Use it for typo-tolerant identifiers, log messages and lexical patterns with explicit edit costs.

### Enable and Search

The index write path uses PostgreSQL's generic WAL facility. Since 4.0.0, `CREATE EXTENSION` is sufficient: `shared_preload_libraries` and a restart are not required to enable a new installation. Extension creation requires a superuser, and the control file is untrusted and non-relocatable.

```sql
CREATE EXTENSION pg_tre;
CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body text NOT NULL
);
INSERT INTO documents (body) VALUES
  ('The PostgreSQL database system'),
  ('The Postgres databse system');
CREATE INDEX documents_body_tre ON documents USING tre (body);
SELECT id, body, body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

PostgreSQL 19 uses the core regular-expression character classifier; ICU character-class edge cases on PostgreSQL 17–18 may differ from native PostgreSQL regex semantics, as described upstream.

### Patterns and Objects

- `tre_pattern(text)` uses `pg_tre.default_max_cost`; `tre_pattern(text, int4)` supplies a maximum cost; the five-argument form also sets insertion, deletion and substitution costs. Patterns use POSIX extended regular expressions and TRE approximate budgets.
- `text %~~ tre_pattern` is the indexable approximate-match predicate; `text <@> tre_pattern` returns the edit cost or NULL when no match fits the budget.
- `tre_distance`, `tre_similarity`, `tre_amatch`, `tre_amatch_cost` and `tre_amatch_detail` expose matching and compatibility functions.
- The index operator family also supports LIKE, ILIKE, regular-expression and equality predicates with heap rechecks. The separately named trigram functions coexist with `pg_trgm`; `pg_tre` does not install the colliding bare trigram operators removed in 3.0.

Indexes accept one text column with `tre_text_ops`. Expression and partial indexes are available, but index-only scans are not. Short patterns, high edit budgets and non-literal regexes can produce broad candidate sets; inspect real plans and keep cancellation and query-time limits enabled.

### Maintenance and Limits

`fastupdate` must be true; setting it false is now rejected because the direct path could not preserve the index format. Existing off settings can still be opened, but inserts use the pending list. `pending_list_limit` is in KiB and defaults to 0, meaning unlimited. A positive value triggers a synchronous merge when the threshold is crossed; that inserting backend pays the work, and if another merge is active the limit can temporarily be exceeded.

`pg_tre.min_trigram_freq` must remain 1. Values above 1 formerly omitted rare postings and could lose results; they are now rejected. Keep regex state and compilation/matching limits bounded for untrusted patterns. Index construction uses tuplesort and temporary disk, with memory governed by maintenance_work_mem; estimate large builds before scheduling them.

```sql
SELECT * FROM tre_estimate_index_build('documents'::regclass, 2);
SELECT * FROM pg_tre_index_pending('documents_body_tre'::regclass);
SELECT * FROM pg_tre_index_check('documents_body_tre'::regclass, 2);
```

`pg_tre_index_check` compares index answers with sequential scans. Each sample performs multiple sequential scans, so bound the sample count on large tables. Version 4.5.1 fixes its search-path handling, pins trusted lookup paths and qualifies calls; use the fixed version before running this helper as a privileged user.

### Upgrade Boundaries

Install matching files, ensure sessions use the new library, and update each database:

```sql
ALTER EXTENSION pg_tre UPDATE TO '4.5.1';
```

Crossing from 3.x to 4.x changes WAL record types even though the index page format is unchanged. Shut down cleanly before replacing the library; a physical standby cannot run a different major extension version from its primary and must be rebuilt from a fresh base backup after the primary upgrade. There is no rolling upgrade across this boundary. Remove any old `wal_consistency_checking = 'pg_tre'` setting; it prevents startup in 4.x. Upstream recommends `wal_consistency_checking = 'all'` when consistency checking is needed.

Version 4.4 writes format 11 and reads formats 6–10. Upgrading from those formats does not require a routine rebuild; optional REINDEX obtains the newer dense-posting representation. Downgrading to 4.3 after a 4.4+ write requires rebuilding the affected indexes for the older format. Older upgrades may cross additional boundaries described in the changelog.

REINDEX every index built with min_trigram_freq greater than 1: installing the fix cannot restore omitted postings. Version 4.5 also disables stale SuRF acceleration on indexes from earlier affected versions, restoring correctness immediately; a rebuild is optional to regain that acceleration. Upgrading 4.5.0 to 4.5.1 changes neither format nor ordinary query behavior and requires no additional REINDEX. Old definitions using the removed tuple bloom options must still be corrected before rebuilding.
