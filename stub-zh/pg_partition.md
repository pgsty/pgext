## 用法

来源：

- [官方扩展控制文件](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/pg_partition.control)
- [官方扩展 SQL](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/sql/pg_partition--0.2.0.sql)
- [官方用法文档](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/doc/pg_partition.md)

`pg_partition` 0.2.0 用于生成声明式分区出现前的继承式时间分区 SQL。给定父表、一个 `timestamptz` 列、UTC 范围和固定间隔后，它会输出子表 DDL、搬移现有行的语句、每个子表的一个索引，以及基于触发器或规则的插入路由。它不会创建原生声明式分区，也不会持续维护分区。

### 核心流程

除非会话时区为 UTC、分区键是 `timestamp with time zone` 且起点是 UTC 午夜，否则生成器会拒绝运行。应先生成脚本、人工检查，再在受控事务中单独执行。

```sql
SET TIME ZONE 'UTC';
CREATE EXTENSION pg_partition;

CREATE TABLE public.events (
    event_id bigint GENERATED ALWAYS AS IDENTITY,
    happened_at timestamptz NOT NULL,
    payload jsonb
);

SELECT create_partitions_trigger_when(
    'public',
    'events',
    'happened_at',
    '2026-01-01 00:00:00+00'::timestamptz,
    '2026-02-01 00:00:00+00'::timestamptz,
    interval '1 day'
);
```

应在不带标题和格式的情况下捕获返回行，检查顺序与标识符，然后只执行一次。`create_partitions_rule` 会生成相同的子表与行搬移语句，但使用规则而不是每个范围一个触发器来路由插入。

### 生成函数

- `validate_inputs` 强制 UTC、分区键类型、午夜起点和允许的间隔。
- `unroll_partition_spec` 把请求的左闭右开时间范围展开成子表规格和后缀。
- `create_partitions_trigger_when` 输出子表、数据搬移、索引、触发器函数与触发器。
- `create_partitions_rule` 输出子表、数据搬移、索引与插入规则。
- 底层的 `create_partition_table`、`move_to_partition`、`create_partition_index`、`create_trigger_function`、`create_trigger` 和 `create_rule` 函数分别以文本返回一段 DDL。

支持的间隔严格限于一年、六个月、三个月、一个月、一周、一天、一小时、一分钟或一秒。

### 运维边界

生成的子表使用 `INHERITS` 和 `CHECK` 约束，并未通过 PostgreSQL 声明式 `PARTITION BY` 机制挂接，因此原生分区裁剪、默认分区、分区挂接与分离以及现代分区管理工具的行为并不相同。

生成脚本会只从父表删除匹配行，然后插入各子表。应在受控写入窗口执行，或显式设计锁定；并发插入与失败可能使行留在意外位置。生成范围之外的行仍在父表中，而且扩展不会自动创建未来分区。

项目源码来自 2014 年，其简短用法文档提到了并不存在的 `create_partitions` 包装函数；应以 0.2.0 实际 SQL 中上述两个函数为准。需要在精确支持的 PostgreSQL 版本上测试生成 DDL、查询计划、触发器搜索路径、备份恢复以及向原生分区的迁移。
