---
title: 配置
description: 预加载扩展与配置扩展参数
icon: FileBox
---

大多数用 SQL 编写的 PostgreSQL 扩展可以[**直接启用**](/zh/usage/create)（`CREATE EXTENSION`），但部分依赖特殊 Postgres Hook 的扩展需**预加载**后方可使用。

------

## 预加载

大多数扩展都包含一个或多个动态库（`.so`、`.dylib`、`.dll`），部分扩展在使用前需预加载。
若未正确预加载即尝试 `CREATE` 这些扩展，将导致报错。
而预加载配置错误则可能导致数据库启动/重启失败。

部分扩展支持“部分预加载”，即部分功能可直接用，完整功能需预加载后方可用。

要预加载扩展，请将其加入[`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#Rshared_preload_libraries)参数，并重启数据库服务。
[扩展目录](/zh/list/attr#need-loading)列出了所有需动态预加载的扩展。

------

## 配置

如需在新建 Postgres 集群时配置预加载，可使用 [`pg_libs`](/zh/docs/pgsql/param#pg_libs) 参数。
该参数会在集群初始化时写入 `shared_preload_libraries`。

<Callout title="示例：配置 Supabase 扩展预加载">

以下示例展示如何通过 `pg_libs` 参数指定需预加载的扩展：

```yaml
all:
  children:
    pg-meta:
      hosts: { 10.10.10.10: { pg_seq: 1, pg_role: primary } }
      vars:
        pg_cluster: pg-meta
        pg_libs: 'timescaledb, plpgsql, plpgsql_check, pg_cron, pg_net, pg_stat_statements, auto_explain, pg_tle, plan_filter'
```

    `shared_preload_libraries` 为逗号分隔的扩展列表。

</Callout>

注意：此方法仅适用于[**集群创建前**](/zh/docs/pgsql/admin#create-cluster)。集群创建后，需通过[**集群配置**](/zh/docs/pgsql/admin#config-cluster)修改 `shared_preload_libraries`（如用 `patronictl`、`ALTER SYSTEM` 等）。

```bash title="将 timescaledb 加入 shared_preload_libraries"
pg edit-config pg-meta --force -p shared_preload_libraries='timescaledb, pg_stat_statements, auto_explain'
pg restart pg-meta    # 重启以应用配置
```

如需手动配置预加载，也可直接编辑 `postgresql.conf` 文件。

--------

## 默认值

[`pg_libs`](/zh/docs/pgsql/param#pg_libs) 默认值为 `pg_stat_statements, auto_explain`，即默认预加载这两个 [Contrib 扩展](/zh/list/repo#contrib)，它们提供基础可观测性：

- [**`auto_explain`**](/zh/e/auto_explain)：自动记录慢 SQL 执行计划
- [**`pg_stat_statements`**](/zh/e/pg_stat_statements)：统计分组 SQL 的执行与规划信息

--------

## 注意事项

预加载库按顺序加载，`shared_preload_libraries` 中扩展顺序有讲究，常见规则如下：

- [**STAT**](/zh/cate/stat) 类扩展应放在 `pg_stat_statements` 之后，以复用 query_id。
- [`timescaledb`](/zh/e/timescaledb) 与 [`citus`](/zh/e/citus) 应置于 `shared_preload_libraries` 开头。
- 若同时用 [`citus`](/zh/e/citus) 与 [`timescaledb`](/zh/e/timescaledb)，应先 `citus` 后 `timescaledb`。
- [**documentdb**](/zh/e/documentdb) 需用 `pg_documentdb` 与 `pg_documentdb_core` 作为库名。
- [`pg_search`](/zh/e/pg_search) 在 PostgreSQL 17 及以上无需预加载，低版本需预加载。

-------

## 参数

部分扩展支持参数配置，可通过多种方式管理：

- [`pg_parameters`](/zh/docs/pgsql/param#pg_parameters)：写入 `/pg/data/postgresql.auto.conf`
- 也可在 [patroni 模板](https://github.com/pgsty/pigsty/blob/main/roles/pgsql/templates/oltp.yml#L404)中自定义
- 或用 [**patronictl**](/zh/docs/pgsql/admin#config-cluster) 动态修改

具体参数说明请查阅各扩展官方文档。