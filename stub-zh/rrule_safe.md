## 用法

来源：

- [官方 README](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/README.md)
- [扩展控制文件](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/rrule_safe.control)
- [pgrx 清单](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/Cargo.toml)

`rrule_safe` 在 PostgreSQL 内计算 iCalendar 重复规则；其 Rust 实现旨在替代已停止维护且存在内存安全问题的 `pg_rrule` 库。

### 启用

0.2.1 版本通过 pgrx 0.18.0 支持 PostgreSQL 13–18。为准确的 PostgreSQL 大版本安装文件后，以超级用户创建扩展：

```sql
CREATE EXTENSION rrule_safe;
```

控制文件将扩展标记为不可迁移、非 trusted，并加载 `rrule_safe`；它不要求预加载或重启服务器。

### 展开重复规则

`get_occurrences` 接收 RRULE 字符串、起始时间戳以及可选的包含式上界。timestamp 与 timestamp-with-time-zone 重载都会返回匹配时间点数组。

```sql
SELECT get_occurrences(
  'FREQ=WEEKLY;COUNT=4',
  '2026-01-01 09:00:00 Europe/Berlin'::timestamptz
);

SELECT get_occurrences(
  'FREQ=DAILY',
  '2026-01-01 09:00:00'::timestamp,
  '2026-01-07 09:00:00'::timestamp
);
```

为兼容旧查询，扩展把 `rrule` 定义为 `text` 上的 domain；它不会复刻 `pg_rrule` 的持久化自定义类型。已保存的旧值必须显式迁移。

### 安全边界

无界或格式异常的重复规则可能消耗大量资源。`rrule_safe.limit` 限制返回的时间点数量，默认值为 65,535。处理用户输入时应设置更低的会话值，并在应用能够确定范围时始终给出上界。

