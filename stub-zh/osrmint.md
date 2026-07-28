## 用法

来源：

- [官方上游 README](https://github.com/fpacheco/osrmint/blob/c1937d4153cf8a38712e9e996e603564a4cd8e1e/README.md)
- [官方扩展控制文件 (osrmint.control)](https://github.com/fpacheco/osrmint/blob/c1937d4153cf8a38712e9e996e603564a4cd8e1e/control/osrmint.control)

`osrmint` — PostgreSQL 路径 OSRM 集成。使用它来进行相应的空间数据或地理空间工作流。在集成之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION osrmint;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 首先安装确认的扩展依赖项：`postgis`。
- 控制文件将该扩展标记为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
