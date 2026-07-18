## 用法

来源：

- [官方上游文档](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/README.md)

`pgafis` — 为 PostgreSQL 提供自动指纹识别系统（AFIS）支持的扩展。

已复核目录快照记录的版本为 `1.2`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgafis";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgafis';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
