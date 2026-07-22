## 用法

来源：

- [官方 README](https://github.com/ketteq-neon/kq_cx/blob/main/README.md)
- [扩展 control 文件](https://github.com/ketteq-neon/kq_cx/blob/main/kq_cx.control)
- [SQL API 与共享内存实现](https://github.com/ketteq-neon/kq_cx/blob/main/src/lib.rs)
- [日历运算实现](https://github.com/ketteq-neon/kq_cx/blob/main/src/math.rs)
- [支持的构建特性](https://github.com/ketteq-neon/kq_cx/blob/main/Cargo.toml)

`kq_cx` 把工作日历日期缓存在 PostgreSQL 共享内存中，并提供快速工作日运算。它面向 KetteQ 模式设计，但可通过四个 SQL 查询 GUC 适配其他兼容的日历模型。使用它必须预加载、重启服务器、以超级用户安装，并谨慎确定数据库范围内共享缓存的唯一权威来源。

### 核心流程

把库加入 `shared_preload_libraries`，重启 PostgreSQL，创建扩展，然后加载或检查缓存：

```conf
shared_preload_libraries = 'kq_cx'
```

```sql
CREATE EXTENSION kq_cx;

SELECT kq_cx_populate_cache();
SELECT * FROM kq_cx_cache_info();
SELECT kq_cx_add_days(DATE '2026-07-22', 5, 42);
SELECT kq_cx_add_days_xuid(DATE '2026-07-22', -3, 'US-WORKDAYS');
```

加载器默认需要 `plan.calendar`、`plan.calendar_date` 和 `plan.data_date`。在填充缓存前，只能通过 `kq.calendar.q_schema_validation`、`kq.calendar.q1_get_calendar_min_max_id`、`kq.calendar.q2_get_calendars_entry_count` 和 `kq.calendar.q3_get_calendar_entries` 定制权威 SQL。

### 主要对象

- `kq_cx_add_days(date, integer, bigint)` 使用数字日历标识符。
- `kq_cx_add_days_xuid(date, integer, text)` 使用外部标识符。
- `kq_cx_populate_cache()` 加载日历数据；`kq_cx_invalidate_cache()` 强制后续重新加载。
- `kq_cx_cache_info()`、`kq_cx_info()`、`kq_cx_display_cache()` 和 `kq_cx_display_page_map()` 暴露诊断状态。

以上当前 SQL API 名称取自实现；README 的一些示例仍使用旧名称。未知日历会返回 `NULL` 并产生警告。运算结果超出缓存范围时会返回哨兵日期 `1970-01-01` 或 `2199-01-01`，而不是 `NULL` 或错误。

### 限制与注意事项

固定的共享内存布局最多支持 64 个日历、每个日历 8192 个条目、512 个页映射条目，以及最长 32 字节的外部日历标识符。加载前应验证源数据。

缓存及其已填充标志属于整个 postmaster，而加载查询通过某个数据库连接执行。在多数据库集群中，应指定一个兼容数据库作为缓存权威来源，并协调失效操作。日历运算函数被声明为 immutable，但管理员可以重新加载共享状态；因此，不要在依赖日历永久不变的表达式索引或生成列中使用这些函数。
