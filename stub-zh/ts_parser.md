## 用法

来源：

- [README.md](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/README.md)
- [modules/ts_parser/ts_parser--1.0.sql](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_parser/ts_parser--1.0.sql)
- [modules/ts_parser/ts_parser.control](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_parser/ts_parser.control)

`ts_parser` 安装一个基于 PostgreSQL 9.6 默认解析器修改的全文检索解析器。它属于 APOD 检索演示，也具有独立的扩展安装入口。

### 核心用法

```sql
CREATE EXTENSION ts_parser;
SELECT * FROM ts_token_type('ts_parser');
SELECT * FROM ts_parse('ts_parser', 'PostgreSQL full text search');
```

### 运行边界

解析器使用版本化 SQL 声明的开始、下一词元、结束、词元类型及摘要回调。解析器本身并不是完整的检索配置；应用使用前还须明确选择字典和词元映射。

安装要求超级用户及编译后的 C 库，不要求预加载。历史源码来源不能证明它兼容当前 PostgreSQL 版本，上游未声明当前支持矩阵。
