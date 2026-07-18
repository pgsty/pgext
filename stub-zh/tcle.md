## 用法

来源：

- [官方上游文档](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/README.md)

`tcle` — 通过表访问方法透明加密指定列类型的实验性 AES-256-CBC 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tcle";
SELECT extversion
FROM pg_extension
WHERE extname = 'tcle';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
