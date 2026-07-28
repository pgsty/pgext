## 用法

来源：

- [官方上游 README](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/README.md)
- [官方扩展控制文件 (pg_fts_greek.control)](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/pg_fts_greek.control)
- [官方扩展 SQL (pg_fts_greek--0.1.sql)](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/sql/pg_fts_greek--0.1.sql)

`pg_fts_greek` — Postgres FTS 改进用于希腊语。使用它来进行相应的文本搜索、解析或语言工作流。在目标 PostgreSQL 构建上测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_fts_greek;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本 `0.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
