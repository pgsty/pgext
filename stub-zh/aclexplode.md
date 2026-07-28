## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/README.md)
- [官方扩展控制文件 (aclexplode.control)](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/aclexplode.control)
- [官方扩展 SQL (aclexplode.sql)](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/sql/aclexplode.sql)

`aclexplode` — 确保你已经安装了 pg_config 并将其添加到路径中。如果你使用的是 RPM 包管理器安装 PostgreSQL，请确保也安装了 -devel 包。如果需要，可以在构建过程中指定其位置：使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 版本上进行测试。

### 核心工作流

```sql
CREATE EXTENSION aclexplode;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `aclexplode(aclitem[], OUT grantor oid, OUT grantee oid, OUT privilege_type text, OUT is_grantable bool)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.3`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
