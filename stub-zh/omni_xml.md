


## 用法

> [omni_xml: XML 工具箱](https://docs.omnigres.org/omni_xml/overview/)

`omni_xml` 扩展使用 pugixml 提供 XML 工具箱功能，当 PostgreSQL 未编译 XML 支持或未安装 `xml2` 扩展时非常有用。

### XPath 查询

**`omni_xml.xpath(xml_text, xpath_query)`** -- 返回包含 `path` 和 `value` 列的表。

**元素查询：**

```sql
SELECT * FROM omni_xml.xpath('<node>value</node>', '/node');
-- 返回完整的 XML 元素
```

**文本提取：**

```sql
SELECT * FROM omni_xml.xpath('<node>value</node>', '/node/text()');
-- 仅返回文本内容
```

**命名空间处理：**

```sql
SELECT * FROM omni_xml.xpath('<ns:node>value</ns:node>', '/ns:node');
```

注意：底层库未完全符合 W3C 标准；不支持 `namespace::` 轴。
