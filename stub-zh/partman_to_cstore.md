## 用法

来源：

- [官方 README](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/README.md)
- [扩展控制文件](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/partman_to_cstore.control)
- [分区迁移函数](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/sql/functions/part_to_cstore.sql)

`partman_to_cstore` 是一个实验性归档桥接扩展，用于旧的 `pg_partman` 时间分区。它把堆分区复制到 `cstore_fdw` 外部表，重建继承、检查约束、所有者与授权，然后删除原始堆表。随后还可依据另一个时间间隔删除过期的 cstore 分区。

### 核心流程

安装两个声明的依赖，并把不可重定位的扩展放到专用模式中：

```sql
CREATE EXTENSION cstore_fdw;
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman WITH SCHEMA partman;
CREATE SCHEMA partman_to_cstore;
CREATE EXTENSION partman_to_cstore WITH SCHEMA partman_to_cstore;

INSERT INTO partman_to_cstore.move_config
    (parent_table, move_int, drop_int, compression)
VALUES
    ('archive.events', '7 days', '180 days', 'pglz');

SELECT partman_to_cstore.part_to_cstore(
    p_parent_table := 'archive.events',
    move_int := '7 days',
    drop_int := '180 days'
);
```

只有在核验候选分区后，才应通过受控调度器运行该函数。`move_singlepart_to_cstore(part_schema, part_table, ...)` 会对单个继承分区执行同样的破坏性转换。

### 对象索引

- `move_config` 保存各父表、迁移年龄、保留年龄、压缩方式、条带/块大小和最后检查时间。
- `part_to_cstore(...)` 发现符合条件的 `pg_partman` 子表，创建 cstore 外部表、复制数据并删除堆表。
- `move_singlepart_to_cstore(...)` 转换单个继承子表。
- `partman_to_cstore_server` 是扩展创建的 `cstore_fdw` 外部服务器。
- `partman_to_cstore_move_data` 是随附的面向 cron 的 Python 驱动程序。

### 运维说明

版本 `1.1.0` 依赖 `pg_partman` 和 `cstore_fdw`，不可重定位，且未声明预加载要求。上游仅记录了对基于时间的 `pg_partman` 分区的支持，其中包括基于 epoch 的命名。SQL 假定旧版 `pg_partman` 目录与函数形态，调度脚本则使用 Python 2 语法；用于当前版本前必须测试兼容性。

转换会复制行，然后对原始堆表执行 `DROP TABLE`。上游警告指出，`pg_dump` 会捕获 cstore 外部表的模式，但不会包含其中的数据。安排迁移前应验证行数、约束、授权、查询行为、物理备份与恢复，并为归档数据保留一份可独立恢复的副本。
