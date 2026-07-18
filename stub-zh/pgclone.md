


## 用法

来源：[README](https://github.com/valehdba/pgclone/blob/main/README.md)、[使用指南](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md)、[异步指南](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md)、[v4.3.2 发行版](https://github.com/valehdba/pgclone/releases/tag/v4.3.2)、[变更日志](https://github.com/valehdba/pgclone/blob/main/CHANGELOG.md)、[SQL 安装脚本](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.3.2.sql)

`pgclone` 可直接从 SQL 克隆表、模式、函数、角色和整个数据库。在 v4.x 中，公共 API 位于 `pgclone` 模式下；上游和 Pigsty 当前都跟踪 PostgreSQL 14-18。

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
- `_ex` 变体暴露显式布尔参数，用于控制索引、约束和触发器。

### 选项与脱敏

- JSON 选项支持 `columns`、`where`、`conflict`，以及 `indexes`、`constraints`、`triggers` 等对象开关。
- JSON 选项也包含 `consistent`；v4.3.0+ 默认使用跨表一致快照，并可通过 `{"consistent": false}` 在单次调用中禁用。
- 上游使用指南记录了脱敏、敏感列自动发现、静态脱敏、动态脱敏、克隆验证，以及 GDPR/合规报告。

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

### 一致性、差异与预检

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

- `pgclone.diff(conninfo, schema)` 报告只读 DDL 漂移，覆盖表、列、索引、约束、触发器、视图和序列。
- `pgclone.preflight(conninfo, schema)` 在克隆前检查源端和目标端的就绪状态，包括连接、版本、权限、容量、命名冲突、缺失角色、缺失扩展和表空间问题。
- v4.3.0+ 默认在源端以 `REPEATABLE READ READ ONLY` 执行克隆读取。多连接的模式、数据库和并行池克隆共享同一个导出快照，从而在活动源端持续写入时保持父子对象一致性。
- 长时间克隆会在源端保持事务打开，可能延迟清理和 WAL 回收；当这个代价比跨表一致性更重要时，可使用 `{"consistent": false}`。

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

- `pgclone.table_async(...)` 和 `pgclone.schema_async(...)` 在后台工作进程中运行。
- `pgclone.jobs_view`、`pgclone.progress_detail()`、`pgclone.resume()` 和 `pgclone.clear_jobs()` 提供任务跟踪与恢复。
- v4.3.2 将快照保持器的弹性修复移植到异步/后台工作进程路径，包括保活注入，以及面向网络源连接的超时保护。

### 注意事项

- 上游要求 PostgreSQL 14+。
- 使用指南说明安装和使用该扩展需要超级用户权限。
- 异步功能需要 `shared_preload_libraries = 'pgclone'`；工作进程池的并行度也依赖 `max_worker_processes`。
- 如果必须绕过源端快照问题，一致性异步克隆仍可使用 `{"consistent": false}` 退出一致快照模式。
- Pigsty 为 PostgreSQL 14-18 打包 `4.3.2`。2026 年 6 月的 RPM 重建使用了 `LLVM_BINPATH` 构建修复；已复核上游，除软件包注意事项和既有 v4.3.2 异步快照说明外，没有实质文档变化。
