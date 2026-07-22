## 用法

来源：

- [官方 README](https://github.com/adjust/pg-device_type/blob/master/README.md)
- [版本 0.0.2 SQL](https://github.com/adjust/pg-device_type/blob/master/device_type--0.0.2.sql)
- [当前 C 实现](https://github.com/adjust/pg-device_type/blob/master/src/device_type.c)
- [当前 control 文件](https://github.com/adjust/pg-device_type/blob/master/device_type.control)

`device_type` 定义一个单字节、类似枚举的基础类型，用于固定的设备类别集合。它比文本紧凑，并支持比较、B-tree 与哈希索引，但排序依据实现内部的序号，而不是词法顺序。

### 核心流程

创建扩展后，用该类型保存受控的设备分类：

```sql
CREATE EXTENSION device_type;

CREATE TABLE session_device (
  session_id bigint PRIMARY KEY,
  device device_type NOT NULL
);

INSERT INTO session_device VALUES (1, 'phone'), (2, 'tablet');
CREATE INDEX ON session_device (device);
SELECT * FROM session_device WHERE device = 'phone';
```

输入会转换大小写，并且必须解析为受支持的标记之一。

### 取值与排序

可接受值按比较顺序依次为 `bot`、`console`、`ipod`、`mac`、`pc`、`phone`、`server`、`simulator`、`tablet`、`tv` 和 `unknown`。该类型是 C 基础类型，不是 PostgreSQL 枚举；扩展也未定义通用的文本或整数转换。

### 版本与运维说明

目录记录的是 0.0.2，而当前上游 control 文件声明 0.0.3。上游从 0.0.2 到 0.0.3 的升级脚本没有面向用户的 SQL 对象变化；已检查的值集合未改变，而当前构建会在受支持服务器上把函数标记为并行安全。由于磁盘字节与排序依赖 C 序号布局，未来插入或重排值都必须谨慎测试升级。部署前应针对选定的准确源码版本执行构建与回归测试。
