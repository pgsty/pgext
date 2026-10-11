## 用法

来源：

- [lagodb-connectors/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-connectors/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-connectors/lagodb_connectors.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/lagodb_connectors.control)

`lagodb_connectors` 0.1.0 通过外部表和 COPY 读写文本、CSV、JSON、Avro 与 Parquet。外部服务器指定存储提供方，用户映射保存凭据。单个对象对应的表只读；目录前缀表支持追加插入，不支持逐行更新或删除。本地原生 COPY 使用 PostgreSQL 服务端文件权限。

### 核心用法

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
CREATE EXTENSION lagodb_connectors;
CREATE SERVER local_files FOREIGN DATA WRAPPER lagodb_connectors;
CREATE FOREIGN TABLE local_events ()
SERVER local_files OPTIONS (path '/tmp/events.parquet');
SELECT * FROM local_events;
```

### 运行边界

上游将其标记为开发软件，不建议用于生产。文档目标为 PostgreSQL 17。扩展由超级用户安装，须先预加载基础运行时并配置提供方库，然后重启数据库；更改提供方库同样需要重启。S3 有上游端到端覆盖，GCS 与 Azure 仍为实验功能。Delta Lake 支持仅存在于另一个框架骨架中，这些扩展不提供该功能。
