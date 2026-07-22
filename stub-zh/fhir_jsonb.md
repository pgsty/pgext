## 用法

来源：

- [扩展控制文件](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb.control)
- [扩展 SQL](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb--1.0.sql)
- [官方回归查询](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/sql/fhir_jsonb.sql)

`fhir_jsonb` 是一个具有破坏性的合成数据夹具，用于在 FHIR 形态的文档上试验 PostgreSQL JSONB 查询。它不是 FHIR 服务器、验证器、模式实现或生产数据模型；只能安装在一次性数据库或隔离模式中。

### 核心流程

安装会立即删除并重建四张未限定名称的表，然后填充生成数据。请使用不存在同名重要对象的专用模式：

```sql
CREATE SCHEMA fhir_fixture;
CREATE EXTENSION fhir_jsonb WITH SCHEMA fhir_fixture;

SELECT count(*) FROM fhir_fixture.patients;
SELECT doc #>> '{name,0,text}' AS patient_name
FROM fhir_fixture.patients
LIMIT 5;
```

安装脚本创建 100 行的 `patients`、1,000 行的 `encounters`，以及各 10,000 行的 `observations` 和 `conditions`。每张表都有整数 `id` 和名为 `doc` 的 `jsonb` 列。

### 对象与查询

除四张夹具表外，扩展还安装 `patient_random_is_active()`、`encounter_random_status(integer)`、`observation_random_name(integer)` 和 `random_array_element(varchar[])` 等辅助函数。它们是数据生成器，并非稳定的医疗 API。

官方回归查询会单独安装 `jsquery`，并使用其 `@@` 操作符。控制文件与安装脚本均未把 `jsquery` 声明为依赖，因此只有准备运行这些历史示例时才应另行安装它。`#>>` 等常规 JSONB 操作符不需要它。

### 安全与兼容性

安装 SQL 会先对 `patients`、`encounters`、`observations` 和 `conditions` 执行 `DROP TABLE IF EXISTS`，再创建夹具。如果扩展目标模式中已有这些同名表，它们可能被销毁。绝不要把它安装到应用模式或生产数据库中。

生成的文档类似旧版 FHIR 资源的一个子集，但包含虚构个人数据，也不保证 FHIR 版本一致性、验证、术语绑定或引用完整性。核验的源码来自 2014 年，没有维护中的 PostgreSQL 兼容矩阵；请在实际服务器版本上测试，并把生成对象视为可丢弃的基准数据。
