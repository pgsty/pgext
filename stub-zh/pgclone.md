## 用法

来源：[README](https://github.com/valehdba/pgclone/blob/main/README.md)、[Usage guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md)、[Async guide](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md)、[Release v4.3.2](https://github.com/valehdba/pgclone/releases/tag/v4.3.2)、[changelog](https://github.com/valehdba/pgclone/blob/main/CHANGELOG.md)、[SQL install script](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.3.2.sql)

`pgclone` 可直接从 SQL 克隆表、schema、函数、角色和整个数据库。在 v4.x 中，公共 API 位于 `pgclone` schema 下；上游和 Pigsty 当前都跟踪 PostgreSQL 14-18。

### 核心克隆函数

```sql
CREATE EXTENSION pgclone;
SELECT pgclone.version();

SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres password=secret',
  'public',
  'customers',
  true
);

SELECT pgclone.schema(
  'host=source-server dbname=mydb user=postgres password=secret',
  'sales',
  true
);

SELECT pgclone.database(
  'host=source-server dbname=mydb user=postgres password=secret',
  true
);
```

- `pgclone.table(...)`、`pgclone.schema(...)`、`pgclone.functions(...)`、`pgclone.database(...)`
- `pgclone.database_create(...)` 会创建本地目标数据库并克隆进去。
- `_ex` 变体暴露显式布尔参数，用于控制 indexes、constraints 和 triggers。

### 选项与 Masking

- JSON options 支持 `columns`、`where`、`conflict`，以及 `indexes`、`constraints`、`triggers` 等对象开关。
- JSON options 也包含 `consistent`；v4.3.0+ 默认使用跨表一致快照，并可通过 `{"consistent": false}` 在单次调用中禁用。
- 上游 usage guide 记录了 masking、敏感列自动发现、static masking、dynamic masking、clone verification，以及 GDPR/compliance reporting。

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

### 一致性、Diff 与 Preflight

```sql
SELECT pgclone.diff(
  'host=source-server dbname=prod user=postgres',
  'app_schema'
)::jsonb;

SELECT pgclone.preflight(
  'host=source-server dbname=prod user=postgres',
  'app_schema'
)::jsonb;
```

- `pgclone.diff(conninfo, schema)` 报告只读 DDL drift，覆盖 tables、columns、indexes、constraints、triggers、views 和 sequences。
- `pgclone.preflight(conninfo, schema)` 在 clone 前检查 source 和 target 的就绪状态，包括 connection、version、permission、capacity、naming-conflict、missing-role、missing-extension 和 tablespace 问题。
- v4.3.0+ clone 默认在 source 上以 `REPEATABLE READ READ ONLY` 读取。多连接 schema、database 和 parallel-pool clones 共享同一个 exported snapshot，从而在 live source 持续写入时保持 parent/child 一致性。
- 长时间 clone 会在 source 上保持事务打开，可能延迟 vacuum cleanup 和 WAL recycling；当这个代价比跨表一致性更重要时，可使用 `{"consistent": false}`。

### 异步与进度

```sql
-- postgresql.conf
shared_preload_libraries = 'pgclone'

SELECT pgclone.schema_async(
  'host=source-server dbname=mydb user=postgres',
  'sales', true, '{"parallel":4}'
);

SELECT * FROM pgclone.jobs_view;
SELECT pgclone.progress(1);
SELECT pgclone.cancel(1);
```

- `pgclone.table_async(...)` 和 `pgclone.schema_async(...)` 在 background workers 中运行。
- `pgclone.jobs_view`、`pgclone.progress_detail()`、`pgclone.resume()` 和 `pgclone.clear_jobs()` 提供任务跟踪与恢复。
- v4.3.2 将 snapshot-keeper resilience fixes 移植到 async/background-worker 路径，包括 keepalive injection，以及面向网络 source connections 的 timeout protection。

### 注意事项

- 上游要求 PostgreSQL 14+。
- usage guide 说明安装和使用该扩展需要 superuser 权限。
- 异步功能需要 `shared_preload_libraries = 'pgclone'`；worker-pool parallelism 也依赖 `max_worker_processes`。
- 如果必须绕过 source-side snapshot 问题，一致性异步 clone 仍可使用 `{"consistent": false}` 退出一致快照模式。
- Pigsty 为 PostgreSQL 14-18 打包 `4.3.2`。2026 年 6 月的 RPM rebuild 使用了 `LLVM_BINPATH` 构建修复；已复核上游，除包 caveat 和既有 v4.3.2 async-snapshot 说明外，没有实质 stub 变化。
