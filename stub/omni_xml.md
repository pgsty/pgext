


## Usage

> [omni_xml: XML toolkit](https://docs.omnigres.org/omni_xml/overview/)

The `omni_xml` extension provides XML toolkit functionality using pugixml, useful when PostgreSQL is built without XML support or the `xml2` extension.

### XPath Querying

**`omni_xml.xpath(xml_text, xpath_query)`** -- Returns a table with `path` and `value` columns.

**Element queries:**

```sql
SELECT * FROM omni_xml.xpath('<node>value</node>', '/node');
-- Returns the complete XML element
```

**Text extraction:**

```sql
SELECT * FROM omni_xml.xpath('<node>value</node>', '/node/text()');
-- Returns just the text content
```

**Namespace handling:**

```sql
SELECT * FROM omni_xml.xpath('<ns:node>value</ns:node>', '/ns:node');
```

Note: The underlying library is not fully W3C-conformant; the `namespace::` axis is unsupported.
