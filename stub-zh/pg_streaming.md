## 用法

来源：

- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/pg_streaming.control)
- [管道 SQL API 与 worker 初始化](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/src/lib.rs)
- [管道定义类型](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/src/dsl/types.rs)

`pg_streaming` 版本 `0.2.0` 是声明式流处理引擎，其管道、状态、偏移量、错误与指标都保留在 PostgreSQL 中并可查询。管道把输入连接到处理器链和输出，并由协调、执行与定时后台工作进程运行。

### 核心流程

预加载库、重启 PostgreSQL、在配置的数据库中创建扩展，然后定义管道：

```conf
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
