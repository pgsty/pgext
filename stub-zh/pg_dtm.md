## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/README.md)
- [PostgreSQL 补丁](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/xtm.patch)
- [外部事务管理器协议](https://github.com/postgrespro/pg_dtm/blob/0a48845d48c690488b22e136877be51e3f1a52df/dtmd/README.md)

`pg_dtm` 是一个基于快照共享的历史分布式事务管理器。它不是独立扩展：上游要求修补并重建 PostgreSQL、运行外部 `dtmd` 服务、通过 `shared_preload_libraries` 加载 `pg_dtm`，并部署至少两个数据库实例。

```sql
CREATE EXTENSION pg_dtm;
SELECT dtm_begin_transaction();
-- Pass the returned global transaction ID to another node.
SELECT dtm_join_transaction(42);
```

参与会话随后开始、修改并提交本地事务；一个提交可能阻塞到其他参与者完成。正确性依赖修补后的服务器、协调器、每个参与者，以及应用传递全局 ID 的过程。

不要把该扩展加载到原版或当前 PostgreSQL。已复核设计属于研究时代软件，包含侵入式服务器补丁、外部协调器，且没有当前故障或兼容性契约。应优先选择持续维护的分布式事务机制。任何归档评估都必须验证协调器崩溃、后端崩溃、网络分区、超时、重复加入、参与者丢失、故障转移与恢复时的原子性，并明确如何备份和修复疑难事务及协调器状态。
