## 用法

来源：

- [扩展 SQL](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser--1.0.sql)
- [扩展 control 文件](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.control)
- [C 实现](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.c)

`test_parser` 版本 `1.0` 安装名为 `testparser` 的 C 文本搜索解析器。可使用 PostgreSQL 的解析器诊断函数检查其输出词元：

```sql
CREATE EXTENSION test_parser;

SELECT *
FROM ts_parse('testparser', 'The quick brown fox 123');
```

### 项目范围

固定版本的上游源代码包含解析器回调和扩展对象，但没有 README、发行说明、许可证文件或 PostgreSQL 兼容矩阵。应将其视为最小解析器示例：使用前必须针对实际服务器大版本进行构建和回归测试，不能根据扩展名称推断其具备生产支持。
