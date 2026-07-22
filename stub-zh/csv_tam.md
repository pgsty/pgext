## 用法

来源：

- [官方 README](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/README.md)
- [表访问方法实现](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam.c)
- [扩展 SQL](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam--0.0.1.sql)

`csv_tam` 0.0.1 是教学原型，通过自定义表访问方法，把关系主分叉存成逗号分隔行。它用于展示 PostgreSQL 的 TableAM 回调机制；它不是事务型 CSV 存储引擎，绝不能用于生产数据。

### 核心流程

安装处理器后，在创建一次性测试表时明确选择 `csv_tam`：

```sql
CREATE EXTENSION csv_tam;

CREATE TABLE tam_demo (
  id integer,
  measured_at timestamptz,
  value numeric
) USING csv_tam;

INSERT INTO tam_demo VALUES (1, now(), 42.5);
SELECT * FROM tam_demo;
```

实现仅序列化 `int2`、`int4`、`int8`、`float4`、`float8`、`date`、`timestamp`、`timestamptz` 和 `numeric`。它没有 CSV 引号或转义机制；NULL 会变成空字段，因而具有歧义。

### 支持范围

顺序插入与扫描是预期的演示路径。更新、删除、截断、索引、清理、分析、并行扫描、重新扫描、采样、推测式插入和关系复制均未实现，会直接报错。规划器行数估计也是固定值，并非来自真实元组统计。

### 运维说明

写入会绕过 MVCC、WAL、崩溃恢复和事务回滚，直接追加到关系文件。因此，失败或回滚的事务仍可能留下字节，并发写入也不安全。若干回调还会输出可能泄露行内容或文件路径的 NOTICE。只能在隔离测试集群中使用一次性且不敏感的数据；绝不能依赖它提供持久性或可见性语义。
