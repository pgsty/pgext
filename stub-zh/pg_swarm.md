## 用法

来源：

- [extensions/pg_swarm/pg_swarm.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/pg_swarm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/src/node.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/node.rs)
- [extensions/pg_swarm/src/scheduler.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_swarm/src/scheduler.rs)

`pg_swarm` 0.3.0 在 `pgswarm` 中登记 SQL 执行函数，将作业拆分为后台工作进程管理的任务，并提供状态、重试与可选结果表。

### 核心用法

```conf
shared_preload_libraries = 'pg_swarm'
```

```sql
CREATE EXTENSION pg_swarm;
CREATE FUNCTION swarm_example(bigint, jsonb, integer, integer)
RETURNS jsonb LANGUAGE SQL AS $$ SELECT $2; $$;
SELECT pgswarm.register_executor('example', 'public.swarm_example');
SELECT pgswarm.submit_job('example', '{"message":"hello"}'::jsonb);
SELECT * FROM pgswarm.list_executors();
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启；所核对的节点管理器和调度器连接 postgres 数据库，应在其中创建扩展及执行函数。`register_executor` 登记接收任务 ID、JSONB 负载、分片索引和分片数的函数，`submit_job` 提交任务。执行函数运行在特权服务中，应限制登记与提交权限。`pg_swarm.workers`、超时和重试设置控制调度。任务重试时，外部副作用应支持幂等处理，不能据此推断跨节点恰好一次保证。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
