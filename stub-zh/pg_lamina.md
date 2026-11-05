## 用法

来源：

- [上游来源 1](https://github.com/justusepostgres/pg_lamina/blob/9ba7509f39b362bc487cc3c5843e03f73769b24a/README.md)
- [上游来源 2](https://github.com/justusepostgres/pg_lamina/blob/9ba7509f39b362bc487cc3c5843e03f73769b24a/pg_lamina.control)

`pg_lamina` 0.2.0 将普通堆表转换为 PostgreSQL 内存储的 Parquet 段，并通过 DataFusion 查询。

### 核心流程

转换成功后会在同一事务内删除源堆表。`analytics.convert_table_sorted` 可先按指定列排序。`analytics.query` 对已注册分析表执行 DataFusion SQL，`analytics.storage_compression` 为后续转换选择压缩方式。

```sql
CREATE EXTENSION pg_lamina;
CREATE TABLE lamina_source (region text, amount bigint);
INSERT INTO lamina_source VALUES ('west', 125), ('west', 200);
BEGIN;
SELECT analytics.convert_table('lamina_source', 'events');
SELECT * FROM analytics.query(
  'SELECT region, sum(amount) AS total FROM events GROUP BY region'
);
COMMIT;
```

### 运行边界

此 PostgreSQL 17 ARM64 研究原型既不是持续同步副本，也不是表访问方法。转换会取得排他锁，源索引、约束、触发器、权限及关系身份不会迁移。切换前应独立保存源数据、核对依赖对象并预留临时空间。SQL API 与存储约定仍处于实验阶段。
