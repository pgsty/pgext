

## 用法

> [pg_clickhouse: PostgreSQL 的 ClickHouse 集成](https://github.com/ClickHouse/pg_clickhouse)

`pg_clickhouse` 允许直接从 PostgreSQL 对 ClickHouse 执行分析查询，无需重写 SQL。支持 PostgreSQL 13+ 和 ClickHouse v23+。

### 创建扩展

```sql
CREATE EXTENSION pg_clickhouse;
```

或安装到指定 schema：

```sql
CREATE SCHEMA env;
CREATE EXTENSION pg_clickhouse SCHEMA env;
```

### 查询下推

该扩展自动将分析查询下推到 ClickHouse 执行，带来显著的性能提升。例如 TPC-H 基准测试显示：

- 查询 1：268ms（标准 PostgreSQL 为 4,693ms）
- 查询 6：53ms（标准 PostgreSQL 为 764ms）

当查询下推激活时，ClickHouse 直接处理执行，避免了复杂分析工作负载的数据传输开销。
