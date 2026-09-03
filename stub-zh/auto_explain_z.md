## 用法

来源：

- [官方文档](https://github.com/vladich/auto_explain_z/blob/b281da41ca5b2fdb03adaaa82a700f5ac75822cc/README.md)
- [扩展控制文件](https://github.com/vladich/auto_explain_z/blob/b281da41ca5b2fdb03adaaa82a700f5ac75822cc/auto_explain_z.control)
- [官方仓库](https://github.com/vladich/auto_explain_z)

`auto_explain_z` 以压缩二进制格式记录执行计划，并支持轮转与离线解析。

### 启用

将 `auto_explain_z` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `auto_explain_z`：

```ini
shared_preload_libraries = 'auto_explain_z'
```

```sql
CREATE EXTENSION auto_explain_z;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT auto_explain_z_rotate_logfile();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `auto_explain_z_rotate_logfile` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `auto_explain_z` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
