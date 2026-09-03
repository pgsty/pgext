## 用法

来源：

- [官方文档](https://github.com/maheshk32/pg_table_guard/blob/65a85d8f451986ee662bb379ee631d159c1d51fa/README.md)
- [扩展控制文件](https://github.com/maheshk32/pg_table_guard/blob/65a85d8f451986ee662bb379ee631d159c1d51fa/table_guard.control)
- [官方仓库](https://github.com/maheshk32/pg_table_guard)

`table_guard` 可配置地阻止 DROP、TRUNCATE 与无条件整表 DELETE。

### 启用

将 `table_guard` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `table_guard`：

```ini
shared_preload_libraries = 'table_guard'
```

```sql
CREATE EXTENSION table_guard;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM table_guard_settings;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `table_guard_help` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `table_guard_settings` | VIEW | 扩展创建的检查或查询视图。 |
| `table_guard_version` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `table_guard` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
