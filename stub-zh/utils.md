## 用法

来源：

- [官方上游 README](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/README.md)

`utils` — 提供整数加法与乘法函数的极简 C 扩展示例。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'utils';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
