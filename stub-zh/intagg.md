

## 用法

> [intagg: 整数聚合器和枚举器（已过时）](https://www.postgresql.org/docs/current/intagg.html)

提供整数聚合器和枚举器。这些现在只是内置 `array_agg()` 和 `unnest()` 函数的封装。

```sql
CREATE EXTENSION intagg;
```

### 函数

| 函数 | 说明 |
|---|---|
| `int_array_aggregate(integer)` | 将整数聚合为数组（`array_agg()` 的封装） |
| `int_array_enum(integer[])` | 将数组展开为行（`unnest()` 的封装） |

### 示例

```sql
-- 将整数聚合为数组
SELECT id_left, int_array_aggregate(id_right) AS rights
FROM many_to_many
GROUP BY id_left;

-- 将整数数组展开为行
SELECT int_array_enum(ARRAY[1, 2, 3, 4]);
-- 返回: 1, 2, 3, 4（作为单独的行）
```

注意：此模块已过时。新代码请使用内置的 `array_agg()` 和 `unnest()` 函数。
