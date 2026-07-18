## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra.control)

`jsonb_extra` — 提供 jsonb 提取、文本转换和更新辅助函数的 C 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "jsonb_extra";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_extra';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
