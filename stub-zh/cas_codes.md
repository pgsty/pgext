## 用法

来源：

- [官方扩展控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/f600a9def937ff5614fbecac805494777d4e69d8/web/web_common/cas_codes/cas_codes.control)

`cas_codes` — PostgreSQL 的 CAS 登记号数据类型，提供格式校验和比较支持。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cas_codes";
SELECT extversion
FROM pg_extension
WHERE extname = 'cas_codes';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
