## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/README.md)
- [1.0 版本 SQL 对象](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup--1.0.sql)
- [C 实现](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup.c)

`calculate_markup` 安装严格且不可变的 C 函数 `price_modifiers._calculate_markup(numeric, numeric[])`。它从二维 `(price, markup)` 数组中选择加价率，或进行线性插值。

### 核心流程

各行必须按价格降序排列。高于第一个断点的价格使用首行加价率，低于最后断点的价格使用末行加价率，断点之间的值线性插值并四舍五入到两位小数。

```sql
CREATE EXTENSION calculate_markup;

SELECT price_modifiers._calculate_markup(
  75::numeric,
  ARRAY[[100, 10], [50, 20], [0, 30]]::numeric[]
);
-- 15.00
```

真实 SQL 名称以下划线开头。README 的两个示例使用 `price_modifiers.c_calculate_markup`，但 1.0 版本并未创建该函数。

### 约束与安全性

实现只检查数组是否为二维。它默认每行恰有两个值、价格断点严格降序且互不相同、至少有一行，并且不存在 NULL 元素。源码没有正确检查 PostgreSQL 的 NULL 位图；空数组、奇异形状、重复或乱序断点及 NULL 元素都可能造成错误结果、除零、越界访问或后端崩溃。调用前必须完整验证数组形状和内容，并显式测试边界价格。

这个已归档项目没有当前 PostgreSQL 兼容性声明。应把它视为不安全的原生代码，只在可丢弃数据库中评估；生产价格规则应优先改写为带明确约束的小型 SQL 实现。
