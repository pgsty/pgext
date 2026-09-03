## 用法

来源：

- [官方文档](https://github.com/hank-cp/pg_op_log/blob/731aeea2b8053a3dde031276729401375a598c4c/README.md)
- [扩展控制文件](https://github.com/hank-cp/pg_op_log/blob/731aeea2b8053a3dde031276729401375a598c4c/op_log.control)
- [官方仓库](https://github.com/hank-cp/pg_op_log)

`op_log` 在事务提交后记录行级变更历史的触发器扩展。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `op_log`：

```sql
CREATE EXTENSION op_log CASCADE;
```

经审查的控制文件或官方流程要求 `plv8`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT op_log_disable('your_table_name'::regclass);
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `op_log_enable` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `data_op_log` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `op_log_disable` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `op_log_diff_array` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `op_log_diff_normal` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `op_log_diff_object` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `op_log_diff_object_array` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `op_log_diff_other` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
