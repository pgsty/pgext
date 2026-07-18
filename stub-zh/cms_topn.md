## 用法

来源：

- [官方上游文档](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/README.md)

`cms_topn` — 用于近似频率统计与 Top-N 查询的 Count-Min Sketch 数据类型及聚合函数

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "cms_topn";
SELECT extversion
FROM pg_extension
WHERE extname = 'cms_topn';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
