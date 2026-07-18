## 用法

来源：

- [官方项目或服务商页面](https://pldotnet.brickabode.com)

`pldotnet` — 通过 .NET 为 PostgreSQL 提供 C# 与 F# 过程语言。

已复核目录快照记录的版本为 `0.9`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pldotnet";
SELECT extversion
FROM pg_extension
WHERE extname = 'pldotnet';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
