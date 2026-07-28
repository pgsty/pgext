## 用法

来源：

- [官方上游 README](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/docs/README.md)
- [官方扩展控制文件 (veil2_demo.control)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/veil2_demo.control)

`veil2_demo` — 提供了 veil2 扩展的示例数据库。在实现相应的安全、审计或访问控制工作流时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION veil2_demo;
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.9.3`。
- 首先安装并验证确认的扩展依赖项：`veil2`。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
