## 用法

来源：

- [官方 README](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/README.md)
- [版本 1.0 扩展 SQL](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions--1.0.sql)
- [C 实现](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions.c)

`pg_variant_functions` 从保存为 `TINYINT[]` 的基因型剂量数组计算检出率与替代等位基因频率。它是一个已停止维护的小型实验项目，并非通用基因组工具集；只有其精确数组约定与数据模型一致时才应使用。

### 核心流程

先安装依赖的 `tinyint` 扩展。基因型位置使用从一开始的 `INTEGER[]` 下标寻址；非 NULL 基因型值会作为替代等位基因剂量求和，再除以已检出基因型数量的两倍。

```sql
CREATE EXTENSION tinyint;
CREATE EXTENSION pg_variant_functions;

SELECT alternate_allele_frequency(
    ARRAY[0, 1, 2, NULL]::tinyint[],
    ARRAY[1, 2, 3, 4]::integer[]
);

SELECT *
FROM summarize_variant(
    ARRAY[0, 1, 2, NULL]::tinyint[],
    ARRAY[1, 2]::integer[],
    ARRAY[3, 4]::integer[]
);
```

`alternate_allele_frequency(genotypes, indices)` 为选中位置返回 `REAL` 值；若这些位置都没有已检出基因型则返回 NULL。

`summarize_variant(genotypes, subset1, subset2)` 返回 `variant_summary` 复合类型，字段如下：

- `call_rate` 与 `maf` 对应完整基因型数组。
- `subset1_call_rate` 与 `subset1_maf` 对应第一个可选下标数组。
- `subset2_call_rate` 与 `subset2_maf` 对应第二个可选下标数组。

子集传入 NULL 时，对应字段保持 NULL。尽管字段名为 `maf`，实现直接报告替代等位基因频率，不会把大于 `0.5` 的值换算成次要等位基因频率。

### 注意事项

代码期望一维剂量数组与合法的正数、从一开始的下标。它不会校验生物学倍性或基因型取值范围，也不应接收零、负数或其他非法下标。空数组与空子集需要由应用处理。两个函数都声明为 immutable 与 parallel safe，但因需要扫描数组，声明成本为 `10000`。版本 `1.0` 已标记为停止维护；用于分析或临床场景前，应与可信实现交叉验证数值结果。
