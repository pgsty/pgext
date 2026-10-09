## 用法

来源：

- [packages/infra-utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/README.md)
- [packages/infra-utils/sql/infra-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/sql/infra-utils--0.47.0.sql)
- [packages/infra-utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/Makefile)
- [packages/infra-utils/infra-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/infra-utils.control)

`infra-utils` 在 PostgreSQL 内编译基础设施资源参数。`infra_utils.compile_resource_spec` 接受基础规格、参数声明、输入值和成员名称；`infra_utils.quantity_to_numeric` 转换所支持的十进制及二进制数量后缀。

### 核心用法

```sql
CREATE EXTENSION "infra-utils";
SELECT infra_utils.quantity_to_numeric('4Gi');
SELECT infra_utils.compile_resource_spec(
  '{}'::jsonb,
  '[{"key":"replicas","type":"int","min":1,"max":10,"bindings":[{"path":"replicas"}]}]'::jsonb,
  '{"replicas":2}'::jsonb,
  'web'
);
```

### 运行边界

声明可限制参数类型、范围与明确的输出路径。`infra_utils.validate_params_schema` 和 `infra_utils.validate_bundle_params` 检查声明与资源包输入，声明接口中未知的成员参数会报错。如果所有成员都未声明参数接口，编译路径会改为合并输入 JSONB，不能将此模式当作白名单。应用准入辅助函数还会查询已配置的基础设施模式。无效或不支持的数量字符串返回 NULL。 依赖为 `plpgsql`, `pgpm-verify`, `errors`, `db-utils`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
