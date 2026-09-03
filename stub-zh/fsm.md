## 用法

来源：

- [官方文档](https://github.com/brunoenten/pg_fsm/blob/f781e13a985330b45eb78195cc1df18eacd84dd3/README.md)
- [扩展控制文件](https://github.com/brunoenten/pg_fsm/blob/f781e13a985330b45eb78195cc1df18eacd84dd3/fsm.control)
- [官方仓库](https://github.com/brunoenten/pg_fsm)

`fsm` 为应用表提供由触发器强制执行的有限状态机。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `fsm`：

```sql
CREATE EXTENSION fsm;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Example business table
CREATE TABLE public.orders (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  customer_id bigint NOT NULL,
  total numeric(12,2) NOT NULL
);

-- Attach FSM columns/triggers
SELECT fsm.add_to_table('public.orders'::regclass);

-- Define transitions from start -> pending -> paid|cancelled
SELECT fsm.add_transition('public.orders'::regclass, 'start',   'create',  'pending');
SELECT fsm.add_transition('public.orders'::regclass, 'pending', 'pay',     'paid');
SELECT fsm.add_transition('public.orders'::regclass, 'pending', 'cancel',  'cancelled');

-- Insert row (FSM columns must remain default on INSERT)
INSERT INTO public.orders (customer_id, total) VALUES (10, 99.99);

-- Append one event to move start -> pending
UPDATE public.orders SET new_event = 'create' WHERE id = 1;

-- Append another event to move pending -> paid
UPDATE public.orders SET new_event = 'pay' WHERE id = 1;

SELECT id, fsm_previous_state, fsm_current_state
FROM public.orders
WHERE id = 1;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `fsm.event` | TYPE | 扩展创建的用户数据类型。 |
| `fsm.add_transition` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fsm.add_to_table` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fsm.machines` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `fsm.run_machine` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fsm.add_callback` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fsm.append_event` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `fsm.events_callbacks` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 9.6, 10, 11, 12, 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 扩展会在 `fsm` 下固定或创建模式对象；权限与备份审查应包含这些对象。
