## 用法

来源：

- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/README.md)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/vectorfast.control)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/LICENSE)
- [官方资料](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/sql/vectorfast--0.2.1.sql)
- [官方资料](https://api.pgxn.org/dist/pgvectorfast/0.2.1/META.json)

vectorfast 0.2.1 将 pgvectorfast 分发中的 vector 扩展与 Calico 连接起来。它实验性地替换整个集群的缓冲区管理器，并提供乐观 HNSW 预取路径，不是单表开关。

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

calico.optimistic_reads 与 hnsw.enable_prefetch 默认开启，文档中的乐观路径针对 vector_l2_ops。0.2.1 修复 VACUUM 回退路径中的私有缓冲区 pin 预留问题。必须针对实际服务端构建验证集群级正确性与恢复，不应推断其支持 LTO、其他 CPU 架构或操作系统。
