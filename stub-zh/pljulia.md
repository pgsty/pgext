## 用法

来源：

- [官方上游来源](https://gitlab.com/pljulia/pljulia)

`pljulia` — 将 Julia 嵌入为非受信 PostgreSQL 过程语言，支持 SPI、DO 代码块、触发器、数组与复合类型转换。

已复核目录快照记录的版本为 `0.8`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pljulia";
SELECT extversion
FROM pg_extension
WHERE extname = 'pljulia';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
