## 用法

来源：

- [官方 README](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/README.md)
- [安装的 SQL 对象](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/init.sql)
- [共享内存与 GUC 实现](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/pg_querylog.c)

`pg_querylog` 在共享内存中捕获每个后端当前或最近完成的查询及渲染后的参数。它只适合在经过明确测试的旧 PostgreSQL 版本上做短期诊断；捕获的文本可能包含凭据、个人数据与应用密钥。

### 加载与查询

将库加入 `shared_preload_libraries` 并重启 PostgreSQL，然后在受限模式中安装 SQL 对象：

```conf
shared_preload_libraries = 'pg_querylog'
pg_querylog.buffer_size = '16kB'
```

```sql
CREATE SCHEMA querylog;
REVOKE ALL ON SCHEMA querylog FROM PUBLIC;
CREATE EXTENSION pg_querylog SCHEMA querylog;

SET pg_querylog.enabled = on;
SELECT * FROM querylog.get_queries(false, true);
SELECT * FROM querylog.running_queries;
```

`get_queries(only_running, skip_overflow)` 返回 `pid`、`query`、`params`、`start_time`、`end_time`、`running` 与 `overflow`。`running_queries` 只保留活动条目。溢出标志表示每后端缓冲区无法保留完整条目。

### 配置与安全

`pg_querylog.buffer_size` 是最小值为 10 kB 的超级用户设置，并按后端分配；容量应按该值乘以 `max_connections` 再加额外开销估算。`pg_querylog.enabled` 也只允许超级用户修改。README 称其默认关闭，但已复核的 C 源码将初始值设为 true；应显式设置，并用 `SHOW pg_querylog.enabled` 核验。

README 也允许使用 `session_preload_libraries`，但警告每个新后端都会重写共享的启用值。为获得确定行为，应优先使用集群预加载。限制模式以及函数/视图权限，定义较短的诊断窗口，完成后关闭采集。该未维护代码挂接执行器内部接口且只列出旧 PostgreSQL 版本，因此必须在准确目标构建上测试崩溃、预备语句、参数渲染、并发后端、截断与内存用量。
