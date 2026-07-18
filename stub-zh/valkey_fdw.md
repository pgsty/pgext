## 用法

来源：

- [上游 README](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/README.md)
- [扩展 control 文件](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/valkey_fdw.control)
- [PGXN 包元数据](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/META.json)

`valkey_fdw` 是用于读取和修改 Valkey 键值服务器数据的外部数据包装器。`1.0` 版本要求 PostgreSQL 17 或更高版本，并依赖 `libvalkey` C 客户端库；无需配置预加载。

```sql
CREATE EXTENSION valkey_fdw;

CREATE SERVER valkey_server
FOREIGN DATA WRAPPER valkey_fdw
OPTIONS (address '127.0.0.1', port '6379');

CREATE FOREIGN TABLE valkey_kv (
    key text,
    val text
)
SERVER valkey_server
OPTIONS (database '0');

SELECT * FROM valkey_kv WHERE key = 'session:42';
```

服务器要求认证时，应通过 user mapping 的 `password` 选项提供凭据。外部表还可通过文档中的 `tabletype`、`tablekeyprefix`、`tablekeyset` 和 `singleton_key` 选项表示 Valkey hash、list、set 与 sorted set。

Valkey 不提供 PostgreSQL 风格的 MVCC 或 SQL cursor，因此扫描不是原子快照：构造 tuple 时 key 可能消失，cursor 扫描也可能重复返回同一元素。只能下推一个针对 `key` 列的等值条件。由于 FDW API 会变化，上游还要求源码分支与 PostgreSQL 大版本匹配。
