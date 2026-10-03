## 用法

来源：

- [extensions/pg_calendar/pg_calendar.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/pg_calendar.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_calendar/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/Cargo.toml)
- [extensions/pg_calendar/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_calendar/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/seed.rs)
- [extensions/pg_calendar/src/compute.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/compute.rs)
- [extensions/pg_calendar/src/pattern.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/pattern.rs)

`pg_calendar` 0.3.0 在 `pgcalendar` 模式中管理命名工作日历、每周规则、节假日与指定日期例外。

### 核心用法

```sql
CREATE EXTENSION pg_calendar;
SELECT pgcalendar.seed_iso();
SELECT pgcalendar.is_working_day('iso', DATE '2026-10-02');
SELECT pgcalendar.add_working_days('iso', DATE '2026-10-02', 3);
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。`create_calendar`、`set_working_days`、`add_holiday` 与 `add_exception` 用于定义日历；`is_working_day`、`add_working_days` 和区间计数函数用于查询。日期例外优先于节假日和每周规则。种子函数提供工作周约定，不包含全球权威节假日数据。`pg_calendar.enabled` 由超级用户控制。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
