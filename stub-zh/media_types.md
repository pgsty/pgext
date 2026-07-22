## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/media_types.control)
- [官方 Media Types README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/extension/usr/share/postgresql/17/extension/media_types--0.1.0.sql)

`media_types` 添加用于校验 `image/gif` 等媒体类型字符串的 `MediaType` 类型。需要解析并约束内容元数据，而不是保存未经检查的文本时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION media_types;

CREATE TABLE assets (
    path text PRIMARY KEY,
    content_type MediaType NOT NULL
);

INSERT INTO assets VALUES ('logo.gif', 'image/gif');
SELECT * FROM assets WHERE content_type = 'image/gif'::MediaType;
```

解析器根据构建时嵌入的注册表检查根类型和子类型，并保留可选参数。未知组合会被拒绝。

### SQL 对象

- `MediaType` 是变长类型，该版本的 PGRX 二进制协议使用 CBOR 序列化。
- 提供比较操作符 `=`、`<>`、`<`、`<=`、`>` 和 `>=`。
- `MediaType_btree_ops` 和 `MediaType_hash_ops` 支持 B-tree 与哈希索引。

### 运维说明

0.1.0 版不可重定位。控制文件设置 `superuser = true` 和 `trusted = false`，因此必须由超级用户创建；它未声明前置扩展或预加载要求。可接受的子类型注册表编译在扩展中。上游仍保留部分已弃用媒体类型并提供替代建议，因此应一起审查应用校验与扩展升级。
