## 用法

来源：

- [官方 pg_registry README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_registry/README.md)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_registry/pg_registry.control)
- [SQL API 与存储定义](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_registry/src/lib.rs)

`pg_registry` 版本 `0.2.0` 在 PostgreSQL 中保存带版本的 JSON Schema、验证 JSONB 值、从表生成模式，并为配套流处理集成记录模式到主题或表到主题的绑定。

### 核心流程

```sql
CREATE EXTENSION pg_registry;

SELECT pgregistry.register_schema(
    'orders-value',
    '{"type":"object","properties":{"order_id":{"type":"integer"}},"required":["order_id"]}'::jsonb,
    'Order event schema'
) AS schema_id;

SELECT pgregistry.validate(
    1,
    '{"order_id":123}'::jsonb
);

SELECT pgregistry.bind_schema_to_topic(
    'orders', 1, 'value', 'STRICT'
);
```

重要函数包括 `register_schema`、`register_schema_from_table`、`generate_schema_from_table`、`get_schema`、`get_latest_schema`、`drop_schema`、`validate` 和 `validate_for_topic`。主题函数管理键/值模式绑定；表函数可以从模式创建表，并保存流式或 upsert 绑定元数据。

### 运维说明

对象位于不可重定位的 `pgregistry` 模式中，control 文件要求超级用户安装。`STRICT` 与 `LOG` 是保存下来的验证模式，但 Kafka 流量中的实际执行依赖配套集成读取这些绑定；`pg_registry` 本身不会启动 Kafka 客户端。生成的模式和表遵循文档所列 PostgreSQL 到 JSON 类型映射，但可能无法保留所有域、约束、生成表达式或应用不变量。生产使用前应审查生成的 DDL 以及模式管理函数的访问权限。
