## 用法

来源：

- [上游 README](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/README.md)
- [扩展 control 文件](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/extensions/postgres/orochi.control)
- [1.0 版 SQL 接口](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/extensions/postgres/sql/orochi--1.0.sql)
- [Beta 测试计划](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/docs/BETA_TESTING_PLAN.md)

`orochi` `1.0` 版是覆盖面广的 PostgreSQL 16+ HTAP 项目，涉及自定义向量、分布式表、超表、列式/压缩存储、分层、连续聚合、近似函数、流水线/CDC、工作负载控制、JSON 处理、认证和 Raft 协调。它在固定的 `orochi` 模式安装大量 SQL/C 接口，并通过钩子和后台组件集成。

### 从盘点开始

```sql
CREATE EXTENSION orochi;
SELECT orochi.orochi_version();
SELECT extversion FROM pg_extension WHERE extname = 'orochi';
SELECT proname FROM pg_proc
WHERE pronamespace = 'orochi'::regnamespace ORDER BY proname;
```

不要仅凭 README 启用分片、分层、流水线、认证钩子或分布式 DDL。仓库一方面宣称生产就绪，另一方面 Beta 计划仍把核心扩展检查列为待验证项，文档与 SQL 之间的示例/名称也有漂移。安装仅限超级用户，部分功能会联系 PostgreSQL 节点、S3、Kafka 或 Web 服务。应在可丢弃数据上逐个子系统评估，审计每项外部凭据和生成对象，并在生产前证明崩溃、升级、事务、故障转移与回滚行为。
