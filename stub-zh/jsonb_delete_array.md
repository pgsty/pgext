## 用法

来源：

- [官方仓库](https://github.com/mhagander/jsonb_delete_array)
- [扩展控制文件](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array.control)
- [1.0 版扩展 SQL](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array--1.0.sql)
- [1.0 版 C 实现](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array.c)
- [当前 PostgreSQL JSON 运算符](https://www.postgresql.org/docs/current/functions-json.html)

jsonb_delete_array 增加一个函数和减法运算符，可在一次调用中删除 JSONB 值的多个顶层对象键，或顶层数组中匹配的字符串元素。它不会递归进入嵌套对象或数组。

### 核心流程

传入原始文档，以及由待删除名称或字符串值组成的一维文本数组。

```sql
CREATE EXTENSION jsonb_delete_array;

SELECT jsonb_delete_array(
    '{"keep":1,"drop_a":2,"drop_b":3}'::jsonb,
    ARRAY['drop_a', 'drop_b']
);

SELECT '["keep","drop","other","drop"]'::jsonb
       - ARRAY['drop']::text[];
```

对于对象，匹配的顶层键及其值会被删除。对于数组，匹配的字符串元素会被删除，非字符串元素则保留。删除数组中的空值会被忽略，空删除数组不会改变输入。

### 边界与错误

函数拒绝标量 JSONB 输入和多维删除数组。匹配要求完全一致并区分大小写。嵌套值会整体复制且不会遍历，因此删除某个名称不会删除顶层以下的同名项。

### PostgreSQL 版本兼容性

当前 PostgreSQL 已记录一个接收文本数组的内置减法重载。由于该扩展会创建相同签名的运算符，在核心已提供此功能的 PostgreSQL 版本上安装可能发生冲突。应检查目标服务器的运算符目录，并在内置行为可用时优先使用受支持的核心实现。

该扩展可重定位，也没有声明预加载或重启要求。移除或替换扩展前应审查应用查询，因为其他 PostgreSQL 版本提供的同形运算符，仍需针对应用预期的对象、数组、标量和嵌套行为逐项测试。
