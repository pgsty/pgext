## 用法

来源：

- [Project README](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/README.md)
- [Extension implementation](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/lib.rs)
- [Extension control file](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/meridian_pg.control)
- [CRDT SQL functions](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/types/mod.rs)
- [Notification trigger](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/trigger.rs)
- [Extension SQL](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/sql/meridian--0.1.0.sql)
- [Package metadata](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/Cargo.toml)
- [Postgres example schema](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/examples/postgres-articles/schema.sql)
- [MIT license](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/LICENSE)

`meridian_pg` 0.1.0 将 Meridian CRDT 状态保存在 PostgreSQL `bytea` 列中，提供计数器、集合、寄存器、大型文档只读视图和变更通知相关的 SQL 函数。该扩展仅源码可用，使用 pgrx 0.13，并声明支持 PostgreSQL 14–17 feature。

### 核心流程

```sql
CREATE EXTENSION meridian_pg;

CREATE TABLE articles (
    id       text PRIMARY KEY,
    views    bytea,
    likes    bytea,
    tags     bytea,
    headline bytea
);

INSERT INTO articles (id) VALUES ('rust-intro');

UPDATE articles
SET views = meridian.gcounter_increment(views, 3, 42),
    likes = meridian.pncounter_decrement(likes, 1, 42)
WHERE id = 'rust-intro';

SELECT meridian.gcounter_value(views),
       meridian.pncounter_value(likes)
FROM articles
WHERE id = 'rust-intro';
```

`orset_add` 接受 JSON 元素字符串以及节点/序列标识；`orset_elements` 返回 JSON 数组。`lww_set` 接受 JSON 值、毫秒墙钟时间和作者 ID，`lww_value` 与 `lww_updated_at_ms` 读取获胜状态。合并函数可合并计数器和集合状态。

### 通知与大型文档

在插入或更新后附加 `meridian.notify_trigger(namespace, id_prefix, primary_key_column, crdt_columns...)`。触发器会为每个发生变化的非空 CRDT 列在 `meridian_ops` 通道发出一个 `pg_notify` 事件，并且不修改行本身。超过扩展 7,800 字节保护的 PostgreSQL 通知负载会被跳过并记录警告，不会截断后继续投递。

可能超过该边界的 RGA 和 Tree 文档应使用 Meridian 服务器的 pg-sync 逻辑复制路径。扩展提供 `rga_text`、`rga_len`、`tree_json`、`tree_node_count` SQL 只读函数；这些 CRDT 类型的写入应来自 Meridian 客户端，而不是 SQL。

### 边界

控制文件设置了 `relocatable = false` 和 `superuser = false`。每列应保持稳定的 CRDT 类型，每次写入也应保持稳定的 client/node ID；解码或序列化非法状态时可能记录警告并返回默认或空状态。扩展本身不认证 Meridian 客户端，也不保证 PostgreSQL 表可被服务器访问。生产使用前，应审查触发器权限、通知暴露、复制槽保留以及 Meridian 服务器的命名空间授权。
