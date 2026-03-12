

## 用法

> [hstore: 键值对数据类型](https://www.postgresql.org/docs/current/hstore.html)

`hstore` 扩展提供了一种键/值对数据类型，用于在单个列中存储半结构化数据。

```sql
CREATE EXTENSION hstore;
```

### 基本用法

```sql
SELECT 'name => Alice, age => 30'::hstore;
SELECT 'name => Alice'::hstore -> 'name';           -- Alice
SELECT 'a => 1, b => 2'::hstore ? 'a';              -- true
SELECT 'a => 1'::hstore || 'b => 2'::hstore;        -- "a"=>"1", "b"=>"2"
```

### 运算符

| 运算符 | 描述 | 示例 |
|--------|------|------|
| `->` | 按键获取值 | `h -> 'key'` |
| `\|\|` | 连接 | `h1 \|\| h2` |
| `?` | 包含键 | `h ? 'key'` |
| `?&` | 包含所有键 | `h ?& ARRAY['a','b']` |
| `?\|` | 包含任一键 | `h ?\| ARRAY['a','b']` |
| `@>` | 包含 | `h @> 'a=>1'` |
| `<@` | 被包含 | `h <@ 'a=>1, b=>2'` |
| `-` | 删除键 | `h - 'key'` 或 `h - ARRAY['a','b']` |

### 下标访问

```sql
SELECT h['name'] FROM mytable;
UPDATE mytable SET h['age'] = '31';
```

### 函数

```sql
-- 构造
SELECT hstore('key', 'value');
SELECT hstore(ARRAY['a','b'], ARRAY['1','2']);
SELECT hstore(ROW(1, 'hello'));

-- 提取
SELECT akeys(h);                    -- 键的 text 数组
SELECT avals(h);                    -- 值的 text 数组
SELECT skeys(h);                    -- 键的集合
SELECT svals(h);                    -- 值的集合
SELECT each(h);                     -- (key, value) 记录集合

-- 查询
SELECT exist(h, 'key');             -- 布尔值
SELECT defined(h, 'key');           -- 值非 NULL 时返回 true

-- 修改
SELECT delete(h, 'key');
SELECT slice(h, ARRAY['a','b']);    -- 键的子集

-- JSON 转换
SELECT hstore_to_json(h);
SELECT hstore_to_jsonb(h);
SELECT hstore_to_json_loose(h);    -- 区分数字/布尔值

-- 记录转换
SELECT populate_record(NULL::my_table, h);
```

### 索引支持

```sql
CREATE INDEX idx ON t USING gin (h);    -- 支持 @>、?、?&、?|
CREATE INDEX idx ON t USING gist (h);   -- 支持 @>、?、?&、?|
CREATE INDEX idx ON t USING btree (h);  -- 支持 =
CREATE INDEX idx ON t USING hash (h);   -- 支持 =
```
