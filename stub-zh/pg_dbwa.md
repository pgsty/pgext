## 用法

来源：

- [官方仓库 README](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/README.md)
- [官方扩展控制文件](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/pg_dbwa.control)
- [官方 0.3.1 版本 SQL](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/sql/pg_dbwa--0.3.1.sql)

`pg_dbwa` 定期保存 `pg_stat_statements` 快照，并为本地或远程 PostgreSQL 数据库报告工作负载增量。它是一个较早期的纯 SQL 工作负载分析器，使用前必须仔细核对模式、依赖、兼容性与安全边界。

### 核心流程

0.3.1 版硬依赖 `dblink`、`pg_stat_statements`、`pg_eyes` 与 `pg_prttn_tools`，并在 `dbwa` 模式中创建对象：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION dblink;
CREATE EXTENSION pg_eyes;
CREATE EXTENSION pg_prttn_tools;
CREATE EXTENSION pg_dbwa;

SELECT dbwa.get_stat('local');
SELECT *
FROM dbwa.show_top_queryes(
    'local',
    timestamp '2026-07-22 00:00:00',
    timestamp '2026-07-22 01:00:00',
    20
);
```

`dbwa.cluster_config` 记录受监控集群与连接信息，`dbwa.cluster_operation` 选择快照操作。`dbwa.get_stat(clustername)` 执行已启用操作并记录状态。`dbwa.show_top_queryes()` 对时间窗口内的增量排序，`dbwa.show_stat_query()` 则返回某个 `(userid, dbid, queryid)` 的区间增量。

### 安全与兼容性

已复核 SQL 将远端用户名和密码存入明文列，并据此构造 `dblink` 连接串。应限制全部 `dbwa` 表与函数、保护备份，并使用最小权限监控角色。若干函数为 `SECURITY DEFINER`，会把配置的操作名拼入动态 SQL，而且归 `postgres` 所有；不可信用户不得修改其配置或搜索路径。

0.3.1 安装脚本使用 `WITH (OIDS=FALSE)` 等旧语法和较早的 `pg_stat_statements` 列。必须针对精确的 PostgreSQL 与依赖版本验证安装，不能假定兼容现代服务器。快照表会持续增长，应安排保留策略并评估监控查询负载。
