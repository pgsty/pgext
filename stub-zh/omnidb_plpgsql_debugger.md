## 用法

来源：

- [官方扩展控制文件](https://github.com/OmniDB/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/omnidb_plpgsql_debugger.control)
- [官方上游文档](https://github.com/OmniDB/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/README.md)

`omnidb_plpgsql_debugger` — 用于在 OmniDB 中启用 PL/pgSQL 调试器的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "omnidb_plpgsql_debugger";
SELECT extversion
FROM pg_extension
WHERE extname = 'omnidb_plpgsql_debugger';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
