## 用法

来源：

- [官方上游 README](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/README.md)
- [官方扩展控制文件 (pg_kafka.control)](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/pg_kafka.control)
- [官方实现源代码](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/src/lib.rs)

`pg_kafka` — PostgreSQL 外部数据封装器用于 Kafka。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_kafka;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
