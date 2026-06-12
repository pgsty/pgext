来源：[fsm_core PGXN README](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/README.md)、[control file](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core.control)、[1.1.0 SQL definition](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core--1.1.0.sql)、[example definitions README](https://github.com/Nirajkashyap/fsm/blob/main/apps/fsm-core-example/README.md)。

## 用法

`fsm_core` 是一个有限状态机工具包，用于在 PostgreSQL 中保存 FSM 定义、实例、转换、队列事件和事件日志。机器定义从 JSON 加载，实例按名称和版本创建，事件通过 SQL 函数发送，并可选择使用 `pgmq` 队列。

该扩展在本地 catalog 中为 PG15+。控制文件要求 `ltree` 和 `pgmq`；上游 README 还列出 `pg_jsonschema` 作为要求，Pigsty 包依赖也包含 `postgresql-$v-pg-jsonschema`。

### 核心表与类型

`fsm_core` 会创建枚举 `fsm_state_type`，包含 `atomic`、`compound`、`parallel`、`final` 和 `history`，并创建下列表：

- `fsm_core.fsm_json`：加载后的 FSM 定义。
- `fsm_core.fsm_states`：展开后的状态节点和 ltree 路径。
- `fsm_core.fsm_transitions`：转换规则。
- `fsm_core.fsm_instance`：运行中的实例。
- `fsm_core.fsm_instance_lock`：advisory/concurrency 状态。
- `fsm_core.fsm_instance_queue_event_logs` 和 `fsm_core.fsm_promise_queue_event_logs`：队列事件历史。

### 加载机器定义

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_type    := 'workflow',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

`load_fsm_from_json_v2()` 会用 `fsm_core.fsm_json_schema()` 校验 JSON，把原始定义缓存到 `fsm_json`，然后展开状态和转换。上游示例使用不可变版本目录，例如 `fsm/creditCheck/v01/xstate-fsm.json` 和 `fsm/creditCheck/v02/xstate-fsm.json`；部署后的版本应保持不可变，使已有实例继续按原始定义运行。

### 创建实例

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name     := 'creditCheck',
  input_fsm_version  := 'v01',
  input_fsm_context  := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue  := true
);
```

该函数会检查指定名称的 FSM 是否存在，插入一条 `fsm_instance`，为该实例复制转换授权行，并在 `create_pgmq_queue` 为 true 时创建一个以实例 UUID 命名的 `pgmq` 队列，然后发送 `initialTransition_event`。

返回的 JSON 包含 `queue_created`、`fsm_name`、`fsm_version`、`fsm_instance_id`、`fsm_instance_context`、`send_event_result`、`message` 和 `extra_message`。

### 发送事件

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id                 := '00000000-0000-0000-0000-000000000000'::uuid,
  input_fsm_instance_id_fsm_type         := 'workflow',
  input_fsm_instance_id_fsm_version      := 'v01',
  input_send_to_parent_queue_id          := fsm_core.pg_system_queue_uuid(),
  input_send_to_parent_queue_type        := fsm_core.pg_system_queue_type(),
  input_send_to_parent_queue_id_event_name := fsm_core.pg_system_event_name(),
  input_event_name                       := 'APPROVE',
  input_event_action_type                := 'user',
  input_event_data                       := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                      := 0
);
```

这个辅助函数会用 `pgmq.send()` 写入实例队列，并在 `fsm_instance_queue_event_logs` 中记录事件。对于嵌套 FSM 和 promise 流，`send_event_to_queue_from_fsm_instance_id_v2()` 会根据 `fsmtype` 分派到子 FSM 或 promise 队列辅助函数。

### 解析并推进状态

```sql
SELECT fsm_core.resolve_state_value_v2(
  input_json        := '{"value":"pending"}'::jsonb,
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);

SELECT fsm_core.macrostep_v2(
  event_name        := 'APPROVE',
  input_state_value := ARRAY['pending']::text[],
  fsm_name_param    := 'creditCheck',
  fsm_version_param := 'v01'
);
```

SQL 接口还包括较底层的 `microstep_v2()`、`fsm_worker_v2()`、锁辅助函数、归档辅助函数和 v1 兼容函数。当两个版本都存在时，新用法应优先选择 v2 入口点。

### 依赖与来源注意事项

- 在 `fsm_core` 前启用 `ltree` 和 `pgmq`；同时考虑 `pg_jsonschema`，因为上游要求和 Pigsty 包依赖都列出了它。
- `pgmq` 队列是异步执行模型的一部分，因此队列名称和保留策略应按应用数据来运维。
- 上游仓库没有 `1.1.0` release tag；本 stub 的权威打包来源是 `packages/database-src/pgxn-dist` 中的 SQL/control 文件集以及 PGXN README。
- 相比 SQL 接口，公开文档较为稀疏。除非 SQL 定义或示例应用展示了用途，否则应把未列出的辅助函数视为内部函数。
