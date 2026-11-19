## 用法

来源：

- [Upstream source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/README.md)
- [Upstream source](https://api.pgxn.org/dist/pgvectorfast/0.2.1/META.json)

- [官方来源](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/README.md)
- [官方来源](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/calico.control)
- [官方来源](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/sql/calico--0.2.sql)
- [官方来源](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/LICENSE)
- [官方来源](https://api.pgxn.org/dist/pgvectorfast/0.2.0/META.json)

`calico` 是 pgvectorfast 0.2.1 分发中的缓冲区管理扩展，SQL 版本为 0.2。激活后会替换整个 PostgreSQL 集群的共享缓冲区管理器，影响普通关系及所有数据库。这是一项实验性运行配置变更，不保证所有工作负载都获得性能提升。

### 启用与检查

使用 Linux x86-64 与 PostgreSQL 18 或 19，并提供匹配的非 LTO 服务端符号或调试符号。运行时必须保留 Capstone。将库加入现有预加载列表并重启，然后以超级用户执行以下操作。激活调用必须位于事务块之外。

```ini
shared_preload_libraries = 'vectorfast'
```

```sql
CREATE EXTENSION calico;
SELECT calico.status();
CALL calico.activate();
SELECT * FROM calico.pool;
```

### 接口与作用范围

`calico.status()` 返回当前模式。`calico.activate(timeout_ms)` 默认等待 30000 毫秒，已撤销公共执行权限。`calico.pool_statistics()` 与 `calico.pool` 展示缓冲区映射及统计。激活状态持续到服务端重启，仅创建 SQL 对象不会启用替换。

该分发还提供已有扩展 `vector` 的兼容类型和索引；它不是 `calico` 控制文件声明的依赖。库名为 `vectorfast`。`calico.optimistic_reads` 启用乐观 HNSW 读取，`hnsw.enable_prefetch` 启用分组预取，两者默认开启。文档中的乐观读取路径适用于 `vector_l2_ops`。不要将拦截能力泛化到任意 PostgreSQL 构建、LTO 二进制、其他操作系统或架构。

### 分发版本 0.2.1

0.2.1 分发修复了 VACUUM 从乐观读取回退时的私有缓冲区 pin 预留问题。calico 的 SQL 版本仍为 0.2；应部署匹配的 vectorfast 库，并遵循预加载、重启和显式激活流程。独立的 vectorfast 扩展有自己的身份，依赖 vector 与 calico。
