## 用法

来源：

- [官方扩展控制文件](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/coordconv.control)
- [官方上游文档](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/README)

`coordconv` — 在 IAU 1958 银道坐标与 J2000 FK5 赤道坐标之间进行转换

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "coordconv";
SELECT extversion
FROM pg_extension
WHERE extname = 'coordconv';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
