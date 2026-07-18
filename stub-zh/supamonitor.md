## 用法

来源：

- [官方扩展控制文件](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/supamonitor.control)
- [官方上游文档](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/src/lib.rs)
- [官方 Rust 包清单](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/Cargo.toml)

`supamonitor` — 通过预加载 pgrx 查询钩子，将查询文本、查询 ID、规划与执行耗时记录为 JSON 日志。

已复核目录快照记录的版本为 `0.0.6`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "supamonitor";
SELECT extversion
FROM pg_extension
WHERE extname = 'supamonitor';
```

该上游项目与 `Supabase` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
