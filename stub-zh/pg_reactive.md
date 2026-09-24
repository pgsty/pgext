## 用法

来源：

- [PGXN 0.1.8 README](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/README.md)
- [Control file](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/pg_reactive.control)
- [Versioned extension SQL](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/pg_reactive--0.1.8.sql)
- [Module initialization](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/src/pg_reactive.c)

`pg_reactive` 将 PostgreSQL SELECT 查询注册为订阅，通过 LISTEN/NOTIFY 发送结果集变化。0.1.8 支持 PostgreSQL 15–18，需要预加载并重启，订阅管理接口位于 `pgr` 模式中，须由受信任的角色使用。

### 启用与订阅

将库加入现有预加载列表，然后重启 PostgreSQL。安装脚本还使用 `plpgsql`，通常已由数据库默认安装。

```conf
shared_preload_libraries = 'pg_reactive'
pg_reactive.max_subscriptions = 1024
```

以下示例由扩展所有者或受信任的管理员执行，各语句按正常方式提交：

```sql
CREATE EXTENSION pg_reactive;

CREATE TABLE public.pgr_demo (
    id bigint PRIMARY KEY,
    status text NOT NULL,
    total numeric
);

LISTEN pgr;
SELECT pgr.subscribe(
    'open_orders',
    $$SELECT id, total FROM public.pgr_demo WHERE status = 'open'$$
);

INSERT INTO public.pgr_demo VALUES (1, 'open', 42);
SELECT * FROM pgr.subscriptions;
SELECT * FROM pgr.stats();
```

扩展跟踪查询的表与列依赖，安装语句级触发器，并与 UNLOGGED 快照比较集合差异。`delta` 模式发送新增和删除的行；`notify` 模式只发送失效通知，由客户端重新获取结果。

```sql
SELECT pgr.subscribe(
    'orders_hint',
    $$SELECT id FROM public.pgr_demo$$,
    'notify'
);

SELECT pgr.unsubscribe('orders_hint');
SELECT pgr.unsubscribe('open_orders');
```

### 接口与配置

- `pgr.subscribe(query_id, query, mode, audience)` 校验 SELECT 并记录订阅；后两个参数默认使用增量模式和空受众。
- `pgr.unsubscribe(query_id)` 删除订阅、触发器、快照和持久化目录记录。
- `pgr.get_subscriptions()` 与 `pgr.subscriptions` 展示共享内存中的活跃订阅，`pgr.stats()` 提供计数器。
- `pgr.subscription_meta(query_id)` 为代理读取已提交的模式、受众和代次信息。
- `pgr.persisted_subscriptions` 保存持久化订阅目录。

`pg_reactive.max_subscriptions`、`pg_reactive.async_recompute` 和 `pg_reactive.database` 需要重启后生效。异步重算默认关闭，其后台进程连接指定数据库。`pg_reactive.batch_invalidation` 默认开启，在事务提交前合并重算；`pg_reactive.notify_channel` 默认是 `pgr`。后两个参数可由超级用户在运行时修改。

### 恢复与安全边界

重启会清空共享内存。受信任的启动流程须在每个相关数据库执行以下语句，并检查恢复失败的订阅警告：

```sql
SELECT pgr.restore_subscriptions();
```

0.1.8 的 SQL 脚本撤销了订阅函数及视图的公共访问权限。仅授权 `pgr.subscribe()` 不够，因为它还调用受限辅助函数和序列。不要向不受信任的角色开放任意订阅 SQL；上游要求使用固定查询模板、由服务端确定标识符的专用包装接口。受众元数据应由受信任的代理执行权限检查，共用通知频道本身不提供逐订阅授权。

发生通知溢出、快照列结构变化、序号缺失或重新连接时，客户端应重新读取完整结果。通知共用一个频道，不是持久化投递队列。注册代次用于区分重新订阅，客户端或代理须丢弃过时代次。即使变化行很少，重算大型结果集也会增加写入开销。

已核对的 PGXN 发布归档包含控制文件、SQL 和 C 源码。核验时，其声明的 GitHub 仓库返回 404；本版本以以上带版本的 PGXN 源码为依据。
