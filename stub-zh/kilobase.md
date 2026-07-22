## 用法

来源：

- [扩展 README](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/README.md)
- [后台工作进程实现](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/src/lib.rs)
- [SQL 对象定义](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/src/sql.rs)
- [扩展控制文件](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/kilobase.control)

`kilobase` 17.4.1 是一个用于调度物化视图刷新的 pgrx 后台工作进程。当前实现是实验性的单数据库调度器：工作进程硬编码连接 `postgres` 数据库，因此扩展对象和任务都必须在该库创建。

### 预加载与数据库设置

把动态库加入 `shared_preload_libraries` 并重启 PostgreSQL；如果只通过 `CREATE EXTENSION` 加载，代码会有意跳过工作进程注册。

```conf
shared_preload_libraries = 'kilobase'
```

连接到 `postgres` 数据库并安装扩展。其 SQL 使用未限定对象名，所以应把扩展保留在默认 `public` 模式，除非已经审计并修补了工作进程中的所有查询。

```sql
CREATE EXTENSION kilobase;
```

安装存在一个特殊副作用：如果数据库内没有物化视图，安装脚本会创建并注册名为 `sample_stats` 的演示视图。不需要它时，请复核并删除该示例及其任务。

### 注册刷新任务

创建物化视图。合适的唯一索引可以让工作进程尝试 `REFRESH MATERIALIZED VIEW CONCURRENTLY`；没有唯一索引时会退回到会阻塞访问的普通刷新。

```sql
CREATE TABLE orders (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id bigint NOT NULL,
    amount numeric NOT NULL
);

CREATE MATERIALIZED VIEW order_totals AS
SELECT customer_id, sum(amount) AS total
FROM orders
GROUP BY customer_id;

CREATE UNIQUE INDEX order_totals_customer_idx
    ON order_totals (customer_id);

SELECT register_matview_refresh('public', 'order_totals', 300, 'orders');

SELECT * FROM matview_refresh_status;
SELECT * FROM matview_refresh_history;
```

可选的第四个参数指定同一模式中的一个源表。提供该参数后，工作进程会比较 `pg_stat_user_tables` 中累计的插入/更新/删除计数，没有变化就跳过刷新。这些统计可能被重置，而且不能描述所有依赖；若正确性要求每次都按计划刷新，请省略该参数。

### 管理对象

- `register_matview_refresh(text, text, integer, text)`：创建或重新启用任务，并返回整数 ID。
- `unregister_matview_refresh(text, text)`：把任务标为停用。
- `matview_refresh_jobs`：调度器状态；每批最多处理十个已到期的活动任务。
- `matview_refresh_log`：刷新结果与耗时。
- `matview_refresh_status` 和 `matview_refresh_history`：运维视图。
- `cleanup_matview_refresh_logs(integer)`：删除早于指定天数的历史；工作进程也会定期执行七天保留期清理。
- `kilobase_health_check()`：汇总工作进程可见性、到期任务、失败、耗时和缺少唯一索引的视图。

### 运维边界

单个工作进程在 `postgres` 数据库中串行处理刷新负载。并发刷新仍可能失败并退回到阻塞刷新；一次慢刷新会延迟后续任务。依赖该计划前，应测试锁影响、工作进程重启行为、任务权限和错误恢复。扩展会创建普通表、视图、函数与触发器，但没有强化它们的权限，因此应只允许可信管理员修改和执行。
