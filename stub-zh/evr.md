## 用法

来源：

- [官方 README](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/README.md)
- [扩展控制文件](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr.control)
- [0.0.2 版扩展 SQL](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr--0.0.2.sql)

`evr` 是一个已归档的纯 PL/pgSQL 辅助扩展，用于把 RPM epoch/version/release 字符串拆成复合值。尽管项目描述称其为 EVR 数据类型，它并未实现 RPM 版本比较、排序运算符或可索引的 EVR 类型。

### 核心流程

安装扩展并查看版本字符串如何拆成数字和字母片段：

```sql
CREATE EXTENSION evr;

SELECT rpmver_array('1.20beta-3');
```

`evr_trigger()` 要求目标表必须具有名称完全一致的 `epoch`、`version`、`release` 和 `evr` 列：

```sql
CREATE TABLE package_version (
  package text,
  epoch integer,
  version text,
  release text,
  evr evr_t
);

CREATE TRIGGER package_version_evr
BEFORE INSERT OR UPDATE OF epoch, version, release ON package_version
FOR EACH ROW EXECUTE FUNCTION evr_trigger();

INSERT INTO package_version(package, epoch, version, release)
VALUES ('demo', 1, '2.4beta', '3.el9');
```

### 对象

- `evr_array_item` 是 `(n numeric, s text)` 复合片段：数字段写入 `n`，字母段写入 `s`。
- `evr_t` 存储 `epoch integer`、`version evr_array_item[]` 和 `release evr_array_item[]`。
- `rpmver_array(varchar)` 移除非字母数字分隔符并返回片段数组。
- `evr_trigger()` 根据四个固定列名填充 `NEW.evr`。
- `empty(text)`、`isalpha(char)`、`isalphanum(char)` 和 `isdigit(char)` 是解析器使用的公开辅助函数。

### 注意事项

解析仅处理 ASCII 并丢弃标点，因此无法保留原始版本字符串。它也没有实现 RPM 的比较规则，包括特殊版本语义；绝不要根据 `evr_t` 的文本表示对软件包排序或比较。

当 `version` 或 `release` 为 NULL 时，触发器会解析字面字符串 `empty`，而不是生成空数组。这些通用名称的辅助函数会安装到目标模式中，可能与应用函数冲突。仓库已经归档，因此应把它视为历史解析器；若需要精确 RPM EVR 语义，应另行实现比较逻辑。
