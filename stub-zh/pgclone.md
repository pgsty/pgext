## 用法

来源：[README](https://github.com/valehdba/pgclone/blob/main/README.md)，[Usage guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md)，[Async guide](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md)，[Release v4.0.0](https://github.com/valehdba/pgclone/releases/tag/v4.0.0)，[SQL install script](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.0.0.sql)

`pgclone` 可以直接通过 SQL 克隆 tables、schemas、functions、roles 以及整个 database。在 v4.0.0 中，公开 API 被放入 `pgclone` schema。

### 核心克隆函数

```sql
CREATE EXTENSION pgclone;

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
- `pgclone.database_create(...)` 会创建本地目标 database 并在其中执行克隆。
- `_ex` 变体显式暴露 indexes、constraints 与 triggers 的布尔开关。

### 选项与脱敏

- JSON 选项支持 `columns`、`where`、`conflict`，以及 `indexes`、`constraints`、`triggers` 等对象开关。
- 上游在 usage guide 中记录了 masking、敏感列自动发现、静态脱敏、动态脱敏、clone verification，以及 GDPR/compliance reporting。

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

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

- `pgclone.table_async(...)` 与 `pgclone.schema_async(...)` 在 background workers 中运行。
- `pgclone.jobs_view`、`pgclone.progress_detail()`、`pgclone.resume()` 与 `pgclone.clear_jobs()` 提供 job 跟踪与恢复能力。

### 注意事项

- 上游要求 PostgreSQL 14+。
- usage guide 说明安装和使用该扩展都需要 superuser 权限。
- 异步功能需要 `shared_preload_libraries = 'pgclone'`；worker pool 并行度还取决于 `max_worker_processes`。
