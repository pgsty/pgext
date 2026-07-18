## 用法

来源：

- [扩展 control 文件](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array.control)
- [扩展 SQL](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array--1.0.sql)
- [C 实现](https://github.com/mhagander/jsonb_delete_array/blob/5880579994603f4350431690594308c3d0c1d961/jsonb_delete_array.c)

`jsonb_delete_array` 增加函数 `jsonb_delete_array(jsonb, text[])` 和运算符 `jsonb - text[]`，可以在一次调用中删除多个顶层对象键或顶层字符串数组元素：

```sql
CREATE EXTENSION jsonb_delete_array;

SELECT jsonb_delete_array(
  '{"a":1,"b":2,"c":3}'::jsonb,
  ARRAY['a', 'c']
);

SELECT '["a","b",1,true]'::jsonb - ARRAY['a', 'b']::text[];
```

### 作用范围

版本 `1.0` 仅处理顶层内容。对于对象，它删除匹配的键；对于数组，它只删除匹配的字符串元素，数字、布尔值、对象和数组元素均会保留。标量 JSONB 输入会报错，空删除列表则原样返回输入。上游仓库提供源代码但没有 README，因此这些严格限定的语义直接来自安装 SQL 与 C 实现。
