

## 用法

> [snowflake: PostgreSQL 的 Snowflake ID 序列](https://github.com/pgEdge/snowflake)

提供基于 `int8` 和 `sequence` 的唯一 ID 生成，使用 Snowflake 格式，适用于分布式系统。

```sql
CREATE EXTENSION snowflake;
```

### 配置

在 `postgresql.conf` 中设置节点标识符（必需，值为 1-1023）：

```ini
snowflake.node = 1
```

### 函数

| 函数 | 描述 |
|---|---|
| `snowflake.nextval([sequence regclass])` | 生成下一个 Snowflake ID（未指定序列时使用内部序列） |
| `snowflake.currval([sequence regclass])` | 返回序列的当前值 |
| `snowflake.get_epoch(snowflake int8)` | 提取时间戳（自 2023-01-01 起的秒数） |
| `snowflake.get_count(snowflake int8)` | 提取计数部分（每毫秒重置） |
| `snowflake.get_node(snowflake int8)` | 提取节点标识符 |
| `snowflake.format(snowflake int8)` | 返回包含 `node`、`ts` 和 `count` 字段的 JSONB |

### 示例

```sql
-- 生成一个 Snowflake ID
SELECT snowflake.nextval();
-- 136169504773242881

-- 配合命名序列使用
CREATE SEQUENCE orders_id_seq;
SELECT snowflake.nextval('orders_id_seq'::regclass);

-- 提取各组成部分
SELECT snowflake.get_epoch(136169504773242881);
-- 1704996539.845

SELECT to_timestamp(snowflake.get_epoch(136169504773242881));
-- 2024-01-11 13:08:59.845-05

SELECT snowflake.get_node(136169504773242881);
-- 1

SELECT snowflake.format(136169504773242881);
-- {"id": 1, "ts": "2024-01-11 13:08:59.845-05", "count": 0}

-- 用作默认列值
CREATE TABLE orders (
  id int8 DEFAULT snowflake.nextval() PRIMARY KEY,
  data text
);
```
