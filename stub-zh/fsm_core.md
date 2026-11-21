## 用法

来源：

- [fsm_core-v1.2.0](https://github.com/pgfsm/fsm/releases/tag/fsm_core-v1.2.0)
- [README.md](https://github.com/pgfsm/fsm/blob/fsm_core-v1.2.0/packages/database-src/README.md)
- [fsm_core-1.2.0.zip](https://github.com/pgfsm/fsm/releases/download/fsm_core-v1.2.0/fsm_core-1.2.0.zip)


`fsm_core` 1.2.0 在 PostgreSQL 中保存有限状态机定义、实例、转换和事件历史。其 SQL 接口用于加载经过校验的 JSON、按名称和版本创建实例，并通过 `pgmq` 队列提交事件。

### 安装与配置

本版本要求 PostgreSQL 15 及以上，发布测试覆盖 15–18。必需扩展为 `ltree` 1.2+、`pgmq` 1.4.4+ 和 `pg_cron` 1.5+。将 `pg_cron` 加入现有 `shared_preload_libraries`，修改预加载配置后重启 PostgreSQL，并安装到 `cron.database_name` 指定的数据库。

安装匹配的扩展文件后，超级用户可以创建固定且不可迁移的 `fsm_core` 模式：

```sql
CREATE EXTENSION fsm_core CASCADE;
SELECT fsm_core.version();
```

安装会创建 `fsm_*` 角色及每五秒运行一次的 `fsm_schedule_all_pending` 调度任务。开放应用访问前，应检查角色成员关系、队列权限和调度器。`pg_jsonschema` 为可选依赖：如果它的校验函数存在于 `fsm_core` 中，加载器会报告模式问题，但不会仅凭该检查拒绝定义。

### 加载定义

先用上游编译器验证状态机，并将 JSON 设置到 psql 变量 fsm_json，再执行此调用。1.2.0 的函数签名不再包含 input_fsm_type 参数：

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

同一名称和版本的并发加载由事务级咨询锁串行化。相同 JSON 会返回缓存定义；在已有名称和版本下提交不同 JSON 会报错。修改机器定义时应使用新版本。可选参数 input_dependent_children 用于记录子状态机依赖。

### 创建实例并发送事件

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01',
  input_fsm_context := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue := true
) AS creation_result
\gset
SELECT :'creation_result'::jsonb AS creation_status;
SELECT :'creation_result'::jsonb ->> 'fsm_instance_id' AS fsm_instance_id
\gset
```

返回的 JSON 包含 `fsm_instance_id`、`queue_created`、`send_event_result`、`message` 和 `extra_message`。提交工作前应检查这些字段，因为队列创建和初始事件发送失败可能记录在返回结果中。初始上下文与定义中的静态上下文进行浅合并，调用时提供的对象键优先。

对于此顶层状态机，发送其定义中声明的事件：

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id            := :'fsm_instance_id'::uuid,
  input_fsm_instance_id_fsm_type   := 'fsm',
  input_fsm_instance_id_fsm_version := 'v01',
  input_event_name                 := 'APPROVE',
  input_event_action_type          := 'user',
  input_event_data                 := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                := 0
);
```

该函数将事件入队并写入事件历史。本版本的签名已去掉旧的父队列参数。事件入队不代表工作进程已经处理它，应监控调度器活动、队列积压和实例状态。

### 对象与升级边界

`fsm_json`、`fsm_states` 和 `fsm_transitions` 保存定义；`fsm_instance` 保存实例；`fsm_instance_queue_event_logs` 记录队列事件。`resolve_state_value_v2` 和 `macrostep_v2` 提供较底层的状态解析和转换逻辑。优先使用经过核对的 v2 入口，并限制对安全定义者函数和底层表的访问。

发布说明描述了由 `require_version` 保护的逐版本升级脚本，但本次核对的 1.2.0 ZIP 只有完整安装脚本，没有版本间升级脚本。升级已有实例前应保留数据，并取得经过核验的升级制品。不能重新执行完整安装脚本，也不能假定 `ALTER EXTENSION ... UPDATE` 已有可用的升级路径。
