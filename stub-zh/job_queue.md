## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/README.md)
- [1.0 版本 SQL 对象](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue--1.0.sql)
- [后台工作进程实现](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue.c)

`job_queue` 是一个已归档的概念验证：语句向 `jobs_queue` 插入数据后，会启动动态后台工作进程。工作进程用 `FOR UPDATE SKIP LOCKED` 领取一行，删除该行，再执行存于 `proc` 的文本以及可选 JSON `args`。

```sql
CREATE EXTENSION job_queue;
INSERT INTO jobs_queue (proc, args)
VALUES ('SELECT process_order($1)', '{"order_id": 42}');
```

它不是安全的生产队列。`proc` 会被拼接为 SQL，而不是解析为有类型的过程，因此任意写入者都可能借后台工作进程的数据库权限执行 SQL。应撤销 `PUBLIC` 对表及启动函数的访问权，且不要让不可信输入控制这些字段。

领取、删除和任务正文位于同一个工作进程事务中，但实现明确留下了未完成的错误路径：任务抛出异常时不会捕获并更新 `error_count` 或 `last_error`。源事务只注册工作进程，不会等待任务成功；动态工作进程槽耗尽时，注册也可能失败。

生产环境应使用具有明确重试、死信、可观测性、取消、幂等和权限语义的受维护队列。此扩展只适合在可丢弃环境中研究 PostgreSQL 后台工作进程。
