## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/CHANGELOG.md)
- [extensions/pg_streaming/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/Cargo.toml)
- [extensions/pg_streaming/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/src/lib.rs)

- [extensions/pg_streaming/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/Cargo.toml)
- [extensions/pg_streaming/src/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/src/worker.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/CHANGELOG.md)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/pgbrew.toml)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/LICENSE)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/pgbrew.toml)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/pg_streaming.control)
- [管道 SQL API 与 worker 初始化](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/src/lib.rs)
- [管道定义类型](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` 版本 `0.3.4` 是声明式流处理引擎，其管道、状态、偏移量、错误与指标都保留在 PostgreSQL 中并可查询。管道把输入连接到处理器链和输出，并由协调、执行与定时后台工作进程运行。

### 核心流程

预加载库、重启 PostgreSQL、在配置的数据库中创建扩展，然后定义管道：

```ini
shared_preload_libraries = 'pg_streaming'
pg_streaming.database = 'postgres'
pg_streaming.worker_count = 2
```

```sql
CREATE EXTENSION pg_streaming;

SELECT pgstreams.create_pipeline(
    'active_orders',
    $json$
    {
      "input": {"table": {"name":"public.order_inbox", "offset_column":"id", "poll":"1s"}},
      "pipeline": {"processors": [{"filter":"value_json->>'active' = 'true'"}]},
      "output": {"table": {"name":"public.active_orders", "mode":"append"}}
    }
    $json$::jsonb
);

SELECT pgstreams.start('active_orders');
SELECT * FROM pgstreams.status();
SELECT * FROM pgstreams.metrics('active_orders');
SELECT pgstreams.stop('active_orders');
```

生命周期函数包括 `create_pipeline`、`update_pipeline`、`drop_pipeline`、`start`、`stop` 和 `restart`。可观测性函数包括 `status`、`errors`、`late_events`、`metrics`、`lag` 与 `trace`；密钥及自定义连接器注册函数用于连接器配置。DSL 覆盖表、CDC、Kafka、分页 HTTP、OpenDAL、自定义输入/输出，以及 filter、mapping、aggregate、window、join、dedupe 和 CEP 等处理器。

### 运维说明

固定的 `pgstreams` 模式和后台工作进程注册要求超级用户安装与服务器级规划。加入库或改变 worker 数量需要重启；应为一个协调 worker、配置数量的执行 worker 及一个定时 worker 预留足够的 `max_worker_processes` 容量。管道表达式和连接器定义在高权限数据库服务中执行，因此应限制管道与密钥管理权限。生产使用前，应针对每种连接器测试投递保证、检查点、重试、迟到数据、模式演进和故障恢复。

### 0.3.0 版本边界

这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

新增 `call` 输出可通过配置的 `set_role` 执行受控调用。Modbus TCP、Siemens S7 输入及写入输出会访问外部设备，启用前应限制连接器配置并确认操作授权。

### 当前版本与升级

扩展版本 0.3.2 通过 `pg_streaming.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.5.0；上游二进制不代表 Pigsty 软件包可用性。

### 工作进程监督

0.3.2 新增 `pgstreams.worker_status()` 与仅超级用户可调用的 `pgstreams.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_streaming.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_streaming UPDATE;
SELECT * FROM pgstreams.worker_status();
```

### 0.3.3 原始数据区读取

仓库发布 0.6.0 包含扩展 0.3.3。OpenDAL 新增 `parse_as: parquet` 与 `parse_as: listing`，`order: lexicographic` 可在文件中途恢复。不可读文件会停止推进，避免后续文件越过它。记录包含用于回放去重的 `source_file` 和 `source_row`，`pgstreams.replay(name, after)` 可回退已停止的管道。数组读取通道仍不支持私有 S3。

### 0.3.4 的空闲工作进程

仓库发布 0.7.2 包含扩展版本 0.3.4。只读轮询或未领取到作业的轮询不再仅因使用 SPI 桥接而分配事务 ID，实际写入行时仍会分配，因而空闲轮询不再产生空提交 WAL。 执行器空闲时，轮询间隔从 `pg_streaming.poll_interval_ms`（默认 10 ms）逐步退避到 `pg_streaming.idle_poll_interval_ms`（默认 250 ms）。 安装匹配的共享库、重启预加载的工作进程并更新 SQL 扩展。既有权限、外部服务和 0.3.0 之前版本的迁移边界仍然适用。
