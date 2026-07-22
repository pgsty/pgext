## 用法

来源：

- [官方 PGXN 分发页面](https://pgxn.org/dist/temporal/)
- [官方函数与运算符参考](https://pgxn.org/dist/temporal/doc/html/reference.html)
- [官方教程](https://pgxn.org/dist/temporal/doc/html/tutorial.html)
- [0.7.1 版源码归档](https://api.pgxn.org/dist/temporal/0.7.1/temporal-0.7.1.zip)

`temporal` 提供固定 16 字节的 `period` 类型，用于带时区时间戳区间，并附带构造器、关系运算符以及 B-tree 与 GiST 索引。它是 2011 年发布的旧扩展，早于 PostgreSQL 内置 range 与 multirange 类型。

### 核心流程

创建扩展，并用默认的双参数构造器存储左闭右开 period。

```sql
CREATE EXTENSION temporal;

CREATE TABLE meeting (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  during period NOT NULL
);

INSERT INTO meeting(during)
VALUES (period(
  timestamptz '2026-07-22 09:00+08',
  timestamptz '2026-07-22 10:00+08'
));
```

默认 GiST 运算符类支持包含与重叠运算符。

```sql
CREATE INDEX meeting_during_gist
ON meeting USING gist (during);

SELECT *
FROM meeting
WHERE during @> clock_timestamp();

SELECT *
FROM meeting
WHERE during && period(
  timestamptz '2026-07-22 09:30+08',
  timestamptz '2026-07-22 11:00+08'
);
```

### 重要对象

- `period`：规范输出是下界包含、上界不包含的时间戳区间。
- `period(timestamptz)`、`period(timestamptz, timestamptz)`：点构造器与默认半开构造器。
- `period_oo`、`period_oc`、`period_co`、`period_cc`：选择端点开闭方式的构造器。
- `empty_period()`、`is_empty(period)`：构造与检查空值。
- `first`、`last`、`prior`、`next`、`length`：端点与持续时间访问器。
- `@>`、`<@`、`&&`、`<<`、`>>`、`&<`、`&>`：包含、被包含、重叠与相对位置运算符。
- `period_intersect`、`period_union`、`minus`：针对 period 的类集合操作。
- `gist_period_ops`、`btree_period_ops`：默认 GiST 与 B-tree 运算符类。

### 语义与迁移边界

`period_union` 遇到不相交且不相邻的输入会报错。`minus` 只能返回一个 period；如果减法会把结果分裂成两个 period，它也会报错。使用显式开闭端点构造器时，必须仔细验证端点语义。

`0.7.1` 相比 `0.7.0` 只修改了 PGXN 元数据；其实现面向远早于当前版本的 PostgreSQL。该类型不是 SQL 标准的 system/application `PERIOD` 语法，也不提供 temporal 主键或外键语义。新设计应优先使用内置 `tstzrange` 与 multirange，除非必须兼容已存储的 `period` 值。迁移前应比较空值、包含端点、相邻关系、union/minus 错误、二进制存储与索引行为，而不能盲目转换。
