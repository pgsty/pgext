---
title: "pg_statecharts_dev"
linkTitle: "pg_statecharts_dev"
description: "SCXML import and migration tooling for pg_statecharts"
weight: 2696
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_statecharts**](https://github.com/kronor-io/statecharts) : SCXML import and migration tooling for pg_statecharts


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2696** | {{< badge content="pg_statecharts_dev" link="https://github.com/kronor-io/statecharts" >}} | {{< ext "pg_statecharts_dev" "pg_statecharts" >}} | `0.1.0` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fsm` |
|   **Requires**    | {{< ext "pg_statecharts" >}} |
|    **Siblings**   | {{< ext "pg_statecharts" >}} |

> [!Note] Development tooling; server file privileges


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_statecharts` | `pg_statecharts` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_statecharts_dev_18" "green" >}} {{< bg "17" "pg_statecharts_dev_17" "green" >}} {{< bg "16" "pg_statecharts_dev_16" "green" >}} {{< bg "15" "pg_statecharts_dev_15" "red" >}} {{< bg "14" "pg_statecharts_dev_14" "red" >}} | `pg_statecharts_dev_$v` | `pg_statecharts_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-statecharts-dev" "green" >}} {{< bg "17" "postgresql-17-pg-statecharts-dev" "green" >}} {{< bg "16" "postgresql-16-pg-statecharts-dev" "green" >}} {{< bg "15" "postgresql-15-pg-statecharts-dev" "red" >}} {{< bg "14" "postgresql-14-pg-statecharts-dev" "red" >}} | `postgresql-$v-pg-statecharts-dev` | `postgresql-$v-pg-statecharts` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
{.matrix .wide .expand}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kronor-io/statecharts" title="Repository" icon="github" subtitle="github.com/kronor-io/statecharts" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_statecharts-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_statecharts;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_statecharts;		# install via package name, for the active PG version
pig install pg_statecharts_dev;		# install by extension name, for the current active PG version

pig install pg_statecharts_dev -v 18;   # install for PG 18
pig install pg_statecharts_dev -v 17;   # install for PG 17
pig install pg_statecharts_dev -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_statecharts_dev CASCADE; -- requires pg_statecharts
```

## Usage

Sources:

- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/README.md)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/pg_statecharts_dev.control)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)

`pg_statecharts_dev` 0.1.0 is the development companion to `pg_statecharts`. It imports SCXML definitions and generates Sqitch migrations. The production state-machine runtime does not need this file-access tooling.

### Core Workflow

After installing the runtime and its dependency, enable the companion:

```sql
CREATE EXTENSION pg_statecharts_dev;
SELECT * FROM fsm.import_scxml_files(
  source_path => '/repo/statecharts',
  recursive => false,
  on_conflict_do_nothing => false
);
```

The source directory is on the database host. Statechart callbacks must exist with the expected signatures. Parse or validation failures roll back the import transaction instead of accepting part of a chart.

### Migration Generation

```sql
SELECT fsm.gen_statechart_sqitch_migrations(
  source_path => '/repo/statecharts',
  sqitch_plan_file_path => '/repo/sqitch/sqitch.plan',
  recursive => false
);
```

Create the destination directories before generation and make them accessible to the database OS user. The generator writes deploy, revert and verify scripts, updates the plan, and can overwrite existing generated chart files. Review the changes before using them as application migrations.

### Permissions and Boundaries

Server file reads require `pg_read_server_files`; writes require `pg_write_server_files`. These are broad privileges, not a sandbox restricted to the supplied path. Restrict the extension and roles to a controlled development database. It shares the `fsm` schema, is pure SQL and requires no preload; PostgreSQL 16–18 are the documented test targets.

Use the runtime's current semantic-version helpers in generated migrations. For old migrations that cast to the removed semver type, follow the upstream migration guide; regenerating every historical chart can introduce callback-order failures. Deploy reviewed generated SQL to production without installing this companion there.
