## 用法

来源：

- [官方上游 README](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/README.md)
- [官方扩展控制文件 (pyramid_generate.control)](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/pyramid_generate/pyramid_generate.control)
- [官方扩展 SQL (pyramid_generate--1.0.sql)](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/pyramid_generate/pyramid_generate--1.0.sql)

`pyramid_generate` — 生成均匀分布的随机 float8 向量，按维度和大小生成。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pyramid_generate;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pyramid_generate(dimension int4, size int8)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件将扩展标记为可信。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
