## 用法

来源：

- [官方上游 README](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/readme.md)
- [官方扩展控制文件 (shardman.control)](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/ext/shardman.control)
- [官方扩展 SQL (shardman--0.0.1.sql)](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/ext/shardman--0.0.1.sql)

`shardman` — 实验性分片、重平衡和跨节点锁图原型。请在相应的分析或存储工作流中使用它。在目标 PostgreSQL 构建上测试时，请使用上述链接的固定上游版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION shardman;

create table pt (id serial, payload real) partition by hash(id);
select shardman.hash_shard_table('pt', 10)
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `bcst_all_sql(cmd text)` 是一个扩展函数，返回 `void`。
- `bcst_sql(cmd text)` 是一个扩展函数，返回 `void`。
- `deny_access()` 是一个扩展函数，返回 `trigger`。
- `drop_repslot(slot_name text, with_force bool default true)` 是一个扩展函数，返回 `void`。
- `eliminate_sub(subname name)` 是一个扩展函数，返回 `void`。
- `ex_sql(rgid int, cmd text)` 是一个扩展函数，返回 `void`。
- `hash_shard_table(relid regclass, nparts int, colocate_with regclass = null)` 是一个扩展函数，返回 `void`。
- `is_subscription_ready(sname text)` 是一个扩展函数。
- `part_moved(relid regclass, pnum int, src_rgid int, dst_rgid int)` 是一个扩展函数，返回 `void`。
- `postgres_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `postgres_fdw_validator(text[], oid)` 是一个扩展函数，返回 `void`。
- `rebalance_cleanup()` 是一个扩展函数，返回 `void`。
- `rebalance_cleanup_local()` 是一个扩展函数，返回 `void`。
- `restore_foreign_tables()` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
