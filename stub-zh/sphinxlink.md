## 用法

来源：

- [官方扩展控制文件](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink.control)
- [官方上游文档](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/README.md)

`sphinxlink` — 通过 MySQL 兼容协议从 PostgreSQL 向 SphinxSearch 或 ManticoreSearch 发起查询的 C 扩展。

已复核目录快照记录的版本为 `1.4`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "sphinxlink";
SELECT extversion
FROM pg_extension
WHERE extname = 'sphinxlink';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
