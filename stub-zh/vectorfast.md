## 用法

来源：

- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/CHANGELOG.md)

- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/README.md)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/vectorfast.control)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/LICENSE)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/sql/vectorfast--0.2.1.sql)
- [官方资料](https://api.pgxn.org/dist/pgvectorfast/0.2.2/META.json)

vectorfast 0.2.2 将 pgvectorfast 分发中的 vector 扩展与 Calico 连接起来。它实验性地替换整个集群的缓冲区管理器，并提供乐观 HNSW 预取路径，不是单表开关。

### 启用

要求 Linux x86-64、PostgreSQL 18 或 19、匹配的非 LTO 服务端或调试符号，以及运行时 Capstone。保留其他预加载项并重启服务。

```ini
shared_preload_libraries = 'vectorfast'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION calico;
CREATE EXTENSION vectorfast;
SELECT calico.status();
SELECT * FROM calico.pool;
```

以超级用户身份在显式事务块之外创建扩展：vectorfast 依赖 vector 和 calico，并激活 Calico。激活影响普通关系及所有数据库。后续服务重启后，需要再次在事务块之外调用 calico.activate()；已有 SQL 对象不会自动重新激活。

```sql
CALL calico.activate();
```

### 运行

calico.optimistic_reads 与 hnsw.enable_prefetch 默认开启，0.2.2 的乐观读取覆盖 L2 及内置内积距离，包括归一化余弦检索。0.2.1 修复 VACUUM 回退路径中的私有缓冲区 pin 预留问题。必须针对实际服务端构建验证集群级正确性与恢复，不应推断其支持 LTO、其他 CPU 架构或操作系统。

### 分发版本 0.2.2

0.2.2 将乐观读取扩展到内置内积距离，包括对归一化向量执行 HNSW 余弦检索；此前仅限 L2 的描述已不能涵盖支持范围。此次修复交错扫描的已访问缓存所有权，并在游标恢复时遵守 `hnsw.enable_prefetch`。迭代检索使用紧凑候选队列和批量 I/O。`calico` 的 SQL 版本仍为 0.2，`vectorfast` 的 SQL 版本为 0.2.2。上文的集群级预加载、重启、显式激活和非 LTO 平台要求继续适用。
