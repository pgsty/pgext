## 用法

来源：

- [已复核 commit 的 pg_git README](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/README.md)
- [已复核 commit 的 pg_git.control](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/pg_git.control)
- [仓库初始化 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/001-init.sql)
- [暂存 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/002-add.sql)
- [提交 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/003-commit.sql)
- [远程操作 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/011-remote.sql)
- [HTTPS 与凭据 SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/014-https.sql)

`pg_git` 在 PostgreSQL 内完整保存类似 Git 的对象模型。对象位于 `pggit` 模式，已实现的函数覆盖仓库初始化、字节内容暂存、提交、历史、分支、标签、差异、快进合并、垃圾回收和完整性检查。

核心入口包括 `init_repository`、`stage_file`、`commit_index` 和 `get_log`。扩展在安装时需要 `plpgsql`、`pgcrypto`、`pg_trgm` 和 `plpython3u`。

### 基本本地工作流

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

### 注意事项

- 远程传输尚未完成。当前 `push` 函数只发出通知，`fetch_remote` 只是物化 `remote_refs` 中已有的行；不要把它当作可直接替代 Git 的网络客户端。
- HTTPS 辅助函数通过不受信任的 `plpython3u` 发出外部请求。启用前应复核数据库出站策略与 Python 权限。
- 存储凭据前应配置强 `pggit.credential_key`。如果没有该设置，当前源码会回退到 `pg_git_default_key`。
- 目录、control 文件和 README 都标识版本 `0.4.0`。README 明确区分已实现与规划中的操作；依赖某项功能前应核实对应函数。
