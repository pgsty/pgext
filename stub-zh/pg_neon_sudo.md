## 用法

来源：

- [官方扩展控制文件](https://github.com/skyzh/pg_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.control)

`pg_neon_sudo` — 提供 Neon 专用的可信 sudo 包装函数，以初始超级用户身份启动和停止 PostgreSQL Anonymizer 动态脱敏。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`anon`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_neon_sudo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_neon_sudo';
```

该上游项目与 `Neon` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
