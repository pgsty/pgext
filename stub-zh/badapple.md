## 用法

来源：

- [官方扩展控制文件](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.control)

`badapple` — 在 psql 客户端中播放 Bad Apple!! 动画。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "badapple";
SELECT extversion
FROM pg_extension
WHERE extname = 'badapple';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
