## 用法

来源：

- [官方上游 README](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/README.md)
- [官方扩展控制文件 (pg_llm_bgw.control)](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/pg_llm_bgw/pg_llm_bgw.control)
- [官方实现源代码](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/pg_llm_bgw/src/lib.rs)

`pg_llm_bgw` — 一个用 Rust 实现的后台工作者示例，用于在 Postgres 客户端 DML SQL 中调用 LLM。使用它来执行相应的向量、模型或检索工作流。在目标 PostgreSQL 构建上测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_llm_bgw;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证安装版本和返回值。

### 重要对象

- `llm_ask` 是一个扩展函数。
- `llm_provider()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
