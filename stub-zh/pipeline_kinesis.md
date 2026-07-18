## 用法

来源：

- [上游 README](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/README.md)
- [扩展 control 文件](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis.control)
- [SQL 安装脚本](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis--0.9.0.sql)
- [扩展实现](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis.c)

`pipeline_kinesis` `0.9.0` 版是历史 PipelineDB 扩展：它启动动态后台工作进程，消费 AWS Kinesis 分片并把记录复制到 PipelineDB 流中，同时在 `pipeline_kinesis` 模式保存端点、消费者和分片序列号元数据。

### 历史工作流

以下内容假定旧 PipelineDB 的 `CREATE STREAM` 等语义，并非当前 PostgreSQL 示例：

```sql
CREATE EXTENSION pipeline_kinesis;
SELECT pipeline_kinesis.add_endpoint('endpoint', 'us-west-2');
SELECT pipeline_kinesis.consume_begin('endpoint', 'test', 'foo_stream');
SELECT pipeline_kinesis.consume_end('endpoint', 'test', 'foo_stream');
```

2016 年的代码面向 PipelineDB 和旧 AWS C++ SDK，没有当前 PostgreSQL 兼容路径，而且会在服务器内启动联网后台进程。凭据文件路径、端点元数据、工作进程上限、重放位置和授权都必须显式管控。消费者查找实现还存在不安全的参数数量处理。不要把它用于现代或生产集群；应改用仍在维护的外部 Kinesis 消费者。
