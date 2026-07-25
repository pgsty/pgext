#!/usr/bin/env python3
"""Generate db/changelog.csv from the RPM/DEB release notes.

Reads content/release/{rpm,deb}.md (English) and their .zh.md twins (Chinese),
parses every per-batch markdown table into (date, package, old, new, comment)
rows, then:

1. aligns batch dates across the RPM and DEB sides (a batch published on
   slightly different days on each side gets one canonical date);
2. normalizes package names onto the catalog package name (pgext pkg column,
   via db/universe.csv): extension names map to their package, aliases and
   typos are fixed, and per-PG kernel builds (babelfish-17, pgedge-15, ...)
   keep their literal package names;
3. merges the RPM and DEB rows of one (date, package) into a single row,
   keeping the English comment in `note` and the Chinese one in `note_zh`.

The output db/changelog.csv feeds pgext.changelog (COPY ... CSV HEADER).
Run from the repository root:  python3 bin/changelog.py
"""

import csv
import re
import sys
from collections import OrderedDict
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
RELEASE = ROOT / "content" / "release"
UNIVERSE = ROOT / "db" / "universe.csv"
OUTPUT = ROOT / "db" / "changelog.csv"

# One canonical date per release batch: the RPM side of the 2025-09 batch went
# out on 09-04, the DEB side on 09-06 — both land on the later date.
DATE_ALIGN = {"2025-09-04": "2025-09-06"}

# Names that resolve to nothing in the catalog on purpose: kernel forks,
# runtime dependencies and tooling that have no extension package page.
# They keep their literal (cleaned) package name.
KNOWN_NON_CATALOG = {
    "agensgraph", "antlr4_runtime413", "babelfishpg", "cloudberry",
    "cloudberry_backup", "cloudberry_pxf", "oriolepg", "pdu", "pgdog",
    "pgedge", "polardb", "openhalodb",
}

# Kernel package families shipped as one package per PostgreSQL major —
# their per-PG names (babelfish-17, pgedge-15, ...) are the real package
# names and must not be collapsed into one row.
KERNEL_PER_PG = {
    "polardb", "agensgraph", "openhalodb", "babelfish", "pgedge",
    "orioledb", "ivorysql", "cloudberry",
}

# Spelling fixes and historical renames → canonical catalog lookup key.
ALIASES = {
    "timesacledb": "timescaledb",
    "timescaledb-toolkit": "timescaledb_toolkit",
    "sentinel": "pgsentinel",
    "pgsql-tweaks": "pgsql_tweaks",
    "pg-orphaned": "pg_orphaned",
    "pg_tokenizer.rs": "pg_tokenizer",
    "tzf-pg": "pg_tzf",
    "postgresql_anonymizer": "pg_anon",
    "postgresbson": "pgbson",
    "pg_safeupdate": "safeupdate",
    "pgddl": "ddlx",
}

EMPTY_VERSIONS = {"", "-", "–", "—", "n/a", "none", "*"}
LINK_RE = re.compile(r"\[([^\]]*)\]\([^)]*\)")
BATCH_RE = re.compile(r"^##\s+(\d{4}-\d{2}-\d{2})\s*$")
SEP_CELL_RE = re.compile(r"^:?-{2,}:?$")


def clean_cell(cell: str) -> str:
    cell = LINK_RE.sub(r"\1", cell).strip()
    return cell.strip("`").strip()


def clean_version(cell: str):
    cell = clean_cell(cell)
    return None if cell.lower() in EMPTY_VERSIONS else cell


def parse_tables(path: Path):
    """Yield (date, [ [cells...] ... ]) for every `## YYYY-MM-DD` batch."""
    batches = OrderedDict()
    date = None
    for line in path.read_text(encoding="utf-8").splitlines():
        m = BATCH_RE.match(line)
        if m:
            date = m.group(1)
            batches.setdefault(date, [])
            continue
        if date is None or not line.strip().startswith("|"):
            continue
        cells = [c.strip() for c in line.strip().strip("|").split("|")]
        if all(SEP_CELL_RE.match(c) for c in cells if c != ""):
            if batches[date]:
                batches[date].pop()  # the row before the separator is the header
            continue
        batches[date].append(cells)
    return batches


class Row:
    __slots__ = ("side", "raw", "old", "new", "note", "note_zh")

    def __init__(self, side, raw, old, new, note, note_zh=None):
        self.side, self.raw = side, raw
        self.old, self.new = old, new
        self.note, self.note_zh = note, note_zh


def parse_side(side: str, en_path: Path, zh_path: Path, warnings: list):
    """Parse one side (rpm or deb) and join English/Chinese comments."""
    en, zh = parse_tables(en_path), parse_tables(zh_path)
    batches_out = {}
    for date, rows in en.items():
        ds = DATE_ALIGN.get(date, date)
        zh_rows = zh.get(date, [])
        zh_by_name = {}
        for cells in zh_rows:
            if len(cells) < 4:
                continue
            key = clean_cell(cells[0]).lower()
            zh_by_name.setdefault(key, []).append(cells)
        out = batches_out.setdefault(ds, [])
        for cells in rows:
            if len(cells) < 4:
                continue
            raw = clean_cell(cells[0])
            if not raw:
                continue
            note = clean_cell(cells[3]) or None
            note_zh = None
            bucket = zh_by_name.get(raw.lower())
            if bucket:
                note_zh = clean_cell(bucket.pop(0)[3]) or None
            elif zh_rows:
                warnings.append(f"{en_path.name} {date}: no zh row for {raw!r}")
            out.append(Row(side, raw, clean_version(cells[1]), clean_version(cells[2]), note, note_zh))
    return batches_out


def load_catalog():
    name_to_pkg, pkgs, leads = {}, set(), {}
    with UNIVERSE.open(encoding="utf-8") as f:
        for r in csv.DictReader(f):
            name_to_pkg[r["name"].lower()] = r["pkg"]
            pkgs.add(r["pkg"])
            if r["name"] == r["lead_ext"]:
                leads.setdefault(r["pkg"].lower(), set()).add(r["name"].lower())
    return name_to_pkg, pkgs, leads


class Normalizer:
    def __init__(self):
        self.name_to_pkg, self.pkgs, self.leads = load_catalog()
        self.pkg_lower = {p.lower(): p for p in self.pkgs}
        self.unresolved = {}

    def resolve(self, raw: str):
        """Return the list of catalog package names for one raw table name."""
        key = ALIASES.get(raw.lower(), raw.lower())

        # wildcard rows: `omni_*` → the omni package; `hunspell_*` → every
        # hunspell_* package in the catalog
        if key.endswith("_*") or key.endswith("*"):
            base = key.rstrip("*").rstrip("_")
            hit = self.lookup(base)
            if hit:
                return [hit]
            members = sorted(p for p in self.pkgs if p.lower().startswith(base + "_"))
            if members:
                return members
            return [self.keep(raw, base)]

        # per-PG kernel builds keep their literal package name (babelfish-17)
        m = re.match(r"^([a-z0-9_]+?)-(\d{2})$", key)
        if m and m.group(1) in KERNEL_PER_PG:
            return [key]

        hit = self.lookup(key)
        if hit:
            return [hit]
        return [self.keep(raw, key)]

    def lookup(self, key: str):
        if key in self.name_to_pkg:  # extension name → its package
            return self.name_to_pkg[key]
        if key in self.pkg_lower:  # already a package name
            return self.pkg_lower[key]
        underscored = key.replace("-", "_")
        if underscored in self.name_to_pkg:
            return self.name_to_pkg[underscored]
        if underscored in self.pkg_lower:
            return self.pkg_lower[underscored]
        return None

    def keep(self, raw: str, key: str):
        pkg = key.replace("-", "_") if key.replace("-", "_") in KNOWN_NON_CATALOG else key
        if pkg not in KNOWN_NON_CATALOG:
            self.unresolved.setdefault(pkg, 0)
            self.unresolved[pkg] += 1
        return pkg


def join_notes(parts, zh=False):
    uniq = []
    for p in parts:
        if p and p not in uniq:
            uniq.append(p)
    return ("；" if zh else "; ").join(uniq) or None


def merge_group(pkg: str, rows: list, leads: set):
    """Merge every row that landed on one (ds, pkg) into a single record.

    Rows come from the same package family (lead + member extensions) and/or
    from both the RPM and DEB side. The best row supplies the version pair
    (package-named row first, then the lead extension, then any row with a
    complete old→new pair); remaining distinct comments are appended. When
    the RPM and DEB sides shipped genuinely different versions, the note gets
    an explicit `RPM x, DEB y` marker.
    """
    def prio(row: Row):
        raw = row.raw.lower().replace("-", "_")
        if raw == pkg.lower().replace("-", "_"):
            return 0
        if raw in leads:
            return 1
        if row.old and row.new:
            return 2
        if row.new:
            return 3
        return 4

    ordered = sorted(rows, key=prio)
    primary = ordered[0]
    old, new = primary.old, primary.new

    notes, notes_zh = [], []
    for r in ordered:
        # a bare "new" fragment from a member extension adds nothing once the
        # family already has a versioned primary row
        if r is not primary and r.note in ("new",) and primary.new:
            continue
        notes.append(r.note)
        notes_zh.append(r.note_zh)
    note = join_notes(notes)
    note_zh = join_notes(notes_zh, zh=True)

    # side-level disagreement (e.g. cloudberry RPM 2.1.0-3 vs DEB 2.1.0-2)
    by_side = {}
    for r in ordered:
        if r.new and r.side not in by_side:
            by_side[r.side] = r.new
    if len(by_side) == 2 and by_side["rpm"] != by_side["deb"]:
        marker = f"RPM {by_side['rpm']}, DEB {by_side['deb']}"
        if not note or (by_side["rpm"] not in note or by_side["deb"] not in note):
            note = join_notes([note, marker])
        marker_zh = f"RPM {by_side['rpm']}，DEB {by_side['deb']}"
        if not note_zh or (by_side["rpm"] not in note_zh or by_side["deb"] not in note_zh):
            note_zh = join_notes([note_zh, marker_zh], zh=True)
    return old, new, note, note_zh


def main():
    warnings = []
    rpm = parse_side("rpm", RELEASE / "rpm.md", RELEASE / "rpm.zh.md", warnings)
    deb = parse_side("deb", RELEASE / "deb.md", RELEASE / "deb.zh.md", warnings)

    norm = Normalizer()
    merged = {}  # (ds, pkg) → [Row]
    for side in (rpm, deb):
        for ds, rows in side.items():
            for row in rows:
                for pkg in norm.resolve(row.raw):
                    merged.setdefault((ds, pkg), []).append(row)

    records = []
    for (ds, pkg), rows in merged.items():
        # identical rows (RPM/DEB twins, duplicated family members) collapse
        uniq, seen = [], set()
        for r in rows:
            sig = (r.old, r.new, r.note, r.note_zh)
            if sig not in seen:
                seen.add(sig)
                uniq.append(r)
        old, new, note, note_zh = merge_group(pkg, uniq, norm.leads.get(pkg.lower(), set()))
        records.append((ds, pkg, old, new, note, note_zh))

    records.sort(key=lambda r: (r[0], r[1]), reverse=True)
    with OUTPUT.open("w", encoding="utf-8", newline="") as f:
        w = csv.writer(f)
        w.writerow(["ds", "pkg", "old_ver", "new_ver", "note_en", "note_zh"])
        for r in records:
            w.writerow(["" if v is None else v for v in r])

    dates = sorted({r[0] for r in records})
    print(f"wrote {len(records)} rows across {len(dates)} batches to {OUTPUT.relative_to(ROOT)}")
    print(f"batches: {dates[0]} .. {dates[-1]}")
    if norm.unresolved:
        print(f"\nkept {len(norm.unresolved)} names verbatim (not in catalog):")
        for name, count in sorted(norm.unresolved.items()):
            print(f"  {name} ×{count}")
    if warnings:
        print(f"\n{len(warnings)} warnings:")
        for wmsg in warnings[:30]:
            print(f"  {wmsg}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
