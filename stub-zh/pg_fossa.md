## 用法

来源：

- [官方文档](https://github.com/fossas/pg_fossa/blob/1ea06cfa082d51ab2a5cf722f6eb57fd7688928f/README.md)
- [扩展控制文件](https://github.com/fossas/pg_fossa/blob/1ea06cfa082d51ab2a5cf722f6eb57fd7688928f/pg_fossa.control)
- [官方仓库](https://github.com/fossas/pg_fossa)

`pg_fossa` 为 FOSSA hasGraph 数据结构与查询提供支持的已归档 SQL 扩展。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_fossa`：

```sql
CREATE EXTENSION pg_fossa;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT fossa_version();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `fossa_dependencies` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fossa_version` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `array_intersect` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `array_intersect_agg` | AGGREGATE | 扩展提供的聚合函数。 |
| `array_sort_unique` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `array_symmetric_difference` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `array_union` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `array_union_agg` | AGGREGATE | 扩展提供的聚合函数。 |

### 运维与边界

- 目录生命周期为 archived；生产使用前应测试升级、备份恢复与服务器兼容性。
