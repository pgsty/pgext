


## 用法

> [pg_logicalinspect: 检查逻辑解码快照文件](https://www.postgresql.org/docs/current/pglogicalinspect.html)

pg_logicalinspect 提供 SQL 函数来检查存储在 `pg_logical/snapshots/` 中的序列化逻辑解码快照，适用于调试和教学目的。

### 函数

**`pg_get_logical_snapshot_meta(filename text)`** -- 快照文件元数据：

```sql
SELECT * FROM pg_get_logical_snapshot_meta('0-40796E18.snap');

 magic      | 1369563137
 checksum   | 1028045905
 version    | 6
```

**`pg_get_logical_snapshot_info(filename text)`** -- 详细快照信息：

```sql
SELECT * FROM pg_get_logical_snapshot_info('0-40796E18.snap');

 state                    | consistent
 xmin                     | 751
 xmax                     | 751
 start_decoding_at        | 0/40796AF8
 two_phase_at             | 0/40796AF8
 initial_xmin_horizon     | 0
 building_full_snapshot   | f
 in_slot_creation         | f
 last_serialized_snapshot | 0/0
 committed_count          | 0
 committed_xip            |
 catchange_count          | 2
 catchange_xip            | {751,752}
```

### 列出可用快照

结合 `pg_ls_logicalsnapdir()` 来发现和检查所有快照：

```sql
SELECT ss.name, meta.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_meta(ss.name) AS meta;

SELECT ss.name, info.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_info(ss.name) AS info;
```

### 访问控制

默认限制为超级用户和 `pg_read_server_files` 成员。
