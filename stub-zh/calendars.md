## 用法

来源：

- [Official extension control file](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/calendars.control)
- [Official upstream documentation](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/README)

`calendars` 1.0 是一个尚未完成的实验，原本计划通过带日历名称的重载函数格式化和解析日期。仓库中的埃塞俄比亚历转换例程没有实现真实转换，因此不适合处理应用数据。

### 核心流程

```sql
CREATE EXTENSION calendars;

SELECT to_char(current_date, 'YYYY-MM-DD', 'Ethiopian');
SELECT to_date('2000-01-01', 'YYYY-MM-DD', 'Ethiopian');
```

三参数 `to_char(date, text, text)` 和 `to_date(text, text, text)` 重载通过名称选择日历。在审阅的源码中，第一个调用返回一段表示未完成状态的固定文本，第二个调用返回固定日期，并不转换输入。

### 注意事项

不要把这些结果用于存储、校验、报表或日期运算。通用的重载名称还可能影响定义了相似签名的数据库中的函数解析。在上游提供实际实现和测试向量之前，应仅限源码检查或隔离实验使用。
