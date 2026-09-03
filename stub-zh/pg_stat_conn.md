## 用法

来源：

- [官方文档](https://github.com/guedes/pg_stat_conn/blob/8e5576e55e43238e94f42712ea3c0b138b64632b/README.md)
- [扩展控制文件](https://github.com/guedes/pg_stat_conn/blob/8e5576e55e43238e94f42712ea3c0b138b64632b/pg_stat_conn.control)
- [官方仓库](https://github.com/guedes/pg_stat_conn)

`pg_stat_conn` 按数据库与角色累计连接、断开与认证失败计数。

### 启用

将 `pg_stat_conn` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_stat_conn`：

```ini
shared_preload_libraries = 'pg_stat_conn'
```

```sql
CREATE EXTENSION pg_stat_conn;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
psql -c "SELECT pg_stat_conn_reset('postgres', 'postgres');"
psql -c "SELECT pg_stat_conn_reset();"
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_stat_conn_reset` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_stat_conn` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
