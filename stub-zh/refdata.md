## 用法

来源：

- [Advanced Storage Pack 官方文档](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`refdata` 是 EDB table access method，适用于高频读取、极少修改的 mostly static reference table。

### 启用

安装 Advanced Storage Pack 1.7.0，预加载 `refdata`、重启并创建扩展：

```ini
shared_preload_libraries = 'refdata'
```

```sql
CREATE EXTENSION refdata;
```

### 创建参考数据

```sql
CREATE TABLE market_symbol (
  symbol_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  symbol    text NOT NULL,
  name      text NOT NULL
) USING refdata;

CREATE TABLE trade (
  symbol_id integer NOT NULL REFERENCES market_symbol(symbol_id),
  traded_at timestamptz NOT NULL,
  price     float8 NOT NULL
);
```

Foreign-key reader 不会在 reference table 中创建 row lock。直接修改 Refdata table 则会获取 table-level `ExclusiveLock`。

### 并发边界

只有在变更很少且能够容忍串行化时才应使用 Refdata。写入会阻塞该表以及引用表的并发修改，因此批量更新参考数据需要显式维护计划。把现有 heap table 迁移前，应使用准确的 EDB/PostgreSQL 版本测试 DDL、vacuum、foreign key、logical replication、backup/restore 与 failover。

