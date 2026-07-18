## 用法

来源：

- [官方扩展控制文件](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg.control)
- [官方上游 README](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/README.md)

`pg_agg` — 实验性的聚合查询支持扩展；当前安装脚本会创建非受信任的 PL/Python 语言。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`。
整理后的兼容版本集合为 `12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_agg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_agg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
