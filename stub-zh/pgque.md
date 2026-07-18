## 用法

来源：

- [扩展 control 文件](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/pgque.control)
- [0.1.0 版 SQL 实现](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/pgque--0.1.0.sql)
- [基础练习](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/test/basic.sql)

`pgque` 0.1.0 是一个小型纯 SQL 作业队列。作业作为行保存在 `pgque_jobs` 中，按 `priority` 和 `id` 升序排列。worker 调用 `pgque_lock()` 获取会话级 advisory lock，然后显式销毁并解锁返回的作业。

### 最小 worker 流程

```sql
CREATE EXTENSION pgque;

SELECT (pgque_enqueue(
  'mail',
  10,
  '{"message_id":123}'::jsonb
)).*;

SELECT * FROM pgque_lock('mail');

-- After processing the returned job in this same session:
SELECT pgque_destroy(42);
SELECT pgque_unlock(42);
```

应把 `42` 替换为返回的 `id`。advisory lock 属于数据库会话，在事务回滚后仍然存在，并在会话结束时释放。`pgque_destroy()` 会删除行但不会释放锁，因此每条 worker 路径都必须调用 `pgque_unlock()` 或关闭会话。

### 投递与协调限制

该设计既不提供 exactly-once 投递，也没有可见性超时。worker 崩溃会释放锁，并让仍存在的行再次可用；外部副作用却可能已经发生。在外部副作用前后删除作业会产生不同的丢失和重复窗口。消费者必须使用幂等键和持久应用协议。

bigint advisory-lock 命名空间与其他应用锁共享，因此作业 ID 可能与无关约定冲突。0.1.0 版没有重试次数、计划时间、死信状态、租约 owner、超时恢复、payload 模式或保留策略，也没有当前兼容性或并发测试套件。只有在压力测试队列饥饿、事务、连接池模式、worker 取消和崩溃恢复后才能使用；生产工作应优先选择受维护的队列实现。
