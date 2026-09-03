## 用法

来源：

- [PGAA 官方文档](https://www.enterprisedb.com/docs/pgaa/1.11/)
- [官方安装指南](https://www.enterprisedb.com/docs/pgaa/1.11/installing/)
- [官方 quickstart](https://www.enterprisedb.com/docs/pgaa/1.11/overview/quick_start/)

`pgaa` 是 EDB Postgres Analytics Accelerator，通过向量化 Seafowl engine 查询 Parquet、Delta Lake 与 Iceberg 数据，同时保持 PostgreSQL SQL 接口的 table access method。

### 启用

PGAA 1.11.0 在文档平台上支持 PostgreSQL、PGE 与 EPAS 16–18。安装匹配软件包，预加载、启动受管 engine、重启，并通过依赖创建扩展：

```ini
shared_preload_libraries = 'pgaa,pgfs'
pgaa.autostart_seafowl = on
```

```sql
CREATE EXTENSION pgaa CASCADE;
```

`CASCADE` 会创建 EDB PGFS。必须确认安装的 `pgfs` 是 EDB provider module，而不是无关的同名扩展。

### 注册存储并查询数据

```sql
SELECT pgfs.create_storage_location(
  'analytics',
  's3://warehouse-bucket',
  '{"region":"us-east-1"}'
);

CREATE TABLE lineitem ()
USING pgaa
WITH (
  pgaa.storage_location = 'analytics',
  pgaa.path = 'tpch/lineitem',
  pgaa.format = 'delta'
);

SELECT count(*) FROM lineitem;
```

PGAA 还可使用 Iceberg catalog，通过 CTAS 写入受支持格式，并把部分 workload offload 到 Spark。

### 运维边界

查询依赖 object-store availability、credential、metadata consistency 与 Seafowl/Spark execution process。应保护 storage-location secret、控制 network egress、监控 worker lifecycle 与 fallback plan，并测试 format/version compatibility。仅有 PostgreSQL backup 无法覆盖外部 lake data；catalog、object storage、PGFS metadata 与数据库恢复必须作为一个系统协调。
