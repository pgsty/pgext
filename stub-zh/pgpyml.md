## 用法

来源：

- [官方上游文档](https://minoro.is-a.dev/pgpyml/get-started/)
- [官方 PGXN 分发页](https://pgxn.org/dist/pgpyml/)
- [官方项目或服务商页面](https://minoro.is-a.dev/pgpyml/)

`pgpyml` — pgpyml 通过 PL/Python 函数与触发器在 PostgreSQL 内运行 scikit-learn 机器学习模型。

已复核目录快照记录的版本为 `0.3.1`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`。

```sql
CREATE EXTENSION "pgpyml";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgpyml';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
