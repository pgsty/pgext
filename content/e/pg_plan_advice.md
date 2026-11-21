---
title: "pg_plan_advice"
linkTitle: "pg_plan_advice"
description: "Describe, reproduce, and constrain query plans with plan advice"
weight: 2988
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["CONTRIB"]
page_width: full
---

[**pg_plan_advice**](https://www.postgresql.org/docs/19/pgplanadvice.html) : Describe, reproduce, and constrain query plans with plan advice


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2988** | {{< badge content="pg_plan_advice" link="https://www.postgresql.org/docs/19/pgplanadvice.html" >}} | {{< ext "pg_plan_advice" >}} | `-` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_stash_advice" >}} |
|   **See Also**    | {{< ext "pg_stash_advice" >}} {{< ext "pg_hint_plan" >}} {{< ext "plan_filter" >}} {{< ext "pg_overexplain" >}} |


## Packages

| **PG19** | **PG18** | **PG17** | **PG16** | **PG15** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 19: version -" "green" >}} | {{< bg "N/A" "PostgreSQL 18: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 16: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 15: not available" "red" >}} |
{.wide .expand}

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_plan_advice';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [PostgreSQL 19 documentation](https://www.postgresql.org/docs/19/pgplanadvice.html)
- [Module definition](https://github.com/postgres/postgres/blob/REL_19_0/contrib/pg_plan_advice/Makefile)
- [Module implementation](https://github.com/postgres/postgres/blob/REL_19_0/contrib/pg_plan_advice/pg_plan_advice.c)
- [Versioned manual](https://github.com/postgres/postgres/blob/REL_19_0/doc/src/sgml/pgplanadvice.sgml)

`pg_plan_advice` is a supplied PostgreSQL 19 module that records planner choices as advice and constrains later planning. This page uses PostgreSQL 19.0 as its version boundary; the module has no independent SQL extension version.

### Session Workflow

```sql
LOAD 'pg_plan_advice';
CREATE TEMP TABLE advice_demo (id integer PRIMARY KEY, body text);
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
SET pg_plan_advice.advice = 'SEQ_SCAN(d)';
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
RESET pg_plan_advice.advice;
```

### Enablement and Feedback

Use `LOAD` with sufficient privileges for one session, `session_preload_libraries` for new sessions, or `shared_preload_libraries` followed by a server restart. There is no extension DDL step. `EXPLAIN (PLAN_ADVICE)` generates advice using relation aliases as targets. `pg_plan_advice.advice` supplies constraints on scans, join order, join methods, or parallelism.

Inspect the returned feedback for unmatched, conflicting, inapplicable, or failed advice. `pg_plan_advice.feedback_warnings` can emit warnings, and `pg_plan_advice.always_store_advice_details` preserves details useful when explaining prepared statements at additional planning cost.

### Limits

Advice only constrains plans the core planner considers viable. It cannot force a semantically invalid plan, control aggregation strategy, or control set operations. Fixed advice can age poorly as data changes. `pg_stash_advice` optionally stores advice by query ID; `pg_plan_guard` monitors drift from approved baselines.
