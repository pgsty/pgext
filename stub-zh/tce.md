## 用法

来源：

- [官方扩展控制文件](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/tce.control)
- [官方上游文档](https://the-cryptic-elephant.readthedocs.io/en/latest/)
- [官方 Rust 包清单](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/Cargo.toml)

`tce` — 使用按角色信封密钥和外部密钥管理服务为 PostgreSQL 提供透明列加密。

已复核目录快照记录的版本为 `0.10.0-dev`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tce";
SELECT extversion
FROM pg_extension
WHERE extname = 'tce';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
