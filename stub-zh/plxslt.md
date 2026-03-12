

## 用法

> [plxslt: PostgreSQL 的 XSLT 过程语言](https://github.com/petere/plxslt)

`plxslt` 允许将 PostgreSQL 函数编写为 XSLT 样式表，用于转换 XML 数据。

```sql
CREATE EXTENSION plxslt;
```

### 创建 XSLT 函数

函数体是一个 XSLT 样式表。第一个参数必须是 `xml` 类型，用于接收输入文档：

```sql
CREATE FUNCTION extract_title(xml) RETURNS xml AS $$
<?xml version="1.0"?>
<xsl:stylesheet version="1.0"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
  <xsl:template match="/">
    <xsl:value-of select="//title"/>
  </xsl:template>
</xsl:stylesheet>
$$ LANGUAGE xslt;

SELECT extract_title('<doc><title>Hello World</title></doc>'::xml);
```

### 返回类型

返回类型必须与样式表的输出方法匹配：

| 输出方法 | 返回类型 |
|----------|----------|
| `xml` | `xml` |
| `text` | `text` 或 `varchar` |
| `html` | `text` 或 `varchar` |

### 传递参数

第一个 `xml` 参数之后的附加函数参数将作为 XSL 样式表参数传递：

```sql
CREATE FUNCTION transform_with_param(xml, text) RETURNS xml AS $$
<?xml version="1.0"?>
<xsl:stylesheet version="1.0"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
  <xsl:param name="arg2"/>
  <xsl:template match="/">
    <result>
      <xsl:value-of select="$arg2"/>
    </result>
  </xsl:template>
</xsl:stylesheet>
$$ LANGUAGE xslt;
```

### 限制

- 第一个参数必须是 `xml` 类型
- 不支持触发器
- 仅支持 XSLT 1.0 转换（通过 libxslt）
