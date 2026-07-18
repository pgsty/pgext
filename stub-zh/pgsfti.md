## 用法

来源：

- [官方扩展控制文件](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/pgsfti.control)
- [官方上游文档](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/docs/source/intro.rst)
- [官方上游 README](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/README.md)

`pgsfti` — pgsfti 为 PostgreSQL 提供简单梯形模糊时间区间数据类型及模糊时态关系运算。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgsfti";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgsfti';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
