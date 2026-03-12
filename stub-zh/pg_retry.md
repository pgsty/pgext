

## 用法

> [pg_retry: 在瞬态错误时使用指数退避重试 SQL 语句](https://github.com/Agent-Hellboy/pg_retry)

### 函数签名

```sql
retry.retry(
  sql TEXT,                          -- 要执行的 SQL 语句（仅限一条）
  max_tries INT DEFAULT 3,           -- 总尝试次数（1 + 重试次数），>= 1
  base_delay_ms INT DEFAULT 50,      -- 初始退避延迟（毫秒）
  max_delay_ms INT DEFAULT 1000,     -- 指数退避延迟上限
  retry_sqlstates TEXT[] DEFAULT ARRAY['40001','40P01','55P03','57014']
) RETURNS INT                       -- 处理的行数
```

默认可重试的 SQLSTATE：`40001`（序列化失败）、`40P01`（检测到死锁）、`55P03`（锁不可用）、`57014`（查询被取消）。

### 示例

使用默认参数重试：

```sql
SELECT retry.retry('UPDATE accounts SET balance = balance - 100 WHERE id = 1');
```

自定义重试参数：

```sql
SELECT retry.retry(
    'INSERT INTO audit_log (event) VALUES (''test'')',
    5,      -- max_tries
    100,    -- base_delay_ms
    5000    -- max_delay_ms
);
```

### GUC 配置

```sql
ALTER SYSTEM SET pg_retry.default_max_tries = 5;
ALTER SYSTEM SET pg_retry.default_base_delay_ms = 100;
ALTER SYSTEM SET pg_retry.default_max_delay_ms = 5000;
ALTER SYSTEM SET pg_retry.default_sqlstates = '40001,40P01,55P03,57014';
SELECT pg_reload_conf();
```

### 安全规则

- 每次调用只能执行一条 SQL 语句（多条语句会失败）
- 禁止使用事务控制语句（BEGIN、COMMIT、ROLLBACK）
- 参数会被校验（max_tries >= 1，延迟不能为负值，base <= max delay）
