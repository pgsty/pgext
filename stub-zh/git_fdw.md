## 用法

来源：

- [官方 1.1.0 版 README](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/README.md)
- [官方 control 文件](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/git_fdw.control)
- [官方代码仓库](https://github.com/franckverrot/git_fdw)
- [官方 PGXN 发布页](https://pgxn.org/dist/git_fdw/)

`git_fdw` 把本地 Git 仓库的提交以只读 PostgreSQL 外部表形式暴露。它使用 libgit2，为指定分支返回提交身份、消息、时间戳，以及新增、删除和变更文件数量。

### 核心流程

创建封装器和无需选项的服务器，然后在 PostgreSQL 9.5 或更高版本中导入仓库模式：

```sql
CREATE EXTENSION git_fdw;
CREATE SERVER git_repo FOREIGN DATA WRAPPER git_fdw;

IMPORT FOREIGN SCHEMA git_data
FROM SERVER git_repo
INTO public
OPTIONS (
  path '/srv/git/project.git',
  branch 'refs/heads/main',
  prefix 'project_'
);

SELECT sha1, name, email, commit_date, message
FROM project_repository
ORDER BY commit_date DESC
LIMIT 10;
```

PostgreSQL 9.4 需要按文档手工定义外部表字段。一张外部表只对应一个仓库/分支；多个表可通过视图组合。

### 选项与字段

- `path`（必需）：PostgreSQL 服务器文件系统上的仓库路径。
- `branch`（必需）：要扫描的完整 Git 引用。
- `git_search_path`（可选）：帮助 libgit2 找到配置。
- `prefix` 在导入时作用于生成的表名。
- 提交字段包括 `sha1`、`message`、作者 `name`/`email`、`commit_date`、`insertions`、`deletions`、`files_changed`。

### 注意事项

PostgreSQL 服务器账号必须能访问仓库文件系统；FDW 不会克隆远程 URL，也不通过用户映射获取凭据。应把仓库内容和提交元数据视为服务器端数据，只向获准读取它们的角色开放外部表。

`IMPORT FOREIGN SCHEMA` 不支持 `LIMIT TO` 或 `EXCEPT`。版本 `1.1.0` 面向最高 libgit2 0.27 和旧版 PostgreSQL，因此必须针对实际软件包验证当前 libgit2/PostgreSQL 兼容性。查询会遍历仓库历史并计算 diff 统计；对于大型历史应使用限制条件并测量成本。
