## 用法

来源：

- [官方扩展控制文件](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer.control)
- [官方上游文档](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/README.md)

`gdpr_anatomizer` — gdpr_anatomizer：基于 Anatomy 方法的 GDPR 匿名化原型扩展，主要使用 PL/Python 并包含 PL/pgSQL 辅助函数；上游未声明许可证。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`。
整理后的兼容版本集合为 `12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "gdpr_anatomizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'gdpr_anatomizer';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
