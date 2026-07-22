## 用法

来源：

- [官方 README](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/README.md)
- [扩展 SQL](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule--1.0.sql)
- [C 类型实现](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule.c)
- [Haskell 调度实现](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/PGSchedule.chs)

`schedule` 1.0 增加经过验证的 cron 格式基础类型，以及测试或枚举匹配时间戳的函数。它只计算计划，不执行任务。所有计算都使用 UTC，因此不适合必须跟随夏令时切换的本地墙上时间周期任务。

### 核心流程

创建扩展，把 cron 表达式转换为自定义类型，再查询其 UTC 触发时间：

```sql
CREATE EXTENSION schedule;

SELECT schedule_contains('0 */6 * * *'::schedule, '2026-07-22 06:00:00+00');
SELECT schedule_next('0 */6 * * *'::schedule, '2026-07-22 06:00:00+00');
SELECT *
FROM schedule_series(
  '0 */6 * * *'::schedule,
  '2026-07-22 00:00:00+00',
  '2026-07-23 00:00:00+00'
);
```

序列包含从下界到闭合上界之间的匹配值。

### 函数与操作符索引

- `schedule_contains` 测试某时间是否匹配。
- `schedule_next` 和 `schedule_previous` 返回严格晚于或早于输入的匹配值。
- `schedule_floor` 和 `schedule_ceiling` 在输入本身匹配时包含该时刻。
- `schedule_series` 返回有界区间内的匹配值。

扩展提供与文本之间的赋值转换。比较操作符与 B-tree 操作符类按存储的 cron 字符串进行词法比较，而不是比较匹配时间戳集合的语义；写法不同但含义等价的计划可能比较为不相等。

### 运维说明

该类型保留原始且通过验证的表达式文本。`schedule_series` 会让 Haskell 层先构造完整时间戳数组，再由 PostgreSQL 返回行，因此必须限制区间和频率以防内存过度消耗。源码已废弃，来自 PostgreSQL 9.5/Nix 时代的构建，并嵌入 Haskell 运行时；使用前必须验证工具链与 PostgreSQL ABI 兼容性。还应处理无结果和时间戳越界错误，并使用持续维护的调度器实际执行任务。
