## 用法

来源：

- [sql/pg_tviews--0.1.0-beta.23--0.1.0-beta.24.sql](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/sql/pg_tviews--0.1.0-beta.23--0.1.0-beta.24.sql)
- [sql/pg_tviews--0.1.0-beta.21--0.1.0-beta.22.sql](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/sql/pg_tviews--0.1.0-beta.21--0.1.0-beta.22.sql)
- [sql/pg_tviews--0.1.0-beta.22--0.1.0-beta.23.sql](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/sql/pg_tviews--0.1.0-beta.22--0.1.0-beta.23.sql)
- [README.md](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/README.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/CHANGELOG.md)
- [scripts/migrate-from-0.1.0.sql](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/scripts/migrate-from-0.1.0.sql)
- [docs/reference/read-contract.md](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/docs/reference/read-contract.md)
- [docs/operations/replication.md](https://github.com/fraiseql/pg_tviews/blob/39637679ade96d05dd8d789096438bed246624c7/docs/operations/replication.md)

`pg_tviews` 0.1.0-beta.24 通过查询分析与基表触发器在事务内维护派生 TVIEW 表。扩展对象现在全部位于固定的 `tviews` 模式，SQL 版本也与发布版本一致。

### 核心用法

```sql
CREATE EXTENSION pg_tviews;
CREATE TABLE tb_post (pk_post bigint PRIMARY KEY, title text);
INSERT INTO tb_post VALUES (1, 'Example');
SELECT tviews.pg_tviews_create('post',
  $$ SELECT pk_post, jsonb_build_object('title', title) AS data FROM tb_post $$);
SELECT * FROM tv_post;
SELECT * FROM tviews.registry;
SELECT * FROM tviews.pg_tviews_health_check();
```

### 运行边界

beta.23 的基本创建／刷新支持惰性加载，无需共享预加载；自动重建工作进程及其服务器启动参数仍需预加载并重启。控制文件不限定超级用户，但调用方需要创建模式／触发器的权限，并必须拥有要替换或删除的 TVIEW。级联以各 TVIEW 所有者身份执行。查询必须提供契约要求的实体键与 JSONB 数据列。UNLOGGED TVIEW 在备库不可读，崩溃／提升后会清空，需选择 LOGGED 存储或安排重建。

工具应通过 `tviews.registry` 与 `tviews.contract_version()` 读取接口。`tviews.pg_tviews_create_or_replace` 应用兼容的定义／选项变化，`tviews.pg_tviews_reregister_all` 在保留行数据的同时重新生成元数据与触发器。库／目录版本不匹配会阻止写入，直至迁移完成。旧 SQL 版本 0.1.0（截至 beta.19）需要链接的事务迁移脚本，它保留 TVIEW 行，但遇到扩展外部依赖会拒绝执行；普通 ALTER EXTENSION 不能替代该迁移。必须先备份并预演。

### beta.21–22 升级

`pg_tviews.uncascaded_policy` 在创建时保存：`warn`（默认）、`error` 或 `full_refresh` 决定如何处理无法将写入映射到 TVIEW 键的基表。beta.22 修复 READ COMMITTED 刷新之间的更新丢失，以及若干外连接、CTE 和 DISTINCT ON 依赖路径。正常 beta 版本升级后需重新登记全部 TVIEW；旧 SQL 版本 0.1.0 仍需执行前述独立迁移。

```sql
ALTER EXTENSION pg_tviews UPDATE;
SELECT * FROM tviews.pg_tviews_reregister_all();
```

### beta.23 行身份

beta.23 将各 TVIEW 的唯一行身份记录在 `tviews.registry.identity`；`DISTINCT ON` 视图使用单列键，复合键或表达式键会被拒绝。该版本修复多组写入、键值变化及向父 TVIEW 的传播。重新登记还会调整主键／唯一索引，应为这些表结构变化安排执行窗口。

### beta.24 生成列

beta.24 修复 PostgreSQL 18 虚拟生成列的依赖跟踪与刷新键映射。安装配套共享库后，执行 ALTER EXTENSION UPDATE 和 `tviews.pg_tviews_reregister_all()`，重新生成已有 TVIEW 的元数据与触发器。这是扩展自身的更新，不会为旧版 PostgreSQL 增加虚拟生成列。
