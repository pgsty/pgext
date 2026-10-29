## 用法

来源：

- [pg_tviews--0.1.0-beta.30--0.1.0-beta.31.sql](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/sql/pg_tviews--0.1.0-beta.30--0.1.0-beta.31.sql)
- [pg_tviews--0.1.0-beta.31--0.1.0-beta.32.sql](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/sql/pg_tviews--0.1.0-beta.31--0.1.0-beta.32.sql)

- [README.md](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/README.md)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/CHANGELOG.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/Cargo.toml)
- [sql/pg_tviews--0.1.0-beta.29--0.1.0-beta.30.sql](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/sql/pg_tviews--0.1.0-beta.29--0.1.0-beta.30.sql)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/docs/user-guides/operators.md)
- [docs/DEPRECATION_WARNINGS.md](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/docs/DEPRECATION_WARNINGS.md)
- [beta.33 SQL 升级](https://github.com/fraiseql/pg_tviews/blob/2a76abdbf8a9ff6cbd365a307d22669574ef0a65/sql/pg_tviews--0.1.0-beta.32--0.1.0-beta.33.sql)

`pg_tviews` 0.1.0-beta.33 通过解析后的查询与基础表触发器，在事务中维护派生的 TVIEW 表。对象固定在 `tviews` 模式中，当前源码支持 PostgreSQL 16–18。

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

### 升级到 beta.30

beta.30 合并了未单独发布的 beta.28 和 beta.29。先备份并演练，安装匹配库并重启 PostgreSQL，再在各数据库更新 SQL 扩展，最后以有权限的维护角色重新注册全部 TVIEW：

```sql
ALTER EXTENSION pg_tviews UPDATE;
SELECT * FROM tviews.pg_tviews_reregister_all();
```

升级会重新推导传播计划、替换若干接口并创建缺少的引用列索引；大型 TVIEW 的索引构建需要预留时间。`tviews.contract_version()` 现在为 2；工具应从 `tviews.registry.options` 读取选项，命名参数统一使用 `tview =>`。`tviews.registry.references` 描述外键证明的引用关系，`tviews.registry.typename` 给出类型名。

`tviews.pg_tviews_create_or_replace` 的 options 是完整声明，省略的选项恢复默认值；应显式传回需要保留的选项。新 TVIEW 默认使用 LOGGED 存储，已有表保留原持久化属性。存储选项变更可能重写全部行。删除已移除的 GUC 配置，尤其是旧 `pg_tviews.uncascaded_policy`；策略现属于各 TVIEW 的选项。跨越更早版本时还应检查被移除的 API、`tviews.tviewschema` 的 CASCADE 删除及维护函数权限变化。

### 刷新与并发边界

提交时仍有待刷新工作会产生 SQLSTATE `55000`；损坏的计划或缺失映射也会使写入失败，避免提交过时结果。并发写入使用值锁和键锁，可能产生 `40P01` 或 `40001`，两类错误都需要重试，并优先使用 READ COMMITTED。`pg_tviews.lock_escalation_threshold` 默认为 64，`tviews.pg_tviews_queue_stats()` 可查看锁等待与升级信息。

依赖时间的定义需要外部调度调用 `tviews.pg_tviews_refresh_time_dependent()`。函数内部读表和读取物化视图需要声明相应维护策略。刷新函数使用受限搜索路径，因此被调用函数应限定关系名。UNLOGGED TVIEW 在崩溃恢复后会丢失内容，也不能在备库读取；需要规划重建或使用 LOGGED 存储。每个 TVIEW 的未级联策略默认报错，现有集成代码不能继续使用已移除的级联、插入、删除和转换接口。

### 统计、通知与写入上限

`tviews.stats` 提供跨会话刷新计数，依赖预加载。`pg_tviews.auto_rebuild_databases` 现在默认覆盖全部数据库，以重建崩溃或提升后被清空的 UNLOGGED TVIEW。`pg_tviews.max_queue_size` 默认 1000000；超级用户可用 `pg_tviews.max_refresh_rows` 限制单条语句的刷新行数，超限以 `54000` 拒绝写入。

`pg_tviews.publish_changes` 默认 off，可由超级用户设为 types 或 rows，在提交时向 `tviews_changes` 发布变化。beta.30 的通知载荷为 `v: 2`，新增 `c` 表示受影响的父级集合；不了解新版本的监听器应使全部缓存失效。NOTIFY 队列使用率超过 90% 时可能提交但跳过通知，并记录警告；它不能作为可靠持久日志。有待发布变化时，PREPARE TRANSACTION 会以 `25000` 失败。

### 从 beta.30 升级到 beta.32

安装匹配库并按预加载库流程重启后执行 `ALTER EXTENSION pg_tviews UPDATE`，使目录修订升级到 9；从 beta.30 或 beta.31 升级无需重新注册 TVIEW。前述重新注册要求适用于跨越更早版本的升级。

`tviews.pg_tviews_feed_gaps()` 和队列统计中的 `feed_gaps` 统计集群启动后提交但未发布通知的事务，依赖预加载。监听器应定期检查此计数，变化时使缓存全部失效。beta.32 移除了会话级 `publish_skipped`，使用者应改读新计数。`pg_tviews_flush_and_report` 在发布 rows 时新增 collections；invalidated_types 覆盖无法逐行列举的整个类型及其集合，不能把旧行列表当作完整结果。

这些版本还修复了写入依赖其他 TVIEW 模式权限的问题，并按实体与计划读取目录，减少无关 TVIEW 对写入的影响。逻辑恢复时若省略 TVIEW 表数据，后续触及它的写入会在检测到源视图有数据时进行补填；不能将这个机制视为恢复后立即完整的保证。

### beta.33 的分区正确性修复

删除分区现在会在同一条语句中，从该分区树关联的所有 TVIEW 移除对应数据；删除外部分区会完整刷新受影响的 TVIEW。普通 DETACH 只刷新受影响的键，ATTACH 与并发 DETACH 仍采用完整刷新。安装匹配库文件并重启预加载服务器后，再从 beta.32 更新。SQL 升级是空操作，目录修订仍为 9，这次转换无须重新登记 TVIEW；跨越更早的受影响版本时，原有迁移要求仍适用。
