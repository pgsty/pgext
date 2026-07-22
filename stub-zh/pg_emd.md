## 用法

来源：

- [Official README](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/README.md)
- [Version 0.1.0 SQL declarations](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/sql/pg_emd--0.1.0.sql)
- [PostgreSQL-facing Rust implementation](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/src/pg_extension.rs)
- [Cargo and PostgreSQL compatibility](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/Cargo.toml)

`pg_emd` 0.1.0 用于计算推土机距离。当前 `emd` 实现针对“数组位置代表分箱”的一维直方图执行精确 O(n) 计算；`emd_weighted` 则使用随机动态树嵌入，近似计算带权多维点之间的距离。该扩展属于研究实现，目前只构建 PostgreSQL 17 版本。

### 一维直方图

```sql
CREATE EXTENSION pg_emd;

SELECT emd(
  ARRAY[0.5, 0.3, 0.2]::double precision[],
  ARRAY[0.2, 0.3, 0.5]::double precision[]
);
```

两个数组长度必须相同；一对空数组返回零。每个值表示位置 `0, 1, ...` 分箱中的质量。调用前应校验值有限且非负、两侧总质量相同，因为函数不会强制这些分布不变量。

### 加权多维点

```sql
SELECT emd_weighted(
  '[{"point":[1.0,2.0],"weight":0.7},
    {"point":[3.0,4.0],"weight":0.3}]'::json,
  '[{"point":[10.0,20.0],"weight":1.0}]'::json
);

SELECT tree_distance(
  ARRAY[0.0, 0.0]::double precision[],
  ARRAY[10.0, 10.0]::double precision[]
);
```

`tree_distance` 暴露底层随机树度量用于检查；它不是精确欧氏距离。

### 规划器与输入的重要注意事项

SQL 把三个函数都声明为 `IMMUTABLE PARALLEL SAFE`，但 `emd_weighted` 和 `tree_distance` 每次调用都会用随机网格偏移构造树嵌入，相同参数的结果可能变化，因此其 volatility 声明不安全。在实现和 SQL 声明修正前，不要把这两个函数用于表达式索引、生成列、分区边界、缓存常量表达式或约束。

`emd_weighted` 从第一个点推导维度，却不会完整校验所有点维度一致、坐标有限、权重非负或总质量相等；转换时还会过滤非数字坐标项。调用方必须先校验并规范化 JSON。任意一侧为空时返回零，这也未必符合应用对缺失数据的语义。

README 中的性能和近似数字只是项目测量，并非工作负载保证。采用前应固定构建，用精确实现验证数值误差与可重复性，并按生产分布规模测试内存和延迟。
