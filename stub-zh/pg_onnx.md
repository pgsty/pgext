## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/kibae/pg_onnx/blob/ce634755381b3fb27b5724a3c58f39228d94eef8/README.md)
- [函数参考](https://github.com/kibae/pg_onnx/blob/ce634755381b3fb27b5724a3c58f39228d94eef8/docs/functions.md)
- [扩展源码](https://github.com/kibae/pg_onnx/tree/ce634755381b3fb27b5724a3c58f39228d94eef8/pg_onnx)

`pg_onnx` 从 PostgreSQL 运行 ONNX 推理。动态后台工作进程托管 `onnxruntime-server`，并为每个导入模型复用一个运行时会话，避免每个客户端后端各自分配大型运行时。

```sql
CREATE EXTENSION pg_onnx;
SELECT pg_onnx_import_model(
  'score_model', 'v1', $1::bytea,
  '{"cuda": false}'::jsonb, 'reviewed model'
);
SELECT pg_onnx_execute_session(
  'score_model', 'v1', '{"input": [[1.0, 2.0]]}'::jsonb
);
```

模型字节应作为参数传入，不要使用调用者控制的服务器文件路径。应限制导入、删除、会话和执行函数。ONNX 模型与可选自定义操作库都是可执行原生输入：只接受经审查、有哈希的工件，并约束扩展库路径。输出是模型结果，不是确定性的数据库事实。

应规划 CPU、GPU、RAM、worker 槽、超时、取消、模型大小、会话回收和并发。应测试 worker 崩溃/重启、模型元数据事务回滚、故障转移、备库、备份恢复、CUDA/ONNX Runtime ABI 和升级。触发器推理会给写事务增加远程计算，可能阻塞或导致事务失败；除非同步语义不可缺少，应优先异步推理。
