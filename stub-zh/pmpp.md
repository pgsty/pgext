## 用法

来源：

- [官方 PGXN 发行页](https://pgxn.org/dist/pmpp/)
- [官方 README](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/README.md)
- [官方函数文档](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/doc/pmpp.md)
- [官方扩展控制文件](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/pmpp.control)

`pmpp` 1.2.3（Poor Man’s Parallel Processing）通过多个异步 libpq 连接运行手工拆分的 SQL 语句，并把兼容的结果行流式返回到一个 PostgreSQL 查询中。它适合由用户自行把工作拆成独立语句的场景；它不是由优化器驱动的并行查询引擎，也不是分布式事务管理器。

### 核心流程

`pmpp.distribute` 需要一个本地复合类型来描述所有返回行。临时表可以只提供这种类型，而无需存储结果。当目标服务器没有安装 `pmpp` 时，应显式传入 worker 数量。

```sql
CREATE EXTENSION pmpp;

CREATE TEMPORARY TABLE pmpp_result(value integer);

SELECT *
FROM pmpp.distribute(
    NULL::pmpp_result,
    'dbname=' || current_database(),
    ARRAY['SELECT 1', 'SELECT 2'],
    1.0,
    2
);
```

安装会创建角色 `pmpp`。只应把该角色授予允许打开所配置远程连接并执行提交 SQL 的用户。

### 主要接口

- `pmpp.distribute(row_type, connection, sql_list, cpu_multiplier, num_workers, setup_commands, result_format)` 处理单个目标，并以提供的行类型返回 `SETOF`。
- `pmpp.distribute` 的重载接受 `query_manifest[]`、`json` 或 `jsonb` 来描述多个目标，并提供类型化及 `record` 返回变体。
- `pmpp.meta` 并行运行 DDL 等命令，并报告每条命令的结果。
- `pmpp.broadcast` 在连接列表上运行一个返回行的查询。
- `pmpp.num_cpus()` 报告自动确定 worker 数量时使用的服务器 CPU 数。

连接参数可以是 libpq 连接字符串或 foreign server 名称。使用 `postgres_fdw` server 与 user mapping 可以避免在每次函数调用中直接写凭据，但凭据仍保存在 PostgreSQL 元数据中。

### 并发、事务与权限

每个 worker 都是独立的数据库会话，使用远程用户的权限，也看不到调用方事务中未提交的修改。结果行顺序取决于 worker 完成顺序；除非外层查询排序，否则不具确定性。

某个子查询失败时，`pmpp` 会尝试取消其他查询，但独立会话不提供原子性的分布式提交或回滚。不要把独立 DDL/DML 任务当成同一个事务；必须显式设计幂等性与恢复。`setup_commands` 和 `pmpp.meta` 可以在 worker 会话执行任意 SQL，因此应限制 `pmpp` 角色、server mapping 和可接受的查询文本。

worker 数会直接转化为额外连接，可能压垮本地或远程服务器。应谨慎设置 `num_workers`，考虑连接上限及每个查询的资源占用，并对完整工作负载进行基准测试。上游项目面向 PostgreSQL 9.4 时代的特性，最后更新于 2018 年，因此必须在准确的目标版本上验证构建和运行行为。
