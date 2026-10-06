## 用法

来源：

- [extensions/pg_xarray/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/Cargo.toml)
- [extensions/pg_xarray/src/server/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/src/server/worker.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_xarray/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/pgbrew.toml)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_xarray/pg_xarray.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/pg_xarray.control)
- [extensions/pg_xarray/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/src/lib.rs)
- [extensions/pg_xarray/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/pgbrew.toml)
- [extensions/pg_xarray/demo/01_register_local.sql](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_xarray/demo/01_register_local.sql)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)

`pg_xarray` 0.4.2 在 `pgx` 中登记科学数据集、变量、数据块和网格，提供科学数组查询与可选的 WMS 服务。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_xarray;
SELECT pgx.register_dataset('weather', 'zarr');
SELECT * FROM pgx.list_datasets();
```

### 运行边界

依赖 `postgis`，需要超级用户安装。数据集／文件登记指向服务器可读取的本地或对象存储资源，应限制路径和凭据。SQL 接口包括 `register_dataset`、`register_file`、`register_variable`、`register_chunk` 和 `list_datasets`。WMS 需要预加载、重启并显式启用 `pg_xarray.wms_enabled`，默认关闭。启用 SQL 目录并不会开启公开瓦片服务。备份时需同时保留外部数组数据与目录。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.4.2 通过 `pg_xarray.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 chunk 目录去重键现在包含 `chunk_key`；升级后重新登记受影响的 Zarr 变量，恢复此前缺失的层级或瓦片。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.4.2 新增 `pgx.worker_status()` 与仅超级用户可调用的 `pgx.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_xarray.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_xarray UPDATE;
SELECT * FROM pgx.worker_status();
```
