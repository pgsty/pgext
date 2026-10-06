## 用法

来源：

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_gpu/alohadb_gpu--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu--1.0.sql)
- [contrib/alohadb_gpu/alohadb_gpu.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu.c)
- [contrib/alohadb_gpu/alohadb_gpu.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu.control)

`alohadb_gpu` 1.0 在基于 PostgreSQL 18 的 AlohaDB 分支中提供批量向量距离、Top-K 选择和矩阵乘法。它使用分支内置的向量类型，不能直接当作原版 PostgreSQL 的 pgvector 插件。

### 核心用法

```sql
CREATE EXTENSION alohadb_gpu;
SELECT alohadb_gpu_available();
SELECT * FROM alohadb_gpu_topk(
  '[0.1,0.2,0.3]'::vector,
  ARRAY['[0.1,0.0,0.2]'::vector, '[0.4,0.5,0.6]'::vector], 1);
```

### 运行边界

`alohadb_gpu_batch_l2` 与 `alohadb_gpu_batch_cosine` 返回各候选的距离，`alohadb_gpu_topk` 返回候选位置与 L2 距离，`alohadb_gpu_matmul` 对二维 float4 数组做矩阵乘法。输入维度必须匹配。

`alohadb.gpu_device_id` 选择 CUDA 设备，`alohadb.gpu_min_batch_size` 控制何时将批量运算交给 GPU。小批次或无 CUDA 支持时使用 CPU 实现，`alohadb_gpu_available` 报告可用性。大数组仍消耗数据库进程内存。安装需要超级用户；核对的模块不要求预加载或重启。功能是否可用取决于 AlohaDB 的构建方式，而不只是机器是否有 GPU。
