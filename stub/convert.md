

## Usage

> [convert: common unit conversion functions for PostgreSQL](https://github.com/rustprooflabs/convert)

Provides functions for common unit conversions: distance, speed, time-to-travel, power, area, and temperature.

```sql
CREATE EXTENSION convert;
```

### Distance Functions

| Function | Description |
|---|---|
| `dist_mi_to_ft(miles)` | Miles to feet |
| `dist_ft_to_mi(feet)` | Feet to miles |
| `dist_ft_to_m(feet)` | Feet to meters |
| `dist_m_to_ft(meters)` | Meters to feet |
| `dist_m_to_km(meters)` | Meters to kilometers |
| `dist_km_to_m(km)` | Kilometers to meters |
| `dist_mi_to_km(miles)` | Miles to kilometers |
| `dist_m_to_mi(meters)` | Meters to miles |
| `dist_km_to_mi(km)` | Kilometers to miles |

### Speed Functions

| Function | Description |
|---|---|
| `speed_mph_to_kmhr(mph)` | MPH to km/h |
| `speed_kmhr_to_mph(kmhr)` | km/h to MPH |
| `speed_kmhr_to_m_s(kmhr)` | km/h to m/s |
| `speed_mph_to_m_s(mph)` | MPH to m/s |
| `speed_m_s_to_kmhr(m_s)` | m/s to km/h |
| `speed_m_s_to_mph(m_s)` | m/s to MPH |

### Area Functions

| Function | Description |
|---|---|
| `area_m2_to_km2(m2)` | sq meters to sq km |
| `area_m2_to_ft2(m2)` | sq meters to sq feet |
| `area_ft2_to_m2(ft2)` | sq feet to sq meters |
| `area_ft2_to_mi2(ft2)` | sq feet to sq miles |
| `area_mi2_to_ft2(mi2)` | sq miles to sq feet |
| `area_mi2_to_acre(mi2)` | sq miles to acres |
| `area_acre_to_mi2(acres)` | Acres to sq miles |
| `area_acre_to_km2(acres)` | Acres to sq km |

### Temperature Functions

| Function | Description |
|---|---|
| `temp_c_to_f(celsius)` | Celsius to Fahrenheit |
| `temp_f_to_c(fahrenheit)` | Fahrenheit to Celsius |

### Power Functions

| Function | Description |
|---|---|
| `power_dbm_to_watts(dbm)` | dBm to watts |
| `power_watts_to_dbm(watts)` | Watts to dBm |

### Examples

```sql
SELECT dist_mi_to_km(26.2);      -- 42.16 (marathon in km)
SELECT temp_f_to_c(98.6);         -- 37.0
SELECT speed_mph_to_kmhr(60.0);   -- 96.56
SELECT area_acre_to_km2(640.0);   -- ~2.59
```
