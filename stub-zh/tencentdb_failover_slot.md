## 用法

来源：

- [TencentDB Failover Slot 官方文档](https://cloud.tencent.com/document/product/409/86587)
- [TencentDB for PostgreSQL 官方产品页](https://cloud.tencent.com/product/postgres)

`tencentdb_failover_slot` 是 TencentDB for PostgreSQL 的服务商扩展，用于把逻辑复制槽状态同步到备库，使逻辑复制在 HA 切换后能够继续。它属于 TencentDB 内核能力，并非可移植的社区扩展，而且只支持逻辑复制槽。

### 创建故障转移槽

实例参数 `tencentdb_force_enable_failover_slot` 启用时，TencentDB 默认把新建逻辑槽创建为故障转移槽：

```sql
CREATE EXTENSION tencentdb_failover_slot;

SELECT *
FROM pg_create_logical_replication_slot(
  'app_slot',
  'pgoutput'
);

SELECT * FROM pg_failover_slots;
SELECT *
FROM pg_replication_slots
WHERE slot_name = 'app_slot';
```

如果关闭了默认强制参数，可用 `pg_create_logical_failover_slot(slotname, pluginname)` 显式创建。腾讯云建议使用 `pgoutput` 解码插件。

### 生命周期操作

- `pg_create_logical_failover_slot()` 显式创建故障转移逻辑槽。
- `pg_failover_slots` 列出当前故障转移槽名称；普通槽详情可连接或查询 `pg_replication_slots`。
- `transform_slot_to_nonfailover()` 把未激活的故障转移槽转换成普通槽。
- `pg_drop_replication_slot()` 删除任意一种复制槽。

复制槽属于实例级，而扩展属于数据库级。切换到另一个数据库调用函数前，需要在那里重新创建 `tencentdb_failover_slot`。

### 故障转移策略与注意事项

TencentDB 参数 `failover_slot_timeline_diverged_option` 控制 HA 期间备库 WAL 接收落后于逻辑消费的情况。默认值 `error` 会暂停逻辑复制并向操作方报告错误。`rewind` 会从切换时间点恢复，因此可能涉及连续性/数据丢失取舍；只有具备应用级对账方案时才能选择。

故障转移前后都应监控保留 WAL、复制槽活跃状态、消费者 LSN、备库重放和磁盘用量。应使用真实订阅端测试计划内与意外切换，确认重复/缺口处理，并及时删除遗弃槽。该功能不支持物理复制槽，其行为、可用性、权限、备份和升级仍受 TencentDB 服务及区域约束。
