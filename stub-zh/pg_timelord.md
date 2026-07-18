## 用法

来源：

- [官方扩展控制文件](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.control)
- [官方上游文档](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/README.md)

`pg_timelord` — 利用提交时间戳与自定义 MVCC 快照进行时间旅行查询的概念验证扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_timelord";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timelord';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
