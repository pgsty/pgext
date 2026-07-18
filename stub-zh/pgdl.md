## 用法

来源：

- [官方扩展控制文件](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/pgdl.control)

`pgdl` — 用于在 PostgreSQL 数据库内执行深度学习模型推理和向量存储的 C++ 扩展。

已复核目录快照记录的版本为 `1.3.0`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgdl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgdl';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
