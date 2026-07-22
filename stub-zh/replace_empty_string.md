## 用法

来源：

- [官方 README](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/README.md)
- [1.0 版扩展 SQL](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/replace_empty_string--1.0.sql)
- [触发器实现](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/replace_empty_string.c)

`replace_empty_string` 提供通用行级触发器，在插入或更新之前把零长度字符串改为 SQL 空值。它适合预期空字符串与空值行为相同的 Oracle 迁移，但会改变所挂接表中每个字符串类别列的数据语义。

### 核心流程

安装扩展，并在每张需要转换的表上把其触发器函数挂为 `BEFORE` 行级触发器：

```sql
CREATE EXTENSION replace_empty_string;

CREATE TABLE contacts (
  contact_id bigint PRIMARY KEY,
  display_name text,
  external_code varchar(32)
);

CREATE TRIGGER contacts_replace_empty_string
BEFORE INSERT OR UPDATE ON contacts
FOR EACH ROW
EXECUTE FUNCTION replace_empty_string();

INSERT INTO contacts VALUES (1, '', 'A-001');
SELECT display_name IS NULL FROM contacts WHERE contact_id = 1;
```

触发器扫描 PostgreSQL 类型类别为字符串的所有属性，将域解包为基础类型，并且只替换存储长度为零的值。仅包含空白字符的字符串不会改变。

### 警告模式

传入字面触发器参数 `on`，可为每个被修改的列发出警告：

```sql
DROP TRIGGER contacts_replace_empty_string ON contacts;

CREATE TRIGGER contacts_replace_empty_string
BEFORE INSERT OR UPDATE ON contacts
FOR EACH ROW
EXECUTE FUNCTION replace_empty_string('on');
```

不传参数或第一个参数为其他值时不会发出警告。批量加载期间，警告模式可能产生大量日志与客户端消息，因此更适合迁移摸底或测试，不宜作为默认可观测机制。

### 运维说明

该函数必须由行级 `BEFORE INSERT OR UPDATE` 触发器调用；C 实现会拒绝直接调用、语句级触发器、后置触发器与删除事件。创建触发器不会重写已有行；若历史数据也必须统一，应显式回填。

1.0 版可重定位，且没有声明预加载或扩展依赖。上游 README 称这段代码已合并到 Orafce，因此新的 Oracle 兼容部署应评估维护中的 Orafce 实现。把空字符串转为空值会影响 `NOT NULL` 约束、唯一索引、谓词、应用序列化与长度检查；大范围挂接前必须测试这些语义。
