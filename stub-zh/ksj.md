## 用法

来源：

- [官方扩展控制文件](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.control)

`ksj` — 日语汉字数字整数类型，支持算术、聚合、类型转换、比较和 B-tree 索引。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "ksj";
SELECT extversion
FROM pg_extension
WHERE extname = 'ksj';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
