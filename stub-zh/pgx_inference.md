## 用法

来源：

- [FEP 18 quick-start guide](https://www.postgresql.fastware.com/knowledge-base/quick-start-guides)
- [FEP 18 knowledge data management manual](https://www.postgresql.fastware.com/hubfs/_Global/Manuals/FEP-v18forx86-KnowledgeDataManagementUserGuide.pdf)

Fujitsu Enterprise Postgres 18 中的 `pgx_inference` 1.0 管理 ONNX 文本嵌入流水线模型，通过数据库所在主机上的 Triton 执行推理。模型必须符合厂商规定的输入、输出及算子要求。

### 核心用法

先配置 Triton、模型路径、gRPC/TLS 与工作进程容量，将 `pgx_inference` 加入 `shared_preload_libraries` 并重启。以扩展所有者身份启用扩展并启动加载调度进程：

```sql
CREATE EXTENSION pgx_inference CASCADE;
SELECT pgx_inference.pgx_launch_load_launcher();
```

使用厂商的 `pgx_aimodel_tool` 导入兼容模型。对于以 `sample-model-v1` 为名导入的模型，请求加载并查看状态：

```sql
SELECT * FROM pgx_inference.pgx_list_model_metadata();
SELECT pgx_inference.pgx_load_model('sample-model-v1');
SELECT * FROM pgx_inference.pgx_triton_model_status;
```

加载是异步过程。确认模型状态为 `is_available = true` 后，再生成嵌入：

```sql
SELECT pgx_inference.pgx_onnx_embed('sample-model-v1', 'PostgreSQL replication');
```

### 运行边界

模型导入、加载、执行及模式访问分别授权。应通过 mTLS 限制 Triton 访问，以落实模型级权限控制。备份必须包含大对象；删除模型前应先卸载，并确认卸载完成。此功能需要 Fujitsu 运行环境。
