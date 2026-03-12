

## 用法

> [pgtt: 为 PostgreSQL 添加全局临时表功能的扩展](https://github.com/darold/pgtt)

### 创建全局临时表

```sql
CREATE EXTENSION pgtt;

-- ON COMMIT PRESERVE ROWS：数据在会话内跨事务保持
CREATE GLOBAL TEMPORARY TABLE test_gtt (
    id integer,
    lbl text
) ON COMMIT PRESERVE ROWS;

-- ON COMMIT DELETE ROWS：数据在事务提交时删除
CREATE GLOBAL TEMPORARY TABLE session_data (
    id integer,
    value text
) ON COMMIT DELETE ROWS;
```

`GLOBAL` 关键字也可以作为注释使用以避免警告：

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE test_gtt (
    LIKE other_table INCLUDING DEFAULTS INCLUDING CONSTRAINTS INCLUDING INDEXES
) ON COMMIT PRESERVE ROWS;
```

### CREATE AS 形式

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE gtt_copy
AS SELECT * FROM source_table WITH DATA;
```

### 使用全局临时表

```sql
INSERT INTO test_gtt VALUES (1, 'one'), (2, 'two');
SELECT * FROM test_gtt;  -- 仅在当前会话中可见
```

### 创建索引

```sql
CREATE INDEX ON test_gtt (id);
```

### 约束

支持除外键以外的所有约束：

```sql
CREATE GLOBAL TEMPORARY TABLE t2 (
    c1 serial PRIMARY KEY,
    c2 VARCHAR(50) UNIQUE NOT NULL,
    c3 boolean DEFAULT false
);
```

### 删除

```sql
DROP TABLE test_gtt;  -- 即使其他会话正在使用也可以删除
```

### 配置

```sql
SET pgtt.enabled TO off;   -- 禁用 GTT 重路由
SET pgtt.enabled TO on;    -- 重新启用 GTT 重路由
```

### 关键行为

- GTT 内容是会话局部的；其他会话无法看到您的数据
- 表结构是持久的（对所有用户可见），但数据是按会话的
- 需要通过 `session_preload_libraries = 'pgtt'` 加载
- 不支持对 GTT 进行分区
