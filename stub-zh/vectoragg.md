## 用法

来源：

- [官方 1.0 版 SQL](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg--1.0.sql)
- [官方 C 实现](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg.c)
- [官方 control 文件](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg.control)
- [官方归档仓库](https://github.com/cybertec-postgresql/vectoragg)

`vectoragg` 提供 C 函数，用于在坐标区间上对等间隔的 `real[]`/`double precision[]` 样本求和或切片，并提供固定倍数降采样辅助函数。它是文档极少且已归档的版本 `1.0` 项目；只有在审查下述源码缺陷后才能使用。

### 核心流程

对于坐标从 `0` 到 `3` 的四个样本，可求和或提取从 `1` 到 `3` 的左闭右开坐标区间：

```sql
CREATE EXTENSION vectoragg;

SELECT array_sum(
  ARRAY[1, 2, 3, 4]::double precision[],
  0, 3, 1, 3
);

SELECT array_clamp(
  ARRAY[1, 2, 3, 4]::double precision[],
  0, 3, 1, 3
);
```

函数按 `(tend - tstart) / (array_length - 1)` 推导采样间隔。本例求和结果为 `5`，切片为 `{2,3}`。

### API

扩展同时安装 `real[]` 与 `double precision[]` 重载：

- `array_sum(samples, tstart, tend, astart, aend) RETURNS double precision`：对选中样本求和。
- `array_clamp(samples, tstart, tend, astart, aend) RETURNS array`：返回选中切片。
- `array_decimate(samples) RETURNS array`：设计用于每 10 个样本求平均。
- `array_hundreth(samples) RETURNS array`：设计用于每 100 个样本求平均；拼写错误是 SQL API 的一部分。

### 注意事项

不要对任意长度调用发布版 `array_decimate` 或 `array_hundreth`。官方 C 源码中处理剩余元素的 `while` 循环没有递增下标，因此长度不是 10 或 100 整数倍的数组可能让 PostgreSQL 后端永久循环。这两个函数必须修补源码并补充回归测试后才能使用。

`array_sum`/`array_clamp` 会拒绝包含空值的数组，但在除以 `n - 1` 前没有验证空数组、单元素数组或坐标顺序。应确认 `n >= 2`、边界有限且有序，并验证采样边界的左闭右开取整行为。仓库已经归档且 README 为空，不存在上游维护或兼容保证。
