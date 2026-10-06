## 用法

来源：

- [extensions/pg_ortools/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/Cargo.toml)
- [extensions/pg_ortools/src/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/src/worker.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_ortools/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pgbrew.toml)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_ortools/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pgbrew.toml)
- [官方 pg_ortools README](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/README.md)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pg_ortools.control)
- [SQL API 实现](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/src/lib.rs)

`pg_ortools` 版本 `0.3.2` 用 SQL 定义混合整数与线性优化问题，并通过 HiGHS 求解。它适合有界指派、资源分配、可行性及目标优化工作负载，这些问题通常可表示为整数或布尔变量和线性约束。

### 核心流程

```sql
CREATE EXTENSION pg_ortools;

SELECT pgortools.create_problem('example');
SELECT pgortools.add_int_var('example', 'x', 0, 100);
SELECT pgortools.add_int_var('example', 'y', 0, 100);
SELECT pgortools.add_constraint('example', 'x + y <= 150');
SELECT pgortools.maximize('example', '2*x + 3*y');

SELECT pgortools.solve_sync('example');
SELECT pgortools.get_solution('example');
```

对于异步任务，`solve` 返回作业 ID；使用 `solve_status`、`cancel_solve`、`LISTEN pgortools_solve` 和 `get_solution` 管理作业。`solve_greedy` 返回第一个可行解，而 `solve_local`、`solve_with_strategy` 与 `solve_auto` 提供其他搜索策略。声明式函数 `solve_assignment` 和 `parse_assignment` 可构建常见的表驱动指派模型。

### 运维说明

扩展在 `pgortools` 中保存问题、变量、约束、作业和解。约束表达式支持线性比较，但官方 README 明确不支持 `!=`。默认求解时限由 `pg_ortools.solver_time_limit` 控制；上游把 worker 启用、轮询间隔和 worker 数据库描述为对重启敏感的服务器设置。依赖异步作业前应确认后台工作进程已经运行。control 文件不可重定位但不限定超级用户安装，安装 SQL 还向 public 授予其模式表的 DML 权限；多租户使用前应审查权限。

### 0.3.0 版本边界

这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

异步任务需要预加载 `pg_ortools` 并重启。默认构建现在还包含基于 Pumpkin 的 CP-SAT 调度引擎，`solve_cp` 支持时间限制和取消。HiGHS 及其 SQL 工作流仍是独立求解路径。新增的 `eidos_catalog_*` 接口描述实时优化目录。

### 当前版本与升级

扩展版本 0.3.2 通过 `pg_ortools.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 `pg_ortools.solver_database` 保留为已弃用的别名。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.3.2 新增 `pgortools.worker_status()` 与仅超级用户可调用的 `pgortools.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_ortools.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_ortools UPDATE;
SELECT * FROM pgortools.worker_status();
```
