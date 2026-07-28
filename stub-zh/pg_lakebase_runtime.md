## 用法

来源：

- [官方上游 README](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/README.md)
- [官方扩展控制文件 (pg_lakebase_runtime.control)](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-lakebase-runtime/pg_lakebase_runtime.control)
- [官方实现源代码](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-lakebase-runtime/src/lib.rs)

`pg_lakebase_runtime` — 共享运行时库，用于 Lakebase 表访问方法扩展。使用它来实现相应的分析或存储工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_lakebase_runtime;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `deregister_worker` 是一个扩展函数。
- `maintenance_worker` 是一个扩展函数。
- `observe_object_tree` 是一个扩展函数。
- `register_worker_impl` 是一个扩展函数。
- `request_worker_wakeup` 是一个扩展函数。
- `retry_maintenance_item` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
