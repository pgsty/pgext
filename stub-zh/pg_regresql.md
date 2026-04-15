## 用法

> 来源：[扩展 README](https://github.com/boringSQL/regresql/blob/master/pg_ext/README.md)，[控制文件](https://github.com/boringSQL/regresql/blob/master/pg_ext/pg_regresql.control)，[portable stats 文章](https://boringsql.com/posts/portable-stats/)

`pg_regresql` 是一个 PostgreSQL 扩展，用来让优化器优先信任 `pg_class` 中的目录统计信息，而不是根据物理文件块数重新估算关系大小。它是 RegreSQL 项目中的扩展部分，主要用于在注入生产统计信息后，复现更接近真实环境的执行计划。

### 要解决的问题

上游扩展 README 指出，PostgreSQL 在估算关系大小时，通常不会完全信任 `pg_class.relpages` 和 `pg_class.reltuples`。规划器会读取当前物理文件大小，并据此重新缩放这些统计值。

这种行为有助于规避陈旧统计信息带来的风险，但在测试环境中会破坏“已从其他环境恢复目录统计信息”的场景，因为本地表文件通常远小于生产环境。

### 它覆盖了什么

`pg_regresql` 通过 `get_relation_info_hook` 在 `estimate_rel_size()` 之后接管规划器估算，并用目录统计值替换默认结果。

| 规划器字段 | 默认来源 | `pg_regresql` 来源 |
| --- | --- | --- |
| `rel->pages` | 通过表访问方法调用 `smgrnblocks()` | `pg_class.relpages` |
| `rel->tuples` | 按物理页数缩放后的密度估算 | `pg_class.reltuples` |
| `rel->allvisfrac` | `relallvisible / physical pages` | `pg_class.relallvisible / relpages` |
| `IndexOptInfo->pages` | `RelationGetNumberOfBlocks()` | 索引的 `pg_class.relpages` |
| `IndexOptInfo->tuples` | 复制自 `rel->tuples` | 索引的 `pg_class.reltuples` |

### 安装

上游 README 给出了三种安装方式：

```bash
sudo pgxn install pg_regresql
```

```bash
make PG_SOURCE=/path/to/postgresql
make install PG_SOURCE=/path/to/postgresql
```

```bash
make USE_PGXS=1
make install USE_PGXS=1
```

控制文件为 `pg_regresql.control`，其中声明了 `default_version = '2.0'` 和 `module_pathname = '$libdir/pg_regresql'`。

### 启用方式

只有在共享库被加载后，这个扩展才会真正接管规划器：

```sql
LOAD 'pg_regresql';

EXPLAIN SELECT ...;
```

如果希望在整个测试实例中启用，README 推荐：

```conf
session_preload_libraries = 'pg_regresql'
```

这里的关键是运行时加载配置，而不是“装完包就自动生效”；只有在当前会话或实例加载了库之后，规划器 hook 才会起作用。

### 典型工作流

它的主要用途是基于已恢复的生产统计信息做执行计划回归测试。在 CI 或测试库中注入目录统计信息后，`pg_regresql` 会强制规划器使用这些恢复值，而不是依赖本地很小的堆表文件大小。

README 给出的示例是：

```sql
SELECT pg_restore_relation_stats(
    'schemaname', 'public',
    'relname', 'test_orders',
    'relpages', 123513::integer,
    'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);

LOAD 'pg_regresql';

EXPLAIN SELECT * FROM test_orders WHERE created_at > '2024-06-01';
```

这种模式适合：

- 在本地复现生产执行计划
- 用更真实的计划估算测试 schema migration
- 模拟表增长后的索引与访问路径选择
- 改进分区规划相关实验

### 兼容性

- 在本仓库中打包支持 PostgreSQL 14 及以上版本
- 上游 README 提到该 hook 自 PostgreSQL 8.3 起就存在
- 设计上应可与 `pg_hint_plan`、`hypopg` 等扩展共存，但上游说明这一点尚未完整测试
