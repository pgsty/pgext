## 用法

来源：[README](https://github.com/EnterpriseDB/pg_failover_slots/blob/master/README.md)，[releases](https://github.com/EnterpriseDB/pg_failover_slots/releases)

`pg_failover_slots` 通过把逻辑复制槽的定义和位置从 primary 同步到 standby，使这些复制槽在故障转移后仍可继续使用。

### 在两端启用

```ini
shared_preload_libraries = 'pg_failover_slots'
```

README 给出的 standby 必要设置是：

```ini
hot_standby_feedback = on
primary_slot_name = 'my_physical_slot'
```

### 主要配置

```ini
pg_failover_slots.synchronize_slot_names = 'name_like:%'
pg_failover_slots.drop_extra_slots = true
pg_failover_slots.primary_dsn = 'host=primary dbname=mydb'
pg_failover_slots.standby_slot_names = 'standby_physical_slot'
pg_failover_slots.standby_slots_min_confirmed = -1
pg_failover_slots.worker_nap_time = 60000
pg_failover_slots.maintenance_db = 'postgres'
```

README 说明 `synchronize_slot_names` 支持按精确槽名、`LIKE` 模式或 plugin 名称过滤。

### 故障转移前检查 standby 是否就绪

```sql
SELECT slot_name, active
FROM pg_replication_slots
WHERE slot_type = 'logical';
```

在 standby 上，逻辑槽只有“存在且 `active = false`”时才算准备完成。README 明确说明 `active = true` 表示该槽仍在初始化过程中。

### 说明

- 上游要求 PostgreSQL 11+。
- `v1.2.1` 是 bug-fix 发布，没有新增面向用户的 SQL 或 GUC。
- `v1.2.0` 增加了 PostgreSQL 18 支持，并补充说明了 `drop_extra_slots` 的行为。
