## 用法

来源：

- [官方扩展控制文件](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/nextid.control)
- [官方上游文档](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/README.md)

`nextid` — nextid：为 PostgreSQL 提供基于 Instagram ID 方案的 ID 生成功能。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "nextid";
SELECT extversion
FROM pg_extension
WHERE extname = 'nextid';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
