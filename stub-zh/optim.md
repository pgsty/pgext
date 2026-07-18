## 用法

来源：

- [官方扩展控制文件](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/optim.control)

`optim` — 实验性 pgrx 扩展，定义优化方向与超参数试验数据类型。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "optim";
SELECT extversion
FROM pg_extension
WHERE extname = 'optim';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
