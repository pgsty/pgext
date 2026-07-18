## 用法

来源：

- [官方扩展控制文件](https://github.com/blkchain/pg_blkchain/blob/master/pg_blkchain.control)
- [官方上游文档](https://github.com/blkchain/pg_blkchain/blob/master/README.md)

`pg_blkchain` — 为 PostgreSQL 提供 Bitcoin 区块链解析和辅助函数的 C 扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_blkchain";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_blkchain';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
