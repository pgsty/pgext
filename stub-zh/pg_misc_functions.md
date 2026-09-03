## 用法

来源：

- [官方文档](https://github.com/BRupireddy2/pg_misc_functions/blob/9e4718148093375fd5b563a5c4fea987558a0502/README.md)
- [扩展控制文件](https://github.com/BRupireddy2/pg_misc_functions/blob/9e4718148093375fd5b563a5c4fea987558a0502/pg_misc_functions.control)
- [官方仓库](https://github.com/BRupireddy2/pg_misc_functions)

`pg_misc_functions` 用于制造后端错误、PANIC、崩溃与故障切换演练的管理函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_misc_functions`：

```sql
CREATE EXTENSION pg_misc_functions;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT pg_current_wal_tli();
SELECT pg_control_checkpoint_tli();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_signal_backend` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_cause_fatal` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_cause_panic` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_control_checkpoint_previous_tli` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_control_checkpoint_tli` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_current_wal_tli` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_last_wal_receive_tli` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_last_wal_replay_tli` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
