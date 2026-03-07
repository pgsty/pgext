
## 用法

请注意，`cron.database` 必须在将 `pg_cron` 添加到 `shared_preload_libraries` 之前设置好。

```
-- 每周六凌晨 3:30（GMT）删除过期数据
SELECT cron.schedule('30 3 * * 6', $$DELETE FROM events WHERE event_time < now() - interval '1 week'$$);
 schedule
----------
       42

-- 每天上午 10:00（GMT）执行 VACUUM
SELECT cron.schedule('nightly-vacuum', '0 10 * * *', 'VACUUM');
 schedule
----------
       43

-- 改为每天凌晨 3:00（GMT）执行 VACUUM
SELECT cron.schedule('nightly-vacuum', '0 3 * * *', 'VACUUM');
 schedule
----------
       43

-- 取消定时任务
SELECT cron.unschedule('nightly-vacuum' );
 unschedule
------------
 t

SELECT cron.unschedule(42);
 unschedule
------------
          t

-- 每周日凌晨 4:00（GMT）在 pg_cron 所在数据库之外的其他数据库中执行 VACUUM
SELECT cron.schedule_in_database('weekly-vacuum', '0 4 * * 0', 'VACUUM', 'some_other_database');
 schedule
----------
       44

-- 每 5 秒调用一次存储过程
SELECT cron.schedule('process-updates', '5 seconds', 'CALL process_updates()');

-- 每月最后一天中午 12:00 执行工资处理
SELECT cron.schedule('process-payroll', '0 12 $ * *', 'CALL process_payroll()');
```

Crontab 格式说明：

```
 ┌───────────── 分钟 (0 - 59)
 │ ┌────────────── 小时 (0 - 23)
 │ │ ┌─────────────── 日期 (1 - 31) 或月末最后一天 ($)
 │ │ │ ┌──────────────── 月份 (1 - 12)
 │ │ │ │ ┌───────────────── 星期 (0 - 6)（0 到 6 表示周日到
 │ │ │ │ │                  周六，也可使用英文名称；7 同样表示周日）
 │ │ │ │ │
 │ │ │ │ │
 * * * * *
```
