## 用法

来源：

- [官方扩展控制文件 (pgscale.control)](https://github.com/kelvich/pgscale/blob/5f24b0db82f3bfb48d8f1e5d6ca1441b543050d8/pgscale.control)
- [官方扩展 SQL (pgscale--1.0.sql)](https://github.com/kelvich/pgscale/blob/5f24b0db82f3bfb48d8f1e5d6ca1441b543050d8/pgscale--1.0.sql)

`pgscale` — 背景工作者通过简单的 HTTP 端点暴露 PostgreSQL 统计视图。在收集或解释相应的 PostgreSQL 统计信息时使用它。在目标 PostgreSQL 构建上测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pgscale;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
