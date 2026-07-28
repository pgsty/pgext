## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/glia/glia-0.0.1/README.md)
- [官方扩展控制文件 (glia.control)](https://api.pgxn.org/src/glia/glia-0.0.1/glia.control)
- [官方扩展 SQL (glia--0.0.1.sql)](https://api.pgxn.org/src/glia/glia-0.0.1/glia--0.0.1.sql)

`glia` — *一个用于数据挖掘的 PostgreSQL 扩展*. 请使用它来执行相应的向量、模型或检索工作流。在目标 PostgreSQL 构建上测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION glia;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 2026-07-28 审查期间前 GitHub 仓库 URL 返回 404；请将上方的固定 PGXN 发行版视为可用源边界。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
