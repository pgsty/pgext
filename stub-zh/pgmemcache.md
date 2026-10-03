## 用法

来源：

- [2.3.0 README](https://github.com/ohmu/pgmemcache/blob/2.3.0/README.rst)
- [2.3.0 release notes](https://github.com/ohmu/pgmemcache/blob/2.3.0/NEWS)
- [SQL API](https://github.com/ohmu/pgmemcache/blob/2.3.0/ext/pgmemcache.sql)

`pgmemcache` 提供与 memcached 服务器交互的 PostgreSQL 用户自定义函数。

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
SELECT memcache_server_add('cache-host');  -- uses default port 11211
```

### 设置和获取值

```sql
-- Set a key (overwrites if exists)
SELECT memcache_set('user:1:name', 'John Doe');
SELECT memcache_set('session:abc', 'data', now() + interval '1 hour');

-- Add a key (fails if exists)
SELECT memcache_add('user:2:name', 'Jane Doe');
SELECT memcache_add('temp_key', 'value', interval '5 minutes');

-- Replace (fails if key doesn't exist)
SELECT memcache_replace('user:1:name', 'John Smith');

-- Get a value
SELECT memcache_get('user:1:name');  -- returns text or NULL

-- Get multiple values
SELECT key, value FROM memcache_get_multi('{key1,key2,key3}'::text[]);
```

### 原子计数器

```sql
SELECT memcache_incr('counter');        -- increment by 1
SELECT memcache_incr('counter', 5);     -- increment by 5
SELECT memcache_decr('counter');        -- decrement by 1
SELECT memcache_decr('counter', 3);     -- decrement by 3
```

### 删除和刷新

```sql
SELECT memcache_delete('user:1:name');
SELECT memcache_flush_all();  -- flush all servers
```

### 统计信息

```sql
SELECT memcache_stats();  -- returns stats from all servers
```

### 触发器示例

表更新时失效缓存：

```sql
CREATE OR REPLACE FUNCTION auth_passwd_upd()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    IF OLD.passwd IS DISTINCT FROM NEW.passwd THEN
        PERFORM memcache_delete('user_id_' || NEW.user_id || '_password');
    END IF;
    RETURN NEW;
END;
$$;

CREATE TRIGGER auth_passwd_upd_trg AFTER UPDATE ON passwd
    FOR EACH ROW EXECUTE PROCEDURE auth_passwd_upd();
```

### 配置与失败边界

库需要可访问的 memcached 服务，以及 libmemcached 或 OMcache。上述预加载配置需要重启 PostgreSQL。通过 `memcache_server_add` 添加的服务器列表属于当前后端；不使用默认配置时，应为每个新连接安排初始化。

`pgmemcache.flush_on_commit` 可以在提交时发送缓冲请求。启用请求缓冲后，设置类调用可能返回 NULL，因为最终结果尚未确定；这不同于 `memcache_get` 在键不存在时返回 NULL。缓存操作是外部副作用，数据库回滚不会撤销已经发送的修改。不要把 memcached 当作持久数据的权威副本。

`memcache_flush_all()` 删除所有已配置缓存服务器上的数据。应限制函数授权和网络访问，尤其是在缓存键或值包含凭证时。
