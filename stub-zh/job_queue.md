## 用法

来源：

- [官方扩展控制文件](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue.control)
- [官方上游文档](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/README.md)

`job_queue` — 使用动态 PostgreSQL 后台工作进程实现的数据库内异步作业队列概念验证。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "job_queue";
SELECT extversion
FROM pg_extension
WHERE extname = 'job_queue';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
