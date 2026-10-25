## 用法

来源：

- [官方资料](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/README.md)
- [官方资料](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/benchmark/README.md)
- [官方资料](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/src/pg_jevplanner.c)
- [官方资料](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/sql/pg_jevplanner--0.1.0.sql)

`pg_jevplanner` 0.1.0 是 PostgreSQL 18 的实验性连接排序扩展。它请求外部 TypeSafe JEV 服务选择合法的内连接合并顺序，物理计划仍由 PostgreSQL 生成。源代码会拒绝其他大版本的服务器头文件。

### 启用与对比

启动服务器前，在服务器进程环境中配置有效的 `TYPESAFE_API_KEY`。由超级用户安装，并在每个后端中加载动态库；创建扩展只加载当前会话。管理员配置会话预加载后，可覆盖后续会话。

```sql
CREATE EXTENSION pg_jevplanner;
LOAD 'pg_jevplanner';
SELECT jevplanner_version();
SET jev.enabled = on;
SET max_parallel_workers_per_gather = 0;
EXPLAIN SELECT a.id FROM a JOIN b ON a.id=b.id JOIN c ON b.id=c.id;
SET jev.enabled = off;
```

示例需要已经存在且可以连接的表。应与禁用扩展时的实际计划和结果比较；单独测量执行耗时不包含规划和外部 API 延迟。这是显式启用的研究型规划器，不支持的 SQL 语义边界仍会保留。

### 配置与数据外发

规划器控制参数需要超级用户权限。`jev.timeout_ms` 默认将每次请求限制为 2000 毫秒，`jev.planning_timeout_ms` 默认将整个规划过程限制为 30000 毫秒。`jev.max_decisions` 和 `jev.max_candidates` 限制决策工作量；超出候选数量上限时会报错，不会静默筛选候选。

请求会将 SQL、名称、字面量、关系统计和连接元数据发送给外部服务。`jev.log_http` 默认关闭，开启后会在通知中展示请求与响应正文。应避免在 SQL 和日志中包含生产凭据，将密钥限制在服务器环境，并在启用钩子前评估外部服务的数据访问边界。准备语句重新规划时可能再次请求模型。
