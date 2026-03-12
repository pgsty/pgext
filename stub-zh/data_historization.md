

## 用法

> [data_historization: 在分区日志表中跟踪数据变更](https://github.com/rodo/postgresql-data-historization)

用于将数据变更历史化到分区表中的 PL/pgSQL 脚本。

### 初始化历史化

为表设置所需的对象（尚不采集数据）：

```sql
SELECT historize_table_init('public', 'my_table');
```

### 启动历史化

安装触发器，开始将变更收集到 `_log` 表中：

```sql
SELECT historize_table_start('public', 'my_table');
```

### 停止历史化

移除触发器，停止收集变更：

```sql
SELECT historize_table_stop('public', 'my_table');
```

### 重置历史化

移除 cron 条目和在源表上创建的列：

```sql
SELECT historize_table_reset('public', 'my_table');
```

### 清理历史化

完全移除日志表：

```sql
SELECT historize_table_clean('public', 'my_table');
```

### 分区管理

手动创建和删除分区：

```sql
SELECT historize_create_partition('public', 'my_table_log', 0);
SELECT historize_drop_partition('public', 'my_table_log', 0);
```

使用 `pg_cron` 自动化：

```sql
SELECT cron.schedule_in_database(
    'create-partitions', '00 08 * * *',
    $$SELECT historize_create_partition('my_table', generate_series(1, 4))$$,
    'my_database'
);
```
