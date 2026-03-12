

## 用法

> [qos: PostgreSQL 会话和查询的 QoS 资源治理扩展](https://github.com/appstonia/pg_qos)

`qos` 扩展为 PostgreSQL 提供服务质量（QoS）资源治理，允许管理员对每个角色和每个数据库设置内存使用、CPU 核心数和并发事务/语句的限制。

### 配置参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `qos.work_mem_limit` | 字节 | 每个会话的最大有效 `work_mem` |
| `qos.cpu_core_limit` | 整数 | 会话可用的最大 CPU 核心数 |
| `qos.max_concurrent_tx` | 整数 | 最大并发事务数 |
| `qos.max_concurrent_select` | 整数 | 最大并发 SELECT 语句数 |
| `qos.max_concurrent_update` | 整数 | 最大并发 UPDATE 语句数 |
| `qos.max_concurrent_delete` | 整数 | 最大并发 DELETE 语句数 |
| `qos.max_concurrent_insert` | 整数 | 最大并发 INSERT 语句数 |

### 角色级限制

```sql
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.cpu_core_limit = '2';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
```

### 数据库级限制

```sql
ALTER DATABASE appdb SET qos.max_concurrent_tx = '200';
```

### 角色+数据库组合限制

```sql
ALTER ROLE app_user IN DATABASE appdb SET qos.work_mem_limit = '4MB';
ALTER ROLE app_user IN DATABASE appdb SET qos.max_concurrent_update = '10';
```

### 执行行为

- **工作内存**：拦截 `SET work_mem` 并拒绝超过配置限制的值
- **CPU 限制**（仅 Linux）：通过 CPU 亲和性将后端绑定到 N 个 CPU 核心；在非 Linux 平台上改为限制并行工作进程
- **并发**：执行器钩子按类型跟踪活动事务/语句；违规将阻止执行

### 可观测性

```sql
SET client_min_messages = 'debug1';  -- 启用 QoS 事件的调试输出
```

角色级和数据库级设置中最严格的组合生效。需要 `shared_preload_libraries = 'qos'` 和 PostgreSQL 15+。
