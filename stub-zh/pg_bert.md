## 用法

来源：

- [官方上游 README](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/README.md)
- [官方扩展控制文件 (pg_bert.control)](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/pg_bert.control)
- [官方实现源代码](https://github.com/usamoi/pg_bert/blob/856a2a418b530ee4540541dcfe7e84278e3fd209/src/lib.rs)

`pg_bert` — BERT 分词器实现为 PostgreSQL 全文搜索解析器。使用它进行相应的文本搜索、解析或语言工作流。在目标 PostgreSQL 构建上测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_bert;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
