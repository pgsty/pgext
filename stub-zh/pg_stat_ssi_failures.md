## 用法

来源：

- [固定提交的上游 README](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/README.md)
- [固定提交的 0.1 版安装 SQL](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures--0.1.sql)
- [固定提交的 C 实现](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.c)
- [固定提交的扩展 control 文件](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.control)

`pg_stat_ssi_failures` `0.1` 维护一个集群级计数器，统计 SQLSTATE 为 `serialization_failure` 的错误。它用于向监控系统暴露可累计的可串行化事务冲突指标，不能指出失败的语句、数据库、角色或事务。

### 核心流程

模块必须在服务器启动时预留共享内存。先把它加入 `shared_preload_libraries` 并重启 PostgreSQL，然后创建 SQL 扩展：

```conf
shared_preload_libraries = 'pg_stat_ssi_failures'
```

```sql
CREATE EXTENSION pg_stat_ssi_failures;

SELECT pg_stat_ssi_failures();
SELECT pg_stat_ssi_failures_reset();
```

进行区间监控时，应在监控系统中保存上次观测值，并从当前累计值中减去它。当前值变小时应按 reset 或 restart 处理，而不是解释为负的失败率。

### 重要对象

- `pg_stat_ssi_failures() RETURNS bigint`：读取集群级共享计数器。
- `pg_stat_ssi_failures_reset() RETURNS void`：把整个实例所有数据库共享的计数器归零。

### 运维说明

该 hook 会统计每个带有 PostgreSQL serialization-failure SQLSTATE 的已发出错误，不提供按查询划分的子集，也没有标签。正常关闭会把计数器写入服务器端统计文件；源码会在启动后特意删除该文件，并排除备份与复制行为。因此，崩溃恢复、恢复备份、提升、reset 或移除扩展都会造成指标断点。

安装 SQL 没有为 `pg_stat_ssi_failures_reset()` 增加权限保护，应撤销应用角色的执行权限。已复核的 reset 实现在只持有 shared LWLock 时写共享状态，因此应避免并发 reset，并只允许管理员执行。该旧源码使用 PostgreSQL 内部 hook 和锁 API；生产监控前必须在准确的目标版本上验证编译与 hook 链，并确认其他预加载模块不会相互干扰。
