## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_streaming/pgbrew.toml)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_streaming/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_streaming/pgbrew.toml)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_streaming/pg_streaming.control)
- [管道 SQL API 与 worker 初始化](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_streaming/src/lib.rs)
- [管道定义类型](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` 版本 `0.3.1` 是声明式流处理引擎，其管道、状态、偏移量、错误与指标都保留在 PostgreSQL 中并可查询。管道把输入连接到处理器链和输出，并由协调、执行与定时后台工作进程运行。

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

扩展版本 0.3.1 通过 `pg_streaming.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.4.1；上游二进制不代表 Pigsty 软件包可用性。
