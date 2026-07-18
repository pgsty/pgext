## 用法

来源：

- [官方扩展控制文件](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml.control)
- [官方上游文档](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/README.md)

`postgresml` — 为实验性朴素贝叶斯文本分类示例提供 tsvector 词元及位置计数函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "postgresml";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgresml';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
