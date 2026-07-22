## 用法

来源：

- [官方 README](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/README.md)
- [C 实现](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.c)
- [回归测试场景](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/sql/pg_sampletolog.sql)
- [扩展控制文件](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.control)

`pg_sampletolog` 2.0.0 是一个面向 PostgreSQL 9.4–11 的无 SQL 对象模块，通过执行器与工具语句钩子把语句或完整事务抽样写入服务器日志。PostgreSQL 12 已加入核心语句抽样设置，因此该模块主要适用于遗留服务器或历史回放实验。它没有 `CREATE EXTENSION` SQL 接口。

### 加载与配置

可以为单个会话加载，也可以在每次连接开始时加载：

```conf
session_preload_libraries = 'pg_sampletolog'

pg_sampletolog.statement_sample_rate = 0.10
pg_sampletolog.transaction_sample_rate = 0
pg_sampletolog.statement_sample_limit = '100ms'
pg_sampletolog.log_statement = 'ddl'
pg_sampletolog.log_before_execution = off
pg_sampletolog.log_level = 'log'
```

诊断会话可使用：

```sql
LOAD 'pg_sampletolog';
SET pg_sampletolog.statement_sample_rate = 0.10;
SET pg_sampletolog.statement_sample_limit = '100ms';
```

两个抽样率默认都为零，因此配置前模块处于禁用状态。语句抽样会对每条语句独立选择；事务抽样在事务开始时选择一次，被选中后记录该事务的所有语句。正的 `statement_sample_limit` 会强制记录超过时长阈值的语句，但至少一个抽样率必须为正。

### 重要设置

- `pg_sampletolog.statement_sample_rate`：记录单条语句的比例。
- `pg_sampletolog.transaction_sample_rate`：记录完整语句集合的事务比例。
- `pg_sampletolog.statement_sample_limit`：在已启用抽样时，始终记录超过时长阈值的执行。
- `pg_sampletolog.log_statement`：取 `none`、`ddl` 或 `mod`，对应历史核心 `log_statement` 分类。
- `pg_sampletolog.log_before_execution`：为回放工具在执行前输出，而不是执行后附带耗时。
- `pg_sampletolog.disable_log_duration`：省略耗时，主要用于测试。
- `pg_sampletolog.log_level`：选择抽样记录使用的 PostgreSQL 消息严重级别。

若已加载 `pg_stat_statements`，模块可以在日志中加入查询标识。对于内部包含多个查询的 SQL 函数，它只记录外层函数调用。不支持 PREPARE；多语句查询字符串还可能在逐语句日志中被重复显示，造成误导。

### 日志与兼容性边界

抽样 SQL 可能包含密码、令牌、个人数据与字面值。应采用与完整语句日志相同的文件权限、传输加密、保留、脱敏和访问控制。抽样是概率行为，不是审计轨迹；若 `ddl` 或 `mod` 模式要求完整性，仍需另行验证。执行前记录会包含后来失败或回滚的语句，执行后记录则可能在后端崩溃时漏掉语句。

这些钩子绑定 PostgreSQL 9.4–11 内部接口，不能由目录中的现代版本矩阵推断支持。在 PostgreSQL 12+ 上应优先使用核心 `log_statement_sample_rate`、`log_transaction_sample_rate`、`log_min_duration_sample` 及相关当前设置。在遗留环境使用前，应测试钩子共存、备库行为、预备与多语句流量、嵌套查询、错误路径、日志量、延迟、随机分布、轮换和回放兼容性。
