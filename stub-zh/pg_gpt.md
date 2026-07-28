## 用法

来源：

- [官方上游 README](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/README.md)
- [官方扩展控制文件 (pg_gpt.control)](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/pg_gpt.control)
- [官方实现源代码](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/src/lib.rs)

`pg_gpt` — 这是一个实验性的 PostgreSQL 扩展，允许在 PostgreSQL 中使用 OpenAI GPT API，从而可以使用自然语言编写查询。请使用此扩展进行相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_gpt;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gpt` 是一个扩展函数。
- `gpt_tables` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
