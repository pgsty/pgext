## 用法

来源：[README](https://github.com/Snehil-Shah/pg_dispatch/blob/master/README.md)，[database.dev page](https://database.dev/Snehil_Shah/pg_dispatch)

`pg_dispatch` 是一个基于 `pg_cron` 的 TLE-compatible 异步 SQL dispatcher。它面向 RDS 或 Supabase 这类沙箱环境，用于把副作用操作延迟到调用者主事务之外执行。

### 依赖与安装

- PostgreSQL 13+
- `pg_cron` 1.5.0+
- `pgcrypto`

```sql
SELECT dbdev.install(Snehil_Shah@pg_dispatch);
CREATE EXTENSION "Snehil_Shah@pg_dispatch";
```

该扩展会安装到 `pgdispatch` schema。

### 主要函数

```sql
SELECT pgdispatch.fire('SELECT pg_sleep(40);');
SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
```

- `pgdispatch.fire(command text)`：将 SQL 放入队列并立即异步执行。
- `pgdispatch.snooze(command text, delay interval)`：将 SQL 放入队列并延迟异步执行。

### 典型用法

官方 README 将 `pg_dispatch` 定位于 PL/pgSQL 或 trigger 驱动的工作流：前台 RPC 需要快速提交，而通知、分析更新或其他昂贵 SQL 则在稍后的独立事务中执行。

### 注意事项

- 除了 `pg_cron` 之外，运行时还依赖 `pgcrypto`。
- `delay` 参数会被截断到秒级精度。
- 项目文档优先介绍 TLE/database.dev 安装方式；README 中也链接了手工 PGXN 安装方式。
