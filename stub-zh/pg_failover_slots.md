

## 用法

> [pg_failover_slots: PG 故障转移槽扩展](https://github.com/EnterpriseDB/pg_failover_slots)

通过将槽状态从主库同步到备库，使逻辑复制槽在物理故障转移后可用。

### 启用

在主库和备库的 `postgresql.conf` 中添加：

```ini
shared_preload_libraries = 'pg_failover_slots'
```

备库所需设置：

```ini
hot_standby_feedback = on
primary_slot_name = 'my_physical_slot'  -- 必须非空
```

### 配置选项

```ini
# 要同步的槽（默认：所有逻辑槽）
pg_failover_slots.synchronize_slot_names = 'name_like:%'

# 同步特定槽
pg_failover_slots.synchronize_slot_names = 'my_slot,plugin:test_decoding'

# 删除备库上主库中不存在的多余槽（默认：true）
pg_failover_slots.drop_extra_slots = true

# 连接到主库的连接字符串（默认：使用 primary_conninfo）
pg_failover_slots.primary_dsn = 'host=primary dbname=mydb'

# 确保物理备库在逻辑消费者之前接收数据
pg_failover_slots.standby_slot_names = 'standby_physical_slot'

# 需要确认的备库槽数（默认：-1 = 全部）
pg_failover_slots.standby_slots_min_confirmed = -1

# 同步间隔（毫秒，默认：60000）
pg_failover_slots.worker_nap_time = 60000
```

### 检查备库就绪状态

在故障转移前验证所有逻辑槽已同步：

```sql
-- 在备库上：所有槽应显示 active = false
SELECT slot_name, active FROM pg_replication_slots WHERE slot_type = 'logical';

--  slot_name        | active
-- ------------------+--------
--  regression_slot1 | f        -- 已同步，就绪
--  regression_slot2 | f        -- 已同步，就绪
--  regression_slot3 | t        -- 仍在同步，未就绪
```

当所有槽显示 `active = false` 时，备库可以安全进行故障转移。

### 关键行为

- 从主库复制缺失的复制槽到备库
- 移除备库上主库中不存在的多余槽
- 定期同步槽位置
- `standby_slot_names` 提供同步复制屏障以防止故障转移时数据丢失
- 需要 PostgreSQL 11 或更高版本
