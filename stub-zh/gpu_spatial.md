## 用法

来源：

- [上游来源 1](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/README.md)
- [上游来源 2](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/ext/gpu_spatial.control)
- [上游来源 3](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/ext/gpu_spatial--0.1.sql)

`gpu_spatial` 0.1 将外部 gpujoind 服务中常驻图层的 CUDA 空间运算暴露为 SQL。

### 核心流程

服务运行且空间源表已准备好时，`gpu_load` 加载命名快照，`gpu_unload` 将其移除。`gpu_join` 返回匹配标识符及可选的相交面积，`gpu_relate`、`gpu_dwithin` 与 `gpu_cluster_intersecting` 提供其他空间运算。可通过 `gpu_last_stats` 检查上次操作。

```sql
CREATE EXTENSION gpu_spatial;
SELECT gpu_status();
SELECT * FROM gpu_layers();
```

### 运行边界

安装仅限超级用户。加载与卸载函数撤销了 PUBLIC 执行权限，查询访问仍需明确审查权限。文档环境使用 NVIDIA Hopper sm90 及独立 CUDA 守护进程，并非支持任意显卡。常驻图层不会自动跟随源表变更，必须显式重新加载。坐标系和地理距离语义应符合上游函数约定。
