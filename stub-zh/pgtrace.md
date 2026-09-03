## 用法

来源：

- [官方文档](https://github.com/Srujan-rai/Pgtrace/blob/0e97f8c4528008ef39d72181b5cf48db05ab74df/README.md)
- [扩展控制文件](https://github.com/Srujan-rai/Pgtrace/blob/0e97f8c4528008ef39d72181b5cf48db05ab74df/pgtrace.control)
- [官方仓库](https://github.com/Srujan-rai/Pgtrace)

`pgtrace` 通过 PostgreSQL 执行器钩子提供查询、锁、失败与延迟可观测性。

### 启用

将 `pgtrace` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pgtrace`：

```ini
shared_preload_libraries = 'pgtrace'
```

```sql
CREATE EXTENSION pgtrace;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Get count of currently tracked queries
SELECT pgtrace_query_count();

-- Clear all query stats
SELECT pgtrace_reset();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pgtrace_query_stats` | VIEW | 扩展创建的检查或查询视图。 |
| `pgtrace_alien_queries` | VIEW | 扩展创建的检查或查询视图。 |
| `pgtrace_failing_queries` | VIEW | 扩展创建的检查或查询视图。 |
| `pgtrace_latency_histogram` | VIEW | 扩展创建的检查或查询视图。 |
| `pgtrace_metrics` | VIEW | 扩展创建的检查或查询视图。 |
| `pgtrace_query_count` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgtrace_reset` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgtrace_slow_queries` | VIEW | 扩展创建的检查或查询视图。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 15, 16；不要推断未列出的主版本。
- 预加载 `pgtrace` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
