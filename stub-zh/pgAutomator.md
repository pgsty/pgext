## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/README.md)
- [0.0.1 版本 schema](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/sql/pgAutomator--0.0.1.sql)
- [扩展 control 文件](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/pgAutomator.control)

`pgAutomator` 安装 pgAutomator 作业调度器的数据库端目录。由于 control 文件名包含大写字母，在区分大小写的系统上应为扩展名加引号。

```sql
CREATE EXTENSION "pgAutomator";
SELECT job_category_id, category
FROM pgautomator.job_category
ORDER BY job_category_id;
```

安装脚本创建固定的 `pgautomator` schema，以及调度枚举、agent 与 SMTP 配置、作业、步骤、计划、连接信息和执行日志。必须另行运行 pgAutomator agent 来轮询和执行作业；只安装该扩展不会启动调度器。

0.0.1 是 2017 年的早期 schema，README 仅有一行。它会在普通表中保存可选 SMTP 与数据库密码，并注明部分调度模式尚不支持。部署前应限制 schema 权限、审查凭据存储，并验证配套外部 agent。
