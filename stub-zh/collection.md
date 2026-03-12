

## 用法

> [collection: 用于 PL/pgSQL 的键值集合数据类型](https://github.com/aws/pgcollection)

`collection` 扩展提供了两种内存优化的集合数据类型，用于 PL/pgSQL 函数中。

```sql
CREATE EXTENSION collection;
```

### 数据类型

- **`collection`**：文本键的键值对（最长 32,767 字符），按创建顺序存储
- **`icollection`**：64 位整数键的键值对，支持稀疏数组

两种类型都支持类型修饰符来指定元素类型：

```sql
DECLARE
    c1  collection('date');
    ic1 icollection('int4');
```

### 下标访问

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['Japan'] := 'Tokyo';
    RAISE NOTICE '%', t_capital['USA'];  -- Washington, D.C.
END $$;
```

### 核心函数

| 函数 | 说明 |
|----------|-------------|
| `add(coll, key, value)` | 添加元素 |
| `count(coll)` | 元素数量 |
| `delete(coll, key)` | 删除元素 |
| `exist(coll, key)` | 检查键是否存在 |
| `find(coll, key)` | 获取值 |
| `first(coll)` | 将迭代器移至起始位置 |
| `last(coll)` | 将迭代器移至末尾位置 |
| `next(coll)` | 迭代器前进 |
| `prev(coll)` | 迭代器后退 |
| `key(coll)` | 当前键 |
| `value(coll)` | 当前值 |
| `copy(coll)` | 创建副本 |
| `sort(coll)` | 按键排序 |
| `keys_to_table(coll)` | 所有键作为集合返回 |
| `values_to_table(coll)` | 所有值作为集合返回 |
| `to_table(coll)` | 键值对作为表返回 |

### 迭代器示例

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['United Kingdom'] := 'London';
    t_capital['Japan'] := 'Tokyo';

    t_capital := first(t_capital);
    WHILE NOT isnull(t_capital) LOOP
        RAISE NOTICE 'Capital of % is %', key(t_capital), value(t_capital);
        t_capital := next(t_capital);
    END LOOP;
END $$;
```

### 稀疏数组（icollection）

`icollection` 支持非连续整数索引，并能区分 NULL 值和未初始化的键：

```sql
DECLARE sparse icollection;
BEGIN
    sparse[1] := 'first';
    sparse[1000000] := 'millionth';  -- 间隙不会浪费内存
END;
```
