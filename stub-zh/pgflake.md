## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/README.md)
- [0.0.1 版本 SQL 对象](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake--0.0.1.sql)
- [C 实现](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake.c)

`pgflake` 0.0.1 试图用 41 位时间、10 位实例和 12 位序列生成可排序的 64 位 Snowflake 风格 ID。但已复核实现既不能可靠保持时间顺序，也不能保证跨会话唯一性，因此只能把它当作已放弃原型，不能当作生产 ID 生成器。

### 演示

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

提取函数只返回编码后的时间、实例与序列组成部分；它们不会修复或验证已生成 ID。评估该原型时仍应保留 `PRIMARY KEY` 或其他唯一约束。

### 完整性缺陷

设计中的配置参数是 `pgflake.instance_id` 与 `pgflake.start_epoch`，但模块通过栈上局部变量注册它们，并且只在库初始化时复制一次值。之后的 `SET` 不会更新生成状态，还可能通过失效存储写入。修复 C 代码之前，不能依赖运行时 GUC 变更。

实现转换 `clock_gettime` 时把 `tv_nsec` 除以 1000，而不是 1000000。因此编码值并非毫秒值，会在每个整秒边界回退，并与其他秒的值区间重叠。它的 `last_time`、`sequence` 与自旋锁也都是后端局部状态，因此使用同一实例 ID 的两个会话可能生成完全相同的值。为每台 PostgreSQL 服务器分配一个实例 ID 并不能区分这些会话。

序列回卷还包含另一处时间算术错误，时钟回拨也不会被拒绝。NTP 无法修复这些实现缺陷。仓库只记录 PostgreSQL 11.4 与 12；除非先修复代码、完成跨后端并发测试并再次审计，否则不要用该版本生成持久标识符。
