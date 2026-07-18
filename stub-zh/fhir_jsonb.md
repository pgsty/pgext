## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb.control)

`fhir_jsonb` — 用于 PostgreSQL 的 FHIR JSONB 合成数据生成与查询基准测试夹具。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "fhir_jsonb";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhir_jsonb';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
