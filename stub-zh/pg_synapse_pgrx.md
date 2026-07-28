## 用法

来源：

- [官方上游 README](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/README.md)
- [官方扩展控制文件 (pg_synapse_pgrx.control)](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/crates/pg-synapse-pgrx/pg_synapse_pgrx.control)
- [官方实现源代码](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/crates/pg-synapse-pgrx/src/lib.rs)

`pg_synapse_pgrx` — pg synapse: Postgres-native agent-loop 运行时。使用它来执行相应的向量、模型或检索工作流。在目标 PostgreSQL 构建上使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_synapse_pgrx;

SELECT synapse.execute('notes_agent', 'Add a note that says "Hello"');
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本 `0.1.1`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
