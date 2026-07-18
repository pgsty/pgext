## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/README.rst)
- [3.5 版本 SQL 对象](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/pgq_node--3.5.sql)
- [扩展 control 文件](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/pgq_node.control)

`pgq_node` 是级联 PGQ 队列的 SQL 支持基础设施。它依赖 `pgq`，以超级用户安装，固定在 `pg_catalog`，应由配套 PGQ/Londiste 工具管理，而不是作为独立工具库调用。

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_node;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN ('pgq', 'pgq_node');
```

完整的级联队列拓扑、节点注册、worker 状态、故障转移与重新同步，应使用完全匹配的文档和守护进程版本。不要通过临时更新捏造节点元数据；队列位置与拓扑表属于协调协议状态。

已复核 README 只称其为“级联队列支持代码”，目录覆盖止于 PostgreSQL 14。必须审计所有安装到 `pg_catalog` 的对象、security-definer 函数、授权和搜索路径。应测试重复投递、顺序、消费者延迟、节点丢失、脑裂、保留、备份恢复和版本偏差。新系统应使用持续维护的逻辑复制或当前受支持 PGQ 栈。
