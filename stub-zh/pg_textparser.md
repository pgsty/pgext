## 用法

来源：

- [上游 README](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/README.md)
- [扩展 SQL](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/pg_textparser--1.0.sql)
- [扩展 control 文件](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/pg_textparser.control)
- [上游仓库](https://github.com/za-arthur/pg_textparser)

`pg_textparser` 是一个已停止维护的自定义文本搜索解析器原型，预期解析器名为 `textparser`。可以先检查软件包可用性，而不要尝试启用：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'pg_textparser';
```

### 安装脚本不完整

已发布的版本 `1.0` SQL 声明了解析器的启动、取词、结束及词素类型函数，但其 `CREATE TEXT SEARCH PARSER` 语句引用了尚未声明为 SQL 函数的 `textparser_headline`。因此，使用未经修改的上游文件执行 `CREATE EXTENSION pg_textparser` 无法成功完成。该仓库没有被标记为 archived，但本目录将此条目标记为 abandoned；任何使用都需要经过审查的分支或补丁。
