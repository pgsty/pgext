## 用法

来源：

- [已复核 commit 的原型 README](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/README.md)
- [扩展 control 文件](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/pg_data_retention.control)
- [后台工作进程源码](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/src/lib.rs)
- [Cargo 特性矩阵](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/Cargo.toml)

`pg_data_retention` 是 0.0.0 版本的后台工作进程实验，根据 `public.data_retention_policy` 中的记录删除数据行。它必须加入 `shared_preload_libraries`，随后重启服务器。

```sql
CREATE EXTENSION pg_data_retention;
SELECT *
FROM public.data_retention_policy
ORDER BY database_name, schema_name, table_name;
```

### 破坏性原型警告

当前源码在启动时连接 `postgres` 数据库，删除已有的 `public.data_retention_policy`，再重新创建它。编排进程每十秒唤醒一次，按顺序启动删除工作进程，并在 SQL 失败时 panic。文档中的 `cron_schedule` 字段尚未实现，仓库也明确列出工作进程与跨版本测试缺失。

不要在含有重要数据的服务器上预加载此代码。重建策略表会破坏配置，而错误的表、时间戳列、保留期或批量大小会删除应用数据。该源码只能作为原型，在隔离环境中审计并重新设计。
