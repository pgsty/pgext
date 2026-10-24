## 用法

来源：

- [v1.18.0 README](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/README.md)
- [v1.18.0 extension control](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/pg_mentat.control)
- [v1.18.0 changelog](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/CHANGELOG.md)
- [v1.9.1 dump repair](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/sql/pg_mentat--1.9.0--1.9.1.sql)
- [v1.18.0 SQL aliases](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/sql/07_function_aliases.sql)
- [v1.18.0 history functions](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/time_travel.rs)
- [v1.18.0 excision function](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/excision.rs)
- [v1.18.0 transaction functions](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/transact.rs)
- [v1.18.0 subscriptions](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/subscriptions.rs)
- [1.18.0 upgrade SQL](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/sql/pg_mentat--1.17.2--1.18.0.sql)
- [1.18.0 命名存储指南](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/docs/src/multi-tenancy.md)
- [1.18.0 查询函数](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/query.rs)
- [1.18.0 pull 函数](https://api.pgxn.org/src/pg_mentat/pg_mentat-1.18.0/crates/pg/pg_mentat/src/functions/pull.rs)

`pg_mentat` 1.18.0 在 PostgreSQL 内实现与 Datomic 兼容的数据模型和 Datalog 查询引擎。它将不可变事实存储为有类型的 datom，并通过 SQL 函数提供模式事务、Datalog 查询、pull 表达式、时间旅行、事务历史和永久切除功能。它适用于需要这种模型的应用；并非关系表或 SQL 的透明替代品。

### 安装并定义模式

```sql
CREATE EXTENSION pg_mentat;

SELECT mentat.t('[
  {:db/ident       :person/name
   :db/valueType   :db.type/string
   :db/cardinality :db.cardinality/one}
  {:db/ident       :person/age
   :db/valueType   :db.type/long
   :db/cardinality :db.cardinality/one}
]');
```

推荐使用的便捷别名位于 `mentat` 模式中。新属性必须先通过模式事务写入，随后事实才能使用它们。

### 写入并查询数据

```sql
SELECT mentat.t('[
  {:person/name "Alice" :person/age 30}
  {:person/name "Bob"   :person/age 25}
]');

SELECT mentat.q('
  [:find ?name ?age
   :where [?e :person/name ?name]
          [?e :person/age ?age]
          [(> ?age 28)]]
');
```

`mentat.t(edn)` 执行 ACID 事务并返回事务报告。`mentat.q(query, inputs)` 将 Datalog 查询编译为 PostgreSQL 执行计划。请使用 EDN 参数和输入绑定，不要把应用字符串插入查询文本。

### Pull、历史记录与假设事务

```sql
SELECT mentat.pull('[*]', 1000001);
SELECT public.log('default', 1000000000001, 1000000000010);
SELECT public.diff(
  'default', 1000000000003, 1000000000007,
  '[:find ?name :where [?e :person/name ?name]]',
  '{}'::jsonb
);

SELECT public.mentat_with('[
  {:person/name "Alice" :person/age 31}
]');
```

`mentat.pull` 返回实体形态的 JSON。`public.log` 返回事务区间内的历史，`public.diff` 比较指定 Datalog 查询在两个事务时点的结果，五个参数都必须提供。`public.mentat_with` 评估事务但不持久化。这些辅助函数在 1.18.0 中导出到 public 模式，没有 mentat 模式别名。应使用自己存储中的实体和事务 ID。

1.18.0 还通过三参数重载提供历史 pull：

```sql
SELECT public.edn_pull(
  '[*]', 1000001, '{"asOf": 1000000000003}'::jsonb
);
```

`asOf` 读取某个事务时点的状态；`since` 选择在某个事务之后断言、当前仍存在的事实，而非期间出现过的所有历史值。组合使用时，只返回在指定区间内断言、到上界仍未撤回的事实。选项键区分大小写；`as_of` 等拼写错误现在会报错，不再静默返回当前数据。

永久切除有意与通常的不可变历史机制分开：

```sql
SELECT public.mentat_excise(ARRAY[1000042]::bigint[], 'default', NULL);
```

执行切除前请检查目标实体和备份；该操作会永久移除 datom，适用于隐私擦除等要求。第一个参数是实体 ID 数组。实体所在分区必须允许切除；模式实体受保护，同批切除范围外的实体引用可能阻止操作。

### 重要对象

- `mentat.t(edn)`：写入模式或数据事务。
- `mentat.q(query, inputs)`：执行 Datalog。
- `mentat.pull(pattern, eid)` 和 `mentat.pull_many(pattern, eids)`：以实体形态读取数据。
- `mentat.entity(eid)` 和 `mentat.schema()`：检查实体或当前模式。
- `public.log(...)` 和 `public.diff(...)`：检查事务历史和查询结果变化。
- `mentat.stats()`、`mentat.storage()` 和 `mentat.cache_stats()`：运行状态检查。
- `public.subscribe(...)`：通过 PostgreSQL `LISTEN`/`NOTIFY` 提供响应式查询通知。

该扩展在 `mentat` 模式下的窄表中存储有类型的 datom，包括引用、整数、字符串、布尔、浮点、时刻、关键字、UUID 和字节值。

### 较早版本的安全修复

1.6.2 修复深层 EDN 嵌套导致栈耗尽、普通用户设置受限参数导致查询失败，以及 UTC 解码错误。可选的 `script` 构建特性提供类 Clojure 求值，默认关闭。上游说明所有更早版本（包括 1.5.7）均受 EDN 嵌套缺陷影响：后端崩溃可能断开全部客户端并触发崩溃恢复。升级已有部署时，应检查已安装的扩展版本。

### 要求与注意事项

- 上游 v1.18.0 的 Cargo 特性覆盖 PostgreSQL 13-19，使用 pgrx 0.19.3。Pigsty 的 1.18.0 软件包面向 PostgreSQL 14-18。
- 该扩展不可重定位，也不要求 `shared_preload_libraries`。
- 可选的 `mentatd` HTTP/Datomic 线协议守护进程是上游配套程序，不包含在 Pigsty `pg_mentat` 软件包中。仅通过 SQL 使用扩展并不需要它。
- Datalog 编译、递归 pull、全文属性、订阅和历史记录可能呈现截然不同的成本特征。请使用文档所述的 explain 辅助函数检查生成的 SQL，并在代表性数据上进行基准测试。
- 切除操作绕过通常的不可变历史模型。请限制权限并审计其使用。

### API 与索引

核心函数为 `edn_t`、`edn_q`、`edn_pull` 和可选脚本接口 `edn_eval`；`mentat.t`、`mentat.q` 便捷别名继续可用，较早的 `mentat_transact`、`mentat_query`、`mentat_pull` 和可选的 `mentat_eval` 已弃用。

`edn_q_rows` 每行返回一个 JSONB 数组。1.18.0 分批读取结果并写入 PostgreSQL tuplestore，超过 `work_mem` 后可溢写临时文件，不再将完整结果保留在后端内存中。它仍会先物化结果再返回，并遵守 `mentat.max_result_rows` 限制。

```sql
SELECT row_data->>0 AS person_name
FROM public.edn_q_rows(
  '[:find ?name :where [?e :person/name ?name]]',
  '{}'::jsonb
) AS r(row_data);
```

`mentat.auto_index` 可设为 `off`、`schema`（默认）或 `adaptive`。当前值表会建立 AVET 索引；自适应模式管理历史属性的部分索引，记录在 `mentat.managed_indexes` 中。`mentat_tune_indexes` 默认只报告计划，显式关闭 dry-run 后才应用；仅回收自己管理的索引，默认空闲窗口为 7 天。

聚合改为 Datalog 集合语义：对 1、1、1、3 求和得到 4，而不是 6；需要保留实体多重性时加入 `:with ?e`。升级前应复核依赖旧结果的查询。

### 升级与逻辑备份

1.9.1 修复了一项重要备份缺陷：更早版本未注册扩展成员表的数据，`pg_dump` 的逻辑备份会漏掉 mentat 数据，恢复后存储为空。升级后应立即重新生成并验证逻辑备份；旧备份不会自动得到修复。物理备份不受该缺陷影响。

安装新版文件后，通过升级链更新到 1.18.0。升级会建立索引并可能阻塞写入，应安排维护窗口：

```sql
ALTER EXTENSION pg_mentat UPDATE TO '1.18.0';
```

已部署数据库仍需执行 SQL 扩展升级并验证新的逻辑备份。

### 1.13–1.14 集成与事务函数

可选的 `pg_fts` 与 `pg_turbovec` 集成提供 `mentat.create_fts_index`、`mentat.attach_turbovec_attribute` 和 `mentat.create_turbovec_index`；两者都不是必需依赖。1.14.0 修正向量辅助函数，改用维度 CHECK 约束，并新增 `bm25-match` / `turbovec-near` Datalog 连接。TurboVec 维度必须能被 8 整除。

启用 `script` 的构建可用 `:db/fn` 定义返回事务数据的事务函数，读取的是事务开始时的数据库状态。执行受 `mentat.script_max_*` 和 `mentat.txfn_max_depth`（默认 16）限制。`:db/retractEntity` 现会撤回指向目标实体的引用，但保留引用方实体。EXPLAIN 辅助函数现能正确处理集合、元组和关系输入。升级到 1.14.0 会加入函数属性并替换向量辅助函数，无需数据迁移。

从 1.17.0 起，控制文件使用 `superuser = true` 与 `trusted = true`，修复普通用户可信安装失败的问题。安装者获得模式使用权及表、序列的可转授权权限，公共角色不会自动获得访问权；旧版本中 `superuser = false` 与可信标记组合导致的失败不再适用于新版。

### 1.17.1 的查询与权限边界

1.17.0 修复计算值上的谓词，以及多条件查询中的 `get-else` 和 `missing?`。1.17.1 主要更新独立 DuckDB 后端和构建依赖，不改变 PostgreSQL 行为。可信安装不代表所有管理操作都对普通角色开放，启用行级安全与实际调整索引仍有所有者或超级用户限制。

旧版本的命名存储写入限制已在 1.18.0 修复，升级前应阅读下面的迁移边界。普通用户备份及恢复仍需使用正确的所有者角色，并验证数据与权限。

### 1.17.2 安装修复

1.17.2 修复了 1.17.0–1.17.1 的可信安装提权漏洞：安装脚本中未限定模式的函数调用可能被数据库所有者预先创建的同名函数劫持。新版限定内置函数并固定搜索路径。升级步骤本身没有模式变化，但不能撤销已经发生的恶意安装。如果非超级用户曾在可创建公共模式函数的数据库中安装受影响版本，应审查非系统模式的同名函数与超级用户角色。旧软件包版本不能据此被视为已更新。

### 1.18.0 命名存储与历史读取

`create_store` 现在创建与默认存储具有同样表布局的可写命名存储。应使用 public 模式下带存储参数的函数；mentat 模式中的便捷别名仍访问默认存储。

```sql
SELECT public.create_store('tenant_demo');
SELECT public.t('tenant_demo', '[
  {:db/ident :sensor/value
   :db/valueType :db.type/long
   :db/cardinality :db.cardinality/one}
]');
SELECT public.t('tenant_demo', '[{:sensor/value 42}]');
SELECT public.q(
  'tenant_demo',
  '[:find ?value :where [?e :sensor/value ?value]]',
  '{}'::jsonb
);
```

历史 `pull` 与查询操作遵循 `asOf` 和 `since`，未知选项键会被拒绝。命名存储仍拒绝文档列出的向量、几何及用户事务函数路径，因为这些路径会错误读取默认存储。新存储继承默认存储的授权，仅使用独立模式并不构成租户访问隔离。

事务报告新增 `datoms_asserted` 和 `datoms_retracted`。重复断言未变化的值不计为变更；替换单值属性会同时报告一次撤回和一次断言。读取事务报告的应用应适配修正后的实际 datom 变更语义。

升级将旧命名存储模式保留为 `mentat_<name>_pre118`，并创建新的空替代模式，对跳过的存储及相关订阅/物化视图给出警告。应检查警告并重建受影响的订阅/视图，不能假定旧数据已经迁移到替代模式。新增 `(store_id, tx)` 索引使用普通 CREATE INDEX，可能阻塞写入。大型部署可先按文档中的准确名称并发预建索引，并在升级前检查有效性。
