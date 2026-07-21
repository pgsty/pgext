## 用法

来源：

- [argm 1.1.1 README](https://github.com/bashtanov/argm/blob/1.1.1/README.md)
- [Extension control file](https://github.com/bashtanov/argm/blob/1.1.1/argm.control)
- [SQL definitions](https://github.com/bashtanov/argm/blob/1.1.1/argm--1.1.1.sql)

`argm` 提供了多态聚合函数 `argmax`、`argmin` 和 `anyold`。这些函数在分组时返回选定行的值，避免了在行可以通过一个或多个可排序键选择时进行连接或窗口函数处理。

### 核心工作流程

```sql
CREATE EXTENSION argm;

SELECT customer_id,
       argmax(order_id, total, ordered_at) AS largest_order
FROM orders
GROUP BY customer_id;
```

`argmax(value, key...)` 返回属于按字典序最大键元组的 `value`。`argmin` 选择最小的键元组。其他键可以打破平局，而无需构建复合值：

```sql
SELECT device_id,
       argmax(reading, measured_at, sequence_no) AS latest_reading
FROM measurements
GROUP BY device_id;
```

使用 `anyold(value)` 时，任何分组成员都可以接受：

```sql
SELECT account_id, anyold(display_name)
FROM account_aliases
GROUP BY account_id;
```

### 关键对象

- `argmax(value, key [, key ...])` 选择与最大键元组关联的值。
- `argmin(value, key [, key ...])` 选择与最小键元组关联的值。
- `anyold(value)` 返回聚合状态中的任意非空值。

这些聚合函数接受任何数据类型；键类型必须支持排序。SQL 定义是并行安全的，并包括组合和序列化函数以进行部分聚合。

### 语义与注意事项

键元组使用整个元组的一个排序方向和一个排序规则，null 键排在最后。如果完整键元组相等，则选择的值未指定；当需要确定结果时，请添加稳定的最终键。与其他 PostgreSQL 聚合函数一样，空输入返回 `NULL`。

`argm` 1.1.x 版本要求 PostgreSQL 9.6 或更高版本。该扩展可重定位。从 1.0.3 升级到 1.1.x 需要删除并重新创建扩展，因为聚合状态发生了变化；而从 1.1.0 到 1.1.1 的升级不会改变公共 SQL 接口。
