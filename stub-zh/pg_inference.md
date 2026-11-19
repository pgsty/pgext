## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference.c)

`pg_inference` `1.0` 提供使用 AlohaDB 内核向量类型的 ONNX 推理。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_inference;
SELECT * FROM pg_list_models();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

pg_load_model 接受模型字节，pg_infer 与 pg_batch_infer 使用 AlohaDB 向量值。库必须编译启用 ONNX Runtime，否则模型操作会失败。pg_inference.device 选择 CPU/CUDA 能力，pg_inference.max_models 限制后端本地缓存。模型需匹配预期张量形状。仅加载可信模型并评估后端内存；加载模型不等于跨会话持久部署。

此模块对应上游旧名 `alohadb_inference`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
