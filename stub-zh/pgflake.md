## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/README.md)
- [0.0.1 版本 SQL 对象](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake--0.0.1.sql)
- [C 实现](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake.c)

`pgflake` 生成可排序的 64 位 Snowflake 风格 ID：相对自定义纪元的 41 位毫秒时间、10 位实例 ID 和 12 位每毫秒序列。生成 ID 前，必须一致地配置 `pgflake.instance_id` 与 `pgflake.start_epoch`。

```sql
CREATE EXTENSION pgflake;
CREATE TABLE event (
  event_id bigint PRIMARY KEY DEFAULT pgflake_generate(),
  payload jsonb NOT NULL
);
SELECT pgflake_extract_time(event_id),
       pgflake_extract_instance(event_id),
       pgflake_extract_sequence(event_id)
FROM event;
```

每个并发活动生成器都必须有唯一的 `pgflake.instance_id`，重复使用会产生碰撞。所有节点还必须共享相同的 `pgflake.start_epoch`，而 41 位时间字段从该纪元起约可使用 69 年。

这个已放弃的 0.0.1 实现只记录了 PostgreSQL 11.4 与 12。应通过 NTP 保持时钟同步，并测试时钟回拨、每毫秒超过 4096 个 ID 的序列耗尽、故障转移实例号分配、后端并发、纪元溢出、备份恢复和升级。即使生成器配置正确，数据库唯一约束仍然不可省略。
