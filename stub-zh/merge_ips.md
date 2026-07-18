## 用法

来源：

- [官方扩展控制文件](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips.control)
- [官方上游文档](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/README)

`merge_ips` — 合并 IP 地址和相邻网段，同时避免结果子网包含间隙。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "merge_ips";
SELECT extversion
FROM pg_extension
WHERE extname = 'merge_ips';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
