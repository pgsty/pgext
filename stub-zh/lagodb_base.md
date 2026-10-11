## 用法

来源：

- [lagodb-base/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-base/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-base/lagodb_base.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/lagodb_base.control)

`lagodb_base` 0.1.0 提供 LagoDB 后台进程管理、存储卷、提供方加载和查询下推，是连接器与 Iceberg 扩展的共同依赖。它分发支持的对象存储 URI COPY 操作，并保留 PostgreSQL 原生查询路径。

### 核心用法

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
```

### 运行边界

上游将其标记为开发软件，不建议用于生产。文档目标为 PostgreSQL 17。扩展由超级用户安装，须先预加载基础运行时并配置提供方库，然后重启数据库；更改提供方库同样需要重启。S3 有上游端到端覆盖，GCS 与 Azure 仍为实验功能。Delta Lake 支持仅存在于另一个框架骨架中，这些扩展不提供该功能。
