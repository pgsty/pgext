## 用法

来源：

- [官方文档](https://github.com/mhagander/pg_sslstatus/blob/906882efacf0c40907910dc09274a3447c6a58b8/README.md)
- [扩展控制文件](https://github.com/mhagander/pg_sslstatus/blob/906882efacf0c40907910dc09274a3447c6a58b8/pg_sslstatus.control)
- [官方仓库](https://github.com/mhagander/pg_sslstatus)

`pg_sslstatus` 面向 PostgreSQL 9.5 之前版本的历史 SSL 连接状态视图。

### 启用

将 `pg_sslstatus` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_sslstatus`：

```ini
shared_preload_libraries = 'pg_sslstatus'
```

```sql
CREATE EXTENSION pg_sslstatus;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM pg_sslstatus;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_get_sslstatus` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 预加载 `pg_sslstatus` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
- 目录生命周期为 deprecated；生产使用前应测试升级、备份恢复与服务器兼容性。
