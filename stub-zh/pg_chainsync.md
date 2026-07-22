## 用法

来源：

- [官方 README](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/README.md)
- [官方扩展控制文件](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/pg_chainsync.control)
- [官方 Rust 软件包清单](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/Cargo.toml)

`pg_chainsync` 是一个概念验证性质的 PostgreSQL 后台工作进程，用来抓取区块链区块、事件和计划任务，并把它们交给用户定义的 SQL 处理函数。当前审阅的上游 README 明确提醒可能出现破坏性变更，因此应把它用于受控实验，不要假定其摄取契约稳定。

### 预加载与生命周期

工作进程要求在 `shared_preload_libraries` 中加入 `pg_chainsync.so` 并重启 PostgreSQL。把 `chainsync.database` 设置为工作进程应操作的数据库，然后在该库创建扩展：

```conf
shared_preload_libraries = 'pg_chainsync.so'
chainsync.database = 'chains'
```

```sql
CREATE EXTENSION pg_chainsync;

SELECT chainsync.restart();
SELECT chainsync.stop();
```

`chainsync.restart()` 把注册任务重新加载到工作进程；`chainsync.stop()` 停止执行这些任务。

### 注册区块处理函数

处理函数接收带类型的链对象以及任务的 `jsonb` 元数据。例如，一个 EVM 区块监视任务可以写成：

```sql
CREATE TABLE evm_blocks (
  block_number bigint PRIMARY KEY,
  block_hash text NOT NULL
);

CREATE FUNCTION store_block(block chainsync.EvmBlock, job jsonb)
RETURNS evm_blocks
LANGUAGE sql
AS $$
  INSERT INTO evm_blocks(block_number, block_hash)
  VALUES (block.number, block.hash)
  ON CONFLICT (block_number) DO UPDATE SET block_hash = EXCLUDED.block_hash
  RETURNING *
$$;

SELECT chainsync.register(
  'ethereum_blocks',
  '{
    "ws": "wss://provider.example/ws",
    "evm": {"block_handler": "store_block"}
  }'::jsonb
);

SELECT chainsync.restart();
```

README 还记录了 `chainsync.EvmLog` 事件处理函数、一次性任务、cron 任务、预加载任务，以及用来协调事件处理与区块可用性的区块等待钩子。由于项目被明确标为不稳定，应按已安装版本确认 `chainsync.register` 签名和任务 JSON 模式。

### 配置与可靠性

- `chainsync.evm_ws_permits` 限制共享同一 EVM WebSocket 服务商的并发任务数。
- `chainsync.evm_blocktick_reset` 控制缩小的区块范围何时重置。
- `chainsync.svm_rpc_permits` 限制单个任务内并发的 SVM RPC 抓取数。
- `chainsync.svm_signatures_buffer` 限制缓存的 SVM 签名数量。
- 服务商速率限制、断连、历史范围限制、链终局性和重组都是外部故障模式。应持久化检查点和链标识，并让处理函数在重试与重放时保持幂等。
- 回调成功本身不代表已达到终局性。同步数据用于不可逆决策前，应定义确认策略和重组修正路径。
- 处理函数使用工作进程的权限执行数据库写入。应约束所有者与 `search_path`、使用模式限定对象、验证任务 JSON、保护服务商凭据，并测试事务和错误行为。
