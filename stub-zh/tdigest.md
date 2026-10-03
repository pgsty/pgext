## 用法

来源：

- [README.md](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/README.md)
- [tdigest.control](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest.control)
- [Changes](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/Changes)
- [tdigest--1.4.6--1.4.7.sql](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest--1.4.6--1.4.7.sql)

`tdigest` 1.4.7 提供可合并的近似排名统计。可存储部分数据的摘要，再合并计算分位数，无需对全部原始数据排序。

### 核心工作流

```sql
CREATE EXTENSION tdigest;
SELECT tdigest_percentile(v, 100, ARRAY[0.5, 0.95, 0.99])
FROM generate_series(1, 1000) AS g(v);
CREATE TABLE digest_daily AS
SELECT current_date AS day, tdigest(v, 100) AS digest
FROM generate_series(1, 1000) AS g(v);
SELECT tdigest_percentile(digest, 0.95) FROM digest_daily;
```

### 函数与精度

`tdigest(value, compression)` 创建摘要；`tdigest(digest)` 合并已存储的摘要。`tdigest_percentile` 估算分位数，`tdigest_percentile_of` 估算排名，均提供标量和数组形式。`tdigest_add` 与 `tdigest_union` 更新或合并摘要；`tdigest_count`、`tdigest_sum` 与 `tdigest_avg` 返回计数及截尾聚合。截尾聚合的上下界参数表示分位数阈值，而非原始数据值边界。`tdigest_is_valid` 检查序列化值。

`compression` 必须在 10 到 10000 之间。较大值通常以更多内存和 CPU 换取更高精度，但不提供固定误差上界。应使用代表性数据与精确结果对照，并在合并状态时保持压缩参数一致。

### 升级与边界

1.4.7 修复了相邻值分位数插值、大计数反向排名、可复用状态的终结处理及 NULL 压缩参数处理，并将更多标量函数标记为可安全并行执行。应同时更新已安装的共享库和 SQL，并执行 `ALTER EXTENSION tdigest UPDATE TO '1.4.7'`。用 `tdigest_is_valid` 检查已存储摘要；更严格的验证可能拒绝旧版本接受的畸形值。安装需要超级用户权限，扩展可重定位且无需预加载。发布元数据声明最低 PostgreSQL 13，更新日志也记录了 PostgreSQL 19 构建修复，但这些信息不构成性能保证。
