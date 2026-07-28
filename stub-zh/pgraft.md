## 用法

来源：

- [官方上游 README](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/README.md)
- [官方扩展控制文件 (pgraft.control)](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/pgraft.control)
- [官方扩展 SQL (pgraft--1.0.sql)](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/pgraft--1.0.sql)

`pgraft` — **pgraft** 是一个高性能的 PostgreSQL 扩展，实现了分布式 PostgreSQL 集群中的 Raft 共识协议。它能够实现自动领导者选举、日志复制和故障容错。在管理或自动化上述数据库行为时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgraft;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pgraft_add_node(node_id integer, address text, port integer)` 是一个扩展函数，返回 `boolean`。
- `pgraft_get_cluster_status()` 是一个扩展函数，返回 `TABLE`。
- `pgraft_get_leader()` 是一个扩展函数，返回 `bigint`。
- `pgraft_get_nodes()` 是一个扩展函数，返回 `TABLE`。
- `pgraft_get_queue_status()` 是一个扩展函数，返回 `TABLE`。
- `pgraft_get_term()` 是一个扩展函数，返回 `integer`。
- `pgraft_get_version()` 是一个扩展函数，返回 `text`。
- `pgraft_get_worker_state()` 是一个扩展函数，返回 `text`。
- `pgraft_init()` 是一个扩展函数，返回 `boolean`。
- `pgraft_is_leader()` 是一个扩展函数，返回 `boolean`。
- `pgraft_log_append(term bigint, data text)` 是一个扩展函数，返回 `boolean`。
- `pgraft_log_apply(index bigint)` 是一个扩展函数，返回 `boolean`。
- `pgraft_log_commit(index bigint)` 是一个扩展函数，返回 `boolean`。
- `pgraft_log_get_entry(index bigint)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
