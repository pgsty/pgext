## 用法

来源：

- [官方扩展控制文件](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/vmagentfdw.control)
- [官方 Rust 包清单](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/Cargo.toml)

`vmagentfdw` — 通过 vmagent HTTP API 读取活跃和已移除抓取目标状态的外部数据包装器。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "vmagentfdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'vmagentfdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
