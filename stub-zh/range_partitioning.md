## 用法

来源：

- [官方使用文档](https://github.com/yazun/range_partitioning/blob/34853779f37b25a253b3c6889b6aaafe2ab56968/doc/range_partitioning.md)
- [官方 README](https://github.com/yazun/range_partitioning/blob/34853779f37b25a253b3c6889b6aaafe2ab56968/README.md)
- [PGXN 发行元数据](https://pgxn.org/dist/range_partitioning/)

`range_partitioning` 使用 PostgreSQL 范围类型实现基于触发器的表分区。它把普通表转换为受管理的父表，将插入行路由到子表，并可拆分或合并相邻范围。该项目早于 PostgreSQL 声明式分区，应视为传统替代方案，不应随意与原生 `PARTITION BY` 表混用。

### 核心流程

将可重定位扩展安装到专用模式，只把其管理角色授予分区管理员，并通过指定分区键转换现有表：

```sql
CREATE SCHEMA range_partitioning;
CREATE EXTENSION range_partitioning SCHEMA range_partitioning;

CREATE SCHEMA app;
CREATE TABLE app.events (
  event_id bigint,
  occurred_on date NOT NULL,
  payload jsonb
);

SELECT range_partitioning.create_parent(
  'app.events',
  'occurred_on'
);

SELECT range_partitioning.create_partition(
  'app.events',
  '[2026-01-01,2027-01-01)'
);
```

`create_parent` 会创建元数据、初始无界子表、约束与插入路由触发器。如果存在多个兼容的范围类型，应显式传入限定后的范围类型。`create_partition` 要求新范围完整包含于某个现有分区，并拆分该分区。

### 管理与检查

- `drop_partition` 把一个子表合并到指定的相邻分区，并非简单的分离操作。
- `master` 与 `partition` 保存管理元数据，`master_partition` 提供联接后的报告视图。
- `get_destination_partition` 判断候选键将进入哪个子表。
- `where_clause` 为一个范围或现有分区派生谓词。
- `value_in_range`、`is_subrange`、`range_add`、`range_subtract` 与 `constructor_clause` 操作动态指定的范围值。

每次拆分或合并前后都应检查受管理布局：

```sql
SELECT master_class::regclass AS parent,
       partition_class::regclass AS child,
       range,
       where_clause(partition_class) AS predicate
FROM range_partitioning.partition
ORDER BY 1, 2;
```

### 权限与注意事项

上游 README 称安装会创建名为 `range_partitioning` 的角色；只将其授予确实需要创建或移除分区的用户。子表创建会按照扩展的传统实现携带表结构与权限，因此应在代表性表上验证索引、授权、触发器、所有权以及转储恢复行为。

1.2.2 版要求 PostgreSQL 9.2 或更高版本，发布于 2016 年。控制文件没有声明预加载要求或扩展依赖。现代 PostgreSQL 已提供与规划器和维护流程集成的声明式范围分区，采用这种触发器加元数据设计前应先比较两者。分区变更会移动或重写数据并获取锁，因此必须评估规模并预演。升级应通过 `ALTER EXTENSION` 使用扩展升级脚本，不要手工编辑其元数据表。
