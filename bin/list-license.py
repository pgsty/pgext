#!/usr/bin/env python3
"""
PostgreSQL Extension License List Generator

Generates license-grouped extension lists in English and Chinese
by reading from the pgext database schema.

Usage:
    python list-license.py [PGURL] [OUTPUT_DIR]

    PGURL: PostgreSQL connection URL (default: postgres:///vonng)
    OUTPUT_DIR: Output directory for generated files (default: content/list/)
"""

import argparse
import logging
from collections import defaultdict
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List

import psycopg2

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(levelname)s: %(message)s')
logger = logging.getLogger(__name__)


@dataclass
class Extension:
    """Extension metadata"""
    id: int
    name: str
    pkg: str
    license: str
    en_desc: str
    zh_desc: str


@dataclass
class LicenseInfo:
    """License metadata"""
    name: str
    url: str
    description: str
    display_order: int


# License metadata mapping
LICENSE_INFO = {
    'PostgreSQL': LicenseInfo('PostgreSQL', 'https://opensource.org/licenses/postgresql',
                              'Very liberal license based on the BSD license, allowing almost unlimited freedom.', 1),
    'Apache-2.0': LicenseInfo('Apache-2.0', 'https://opensource.org/licenses/Apache-2.0',
                              'Permissive license with patent protection and attribution requirements.', 2),
    'MIT': LicenseInfo('MIT', 'https://opensource.org/licenses/MIT',
                       'A permissive license that allows commercial use, modification, and private use.', 3),
    'BSD 3-Clause': LicenseInfo('BSD 3-Clause', 'https://opensource.org/license/bsd-3-clause',
                                'Permissive license with attribution and endorsement restriction clauses.', 4),
    'BSD 2-Clause': LicenseInfo('BSD 2-Clause', 'https://opensource.org/license/bsd-2-clause',
                                'Permissive license requiring attribution but allowing commercial use.', 5),
    'GPL-2.0': LicenseInfo('GPL-2.0', 'https://opensource.org/licenses/GPL-2.0',
                           'Strong copyleft license requiring derivative works to be open source.', 6),
    'GPL-3.0': LicenseInfo('GPL-3.0', 'https://opensource.org/licenses/GPL-3.0',
                           'Strong copyleft license with additional patent and hardware restrictions.', 7),
    'AGPL-3.0': LicenseInfo('AGPL-3.0', 'https://opensource.org/licenses/AGPL-3.0',
                            'Network copyleft license extending GPL to cover network-distributed software.', 8),
    'ISC': LicenseInfo('ISC', 'https://opensource.org/licenses/ISC',
                       'A permissive license similar to MIT, allowing commercial use and modification.', 9),
    'Artistic': LicenseInfo('Artistic', 'https://opensource.org/license/artistic-2-0',
                            'Copyleft license allowing modification with certain distribution requirements.', 10),
    'Timescale': LicenseInfo('Timescale', 'https://www.timescale.com/legal/licenses',
                             'Proprietary license with restrictions on commercial use and distribution.', 11),
    'BSD 0-Clause': LicenseInfo('BSD 0-Clause', 'https://opensource.org/license/0bsd',
                                'Public domain equivalent license with no restrictions on use.', 12),
    'LGPL-3.0': LicenseInfo('LGPL-3.0', 'https://opensource.org/licenses/LGPL-3.0',
                            'Weak copyleft license with additional patent and hardware provisions.', 13),
    'MPL-2.0': LicenseInfo('MPL-2.0', 'https://opensource.org/licenses/MPL-2.0',
                           'Weak copyleft license allowing proprietary combinations with file-level copyleft.', 14),
    'LGPL-2.1': LicenseInfo('LGPL-2.1', 'https://opensource.org/licenses/LGPL-2.1',
                            'Weak copyleft license allowing proprietary applications to link dynamically.', 15),
}

# License groupings
PERMISSIVE_LICENSES = ['MIT', 'ISC', 'PostgreSQL', 'BSD 0-Clause', 'BSD 2-Clause', 'BSD 3-Clause',
                       'Artistic', 'Apache-2.0', 'MPL-2.0']
COPYLEFT_LICENSES = ['GPL-2.0', 'GPL-3.0', 'LGPL-2.1', 'LGPL-3.0', 'AGPL-3.0', 'Timescale']


def fetch_extensions(pgurl: str) -> Dict[str, List[Extension]]:
    """Fetch all extensions grouped by license"""
    with psycopg2.connect(pgurl) as conn, conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, pkg, license, en_desc, zh_desc
            FROM pgext.extension
            ORDER BY license, id
        """)

        extensions_by_license = defaultdict(list)
        for row in cur.fetchall():
            ext = Extension(*row)
            extensions_by_license[ext.license].append(ext)

        return extensions_by_license


def sanitize(text: str) -> str:
    """Sanitize text for markdown table"""
    if not text:
        return ""
    return text.replace('|', '\\|').replace('\n', ' ').strip()


def alias_shortcode(ext_name: str, pkg_name: str) -> str:
    """Generate alias shortcode for extension"""
    if ext_name == pkg_name:
        return f'{{{{< alias "{ext_name}" >}}}}'
    return f'{{{{< alias "{ext_name}" "{pkg_name}" >}}}}'


def license_badge_line(licenses: List[str]) -> str:
    """Generate a line of license badges"""
    return ' '.join(f'{{{{< license "{lic}" >}}}}' for lic in licenses)


def render_summary_table(sorted_licenses: List[tuple], is_zh: bool) -> List[str]:
    """Render summary table with license counts"""
    headers = ("许可证", "数量", "参考", "描述") if is_zh else ("License", "Count", "Reference", "Description")
    ref_text = "许可证文本" if is_zh else "License Text"

    lines = [
        f"| {headers[0]} | {headers[1]} | {headers[2]} | {headers[3]} |",
        "|:--------|:-----:|:-------:|:----------|"
    ]

    for license_name, exts in sorted_licenses:
        count = len(exts)
        info = LICENSE_INFO.get(license_name, LicenseInfo(license_name, '#', 'Unknown license', 999))
        lines.append(
            f'| {{{{< license "{license_name}" >}}}} | {count} | '
            f'[{ref_text}]({info.url}) | {info.description} |'
        )

    return lines


def render_license_section(license_name: str, exts: List[Extension], is_zh: bool) -> List[str]:
    """Render a single license section with extension table"""
    info = LICENSE_INFO.get(license_name, LicenseInfo(license_name, '#', 'Unknown license', 999))
    count = len(exts)

    # Section header
    badge_text = f"{count} 个扩展" if is_zh else f"{count} Extensions"
    ref_text = "许可证文本" if is_zh else "License Text"

    lines = [
        f"## {license_name}",
        "",
        f'{{{{< license "{license_name}" >}}}} {{{{< badge content="{badge_text}" color="gray" icon="cube" >}}}}',
        "",
        f"[{license_name} {ref_text}]({info.url}) : {info.description}",
        ""
    ]

    # Table header
    headers = ("ID", "扩展", "描述") if is_zh else ("ID", "Extension", "Description")
    lines.extend([
        f"| {headers[0]} | {headers[1]} | {headers[2]} |",
        "|:---:|:---|:---|"
    ])

    # Table rows
    for ext in exts:
        desc = sanitize(ext.zh_desc or ext.en_desc) if is_zh else sanitize(ext.en_desc or ext.zh_desc)
        lines.append(f"| {ext.id} | {alias_shortcode(ext.name, ext.pkg)} | {desc or '暂无描述' if is_zh else desc} |")

    lines.append("")
    return lines


def generate_content(extensions_by_license: Dict[str, List[Extension]], is_zh: bool) -> str:
    """Generate markdown content for license list"""
    # Sort licenses by count (descending) and then by display order
    sorted_licenses = sorted(
        extensions_by_license.items(),
        key=lambda x: (-len(x[1]), LICENSE_INFO.get(x[0], LicenseInfo(x[0], '', '', 999)).display_order)
    )

    lines = []

    # Front matter
    if is_zh:
        lines.extend(["---", "title: 按许可证", "description: 按开源许可证组织的 PostgreSQL 扩展", "weight: 300","---", ""])
    else:
        lines.extend(["---", "title: By License", "description: PostgreSQL extensions organized by open source license", "weight: 300", "---", ""])

    # License badges
    permissive = [lic for lic in PERMISSIVE_LICENSES if lic in extensions_by_license]
    copyleft = [lic for lic in COPYLEFT_LICENSES if lic in extensions_by_license]

    lines.extend([
        license_badge_line(permissive),
        "",
        license_badge_line(copyleft),
        ""
    ])

    # Summary section
    lines.append("## 概览" if is_zh else "## Summary")
    lines.append("")
    lines.extend(render_summary_table(sorted_licenses, is_zh))
    lines.extend(["", "---------", ""])

    # Individual license sections
    for license_name, exts in sorted_licenses:
        lines.extend(render_license_section(license_name, exts, is_zh))

    return '\n'.join(lines)


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(description='Generate license-grouped extension lists')
    parser.add_argument('pgurl', nargs='?', default='postgres:///vonng',
                        help='PostgreSQL connection URL (default: postgres:///vonng)')
    parser.add_argument('output_dir', nargs='?', default=None,
                        help='Output directory for generated files')
    parser.add_argument('--verbose', '-v', action='store_true',
                        help='Enable verbose logging')

    args = parser.parse_args()

    # Set output directory
    if args.output_dir is None:
        output_dir = Path(__file__).parent.parent / 'content' / 'list'
    else:
        output_dir = Path(args.output_dir)

    # Configure logging
    if args.verbose:
        logging.getLogger().setLevel(logging.DEBUG)

    # Generate files
    output_dir.mkdir(parents=True, exist_ok=True)

    logger.info("Connecting to database")
    extensions_by_license = fetch_extensions(args.pgurl)
    logger.info(f"Found {len(extensions_by_license)} licenses")

    # Generate English content
    en_content = generate_content(extensions_by_license, is_zh=False)
    en_path = output_dir / "license.md"
    en_path.write_text(en_content, encoding='utf-8')
    logger.info(f"Generated English license list: {en_path}")

    # Generate Chinese content
    zh_content = generate_content(extensions_by_license, is_zh=True)
    zh_path = output_dir / "license.zh.md"
    zh_path.write_text(zh_content, encoding='utf-8')
    logger.info(f"Generated Chinese license list: {zh_path}")


if __name__ == "__main__":
    main()
