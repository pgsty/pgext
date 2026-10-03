## 用法

来源：

- [README.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/README.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/CHANGELOG.md)
- [scripts/migrate-from-0.1.0.sql](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/scripts/migrate-from-0.1.0.sql)
- [docs/reference/read-contract.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/docs/reference/read-contract.md)
- [docs/operations/replication.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/docs/operations/replication.md)

`pg_tviews` 0.1.0-beta.20 通过查询分析与基表触发器在事务内维护派生 TVIEW 表。扩展对象现在全部位于固定的 `tviews` 模式，SQL 版本也与发布版本一致。

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

beta.20 的基本创建／刷新支持惰性加载，无需共享预加载；自动重建工作进程及其服务器启动参数仍需预加载并重启。控制文件不限定超级用户，但调用方需要创建模式／触发器的权限，并必须拥有要替换或删除的 TVIEW。级联以各 TVIEW 所有者身份执行。查询必须提供契约要求的实体键与 JSONB 数据列。UNLOGGED TVIEW 在备库不可读，崩溃／提升后会清空，需选择 LOGGED 存储或安排重建。

工具应通过 `tviews.registry` 与 `tviews.contract_version()` 读取接口。`tviews.pg_tviews_create_or_replace` 应用兼容的定义／选项变化，`tviews.pg_tviews_reregister_all` 在保留行数据的同时重新生成元数据与触发器。库／目录版本不匹配会阻止写入，直至迁移完成。旧 SQL 版本 0.1.0（截至 beta.19）需要链接的事务迁移脚本，它保留 TVIEW 行，但遇到扩展外部依赖会拒绝执行；普通 ALTER EXTENSION 不能替代该迁移。必须先备份并预演。
