## 用法

来源：

- [官方文档](https://github.com/asotolongo/pg_login_stat/blob/a0fbc06b6e3d7358561647a3893a702faa2a4774/README.md)
- [扩展控制文件](https://github.com/asotolongo/pg_login_stat/blob/a0fbc06b6e3d7358561647a3893a702faa2a4774/pg_login_stat.control)
- [官方仓库](https://github.com/asotolongo/pg_login_stat)

`pg_login_stat` 按用户和数据库统计认证成功与失败次数。

### 启用

将 `pg_login_stat` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_login_stat`：

```ini
shared_preload_libraries = 'pg_login_stat'
```

```sql
CREATE EXTENSION pg_login_stat;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM pg_login_stats;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_login_stats` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_login_stat_reset` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 15, 16, 17；不要推断未列出的主版本。
- 预加载 `pg_login_stat` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
