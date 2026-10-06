## 用法

来源：

- [extensions/pg_delta/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/Cargo.toml)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pgbrew.toml)
- [extensions/pg_delta/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/README.md)
- [extensions/pg_delta/pg_delta.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pg_delta.control)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/src/lib.rs)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pgbrew.toml)
- [docs/pg_delta.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/docs/pg_delta.md)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)

`pg_delta` 0.3.3 通过读取、导出和受管理的流将 PostgreSQL 与 Delta Lake 集成。新增索引模式登记 Delta 事务日志并裁剪文件，以 FDW 原地查询。

### 核心用法

```ini
shared_preload_libraries = 'pg_delta'
```

```sql
CREATE EXTENSION pg_delta;
SELECT * FROM delta.list_tables();
SELECT delta.status();
```

### 运行边界

安装需要超级用户。流管理器要求预加载 `pg_delta` 并重启，应明确配置工作数据库和存储凭据。`delta` 模式提供表／流创建、刷新、状态、历史与导出接口，索引模式使用 `pg_delta_server`。云路径和 SQL 定义属于特权操作。手册将基于逻辑复制的 CDC 导出标为未实现；轮询和快照模式各有更新／删除及恢复语义。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.3.3 通过 `delta.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.3.3 新增 `delta.worker_status()` 与仅超级用户可调用的 `delta.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`delta.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_delta UPDATE;
SELECT * FROM delta.worker_status();
```
