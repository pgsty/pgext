## 用法

来源：

- [固定提交的上游 README](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/README.md)
- [固定提交的使用指南](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/docs/usage.md)
- [固定提交的安装指南](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/docs/setup.md)
- [固定提交的 0.1 版安装 SQL](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/sql/event_manager--0.1.sql)

`event_manager` `0.1` 通过事件队列与后续工作队列实现松耦合的 DML 触发处理。work-item 查询把 `INSERT`、`UPDATE` 或 `DELETE` 事件转换为 JSON 参数；随后 action 可以同步执行本地 SQL，也可以交给独立的 `event_manager` 守护进程异步处理。

### 核心流程

扩展由超级用户安装在固定的 `event_manager` schema 中。以下面向 psql 的骨架注册一个同步本地 action；最后一次插入使用 `\gset` 返回的标识：

```sql
CREATE EXTENSION event_manager;

CREATE TABLE public.orders (id integer PRIMARY KEY, note text);
CREATE TABLE public.order_events (order_id integer, event_op char(1));

INSERT INTO event_manager.tb_event_table(schema_name, table_name)
VALUES ('public', 'orders')
RETURNING event_table \gset

INSERT INTO event_manager.tb_action(query, static_parameters)
VALUES (
  'INSERT INTO public.order_events(order_id, event_op) VALUES (?order_id?::integer, ?event_op?::char(1))',
  '{}'::jsonb
)
RETURNING action \gset

INSERT INTO event_manager.tb_event_table_work_item(
  source_event_table, action, work_item_query, op, execute_asynchronously
)
VALUES (
  :'event_table',
  :'action',
  'SELECT jsonb_build_object(''order_id'', ?pk_value?, ''event_op'', ?op?) AS parameters',
  ARRAY['I']::char(1)[],
  false
);

INSERT INTO public.orders VALUES (1, 'first order');
SELECT * FROM public.order_events;
```

异步处理需要启动上游服务，并至少配置一个 event worker 与一个 work worker（`-E 1 -W 1`）。远程 HTTP action 不能同步执行。

### 重要对象

- `event_manager.tb_event_table`：注册由生成的行级触发器监视的源表。
- `event_manager.tb_event_table_work_item`：把操作类型、可选 when-function、参数生成查询、action 和同步/异步模式关联起来。
- `event_manager.tb_action`：保存本地 DML 或 `GET`、`POST`、`PUT` HTTP action 及静态参数。
- `event_manager.tb_event_queue` 与 `event_manager.tb_work_queue`：保存事件快照和可执行 action 参数。
- `event_manager.tb_setting`：持久化扩展 GUC，并将其应用于当前数据库。
- `?pk_value?`、`?op?`、`?OLD.column?`、`?NEW.column?` 和具名 session-GUC 绑定标记：把事件状态注入 work-item 查询；查询必须返回别名为 `parameters` 的 `jsonb`。

### 运维说明

`0.1` 版 schema 把源主键值保存为 `integer`，因此它不是适用于 UUID、组合键或 bigint 键的通用队列。同步 action 在触发事务中执行；异步 action 稍后运行，因此回滚和可见性语义不同。失败的队列项会被标为失败，但不会自动重试。如果安装时选择 unlogged queue，它们会在崩溃恢复时被截断，也不会复制。

写入 `tb_action`、`tb_event_table_work_item` 和 `tb_setting` 可以触发任意 SQL、出站 HTTP、触发器创建或数据库级配置变更。应只允许可信管理员操作这些表，并验证绑定标记替换，因为官方指南明确警告存在 SQL 注入风险。源码树包含 `0.2` 文档与升级 SQL，但已复核 control 的默认版本是 `0.1`；在安装并测试升级前，不应假定具备 `0.2` 的配置管理器或 worker 扩缩容行为。
