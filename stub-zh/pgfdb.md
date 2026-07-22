## 用法

来源：

- [官方 README 与限制](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/README.md)
- [表访问方法实现](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/tam/mod.rs)
- [索引访问方法实现](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/iam/mod.rs)
- [FoundationDB 健康检查函数](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/health.rs)

`pgfdb` 0.0.2 是实验性 pgrx 扩展，实现由 FoundationDB 支撑的 PostgreSQL 表和索引访问方法。它是分布式存储与事务研究原型，不是生产中间件。上游明确警告，即使已有生产 FoundationDB 集群也不要让它连接。

### 支持的试用流程

上游以 Docker 镜像分发项目，并要求把 FoundationDB 7.3 集群文件挂载到容器中。预制环境启动后，先验证连接，再显式选择自定义访问方法：

```sql
CREATE EXTENSION pgfdb;

SELECT fdb_is_healthy();

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(),
    name text
) USING pgfdb_table;

CREATE INDEX users_id_fdb ON users USING pgfdb (id);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');

SET enable_seqscan = off;
SELECT * FROM users WHERE id = '00000000-0000-0000-0000-000000000000';
```

`pgfdb_table` 是表访问方法。`pgfdb` 是当前源码中的索引访问方法，为有限的整数、浮点数、文本和 UUID 比较提供操作符类。README 中仍有较旧的 `pgfdb_idx` 拼写；固定版本实现注册的是 `pgfdb`，因此应检查实际镜像中存在的对象。

### 原型限制

DDL 与 PostgreSQL 系统目录不会持久化到 FoundationDB。即使 FoundationDB 仍保留数据，新 PostgreSQL 实例也会丢失模式，因此此版本尚不能实现所宣传的无状态多节点模型。代价估算不完整，规划器可能不会选择自定义索引，测试时需要使用 `enable_seqscan=off` 之类设置。

FoundationDB 的五秒读写事务限制也会约束 PostgreSQL 事务，因此不适合长语句和 OLAP 负载。自定义索引访问方法不支持主键约束，只有有限数据类型可以建立索引。上游还明确说明当前性能不具代表性。

只能使用一次性数据与隔离的 FoundationDB 集群。扩展会在每个加载它的后端中初始化 FoundationDB 客户端和事务回调；仅安装软件包并不足够，还需要有效集群文件与兼容的 FoundationDB 原生库。应把 Docker 镜像、服务器大版本、扩展源码修订和故障恢复作为一个紧密耦合的测试环境验证。
