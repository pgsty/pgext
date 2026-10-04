## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/pg_swarm.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/pg_swarm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/src/node.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/src/node.rs)
- [extensions/pg_swarm/src/scheduler.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_swarm/src/scheduler.rs)

`pg_swarm` 0.3.1 在 `pgswarm` 中登记 SQL 执行函数，将作业拆分为后台工作进程管理的任务，并提供状态、重试与可选结果表。

### 核心用法

```ini
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

控制文件要求超级用户安装。需要预加载并重启；节点管理器和调度器使用 `pg_swarm.database`（默认 postgres），应在该库创建扩展及执行函数。`register_executor` 登记接收任务 ID、JSONB 负载、分片索引和分片数的函数，`submit_job` 提交任务。执行函数运行在特权服务中，应限制登记与提交权限。`pg_swarm.workers`、超时和重试设置控制调度。任务重试时，外部副作用应支持幂等处理，不能据此推断跨节点恰好一次保证。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.3.1 通过 `pg_swarm.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.4.1；上游二进制不代表 Pigsty 软件包可用性。
