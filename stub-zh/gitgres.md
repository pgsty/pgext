## 用法

来源：

- [官方 README](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/README.md)
- [官方扩展 SQL](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/ext/sql/gitgres--0.1.sql)
- [官方 libgit2 后端源码](https://github.com/andrew/gitgres/tree/eaf8743f2c137d61a75432e44e13467cad7eceaa/backend)

`gitgres` 将 Git 对象、仓库、引用与 reflog 记录保存在 PostgreSQL 中，并提供 SQL 解析器以及独立的 libgit2/libpq 后端，用于 push 和 clone。版本 `0.1` 是一个较新的实现，以 20 字节 SHA-1 Git 对象标识符为核心；在替代文件系统 Git 存储前，应使用真实仓库规模和恢复要求进行评估。

### 核心流程

安装扩展及其 `pgcrypto` 依赖，构建配套后端，然后初始化命名仓库：

```sql
CREATE EXTENSION gitgres CASCADE;
```

```bash
./backend/gitgres-backend init "dbname=gitgres" myrepo
./backend/gitgres-backend push "dbname=gitgres" myrepo /path/to/repo
./backend/gitgres-backend clone "dbname=gitgres" myrepo /path/to/dest
./backend/gitgres-backend ls-refs "dbname=gitgres" myrepo
```

导入或推送对象后，需要先刷新面向查询的物化视图再读取：

```sql
REFRESH MATERIALIZED VIEW commits_view;
REFRESH MATERIALIZED VIEW tree_entries_view;

SELECT sha, author_name, authored_at, message
FROM commits_view
ORDER BY authored_at DESC;
```

### 重要对象

- `repositories`、`objects`、`refs` 与 `reflog` 是主要存储表。
- `git_oid` 是固定 20 字节类型，支持比较、B-tree 与哈希。
- `git_object_write(...)`、`git_object_read(...)` 与 `git_object_read_prefix(...)` 管理原始 Git 对象。
- `git_tree_entries(bytea)`、`git_ls_tree_r(...)` 与 `git_commit_parse(bytea)` 展开树和提交结构。
- `git_ref_update(...)` 对引用执行比较并交换更新；`git_ref_set_symbolic(...)` 管理符号引用。
- `commits_view` 与 `tree_entries_view` 是物化报表视图，不会自动刷新。

### 运维注意事项

扩展依赖 `pgcrypto`；push/clone 命令还需要 libgit2、libpq、OpenSSL、源或目标仓库的文件系统访问权，以及拥有适当表/函数权限的数据库角色。应保护连接字符串，并限制对存储表的直接写入，避免绕过引用不变量。Git 对象数据可能体积大且写入密集，需要测试 WAL 量、vacuum、备份、副本、物化视图刷新成本和恢复往返流程。`git_oid` 格式只支持 SHA-1 仓库，不能据此推断支持 Git 的 SHA-256 对象格式。数据库备份只有在恢复后克隆出的仓库通过 Git 完整性检查时，才算经过验证的 Git 恢复。
