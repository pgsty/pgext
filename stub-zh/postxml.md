## 用法

来源：

- [官方文档](https://github.com/albaike/postxml/blob/cc25905bf2fc0e77b65235ac10f67794ab4e61cd/README.md)
- [扩展控制文件](https://github.com/albaike/postxml/blob/cc25905bf2fc0e77b65235ac10f67794ab4e61cd/postxml.control)
- [官方仓库](https://github.com/albaike/postxml)

`postxml` 在 PostgreSQL 中构建超媒体文档的 XML 与 XSLT 转换助手。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `postxml`：

```sql
CREATE EXTENSION postxml;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT postxml.xml_docs_equal('<a/>'::xml, '<a/>'::xml);
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `postxml.init_xsd` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.read_text` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.xml_docs_equal` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.xsd_schemas` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `postxml.xsd_validate` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.xsd_validate_name` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.xsl_transform` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `postxml.xsl_transform_name` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 16；不要推断未列出的主版本。
