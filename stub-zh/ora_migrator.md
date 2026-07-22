## 用法

来源：

- [官方 README](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/README.md)
- [官方迁移示例](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/sql/migrate.sql)
- [官方扩展控制文件](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/ora_migrator.control)

`ora_migrator` 是 `db_migrator` 的 Oracle 插件。它通过 `oracle_fdw` 检查 Oracle 数据库、生成可编辑的迁移元数据暂存区、复制数据与对象，并可选用基于触发器的变更捕获完成低停机切换。

### 核心流程

先安装声明的依赖，创建 `oracle_fdw` 服务器与用户映射，再准备并检查暂存模式，确认后才开始复制。

```sql
CREATE EXTENSION oracle_fdw;
CREATE EXTENSION db_migrator;
CREATE EXTENSION ora_migrator;

SELECT db_migrate_prepare(
  plugin => 'ora_migrator',
  server => 'oracle_src',
  only_schemas => ARRAY['APP'],
  options => '{"max_long": 32767}'::jsonb
);

SELECT schema, table_name, migrate
FROM pgsql_stage.tables
ORDER BY schema, table_name;

SELECT db_migrate_mkforeign(
  plugin => 'ora_migrator',
  server => 'oracle_src',
  options => '{"max_long": 32767}'::jsonb
);

SELECT db_migrate_tables(plugin => 'ora_migrator');
SELECT db_migrate_indexes(plugin => 'ora_migrator');
SELECT db_migrate_constraints(plugin => 'ora_migrator');

SELECT db_migrate_finish();
```

Oracle 模式名通常为大写，因此通过 `only_schemas` 传入的值必须使用 Oracle 端的拼写。`max_long` 控制应用于 LONG、LONG RAW 与 XMLTYPE 列的 `oracle_fdw` 限制。

### 迁移与复制对象

- `create_oraviews` 创建插件使用的 Oracle 数据字典外部表。
- `oracle_test_table` 与 `oracle_migrate_test_data` 在准备完成后检查零字节、编码失败及其他数据转换问题。
- `db_migrator_callback` 将 Oracle 专用的发现和转换逻辑接入通用迁移管线。
- `oracle_replication_start`、`oracle_replication_catchup` 与 `oracle_replication_finish` 管理临时 Oracle 日志表和触发器，以及 PostgreSQL 暂存对象。

### 低停机边界

使用复制时，在建立 `oracle_replication_start` 和初始 `db_migrate_tables` 过渡期间暂停 Oracle 写入。追赶过程使用 Oracle `SERIALIZABLE` 隔离级别；最后一次追赶也应停写，然后再切换应用。源端用户需要访问数据字典；复制还需要较高的 Oracle 建表与触发器权限。追赶期间可能需要将 `session_replication_role` 设为 `replica`，以避免触发目标端触发器与外键检查。

并非所有 Oracle 分区功能都受支持。应确保 Oracle 有足够的 UNDO 支持一致性快照，检查 `pgsql_stage.migrate_log`，并注意 `oracle_replication_finish` 会删除两端的相关对象。上游前置条件为 PostgreSQL `9.5+`、`oracle_fdw` 和 `db_migrator`。
