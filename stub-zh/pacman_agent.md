## 用法

来源：

- [官方上游 README](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/README.md)
- [官方扩展控制文件 (pacman_agent.control)](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/pacman_agent.control)
- [官方扩展 SQL (pacman_agent--0.1.0.sql)](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/sql/pacman_agent--0.1.0.sql)

`pacman_agent` — 该目录包含初始的 PACMAN PostgreSQL 背景工作者扩展框架。在管理或自动化上述数据库行为时使用它。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pacman_agent;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
