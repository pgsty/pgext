## 用法

来源：

- [上游 README](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/README.md)
- [扩展 control 文件](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/supa_queue.control)
- [1.0.4 版 SQL](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/supa_queue--1.0.4.sql)
- [dbdev 软件包页面](https://database.dev/mansueli/supa_queue)

`supa_queue` `1.0.4` 版是基于 `pg_net` 和 `pg_cron`、面向 Supabase 的 SQL 队列。插入 `job_queue` 会触发异步 GET、POST 或 DELETE 请求；数据表跟踪进行中任务和工作槽，定时函数收集响应并重试失败任务。消费者 URL 和服务角色密钥从 Vault secret 读取。

### 示例

复核安装 SQL、配置 Vault 并创建依赖扩展后执行：

```sql
CREATE EXTENSION pg_net;
CREATE EXTENSION pg_cron;
CREATE EXTENSION supa_queue;
INSERT INTO job_queue (http_verb, payload, url_path)
VALUES ('POST', '{"task":"refresh"}', '/jobs');
SELECT job_id, status, retry_count FROM job_queue ORDER BY job_id DESC;
```

安装会立即创建 cron 任务，其中一个定时函数会两次休眠 20 秒。若干函数使用 `SECURITY DEFINER`，两个函数被强制放入 `public`；若不撤销执行权限，`request_wrapper(...)` 可以发起任意 HTTP 请求。数据表启用了 RLS 却没有策略。安装前必须修正所有权/搜索路径和模式位置、撤销公共执行、添加授权/策略、约束目标地址、保护 Vault 访问，并定义重试/幂等及保留行为。
