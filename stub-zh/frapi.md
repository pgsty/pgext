## 用法

来源：

- [官方扩展控制文件](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi.control)
- [官方上游文档](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/README.md)

`frapi` — 为法国公共管理数据 API 提供数据库内封装函数。

已复核目录快照记录的版本为 `0.1.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plsh`, `postgis`。

```sql
CREATE EXTENSION "frapi";
SELECT extversion
FROM pg_extension
WHERE extname = 'frapi';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
