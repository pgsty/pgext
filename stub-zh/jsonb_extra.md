## 用法

来源：

- [官方扩展 SQL](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra--1.0.sql)
- [官方回归测试 SQL](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/sql/jsonb_extra.sql)
- [官方实现](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra.c)

`jsonb_extra` 1.0 是一个小型 C 扩展，用于提取键路径到达的所有 JSON 值、把 JSON 标量转换为文本，以及替换已有的嵌套值。它早于 PostgreSQL 当前的大部分 JSON path 能力，应按实验性代码评估。

### 核心流程

```sql
CREATE EXTENSION jsonb_extra;

SELECT jsonb_extract(
  '{"a":{"items":[{"x":5},{"x":6},{"other":7}]}}'::jsonb,
  ARRAY['a', 'items', 'x']
);

SELECT jsonb_extract_text(
  '{"a":{"items":[{"x":5},{"x":6}]}}'::jsonb,
  ARRAY['a', 'items', 'x']
);

SELECT jsonb_as_text('true'::jsonb);

SELECT jsonb_update(
  '{"a":{"b":{"c":1,"d":7}}}'::jsonb,
  ARRAY['a', 'b', 'c'],
  '5'::jsonb
);
```

`jsonb_extract` 沿对象键遍历，并在路径遇到数组时展开所有元素，返回 `jsonb[]`。`jsonb_extract_text` 执行相同遍历并返回 `text[]`。路径没有产生任何值时，两者都返回 null。

### 重要对象

- `jsonb_extract(jsonb, text[])` 将所有匹配值作为 JSONB 元素返回。
- `jsonb_extract_text(jsonb, text[])` 返回转换为文本的所有匹配值。
- `jsonb_as_text(jsonb)` 将顶层标量转换为文本；JSON null 变为空字符串，容器使用 JSON 文本表示。
- `jsonb_update(jsonb, text[], jsonb)` 替换已有对象路径上的值，并返回重建后的文档。

### 运维说明

四个函数都声明为 strict 和 immutable。实现明确把路径创建和数组更新标记为未完成，因此 `jsonb_update` 只应用于已有对象路径。源码中还保留了与 JSONB 值处理有关、尚未解决的内存安全问题注释。应在精确的 PostgreSQL 构建上测试畸形输入、缺失路径、嵌套数组和大文档，并在内置 JSONB 函数能够完成同类操作时优先使用 PostgreSQL 维护的内置能力。
