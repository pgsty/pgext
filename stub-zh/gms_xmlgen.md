## 用法

来源：

- [官方资料](https://docs.opengauss.org/en/docs/7.0.0/extension_reference/extension_reference/server/gms_xmlgen.html)
- [官方资料](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlgen/gms_xmlgen.control)
- [官方资料](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlgen/gms_xmlgen--1.0.sql)

`gms_xmlgen` 1.0 在 openGauss 中将查询结果转换为 XML。它面向该内核的 Oracle 兼容 A 模式数据库，依赖 libxml，轻量版本不可用；这不代表它兼容原生 PostgreSQL。

### 核心流程

安装匹配的 openGauss 扩展文件并具备所需安装权限后：

```sql
CREATE EXTENSION gms_xmlgen;
SELECT gms_xmlgen.getxml('SELECT 1 AS id');
```

接口返回 CLOB 或 XMLTYPE 值。`gms_xmlgen.newcontext` 建立可复用的查询上下文，`gms_xmlgen.getxml` 获取 XML，`gms_xmlgen.closecontext` 释放上下文。`gms_xmlgen.setmaxrows` 与 `gms_xmlgen.setskiprows` 控制行数和跳过行；对应辅助接口还可配置行标签、结果集标签与空值处理。

### 兼容性与维护

XMLTYPE 接口要求设置 `behavior_compat_options = 'bind_procedure_searchpath'`。所需 XMLTYPE 支持在 openGauss 7.0.0-RC1 中引入。这是通过 SQL 启用的扩展，没有预加载要求。

应核对查询权限，避免将不可信 SQL 字符串交给生成器，并及时释放上下文。由于 XMLTYPE 依赖，厂商要求在相关内核升级或回滚流程之前移除此扩展；在有数据的数据库中操作前，应清点依赖对象并演练恢复。
