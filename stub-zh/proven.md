## 用法

来源：

- [官方扩展控制文件](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven.control)

`proven` — proven：将 libproven 形式化验证安全函数暴露给 PostgreSQL 的 C 语言绑定

已复核目录快照记录的版本为 `0.5.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "proven";
SELECT extversion
FROM pg_extension
WHERE extname = 'proven';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
