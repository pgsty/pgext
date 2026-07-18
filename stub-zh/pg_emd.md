## 用法

来源：

- [官方扩展控制文件](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/pg_emd.control)
- [官方上游文档](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/README.md)
- [官方 Rust 包清单](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/Cargo.toml)

`pg_emd` — 使用动态树嵌入快速近似计算推土机距离（Earth Mover's Distance）

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_emd";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_emd';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
