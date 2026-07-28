## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/check_updates/check_updates-1.0.1/README.md)
- [官方扩展控制文件 (check_updates.control)](https://api.pgxn.org/src/check_updates/check_updates-1.0.1/check_updates.control)

`check_updates` — 确保你已经安装了 pg_config 并将其添加到路径中。如果你使用的是 RPM 包管理器安装 PostgreSQL，请确保也安装了 -devel 包。如果必要，告诉构建过程 pg_config 的位置：使用它来管理和自动化上述数据库行为。在目标 PostgreSQL 版本上使用链接的固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION check_updates;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.1`。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
