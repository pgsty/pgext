## 用法

来源：

- [官方扩展控制文件](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/graph_accel.control)
- [官方上游文档](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/README.md)
- [官方 Rust 包清单](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/Cargo.toml)

`graph_accel` — graph_accel：面向 Apache AGE 的内存图遍历加速扩展。

已复核目录快照记录的版本为 `0.5.0`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`age`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "graph_accel";
SELECT extversion
FROM pg_extension
WHERE extname = 'graph_accel';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
