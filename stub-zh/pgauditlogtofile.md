
## 用法

> [pgauditlogtofile: 将 pgAudit 日志重定向到独立文件](https://github.com/fmbiete/pgauditlogtofile)

`pgauditlogtofile` 是 pgAudit 的附加组件，会将审计日志行重定向到独立文件，而不是 PostgreSQL 服务器日志，并支持自动轮转。

```sql
CREATE EXTENSION pgauditlogtofile;
```

### 配置参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pgaudit.log_format` | `csv` | 输出格式：`csv` 或 `json` |
| `pgaudit.log_directory` | `log` | 审计文件目录，留空则禁用 |
| `pgaudit.log_filename` | `audit-%Y%m%d_%H%M.log` | 文件名模式，支持时间格式 |
| `pgaudit.log_file_mode` | `0600` | 审计日志文件权限 |
| `pgaudit.log_rotation_age` | `1440` | 轮转间隔，单位分钟，1 天 |
| `pgaudit.log_connections` | `off` | 记录连接事件，需要 `log_connections = on` |
| `pgaudit.log_disconnections` | `off` | 记录断开事件，需要 `log_disconnections = on` |
| `pgaudit.log_autoclose_minutes` | `0` | 空闲 N 分钟后自动关闭文件句柄 |
| `pgaudit.log_execution_time` | `off` | 统计语句执行时间 |
| `pgaudit.log_execution_memory` | `off` | 统计语句内存占用 |

### 设置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'pgaudit, pgauditlogtofile'
pgaudit.log_directory = 'log'
pgaudit.log_filename = 'audit-%Y%m%d_%H%M.log'
pgaudit.log_rotation_age = 1440
```

审计条目会写入独立文件，服务器日志保持整洁。
