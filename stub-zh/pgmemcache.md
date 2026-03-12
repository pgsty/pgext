

## 用法

> [pgmemcache: memcached 接口](https://github.com/ohmu/pgmemcache)

提供与 memcached 服务器交互的 PostgreSQL 用户自定义函数。

### 启用

```sql
CREATE EXTENSION pgmemcache;
```

在 `postgresql.conf` 中配置默认服务器：

```ini
shared_preload_libraries = 'pgmemcache'
pgmemcache.default_servers = 'localhost:11211'
pgmemcache.default_behavior = 'DEAD_TIMEOUT:2'
```

### 服务器管理

```sql
SELECT memcache_server_add('localhost:11211');
SELECT memcache_server_add('cache-host');  -- 使用默认端口 11211
```

### 设置和获取值

```sql
-- 设置键（存在则覆盖）
SELECT memcache_set('user:1:name', 'John Doe');
SELECT memcache_set('session:abc', 'data', now() + interval '1 hour');

-- 添加键（存在则失败）
SELECT memcache_add('user:2:name', 'Jane Doe');
SELECT memcache_add('temp_key', 'value', interval '5 minutes');

-- 替换（键不存在则失败）
SELECT memcache_replace('user:1:name', 'John Smith');

-- 获取值
SELECT memcache_get('user:1:name');  -- 返回 text 或 NULL

-- 获取多个值
SELECT key, value FROM memcache_get_multi('{key1,key2,key3}'::text[]);
```

### 原子计数器

```sql
SELECT memcache_incr('counter');        -- 增加 1
SELECT memcache_incr('counter', 5);     -- 增加 5
SELECT memcache_decr('counter');        -- 减少 1
SELECT memcache_decr('counter', 3);     -- 减少 3
```

### 删除和刷新

```sql
SELECT memcache_delete('user:1:name');
SELECT memcache_flush_all();  -- 刷新所有服务器
```

### 统计信息

```sql
SELECT memcache_stats();  -- 返回所有服务器的统计信息
```

### 触发器示例

表更新时失效缓存：

```sql
CREATE OR REPLACE FUNCTION auth_passwd_upd()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    IF OLD.passwd <> NEW.passwd THEN
        PERFORM memcache_delete('user_id_' || NEW.user_id || '_password');
    END IF;
    RETURN NEW;
END;
$$;

CREATE TRIGGER auth_passwd_upd_trg AFTER UPDATE ON passwd
    FOR EACH ROW EXECUTE PROCEDURE auth_passwd_upd();
```
