## 用法

来源：

- [官方上游 README](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/README.md)
- [官方扩展控制文件 (pg_tm_aux.control)](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/pg_tm_aux.control)
- [官方扩展 SQL (pg_tm_aux--1.0.sql)](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/pg_tm_aux--1.0.sql)

`pg_tm_aux` — **(自 Postgres 17 起不再需要，使用逻辑槽故障转移)**。在移动、转换或集成相应数据时使用它。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_tm_aux;

SELECT * from  pg_create_logical_replication_slot_lsn('dtt3gjq2tfmocenb6vru', 'wal2json', false, pg_lsn('1/20030948'));
SELECT * from pg_logical_slot_peek_changes('dtt3gjq2tfmocenb6vru', null, null);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_create_logical_replication_slot_lsn(IN slot_name name, IN plugin name, IN temporary boolean DEFAULT false, IN restart_lsn pg_lsn DEFAULT null, IN force boolean DEFAULT true, OUT slot_name name, OUT lsn pg_lsn)` 是一个扩展函数，返回 `RECORD`。
- `pg_create_logical_replication_slot_lsn(IN slot_name name, IN plugin name, IN temporary boolean DEFAULT false, IN restart_lsn pg_lsn DEFAULT null, OUT slot_name name, OUT lsn pg_lsn)` 是一个扩展函数，返回 `RECORD`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.1.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
