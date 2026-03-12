

## 用法

> [pg_rational: 64 位精确分数运算](https://github.com/begriffs/pg_rational)

`pg_rational` 扩展提供了以 64 位（与 float 大小相同）存储的精确分数运算，并支持使用 Stern-Brocot 树查找中间分数。

```sql
CREATE EXTENSION pg_rational;
```

### 数据类型

- **`rational`**：分数类型（分子/分母）
- **`ratt`**：用于元组转换的辅助类型

### 基本运算

```sql
SELECT '1/3'::rational + '2/7'::rational;  -- 13/21
SELECT '3/4'::rational * '2/3'::rational;  -- 1/2
```

### 函数

```sql
-- 化简分数
SELECT rational_simplify('36/12');  -- 3/1

-- 查找中间分数（Stern-Brocot 树）
SELECT rational_intermediate('1/2', '2/3');  -- 3/5
```

### 类型转换

```sql
SELECT 0.263157894737::float::rational;  -- 5/19
SELECT '-1/2'::rational::float;          -- -0.5
SELECT 1::rational;                       -- 1/1
```

### 动态行排序

一个核心使用场景是在不重新编号的情况下维护可排序的行顺序：

```sql
CREATE TABLE todos (
    prio rational UNIQUE DEFAULT nextval('todos_seq')::integer,
    what text NOT NULL
);

-- 在优先级 1 和 2 之间插入
INSERT INTO todos VALUES (rational_intermediate('1', '2'), 'new task');
-- prio 变为 3/2，其他行不受影响
```

### 索引支持

`rational` 列支持 Btree 和 Hash 索引。
