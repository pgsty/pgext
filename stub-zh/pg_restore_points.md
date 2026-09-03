## 用法

来源：

- [官方文档](https://github.com/dalibo/pg_restore_points/blob/d1c8b71c10cd223b601d231bf2b726605cce5371/README.md)
- [扩展控制文件](https://github.com/dalibo/pg_restore_points/blob/d1c8b71c10cd223b601d231bf2b726605cce5371/pg_restore_points.control)
- [官方仓库](https://github.com/dalibo/pg_restore_points)

`pg_restore_points` 创建、登记、列出与删除具名 PostgreSQL 恢复点。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_restore_points`：

```sql
CREATE EXTENSION pg_restore_points;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT pg_purge_restore_points('interval_value');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `restore_points` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pg_extend_create_restore_point` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_purge_restore_points` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `restore_point_mode` | TYPE | 扩展创建的用户数据类型。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
