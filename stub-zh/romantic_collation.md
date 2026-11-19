## 用法

来源：

- [官方资料](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/docs/DATABASES.md)
- [官方资料](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/db/romantic_collation.control)
- [官方资料](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/LICENSE)
- [官方资料](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/db/romantic_collation--1.0.sql)

romantic_collation 以 bytea 排序键提供可复现的自然顺序，不注册 PostgreSQL COLLATE 提供者。原生适配器要求数据库编码为 UTF8，并使用匹配服务端头文件的库；未声明固定主版本支持列表。

### 生成排序键与索引

```sql
CREATE EXTENSION romantic_collation;
SELECT label FROM (VALUES ('item10'), ('item2')) AS x(label)
ORDER BY romantic_key_v1(label, 0);
CREATE TABLE labels(label text);
CREATE INDEX labels_natural ON labels(romantic_key_v1(label, 0));
```

romantic_key_v1(text, integer DEFAULT 0) 标记为 immutable、strict 和 parallel safe。查询与索引表达式必须采用相同选项位：码点模式为 1，小数解析为 8，科学计数法为 32，罗马数字解析为 64。默认采用库的自然排序语义。

### 持久化

持久化或索引的键应固定版本与选项。排序键语义变化时，应重新计算受影响的键并重建表达式索引，较早版本的科学计数法键不能混用。键相等不代表输入字节完全一致；过长的键可能超过 PostgreSQL B-tree 键长度限制。无需预加载。
