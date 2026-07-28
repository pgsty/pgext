## 用法

来源：

- [官方上游 README](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/README.md)
- [官方扩展控制文件 (pgvecutils.control)](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/pgvecutils.control)
- [官方扩展 SQL (pgvecutils--0.0.1.sql)](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/pgvecutils--0.0.1.sql)

`pgvecutils` — 这个项目是一个 PostgreSQL 扩展，提供了向量操作工具。使用它来进行相应的向量、模型或检索工作流。在安装扩展及其依赖项之前，请确保已正确安装并验证了这些依赖项。

### 核心工作流

```sql
CREATE EXTENSION pgvecutils;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `vector_rand(int)` 是一个扩展函数，返回 `vector`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 首先安装并验证确认的扩展依赖项：`vector`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
