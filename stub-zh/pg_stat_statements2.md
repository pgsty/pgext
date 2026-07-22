## 用法

来源：

- [官方 README](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/README)
- [官方扩展 SQL](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2--1.2.sql)
- [官方 C 实现](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2.c)

`pg_stat_statements2` 是 PostgreSQL 9.4.4 时代标准语句统计扩展的分支，经过修改以复用 `sql_firewall` 的查询指纹。它服务于这一历史集成场景，不是 PostgreSQL 内置模块的现代直接替代品。

### 前置条件与启动

两个库必须按以下顺序预加载，修改该配置后必须重启服务端：

```conf
shared_preload_libraries = 'sql_firewall,pg_stat_statements2'
```

重启后安装 SQL 对象。使用独立 schema 可减少与其他扩展的名称冲突：

```sql
CREATE SCHEMA firewall_stats;
CREATE EXTENSION pg_stat_statements2 WITH SCHEMA firewall_stats;

SELECT query, calls, total_time, rows
FROM firewall_stats.pg_stat_statements
ORDER BY total_time DESC
LIMIT 20;
```

### SQL 接口

- `pg_stat_statements(showtext boolean)` 返回 PostgreSQL 9.4 风格的计数器，包括调用次数、总耗时、行数、块活动与块 I/O 时间。
- `pg_stat_statements` 是该函数之上的视图，安装脚本将其权限授予 `PUBLIC`。
- `pg_stat_statements_reset()` 清除共享统计，其默认 `PUBLIC` 权限已被撤销。

该视图会向所有拥有查询权限的角色暴露捕获的查询文本。应检查并在必要时撤销对 `firewall_stats.pg_stat_statements` 的访问；重置函数则应视为具有破坏性的管理操作。

### 兼容性边界

该扩展刻意删除了自己的 `JumbleQuery` 实现，并调用 `sql_firewall` 导出的符号。因此库加载顺序以及 ABI 兼容的 firewall 构建都是强制要求；符号缺失或不匹配会阻止启动。

其 SQL 对象名称与内置语句统计扩展重叠。schema 隔离能限制 SQL 名称冲突，但不能证明两个共享库、hook、GUC 或 ABI 可以共存。应在隔离服务器上测试完整预加载组合。

输出结构早于当前 PostgreSQL 中的规划计数器及其他字段。上游只记录 PostgreSQL 9.4.4 与历史 `sql_firewall` 的组合，本次核验的仓库最后修改于 2015 年。除非经过有意识的移植、源码审查与负载测试，否则不要在现代服务端加载它。
