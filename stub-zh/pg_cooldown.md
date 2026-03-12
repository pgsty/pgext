

## 用法

> [pg_cooldown: 从共享缓冲区中移除特定关系的缓存页](https://github.com/rbergm/pg_cooldown)

pg_cooldown 是 `pg_prewarm` 的互补工具：它从共享缓冲区中移除特定关系的所有缓存页，适用于研究和测试中模拟冷启动场景。

### 从共享缓冲区中移除页面

```sql
-- 移除表的所有数据页
SELECT pg_cooldown('my_relation');

-- 移除主键索引的页面
SELECT pg_cooldown('my_relation_pkey');

-- 移除任意索引的页面
SELECT pg_cooldown('my_index_name');
```
