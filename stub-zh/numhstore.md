## 用法

来源：

- [官方 README](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/README.md)
- [0.1.7 版扩展 SQL](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/numhstore--0.1.7.sql)
- [扩展 control 文件](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/numhstore.control)

`numhstore` 增加名为 `inthstore` 和 `floathstore` 的数值键值存储。它们沿用 `hstore` 的文本形态与二进制表示，同时为稀疏整数或浮点映射增加算术、汇总与聚合能力。

### 核心流程

该扩展依赖 PostgreSQL 的 `hstore` 扩展。

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION numhstore;
```

先把 `hstore` 值转换为与存储值匹配的数值类型，再执行逐键算术。

```sql
SELECT 'a=>1,b=>2'::hstore::inthstore
     + 'a=>4,c=>8'::hstore::inthstore;

SELECT 'cpu=>1.5,mem=>2.25'::hstore::floathstore * 2::numeric;
```

可以跨行聚合稀疏映射。缺失键遵循各运算符自己的语义，因此应测试具体操作，不能直接套用普通标量 NULL 行为。

```sql
CREATE TEMP TABLE counters(v inthstore);
INSERT INTO counters VALUES
  ('ok=>2,error=>1'),
  ('ok=>3,retry=>4');

SELECT sum(v), avg(v)
FROM counters;
```

### 重要对象

- `inthstore`：整数值存储，可与 `hstore` 双向隐式转换，也可隐式转换为 `floathstore`。
- `floathstore`：浮点/数值存储，可与 `hstore` 双向隐式转换。
- `+`、`-`、`*`、`/`：存储之间或存储与标量之间的逐键算术。
- `sum(inthstore)`、`sum(floathstore)`、`avg(inthstore)`、`avg(floathstore)`：聚合稀疏映射。
- `array_count(integer[])`、`array_count(text[])`、`array_add(text[], integer[])`：构造或合并计数映射。
- `hstore_sum_up(hstore)`、`hstore_array(hstore)`、`hstore_length(hstore)`：求值总和、以数组返回值或统计键数量。

### 语义与限制

自定义类型与 `hstore` 二进制兼容，但数值仍通过 hstore 风格文本表示。无效数值字符串会在算术或转换解析时失败。仅一侧存在的键在不同运算符下行为并不统一：加减法保留类似并集的键，而乘除法可能按实现产生 NULL 或类似零的结果。应使用实际数据测试缺失键、NULL、除零、溢出与浮点舍入。

隐式转换可能让重载表达式产生意外结果；在 API 与模式边界应使用显式转换。固定的上游源码没有给出当前 PostgreSQL 兼容矩阵，因此生产使用前需在目标服务器验证 `0.1.7` 版。
