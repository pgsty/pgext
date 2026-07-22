## 用法

来源：

- [官方 1.0 版 SQL](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd--1.0.sql)
- [官方验证器实现](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.c)
- [官方回归测试运行器](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/regress/run.pl)
- [官方 control 文件](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.control)

`pgxsd` 使用数据库中保存的 XML Schema 文档验证 PostgreSQL `xml` 值。它把 XSD 文档保存在固定的 `pgxsd` 模式中，并从该表解析导入的 schema location，而不是访问网络。

### 核心流程

以验证函数第二个参数或 XSD import 使用的准确 location 保存每个模式，然后验证 XML 文档：

```sql
CREATE EXTENSION pgxsd;

INSERT INTO pgxsd.schemata (schema_location, document)
VALUES (
  'invoice.xsd',
  '<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
     <xs:element name="invoice" type="xs:integer"/>
   </xs:schema>'::xml
);

SELECT pgxsd.schema_validate(
  '<invoice>42</invoice>'::xml,
  'invoice.xsd'
);
```

成功时返回 `void`。文档或模式无效时会抛出 `invalid_xml_document`；上游回归测试期望验证失败的 SQLSTATE 为 `2200M`。

### 对象

- `pgxsd.schemata(schema_location text PRIMARY KEY, document xml)`：XSD 注册表与 import 解析器。
- `pgxsd.schema_validate(doc xml, schemaname text) RETURNS void`：使用 libxml2 解析命名模式并验证 `doc`。

### 注意事项

扩展要求 PostgreSQL 构建包含 libxml/libxml2 schema 支持。Schema location 是应用键，不是会被获取的 URL；每个导入 XSD 的 location 必须在 `pgxsd.schemata` 中恰好存在一条。

版本 `1.0` 把 `schema_validate` 声明为 `IMMUTABLE`，但函数实际上通过 SPI 读取可变的 `schemata` 表。这个易变性声明不正确：不要在存储生成表达式或索引中使用它，模式改变后还应强制重新求值。XML/XSD 解析与实体展开也可能消耗大量 CPU/内存；应限制谁能插入模式或提交文档，设置工作负载限制，并测试恶意嵌套输入。
