## 用法

来源：

- [packages/utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/README.md)
- [packages/utils/sql/pgpm-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/sql/pgpm-utils--0.47.0.sql)
- [packages/utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/Makefile)
- [packages/utils/pgpm-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/pgpm-utils.control)

`pgpm-utils` 在工具模式中安装辅助函数。`utils.mask_pad` 保留指定长度的末尾字符，输入较短时在左侧填充；`utils.bitmask_pad` 对位串执行同类操作。此操作并非隐私脱敏。

### 核心用法

```sql
CREATE EXTENSION "pgpm-utils";
SELECT utils.mask_pad('123456', 4, '0');
SELECT utils.bitmask_pad(B'101', 8, '0');
```

### 运行边界

固定版本的 SQL 将 `utils.throw()` 定义为抛出错误的触发器，而非接收文本参数的标量函数。`utils.ensure_singleton()` 统计行数，只要存在任何行就报错，既不独立保证并发下的唯一性，也不支持 README 中通用的 UPDATE 示例。其他触发器包括 `utils.enforce_identity_providers_quota` 与 `utils.default_self_reference`。文档示例有冲突时，以版本化 SQL 为准；数据约束仍需真实约束和恰当的锁策略保障。 依赖为 `plpgsql`, `pgpm-verify`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
