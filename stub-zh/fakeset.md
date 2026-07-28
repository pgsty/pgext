## 用法

来源：

- [官方上游 README](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/README.md)
- [官方扩展控制文件 (fakeset.control)](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/fakeset.control)
- [官方实现源代码](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/src/lib.rs)

`fakeset` — 一个用于生成假数据的 postgres 扩展。当 SQL 需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION fakeset;

select lorem(3, 10) from generate_series(1, 2000);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `lorem` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
