## 用法

来源：

- [官方 README](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/README.md)
- [扩展 SQL 定义](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath--1.0.sql)
- [扩展控制文件](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath.control)

`fhirpath` 是一个用 C 原生实现的 FHIRPath/FluentPath 求值器，面向 PostgreSQL `jsonb`。它可以把路径编译成 `fhirpath` 类型，并从 FHIR 形态的 JSON 文档中提取或规范化值。上游仓库只提供精简的底层 SQL 接口，并不是完整的应用框架。

### 核心流程

安装扩展，将路径表达式转换为 `fhirpath`，再把它与 JSON 文档一同传给提取函数：

```sql
CREATE EXTENSION fhirpath;

SELECT fhirpath_extract(
    '{"resourceType":"Patient","active":true}'::jsonb,
    'resourceType'::fhirpath
);
```

专用访问函数还接受额外的文本参数，供实现选择或规范化值。将它们用于生成列或索引前，应以有代表性的 FHIR 文档验证实际行为。

### 用户接口

- `fhirpath`：可变长度类型，其输入解析器会编译路径表达式。
- `fhirpath_extract(jsonb, fhirpath)`：以 `jsonb` 返回选中的值。
- `fhirpath_as_string(jsonb, fhirpath, text)`：以 `text` 返回选中的值。
- `fhirpath_as_token(jsonb, fhirpath, text)` 与 `fhirpath_as_reference(jsonb, fhirpath, text)`：返回 `text[]` 形式的令牌或引用。
- `fhirpath_as_number(jsonb, fhirpath, text, text)`：返回 `numeric`。
- `fhirpath_as_date(jsonb, fhirpath, text, text)` 与 `fhirpath_date_bound(text, text)`：返回 `timestamptz` 值。
- `fhirpath_exists(jsonb, fhirpath, text)`：检查所选值是否存在。

### 运维说明

控制文件将 `fhirpath` 标记为可迁移，且不要求预加载。其 SQL 函数声明为 `STRICT IMMUTABLE`，因此在输入不可变时可用于表达式索引。上游没有发布正式 Release，README 也主要是开发材料，没有声称完整符合某个特定版本的 FHIRPath 规范。应把版本 `1.0` 视为实验性 API，直接测试路径语法与边界情况，并固定已核验的源码修订。
