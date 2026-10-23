## 用法

来源：

- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/README.md)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/pg_nanotime.control)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/sql/pg_nanotime--0.1.0.sql)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/docs/TECHNICAL.md)

`pg_nanotime` 在 PostgreSQL 17–18 中提供纳秒时间戳与时间间隔类型。0.1.0 是源码扩展；纳秒精度不代表系统时钟具有纳秒准确度。

### 核心用法

```sql
CREATE EXTENSION pg_nanotime;
CREATE TABLE events (at nanotimestamptz DEFAULT clock_ns(), body text);
INSERT INTO events (body) VALUES ('received');
SELECT at, at::timestamptz FROM events;
```

`nanotimestamptz` 与 `nanotimestamp` 分别是带时区和不带时区的 16 字节类型。`nanotimestamptz8` 使用 8 字节，范围缩小至 1677–2262 年。`nanointerval` 表示固定经过时长，与 PostgreSQL 按日历计算的时间间隔不同。转换为内置时间戳会丢失不足一微秒的精度。

### 时钟与使用边界

`clock_ns()` 读取墙上时钟；`now_ns()` 为事务缓存扩展读取的时钟值，与 PostgreSQL 的事务时间戳并非同一次读取。`clock_ns_monotonic()` 保证单个后端内调用递增。`clock_ns_global()` 保证跨后端调用递增，需要添加以下配置并重启服务器：

```ini
shared_preload_libraries = 'pg_nanotime'
```

保留配置中已有的其他库。普通类型和时钟不要求预加载。全局递增状态位于共享内存，重启后不保留，不能当作持久序列或分布式排序机制。安装需要超级用户；扩展可迁移到其他模式，但未标记为可信扩展。
