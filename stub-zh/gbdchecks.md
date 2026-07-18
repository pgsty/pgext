## 用法

来源：

- [官方扩展控制文件](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/gbdchecks.control)
- [官方上游文档](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/README.md)

`gbdchecks` — 提供膨胀、锁、查询统计和权限检查函数及视图的数据库巡检扩展。

已复核目录快照记录的版本为 `1.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pg_stat_statements`。
整理后的兼容版本集合为 `11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "gbdchecks";
SELECT extversion
FROM pg_extension
WHERE extname = 'gbdchecks';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
