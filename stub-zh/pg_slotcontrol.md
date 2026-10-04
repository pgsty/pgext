## 用法

来源：

- [pg_slotcontrol/pg_slotcontrol.control](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol.control)
- [pg_slotcontrol/README.md](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/README.md)
- [pg_slotcontrol/pg_slotcontrol--1.0.sql](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol--1.0.sql)
- [pg_slotcontrol/pg_slotcontrol.c](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol.c)
- [LICENSE](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/LICENSE)

`pg_slotcontrol` 通过 `pg_slotmove(text, pg_lsn)` 向前推进物理复制槽的保留位置。这是 2017 年源码快照，尚未确认兼容当前 PostgreSQL。

### 核心用法

```sql
CREATE EXTENSION pg_slotcontrol;
SELECT slot_name, slot_type, restart_lsn FROM pg_replication_slots;
```

```sql
SELECT pg_slotmove('physical_standby', '1/86000100'::pg_lsn);
```

### 运行边界

只有确认其他位置保存着所需 WAL 后才能推进复制槽。C 实现会拒绝逻辑槽，将目标限制在当前 WAL 写入位置以内，并且不会向后移动。返回的布尔值表示位置是否改变。安装需要超级用户；PUBLIC 执行权限已撤销，应谨慎委派该函数。上游未要求预加载。创建此扩展不会安装独立的 `pg_slotsync` 程序。
