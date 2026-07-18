## 用法

来源：

- [官方上游文档](https://pgxn.org/dist/test_factory/doc/test_factory.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/test_factory/)
- [官方上游 README](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/README.md)

`test_factory` — 用于注册、创建、缓存并复用命名 PostgreSQL 测试数据集的框架。

已复核目录快照记录的版本为 `0.5.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "test_factory";
SELECT extversion
FROM pg_extension
WHERE extname = 'test_factory';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
