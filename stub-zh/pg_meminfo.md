## 用法

来源：

- [官方文档](https://github.com/bonesmoses/pg_meminfo/blob/da919d5ff0eb8da28de934915a745526739d2cd2/README.md)
- [扩展控制文件](https://github.com/bonesmoses/pg_meminfo/blob/da919d5ff0eb8da28de934915a745526739d2cd2/pg_meminfo.control)
- [官方仓库](https://github.com/bonesmoses/pg_meminfo)

`pg_meminfo` 通过 SQL 检查数据库后端的内存与进程资源使用。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_meminfo`：

```sql
CREATE EXTENSION pg_meminfo;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT pss FROM smap_summary WHERE pid = 4242;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `smap_summary` | VIEW | 扩展创建的检查或查询视图。 |
| `get_all_smaps` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
