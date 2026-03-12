


## 用法

> [index_advisor: 查询索引顾问](https://github.com/supabase/index_advisor)

`index_advisor` 分析 SQL 查询并推荐能够改善执行时间的索引，报告创建建议索引前后的成本估算以及相应的 DDL 语句。

### 函数

```sql
index_advisor(query text)
RETURNS TABLE (
    startup_cost_before jsonb,
    startup_cost_after  jsonb,
    total_cost_before   jsonb,
    total_cost_after    jsonb,
    index_statements    text[],
    errors              text[]
)
```

### 基本示例

```sql
CREATE TABLE book (
    id    int  PRIMARY KEY,
    title text NOT NULL
);

SELECT * FROM index_advisor('
    SELECT book.id FROM book WHERE title = $1
');
```

返回成本改善情况和推荐的 `CREATE INDEX` 语句。

### 多表示例

```sql
SELECT * FROM index_advisor('
    SELECT book.id, book.title, publisher.name, author.name, review.body
    FROM book
    JOIN publisher ON book.publisher_id = publisher.id
    JOIN author    ON book.author_id = author.id
    JOIN review    ON book.id = review.book_id
    WHERE author.id = $1 AND publisher.id = $2
');
```

输出包含跨多个关联表的索引推荐，以及创建前后的成本对比。

### 功能特性

- 支持使用 `$1`、`$2` 语法的参数化查询
- 支持物化视图
- 能够解析视图背后的底层表和列
- 依赖 HypoPG 进行假想索引评估
