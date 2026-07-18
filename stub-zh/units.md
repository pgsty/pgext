## 用法

来源：

- [官方上游文档](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/README.md)

`units` — 提供单位换算表和 PL/pgSQL 换算函数，当前脚本内置长度单位换算数据。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "units";
SELECT extversion
FROM pg_extension
WHERE extname = 'units';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
