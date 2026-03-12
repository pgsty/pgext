

## Usage

> [pg_sphere: spherical geometry data types and operations for PostgreSQL](https://github.com/postgrespro/pgsphere)

The `pg_sphere` extension provides spherical geometry types and operations, useful for astronomy, geospatial, and other applications involving spherical coordinates.

```sql
CREATE EXTENSION pg_sphere;
```

### Data Types

- **Spherical point (`spoint`)**: A location on a sphere (longitude, latitude in radians)
- **Spherical circle (`scircle`)**: A circular region defined by center and radius
- **Spherical line (`sline`)**: A great circle segment
- **Spherical ellipse (`sellipse`)**: An elliptical region on a sphere
- **Spherical polygon (`spoly`)**: A multi-vertex shape on a sphere
- **Spherical path (`spath`)**: A connected sequence of points
- **Spherical box (`sbox`)**: A bounding box on a sphere

### Core Operations

- Membership testing (point in polygon, point in circle, etc.)
- Overlap detection between spherical objects
- Circumference and area calculations
- Object rotation using Euler angles for coordinate transformations
- Distance calculations between spherical objects

### Index Support

- **GiST index**: R-tree implementation for efficient spatial queries
- **BRIN index**: Block range indexing for large datasets

### Input/Output

The extension handles input and output in various coordinate formats commonly used in astronomy and geospatial applications, including degrees, radians, and HMS/DMS notation.

```sql
-- Create a spherical point (RA, Dec)
SELECT spoint '(10.25d, 45.5d)';

-- Create a spherical circle (center, radius)
SELECT scircle '<(10.25d, 45.5d), 1d>';

-- Check containment
SELECT spoint '(10.25d, 45.5d)' @ scircle '<(10d, 45d), 2d>';
```
