

## 用法

> [xml2: XPath 查询与 XSLT 转换功能](https://www.postgresql.org/docs/current/xml2.html)

`xml2` 扩展提供了 XML 文档的 XPath 查询和 XSLT 转换函数。注意：该模块已被弃用，建议使用核心 SQL/XML 功能替代。

```sql
CREATE EXTENSION xml2;
```

### XML 验证

```sql
SELECT xml_valid('<doc><item>test</item></doc>');  -- true
```

### XPath 查询函数

```sql
-- 提取文本值
SELECT xpath_string('<doc><name>Alice</name></doc>', '/doc/name');

-- 提取数值
SELECT xpath_number('<doc><count>42</count></doc>', '/doc/count');

-- 提取布尔值
SELECT xpath_bool('<doc><active>true</active></doc>', '/doc/active');

-- 提取带标签的节点集
SELECT xpath_nodeset('<doc><a>1</a><a>2</a></doc>', '/doc/a', 'results', 'item');

-- 提取值为逗号分隔的列表
SELECT xpath_list('<doc><a>1</a><a>2</a></doc>', '/doc/a');  -- 1,2
```

### xpath_table 函数

对一组文档执行多个 XPath 查询：

```sql
SELECT * FROM
  xpath_table('id', 'xml_col', 'documents',
              '/doc/author|/doc/title',
              'true')
  AS t(id int, author text, title text);
```

参数说明：
- 键字段名（第一个输出列）
- XML 文档字段名
- 表/视图名
- 以 `|` 分隔的 XPath 表达式
- WHERE 子句（使用 `'true'` 表示所有行）

### XSLT 转换

```sql
-- 对文档应用 XSL 样式表
SELECT xslt_process(xml_document, xsl_stylesheet);

-- 带参数
SELECT xslt_process(xml_document, xsl_stylesheet, 'param1=value1,param2=value2');
```
