## 用法

来源：

- [阿里云密文索引指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-encdb-btree-to-facilitate-operations-on-ciphertext-indexes)

`encdb_btree` 1.0.0 是阿里云 ApsaraDB RDS for PostgreSQL 的服务商扩展。它为全密态数据库增加 `enc_btree` 索引访问方法，用于密文列的等值和范围操作。它不是可移植的社区扩展。

### 前提与安装

使用小版本为 20230830 或更高的全密态数据库，并先安装服务商提供的基础 `encdb` 扩展：

```sql
CREATE EXTENSION IF NOT EXISTS encdb;
CREATE EXTENSION encdb_btree;
```

使用新的访问方法创建单列、唯一或多列密文索引：

```sql
CREATE INDEX ON test USING enc_btree (t1);
CREATE UNIQUE INDEX ON test USING enc_btree (t2);
CREATE INDEX ON test USING enc_btree (t1, t2, t3);
```

当服务商的规划器生成兼容计划时，现有 SQL 可以使用这些索引。请在有代表性的加密数据上通过 `EXPLAIN` 确认索引是否被选中。

### 运行限制

`enc_btree` 唯一索引不能支持 `INSERT ... ON CONFLICT`，因为它不支持 speculative insertion；它也与外键约束不兼容。这些限制会影响模式与应用设计，因此在生产环境启用加密前应测试迁移。

仍有依赖的密文索引时不能删除该扩展。`DROP EXTENSION encdb_btree CASCADE` 会删除这些索引，但不会删除对应的表或表数据。执行任何删除或引擎变更前，应清点依赖索引并保留回退方案。可用性、权限、升级和备份遵循托管服务约定，而不是上游 PostgreSQL 打包约定。
