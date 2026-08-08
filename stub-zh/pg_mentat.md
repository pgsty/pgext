## 用法

来源：

- [pg_mentat v1.5.7 README](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/README.md)
- [pg_mentat v1.5.7 控制文件](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/pg_mentat.control)
- [pg_mentat v1.5.6 到 v1.5.7 升级 SQL](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/sql/pg_mentat--1.5.6--1.5.7.sql)
- [Pigsty 软件包矩阵](https://pgext.cloud/ext/pg_mentat)

`pg_mentat` 在 PostgreSQL 内实现与 Datomic 兼容的数据模型和 Datalog 查询引擎。它将不可变事实存储为有类型的 datom，并通过 SQL 函数提供模式事务、Datalog 查询、pull 表达式、时间旅行、事务历史和永久切除功能。它适用于需要这种模型的应用；并非关系表或 SQL 的透明替代品。

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
SELECT mentat.pull('[*]', 10001);
SELECT mentat.log('default', 1000001, 1000010);
SELECT mentat.diff('default', 1000003, 1000007);

SELECT mentat.mentat_with('[
  {:person/name "Alice" :person/age 31}
]');
```

`mentat.pull` 返回实体形态的 JSON。`mentat.log` 和 `mentat.diff` 提供事务历史，`mentat.mentat_with` 则评估事务但不持久化。查询还可以使用文档所述的数据库参数，以某个事务时点或从某个事务之后开始求值。

永久切除有意与通常的不可变历史机制分开：

```sql
SELECT mentat.mentat_excise('default', 10042, NULL);
```

执行切除前请检查目标实体和备份；该操作会永久移除 datom，适用于隐私擦除等要求。

### 重要对象

- `mentat.t(edn)`：写入模式或数据事务。
- `mentat.q(query, inputs)`：执行 Datalog。
- `mentat.pull(pattern, eid)` 和 `mentat.pull_many(pattern, eids)`：以实体形态读取数据。
- `mentat.entity(eid)` 和 `mentat.schema()`：检查实体或当前模式。
- `mentat.log(...)` 和 `mentat.diff(...)`：检查事务历史。
- `mentat.stats()`、`mentat.storage()` 和 `mentat.cache_stats()`：运行状态检查。
- `mentat.subscribe(...)`：通过 PostgreSQL `LISTEN`/`NOTIFY` 提供响应式查询通知。

该扩展在 `mentat` 模式下的窄表中存储有类型的 datom，包括引用、整数、字符串、布尔、浮点、时刻、关键字、UUID 和字节值。

### 要求与注意事项

- 上游 v1.5.7 支持 PostgreSQL 13-18。当前 Pigsty 软件包面向 PostgreSQL 14-18，并使用 pgrx 0.19.1 重新构建；上游标签源码声明使用 pgrx 0.17。请将打包后的二进制作为兼容性边界。
- 该扩展不可重定位，也不要求 `shared_preload_libraries`。
- 可选的 `mentatd` HTTP/Datomic 线协议守护进程是上游配套程序，不包含在 Pigsty `pg_mentat` 软件包中。仅通过 SQL 使用扩展并不需要它。
- Datalog 编译、递归 pull、全文属性、订阅和历史记录可能呈现截然不同的成本特征。请使用文档所述的 explain 辅助函数检查生成的 SQL，并在代表性数据上进行基准测试。
- 切除操作绕过通常的不可变历史模型。请限制权限并审计其使用。
