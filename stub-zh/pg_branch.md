## 用法

来源：

- [上游 pre-alpha 文档](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/README.md)
- [ProcessUtility hook 实现](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/src/hooks.rs)
- [扩展控制文件](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/pg_branch.control)

`pg_branch` 是 pre-alpha 概念验证项目，通过 BTRFS 写时复制快照加速 `CREATE DATABASE ... WITH TEMPLATE`。它面向 Linux 集群，并要求事先把数据目录转换成所需的 BTRFS 子卷布局。

```sql
CREATE EXTENSION pg_branch;
CREATE DATABASE branch_db WITH TEMPLATE source_db;
```

创建扩展会加载后端局部 utility hook；上游没有说明共享预加载配置。该工具会改变数据库文件创建语义并调用文件系统操作，缺陷或布局不匹配可能损坏集群。只能在一次性数据上使用，并先验证备份、子卷、挂载选项、并发会话、崩溃清理、表空间和 PostgreSQL 大版本兼容性。它不适合生产。
