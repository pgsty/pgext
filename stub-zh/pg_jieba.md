## 用法

来源：

- [官方上游文档](https://github.com/jaiminpan/pg_jieba/blob/d0ffac8028328b2566a889ff4db3d74ba63d1b42/README.md)

`pg_jieba` — 使用 Jieba 分词的 PostgreSQL 中文全文检索解析器扩展

已复核目录快照记录的版本为 `1.1.1`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_jieba";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_jieba';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
