## 用法

来源：

- [pgiceberg.control](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/pgiceberg.control)
- [README.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/README.md)
- [sql/pgiceberg--0.1.0.sql](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/sql/pgiceberg--0.1.0.sql)
- [docs/design/postgres-extension-surfaces.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/docs/design/postgres-extension-surfaces.md)
- [docs/design/postgres-iceberg-commit.md](https://github.com/pgiceberg/pgiceberg/blob/4e604340a008f1ab962584fba9360dcc3fb41f5e/docs/design/postgres-iceberg-commit.md)

`pgiceberg` 提供三个独立的 Iceberg 接口：外部表、原生 Iceberg 表和逻辑解码镜像，使用 Apache Iceberg C++ 处理目录与 Parquet 数据。扩展版本为 0.1.0，不是其依赖库版本。

### 核心用法

```sql
CREATE EXTENSION pgiceberg;
SELECT pgiceberg.add_catalog('demo', 'sqlite',
  '/tmp/pgiceberg_demo.db', '/tmp/pgiceberg_demo_warehouse');
CREATE SERVER iceberg_demo FOREIGN DATA WRAPPER pgiceberg OPTIONS (catalog 'demo');
SELECT * FROM pgiceberg.commit_recovery_log();
```

### 运行边界

需要超级用户安装及匹配的 C++ 运行时。目录路径与凭据赋予服务器访问外部存储的能力。支持 `IMPORT FOREIGN SCHEMA`、模式差异／刷新、快照读取及非分区表 UPDATE/DELETE。历史快照拒绝写入，逻辑镜像需单独配置逻辑解码。Iceberg 发布早于 PostgreSQL 提交，并非两阶段事务，崩溃或多表失败可能留下部分外部提交。应检查恢复日志并遵循上游协调修复流程。PostgreSQL 隔离级别不能隔离其他 Iceberg 写入方。
