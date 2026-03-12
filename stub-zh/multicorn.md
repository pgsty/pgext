

## 用法

> [multicorn: 在 PostgreSQL 服务器中使用 Python 获取外部数据](https://github.com/pgsql-io/multicorn2)

Multicorn2 允许您用 Python 编写外部数据包装器。您需要实现一个继承自 `multicorn.ForeignDataWrapper` 的 Python 类，Multicorn 负责将其桥接到 PostgreSQL 的 FDW 接口。

### 定义 Python FDW 类

创建一个 PostgreSQL 进程可访问的 Python 模块（例如 `myfdw.py`）：

```python
from multicorn import ForeignDataWrapper

class MyFDW(ForeignDataWrapper):
    def __init__(self, options, columns):
        super().__init__(options, columns)
        self.options = options
        self.columns = columns

    def execute(self, quals, columns):
        """以字典形式返回行。quals 包含 WHERE 下推信息。"""
        yield {"id": 1, "name": "example"}

    def insert(self, new_values):
        """处理 INSERT 操作。"""
        pass

    def update(self, old_values, new_values):
        """处理 UPDATE 操作。"""
        pass

    def delete(self, old_values):
        """处理 DELETE 操作。"""
        pass
```

### 创建服务器和外部表

```sql
CREATE EXTENSION multicorn;

CREATE SERVER multicorn_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'myfdw.MyFDW');

CREATE FOREIGN TABLE my_table (
  id integer,
  name text
)
SERVER multicorn_srv
OPTIONS (
  option1 'value1'
);

SELECT * FROM my_table;
```

`wrapper` 选项指定完全限定的 Python 类名。任何额外选项都会传递给类构造函数的 `options` 参数。

### 内置 FDW 示例

Multicorn 附带了几个可以直接使用或作为参考的 FDW 实现：

- **CsvFdw** -- 读取 CSV 文件
- **ProcessFdw** -- 执行系统命令并解析输出
- **GCalFdw** -- 访问 Google 日历
- **ImapFdw** -- 查询 IMAP 邮箱
- **RssFdw** -- 读取 RSS/Atom 订阅

```sql
CREATE SERVER csv_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'multicorn.csvfdw.CsvFdw');

CREATE FOREIGN TABLE csvtest (
  col1 text,
  col2 text
)
SERVER csv_srv
OPTIONS (
  filename '/tmp/data.csv',
  skip_header '1',
  delimiter ','
);
```
