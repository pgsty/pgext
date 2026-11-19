## 用法

来源：

- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/README.md)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/kingres_timepart.control)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/kingres_timepart--0.1.sql)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_timepart` 0.1 通过 SQL 和 PL/pgSQL 管理普通时间范围分区的创建与保留。它没有共享库、预加载要求或后台调度器。同一安装 SQL 面向 PostgreSQL 18/19。

### 定义并检查策略

安装会创建事件触发器，因此需要超级用户。示例属于管理操作，维护运行可能删除过期分区。应先检查计划，并在事务块外调用过程。

```sql
CREATE EXTENSION kingres_timepart;
CREATE TABLE public.events (ts timestamptz, payload text) PARTITION BY RANGE (ts);
SELECT kingres.timepart_set_policy(
    'public.events', '1 day', create_ahead => '7 days', retention => '30 days'
);
SELECT * FROM kingres.timepart_plan('public.events');
CALL kingres.timepart_run('public.events');
```

### 维护与权限

表必须使用单一普通 `date`、`timestamp` 或 `timestamptz` RANGE 分区键，且没有 DEFAULT 分区。获得显式执行授权的所有者可配置策略，并使用 `kingres.create_time_partitions`、`kingres.drop_old_time_partitions`、`kingres.timepart_plan` 与 `kingres.timepart_check`。维护过程仅限超级用户，操作之间会提交；部分失败不会回滚已提交的工作。

`kingres.timepart_policies` 与运行历史视图按所有者过滤状态。没有外部调用者执行时，策略不会自动运行。时间网格默认使用 UTC，并区分日历日与固定小时。不对齐的分区会报告而不自动清理。分布式表、自动调度和列存转换过程均不是本版本已实现的功能，应避免使用保留的列存策略字段。
