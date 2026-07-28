## 用法

来源：

- [官方上游 README](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/readme.md)
- [官方扩展控制文件 (pg_embed.control)](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/pg_embed.control)
- [官方扩展 SQL (pg_embed--1.0.sql)](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/pg_embed--1.0.sql)

`pg_embed` — pg_embed 是一个使用 PL/Python 编写的 PostgreSQL 扩展，它允许您直接在数据库中生成和存储文本嵌入，使用 HuggingFace 的 Inference API。此扩展适用于为基于 PostgreSQL 的应用程序添加语义搜索、相似性和 AI 功能。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_embed;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `generate_and_store_embeddings(tbl_name TEXT, txt_col TEXT, hf_api_key TEXT)` 是一个扩展函数，返回 `VOID`。
- `get_embedding(input_text TEXT, hf_api_key TEXT)` 是一个扩展函数，返回 `FLOAT8[]`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
