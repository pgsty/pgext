## 用法

来源：

- [目录版本对应的官方 README](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/README.md)
- [目录版本对应的实现](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/nextid.c)

`nextid` 0.1 根据序列值、分片编号和当前系统时钟生成 64 位时间分片 ID。它可用于本地生成近似按时间排序的 ID，但唯一性依赖调用方维持速率、分片和时钟约束。

### 核心流程

```sql
CREATE EXTENSION nextid;
CREATE SEQUENCE order_id_seq;

-- Use a globally assigned shard number in the range 0..8191.
SELECT c_next_id(nextval('order_id_seq'), 42);

CREATE TABLE orders (
    id bigint PRIMARY KEY DEFAULT c_next_id(nextval('order_id_seq'), 42),
    payload jsonb NOT NULL
);
```

### 函数与布局

- `c_next_id(bigint, integer)` 将实现中固定 epoch 之后的毫秒数、左移 10 位的分片号以及序列值对 1024 的余数组合起来。
- 低 10 位每经过 1024 个序列值就重复；其余低位空间预期容纳 13 位分片标识。
- 函数为 strict 并返回 bigint，但不会校验或掩码处理分片参数。

### 唯一性边界

- 集中分配分片 ID，并将其限制在 0 到 8191。负数或更宽的值可能覆盖时间戳部分并破坏布局。
- 单个分片在同一毫秒内不要生成超过 1024 个 ID，否则序列余数会重复。
- 系统时钟回拨可能产生重复值，或使 ID 不再按创建顺序排列。应监控时钟并使用数据库唯一约束，不能只把位布局当作保证。
- 更换 epoch、分片分配或算法时，必须为历史 ID 制定兼容方案。
