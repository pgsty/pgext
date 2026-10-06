## 用法

来源：

- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pgbrew.toml)
- [extensions/pg_s3/pg_s3.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pg_s3.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pgbrew.toml)

`pg_s3` 0.3.2 提供实验性的 S3 兼容 HTTP API，在 `pgs3` 中保存桶／对象元数据，在本地磁盘保存二进制内容。

### 核心用法

```ini
shared_preload_libraries = 'pg_s3'
pg_s3.database = 'postgres'
pg_s3.host = '127.0.0.1'
pg_s3.port = 9100
```

```sql
CREATE EXTENSION pg_s3;
SELECT pgs3.create_bucket('sample-bucket');
SELECT * FROM pgs3.list_buckets();
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启。通过 `pg_s3.database`、`pg_s3.host`、`pg_s3.port` 和 `pg_s3.data_directory` 配置服务，默认包括所有接口、9100 端口与 postgres 数据库。SQL API 管理桶、对象元数据及列表。备份必须一致地覆盖数据库元数据和对象文件。应限制文件路径与网络访问；源码可用不代表完整的 S3 API、认证或多节点持久性兼容。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.3.2 通过 `pg_s3.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.3.2 新增 `pgs3.worker_status()` 与仅超级用户可调用的 `pgs3.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_s3.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_s3 UPDATE;
SELECT * FROM pgs3.worker_status();
```
