## 用法

来源：

- [项目 README](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/README.md)
- [扩展 control 文件](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/query_histogram.control)
- [1.1 版 SQL API](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/sql/query_histogram--1.1.sql)
- [执行器与共享内存实现](https://github.com/tvondra/query_histogram/tree/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/src)

`query_histogram` 1.1 在共享内存中记录集群级语句和事务耗时直方图。它按耗时区间报告计数与累计时间，但不保留查询标识或文本。

### 预加载并读取直方图

配置库和采集参数，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'query_histogram'
query_histogram.sample_pct = 5
query_histogram.bin_width = 10
query_histogram.bin_count = 1000
query_histogram.dynamic = true
```

安装 SQL 对象并查询视图：

```sql
CREATE EXTENSION query_histogram;

SELECT * FROM query_histogram;
SELECT * FROM xact_histogram;
SELECT query_histogram_get_reset();
```

`query_histogram.bin_width` 以毫秒计，`query_histogram.bin_count` 上限为 1000。区间数为零会禁用采集，但 hook 仍然存在。动态模式允许运行时修改，但会增加共享内存加锁。

### 采样、重置与开销

每次被采样的执行都会更新由 System V semaphore 保护的共享段。应从较低的 `query_histogram.sample_pct` 开始，用代表性并发进行基准测试，并记住采样计数不是准确总数。区间边界也会丢失分布细节，而且该扩展不能识别哪个归一化语句造成了慢区间。

`query_histogram_reset()` 会清除所有用户的共享观测，因此应限制执行权限，并在重置前协调采集器。SQL 把读取状态和重置函数声明为 `IMMUTABLE`，但结果实际依赖可变共享状态；不要依赖规划器围绕这些函数的缓存语义。源码早于当前 PostgreSQL 大版本，且没有发布当前兼容矩阵，因此必须针对准确目标版本和其他每个执行器 hook 扩展进行编译与加载测试。
