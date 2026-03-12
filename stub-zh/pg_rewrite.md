

## 用法

> [pg_rewrite: 允许在表重写期间继续读写的工具](https://github.com/cybertec-postgresql/pg_rewrite)

`pg_rewrite` 需要设置 `wal_level = logical`，并且必须添加到 `shared_preload_libraries` 中。常见用例包括更改列数据类型、分区表、重新排列列以及将表移动到不同的表空间——同时允许并发读写。

### 重写表

创建具有目标架构的目标表，然后调用 `rewrite_table()`：

```sql
-- 源表
CREATE TABLE measurement (id int, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate));

-- 具有新模式的目标表（例如 bigint id + 分区）
CREATE TABLE measurement_aux (id bigint, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate))
  PARTITION BY RANGE (logdate);

CREATE TABLE measurement_y2006m02 PARTITION OF measurement_aux FOR VALUES FROM ('2006-02-01') TO ('2006-03-01');

-- 执行重写（复制数据、应用并发变更，然后重命名）
SELECT rewrite_table('measurement', 'measurement_aux', 'measurement_old');
```

源表和目标表都必须有一个标识索引（通常是主键）。该函数会复制所有行，通过逻辑解码重放并发变更，然后原子性地重命名表。

### 进度监控

```sql
SELECT * FROM pg_rewrite_progress;
```

显示 `ins_initial`（初始复制的行数）、`ins`、`upd`、`del`（应用的并发变更）。

### 配置

- **`rewrite.max_xlock_time`** -- 最终重命名阶段持有排他锁的最大时间（毫秒）。默认 `0`（不限制）。设置为例如 `100` 可将锁定时间限制在 0.1 秒；超出时函数会重试。

```sql
SET rewrite.max_xlock_time TO 100;
```

### 约束处理

- PRIMARY KEY、UNIQUE、EXCLUDE：在调用 `rewrite_table()` 之前添加到目标表
- CHECK、NOT NULL（PG18+）、FOREIGN KEY：自动创建为 NOT VALID；使用 `ALTER TABLE ... VALIDATE CONSTRAINT ...` 验证
