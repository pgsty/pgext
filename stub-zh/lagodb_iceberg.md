## 用法

来源：

- [lagodb-iceberg/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-iceberg/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-iceberg/lagodb_iceberg.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/lagodb_iceberg.control)

`lagodb_iceberg` 0.1.0 提供用于受管表的 iceberg 表访问方法，以及访问已有 REST 目录的外部数据包装器。受管表支持 SQL DML、事务、部分模式变更与维护。外部目录中的表仍由外部目录管理。Iceberg v1 不支持逐行更新和删除；不支持可重复读隔离级别，串行化模式也不提供完整的 PostgreSQL SSI 语义。

### 核心用法

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
CREATE EXTENSION lagodb_iceberg;
CREATE TABLE local_events (event_time timestamptz, device_id bigint) USING iceberg;
INSERT INTO local_events VALUES (now(), 101);
SELECT * FROM local_events;
```

### 运行边界

上游将其标记为开发软件，不建议用于生产。文档目标为 PostgreSQL 17。扩展由超级用户安装，须先预加载基础运行时并配置提供方库，然后重启数据库；更改提供方库同样需要重启。S3 有上游端到端覆盖，GCS 与 Azure 仍为实验功能。Delta Lake 支持仅存在于另一个框架骨架中，这些扩展不提供该功能。
