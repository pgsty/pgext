## 用法

来源：

- [官方上游文档](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/README.md)
- [官方扩展 SQL](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/sql/pg_libnumbertext--0.1.0.sql)
- [官方 C++ 实现](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/src/pg_libnumbertext.cpp)

`pg_libnumbertext` 0.1.0 将 PostgreSQL 绑定到 `libnumbertext`，把以文本提供的数字转换为指定语言的单词。项目自称处于早期 alpha 状态。它要求数据库服务器具备匹配的 C++ 库与语言数据文件，并且只支持 UTF-8 数据库。

### 将数字转换为单词

创建扩展，并以数字字符串和语言代码调用 `numbertext(text, text)`：

```sql
CREATE EXTENSION pg_libnumbertext;

SELECT numbertext('1', 'en');
SELECT numbertext('3', 'ja');
```

文档中的结果分别为 `one` 与 `三`。以文本传入数字可以保留数值转换可能丢失的格式选择；实际接受的形式和语言代码取决于已安装的 `libnumbertext` 数据。

### API 与数据文件

`numbertext(text, text)` 声明为 `IMMUTABLE` 与 `STRICT`：NULL 输入返回 NULL，PostgreSQL 可以在表达式索引中使用结果。运行时，库会加载安装在 PostgreSQL 扩展共享目录下的语言资源。应将二进制文件与这些数据文件作为一个带版本的整体部署到每台可能执行查询的服务器。

已复核的扩展 SQL 还会执行 `CREATE TYPE phone_number;`，但从未提供输入或输出函数来完成这个 shell type。它与有文档的数字转单词 API 无关，不能用作真实列类型。

### 兼容性与索引安全

数据库编码必须是 UTF-8。不支持的语言、畸形数字、缺失数据文件或库错误会抛出 PostgreSQL 错误。SQL 函数被标为不可变，但输出依赖外部资源文件，因此修改 `libnumbertext` 或语言数据可能使既有预期和表达式索引失效。应一起固定库、资源与扩展；升级前比较输出，结果可能变化时重建依赖索引。应为每种启用语言测试小数、符号、大数、区域专用语法和应用规范化。
