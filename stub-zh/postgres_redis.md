## 用法

来源：

- [官方 README](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/README.md)
- [官方扩展源码](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/src/lib.rs)
- [官方 GUC 定义](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/src/gucs.rs)

`postgres_redis` 是实验性的钩子与后台工作进程扩展，把配置表中的指定值复制到 Redis。它观察符合条件的 `SELECT` 与 `UPDATE` 语句，在提交后把键值对放入队列，再由工作进程异步发送 Redis `SET` 命令。

### 配置与流程

在服务器启动前配置一张表及一组键/值列，预加载动态库，然后重启 PostgreSQL。

```conf
shared_preload_libraries = 'postgres_redis'

postgres_redis.redis_url = 'redis://127.0.0.1'
postgres_redis.table = 'users'
postgres_redis.key_column = 'first_name'
postgres_redis.value_column = 'last_name'
postgres_redis.bg_delay = 10
```

```sql
CREATE EXTENSION postgres_redis;

SELECT last_name
FROM users
WHERE first_name = 'Adebayo';

UPDATE users
SET last_name = 'Updated'
WHERE first_name = 'Adebayo';
```

### GUC 与行为

- `postgres_redis.redis_url` 是 Redis 连接 URL。
- `postgres_redis.table`、`postgres_redis.key_column` 与 `postgres_redis.value_column` 选择唯一一组表映射。
- `postgres_redis.bg_delay` 是以秒为单位的工作进程间隔，最小值为一。
- 只有当键列出现在与常量比较的等值谓词中时，查询钩子才识别配置表。其他谓词和命令会被忽略。

### 实验性限制

本次复核的源码使用容量为 400 的共享队列，以及固定 127 字符的键和值数组；队列已满时会发出警告并丢弃更新。Redis 投递发生在 PostgreSQL 提交后，但不与 Redis 构成事务，因此崩溃和 Redis 故障可能导致变更丢失或延迟。工作进程只打开一个 Redis 连接，设计也只覆盖一组表映射。

`CREATE EXTENSION` 还嵌入了仓库的 `sql/test.sql`，会创建并填充 `test` 与 `users` 表。在已经使用这些名称的数据库安装前，必须检查或修补打包 SQL。control 文件要求超级用户安装；唯一显式 SQL 函数是 `hello_postgres_redis`。这个项目应被视为学习原型，而不是通用变更数据捕获系统。
