## 用法

来源：

- [官方上游文档](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/doc/tinyint.md)
- [官方扩展 SQL 定义](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/tinyint.sql.in.c)
- [官方回归测试](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/sql/tinyint.sql)

`tinyint` 增加单字节有符号整数类型，文档给出的取值范围为 -127 至 127。它支持整数式类型转换、比较、算术、聚合与索引，但行对齐可能抵消预期的存储节省，而且这份较老的上游代码需要在目标 PostgreSQL 版本上进行兼容性测试。

### 核心流程

创建扩展，并仅将该类型用于确定不会超出其狭窄范围的值：

```sql
CREATE EXTENSION tinyint;

CREATE TABLE sensor_flags (
    sensor_id bigint PRIMARY KEY,
    state tinyint NOT NULL CHECK (state BETWEEN 0 AND 3)
);

INSERT INTO sensor_flags VALUES (1, 2::tinyint);

SELECT state, count(*)
FROM sensor_flags
GROUP BY state
ORDER BY state;
```

从 `smallint`、`integer` 与 `bigint` 转换时使用赋值类型转换，超出范围会被拒绝。从 `tinyint` 到更宽整数类型的转换则是隐式的。

### 运算符、聚合与索引

- 算术支持一元正负号以及 `+`、`-`、`*`、`/`；与更宽整数类型混合运算时返回更宽的类型。
- `tinyint` 可与 `smallint`、`integer` 或 `bigint` 比较，两种参数顺序都受支持。
- 扩展定义了 `min`、`max`、`sum` 与 `avg`；`sum` 使用 `bigint` 作为转换状态及结果，`avg` 使用 PostgreSQL 的整数平均值终结函数。
- `btree_tinyint_ops` 与 `hash_tinyint_ops` 支持标量 B-tree 和哈希索引。`_tinyint_ops` 支持为 `tinyint[]` 数组建立 GIN 索引。

### 限制

算术与类型转换发生溢出时会报错，因此批量加载与迁移前应验证输入。它是有符号类型，不能直接替代其他数据库中的无符号 `TINYINT`。

尽管数据本身只占一个字节，PostgreSQL 仍可能依据列顺序和相邻类型加入对齐填充。若只是为了节省空间，应先测量实际元组与索引大小。本次核验的上游修订停留在 2017 年，公开兼容性材料也较旧；每个计划使用的 PostgreSQL 大版本都必须经过构建、安装、升级、转储恢复、运算符、聚合与索引测试。
