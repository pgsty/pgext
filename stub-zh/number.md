## 用法

来源：

- [官方 README](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/README.md)
- [版本 1 扩展 SQL](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number--1.sql)
- [类型实现](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number.c)

`number` 是可变宽度的有符号 64 位整数类型。零值占一个字节，其他值按数值大小占两到九个字节，同时保持与 `bigint` 相同的数值范围。它面向以小值为主的整数列，只有当行与索引空间节省超过扩展和转换开销时才有价值。

### 核心流程

安装扩展，以 `number` 作为列类型，并在代表性数据上与 `bigint` 比较存储大小：

```sql
CREATE EXTENSION number;

CREATE TABLE compact_counters (
  counter_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value number NOT NULL
);

INSERT INTO compact_counters (value)
VALUES (0), (127), (128), (2147483648);

SELECT value, pg_column_size(value)
FROM compact_counters
ORDER BY value;
```

文本输入输出使用普通有符号整数写法。超出有符号 64 位范围的值会被拒绝；它不是任意精度 decimal 类型，也没有 scale。

### 转换、比较与索引

- 赋值转换把 `integer` 和 `bigint` 转为 `number`。
- 隐式转换把 `number` 转为 `bigint`，从而由 PostgreSQL 整数算术处理表达式。
- 版本 1 定义了两个 `number` 之间以及与 `bigint` 交叉类型的比较操作符。
- 默认 `number_ops` B-tree 操作符类支持有序索引。

```sql
CREATE INDEX compact_counters_value_idx
ON compact_counters (value);

SELECT *
FROM compact_counters
WHERE value >= 100::bigint
ORDER BY value;
```

README 中“没有操作符”的表述早于版本 1 SQL 已包含的比较支持；应以 SQL 文件为权威对象接口。扩展没有返回 `number` 的原生算术操作符，因此在结果类型很重要的表达式中要检查隐式转换与结果类型。

### 运维说明

小值可以节省空间，但大值在普通元组开销之前最多占九个字节，可能比定宽 `bigint` 更大。应针对实际分布、表压缩、索引、比较和写入负载做基准，而不是假定一定节省。

版本 1 发布于 2017 年，可重定位，且没有声明预加载或扩展依赖。自定义磁盘表示使扩展二进制成为备份和升级兼容性的一部分：恢复依赖列前，必须在目标端安装并测试扩展。生产采用前还要验证当前 PostgreSQL 头文件兼容性、逻辑复制、转储恢复、二进制升级、ORM 所用转换与查询计划。
