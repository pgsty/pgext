## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/README.md)
- [版本 0.0.1 的数据库 schema](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/sql/pgAutomator--0.0.1.sql)
- [扩展 control 文件](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/pgAutomator.control)

`pgAutomator` 安装早期 pgAutomator 作业调度器的数据库端目录。它在 `pgautomator` schema 中创建调度、代理、作业、步骤、队列、连接、邮件和执行日志对象。扩展本身不含后台工作进程或可执行代理，因此单独安装不会执行作业；还需要与之匹配的外部 pgAutomator 代理。

规范扩展名包含大写字母，因此必须在 SQL 中引用。

```sql
CREATE EXTENSION "pgAutomator";

SELECT pgautomator.schema_version();
SELECT job_category_id, category
FROM pgautomator.job_category
ORDER BY job_category_id;
```

### 对象模型

- `pgautomator.agent` 和 `pgautomator.agent_active` 跟踪外部代理及其轮询心跳。
- `pgautomator.job`、`pgautomator.step`、`pgautomator.job_schedule` 和各调度子类型表描述工作、控制流与日历时间。
- `pgautomator.job_execution_queue` 保存到期执行项；`pgautomator.queue_job_execution()` 根据已启用的作业和调度填充队列。
- `pgautomator.job_log` 和 `pgautomator.step_log` 记录执行状态。`pgautomator.job_kill` 是供外部代理解释的请求表。
- `pgautomator.step_connection`、`pgautomator.agent_smtp`、`pgautomator.job_email` 和 `pgautomator.step_email` 保存连接与通知配置。

安装脚本会写入初始作业类别，并物化从扩展创建时起约一年的 `pgautomator.dim_date`。长期运行的安装必须刷新该物化视图，并在初始范围过期前验证调度器日期逻辑。

### 运维注意事项

- 上游版本 `0.0.1` 发布于 2017 年，README 只有一行说明。它没有维护中的部署、代理协议、升级或当前 PostgreSQL 兼容文档。
- `pgautomator.schedule_type` 中存在 `ON_IDLE` 与 `ON_TRIGGER_FILE` 成员，但源码注释明确说这些模式不受支持。只能使用已经与匹配代理验证过的行为。
- SMTP 密码和数据库密码以普通表的明文列保存。写入凭据前，必须限制 `pgautomator` schema、备份、副本和日志的访问。
- SQL 与批处理步骤可通过外部代理执行高权限命令。应把作业定义写入视为特权管理，并限制代理的操作系统账户和数据库角色。
- 该 schema 是早期且紧耦合的数据模型，包含循环的延迟外键与生成的调度状态。生产使用前必须测试创建、排队、重试、时区、夏令时切换、故障恢复和代理兼容性。
