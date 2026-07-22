## 用法

来源：

- [Official README](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/README.md)
- [Extension control file](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/pg_snowid.control)
- [Function and shared-memory implementation](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/src/lib.rs)

pg_snowid 在 PostgreSQL 中生成按时间排序的 64 位 Snowflake 类标识符及紧凑 Base62 表示。它在共享内存中维护各键对应的生成器，分布式部署中每个 PostgreSQL 实例都需要不同的节点 ID。

### 核心流程

把 `pg_snowid` 加入 `shared_preload_libraries`，重启 PostgreSQL，创建扩展，并在生成任何 ID 前设置节点：

```ini
shared_preload_libraries = 'pg_snowid'
```

```sql
CREATE EXTENSION pg_snowid;

SELECT snowid_set_node(5);

CREATE TABLE event (
  id bigint PRIMARY KEY DEFAULT snowid_generate(1),
  payload jsonb NOT NULL
);

INSERT INTO event(payload) VALUES ('{"kind":"created"}');
SELECT id, snowid_get_timestamp(id) FROM event;
```

传给 `snowid_generate` 的生成器键是类似整数的 OID 值。应为每个 ID 域维护稳定且有记录的键；它只选择共享生成器状态，并不会编码进最终 ID。

### 重要对象

- `snowid_set_node` 要在任何生成器创建前设置 0 到 1023 的节点 ID；默认值为 1。
- `snowid_get_node` 返回当前节点 ID。
- `snowid_generate` 和 `snowid_generate_base62` 生成一个数值或 Base62 ID。
- `snowid_try_generate` 和 `snowid_try_generate_base62` 不会推进逻辑时间，而会返回 null。
- `snowid_generate_batch` 和 `snowid_try_generate_batch` 一次最多保留 100,000 个 ID。
- `snowid_get_timestamp` 和 `snowid_get_timestamp_base62` 提取毫秒时间戳。
- `snowid_stats` 报告节点及生成器状态。

### 运维说明

在持续高负载下，普通生成函数会推进逻辑时间而不是等待，因此嵌入时间戳可能领先于墙上时钟。try 变体保持墙上时间语义，但可能返回 null 或短批次。节点 ID 不会自动推导；不同实例复用节点 ID 有碰撞风险，而且生成器一旦存在，更改节点需要重启 PostgreSQL。不同生成器键彼此独立，而键不在位布局中，因此在没有应用级测试及约束策略前，不要假设不同键之间全局唯一。
