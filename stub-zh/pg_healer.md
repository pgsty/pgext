## 用法

来源：

- [官方扩展控制文件](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.control)
- [官方上游 README](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/README.md)

`pg_healer` — 自动修复 PostgreSQL 数据损坏的概念验证扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plperlu`。

```sql
CREATE EXTENSION "pg_healer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_healer';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
