## 用法

来源：

- [官方上游 README](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/README.md)
- [官方扩展控制文件 (pg_ctblmgr.control)](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/server/pg_ctblmgr.control)
- [官方扩展 SQL (pg_ctblmgr--0.1.sql)](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/server/sql/pg_ctblmgr--0.1.sql)

`pg_ctblmgr` — 基于逻辑复制的异步预取化视图，适用于 PostgreSQL。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_ctblmgr;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
