## 用法

来源：

- [Official README](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/README.md)
- [Extension control](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/ch_table_engines.control)
- [Versioned SQL](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/sql/ch_table_engines--0.1.sql)
- [Preload boundary test](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/sql/nopreload.sql)
- [Settings](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/src/module.c)

`ch_table_engines` 0.1 在未经修改的 PostgreSQL 19 上提供 ClickHouse MergeTree 系列表访问方法，由 PostgreSQL 系统表控制快照可见性。它依赖另行提供的 `libchte_engine.so`，该库由 ClickHouse 26.8.8.8-lts 源码构建。这是早期项目，与 ClickHouse, Inc. 无隶属关系。

### 核心流程

持久表要求预加载并重启 PostgreSQL。如果引擎库的默认位置与安装位置不符，通过 `ch.engine_library` 配置：

```ini
shared_preload_libraries = 'ch_table_engines'
```

以超级用户安装不可迁移的扩展，其对象位于模式 `ch`：

```sql
CREATE EXTENSION ch_table_engines;
CREATE TABLE events (account_id int, event_day date, amount bigint)
    USING mergetree WITH (order_by = 'account_id, event_day', storage = 'pages');
INSERT INTO events VALUES (1, DATE '2026-10-22', 100);
SELECT account_id, sum(amount) FROM events GROUP BY account_id;
SELECT * FROM ch.parts('events'::regclass);
```

TEMP 与 UNLOGGED 表可在未预加载时工作，持久表则需要服务器启动时注册的 WAL 资源管理器。常规持久化部署应采用预加载配置，同时启用合并工作进程。

### 引擎、选项与维护

安装的访问方法包括 `mergetree`、`replacingmergetree`、`summingmergetree`、`aggregatingmergetree`、`collapsingmergetree`、`versionedcollapsingmergetree` 和 `coalescingmergetree`。此版本的自动与会话合并只适用于 MergeTree 表；不能假定其他访问方法已具备相应的合并归约行为。

- `storage` 选择缓冲池管理的关系页或文件目录；`order_by`、`primary_key`、`index_granularity`、`compression` 与 `low_cardinality` 描述数据片布局。布局选项只能在空表上修改。`primary_key` 选项是存储键，不是 PostgreSQL 唯一性约束。
- `ch.parts`、`ch.definition`、`ch.table_size` 和 `ch.engine_info` 用于检查表和引擎。
- `ch.fold` 重写数据片，移除所有快照都已不可见的删除行；`ch.optimize(rel, final)` 合并数据片。系统表可见性允许后，VACUUM 才会释放过时数据片的存储。
- `ch.max_merge_workers` 默认为 2，`ch.merge_naptime` 默认为 10 秒。工作进程不会在备库合并。

### 使用限制

不支持唯一索引、排除约束、并发索引构建或 BRIN，因此也不支持主键约束及 ON CONFLICT 所需的唯一仲裁索引。目前所有行锁模式都按排他方式工作。采用前须验证所需类型、事务、备份和恢复行为；索引与合并限制直接影响存储使用方式。
