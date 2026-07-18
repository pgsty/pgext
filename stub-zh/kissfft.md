## 用法

来源：

- [项目 README](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/README.md)
- [扩展 control 文件](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/kissfft.control)
- [SQL API](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/sql/kissfft.sql)
- [原生聚合实现](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/src/kissfft.c)
- [回归示例](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/test/test.sql)

`kissfft` 0.0.1 是一个 2012 年的 C 扩展，包装 Kiss FFT 算法。它定义 `fft_agg(real)`，收集有序标量序列并返回 `real[]` 形式的类功率谱，还提供 `fft(real[], real)` 将数组位置与频率配对，以及 `fft_print(real[])` 通过消息输出数值。

### 实验性查询形态

输入顺序属于信号的一部分，因此应在聚合内部明确指定：

```sql
CREATE EXTENSION kissfft;

SELECT fft_agg(sample_value ORDER BY sampled_at)
FROM signal_sample;
```

该聚合会把全部输入缓存在 PostgreSQL 数组中，只在最终函数执行变换。内存用量随分组大小增长，大型分组可能耗尽后端内存。应有意地分段和加窗，并用受信数值库验证缩放、频率区间、归一化和实信号对称性。

### 不安全的旧实现

原生最终函数先拆解 PostgreSQL `Datum` 数组，再通过类型转换把紧凑的 `float` 值写入同一存储，然后重建 SQL 数组。这依赖在现代 64 位 PostgreSQL 构建上不安全的内部布局，可能产生错误结果或内存故障。代码还依靠断言检查输入形态，且没有当前兼容矩阵。

不要在生产后端加载未经修改的二进制。必须审计并修复 C 内存处理，增加确定性数值测试和错误输入测试，并在准确 PostgreSQL 大版本上使用 sanitizer 运行。普通分析应使用受维护的数值库处理导出数据，而不是把这个已废弃且 ABI 敏感的实现放入数据库服务器。
