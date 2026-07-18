## 用法

来源：

- [官方扩展控制文件](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/sys_syn_dblink.control)

`sys_syn_dblink` — 基于 PL/pgSQL 的同步处理器：通过具名 dblink 连接拉取远端 sys_syn 队列，批量转换并写入本地表，再将处理状态回传源数据库。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`dblink`, `hstore`, `plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "sys_syn_dblink";
SELECT extversion
FROM pg_extension
WHERE extname = 'sys_syn_dblink';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
