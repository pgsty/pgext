## 用法

来源：

- [官方 README 与路径模型](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/README.md)
- [扩展 SQL API](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/jsonknife--1.0.sql)
- [官方回归测试](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/sql/test.sql)

`jsonknife` 1.0 提供 C 函数，用声明式路径从 `jsonb` 中提取值并进行归约。路径元素为字符串时选择对象键，为数字时选择数组下标，为对象时按匹配模式过滤数组元素。它适合紧凑的只读遍历，但不提供索引访问方法或 JSON 修改 API。

### 核心流程

第二个参数是由路径组成的 JSON 数组，而每条路径本身也是 JSON 数组：

```sql
CREATE EXTENSION jsonknife;

SELECT knife_extract_text(
  '{"friends":[
     {"best":true,"name":"Niquola"},
     {"best":false,"name":"Avraam"},
     {"best":true,"name":"Ivan"}
   ]}'::jsonb,
  '[["friends", {"best":true}, "name"]]'::jsonb
);
```

结果是包含匹配姓名的 PostgreSQL 数组。外层数组可提供多条路径，遍历过程也会展开 JSON 数组中找到的匹配项。

### 函数索引

- `knife_extract(jsonb, jsonb) RETURNS jsonb[]` 保留 JSON 值。
- `knife_extract_text`、`knife_extract_numeric` 和 `knife_extract_timestamptz` 返回带类型的 PostgreSQL 数组。
- `knife_extract_min_numeric`、`knife_extract_max_numeric`、`knife_extract_min_timestamptz` 和 `knife_extract_max_timestamptz` 把匹配值归约为单个边界。
- `knife_date_bound(text, text) RETURNS timestamptz` 使用 `min` 或 `max` 把部分日期解释为缺失字段的下界或上界。

所有函数都声明为 `STRICT` 和 `IMMUTABLE`；NULL 输入会产生 NULL，类型转换错误会中止语句。实现直接使用 PostgreSQL 内部 `jsonb` C 结构，并源自 `jsonb` 的早期版本。部署前应确认目标 PostgreSQL 大版本的源码/二进制兼容性；处理大型文档时还应基准测试，因为提取过程会遍历传入 JSON，而不会使用数据库索引。
