## 用法

来源：

- [官方 README](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/README.md)
- [扩展 SQL](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/compress--0.0.2.sql)
- [归档实现](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/compress_cpp.cpp)

`compress` 版本 `0.0.2` 用于在单文件 ZIP 归档的 Base64 文本与普通文本之间转换。它只适用于需要文本型 SQL 接口的场景，并不是通用的归档管理 API。

### 核心流程

```sql
CREATE EXTENSION compress;

SELECT zip_base64('message.txt', 'hello from PostgreSQL');

SELECT unzip_base64(
  'UEsDB...base64-encoded-zip...'
);
```

`zip_base64(filename, data)` 创建包含一个命名文件的 ZIP，并返回其 Base64 表示。`unzip_base64(archive)` 解码 Base64 输入、打开 ZIP，并以文本返回序号为零的条目内容。

### 对象

- `zip_base64(text, text) returns text`：创建单文件归档；第一个参数是归档内文件名，第二个参数是文件内容。
- `unzip_base64(text) returns text`：从 Base64 编码的 ZIP 归档中提取第一个条目。

两个函数都声明为 `IMMUTABLE` 和 `STRICT`。

### 运维说明

接口使用 `text` 而非 `bytea`，实现会在后端内存中处理归档。它不提供条目选择，也不会提取多个文件。接收不可信压缩数据前应限制输入大小，因为解压后的输出可能远大于 Base64 输入。安装库后不需要预加载或重启。
