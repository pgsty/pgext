## 用法

来源：

- [阿里云官方在线分区指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-online-migrate-extension-to-partition-tables-online)
- [ApsaraDB RDS for PostgreSQL 官方产品文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_online_migrate` 是阿里云 ApsaraDB RDS for PostgreSQL 扩展，可在持续写入期间将普通表转换为分区表，或对已有分区表重新分区。它复制现有数据、通过逻辑复制流式同步变更，并在同步完成后原子重命名源表和目标表。这是服务商托管功能，不是可移植的社区扩展。

### 前置条件与核心流程

阿里云当前指南要求 RDS PostgreSQL 13 或更高版本及符合要求的小版本内核、`wal_level = logical`、由主键或 `REPLICA IDENTITY` 提供的行标识、创建发布/订阅/复制槽和重命名表的权限、已验证备份，以及至少为源表及索引总大小两倍的可用存储。

在同一模式中创建目标分区表，然后开始转换：

```sql
CREATE EXTENSION rds_online_migrate;

CREATE TABLE public.events_new
    (LIKE public.events INCLUDING ALL)
    PARTITION BY HASH (id);
CREATE TABLE public.events_new_0 PARTITION OF public.events_new
    FOR VALUES WITH (MODULUS 2, REMAINDER 0);
CREATE TABLE public.events_new_1 PARTITION OF public.events_new
    FOR VALUES WITH (MODULUS 2, REMAINDER 1);

SELECT rds_online_migrate.rewrite_table(
    'public.events',
    'public.events_new'
);
```

调用成功返回 `true`。目标表会取得原表名，源表则以 `_rds_bkp` 后缀保留以供验证。

### 迁移后处理与故障恢复

与原表关联的视图、触发器和外键约束不会自动转移到替换表，必须手工重建。若源表属于逻辑复制发布，需要把重命名后的替换表重新加入发布。还应重新应用并验证业务角色权限、比较行数与分区，并在应用验证完成前保留备份表。

该过程并非原子操作。中断可能遗留元数据行、发布、订阅或复制槽，需要按官方指南手工清理。并发任务可能需要增加 `max_worker_processes`。应在非生产 RDS 实例中测试完整转换与恢复路径，并通过服务商当前资格要求与支持流程启用此功能。
