

## Usage

> [pgmemcache: memcached interface](https://github.com/ohmu/pgmemcache)

Provides PostgreSQL user-defined functions for interacting with memcached servers.

### Enabling

```sql
CREATE EXTENSION pgmemcache;
```

Configure default servers in `postgresql.conf`:

```ini
shared_preload_libraries = 'pgmemcache'
pgmemcache.default_servers = 'localhost:11211'
pgmemcache.default_behavior = 'DEAD_TIMEOUT:2'
```

### Server Management

```sql
SELECT memcache_server_add('localhost:11211');
SELECT memcache_server_add('cache-host');  -- uses default port 11211
```

### Setting and Getting Values

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

### Atomic Counters

```sql
SELECT memcache_incr('counter');        -- increment by 1
SELECT memcache_incr('counter', 5);     -- increment by 5
SELECT memcache_decr('counter');        -- decrement by 1
SELECT memcache_decr('counter', 3);     -- decrement by 3
```

### Deleting and Flushing

```sql
SELECT memcache_delete('user:1:name');
SELECT memcache_flush_all();  -- flush all servers
```

### Statistics

```sql
SELECT memcache_stats();  -- returns stats from all servers
```

### Trigger Example

Invalidate cache on table updates:

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
