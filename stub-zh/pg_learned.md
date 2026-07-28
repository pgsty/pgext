## 用法

来源：

- [官方上游 README](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/README.md)
- [官方扩展控制文件 (pg_learned.control)](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/pg_learned.control)
- [官方实现源代码](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/c_impl/pg_learned.c)

`pg_learned` — **一个展示学习索引技术以加快数据库查询速度的 PostgreSQL 扩展。** 当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_learned;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
