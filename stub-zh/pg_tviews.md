## 用法

来源：

- [官方 README](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/README.md)
- [官方扩展 SQL](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/sql/pg_tviews--0.1.0.sql)
- [官方创建 API 源码](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/src/ddl/mod.rs)
- [官方运行时配置源码](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/src/config/mod.rs)

`pg_tviews` 通过分析 SELECT 查询、跟踪依赖关系并在基表上安装刷新触发器，以增量且事务一致的方式维护派生 `tv_*` 表。目录/控制文件版本为 `0.1.0`，而复核的 README 将当前代码标为 `0.1.0-beta.11`，并明确限制为评估用途，不适合关键任务系统。

### 核心流程

创建扩展，然后使用显式创建 API，传入能返回所需实体键与 JSONB 数据字段的查询：

```sql
CREATE EXTENSION pg_tviews;

SELECT pg_tviews_create(
  'post',
  $query$
    SELECT
      p.pk_post,
      p.fk_user,
      jsonb_build_object(
        'id', p.id,
        'title', p.title,
        'author', u.name
      ) AS data
    FROM tb_post AS p
    JOIN tb_user AS u ON u.pk_user = p.fk_user
  $query$
);

SELECT data FROM tv_post WHERE pk_post = 42;
```

该 API 会创建 `tv_post`、记录依赖元数据并安装触发器，使源记录变化能在同一事务中刷新受影响的派生记录。README 也记录了被拦截的 `CREATE TABLE tv_* AS SELECT ...` 语法；显式函数能让扩展加载和错误报告更加清晰。

### 重要操作

- `pg_tviews_create(name, select_sql)` 创建 TVIEW；`pg_tviews_drop(name, if_exists, cascade)` 删除它。
- `pg_tviews_health_check()` 与 `pg_tviews_queue_realtime` 提供健康和队列信息。
- `pg_tviews_suspend_triggers()`、`pg_tviews_resume_triggers()` 与 `pg_tviews_refresh_all()` 支持在批量装载前后按依赖顺序刷新。
- `pg_tviews_recover_after_crash(name)` 检查并重建崩溃后被清空的 UNLOGGED TVIEW。
- 复核源码中的 `pg_tviews.unlogged_by_default` 默认为 true；其他 GUC 控制队列大小、传播深度、缓存、指标、审计日志与触发器暂停。

### 批量装载模式

```sql
BEGIN;
SELECT pg_tviews_suspend_triggers();
INSERT INTO tb_post SELECT * FROM staging_post;
SELECT pg_tviews_resume_triggers();
COMMIT;

SELECT pg_tviews_refresh_all();
```

### 运维注意事项

TVIEW 定义遵循 FraiseQL 约定：查询必须公开预期的 `pk_<entity>` 血缘键与 `data` JSONB 字段，并使用外键字段传播级联。默认 UNLOGGED 存储可提高写入速度，但会在 PostgreSQL 崩溃后被清空，必须从基表重建。上游基准只是测量结果而不是保证；需要用代表性数据测量触发器延迟、级联、批量装载、WAL/存储权衡、崩溃恢复、备份恢复与模式变更。API 在 `1.0.0` 前可能变化；应固定精确构建、演练升级，并在 beta 警告移除且工作负载通过验证前避免用于关键任务。
