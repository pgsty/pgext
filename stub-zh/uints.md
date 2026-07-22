## 用法

来源：

- [官方 README](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/README.md)
- [版本 0.9 扩展 SQL](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/uints--0.9.sql)
- [无符号整数实现](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/uint.c)

`uints` 增加名为 `uint2` 与 `uint4` 的无符号 16 位和 32 位整数类型。当数据库需要拒绝负数输入，同时保留定宽存储、算术、比较、位操作以及普通 B-tree 或 hash 索引时，可以使用它们。

### 核心流程

```sql
CREATE EXTENSION uints;

CREATE TABLE device_counters (
    device_id bigint PRIMARY KEY,
    small_counter uint2 NOT NULL,
    large_counter uint4 NOT NULL
);

INSERT INTO device_counters VALUES
    (1, 65535, 4294967295);

SELECT small_counter + 1::uint2,
       large_counter & 255::uint4,
       large_counter >> 8
FROM device_counters;

CREATE INDEX device_counters_large_idx
ON device_counters(large_counter);
```

`uint2` 范围为 `0` 到 `65535`；`uint4` 范围为 `0` 到 `4294967295`。输入超出无符号范围，以及算术溢出或下溢都会报错。

### 操作符、转换与索引

两种类型都支持比较操作符、`+`、`-`、`*`、`/`、`%`，位操作 `&`、`|`、`#`、`~`，以及移位 `<<` 与 `>>`。已定义的 `uint2` 和 `uint4` 混合比较与算术会返回较宽类型。`uint2` 可隐式转换为 `uint4`，窄化转换则是显式的。默认 B-tree 与 hash 操作符类提供等值、排序、分组及索引查找能力。

### 注意事项

扩展 SQL 只定义 `uint2` 与 `uint4`；不会安装 `uint8` SQL 类型。PostgreSQL 有符号类型与这些无符号类型之间的同宽转换被声明为不使用转换函数，因此最高位为一的值可能按原始位模式解释，而不是按数学上等价的有符号值解释。不要依赖 `uint4::integer` 或负数 `integer::uint4` 转换；涉及完整无符号范围时，应通过已校验文本或 numeric 逻辑转换。

除零与范围溢出都会报错。版本 `0.9` 是较旧的 C 实现，生产使用前应验证 ABI 与 PostgreSQL 版本兼容性、转储恢复行为、客户端驱动解码及全部边界值。不需要预加载或重启。
