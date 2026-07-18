## 用法

来源：

- [官方扩展控制文件](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/cidr.control)

`cidr` — 将 IPv4 或 IPv6 地址范围转换为最小 CIDR 网段集合

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "cidr";
SELECT extversion
FROM pg_extension
WHERE extname = 'cidr';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
