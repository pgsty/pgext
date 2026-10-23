## 用法

来源：

- [README.md](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/README.md)
- [ora_timestamptz--1.0.sql](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_timestamptz--1.0.sql)
- [ora_tstz_ddl.c](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_tstz_ddl.c)
- [ora_timestamptz.control](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_timestamptz.control)

`ora_timestamptz` 同时保存 UTC 时刻及原始时区区域名或固定偏移量。当前源码面向 PostgreSQL 17，其他主版本尚未经上游验证。该扩展不可迁移模式，并会在 `pg_catalog` 中创建函数，因此必须由超级用户安装。普通类型操作无须预加载或重启。

### 核心用法

```sql
CREATE EXTENSION ora_timestamptz;
CREATE TABLE events (id bigint PRIMARY KEY, occurred_at ora_timestamptz(9));
INSERT INTO events VALUES
  (1, '2026-10-16 14:30:00 Europe/Berlin'),
  (2, '2026-10-16 08:30:00 -04:00');
SELECT id, ora_timestamptz_iso(occurred_at) FROM events;
SELECT count(DISTINCT occurred_at) FROM events;
```

### 运行边界

小数秒精度范围为 0 到 9，默认 6。相等、排序与唯一性按 UTC 时刻比较，因此不同区域的值也可能相等。`ora_timestamptz_identical` 还会比较时区，`ora_timestamptz_iso` 输出无损文本。`from_tz`、`sys_extract_utc`、`to_timestamp_tz` 和 `to_char` 用于转换与格式化，也提供 B-tree、哈希和 BRIN 操作符类。

夏令时跳跃造成的不存在时间会报错。重叠时间默认选择标准时间，也可通过 `ora_timestamptz.error_on_overlap_time` 拒绝歧义值。时间运算基于 UTC，EXTRACT 返回 UTC 日期与时间字段。`ora_timestamptz.output_format` 默认为 iso；nls 输出可能丢失信息，不适合备份。

可选的 DDL 重写由 `ora_timestamptz.replace_timestamptz` 控制，并可通过 `ora_timestamptz.replace_schemas` 限定模式。需要为新会话设置 `session_preload_libraries`，或配置 `shared_preload_libraries` 并重启服务器。恢复原生时间戳列时应关闭重写，否则备份中的类型语法可能被改写。重写不会转换现有列，也不影响查询字面量或类型转换。访问已有数据前，备库与恢复目标都必须安装扩展库。
