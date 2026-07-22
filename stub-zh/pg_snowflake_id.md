## 用法

来源：

- [官方 README](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/README.md)
- [官方控制文件](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/pg_snowflake_id.control)
- [官方 GUC 定义](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/config.rs)
- [官方 SQL 函数](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/generate_snowflake_id.rs)
- [官方生成器实现](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/snowflake.rs)

`pg_snowflake_id` 提供 `generate_snowflake_id()`，用于生成 Snowflake 结构的 64 位整数。已复核的 `0.0.1` 实现为每个 PostgreSQL 后端保存一套生成器，因此当并发后端使用相同节点标识时，它不能保证数据库范围内的唯一性。

### 核心流程

每个后端都要在第一次生成数值之前设置数据中心与工作节点标识，然后再调用函数。

```sql
CREATE EXTENSION pg_snowflake_id;

SET pg_snowflake_id.data_center_id = 1;
SET pg_snowflake_id.worker_id = 7;
SET pg_snowflake_id.epoch = '2021-01-01 00:00:00 UTC';

SELECT generate_snowflake_id();
```

位布局依次为 41 位时间戳、5 位数据中心 ID、5 位工作节点 ID 与 12 位序列。单个生成器每毫秒最多产生 4096 个值，耗尽后等待下一毫秒。

### 配置

- `pg_snowflake_id.data_center_id` 接受 `0` 到 `31`，默认值为 `1`。
- `pg_snowflake_id.worker_id` 接受 `0` 到 `31`，默认值为 `1`。
- `pg_snowflake_id.epoch` 选择时间戳纪元，默认值为 `2021-01-01 00:00:00 UTC`。

这三个设置都可由用户修改。后端在第一次调用时初始化并缓存生成器，因此同一会话后续修改设置不会改变已经启用的生成器。

### 唯一性边界

不同 PostgreSQL 后端之间不共享 Rust 互斥量或序列计数器。因此，使用相同数据中心和工作节点 ID 的两个并发会话可能在同一毫秒生成相同数值。会话级设置与连接池复用也使得为每个后端分配持久唯一标识变得困难。

除非经过审计的下游补丁或外部分配方案能在所有后端与节点之间建立唯一性，否则不要在并发环境下把 `generate_snowflake_id()` 用作主键默认值。还要处理明确的时钟回退错误，并验证所选纪元的可用时间范围。
