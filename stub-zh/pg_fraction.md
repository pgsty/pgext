## 用法

来源：

- [官方 README](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/README.md)
- [官方扩展控制文件](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/pg_fraction.control)

`pg_fraction` 增加 `fraction` 类型，用于精确存储分子和分母，并提供基本算术、比较、聚合、类型转换和 B-tree 索引。上游把它描述为示例扩展，而不是面向商业用途的软件。

### 核心流程

创建扩展后，可以存储分数字面量并使用其运算符：

```sql
CREATE EXTENSION pg_fraction;

CREATE TABLE ratios (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value fraction NOT NULL
);

INSERT INTO ratios (value) VALUES ('1/2'), ('2/4'), ('3/5');

SELECT value,
       value + '1/3'::fraction AS increased,
       get_value(value) AS approximate_value
FROM ratios
ORDER BY value;
```

分数会自动约分，因此 `2/4` 与 `1/2` 这类等价输入会规范化为相同值。

### 重要对象

- `fraction` 存储分子和分母，并接受以斜杠分隔的输入。
- `+`、`-`、`*` 和 `/` 执行分数算术。
- 比较运算符与 B-tree 运算符类支持排序和索引。
- `max(fraction)` 与 `min(fraction)` 对分数值执行聚合。
- `get_value(fraction)` 返回近似浮点值。

### 限制与注意事项

- 实现要求每个输入部分最多五位十进制数字且数值不超过 `99999`；算术结果超出限制时会报错。
- 位数检查发生在约分之前，因此即使规范化结果可以容纳，写法较大的可约分输入仍可能被拒绝。
- 上游比较实现会把分数转换为浮点数。对于非常接近或较大的值，应验证精度与排序行为；如果业务严格要求精确有理数比较，不应采用该类型。
- 采用这个示例扩展前，应在目标 PostgreSQL 构建上测试除零、符号处理、类型转换、溢出、转储恢复和索引排序。
