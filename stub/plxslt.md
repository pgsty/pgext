


## Usage

> [plxslt: XSLT procedural language for PostgreSQL](https://github.com/petere/plxslt)

`plxslt` enables writing PostgreSQL functions as XSLT stylesheets for transforming XML data.

```sql
CREATE EXTENSION plxslt;
```

### Create XSLT Functions

The function body is an XSLT stylesheet. The first parameter must be of type `xml` and receives the input document:

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

### Return Types

The return type must match the stylesheet's output method:

| Output Method | Return Type |
|---------------|-------------|
| `xml` | `xml` |
| `text` | `text` or `varchar` |
| `html` | `text` or `varchar` |

### Passing Parameters

Additional function parameters beyond the first `xml` parameter are passed as XSL stylesheet parameters:

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

### Limitations

- First parameter must be `xml` type
- Triggers are not supported
- Only XSLT 1.0 transformations are supported (via libxslt)
