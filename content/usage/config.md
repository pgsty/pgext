---
title: Config
description: Preload extensions and configure extension parameters
icon: FileBox
---


While most PostgreSQL extensions written in SQL can be [**directly enabled**](/usage/create) with `CREATE EXTENSION`, some extensions that use special postgres hook will require an extra step to **preload** them before using.



------

## Preloading


Most extensions have one or more corresponding dynamic library (`.so`, `.dylib`, `.dll`), some of them require preloading before using.
Attempting to `CREATE` these extensions without proper preloading will result in an error. 
And a wrongly configured preload library may lead to a failure on database restart/start.

Some extensions can partially work without preloading, which means part of the extension features are available directly, and the rest of the features are available after preloading.

To preload an extension, add it to the [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#Rshared_preload_libraries) and restart the database server.
The [Extension Catalog](/list/attr#need-loading) gives the complete list of extensions that require dynamic preloading.



------

## Configure

To configure a preload on new postgres cluster, the  [`pg_libs`](/docs/pgsql/param#pg_libs) parameter can be used.
It will be populated to the `shared_preload_libraries` parameter during postgres cluster bootstrap.

<Callout title="Example: Setup Supabase Extension Preloading">

This example show how to specify pre-loaded extensions with `pg_libs` parameter.

```yaml
all:
  children:
    pg-meta:
      hosts: { 10.10.10.10: { pg_seq: 1, pg_role: primary } }
      vars:
        pg_cluster: pg-meta
        pg_libs: 'timescaledb, plpgsql, plpgsql_check, pg_cron, pg_net, pg_stat_statements, auto_explain, pg_tle, plan_filter'
```

    `shared_preload_libraries` is a comma-separated list of extensions.

</Callout>


Beware that only works before [**cluster creation**](/docs/pgsql/admin#create-cluster). After that,
you'll have to [**config cluster**](/docs/pgsql/admin#config-cluster) to change the `shared_preload_libraries` parameter on existing cluster. (with `patronictl`, `ALTER SYSTEM`, etc...)

```bash title="add timescaledb to shared_preload_libraries"
pg edit-config pg-meta --force -p shared_preload_libraries='timescaledb, pg_stat_statements, auto_explain'
pg restart pg-meta    # restart to apply changes
````

If you want to configure preloading manually, you can just change the `postgresql.conf` by yourself



--------

## Default

The default value of [`pg_libs`](/docs/pgsql/param#pg_libs) is `pg_stat_statements, auto_explain`,
which preload these two [Contrib extensions](/list/repo#contrib) by default, these two extensions provide essential observability:

- [**`auto_explain`**](/e/auto_explain): Automatic logging of slow query execution plans
- [**`pg_stat_statements`**](/e/pg_stat_statements): Tracks planning and execution statistics for grouped SQL statements


--------

## Caveats

Preload libraries are loaded one by one, so the order of extensions in `shared_preload_libraries` matters,
Here are some known rules to follow:

- For [**STAT**](/cate/stat) extension, add them **AFTER** `pg_stat_statements` to ensure using the same query_id.
- [`timescaledb`](/e/timescaledb) and [`citus`](/e/citus) should be placed at the **BEGINNING** of `shared_preload_libraries`
- If you use [`citus`](/e/citus) and [`timescaledb`](/e/timescaledb) together, place `citus` before `timescaledb`.
- Use `pg_documentdb` and `pg_documentdb_core` as library name for [**documentdb**](/e/documentdb).
- [`pg_search`](/e/pg_search) does not require preloading in PostgreSQL 17 and later, but earlier versions do.

-------

## Parameter

Some extensions have configurable parameters, you can manage them in different places.

- [`pg_parameters`](/docs/pgsql/param#pg_parameters): write to `/pg/data/postgresql.auto.conf`
- You can also customize them in [patroni templates](https://github.com/pgsty/pigsty/blob/main/roles/pgsql/templates/oltp.yml#L404)
- Or dynamic change them with [**patronictl**](/docs/pgsql/admin#config-cluster)

Consult the official docs of each extension for details.