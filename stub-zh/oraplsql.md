## 用法

来源：

- [官方上游 README](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/README.md)
- [官方扩展控制文件 (oraplsql.control)](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/src/pl/oraplsql/src/oraplsql.control)
- [官方扩展 SQL (oraplsql--1.0.sql)](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/src/pl/oraplsql/src/oraplsql--1.0.sql)

`oraplsql` — Oracle兼容的程序化SQL语言，用于 OpenTenBase。在移植或模拟相应的数据库API时使用它。使用上述链接的固定上游版本作为API边界，并在目标PostgreSQL构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION oraplsql;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序SQL之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的PostgreSQL版本、升级行为和失败情况与固定源代码中的信息一致。
