## 用法

来源：

- [官方上游 README](https://github.com/jwdeitch/pg_cmake_template/blob/6eafd52da9790db53fb32b669f77350b8476a1e5/readme)
- [官方扩展控制文件 (demopgextension.control)](https://github.com/jwdeitch/pg_cmake_template/blob/6eafd52da9790db53fb32b669f77350b8476a1e5/demopgextension.control)

`demopgextension` — 该项目提供了一个 CMake 模板用于开发 PostgreSQL 扩展，并旨在替代 PGXS。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建上测试链接的上游固定版本作为 API 边界，并验证安装的版本和返回值，然后再将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION demopgextension;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
