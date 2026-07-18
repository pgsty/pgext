## 用法

来源：

- [官方扩展控制文件](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/italian_codes.control)
- [官方上游文档](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/README.rst)
- [官方 PGXN 分发页](https://pgxn.org/dist/italian_codes/)

`italian_codes` — 提供意大利税号与增值税号的校验域。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "italian_codes";
SELECT extversion
FROM pg_extension
WHERE extname = 'italian_codes';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
