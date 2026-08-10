## 用法

来源：

- [PGXN 上的 pg_local_cache 1.2.0](https://pgxn.org/dist/pg_local_cache/1.2.0/)
- [pg_local_cache v1.2.0 README](https://github.com/profundium/pg_local_cache/blob/v1.2.0/README.md)
- [pg_local_cache v1.2.0 控制文件](https://github.com/profundium/pg_local_cache/blob/v1.2.0/pg_local_cache.control)
- [pg_local_cache 1.2.0 扩展 SQL](https://github.com/profundium/pg_local_cache/blob/v1.2.0/sql/pg_local_cache--1.2.0.sql)
- [技术参考](https://github.com/profundium/pg_local_cache/blob/v1.2.0/docs/TECHNICAL.md)
- [现有服务器安装指南](https://github.com/profundium/pg_local_cache/blob/v1.2.0/docs/INSTALL_EXISTING.md)
- [pg_local_cache v1.2.0 发行版](https://github.com/profundium/pg_local_cache/releases/tag/v1.2.0)

`pg_local_cache` 1.2.0 是一个面向重复主键读取的事务感知 PostgreSQL 进程内缓存。它在共享内存中保存有界的整行条目，可以透明加速符合条件的普通 `SELECT`，同时始终保留原始 PostgreSQL 主键计划作为权威回退路径。它适合单个可写主库上的热点工作集；不是通用查询结果缓存、持久化层，也不是分布式 Redis/Valkey 替代品。

### 核心流程

该动态库必须在 postmaster 启动时加载。下面的 SQL-only 配置关闭可选 RESP 监听器，并为一个应用数据库提供服务：

```conf
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.port = 0
pg_local_cache.cache_entries = 16384
pg_local_cache.memory_budget_mb = 384
```

应把 `pg_local_cache` 追加到现有逗号分隔的预加载列表中，而不是覆盖其他库；验证配置后执行一次受控 PostgreSQL 重启。控制文件把扩展固定在 `local_cache` 模式，设置 `superuser=true`，并声明不可重定位，因此必须由超级用户在每个需要使用它的数据库中创建：

```sql
CREATE EXTENSION pg_local_cache;
```

创建符合要求的永久表，然后将其附加到缓存。`attach_table` 会取得 `ShareRowExclusiveLock`，把完整主键记录到 `local_cache.mapping`，安装扩展自有的失效触发器，并把映射发布到共享内存。在线系统应使用有界的锁等待时间：

```sql
CREATE TABLE public.items (
    id bigint PRIMARY KEY,
    value text NOT NULL,
    enabled boolean NOT NULL DEFAULT true,
    metadata jsonb
);

INSERT INTO public.items VALUES
    (1, 'hello', true, '{"source":"postgres"}');

BEGIN;
SET LOCAL lock_timeout = '2s';
SELECT local_cache.attach_table('public.items'::regclass);
COMMIT;
```

默认的 `p_writable=false` 只会禁止 RESP `SET` 与 `DEL`，不会妨碍普通 PostgreSQL DML。应用继续使用原有 PostgreSQL 连接、行类型和 SQL：

```sql
SELECT * FROM public.items WHERE id = $1::bigint;

SELECT value, metadata
FROM public.items
WHERE id = ANY($1::bigint[]);

EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM public.items WHERE id = 1;
```

符合条件的计划会显示 `Custom Scan (pg_local_cache_sql)`。缓存未命中，或遇到任何不安全、不支持的条件时，会执行保留的主键索引计划；PostgreSQL 始终是事实来源。

### 显式 JSON API

普通 SQL 是返回原生元组的标准接口。确实需要缓存式 JSON 接口的调用方可以使用以下 `SECURITY INVOKER` 函数：

```sql
SELECT local_cache.get('public.items'::regclass, 1::bigint);

SELECT local_cache.mget(
    'public.items'::regclass,
    ARRAY[1, 7, 1]::bigint[]
);
```

`get(regclass, anyelement)` 以 `text` 返回完整行 JSON；`mget(regclass, anyarray)` 返回与输入位置对齐的 `text[]`，保留重复项与 `NULL` 位置。对于复合或异构主键，可调用 `get(regclass, text[])`，各分量顺序必须与 `attach_table` 记录的主键顺序一致。显式 API 调用者需要 `local_cache` 模式的 `USAGE`、对应重载函数的 `EXECUTE`，以及源表上的常规 `SELECT` 权限。

### 重要对象与控制项

- `local_cache.attach_table(regclass, boolean, text)` 验证并注册表。只有可选 RESP worker 需要写源表时才设置 `p_writable=true`；`p_namespace` 可覆盖自动生成的映射名称。
- `local_cache.detach_table(regclass)` 删除映射、托管触发器、共享条目状态和 worker 角色的直接权限；关系未附加时返回 `false`。
- `local_cache.reconcile_table(regclass)` 与 `local_cache.reconcile_all()` 在受控 DDL 或权限变更后重新验证关系形状、主键、触发器来源与 worker 授权。
- `local_cache.mapping` 是扩展自有的映射注册表，并会进入扩展配置转储。不要把直接修改它当成管理函数的替代方案。
- `local_cache.metrics()` 返回有类型的计数器以及内存、worker 指标，`local_cache.health()` 返回精简的 JSON 就绪状态，`local_cache.stats()` 返回详细 JSON 诊断。这些函数与管理函数都已撤销 `PUBLIC` 权限，只应授予指定的部署或监控角色。
- `local_cache.invalidate(namespace)` 使一个映射命名空间失效，并返回受影响的条目数。普通 DML、`TRUNCATE` 与相关 DDL 会自动执行事务性感知的失效处理。

关键设置如下：

| 设置 | 默认值 | 作用 |
|---|---:|---|
| `pg_local_cache.port` | `6380` | RESP2 端口；SQL-only 模式应设为 `0`。 |
| `pg_local_cache.database` | `postgres` | 一个扩展实例所服务的唯一数据库。 |
| `pg_local_cache.cache_entries` | `16384` | 共享行缓存条目容量。 |
| `pg_local_cache.relation_states` | `1024` | 共享关系版本状态容量。 |
| `pg_local_cache.memory_budget_mb` | `384` | 确定性扩展内存分配的启动预算。 |
| `pg_local_cache.max_dirty_keys` | `4096` | 单事务逐键失效上限；超过后扩大为整关系失效。 |
| `pg_local_cache.sql_cache` | `on` | 普通 SQL 快速路径的 `USERSET` 开关；修改无需重启。 |

除 `pg_local_cache.sql_cache` 外，文档列出的 GUC 都是 postmaster 级设置。内存预算覆盖扩展确定性的共享哈希和可选 RESP 缓冲区，不包括 `shared_buffers`、后端内存、操作系统或其他服务。

### 快速路径与一致性边界

透明路径有意保持严格。它要求 `READ COMMITTED`、单个已附加基表、直接列投影，以及完整主键每一列的等值谓词。单列主键还支持有界的 `IN` 与 `= ANY(array)`。连接、CTE、子查询、聚合、分组、排序、行锁、额外谓词、恢复模式、并行执行、`REPEATABLE READ` 和 `SERIALIZABLE` 都会使用普通 PostgreSQL 计划。标量读取可以没有 `LIMIT`，或使用常量 `LIMIT 1`；批量读取不能使用 `LIMIT`。

对于 `IN`/`ANY`，执行器最多去重 1024 个非空键，并把最多 16 MiB 的已验证行复制到查询本地内存。批处理是全有或全无：任意一次未命中、不安全快照、损坏条目或预算溢出，都会执行完整源计划，而不会混合缓存行和源表行。

源表写入仍是普通 PostgreSQL 事务。托管触发器收集发生变化的键，预提交回调会在事务变得可见之前发布失效栅栏。回滚不会发布未提交的行数据。当前事务写过已附加关系后，该事务后续读取会绕过缓存，以保留读己之写语义。发生这种写入后再执行 `PREPARE TRANSACTION` 会被拒绝。

条目没有 TTL，会一直保留到失效、淘汰、替换、损坏检测或 MVCC 安全检查将其退役。编码后的缓存值上限为 8 KiB；更宽的行只会回退 PostgreSQL，不会成为缓存条目。

### 表与部署要求

版本 1.2.0 的已发布 Linux amd64 构建支持 PostgreSQL 14–18、一个配置数据库和一个可写主库。附加对象必须是没有继承和行级安全策略的永久非分区表，并使用即时、非部分的 B-tree 主键。支持的键列类型包括 `int2`、`int4`、`int8`、`text`、`varchar`、`bpchar` 和 `uuid`；复合主键可包含 1–16 列。临时表、无日志表、视图、分区表、表达式或部分主键、不确定性排序规则，以及非默认主键操作符类都会被拒绝。

每个实例最多发布 128 个映射。删除表会遗忘其映射；用相同名称重建表不会自动重新附加。备库不提供缓存服务，也不支持多主协调、TTL、集群、Pub/Sub，或通用范围、连接、聚合缓存。

### 可选 RESP2 安全边界

RESP 模式通过有限的 RESP2 协议提供整行 `GET`、`SET` 和 `DEL`，但所有已接受的映射共享一个令牌和一个 `LOGIN NOSUPERUSER NOINHERIT` worker 角色。它不提供 TLS，也没有逐客户端 PostgreSQL 身份或 ACL 上下文。除非确实需要该接口，否则应保持 `pg_local_cache.port=0`。启用时，应保留默认回环地址，或把远程访问置于网络隔离和认证 TLS 之后；通过 `pg_local_cache.auth_token_file` 使用 PostgreSQL 操作系统用户所有、权限为 `0400` 或 `0600` 的令牌文件。写操作响应丢失不能证明对应 PostgreSQL 事务没有提交。
