## 用法

来源：

- [官方 README](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/README.md)
- [扩展控制文件](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/uuid47.control)
- [扩展 SQL](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/sql/uuid47--1.0.sql)
- [PostgreSQL 实现](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/src/uuid47_pg.c)

`uuid47` 在内部保存按时间排序的 UUIDv7 字节，但通过文本输出和类型转换呈现确定性的 UUIDv4 外观。128 位密钥驱动由 SipHash 派生的可逆时间戳位掩码；这是一种有密钥管理要求的标识符混淆，不是对行数据的加密。

### 核心流程

任何文本输入、输出、转换、转储或客户端交换之前都要设置稳定密钥。应通过受保护的会话或角色配置注入，不要把密钥写进应用 SQL。

```sql
CREATE EXTENSION uuid47;
SET uuid47.key = '0011223344556677:8899aabbccddeeff';

CREATE TABLE event (
  event_id uuid47 PRIMARY KEY DEFAULT uuid47_generate_monotonic(),
  payload jsonb NOT NULL
);

INSERT INTO event(payload) VALUES ('{"kind":"created"}');

SELECT event_id,
       uuid47_timestamp(event_id),
       uuid47_as_v7(event_id),
       uuid47_key_fingerprint()
FROM event;
```

`uuid47_generate()` 生成随机 UUIDv7 后缀；`uuid47_generate_monotonic()` 在同一毫秒内提供每后端单调顺序；`uuid47_generate_at(timestamptz)` 使用给定时间。`uuid47_timestamp`、`uuid47_as_v7`、`uuid47_explain` 用于检查内部 UUIDv7 状态。`uuid47` 与 `uuid` 间的赋值转换使用会话密钥，显式密钥版本 `uuid47_to_uuid_with_key` 与 `uuid_to_uuid47_with_key` 接受 16 字节密钥。

该类型提供比较操作符、默认 B-tree 与哈希操作符类，以及非默认的 `uuid47_brin_minmax_multi_ops`。比较和索引基于内部 UUIDv7 字节，因此按时间排序，不受显示外观影响。

### 密钥与交换注意事项

`uuid47.key` 是 `USERSET`，每个会话都能选择不同密钥。同一个存储值会显示成不同的 UUIDv4 外观，而用一个密钥生成的外观在另一个密钥下会解码为错误 UUIDv7。应在可信客户端间统一密钥，避免写入日志，在带外保存密钥标识，并用预期密钥测试备份恢复。纯文本转储会调用带密钥的输出；丢失密钥会改变外部标识符表面。

密钥轮换必须兼容已经发出的外观标识；只修改设置无法保持它们的解码结果。掩码会隐藏直接时间戳外观，但保留随机位，不能替代授权或机密性。若消费者必须跨密钥轮换持久保存标识，或不能接受关联风险，应使用真正独立的外部 ID。不需要预加载，但二进制和 BRIN 支持必须匹配目标 PostgreSQL 大版本。
