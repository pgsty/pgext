## 用法

来源：

- [官方 README](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/README.md)
- [1.0 版加载脚本](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table--1.0.sql)
- [ProcessUtility 钩子实现](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table.c)

`pg_force_unlogged_create_table` 安装一个工具命令钩子，把普通 `CREATE TABLE` 与 `CREATE TABLE AS` 改为创建 `UNLOGGED` 表。它适合明确愿意用崩溃持久性与复制覆盖换取装载速度的环境。

### 核心流程

创建扩展时会执行 `LOAD`，因此钩子在当前会话中立即生效。若要让每个新会话都继承该行为，应把库加入 `shared_preload_libraries` 并重启 PostgreSQL。

```sql
CREATE EXTENSION pg_force_unlogged_create_table;
CREATE TABLE staging (id bigint, payload jsonb);

SELECT relpersistence
FROM pg_class
WHERE oid = 'staging'::regclass;
```

预期 `relpersistence` 为 `u`。显式的 `CREATE TEMP TABLE` 与 `CREATE UNLOGGED TABLE` 命令保持不变，`CREATE MATERIALIZED VIEW` 也不受影响。该扩展不定义 SQL 函数或类型。

### 持久性边界

非日志关系会在非正常停机后被清空，也不会为物理复制写入 WAL。非日志表上的索引同样不记录日志。只能将它用于可丢弃、可重建的数据，并确认备份、故障切换、逻辑处理及恢复流程都不依赖这些表的持久性。

会话级加载只影响库加载后执行的语句。删除扩展会移除其成员对象，但不会把已有非日志表转回日志表；如果持久性要求发生变化，应显式转换或重建这些关系。
