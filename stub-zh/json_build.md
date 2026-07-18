## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/README.md)
- [1.1.0 版本 SQL 对象](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/json_build--1.1.0.sql)
- [扩展 control 文件](https://github.com/pgexperts/json_build/blob/9a9ce7a38674499528e91dd46e280b95b94c7114/json_build.control)

`json_build` 是用于构造嵌套 JSON 的历史辅助扩展。它提供 `build_json_object(VARIADIC "any")`、`build_json_array(VARIADIC "any")` 与 `json_object_agg("any", "any")`，三者均返回 `json`。

```sql
CREATE EXTENSION json_build;
SELECT build_json_object(
  'user', build_json_object('id', 7, 'active', true),
  'roles', build_json_array('reader', 'writer')
);
```

对象键必须是非空标量，且 `build_json_object(VARIADIC "any")` 要求参数个数为偶数。数组会转成 JSON 数组，记录会转成对象，已有 JSON 值则直接传递。

现代 PostgreSQL 已内置名称和行为重叠的 JSON 构造器与聚合函数。尤其在创建本扩展时，可能与已有的 `json_object_agg("any", "any")` 签名冲突。新代码应优先使用核心 `json_build_object`、`json_build_array` 与受支持的 JSON 聚合；只有在隔离数据库验证对象名冲突与应用语义后，才应安装这个老旧扩展。
