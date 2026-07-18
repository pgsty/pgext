## 用法

来源：

- [官方扩展控制文件](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.control)
- [官方上游文档](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/README.md)

`recommend` — 实验性 C 扩展，提供标量距离、文本相似度与一维数组相似度函数，代码派生自 smlar。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "recommend";
SELECT extversion
FROM pg_extension
WHERE extname = 'recommend';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
