## 用法

来源：

- [README.md](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/README.md)
- [pg_cuvs.control](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/pg_cuvs.control)
- [docs/playbooks/release-upgrade.md](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/docs/playbooks/release-upgrade.md)
- [sql/pg_cuvs--0.9.0--0.9.1.sql](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/sql/pg_cuvs--0.9.0--0.9.1.sql)
- [sql/pg_cuvs--0.9.1--0.9.2.sql](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/sql/pg_cuvs--0.9.1--0.9.2.sql)
- [Release](https://github.com/pg-cuvs/pg_cuvs/releases/tag/v0.9.2)

`pg_cuvs` 0.9.2 使用独立的 NVIDIA cuVS 守护进程执行 GPU 向量检索与 HNSW 构建加速。文档目标为 Linux/PostgreSQL 16，要求 pgvector 0.7.0 或更新版本。扩展、守护进程、CUDA/cuVS 库及所需 SDK 补丁必须匹配；CPU CI 通过不能证明 GPU 正确性。

### 启用与查询

将 `pg_cuvs` 加入现有预加载列表，保留其他库，并为后端与守护进程配置一致的路径：

```ini
shared_preload_libraries = 'pg_cuvs'
cuvs.socket_path = '/var/run/postgresql/.s.pg_cuvs'
cuvs.index_dir = '/var/lib/postgresql/cuvs_indexes'
```

重启 PostgreSQL，启动可访问相同套接字与持久索引目录的匹配 `pg_cuvs_server`，等待其开始监听，再由有权限的管理员启用不可迁移的扩展：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_cuvs;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
INSERT INTO items VALUES (1, '[1,0,0]'), (2, '[0,1,0]'), (3, '[0,0,1]');
CREATE INDEX items_cagra ON items USING cagra (embedding vector_l2_ops);
SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
EXPLAIN SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
```

`cagra` 提供近似 GPU 检索，`flat` 提供穷举 GPU 检索，`ivfpq` 提供压缩近似检索。`pg_cuvs_hnsw` 使用 GPU 构建，在 CPU 上查询标准 pgvector HNSW 页面。`cuvs.k` 独立于 SQL LIMIT 控制候选数量；过滤和近似检索需要针对实际负载验证召回率。用 EXPLAIN 确认实际路径。

### 运维与升级

`pg_stat_gpu_search` 提供检索统计。`pg_cuvs_hnsw_connectivity_status(regclass)` 返回可选的构建时结构快照，不跟踪后续 INSERT/DELETE/VACUUM，也不保证召回率。旧的两步导入 API 已移除，当前构建应使用 `CREATE INDEX ... USING pg_cuvs_hnsw`。

安装两种匹配的二进制，先重启 PostgreSQL，再重启守护进程，随后在每个数据库执行 `ALTER EXTENSION pg_cuvs UPDATE TO '0.9.2'`。仓库包含从 0.1.0 开始的 SQL 升级链。0.9.2 不改变 SQL 对象，修复并发微批次下的套接字重试及守护进程监听积压问题。真实 GPU 证据针对特定 A100/PG16 环境，不覆盖所有 GPU 或驱动。

### 故障与持久化边界

规划完成后守护进程失联会报错，执行中的语句不会自动切换 CPU 节点。重试前应回滚并检查新执行计划。`enable_cuvs=off` 可禁用 GPU 规划路径。GPU 辅助文件不在 PostgreSQL WAL 内，需要独立的备份、恢复与副本流程。IVF-PQ 不增量索引插入行，数据变化后应重建并核验辅助文件恢复。故障期间的陈旧标记投递仍为尽力而为，内存预检也不会替其他进程预留内存。项目仍在维护，某个辅助函数弃用不代表整个扩展弃用。
