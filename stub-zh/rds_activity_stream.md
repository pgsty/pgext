## 用法

来源：

- [AWS：启动 Aurora 数据库活动流](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.Enabling.html)
- [AWS：Aurora PostgreSQL 扩展版本矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`rds_activity_stream` 是 Amazon Aurora PostgreSQL Database Activity Streams 背后的云厂商托管组件。AWS 当前扩展矩阵列出的版本为 `1.7`。它采集已配置的访问与变更事件，并将每个 Aurora 集群的加密审计数据发送到同一 AWS Region 内的 Kinesis stream。

这是没有 SQL 激活步骤的 headless 托管功能；应通过 RDS 控制台、CLI 或 `StartActivityStream` API 为 Aurora 集群启用。例如：

```shell
aws rds start-activity-stream \
    --mode async \
    --kms-key-id my-kms-key-arn \
    --resource-arn my-cluster-arn \
    --apply-immediately
```

Aurora PostgreSQL 支持同步或异步模式。应提供用户选择的 KMS key，而不是默认 key。使用 `--apply-immediately` 启动活动流会立即重启集群；不使用该参数时，启用操作会等到维护窗口，且在重启完成前活动流不会启动。安排变更前，应核对受支持的引擎补丁版本与实例类型要求。
