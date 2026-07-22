## 用法

来源：

- [1.0 版本 SQL 对象](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj--1.0.sql)
- [类型实现](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.c)
- [扩展控制文件](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.control)

`ksj` 定义了一种定长 PostgreSQL 类型：文本形式使用日文汉字数字，内部值则为 32 位整数。它提供算术、比较、`sum`、`min`、`max`、整数转换及 B-tree 操作符类。

```sql
CREATE EXTENSION ksj;
SELECT '百二十三'::ksj + '七'::ksj;
CREATE INDEX ON ledger (amount_kanji);
```

所复核仓库没有用户文档或回归测试，因此必须针对精确构建，从 C 实现验证可接受语法和输出写法。32 位表示也意味着需要显式测试算术和转换的边界及溢出行为。

不要把显示文本用作稳定交换格式或排序键。若写法本身有意义，应保留原始文本；计算、约束和外部 API 应使用核心数值类型。只有在验证负数、零、畸形输入、转储恢复和升级行为后，才能把 `ksj` B-tree 排序视为数值排序。
