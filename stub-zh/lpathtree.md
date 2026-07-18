## 用法

来源：

- [官方扩展控制文件](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/lpathtree.control)

`lpathtree` — 使用斜杠分隔标签的层级路径类型，源自 PostgreSQL ltree。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "lpathtree";
SELECT extversion
FROM pg_extension
WHERE extname = 'lpathtree';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
