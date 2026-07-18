## 用法

来源：

- [官方扩展控制文件](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/pg_ta.control)
- [官方上游文档](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/README.md)

`pg_ta` — pg_ta：将 TA-Lib 技术分析函数封装为 PostgreSQL 函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_ta";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ta';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
