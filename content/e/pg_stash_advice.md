---
title: "pg_stash_advice"
linkTitle: "pg_stash_advice"
description: "Store and automatically apply plan advice by query ID"
weight: 2989
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["CONTRIB"]
page_width: full
---

[**pg_stash_advice**](https://www.postgresql.org/docs/19/pgstashadvice.html) : Store and automatically apply plan advice by query ID


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2989** | {{< badge content="pg_stash_advice" link="https://www.postgresql.org/docs/19/pgstashadvice.html" >}} | {{< ext "pg_stash_advice" >}} | `1.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_plan_advice" >}} |
|   **See Also**    | {{< ext "pg_plan_advice" >}} {{< ext "pg_hint_plan" >}} {{< ext "plan_filter" >}} {{< ext "pg_stat_plans" >}} |


## Packages

| **PG19** | **PG18** | **PG17** | **PG16** | **PG15** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 19: version 1.0" "green" >}} | {{< bg "N/A" "PostgreSQL 18: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 16: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 15: not available" "red" >}} |
{.wide .expand}

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_plan_advice, pg_stash_advice';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_stash_advice CASCADE; -- requires pg_plan_advice
```

## Usage

Sources:

- [PostgreSQL 19 documentation](https://www.postgresql.org/docs/19/pgstashadvice.html)
- [PostgreSQL 19.0 extension control file](https://github.com/postgres/postgres/blob/REL_19_0/contrib/pg_stash_advice/pg_stash_advice.control)
- [SQL API and default privileges](https://github.com/postgres/postgres/blob/REL_19_0/contrib/pg_stash_advice/pg_stash_advice--1.0.sql)

`pg_stash_advice` stores plan-advice strings in named, dynamically allocated shared-memory stashes keyed by PostgreSQL query ID. A session selects a stash with `pg_stash_advice.stash_name`; when a matching query is planned, the saved advice is applied automatically.

### Create and Select a Stash

Preloading is the normal way to make advice application and automatic persistence available in every session.

```ini
shared_preload_libraries = 'pg_plan_advice,pg_stash_advice'
pg_stash_advice.persist = on
pg_stash_advice.persist_interval = 30
```

After restarting, obtain the query ID with `EXPLAIN (VERBOSE)` and the advice string with `EXPLAIN (PLAN_ADVICE)`, then store them:

```sql
CREATE EXTENSION pg_stash_advice;

SELECT pg_create_advice_stash('critical_queries');
SELECT pg_set_stashed_advice(
  'critical_queries',
  :'query_id'::bigint,
  :'plan_advice'
);

SET pg_stash_advice.stash_name = 'critical_queries';
```

Inspection functions are `pg_get_advice_stashes()` and `pg_get_advice_stash_contents(text)`. `pg_drop_advice_stash(text)` removes a stash. The install script revokes public execution of all management and inspection functions.

### Caveats

- This extension is included in PostgreSQL 19 with SQL version `1.0`. It is not supplied by PostgreSQL 18 or earlier.
- Stashes consume dynamic shared memory. Applying advice adds planning overhead and can prevent the planner from adapting as data distribution changes, so constrain both stash size and scope.
- Persistence is enabled by default and writes `pg_stash_advice.tsv` in the data directory through a background worker. Automatic worker startup requires preloading; `pg_start_stash_advice_worker()` can start persistence manually but does not load an existing file.
- The security model is intentionally limited. Although management functions are not public by default, any user may select a known stash name, which may reveal its advice strings. Do not store secrets in plan advice.
- Query IDs require `compute_query_id`; loading the module enables computation when that setting is `auto`. Advice tied to obsolete query IDs must be maintained as queries and schemas change.

The reviewed release is PostgreSQL 19.0. Load `pg_plan_advice` as well as `pg_stash_advice`: the former implements plan advice, while the latter stores and selects it. A nonempty `pg_stash_advice.stash_name` selects the stash to apply.
