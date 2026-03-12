

## 用法

> [btree_gist: B 树等价的 GiST 操作符类](https://www.postgresql.org/docs/current/btree-gist.html)

为通常仅支持 B 树索引的数据类型提供 GiST 索引操作符类。支持将等值运算与范围运算结合使用的排除约束。

```sql
CREATE EXTENSION btree_gist;
```

### 支持的数据类型

`int2`、`int4`、`int8`、`float4`、`float8`、`numeric`、`timestamp with time zone`、`timestamp without time zone`、`time with time zone`、`time without time zone`、`date`、`interval`、`oid`、`money`、`char`、`varchar`、`text`、`bytea`、`bit`、`varbit`、`macaddr`、`macaddr8`、`inet`、`cidr`、`uuid`、`bool` 以及所有 `enum` 类型。

### 距离运算符

为数值和时间类型提供 `<->` 运算符，用于最近邻搜索。

### 示例

```sql
-- 整数列上的 GiST 索引
CREATE INDEX idx ON test USING GIST (a);
SELECT * FROM test WHERE a < 10;

-- 最近邻搜索
SELECT *, a <-> 42 AS dist FROM test ORDER BY a <-> 42 LIMIT 10;

-- 排除约束：每个笼子只能放一种动物
CREATE TABLE zoo (
  cage   integer,
  animal text,
  EXCLUDE USING GIST (cage WITH =, animal WITH <>)
);

INSERT INTO zoo VALUES (1, 'lion');    -- OK
INSERT INTO zoo VALUES (1, 'tiger');   -- ERROR: 冲突的键值
INSERT INTO zoo VALUES (2, 'tiger');   -- OK

-- 每个房间不重叠时间范围的排除约束
CREATE TABLE reservations (
  room int,
  during tsrange,
  EXCLUDE USING GIST (room WITH =, during WITH &&)
);
```
