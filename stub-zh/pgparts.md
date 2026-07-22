## 用法

来源：

- [官方 README](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/README.md)
- [官方扩展 SQL](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/sql/pgparts.sql)
- [官方扩展控制文件](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/pgparts.control)

`pgparts` 0.0.15 是用于传统 inheritance-based 表分区的 PL/pgSQL 工具包。它把基表的约束、索引、trigger、owner 和权限复制到子表，并维护路由 trigger 与元数据。它适合有意保留这种设计的数据库，不应替代 PostgreSQL 当前的 declarative partitioning。

### 核心流程

把不可迁移的扩展安装到专用 schema，准备基表，创建覆盖某个值的分区，然后插入数据：

```sql
CREATE SCHEMA parts;
CREATE EXTENSION pgparts WITH SCHEMA parts;

CREATE TABLE events (
  id bigint PRIMARY KEY,
  created_at timestamptz NOT NULL,
  payload jsonb
);

SELECT parts.setup(
  'events'::regclass,
  'created_at',
  'monthly'
);

SELECT parts.create_for(
  'events'::regclass,
  '2026-07-22 12:00:00+00'
);

INSERT INTO events VALUES (
  1,
  '2026-07-22 12:00:00+00',
  '{"event":"created"}'::jsonb
);

SELECT * FROM parts.existing_partition;
```

`parts.setup` 记录分区键与策略，并生成 insert routing function 和 trigger。对于已经存在的范围，`parts.create_for` 是幂等的，并以 `regclass` 返回子表。

### 重要对象

- `setup(regclass, name, name, params)` 准备表；第四个参数默认是空 `params` 值。
- `create_for(regclass, text)` 创建覆盖示例 key value 的分区。
- `partition_for(regclass, text)`、`info(regclass, text)`、`start_for(regclass, text)` 和 `end_for(regclass, text)` 检查路由与边界。
- `detach_for(regclass, text)` 和 `attach_for(regclass, text)` 修改 inheritance，但不删除子表。
- `create_archive(regclass)`、`archive_before(regclass, timestamptz)`、`archive_partition(regclass)` 和 `unarchive_partition(regclass)` 管理归档子表。
- `partitioned_table`、`partition` 和 `existing_partition` 暴露配置与实时子表元数据。

### 策略与参数

内置 `monthly` 和 `daily` 策略支持 date、timestamp、timestamp with time zone 及其他已配置类型。参数包括 `nmonths` 或 `ndays`、边界 `timezone`、`drop_old`、`on_conflict_drop` 以及信息性的 `retention`；`start_dow` 用于按日/周分组。可检查 `parts.partition_schema` 和 `parts.schema_param` 获取准确的可接受组合与说明。

### 运维说明

路由依赖生成的 trigger 与表 inheritance。基表用作蓝图，但后续 DDL 变更并不保证自动传播到所有既有子表；必须有计划地管理 schema 变更。路由 trigger 从最新分区向最旧分区检查，因此随着子表数量增加，随机历史写入会变慢。归档操作改变 inheritance 而不是移动行，扩展也不会执行 `retention`。应在目标 PostgreSQL 版本上测试唯一性、外键、update routing、备份恢复、权限、DDL 传播和并发分区创建。
