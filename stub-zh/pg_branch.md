## 用法

来源：

- [官方 pre-alpha 文档](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/README.md)
- [ProcessUtility hook 实现](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/src/hooks.rs)
- [扩展控制文件](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/pg_branch.control)

`pg_branch` 是一个 pre-alpha 概念验证项目，通过用 BTRFS 写时复制快照取代 PostgreSQL 的常规文件复制，加速 `CREATE DATABASE ... WITH TEMPLATE`。它面向可丢弃的 Linux 测试集群，而非生产数据库。

### 核心流程

安装扩展之前，将集群放在 BTRFS 上，并按上游初始化流程把所需数据目录子目录转换为 BTRFS 子卷。然后在管理会话中加载 hook，并根据模板创建数据库：

```sql
CREATE EXTENSION pg_branch;

CREATE DATABASE branch_db
WITH TEMPLATE source_db;
```

创建扩展会加载后端局部的 `ProcessUtility` hook；上游没有说明 `shared_preload_libraries`。应从已加载扩展库的会话执行命令，并独立验证快照布局。

### 安全边界

版本 `0.0.1` 会调用 BTRFS 文件系统操作，并改变数据库文件创建语义。README 明确说明该项目是概念验证且尚未准备好用于生产，并把更广泛的选项和文件系统支持列为未来工作。

布局不匹配、不支持的 `CREATE DATABASE` 选项、并发会话、崩溃、表空间、备份工具或 PostgreSQL 大版本变更都可能使其假设失效并损坏集群。只能对可丢弃数据使用，并预先验证可恢复备份、挂载选项、子卷、所有权、崩溃清理、表空间和模板连接行为。不要根据项目名推断它支持 ZFS、XFS、集群分支或生产故障转移。
