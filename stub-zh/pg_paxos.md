## 用法

来源：

- [官方 README](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/README.md)
- [扩展控制文件](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/pg_paxos.control)
- [SQL 实现](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/sql/pg_paxos.sql)
- [规划器与执行器钩子](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/src/pg_paxos.c)

`pg_paxos` 是一个已废弃的实验性 Multi-Paxos 表复制实现，面向 PostgreSQL 9.4–9.6 时代的服务器。它记录 SQL 文本并在组成员上重新执行；只能用于历史研究或隔离实验，不能作为受支持的高可用或生产复制系统。

### 配置与核心流程

在每个节点安装声明的 `dblink` 依赖，预加载 `pg_paxos`，分配唯一节点 ID，并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_paxos'
pg_paxos.node_id = 'node-a'
```

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pg_paxos;

CREATE TABLE coordinates (x integer, y integer);
SELECT paxos_create_group('geo', 'host=10.0.0.1 dbname=postgres');
SELECT paxos_replicate_table('geo', 'coordinates'::regclass);
```

在具有相同表定义与扩展的新节点上，使用现有节点和新节点各自的连接字符串加入组：

```sql
SELECT paxos_join_group(
    'geo',
    'host=10.0.0.1 dbname=postgres',
    'host=10.0.0.2 dbname=postgres'
);
```

如果在组建立后添加其他表，按上游要求在每个节点调用 `paxos_replicate_table(text, regclass)`。

### 配置与一致性

- `pg_paxos.node_id` 必须唯一标识节点。
- `pg_paxos.enabled` 控制是否拦截受管理表。
- `pg_paxos.consistency_model` 默认为 `strong`；`optimistic` 会减少读取协调，但故障后可能返回旧数据。
- `pg_paxos.allow_mutable_functions` 默认为关闭，因为非确定性函数会使副本分歧。

规划器和执行器钩子要求受管理查询引用的每个关系都属于同一个 Paxos 组。写入会转换成 SQL 文本，追加到组日志并重放。复制表上的写入禁止出现在显式事务块中，重放路径也只支持单语句执行。

### 安全与兼容性

连接字符串保存在 `pgp_metadata` 表中，并由 `dblink` 使用；应保护该模式，并尽量避免在其中嵌入长期密码。SQL 日志设计要求确定性行为。序列、触发器、可变函数、时间与随机值、外部副作用和模式漂移都可能导致分歧或重复效果。启用 `pg_paxos.allow_mutable_functions` 只会移除一道检查，并不能使这些操作安全。

版本 `0.2` 来自 2016 年，服务器内部兼容代码只覆盖 PostgreSQL 9.4、9.5 与 9.6。它缺少生产共识系统应具备的运维成熟度、升级路径、监控、恢复流程和现代 PostgreSQL 支持。应把 `strong` 视为该实验实现的预期模式，而不是当前服务保证；只能在一次性环境中独立测试故障、分区、成员变更、崩溃恢复和数据对账。
