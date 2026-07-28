## 用法

来源：

- [官方上游 README](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/README.md)
- [官方扩展控制文件 (pgtoken.control)](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/pgtoken.control)
- [官方扩展 SQL (pgtoken--1.0.sql)](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/pgtoken--1.0.sql)

`pgtoken` — PostgreSQL 扩展，用于 rank-varint 令牌存储。将 LLM 令牌 ID 压缩为 bytea 存储。读取时不进行重新标记化。将其用于相应的向量、模型或检索工作流。在目标 PostgreSQL 构建中使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgtoken;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgtoken_count(encoded bytea)` 是一个扩展函数，返回 `integer`。
- `pgtoken_decode(encoded bytea, codebook text DEFAULT 'cl100k_base')` 是一个扩展函数，返回 `integer[]`。
- `pgtoken_encode(token_ids integer[], codebook text DEFAULT 'cl100k_base')` 是一个扩展函数，返回 `bytea`。
- `pgtoken_reload_codebooks()` 是一个扩展函数，返回 `void`。
- `pgtoken_codebooks` 是一个扩展定义的视图。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
