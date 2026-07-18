## 用法

来源：

- [官方扩展控制文件](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs.control)
- [官方上游文档](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/README.md)

`ludia_funcs` — ludia_funcs：全文搜索相关 PostgreSQL 扩展；提供 Ludia 辅助函数。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ludia_funcs";
SELECT extversion
FROM pg_extension
WHERE extname = 'ludia_funcs';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
