## 用法

来源：

- [官方上游文档](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/readme.md)
- [官方 Rust 包清单](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/Cargo.toml)

`sixql` — 实验性 pgrx 扩展，尝试通过 PostgreSQL 服务端输出在兼容终端中渲染 SIXEL 图像。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "sixql";
SELECT extversion
FROM pg_extension
WHERE extname = 'sixql';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
