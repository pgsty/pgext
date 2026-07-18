## 用法

来源：

- [项目 README](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/README.md)
- [扩展 control 文件](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson.control)
- [1.0 版 SQL API](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson--1.0.sql)
- [libbson 转换实现](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson.c)

`modeshape_bson` 1.0 把以 `bytea` 保存的 ModeShape BSON 转换为 canonical Extended JSON 文本，也能把 JSON 文本转回 BSON。其设计目标是在 ModeShape 所需的可写代理表背后使用 PostgreSQL JSONB 表。

### 测试转换

```sql
CREATE EXTENSION modeshape_bson;

SELECT modeshape_bson_version();

WITH encoded AS (
  SELECT json_text_to_modeshape_bson('{"name":"node","count":2}') AS value
)
SELECT modeshape_bson_to_json_text(value)
FROM encoded;
```

返回内容是 libbson 的 canonical Extended JSON，因此即使值保持等价，往返转换后的文本格式和表示也可能不同。应验证应用使用的每一种 BSON 类型，包括日期、二进制值、数值宽度、对象标识符、数组、null 和嵌套文档。

### 旧格式与代理表风险

上游记录了一个 ModeShape `repository:info` 记录，其 `$date` 值使用 BSON 类型 9，而不是通常观察到的 ISO 字符串，并明确表示不确定能否安全写回。在逐字节与语义测试通过前，应把该记录和所有罕见类型视为迁移阻断项。

源码最后变更于 2018 年，并内嵌旧 Mongo C driver/libbson 源码树。它没有发布当前 PostgreSQL、ModeShape 或 libbson 兼容矩阵。建议的可写视图设计用 rule 重定向全部读写，会增加并发、权限、触发器和错误传播复杂性。应使用最小权限 owner，不要在存储配置示例中放入凭据，迁移期间保留原始字节；除非已针对准确平台重建并审计原生代码，否则优先使用受维护的应用侧转换器。
