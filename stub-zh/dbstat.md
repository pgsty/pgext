## 用法

来源：

- [官方扩展控制文件](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/dbstat.control)
- [官方上游文档](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/README.md)

`dbstat` — 通过 SQL 对象和配套的 libpq 进程收集表的精确行数及增删改时间。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "dbstat";
SELECT extversion
FROM pg_extension
WHERE extname = 'dbstat';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
