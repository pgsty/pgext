## 用法

来源：

- [官方 README](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/README.md)
- [0.1.1 扩展 SQL](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/jx_io--0.1.1.sql)

`jx_io` 在 PostgreSQL 与数据库服务器可见的文件之间导入或导出 JSON、XML。它会在 `jx_io` 模式中创建对象，适用于受控的服务端数据交换，不能用来访问任意客户端本地路径。

### 核心流程

```sql
CREATE EXTENSION jx_io;

SELECT jx_io.export_json_query(
  '/srv/exports/active_users.json',
  'SELECT id, name FROM app_user WHERE active'
);

SELECT jx_io.import_json_doc('/srv/imports/document.json');
```

导出函数在文件写入成功后返回 true。导入函数会返回 JSON/XML 文档，或者将按行保存的 JSON 文档复制到目标列中。

### 主要函数

- `export_json_query(filename text, query text)` 与 `export_json_table(filename text, tbl text)`：写出 JSON 行。
- `export_xml_query(filename text, query text)` 与 `export_xml_table(filename text, tbl text)`：写出 XML。
- `import_json_doc(filename text)`：返回一个 JSON 文档。
- `import_json_docs_into_table(filename text, tbl text, col text)`：将文档复制到表列中。
- `import_xml(filename text)`：返回一个 XML 文档。

### 运维与安全边界

路径由 PostgreSQL 服务器解释。导出函数会构造并执行服务端 `COPY` 命令，导入函数使用 PostgreSQL 大对象文件访问能力。因此调用者需要相应的服务端文件权限，PostgreSQL 操作系统账号也必须能够访问该路径。

0.1.1 版本通过拼接传入的查询、表名、列名、文件名和路径文本来构造动态 SQL。只能传入管理员控制的值，不要将这些函数直接开放给不可信用户或应用输入。上游 README 还明确警告该版本存在已知缺陷，因此投入生产前应验证数据与错误处理行为。
