## 用法

来源：

- [官方文档](https://github.com/psqlmaster/pgsyswatch/blob/0dde3e65122066df75f47005aeb6b0fc951b1223/readme.md)
- [扩展控制文件](https://github.com/psqlmaster/pgsyswatch/blob/0dde3e65122066df75f47005aeb6b0fc951b1223/pgsyswatch.control)
- [官方仓库](https://github.com/psqlmaster/pgsyswatch)

`pgsyswatch` 通过 SQL 暴露主机进程、CPU、内存、交换区、网络与负载统计。

### 启用

将 `pgsyswatch` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pgsyswatch`：

```ini
shared_preload_libraries = 'pgsyswatch'
```

```sql
CREATE EXTENSION pgsyswatch;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
select * from  pgsyswatch.net_and_loadavg_snapshots;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `proc_activity_snapshots` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pgsyswatch.net_and_loadavg` | VIEW | 扩展创建的检查或查询视图。 |
| `proc_monitor_all` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `net_and_loadavg_snapshots` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pg_proc_activity` | VIEW | 扩展创建的检查或查询视图。 |
| `cpu_frequencies` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `manage_partitions_maintenance` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 预加载 `pgsyswatch` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
