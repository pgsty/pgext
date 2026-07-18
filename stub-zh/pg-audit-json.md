## 用法

来源：

- [官方扩展控制文件](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/pg-audit-json.control)
- [官方上游文档](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/README.md)

`pg-audit-json` — 基于触发器的 PostgreSQL 表审计扩展，使用 JSONB 保存行快照并递归记录字段差异。

已复核目录快照记录的版本为 `1.0.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg-audit-json";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg-audit-json';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
