


## Usage

> [pg_render: Liquid template rendering for PostgreSQL](https://github.com/mkaski/pg_render)

### `render(template text, input json|array|value)`

Render a template with query results using [Liquid](https://shopify.github.io/liquid/) syntax:

```sql
-- Single value
SELECT render('Total: {{ value }}', (SELECT count(*) FROM posts));

-- Multiple columns from one row
SELECT render(
    '<h1>{{ title }}</h1><p>{{ text }}</p>',
    (SELECT to_json(r) FROM (SELECT title, text FROM posts WHERE id = 1) r)
);

-- Loop over an array
SELECT render(
    '{% for v in values %} {{ v }} {% endfor %}',
    (SELECT array(SELECT title FROM posts))
);

-- Loop over multiple rows with multiple columns
SELECT render(
    '{% for row in rows %} {{ row.title }} - {{ row.author }} {% endfor %}',
    json_agg(to_json(posts.*))
) FROM posts;
```

### `render_agg(template text, input record|json|value)`

Aggregate render function -- renders a template for each row:

```sql
-- Render each row from a derived table
SELECT render_agg('{{ title }} {{ text }}', props)
FROM (SELECT title, text FROM posts) AS props;

-- Render using json_build_object
SELECT render_agg(
    '<article><h1>{{ title }}</h1></article>',
    json_build_object('title', title)
) FROM posts;
```

### Using Stored Templates

```sql
SELECT render(
    (SELECT template FROM templates WHERE id = 'my_tpl'),
    (SELECT to_json(r) FROM (SELECT title, text FROM posts WHERE id = 1) r)
);
```

### PostgREST Integration

```sql
CREATE FUNCTION api.index() RETURNS "text/html" AS $$
SELECT render(
    '<html><body><h1>{{ title }}</h1></body></html>',
    (SELECT to_json(r) FROM (SELECT title FROM posts WHERE id = 1) r)
) $$;
```
