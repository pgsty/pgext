## 用法

来源：

- [官方上游 README](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/README.md)
- [官方扩展控制文件 (pg_show_vm.control)](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/pg_show_vm.control)
- [官方扩展 SQL (pg_show_vm--1.0.sql)](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/pg_show_vm--1.0.sql)

`pg_show_vm` — 该扩展支持 PostgreSQL 版本 16 和 17。在管理或自动化上述数据库行为时使用它。使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_show_vm;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_show_rel_vm(IN relname text, IN index bool, IN partition bool, OUT relid int, OUT relpages int, OUT all_visible int, OUT all_frozen int, OUT type int)` 是一个扩展函数，返回 `SETOF`。
- `pg_show_vm(IN relid oid, OUT relid int, OUT relpages int, OUT all_visible int, OUT all_frozen int, OUT type int)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
