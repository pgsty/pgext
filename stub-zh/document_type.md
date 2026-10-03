## 用法

来源：

- [Official document_type.control](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/document/document_type.control)
- [Official document_type--1.0.0.sql](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/document/document_type--1.0.0.sql)
- [Official README.md](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/README.md)
- [Official bugs.txt](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/documentation/bugs.txt)

`document_type` 1.0.0 是早期 Sinew 研究系统使用的文档表示，保存结构化键值数据，并提供类型化访问与修改函数。已公布的源码停留在 2014 年，尚未确认与当前 PostgreSQL 兼容。

### 文档值

```sql
CREATE EXTENSION document_type;
SELECT document_get_int('{"n":7}'::document, 'n');
SELECT document_put_int('{"n":7}'::document, 'n', 8);
```

### 接口与元数据

`document` 类型提供字符串输入输出转换。`document_get()` 与 `document_delete()` 接受键及类型信息；类型化读取/写入变体覆盖整数、浮点、布尔、文本与嵌套文档。修改函数返回新文档值，调用方需要保存结果才能持久化变更。`document_schema._attributes` 为内部表示跟踪键名与类型。`schema_analyzer` 是可选的 Sinew 配套扩展，用于跟踪字段频率。

### 研究系统限制

安装需要管理员权限与匹配的 C 库，未要求预加载。虽然控制文件声明支持重定位，SQL 仍创建固定的文档元数据模式。上游记录了载入/验证不完整，以及畸形记录可能导致查询失败的问题。应将其视为保留的研究软件，不是持续维护的 JSONB 替代品；处理重要数据前应验证完整的 Sinew 环境。
