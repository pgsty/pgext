## 用法

来源：

- [官方扩展控制文件](https://github.com/funbringer/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/pg_linegazer.control)
- [官方上游文档](https://github.com/funbringer/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/README.md)

`pg_linegazer` — 为 PL/pgSQL 提供透明的代码覆盖率统计。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pg_linegazer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_linegazer';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
