## 用法

来源：

- [官方上游文档](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/doc/validadores.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/validadores/)

`validadores` — 巴西 CPF 校验位验证函数、运算符与域类型

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "validadores";
SELECT extversion
FROM pg_extension
WHERE extname = 'validadores';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
