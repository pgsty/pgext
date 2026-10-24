## 用法

来源：

- [0.2.0 README](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/README.md)
- [0.2.0 控制文件](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/pg_nanotime.control)
- [0.2.0 SQL 对象](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/sql/pg_nanotime--0.2.0.sql)
- [技术参考](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/docs/TECHNICAL.md)
- [0.2.0 变更记录](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/CHANGELOG.md)
- [PGXN 发布元数据](https://github.com/bunko666-cloud/pg_nanotime/blob/f7583f4609223a53818a7841d5bc52903305ed2f/META.json)

`pg_nanotime` 0.2.0 在 PostgreSQL 17–18 中提供纳秒时间戳与时间间隔类型，适用于事件时间、延迟测量和外部纳秒数据。纳秒存储精度不代表系统时钟具有纳秒准确度。

### 核心用法

```sql
CREATE EXTENSION pg_nanotime;
CREATE TABLE events (at nanotimestamptz DEFAULT clock_ns(), body text);
INSERT INTO events (body) VALUES ('received');
SELECT at, at::timestamptz FROM events;
```

| 类型 | 存储与含义 |
| --- | --- |
| `nanotimestamptz` | 16 字节；按会话时区显示的时刻，覆盖 PostgreSQL 的完整时间戳范围 |
| `nanotimestamp` | 16 字节；不带时区的时间戳 |
| `nanotimestamptz8` | 8 字节；有符号 Unix 纳秒数，范围缩小到 1677–2262 年 |
| `nanointerval` | 16 字节；固定经过时长，不包含日历月份 |

转换为内置时间戳时向下取整到微秒，Unix 纪元之前的时间也一样。绝对时刻转为 bigint 得到 Unix 纳秒数，超出 bigint 范围会报错。不要用双精度浮点数传递较大的整数纳秒值。

### 时长、聚合与序列

时间戳相减得到纳秒时长。0.2.0 新增时间间隔聚合与序列生成：

```sql
WITH samples(ns) AS (VALUES (1250::bigint), (2751::bigint))
SELECT sum(nanointerval_from_ns(ns)) AS total,
       avg(nanointerval_from_ns(ns)) AS mean,
       abs(nanointerval_from_ns(-1250::bigint)) AS magnitude
FROM samples;

SELECT *
FROM generate_series(
  '2026-10-24 00:00:00+00'::nanotimestamptz,
  '2026-10-24 00:00:00.000000003+00'::nanotimestamptz,
  nanointerval_from_ns(1::bigint)
);
```

`sum` 与 `avg` 精确累加，忽略 NULL，无输入时返回 NULL，溢出时报错。平均值舍入到最近的纳秒，恰好一半时远离零。序列包含步长能够到达的终点，支持负步长，拒绝零步长。

整数输入用 `nanointerval_from_ns(bigint)` 保持精确，更大的值用 `nanointerval_from_ns(numeric)`；小数输入会舍入到纳秒。双精度浮点重载在超过 2^53 纳秒后会丢失精度。

其他常用对象包括固定时长分桶函数 `time_bucket_ns`、精确字段提取函数 `extract_ns`，以及 `date_trunc` 和 `timezone`。应显式标注时长字面量的类型：时间戳运算中的无类型字符串可能被解析为 PostgreSQL 的日历间隔。时间戳类型支持 B-tree、hash 与 BRIN minmax 索引；时间间隔类型支持 B-tree 和 hash。

### 时钟与使用边界

`clock_ns()` 每次调用都读取墙上时钟。`now_ns()` 为事务缓存扩展自己的时钟读数，`statement_ns()` 为语句缓存；这些读数与 PostgreSQL 内置时间戳并非同一次采样。普通时钟可能随系统时钟回拨。

`clock_ns_monotonic()` 在单个后端内产生严格递增的值。`clock_ns_global()` 在同一 PostgreSQL 实例的各后端之间实现严格递增，需要以下配置并重启服务器：

```ini
shared_preload_libraries = 'pg_nanotime'
```

保留配置中已有的其他库。普通类型和时钟不要求预加载。全局递增状态位于共享内存，重启后不保留，不能当作持久序列或分布式排序机制。安装需要超级用户；扩展可迁移到其他模式，但未标记为可信扩展。

`nanotime_clock_status()` 提供时钟与单调状态诊断。`nanotime.clock_source` 默认为 `realtime`；可选的 `tai` 依赖操作系统配置的 TAI 偏移量。TAI 值仍按 UTC 格式输出，因此不要在同一列中混用这两种来源。

### 升级与迁移

上游将 0.2.0 标记为 PGXN 测试版。替换库文件与 SQL 文件后，逐库执行升级：

```sql
ALTER EXTENSION pg_nanotime UPDATE TO '0.2.0';
```

列转换辅助过程 `nanotime_migrate_column` 默认使用 16 字节类型，bigint 来源也一样；它通过 ALTER TABLE 修改列并重建索引。执行前应检查转换规则和锁开销。TimescaleDB 集成是可选功能；整数时间与连续聚合有各自的要求，不能仅凭能够存储纳秒时间戳就推断兼容。
