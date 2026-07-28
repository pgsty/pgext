## 用法

来源：

- [官方上游 README](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/README.md)
- [官方扩展控制文件 (pg_delta_am.control)](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-delta-am/pg_delta_am.control)
- [官方实现源代码](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-delta-am/src/lib.rs)

`pg_delta_am` — Delta 表访问方法骨架。使用它来实现相应的分析或存储工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_delta_am;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `delta_table_am_handler()` 是一个扩展函数，返回 `pg_lakebase_core::TableAmRoutine`。

### 要求与注意事项

- 元数据记录版本 `0.1.0`。
- 首先安装已确认的扩展依赖项：`pg_lakebase_runtime`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
