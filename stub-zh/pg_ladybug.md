## 用法

来源：

- [官方文档](https://github.com/LadybugDB/pg_ladybug/blob/895cbdb5d1d0fb7db65e78ac7352795489b91f5d/README.md)
- [扩展控制文件](https://github.com/LadybugDB/pg_ladybug/blob/895cbdb5d1d0fb7db65e78ac7352795489b91f5d/pg_ladybug.control)
- [官方仓库](https://github.com/LadybugDB/pg_ladybug)

`pg_ladybug` 通过嵌入式 Ladybug 图引擎对 PostgreSQL 表执行 Cypher 查询。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_ladybug`：

```sql
CREATE EXTENSION pg_ladybug;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
┌─────────────────────────────────────────────┐
│          PostgreSQL backend                  │
│                                              │
│  Cypher query → ladybug.cypher()             │
│       │                                      │
│       ▼                                      │
│  Ladybug planner (liblbug, compile-time      │
│  linked via -llbug)                          │
│    - Parse Cypher                            │
│    - Discover tables via pg_client           │
│      (libpq → information_schema)            │
│    - Produce EXPLAIN plan                    │
│       │                                      │
│       ▼                                      │
│  extract_pushed_sql()                        │
│    - Parse plan text                         │
│    - Extract "Function:" section (SQL)       │
│    - OR construct SELECT from plan nodes     │
│       │                                      │
│       ▼                                      │
│  SPI_execute(pushed_sql)                     │
│    - Native Postgres execution               │
│    - Return rows via SRF                     │
│                                              │
└─────────────────────────────────────────────┘
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ladybug.cypher` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.explain` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.pushed_sql` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.register_node` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.register_edge` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.list_labels` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.reset_graph` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ladybug.sql_query` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 17, 18；不要推断未列出的主版本。
- 扩展会在 `ladybug` 下固定或创建模式对象；权限与备份审查应包含这些对象。
