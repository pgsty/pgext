## Usage

Sources:

- [Official README](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/README.md)
- [Extension SQL](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/sql/colors.sql)
- [C implementation](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/src/colors.c)

`colors` 0.0.1 adds four perceptual color-distance calculations for pairs of CIE L*a*b* coordinates. Use it when colors have already been converted to the same CIELAB white-point and observer convention; the extension does not convert RGB/XYZ values or carry color-management metadata.

### Core Workflow

Install the extension, then pass the six L*, a*, and b* coordinates as `double precision` values:

```sql
CREATE EXTENSION colors;

SELECT delta_e_cie_1976(50, 2.5, -20, 50, 0, -18);
SELECT delta_e_cie_2000(50, 2.5, -20, 50, 0, -18);
```

Smaller results mean the samples are closer under the selected formula. Pick one formula for a workflow and keep its parameters stable; values from different formulas are not directly interchangeable.

### Function Index

- `delta_e_cie_1976(...)` implements the six-coordinate CIE76 Euclidean distance.
- `delta_e_cie_1994(...)` has a six-coordinate overload with graphic-arts defaults and an eleven-argument overload exposing the weighting constants.
- `delta_e_cmc(...)` has a six-coordinate overload with the 2:1 acceptability defaults and an eight-argument overload exposing the lightness and chroma weights.
- `delta_e_cie_2000(...)` has a six-coordinate overload and a nine-argument overload exposing the lightness, chroma, and hue weights.

All functions are strict and immutable and return `double precision`.

### Operational Notes

The implementation neither validates the physical range of the coordinates nor transforms between color spaces. Normalize data before comparison and test the chosen formula against domain reference samples. The source is an old, small C implementation, so verify it on the exact PostgreSQL build and include edge cases in numerical regression tests.
