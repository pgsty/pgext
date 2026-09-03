## 用法

来源：

- [官方文档](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/README.md)
- [扩展控制文件](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/pg_chaos_shuffle.control)
- [构建清单](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/Cargo.toml)

`pg_chaos_shuffle` 随机打乱未使用 ORDER BY 的 SELECT 结果，用于混沌测试。

### 启用

将 `pg_chaos_shuffle` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_chaos_shuffle`：

```ini
shared_preload_libraries = 'pg_chaos_shuffle'
```

```sql
CREATE EXTENSION pg_chaos_shuffle;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
LOAD 'pg_chaos_shuffle';

CREATE TEMP TABLE users AS
SELECT g AS id FROM generate_series(1, 100) AS g;

SELECT * FROM users LIMIT 5;
SELECT * FROM users ORDER BY id LIMIT 5;
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_chaos_shuffle` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
