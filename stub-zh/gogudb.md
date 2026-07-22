## 用法

来源：

- [官方 README](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/README.md)
- [官方安装指南](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/Install.md)
- [官方 control 文件](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/gogudb.control)

`gogudb` 是一个 Postgres-XL 时代的分布式表分区扩展。它通过自带的外部数据封装器把哈希或范围分区数据路由到 PostgreSQL 服务器，并在扩展元数据中保存分区规则。它只适合维护兼容的遗留部署，不能替代现代声明式分区。

### 核心流程

必须预加载该库。在服务器配置中加入下列设置，重启 PostgreSQL 后再创建扩展：

```conf
shared_preload_libraries = 'gogudb'
```

```sql
CREATE EXTENSION gogudb;

CREATE SERVER shard1
FOREIGN DATA WRAPPER gogudb_fdw
OPTIONS (host '10.0.0.11', port '5432', dbname 'app');

CREATE USER MAPPING FOR CURRENT_USER
SERVER shard1 OPTIONS (user 'app', password 'secret');
```

随后在 `_gogu.server_map` 中定义服务器/范围映射，用 `_gogu.reload_range_server_set()` 重新加载，再在 `_gogu.table_partition_rule` 中登记父表的哈希或范围规则。这些元数据布局属于扩展实现契约，必须严格遵循已安装版本的上游示例。

### 主要对象

- `gogudb_fdw`：用于远程分片的外部数据封装器。
- `_gogu.server_map`：分区范围到外部服务器的映射。
- `_gogu.table_partition_rule`：表级分区规则；上游使用分区类型 `1` 表示哈希、`2` 表示范围。
- `_gogu.reload_range_server_set()`：元数据改变后重载内存中的范围/服务器映射。
- `_gogu` 与 `gogudb_partition_table`：固定扩展模式；扩展不可重定位。

### 注意事项

官方材料面向旧版 PostgreSQL 9.6/10 与 CentOS 7 时代环境。不要推断其支持当前 PostgreSQL，必须确认二进制兼容性。扩展通过 `shared_preload_libraries` 安装钩子，因此实例中所有数据库都共同承担所加载库的运行风险。

上游明确不支持重命名分区父表、子表、远程表或其模式；直接修改或删除子表/远程表；分区拆分或合并；以及读写分离。应把元数据表视为高权限控制数据，保护用户映射中的凭据，并在一次性分片环境中演练故障与恢复。
