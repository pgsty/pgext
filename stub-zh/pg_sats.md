## 用法

来源：

- [官方文档](https://github.com/jasonmassie01/pg-sats/blob/3c1a55bf8d550f6e34fc000dcd3f6bd6480d6651/README.md)
- [扩展控制文件](https://github.com/jasonmassie01/pg-sats/blob/3c1a55bf8d550f6e34fc000dcd3f6bd6480d6651/pg_sats.control)
- [官方仓库](https://github.com/jasonmassie01/pg-sats)

`pg_sats` 使用聪计价余额按角色计量查询，并可选择强制限额。

### 启用

将 `pg_sats` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_sats`：

```ini
shared_preload_libraries = 'pg_sats'
```

```sql
CREATE EXTENSION pg_sats;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Fund an account with 1000 sats
SELECT pg_sats.deposit('alice', 1000);

-- Check balance
SELECT pg_sats.balance('alice');
-- Returns: 1000

-- Alice runs queries... each costs 1 sat (default)
-- After 10 queries:
SELECT pg_sats.balance('alice');
-- Returns: 990

-- View spending history
SELECT * FROM pg_sats.history('alice');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_sats.balance` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_sats.deposit` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_sats.history` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_sats.balances` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pg_sats.query_history` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_sats` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
- 扩展会在 `pg_sats` 下固定或创建模式对象；权限与备份审查应包含这些对象。
