## 用法

来源：

- [README.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/README.md)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/CHANGELOG.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/Cargo.toml)
- [sql/pg_tviews--0.1.0-beta.26--0.1.0-beta.27.sql](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/sql/pg_tviews--0.1.0-beta.26--0.1.0-beta.27.sql)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/docs/user-guides/operators.md)
- [docs/DEPRECATION_WARNINGS.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/docs/DEPRECATION_WARNINGS.md)

`pg_tviews` 0.1.0-beta.27 通过解析后的查询与基础表触发器，在事务中维护派生的 TVIEW 表。对象固定在 `tviews` 模式中，当前源码支持 PostgreSQL 16–18。

### 核心用法

基础 SQL 函数可以延迟加载库，但上游快速入门采用预加载，确保各会话的钩子都已生效。自动重建工作进程需要预加载并重启服务器。修改配置时应保留已有的其他库。

```ini
shared_preload_libraries = 'pg_tviews'
```

```sql
CREATE EXTENSION pg_tviews;
SET search_path = public, tviews;
CREATE TABLE tb_post (pk_post bigint PRIMARY KEY, title text);
INSERT INTO tb_post VALUES (1, 'Example');
SELECT tviews.pg_tviews_create('post',
  $$ SELECT pk_post, jsonb_build_object('title', title) AS data FROM tb_post $$);
SELECT * FROM tv_post;
SELECT * FROM tviews.registry;
SELECT * FROM tviews.pg_tviews_health_check();
```

定义必须恰好是一条 SELECT，并提供约定的整数实体键和 JSONB 数据。调用者需要模式、表和触发器权限，也必须拥有所维护的 TVIEW。`tviews.registry`、`tviews.contract_version()` 和 `tviews.registry.managed_indexes` 描述注册对象，应通过这些接口检查，避免手工修改内部计划。`tviews.pg_tviews_create_or_replace` 修改定义，`tviews.pg_tviews_refresh` 则以 TVIEW 所有者身份刷新。

### 升级到 beta.27

先备份并演练。对于 beta.20 及更新版本，安装匹配的库后执行以下语句，会自动重新推导全部 TVIEW 并重装触发器，无须再单独执行一轮重新注册。如果解析失败，应在旧版本修复或删除错误中指出的定义，再重试升级。旧 SQL 版本 0.1.0 则需要使用上游迁移脚本。

```sql
ALTER EXTENSION pg_tviews UPDATE;
```

升级前应检查依赖：废弃接口和 `tviews.tviewschema` 类型会被删除，后者使用 CASCADE，可能连带删除依赖它的用户对象。全局维护函数会撤销 PUBLIC EXECUTE，部署与恢复工具所需函数必须显式授权。单个 TVIEW 的维护操作还会检查所有权。受管理索引的保留名称不能用于不同定义的用户索引；应在后续备份恢复前重命名冲突索引。

### 刷新与并发边界

提交时仍有待刷新工作会产生 SQLSTATE `55000`；损坏的计划或缺失映射也会使写入失败，避免提交过时结果。并发写入使用值锁和键锁，可能产生 `40P01` 或 `40001`，两类错误都需要重试，并优先使用 READ COMMITTED。`pg_tviews.lock_escalation_threshold` 默认为 64，`tviews.pg_tviews_queue_stats()` 可查看锁等待与升级信息。

依赖时间的定义需要外部调度调用 `tviews.pg_tviews_refresh_time_dependent()`。函数内部读表和读取物化视图需要声明相应维护策略。刷新函数使用受限搜索路径，因此被调用函数应限定关系名。UNLOGGED TVIEW 在崩溃恢复后会丢失内容，也不能在备库读取；需要规划重建或使用 LOGGED 存储。`pg_tviews.uncascaded_policy` 默认报错，现有集成代码不能继续使用已移除的级联、插入、删除和转换接口。
