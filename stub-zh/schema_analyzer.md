## 用法

来源：

- [Official schema_analyzer.control](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer.control)
- [Official schema_analyzer--1.0.0.sql](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer--1.0.0.sql)
- [Official schema_analyzer.c](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer.c)
- [Official README.md](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/README.md)

`schema_analyzer` 1.0.0 为 Sinew 提供触发器函数，统计文档字段出现次数并识别高频字段。它是 `document_type` 的配套扩展，不是面向任意 JSONB 表的通用触发器。此源码来自 2014 年的早期研究系统。

### 启用配套扩展

```sql
CREATE EXTENSION document_type;
CREATE EXTENSION schema_analyzer;
```

### 触发器约定

`analyze_document()` 处理行变更，在 Sinew 元数据表中增加或减少字段计数。实现从第 2 列读取文档，并要求 `document_schema` 中存在与关系对应的表，含字段 ID、计数、脏标记与提升标记。`analyze_schema()` 使用行数统计与 0.5 的频率阈值，标记需要提升或降级的字段。这些函数返回触发器值，应由触发器管理器调用，不能作为普通 SELECT 函数调用。

### 运行限制

请先安装文档扩展，因为安装脚本会在其缺失时尝试创建。仅安装本扩展不会配置应用表、各表元数据或触发器；这些对象必须遵循 Sinew 布局。C 代码将关系名插入元数据 SQL，并假定特定的二进制文档布局，因此只应在可信、受控的研究模式中使用。安装需要管理员权限，未要求预加载，尚未确认与当前 PostgreSQL 兼容。
