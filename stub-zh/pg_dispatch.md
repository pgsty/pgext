
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION "Snehil_Shah@pg_dispatch";
> SELECT pgdispatch.fire('SELECT pg_sleep(40);');
> SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
> ```
>
> 来源：[README](https://github.com/Snehil-Shah/pg_dispatch)，[database.dev 页面](https://database.dev/Snehil_Shah/pg_dispatch)

`pg_dispatch` 是 PostgreSQL 的异步 SQL 调度器。它被设计为兼容 TLE 的 `pg_later` 替代方案，构建在 `pg_cron` 之上，因此可用于 Supabase 和 AWS RDS 等受限环境。

## 先决条件

上游 README 列出了以下要求：

- PostgreSQL 13 或更高版本
- `pg_cron` 1.5.0 或更高版本
- `pgcrypto`

## 安装

文档中给出的 TLE 安装路径如下：

```sql
SELECT dbdev.install(Snehil_Shah@pg_dispatch);
CREATE EXTENSION "Snehil_Shah@pg_dispatch";
```

README 提醒，该扩展会安装到 `pgdispatch` 模式中，如果系统里已经存在同名模式，可能会发生命名空间冲突。

## 函数

### `pgdispatch.fire(command text)`

异步派发一条 SQL 命令：

```sql
SELECT pgdispatch.fire('SELECT pg_sleep(40);');
```

### `pgdispatch.snooze(command text, delay interval)`

延迟派发一条 SQL 命令：

```sql
SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
```

README 指出，这个延迟调度是异步完成的，不会阻塞主事务。

## 适用场景

该项目主要用于数据库内的异步副作用，尤其适合 PL/pgSQL 或触发器流程。示例场景是把昂贵的通知或分析任务从 `AFTER INSERT` 触发器中拆出去，让主 RPC 或应用请求更快返回。
