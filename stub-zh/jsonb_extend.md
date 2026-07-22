## 用法

来源：

- [目录版本对应的官方 README](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/README.md)
- [目录版本对应的扩展 SQL](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/jsonb_extend--1.0.sql)

`jsonb_extend` 1.0 提供浅层与递归 JSONB 合并，适合需要明确合并顺序的场景：对象中靠后的值覆盖靠前的值，数组则进行拼接。源码依赖较老的 PostgreSQL 内部接口，因此应针对实际运行的服务器大版本构建和测试。

### 核心流程

```sql
CREATE EXTENSION jsonb_extend;

-- Shallow object merge: the rightmost value wins.
SELECT jsonb_extend('{"a":1,"nested":{"x":1}}'::jsonb,
                    '{"a":2,"nested":{"y":2}}'::jsonb);

-- Recursive object merge when the first flag is false.
SELECT jsonb_deep_extend(false,
                         '{"nested":{"x":1}}'::jsonb,
                         '{"nested":{"y":2}}'::jsonb);

-- Arrays are concatenated; a later scalar becomes another element.
SELECT jsonb_extend('[1,2]'::jsonb, '3'::jsonb);
```

### 函数索引

- `jsonb_extend(variadic jsonb[])` 执行浅层合并。首个值为对象时，后续值也必须是对象；首个值为数组时，后续值可以是数组、对象或标量。
- `jsonb_deep_extend(boolean, variadic jsonb[])` 接受对象；布尔标志为 false 时递归合并嵌套对象，为 true 时把嵌套值视为不可继续下钻的整体。
- 两个函数都声明为 immutable 和 strict。单个输入原样返回，形状不兼容时会报错。

### 注意事项

- 合并顺序会影响结果，替换语义也不等同于 RFC JSON Merge Patch。
- 实现基于 PostgreSQL 9.4 时代的 JSONB 内部接口；某个大版本可以成功构建，不代表其他大版本也兼容。
- 合并大型数组或深层对象前应评估结果大小，因为该操作会构造新的 JSONB 值。
