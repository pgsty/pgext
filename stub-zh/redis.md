

## 用法

> [redis: 从 PostgreSQL 直接向 Redis 发送发布/订阅消息](https://github.com/brettlaforge/pg_redis_pubsub)

`redis` 扩展（pg_redis_pubsub）允许 PostgreSQL 向 Redis 发布/订阅频道发布消息。目前仅支持发布端。

### 配置

通过 GUC 变量设置 Redis 连接参数：

```sql
ALTER SYSTEM SET redis.host = '127.0.0.1';
ALTER SYSTEM SET redis.port = '6379';
SELECT pg_reload_conf();
```

也可以在数据库、角色或会话级别设置：

```sql
SET redis.host = '127.0.0.1';
SET redis.port = '6379';
```

### 基本用法

```sql
CREATE EXTENSION redis;

SELECT redis_connect();
SELECT redis_publish('mychannel', 'Hello World');
SELECT redis_disconnect();
```

### 可用函数

| 函数 | 描述 |
|----------|-------------|
| `redis_connect()` | 使用 `redis.host` 和 `redis.port` 设置连接 Redis |
| `redis_disconnect()` | 断开 Redis 连接 |
| `redis_publish(channel text, message text)` | 在指定频道发布消息 |
| `redis_status()` | 返回 Redis 客户端状态 |

注意：`redis_publish` 在没有连接时会自动连接。

### 触发器示例

当表被修改时向 Redis 发布变更事件：

```sql
CREATE OR REPLACE FUNCTION notify_changes()
RETURNS TRIGGER AS $$
BEGIN
  PERFORM redis_publish(
    'products:' || NEW.id::text,
    to_jsonb(NEW)::text
  );
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER on_product_change
  AFTER INSERT OR UPDATE ON products
  FOR EACH ROW EXECUTE PROCEDURE notify_changes();
```

这允许监听 Redis 频道的外部订阅者实时响应 PostgreSQL 数据变更。
