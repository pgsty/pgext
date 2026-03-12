

## 用法

> [pg_render: PostgreSQL 的 Liquid 模板渲染](https://github.com/mkaski/pg_render)

### `render(template text, input json|array|value)`

使用 [Liquid](https://shopify.github.io/liquid/) 语法，用查询结果渲染模板：

```sql
-- 单个值
SELECT render('Total: {{ value }}', (SELECT count(*) FROM posts));

-- 单行多列
SELECT render(
    '<h1>{{ title }}</h1><p>{{ text }}</p>',
    (SELECT to_json(r) FROM (SELECT title, text FROM posts WHERE id = 1) r)
);

-- 遍历数组
SELECT render(
    '{% for v in values %} {{ v }} {% endfor %}',
    (SELECT array(SELECT title FROM posts))
);

-- 遍历多行多列
SELECT render(
    '{% for row in rows %} {{ row.title }} - {{ row.author }} {% endfor %}',
    json_agg(to_json(posts.*))
) FROM posts;
```

### `render_agg(template text, input record|json|value)`

聚合渲染函数 -- 为每一行渲染模板：

```sql
-- 从派生表逐行渲染
SELECT render_agg('{{ title }} {{ text }}', props)
FROM (SELECT title, text FROM posts) AS props;

-- 使用 json_build_object 渲染
SELECT render_agg(
    '<article><h1>{{ title }}</h1></article>',
    json_build_object('title', title)
) FROM posts;
```

### 使用存储的模板

```sql
SELECT render(
    (SELECT template FROM templates WHERE id = 'my_tpl'),
    (SELECT to_json(r) FROM (SELECT title, text FROM posts WHERE id = 1) r)
);
```

### PostgREST 集成

```sql
CREATE FUNCTION api.index() RETURNS "text/html" AS $$
SELECT render(
    '<html><body><h1>{{ title }}</h1></body></html>',
    (SELECT to_json(r) FROM (SELECT title FROM posts WHERE id = 1) r)
) $$;
```
