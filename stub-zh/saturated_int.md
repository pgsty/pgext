## 用法

来源：

- [官方 README](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/README.md)
- [扩展 SQL 与对象定义](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/saturated_int--0.0.1.sql)
- [饱和算术实现](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/src/saturated_int.c)

`saturated_int` 版本 `0.0.1` 提供有符号 32 位整数类型，其输入与算术溢出时会钳制到 `-2147483648` 或 `2147483647`，而不是抛出 `integer out of range`。

### 核心流程

```sql
CREATE EXTENSION saturated_int;

SELECT 999999999999999::saturated_int;
SELECT 2147483640::saturated_int + 100::saturated_int;
SELECT (-2147483648)::saturated_int / (-1)::saturated_int;
```

这些表达式的结果依次是 `2147483647`、`2147483647` 和 `2147483647`。

### 对象

- `saturated_int`：与 int4 等宽的自定义类型，范围从 `-2147483648` 到 `2147483647`。
- 比较运算符：`=`、`<>`、`<`、`<=`、`>` 和 `>=`。
- 算术运算符：`+`、`-`、`*`、`/` 和 `%`。
- `sum(saturated_int)`：累加器采用饱和语义的聚合。
- 用于索引、排序、分组与等值查找的 B-tree 和 hash 运算符类。

### 转换与算术说明

从 `bigint` 转换时会钳制到类型范围。`integer` 与 `saturated_int` 之间存在赋值转换，但扩展没有定义混合类型运算符；与内置整数比较或计算时，应显式转换两个操作数。

饱和是一种数据模型选择，而不只是错误处理：不同的越界输入会折叠成同一个边界值，聚合溢出也会钳制。除以零仍然报错。只有在确实希望丢失数值幅度时才应使用该类型。不需要预加载或重启。
