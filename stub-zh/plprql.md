

## 用法

> [plprql: 在 PostgreSQL 中使用 PRQL —— 管道式关系查询语言](https://github.com/kaspermarstal/plprql)

`plprql` 允许使用 [PRQL](https://prql-lang.org/)（管道式关系查询语言）编写 PostgreSQL 函数，这是一种编译为 SQL 的现代语言，采用管道语法。

```sql
CREATE EXTENSION plprql;
```

### 使用 PRQL 创建函数

```sql
CREATE FUNCTION match_stats(int)
RETURNS TABLE(player text, kd_ratio float) AS $$
  from matches
  filter match_id == $1
  group player (
    aggregate {
      total_kills = sum kills,
      total_deaths = sum deaths
    }
  )
  filter total_deaths > 0
  derive kd_ratio = total_kills / total_deaths
  select { player, kd_ratio }
$$ LANGUAGE plprql;

SELECT * FROM match_stats(42);
```

### 直接执行 PRQL

```sql
SELECT prql('from matches | filter player == "Player1"')
  AS (id int, match_id int, player text, kills int)
  LIMIT 2;
```

### 使用游标处理大结果集

```sql
SELECT prql('from matches | filter player == "Player1"', 'cursor_name');
FETCH 2 FROM cursor_name;
```

### 查看编译后的 SQL

```sql
SELECT prql_to_sql('from matches | filter player == "Player1"');
```

### PRQL 语法概览

PRQL 使用管道式转换：

```
from employees                    # 数据源
filter department == "Engineering" # 行过滤
derive monthly_salary = salary / 12 # 计算列
sort {-monthly_salary}            # 排序（- 表示降序）
select {name, monthly_salary}     # 列投影
take 10                           # 限制行数
```

### 限制

PRQL 仅支持 `SELECT` 语句。对于 `INSERT`、`UPDATE` 和 `DELETE`，请使用 SQL 或 PL/pgSQL。
