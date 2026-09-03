## 用法

来源：

- [官方文档](https://github.com/VladUl287/table_change_tracker/blob/407d7b365b1ec2af9a32ba0203d56b7e838e2386/README.md)
- [扩展控制文件](https://github.com/VladUl287/table_change_tracker/blob/407d7b365b1ec2af9a32ba0203d56b7e838e2386/table_change_tracker.control)
- [官方仓库](https://github.com/VladUl287/table_change_tracker)

`table_change_tracker` 在共享内存中跟踪指定表的最近修改时间。

### 启用

将 `table_change_tracker` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `table_change_tracker`：

```ini
shared_preload_libraries = 'table_change_tracker'
```

```sql
CREATE EXTENSION table_change_tracker;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Get timestamps for multiple tables
SELECT get_last_timestamps(ARRAY['public.users', 'public.orders']::regclass[]);
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `enable_table_tracking` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `disable_table_tracking` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `get_last_timestamp` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `get_last_timestamps` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `is_table_tracked` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `set_last_timestamp` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 预加载 `table_change_tracker` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
