## 用法

来源：

- [官方资料](https://docs.opengauss.org/en/docs/latest/extension_reference/extension_reference/server/gms_xmlparser.html)
- [官方资料](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlparser/gms_xmlparser.control)
- [官方资料](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlparser/gms_xmlparser--1.0.sql)

`gms_xmlparser` 1.0 为 openGauss 的 Oracle 兼容场景提供 XML 解析与文档句柄。它依赖 libxml，轻量版本不可用。其 SQL 使用 openGauss 特有类型，不代表兼容原生 PostgreSQL。

### 启用与接口

```sql
CREATE EXTENSION gms_xmlparser;
```

扩展创建 `gms_xmlparser` 与 `gms_xmldom` 模式。`gms_xmlparser.newparser` 分配解析器，`gms_xmlparser.parse`、`gms_xmlparser.parsebuffer` 和 `gms_xmlparser.parseclob` 接收 XML 输入，`gms_xmlparser.getdocument` 返回解析后的文档。`gms_xmlparser.freeparser` 与 `gms_xmldom.freedocument` 分别释放相应资源。

### 资源与文件边界

解析器与文档值是内部句柄，不是可由用户直接构造的标量值。应在同一会话流程中完成分配、使用与释放，避免持续占用会话内存。

单字符串解析接口先尝试 XML，解析无效时会将输入解释为本地文件路径。输入验证和数据库主机的文件权限属于安全边界，不应将任意输入直接暴露成文件读取接口。安装应使用 openGauss 部署所要求的权限；文档没有预加载要求。删除扩展之前需要先处理依赖对象。
