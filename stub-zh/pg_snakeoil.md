

## 用法

> [pg_snakeoil: PostgreSQL 数据的 ClamAV 防病毒扫描](https://github.com/credativ/pg_snakeoil)

`pg_snakeoil` 提供对 PostgreSQL 中存储数据的 ClamAV 病毒扫描功能，不干扰正常的数据库操作。

```sql
CREATE EXTENSION pg_snakeoil;
```

### 函数

| 函数 | 返回类型 | 描述 |
|----------|---------|-------------|
| `so_is_infected(text)` | `bool` | 检查文本数据是否匹配病毒特征 |
| `so_is_infected(bytea)` | `bool` | 检查 bytea 数据是否匹配病毒特征 |
| `so_virus_name(text)` | `text` | 如果感染则返回病毒名称，否则返回空字符串 |
| `so_virus_name(bytea)` | `text` | 如果感染则返回病毒名称，否则返回 NULL |
| `so_update_signatures()` | `bool` | 重新加载病毒特征库，如有变化返回 true |

### 即时扫描

```sql
SELECT so_is_infected('Not a virus!');
-- f

SELECT so_is_infected('X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*');
-- t

SELECT so_virus_name('X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*');
-- Eicar-Test-Signature
```

### 使用域进行实时防护

```sql
CREATE DOMAIN safe_text AS text CHECK (NOT so_is_infected(value));
CREATE TABLE t1 (safe safe_text);

INSERT INTO t1 VALUES ('This text is safe!');
-- INSERT

INSERT INTO t1 VALUES('X5O!P%@AP...');
-- NOTICE: Virus found: Eicar-Test-Signature
-- ERROR: value for domain safe_text violates check constraint "safe_text_check"
```

### 使用触发器进行实时防护

```sql
CREATE OR REPLACE FUNCTION check_virus() RETURNS trigger AS $$
BEGIN
    IF so_is_infected(NEW.content) THEN
        RAISE EXCEPTION 'Virus detected: %', so_virus_name(NEW.content);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER virus_check BEFORE INSERT OR UPDATE ON uploads
    FOR EACH ROW EXECUTE FUNCTION check_virus();
```
