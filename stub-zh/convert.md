

## 用法

> [convert: PostgreSQL 的常用单位转换函数](https://github.com/rustprooflabs/convert)

提供常用单位转换函数：距离、速度、行程时间、功率、面积和温度。

```sql
CREATE EXTENSION convert;
```

### 距离函数

| 函数 | 说明 |
|---|---|
| `dist_mi_to_ft(miles)` | 英里转英尺 |
| `dist_ft_to_mi(feet)` | 英尺转英里 |
| `dist_ft_to_m(feet)` | 英尺转米 |
| `dist_m_to_ft(meters)` | 米转英尺 |
| `dist_m_to_km(meters)` | 米转千米 |
| `dist_km_to_m(km)` | 千米转米 |
| `dist_mi_to_km(miles)` | 英里转千米 |
| `dist_m_to_mi(meters)` | 米转英里 |
| `dist_km_to_mi(km)` | 千米转英里 |

### 速度函数

| 函数 | 说明 |
|---|---|
| `speed_mph_to_kmhr(mph)` | 英里/时转千米/时 |
| `speed_kmhr_to_mph(kmhr)` | 千米/时转英里/时 |
| `speed_kmhr_to_m_s(kmhr)` | 千米/时转米/秒 |
| `speed_mph_to_m_s(mph)` | 英里/时转米/秒 |
| `speed_m_s_to_kmhr(m_s)` | 米/秒转千米/时 |
| `speed_m_s_to_mph(m_s)` | 米/秒转英里/时 |

### 面积函数

| 函数 | 说明 |
|---|---|
| `area_m2_to_km2(m2)` | 平方米转平方千米 |
| `area_m2_to_ft2(m2)` | 平方米转平方英尺 |
| `area_ft2_to_m2(ft2)` | 平方英尺转平方米 |
| `area_ft2_to_mi2(ft2)` | 平方英尺转平方英里 |
| `area_mi2_to_ft2(mi2)` | 平方英里转平方英尺 |
| `area_mi2_to_acre(mi2)` | 平方英里转英亩 |
| `area_acre_to_mi2(acres)` | 英亩转平方英里 |
| `area_acre_to_km2(acres)` | 英亩转平方千米 |

### 温度函数

| 函数 | 说明 |
|---|---|
| `temp_c_to_f(celsius)` | 摄氏度转华氏度 |
| `temp_f_to_c(fahrenheit)` | 华氏度转摄氏度 |

### 功率函数

| 函数 | 说明 |
|---|---|
| `power_dbm_to_watts(dbm)` | dBm 转瓦特 |
| `power_watts_to_dbm(watts)` | 瓦特转 dBm |

### 示例

```sql
SELECT dist_mi_to_km(26.2);      -- 42.16（马拉松的千米数）
SELECT temp_f_to_c(98.6);         -- 37.0
SELECT speed_mph_to_kmhr(60.0);   -- 96.56
SELECT area_acre_to_km2(640.0);   -- ~2.59
```
