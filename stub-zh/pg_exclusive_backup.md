## 用法

来源：

- [官方文档](https://github.com/MasaoFujii/pg_exclusive_backup/blob/d6c38fc0e00a7fdbb9f5a9604c66d10a857c3d19/README.md)
- [扩展控制文件](https://github.com/MasaoFujii/pg_exclusive_backup/blob/d6c38fc0e00a7fdbb9f5a9604c66d10a857c3d19/pg_exclusive_backup.control)
- [官方仓库](https://github.com/MasaoFujii/pg_exclusive_backup)

`pg_exclusive_backup` 为 PostgreSQL 15 及以上版本提供排他式物理备份兼容函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_exclusive_backup`：

```sql
CREATE EXTENSION pg_exclusive_backup;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT pg_start_backup('nightly', true);
SELECT pg_is_in_backup();

-- Copy the physical backup while backup mode is active.
SELECT pg_stop_backup();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_catalog.pg_start_backup` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_catalog.pg_stop_backup` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_catalog.pg_backup_start_time` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_catalog.pg_is_in_backup` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 15, 16, 17, 18；不要推断未列出的主版本。
- 扩展会在 `pg_catalog` 下固定或创建模式对象；权限与备份审查应包含这些对象。
- 目录生命周期为 deprecated；生产使用前应测试升级、备份恢复与服务器兼容性。
