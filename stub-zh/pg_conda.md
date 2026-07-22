## 用法

来源：

- [官方 README](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/README.md)
- [类型与函数实现](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/src/lib.rs)
- [扩展控制文件](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/pg_conda.control)
- [构建兼容性声明](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/Cargo.toml)

`pg_conda` 提供 `condaversion` 类型，使用 Rattler 的 Conda 排序规则存储和比较 Conda 生态中的版本字符串。它适合包元数据与依赖分析，但应先用实际数据中的版本拼写测试其比较语义。

### 核心流程

```sql
CREATE EXTENSION pg_conda;

CREATE TABLE package_release (
    package text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO package_release VALUES
    ('demo', '1.5.0'),
    ('demo', '2.0.0'),
    ('demo', '1.0.0.post1');

SELECT package, min(version), max(version)
FROM package_release
GROUP BY package;

SELECT conda_is_post('1.0.0.post1'::condaversion),
       conda_segments('1.2.3.alpha1'::condaversion);
```

输入由 Rattler 的 Conda 版本解析器处理，输出则保留原始文本拼写。

### 对象

- `condaversion` 支持相等、排序、哈希、B-tree 操作符，以及与 `text` 之间的转换。
- `min(condaversion)` 和 `max(condaversion)` 使用扩展自己的排序实现。
- `conda_is_dev(condaversion)` 与 `conda_is_post(condaversion)` 识别开发版和后发布组件。
- `conda_has_epoch(condaversion)` 与 `conda_has_local(condaversion)` 识别 epoch 和本地组件。
- `conda_segments(condaversion) → text[]` 返回解析后的版本组件。
- `conda_major(condaversion) → bigint` 与 `conda_minor(condaversion) → bigint` 在可用时返回数值组件；遇到不支持的组件形态时返回 `NULL`。

### 比较注意事项

在核验的 `1.0.1` 源码中，相等与哈希比较规范化后的 Rattler 版本；全序比较则先比较该版本，再比较保留的源文本。因此，`1.1.0` 与 `1.01.0` 等等价拼写可能判定为相等，却具有不同排序结果。这违反 PostgreSQL B-tree 比较所期望的一致性，可能影响 B-tree/唯一索引，以及 `min(condaversion)` 或 `max(condaversion)` 返回的具体值。

在上游修复前，应规范化存储的拼写，或避免依赖这些索引与聚合语义。请用实际版本语料测试相等、排序、唯一性和升级行为。

### 兼容性

控制文件没有声明为受信扩展，扩展也不可重定位，因此安装通常需要特权角色，且对象保留在安装模式中。当前 Cargo 特性声明支持 PostgreSQL 13 到 17，并默认 PostgreSQL 17；应确认软件包是为目标主版本专门构建的。

核验的源码树使用 `pgrx` 0.18.1 与 Rattler 0.36.0。无效输入会抛出 PostgreSQL 错误。由于 Rattler 解析器或排序规则的变化可能改变数据库可见的比较行为，应固定扩展与依赖版本。
