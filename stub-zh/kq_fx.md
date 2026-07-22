## 用法

来源：

- [官方 README](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/README.md)
- [官方 Rust 实现](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/src/lib.rs)
- [官方扩展控制文件](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/kq_fx.control)

`kq_fx` 将货币汇率历史载入 PostgreSQL 共享内存，以便快速按日期查询。版本 `1.0.1` 要求预加载并重启，其固定共享缓存作用于整个集群，而非按数据库隔离。

### 核心流程

准备预期的源表，然后预加载库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'kq_fx'
```

```sql
CREATE EXTENSION kq_fx;

SELECT kq_fx_check_db();
SELECT kq_fx_populate_cache();
SELECT kq_fx_get_rate(1, 2, DATE '2026-01-31');
SELECT kq_fx_get_rate_xuid('USD', 'EUR', DATE '2026-01-31');
```

默认查询读取 `plan.currency(id,xuid)` 和 `plan.fx_rate(currency_id,to_currency_id,date,rate)`。可由超级用户设置的 GUC `kq.currency.q1_validation`、`kq.currency.q2_get_currencies_xuid` 与 `kq.currency.q3_get_currency_entries` 可以替换这些查询。

### 缓存语义与限制

`kq_fx_invalidate_cache()` 清空缓存，`kq_fx_populate_cache()` 重新加载，`kq_fx_display_cache()` 报告条目。查询也会惰性填充缓存。相同货币返回 1.0；早于第一条观测的日期返回 NULL；观测之间的日期使用不晚于该日期的最新汇率；晚于最后观测的日期使用最后一条汇率。

已核验实现将容量固定为 64 种货币、1024 个货币对，以及每个货币对 512 条带日期记录，并将货币 XUID 限制为 16 字节。默认条目查询只保留每个货币对最新的 512 行。

共享内存没有按数据库区分。首个填充缓存的会话决定其他数据库会话看到的数据，因此同一 PostgreSQL 集群中不能使用不同的汇率数据集。应限制 GUC 修改与缓存管理函数，加载前验证容量，并在源表改变后显式失效缓存。
