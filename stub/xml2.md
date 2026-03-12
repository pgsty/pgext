

## Usage

> [xml2: XPath querying and XSLT functionality](https://www.postgresql.org/docs/current/xml2.html)

The `xml2` extension provides XPath querying and XSLT transformation functions for XML documents. Note: this module is deprecated in favor of the core SQL/XML functionality.

```sql
CREATE EXTENSION xml2;
```

### XML Validation

```sql
SELECT xml_valid('<doc><item>test</item></doc>');  -- true
```

### XPath Query Functions

```sql
-- Extract text value
SELECT xpath_string('<doc><name>Alice</name></doc>', '/doc/name');

-- Extract numeric value
SELECT xpath_number('<doc><count>42</count></doc>', '/doc/count');

-- Extract boolean
SELECT xpath_bool('<doc><active>true</active></doc>', '/doc/active');

-- Extract node set with tags
SELECT xpath_nodeset('<doc><a>1</a><a>2</a></doc>', '/doc/a', 'results', 'item');

-- Extract values as comma-separated list
SELECT xpath_list('<doc><a>1</a><a>2</a></doc>', '/doc/a');  -- 1,2
```

### xpath_table Function

Evaluate multiple XPath queries across a set of documents:

```sql
SELECT * FROM
  xpath_table('id', 'xml_col', 'documents',
              '/doc/author|/doc/title',
              'true')
  AS t(id int, author text, title text);
```

Parameters:
- Key field name (first output column)
- XML document field name
- Table/view name
- XPath expressions separated by `|`
- WHERE clause (use `'true'` for all rows)

### XSLT Transformation

```sql
-- Apply XSL stylesheet to document
SELECT xslt_process(xml_document, xsl_stylesheet);

-- With parameters
SELECT xslt_process(xml_document, xsl_stylesheet, 'param1=value1,param2=value2');
```
