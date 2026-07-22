## 用法

来源：

- [已复核 commit 的 pg_git README](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/README.md)
- [仓库初始化 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/001-init.sql)
- [暂存 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/002-add.sql)
- [提交 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/003-commit.sql)
- [远程操作 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/011-remote.sql)
- [HTTPS 与凭据 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/014-https.sql)

`pg_git` 在 PostgreSQL 表中实现类似 Git 的对象模型。版本 `0.4.0` 在固定的 `pggit` schema 中提供本地仓库初始化、暂存、提交、日志、分支、标签、差异、快进合并原语、重置、垃圾回收和完整性检查。它适合试验事务性、可查询的版本图；远程传输仍不完整，不能直接替代 Git 命令行客户端。

扩展依赖 `plpgsql`、`pgcrypto`、`pg_trgm` 和非可信语言 `plpython3u`。安装通常需要超级用户，以及与服务器匹配的 Python 过程语言软件包。

### 本地工作流

```sql
CREATE EXTENSION pg_git CASCADE;

SELECT pggit.init_repository('demo', '/logical/demo');

SELECT pggit.stage_file(
  (SELECT id FROM pggit.repositories WHERE name = 'demo'),
  'README.md',
  convert_to('hello from PostgreSQL', 'UTF8')
);

SELECT pggit.commit_index(
  (SELECT id FROM pggit.repositories WHERE name = 'demo'),
  'alice',
  'initial content'
);

SELECT * FROM pggit.get_log(
  (SELECT id FROM pggit.repositories WHERE name = 'demo')
);
```

仓库内容以字节为单位。`stage_file` 接受路径和 `bytea`；`commit_index` 写入暂存树并推进仓库历史。其他已实现入口包括 `unstage_file`、`get_status`、`get_decorated_log`、`create_branch`、`checkout_branch`、`merge_branches`、`diff_blobs`、`diff_commits`、`reset_soft`、`reset_mixed`、`create_tag`、`gc`、`verify_integrity` 和 `optimize_indexes`。

### 远程访问与凭据边界

schema 中包含 `add_remote`、`fetch_remote`、`push`、`pull`、`clone`、`store_credentials` 和 `http_fetch`，但这些名称夸大了当前互操作能力。在已复核 commit 中，`push` 只发出通知而不上传对象，`fetch_remote` 则从数据库中已经存在的引用行工作。把任何远程函数视为 Git 网络行为前，必须核对其 SQL 实现。

HTTPS 辅助函数在 `plpython3u` 中执行外部请求。凭据存储使用可配置的 `pggit.credential_key`；未设置时会退回 `pg_git_default_key`。保存任何密钥前，应设置部署专用的强密钥，并保护设置本身。

### 运维注意事项

- 非可信 Python 和外部 HTTPS 扩大了数据库安全边界。应限制函数执行、数据库出口、扩展所有权以及可写入远程 URL 的角色。
- 本地表更新参与 PostgreSQL 事务，但外部 HTTP 效果不会随事务回滚。必须显式设计重试与对账。
- 对象内容与历史会消耗数据库存储和 WAL。生产使用前，应测量大型仓库、备份增长、复制负载、垃圾回收和查询计划。
- README 混合了已实现函数与规划功能。只能依赖在确切 `0.4.0` commit 中验证过的可调用 SQL，并用有代表性的仓库演练恢复与完整性检查。
