## 用法

来源：

- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/CHANGELOG.md)
- [docs/reference/api.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/reference/api.md)
- [docs/concurrency.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/concurrency.md)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/user-guides/operators.md)
- [docs/operations/upgrade.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/operations/upgrade.md)
- [docs/operations/upgrades.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/operations/upgrades.md)

`pg_tviews` 0.1.0-beta.35 根据已分析的查询和基础表触发器，在事务中维护派生的 TVIEW 表。它支持 PostgreSQL 16–18，将 API 安装在 `tviews` 模式中，并依赖 `jsonb_delta`。

### 核心流程

在每个使用它的服务器上预加载库，保留原有预加载条目，然后重启 PostgreSQL。预加载是强制要求；缺少它时，安装以及后续函数调用或写入都会以 SQLSTATE `55000` 失败。

```ini
shared_preload_libraries = 'pg_tviews'
```

以超级用户安装并创建一个简单的 TVIEW：

```sql
CREATE EXTENSION jsonb_delta;
CREATE EXTENSION pg_tviews;
CREATE TABLE tb_user (
    pk_user bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text NOT NULL
);
INSERT INTO tb_user (name) VALUES ('Alice'), ('Bob');
CREATE TABLE tv_user AS
SELECT pk_user, id, jsonb_build_object('id', id, 'name', name) AS data
FROM tb_user;
SELECT * FROM tv_user;
SELECT * FROM tviews.pg_tviews_check_health();
```

定义包含一条 SELECT、约定的整数实体键、UUID id 和 JSONB data。维护以 TVIEW 所有者身份执行，所有者需能 SELECT 源表；写入者需要基础表权限，读取者需要 TVIEW 的 SELECT 权限。`tviews.pg_tviews_create_or_replace` 要求拥有已有 TVIEW，并把传入的选项视为完整声明：省略项恢复默认值。默认使用 LOGGED 存储。

### 并发与对象

写入以及 flush/refresh/rebuild 操作要求 READ COMMITTED。更强的隔离级别会被 SQLSTATE `0A000` 拒绝；此版本不再主动产生 `40001`。死锁仍可能返回 `40P01`，需要重试事务。提交时仍有待刷新或挂起的工作会失败，避免提交过时结果。

通过 `tviews.registry` 查看注册状态，用 `tviews.contract_version()` 查询 API 契约版本，当前为 3。现行监控接口包括 `tviews.pg_tviews_show_queue`、`tviews.pg_tviews_show_suspended`、`tviews.pg_tviews_count_feed_gaps` 和 `tviews.pg_tviews_show_replication`。早期 health、queue 与 feed-gap 名称已删除，没有兼容别名。`tviews.stats` 提供跨会话维护计数。

扩展拒绝在由逻辑订阅写入的表上创建 TVIEW，应在发布端发布已维护的 TVIEW。UNLOGGED TVIEW 在崩溃恢复或提升后为空，且不能在备库读取；需要规划重建或使用 LOGGED 存储。

### 配置与变化通知

面向运维的五项设置为 `pg_tviews.auto_rebuild_databases`（默认所有数据库）、`pg_tviews.publish_changes`（off）、`pg_tviews.max_refresh_rows`（0，无限制）、`pg_tviews.lock_escalation_threshold`（64）和 `pg_tviews.audit_enabled`（off）。早期锁升级取值 -1 已删除。调整重建数据库范围需要重启，其余设置由超级用户控制。

隐藏的超级用户诊断设置 `pg_tviews.max_queue_size` 默认为 1000000。早期可配置的深度、缓存和批量限制现为固定内部限制，应删除过时设置，不能假定它们仍有效。超过刷新限制可能以 `54000` 拒绝写入。

`tviews_changes` 上的变化通知不是持久日志。通知队列压力较大时可能提交事务却跳过通知；消费者应监控 feed gap，并在计数增加时使缓存失效。存在待发布变化时会拒绝预备事务。

### 升级到 beta.35

先备份并演练，然后暂停写入。替换已加载的库前先停服，安装匹配文件后重启。SQL 直接升级仅支持前一发布 beta.34：

```sql
ALTER EXTENSION pg_tviews UPDATE TO '0.1.0-beta.35';
SELECT tviews.pg_tviews_version(), tviews.contract_version();
SELECT * FROM tviews.pg_tviews_check_health() WHERE status <> 'OK';
SELECT schema, name, needs_reregister FROM tviews.registry;
```

重启至完成 SQL 升级之间，基础表写入会以 `55000` 失败。升级以各 TVIEW 所有者身份重新推导全部计划和触发器，保留已有行，并将每个 TVIEW 的触发器改为每张基础表一组。任一定义无法分析时，整个升级回滚。恢复写入前，应恢复被替换函数的授权，包括 `pg_tviews_ensure_propagation_indexes_all(boolean)` 与 `pg_tviews_reset_stats_all()`。

更早版本应按官方删除重建流程，使用保存的定义按依赖顺序迁移；保留基础表行，并暂停写入直到全部 TVIEW 重建完成。不要把历史 beta.30–33 升级说明当作直达 beta.35 的路径。没有降级脚本：回滚需要匹配的旧版二进制，以及恢复到新数据库的已验证备份。
