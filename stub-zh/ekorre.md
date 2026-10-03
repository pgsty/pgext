## 用法

来源：

- [Official README](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/README.md)
- [Extension control file](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/ekorre.control)
- [Installation SQL](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/ekorre--1.0.0.sql)
- [Build configuration](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/Makefile)

`ekorre` 将数据库服务器本地的 Git 仓库映射为 PostgreSQL 外部表。上游将其标为未发布软件，用于仓库信息查询，不保证谓词下推。

### 基本用法

安装扩展库和 libgit2 后，由管理员创建扩展并导入仓库结构。仓库路径在数据库主机上解析。

```sql
CREATE EXTENSION ekorre;
CREATE SERVER git_server FOREIGN DATA WRAPPER ekorre;
IMPORT FOREIGN SCHEMA git FROM SERVER git_server INTO public
  OPTIONS (repopath '/srv/git/project');
```

### 对象与限制

`ekorre_handler()`、`ekorre_validator(text[], oid)` 实现该包装器。导入通过 `repopath` 选择 Git 仓库，并创建实现中定义的表。control 版本为 `1.0.0`，扩展可迁移。上游未声明预加载要求或 PostgreSQL 大版本支持矩阵。服务器和导入表的权限应只授予有权读取该仓库的角色，查询过滤不能替代主机文件访问控制。
