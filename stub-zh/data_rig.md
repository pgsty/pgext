## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/README.md)
- [1.0 版本 SQL 对象](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig--1.0.sql)
- [C 实现](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.c)
- [扩展 control 文件](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.control)

`data_rig` 是实验性原生类型，把多维 OLAP fact 表示为已排序的编码整数集合，并提供包含操作符和默认 GiST 操作符类。

### 核心流程

SQL 脚本在 GiST 支持函数签名中引用 `cube` 类型，但 control 文件的 `requires` 没有声明 `cube`。必须先显式安装它；`CREATE EXTENSION data_rig CASCADE` 无法推断这项依赖。

```sql
CREATE EXTENSION cube;
CREATE EXTENSION data_rig;

CREATE TABLE fact_demo (id bigint, dimensions fact);
INSERT INTO fact_demo
VALUES (1, fact(ARRAY[to_fact_number(1, 10), to_fact_number(2, 20)]));
CREATE INDEX fact_demo_dimensions_idx
  ON fact_demo USING gist (dimensions gist_fact_ops);

SELECT id
FROM fact_demo
WHERE dimensions @> fact(ARRAY[to_fact_number(1, 10)]);
```

`fact(integer[])` 会对编码成员排序、去重，并拒绝含 NULL 的数组。`@>` 与 `<@` 测试包含关系，`fact_intersect(fact, fact)` 返回公共成员，`to_fact_number(dimension, value)` 把维度和值打包成一个 `int4`。

### 限制与安全性

文本输入不可用：类型的 `fact_in(cstring)` 始终抛出 `fact_in() not implemented`，因此必须通过 `fact(integer[])` 构造值，不能使用字面量或文本转换。输出函数主要用于显示，不能提供可往返的文本格式。

上游只有一行 README，没有兼容矩阵或有意义的回归测试。自定义类型与 GiST 实现包含未经检查的指针和数组运算；畸形值或索引边界情况可能产生错误结果或导致后端崩溃。不要把 `data_rig` 用于生产数据。若要评估 `data_rig`，应使用可丢弃数据库，把每个索引结果与顺序扫描对比，覆盖空集合和边界集合，并准备在任何二进制变更后删除、重建索引。
