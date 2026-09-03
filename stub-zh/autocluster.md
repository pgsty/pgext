## 用法

来源：

- [Advanced Storage Pack 官方文档](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`autocluster` 是 EDB table access method，会在插入新数据时把相同 clustering key 的行放置在相邻位置。

### 启用

安装 Advanced Storage Pack 1.7.0，预加载 `autocluster`、重启并创建扩展：

```ini
shared_preload_libraries = 'autocluster'
```

```sql
CREATE EXTENSION autocluster;
```

### 创建聚簇表

```sql
CREATE TABLE iot (
  thermostat_id        bigint NOT NULL,
  recordtime           time NOT NULL,
  measured_temperature float4
) USING autocluster;

CREATE INDEX ON iot (thermostat_id);

SELECT autocluster.autocluster(
  rel := 'iot'::regclass,
  cols := '{1}',
  max_objects := 10000
);
```

`cols` 指定 clustering-key attribute number。相同 key 的插入会被导向相邻 block，从而减少 key-local 访问模式的 page read。

### 运维边界

Autocluster 是物理表布局，不是 query hint。应根据稳定访问模式选择 key，验证写并发与空间行为，并使用准确 EDB 软件包测试 backup、restore、logical replication 与大版本升级。预加载作用于整个集群，而 `CREATE EXTENSION` 与 TAM table 只属于具体数据库。

