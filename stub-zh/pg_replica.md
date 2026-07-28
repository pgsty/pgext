## 用法

来源：

- [官方上游 README](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/README.md)
- [官方扩展控制文件 (pg_replica.control)](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/pg_replica.control)
- [官方实现源代码](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/src/lib.rs)

`pg_replica` — 一个 PostgreSQL 扩展，为一组 **纯正的 Postgres** 节点提供自动、共识驱动的故障转移功能 —— 全集群复制（表、角色、DDL，一切内容）并内置了 Raft 组，无需外部依赖：无需 etcd、Consul 或 Kubernetes。在管理或自动化上述数据库行为时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_replica;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `rotate_credential` 是一个扩展函数。
- `status()` 是一个扩展函数。

### 要求与注意事项

- 表格记录版本 `0.7.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和故障情况，与固定源代码进行比对。
