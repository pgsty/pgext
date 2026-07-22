## 用法

来源：

- [AWS Database Activity Streams 概览](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.html)
- [AWS：启动 Aurora 数据库活动流](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.Enabling.html)
- [AWS Aurora PostgreSQL 扩展版本矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`rds_activity_stream` 是 Aurora PostgreSQL Database Activity Streams 背后的 AWS 托管内部组件。AWS 扩展矩阵列出的版本为 1.7，但用户不能通过 SQL 安装或调用它。应在 Aurora DB 集群上启用托管功能，将加密的数据库活动记录发送到同一区域内自动创建的 Amazon Kinesis 数据流。

### 启动托管流

通过 RDS 控制台、API 或 AWS CLI 在集群范围操作。以下命令立即启动异步流：

```shell
aws rds start-activity-stream \
  --mode async \
  --kms-key-id <kms-key-arn> \
  --resource-arn <cluster-arn> \
  --apply-immediately
```

Aurora PostgreSQL 支持 `sync` 和 `async` 模式。应根据所需审计保证与工作负载延迟选择模式。提供服务可以使用的客户管理 KMS key；此流程不接受默认 AWS KMS key。

### 重启与范围

启动 Database Activity Streams 需要重启数据库集群。使用 `--apply-immediately` 时，AWS 会立即重启集群；否则变更会等待下一个维护窗口，且在重启前流保持停止。设置适用于集群中当前与未来的所有实例。对于 Aurora global database，应在每个区域集群上独立启用和管理活动流。

### 运维与安全

安排变更前应核对 AWS 引擎版本、补丁、区域和实例类型前置条件。该功能会创建 Kinesis stream，并产生 Kinesis 及相关消费和存储费用。应建设并监控消费端、加密访问、保留策略、告警、吞吐与故障处理；只启动活动流并不等于完成了审计管道。活动记录属于敏感数据，应以最小权限授予 KMS、RDS、Kinesis 与消费端权限。不要执行 `CREATE EXTENSION rds_activity_stream`，也不要假定它暴露受支持的 SQL API。
