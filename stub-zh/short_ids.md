## 用法

来源：

- [官方 README](https://gitlab.com/depesz/short_ids/-/blob/26d4f30f1b0eef49b1ac81cf9ffae295d4cd33f2/README)
- [官方扩展 SQL](https://gitlab.com/depesz/short_ids/-/blob/26d4f30f1b0eef49b1ac81cf9ffae295d4cd33f2/short_ids--0.1.0.sql)
- [官方 PGXN 版本](https://pgxn.org/dist/short_ids/0.1.0/)

`short_ids` 生成适合作为列默认值的短随机文本标识符。它会在目标表中检查冲突，并通过事务级咨询锁防止并发调用同一函数时返回相同候选值；但唯一约束或主键仍必须作为权威保护。

### 核心流程

```sql
CREATE EXTENSION short_ids;

CREATE TABLE public.links (
  id text PRIMARY KEY DEFAULT generate_random_id(
    'public', 'links', 'id', 4
  ),
  target_url text NOT NULL
);

INSERT INTO public.links(target_url)
VALUES ('https://example.com/docs')
RETURNING id;
```

如需自定义字符集，可传入可选的第五个参数：

```sql
SELECT generate_random_id(
  'public', 'links', 'id', 6,
  'abcdefghijklmnopqrstuvwxyz'
);
```

发生冲突或咨询锁竞争时，`generate_random_id()` 会把候选长度增加一位并重试。

### 函数索引

- `generate_random_id(table_schema text, table_name text, column_name text, string_length integer, possible_chars text DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz')` 返回当前未使用的候选值。
- `get_random_string(string_length integer, possible_chars text DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz')` 只生成字符串，不检查表。

### 注意事项

版本 `0.1.0` 使用 PostgreSQL `random()`，不是密码学安全生成器。标识符可被猜测，不能用作密码、持有者令牌或访问控制秘密。安全性依赖所有写入者都遵守数据库唯一约束；直接并发插入可能与预检查竞争，调用方必须在唯一冲突后重试。应验证长度为正且字符集非空，扩展若未安装在 `public` 中则需限定函数模式，监控冲突导致的长度增长，并注意该项目年代较早且目录状态为已废弃。
