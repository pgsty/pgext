## 用法

来源：

- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/pg_swarm.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pg_swarm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_swarm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/Cargo.toml)
- [extensions/pg_swarm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_swarm/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/pgbrew.toml)
- [extensions/pg_swarm/src/node.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/node.rs)
- [extensions/pg_swarm/src/scheduler.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_swarm/src/scheduler.rs)

`pg_swarm` 0.3.2 在 `pgswarm` 中登记 SQL 执行函数，将作业拆分为后台工作进程管理的任务，并提供状态、重试与可选结果表。

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

扩展版本 0.3.2 通过 `pg_swarm.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.3.2 新增 `pgswarm.worker_status()` 与仅超级用户可调用的 `pgswarm.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_swarm.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_swarm UPDATE;
SELECT * FROM pgswarm.worker_status();
```
