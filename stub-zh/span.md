## 用法

来源：

- [官方 README](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/README.md)
- [官方用法文档](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/doc/span.mmd)
- [扩展 SQL](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/sql/span.sql)

`span` 增加一种紧凑数据类型，用于标识文档中的文本区间。值的形式为 `document:begin-end`，例如 `article-42:10-25`，可以存储、比较、分组并建立索引。

### 核心流程

```sql
CREATE EXTENSION span;

CREATE TABLE annotations (
    annotation_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    location span NOT NULL,
    label text NOT NULL
);

INSERT INTO annotations(location, label) VALUES
    ('article-42:10-25', 'person'),
    ('article-42:30-44', 'place');

CREATE INDEX annotations_location_idx
ON annotations(location);

SELECT get_span_doc(location),
       get_span_begin(location),
       get_span_end(location),
       label
FROM annotations
ORDER BY location;
```

默认 B-tree 操作符类支持 `=`、`<>`、`<`、`<=`、`>` 与 `>=`。hash 操作符类支持等值比较与哈希聚合。扩展还定义了 `min(span)` 与 `max(span)` 聚合。

### 构造与访问函数

- `span(text)` 把经过校验的文本转换为 `span`；`text(span)` 执行反向转换。
- `is_span(text)` 检查文本是否具有合法文档标识与有序偏移。
- `get_span_doc(span)` 返回文档标识符。
- `get_span_begin(span)` 与 `get_span_end(span)` 返回整数偏移。

```sql
SELECT is_span('article-42:10-25'),
       span('article-42:10-25'),
       text('article-42:10-25'::span);
```

### 注意事项

该类型只建模一个标识符与两个整数偏移；不提供重叠、包含、并集或范围索引操作符。排序适用于确定性排序与 B-tree 查找，并不具备区间搜索语义。非法形式、缺失端点及起点大于终点都会被拒绝。应用必须明确偏移表示字节、字符还是其他单位，并在类型之外保持一致约定。版本 `0.1.0` 较旧且采用固定长度，因此迁移前应验证文档标识符长度与 PostgreSQL 版本兼容性。不需要预加载或重启。
