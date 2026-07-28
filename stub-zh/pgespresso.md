## 用法

来源：

- [官方扩展控制文件 (pgespresso.control)](https://api.pgxn.org/src/pgespresso/pgespresso-1.2.0/pgespresso.control)
- [官方扩展 SQL (pgespresso--1.2.sql)](https://api.pgxn.org/src/pgespresso/pgespresso-1.2.0/pgespresso--1.2.sql)

`pgespresso` — 可选扩展用于 Barman，PostgreSQL 的备份和恢复管理器。在管理或自动化上述数据库行为时使用此扩展。请使用链接中的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgespresso;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pgespresso_abort_backup()` 是一个扩展函数，返回 `VOID`。
- `pgespresso_start_backup(label TEXT, fast BOOL)` 是一个扩展函数，返回 `TEXT`。
- `pgespresso_stop_backup(label_content TEXT)` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.2`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
