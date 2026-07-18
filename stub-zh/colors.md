## 用法

来源：

- [官方上游文档](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/colors/)

`colors` — 用于计算 CIE L*a*b* 色彩空间感知色差的 C 函数

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "colors";
SELECT extversion
FROM pg_extension
WHERE extname = 'colors';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
