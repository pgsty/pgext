## 用法

来源：

- [官方文档](https://github.com/akonorev/pg_stat_errors/blob/af98390e57f5cc529765c5efcf20a4bd6ddb04af/README.rst)
- [扩展控制文件](https://github.com/akonorev/pg_stat_errors/blob/af98390e57f5cc529765c5efcf20a4bd6ddb04af/pg_stat_errors.control)
- [官方仓库](https://github.com/akonorev/pg_stat_errors)

`pg_stat_errors` 提供集群级错误类别计数与有界近期错误样本。

### 启用

将 `pg_stat_errors` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_stat_errors`：

```ini
shared_preload_libraries = 'pg_stat_errors'
```

```sql
CREATE EXTENSION pg_stat_errors;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
displays the last errors that occured in the database. This view contains up to
``pg_stat_errors.max_last`` rows.

+---------------+----------------+-------------------------------------------------------+
| Name          | Type           | Description                                           |
+===============+================+=======================================================+
| error_time    | timestamp with | Time of occurrence of the error                       |
|               | time zone      |                                                       |
+---------------+----------------+-------------------------------------------------------+
| userid        | oid            | User OID                                              |
+---------------+----------------+-------------------------------------------------------+
| dbid          | oid            | Database OID                                          |
+---------------+----------------+-------------------------------------------------------+
| query         | text           | Text of the query                                     |
+---------------+----------------+-------------------------------------------------------+
| error_level   | text           | Error level (WARNING, ERROR, FATAL and PANIC)         |
+---------------+----------------+-------------------------------------------------------+
| error_state   | text           | Error state as a five-character code                  |
+---------------+----------------+-------------------------------------------------------+
| error_message | text           | Error message                                         |
+---------------+----------------+-------------------------------------------------------+


dba_stat_errors_last view
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_stat_errors_total_errors` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stat_errors_info` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stat_errors_last` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `dba_stat_errors` | VIEW | 扩展创建的检查或查询视图。 |
| `dba_stat_errors_last` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_errors_reset` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 预加载 `pg_stat_errors` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
