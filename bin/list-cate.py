#!/usr/bin/env python3
"""
Generate PostgreSQL Extension Category List Pages

Usage:
    python list-cate.py [PGURL] [OUTPUT_DIR]

    PGURL: PostgreSQL connection URL (default: postgres:///)
    OUTPUT_DIR: Output directory for list files (default: content/list/)
"""

import argparse
import psycopg2
from pathlib import Path
from typing import List, Optional


# ============================================================
# Data Models
# ============================================================
class Category:
    """Category metadata"""
    def __init__(self, id: int, name: str, en_desc: str, zh_desc: str):
        self.id = id
        self.name = name
        self.en_desc = en_desc or ""
        self.zh_desc = zh_desc or ""


class Extension:
    """Extension summary for category lists"""
    def __init__(self, id: int, name: str, pkg: str, version: str,
                 en_desc: str, zh_desc: str, category: str, lead_ext: str):
        self.id = id
        self.name = name
        self.pkg = pkg
        self.version = version or "-"
        self.en_desc = en_desc or ""
        self.zh_desc = zh_desc or ""
        self.category = (category or "").upper()
        self.lead_ext = lead_ext or name


# ============================================================
# Database Operations
# ============================================================
def fetch_categories(conn) -> List[Category]:
    """Fetch all categories ordered by ID"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, en_desc, zh_desc
            FROM pgext.category
            ORDER BY id
        """)
        return [Category(*row) for row in cur.fetchall()]


def fetch_extensions(conn) -> List[Extension]:
    """Fetch extension summaries grouped by category"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT e.id, e.name, e.pkg, e.version, e.en_desc, e.zh_desc,
                   e.category, e.lead_ext
            FROM pgext.extension e
            LEFT JOIN pgext.category c ON upper(e.category) = upper(c.name)
            ORDER BY c.id NULLS LAST, e.id
        """)
        return [Extension(*row) for row in cur.fetchall()]


def count_packages(extensions: List[Extension]) -> int:
    """Count unique packages"""
    return len({ext.pkg for ext in extensions})


# ============================================================
# Shortcode Helper
# ============================================================
def alias_shortcode(name: str, label: Optional[str] = None) -> str:
    """Generate alias shortcode: {{< alias "name" "label" >}}"""
    if label and label != name:
        return f'{{{{< alias "{name}" "{label}" >}}}}'
    return f'{{{{< alias "{name}" >}}}}'


# ============================================================
# Content Generation
# ============================================================
def sanitize_text(text: str) -> str:
    """Sanitize text for markdown table"""
    if not text:
        return ""
    return text.replace('|', '\\|').replace('\n', ' ').strip()


def generate_frontmatter(locale: str, cat_count: int, ext_count: int, pkg_count: int) -> str:
    """Generate frontmatter"""
    if locale == "en":
        title = "By Category"
        summary = f"PostgreSQL Extensions ({ext_count} ext in {pkg_count} pkg) categorized into {cat_count} categories."
    else:  # zh
        title = "按分类"
        summary = f"PostgreSQL 扩展（{ext_count} ext / {pkg_count} pkg）归属 {cat_count} 个分类。"

    return f"""---
title: "{title}"
weight: 100
---

{summary}

"""


def generate_category_section(category: Category, extensions: List[Extension], locale: str) -> str:
    """Generate markdown section for one category"""
    lines = [f"## {category.name}", ""]

    # Add category description
    desc = category.zh_desc if locale == "zh" else category.en_desc
    desc = desc or (category.en_desc if locale == "zh" else category.zh_desc)
    if desc:
        # Remove redundant category name prefix
        prefix = f"{category.name}:"
        if desc.startswith(prefix):
            desc = desc[len(prefix):].lstrip()
        lines.append(sanitize_text(desc))
        lines.append("")

    # Table header
    if locale == "en":
        lines.extend([
            "| ID | Extension / Package | Version | Description |",
            "|:---:|:---|:---|:---|"
        ])
        empty_msg = "No extensions available in this category"
    else:  # zh
        lines.extend([
            "| ID | 扩展/包 | 版本 | 描述 |",
            "|:---:|:---|:---|:---|"
        ])
        empty_msg = "此分类暂无扩展"

    # Table rows
    if extensions:
        for ext in extensions:
            desc = sanitize_text(ext.zh_desc if locale == "zh" else ext.en_desc)
            desc = desc or sanitize_text(ext.en_desc if locale == "zh" else ext.zh_desc) or "-"

            # Merged Extension/Package column using alias shortcode
            ext_pkg_cell = alias_shortcode(ext.name, ext.pkg)

            lines.append(f"| {ext.id} | {ext_pkg_cell} | {ext.version} | {desc} |")
    else:
        lines.append(f"| - | - | - | {empty_msg} |")

    lines.append("")
    return '\n'.join(lines)


def generate_category_content(categories: List[Category],
                              extensions_by_cat: dict,
                              locale: str) -> str:
    """Generate complete category list content body"""
    sections = []
    for cat in categories:
        cat_key = cat.name.upper()
        exts = extensions_by_cat.get(cat_key, [])
        sections.append(generate_category_section(cat, exts, locale))

    return '\n'.join(sections)


# ============================================================
# Main Logic
# ============================================================
def generate_category_lists(pgurl: str, output_dir: Path):
    """Generate both English and Chinese category list pages"""
    # Connect to database
    conn = psycopg2.connect(pgurl)

    try:
        # Fetch data
        categories = fetch_categories(conn)
        extensions = fetch_extensions(conn)

        # Group extensions by category
        ext_by_cat = {}
        for ext in extensions:
            if ext.category:
                ext_by_cat.setdefault(ext.category, []).append(ext)

        # Calculate statistics
        cat_count = len(categories)
        ext_count = len(extensions)
        pkg_count = count_packages(extensions)

        # Generate English list
        en_content = generate_frontmatter("en", cat_count, ext_count, pkg_count)
        en_content += generate_category_content(categories, ext_by_cat, "en")

        # Generate Chinese list
        zh_content = generate_frontmatter("zh", cat_count, ext_count, pkg_count)
        zh_content += generate_category_content(categories, ext_by_cat, "zh")

        # Write files
        output_dir.mkdir(parents=True, exist_ok=True)
        (output_dir / "cate.md").write_text(en_content, encoding='utf-8')
        (output_dir / "cate.zh.md").write_text(zh_content, encoding='utf-8')

        print(f"✓ Generated: {output_dir}/cate.md")
        print(f"✓ Generated: {output_dir}/cate.zh.md")
        print(f"  Categories: {cat_count}, Extensions: {ext_count}, Packages: {pkg_count}")

    finally:
        conn.close()


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(
        description='Generate category list pages for Hugo'
    )
    parser.add_argument(
        'pgurl', nargs='?', default='postgres:///',
        help='PostgreSQL connection URL (default: postgres:///)'
    )
    parser.add_argument(
        'output_dir', nargs='?', default=None,
        help='Output directory (default: content/list/)'
    )

    args = parser.parse_args()

    # Determine output directory
    if args.output_dir:
        output_dir = Path(args.output_dir)
    else:
        output_dir = Path(__file__).parent.parent / 'content' / 'list'

    generate_category_lists(args.pgurl, output_dir)


if __name__ == "__main__":
    main()
