## 用法

来源：

- [Official PGXN documentation](https://pgxn.org/dist/pair/0.1.8/README.html)
- [Official extension SQL](https://api.pgxn.org/src/pair/pair-0.1.8/sql/pair.sql)

`pair` 增加一个小型有序键值复合类型。两个字段都是 `text`；其主要用途是向函数传递有序、可变数量的键值对，而不必采用 `hstore` 或 JSON 之类更大的容器。

### 核心流程

用 `pair()` 或 `~>` 运算符构造值，再访问其 `k` 和 `v` 字段：

```sql
CREATE EXTENSION pair;

SELECT pair('environment', 'production');
SELECT 'retries' ~> 3;
SELECT ('region' ~> 'us-east-1').k,
       ('region' ~> 'us-east-1').v;
```

重载接受 `text` 和多态 `anyelement` 参数，但生成的复合值仍会把每个值存为 `text`。

### 可变参数用法

该类型尤其适合作为可变参数函数的参数：

```sql
CREATE FUNCTION show_pairs(VARIADIC items pair[])
RETURNS TABLE (key text, value text)
LANGUAGE sql
AS $$
  SELECT (item).k, (item).v
  FROM unnest(items) AS item
$$;

SELECT * FROM show_pairs(
  'environment' ~> 'production',
  'retries' ~> 3
);
```

### 重要对象与说明

- 类型：`pair`，字段为 `k text` 和 `v text`。
- 构造函数：`pair(anyelement, text)`、`pair(text, anyelement)`、`pair(anyelement, anyelement)`、`pair(text, text)`。
- 运算符：调用上述构造函数的对应 `~>` 重载。

`pair` 除 PostgreSQL 外没有运行时依赖，可以迁移，无需预加载或重启。PGXN 0.1.8 分发包声明的 SQL 扩展版本仍是 0.1.2，与目录版本一致；这次分发更新并不是扩展模式版本升级。
