## 用法

来源：

- [Official PGXN1.1.0 distribution](https://pgxn.org/dist/fsm_core/1.1.0/)
- [1.1.0 README](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/README.md)
- [1.1.0 control](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/fsm_core.control)
- [1.1.0 SQL](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/fsm_core--1.1.0.sql)
- [1.1.0 metadata](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/META.json)

`fsm_core` 是一个有限状态机工具包，用于在 PostgreSQL 中保存 FSM 定义、实例、转换和事件日志。机器定义从 JSON 加载，实例按名称和版本创建，事件通过 SQL 函数发送，并可选择使用 `pgmq` 队列。

本文描述 PGXN 1.1.0 发行版，要求 PostgreSQL 15 及以上、`ltree` 1.2 及以上，以及 `pgmq` 1.4.4 及以上。固定模式为 `fsm_core`，安装需要超级用户。这个 SQL 扩展不要求预加载或重启。当前仓库迁移是另外的来源，可能提供不同的 API。

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

`load_fsm_from_json_v2()` 使用 `fsm_core.fsm_json_schema()` 检查 JSON，展开状态和转换，然后把原始定义缓存到 `fsm_json`。在 1.1.0 中，违反 JSON 模式只会产生 NOTICE，加载仍会继续；模式校验的异常语句已被注释。应在加载前验证定义，不能依赖这个检查拒绝所有无效定义。部署后的定义和版本标识应保持稳定，使已有实例继续按原始定义运行。执行示例前，应将经过校验的状态机 JSON 设置到 psql 变量 fsm_json。

### 创建实例

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name     := 'creditCheck',
  input_fsm_version  := 'v01',
  input_fsm_context  := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue  := true
) AS creation_result
\gset
SELECT :'creation_result'::jsonb AS creation_status;
SELECT :'creation_result'::jsonb ->> 'fsm_instance_id' AS fsm_instance_id
\gset
```

PGXN 1.1.0 函数会检查指定名称的 FSM 是否存在，插入一条 `fsm_instance`，并为该实例复制转换授权行。在 `create_pgmq_queue` 为 true 时，它尝试创建以实例 UUID 命名的 `pgmq` 队列，并发送 `initialTransition_event`。队列创建和初始事件发送失败会被捕获，并记录在返回的 JSON 中。发送后续事件前，必须确认 `queue_created` 为 true，并检查 `send_event_result`、`message` 和 `extra_message`，确认初始事件发送成功。psql 示例保留返回的 `fsm_instance_id`，供下一次调用使用。

### 发送事件

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id                 := :'fsm_instance_id'::uuid,
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

### 依赖与运行

安装 `fsm_core` 前启用 `ltree` 和 `pgmq`。发行版 README 还要求 `pg_jsonschema` 0.3.3 及以上，但控制文件和 META 的依赖列表没有列出它。JSON 加载器调用 `fsm_core.jsonschema_validation_errors`，而发行版 SQL 脚本没有定义这个函数；加载定义前应确认该模式下已有上游要求的 JSON 模式校验辅助函数。仅将依赖安装到其他模式，并不能提供这个带模式限定的函数。

队列事件作为应用数据持久保存。设置 `create_pgmq_queue => true` 会请求创建实例队列及其初始事件；继续操作前应检查返回状态。仍需要消费者处理队列工作。发送事件本身不保证异步工作进程已经执行转换。应一并检查队列保留策略、消费者和权限。

向应用角色开放机器创建或任意事件提交前，应检查函数和表授权。SQL 接口还包含旧 v1 和较底层的辅助函数；本版本应使用已经核实的 v2 入口。
