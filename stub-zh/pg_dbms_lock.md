

## 用法

> [pg_dbms_lock: 为 PostgreSQL 添加 Oracle DBMS_LOCK 完整兼容的扩展](https://github.com/HexaCluster/pg_dbms_lock)

使用 PostgreSQL 咨询锁模拟 Oracle DBMS_LOCK 行为。

### 启用

```sql
CREATE EXTENSION pg_dbms_lock;
```

### ALLOCATE_UNIQUE

为命名锁分配唯一锁标识符：

```sql
DO $$
DECLARE
    lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique(
        lockname => 'printer_lock',
        lockhandle => lockhandle
    );
    RAISE NOTICE 'Handle: %', lockhandle;
END;
$$;
```

### REQUEST

以指定模式请求锁（Exclusive=6，Shared=4）：

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.request(
        id => 123,
        lockmode => 6,           -- 排他锁
        timeout => 300,
        release_on_commit => false
    );
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Lock request failed: %', lock_res;
    END IF;
END;
$$;
```

返回值：0=成功，1=超时，3=参数错误，4=已持有锁，5=非法句柄。

### RELEASE

显式释放之前获取的锁：

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.release(id => 123);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Release failed: %', lock_res;
    END IF;
END;
$$;
```

### SLEEP

暂停会话指定时长：

```sql
CALL dbms_lock.sleep(0.70);  -- 休眠 0.7 秒
```

### 完整示例

```sql
DO $$
DECLARE
    lock_res int;
    printer_lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique('printer_lock', printer_lockhandle);
    lock_res := dbms_lock.request(lockhandle => printer_lockhandle, lockmode => 6, timeout => 5);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'REQUEST failed: %', lock_res;
    END IF;
    -- 在此执行排他操作
    lock_res := dbms_lock.release(lockhandle => printer_lockhandle);
END;
$$;
```
