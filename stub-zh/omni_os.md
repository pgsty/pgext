


## 用法

> [omni_os: 操作系统集成](https://docs.omnigres.org/omni_os/intro/)

`omni_os` 扩展提供对操作系统功能的访问。仅供超级用户使用。

### 环境变量

`omni_os.env` 视图暴露系统环境变量：

```sql
SELECT * FROM omni_os.env;
```

| 列 | 类型 |
|:---|:-----|
| `variable` | text |
| `value` | text |

类似于命令行的 `env` 工具，返回 PostgreSQL 进程可见的所有环境变量。
