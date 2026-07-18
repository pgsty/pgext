## 用法

来源：

- [上游固定版本 README](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/README.md)
- [0.1 版本安装 SQL](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/jsonfun--0.1.sql)
- [固定版本 JSON 解析实现](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/json_extract_keys_array.c)
- [固定版本扩展控制文件](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/jsonfun.control)

jsonfun 0.1 提供两个 C 函数，把一个 json 值中的多个对象路径提取到同一个文本数组。默认函数使用点号分隔路径，另一个函数接收自定义分隔符。如果选中的值是 JSON 数组，其直接元素会依次追加到结果中。

### 提取示例

```sql
CREATE EXTENSION jsonfun;

SELECT json_extract_keys_array(
    '{"a":{"b":"hello"},"c":{"d":[1,2,null]}}'::json,
    'a.b',
    'c.d',
    'missing.path'
);
```

结果是包含 hello、1 与 2 的文本数组。不存在的路径和 JSON null 会被忽略。路径只能遍历对象：不支持在路径中使用数组下标，尽管终点数组会被展开。

### 兼容性与输入限制

实现复制并使用了 2014 年 PostgreSQL 的内部 JSON 词法器和语义动作接口。这些并非稳定扩展 API，仓库也没有现代大版本兼容矩阵。函数只接收 json 而不是 jsonb，并且每个请求路径都会重新解析整份文档，因此开销会随文档大小和键数量同时增长。

自定义分隔符版本把文本转换成 C 字符串，却只使用第一个字节；源码自己警告特殊多字节编码会出问题。结果统一转换成文本，JSON 类型差异会丢失。深层输入、大量键、大型终点数组、重复路径、空路径段、多字节分隔符和当前服务器头文件都需要回归测试。

仓库自 2014 年后没有更新。应把 jsonfun 视为遗留源码；除非必须复现其历史行为，否则优先使用当前 jsonb 操作符和 SQL/JSON 路径功能。
