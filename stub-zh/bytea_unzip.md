## 用法

来源：

- [README.md](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/README.md)
- [bytea_unzip.control](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/bytea_unzip.control)
- [bytea_unzip--1.0.sql](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/bytea_unzip--1.0.sql)
- [Makefile](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/Makefile)
- [License](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/License)

`bytea_unzip` 从 bytea 值读取 ZIP 成员，并将 OpenDocument 电子表格中的表节点拆成逐行文本数组。SQL 版本 1.0 使用 libzip 和 libxml2。上游文档要求 PostgreSQL 16 或更高版本，面向 Linux/POSIX 环境；函数处理调用者提供的二进制值。

### 核心用法

由超级用户安装。control 文件允许迁移模式，不要求预加载。

```sql
CREATE EXTENSION bytea_unzip;
CREATE TABLE archives (id bigint PRIMARY KEY, payload bytea);
SELECT id, bytea_is_zip(payload) FROM archives;
SELECT a.id, f.*
FROM archives AS a
CROSS JOIN LATERAL bytea_unzip_list_files(a.payload) AS f;
SELECT id, convert_from(bytea_unzip(payload, 'mimetype'), 'UTF8')
FROM archives;
```

通过应用程序将真实归档字节写入表中。`bytea_is_zip()` 检查 libzip 能否识别归档，不保证其中包含有效文件。`bytea_unzip()` 以 bytea 返回指定成员，成员不存在时返回 NULL，输入不是有效 ZIP 时报告错误。文件列表返回名称、解压前后大小及修改时间。

### OpenDocument 表格

```sql
WITH sheets AS (
  SELECT unnest(xpath(
    '//office:body/office:spreadsheet/table:table',
    convert_from(bytea_unzip(payload, 'content.xml'), 'UTF8')::xml,
    ARRAY[
      ARRAY['office', 'urn:oasis:names:tc:opendocument:xmlns:office:1.0'],
      ARRAY['table', 'urn:oasis:names:tc:opendocument:xmlns:table:1.0']
    ]
  )) AS sheet
  FROM archives
)
SELECT decompose_ods_table_node(sheet) FROM sheets;
```

`decompose_ods_table_node()` 的输入是一个 XML 表节点，返回文本数组集合。`bytea_unzip_version()` 和 `bytea_unzip_libzip_version()` 提供扩展与库的版本号。函数声明为严格函数，NULL 参数会得到 NULL。

### 边界

已核实的 SQL 接口是一组函数，不是外部数据包装器。README 中的 HTTP 下载示例使用独立的 `http` 扩展，它是可选项，不是这些函数的依赖。README 中无关的空间结果描述与安装 SQL 不符。服务器必须能加载 libzip 和 libxml2；应用程序应按自身资源限制管理归档输入。
