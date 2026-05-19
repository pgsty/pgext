(() => {
  const payload = window.PGEXT_UNIVERSE_DATA;

  if (!payload || !Array.isArray(payload.extensions)) {
    document.body.innerHTML = '<main class="stage"><p style="padding:32px">Missing pgext.cloud data.</p></main>';
    return;
  }

  const canvas = document.getElementById('galaxy');
  const stageEl = document.querySelector('.stage');
  const ctx = canvas.getContext('2d');
  const legendEl = document.getElementById('legend');
  const legendTitleEl = document.getElementById('legendTitle');
  const legendCountEl = document.getElementById('legendCount');
  const legendDimensionEl = document.getElementById('legendDimension');
  const tooltipEl = document.getElementById('tooltip');
  const linkMenuEl = document.getElementById('linkMenu');
  const searchInput = document.getElementById('searchInput');
  const dimensionListEl = document.getElementById('dimensionList');
  const dimensionTitleEl = document.querySelector('.dimension-list-head span');
  const dimensionCountEl = document.getElementById('dimensionCount');
  const clearGroupButton = document.getElementById('clearGroup');
  const toggleLabelsButton = document.getElementById('toggleLabels');
  const resetViewButton = document.getElementById('resetView');
  const quickFiltersEl = document.getElementById('quickFilters');
  const modeButtons = [...document.querySelectorAll('[data-entity-mode]')];
  const toggleLeftPanelButton = document.getElementById('toggleLeftPanel');

  const TAU = Math.PI * 2;
  const GOLDEN = Math.PI * (3 - Math.sqrt(5));
  const dimensions = payload.dimensions || [];
  const dimensionByKey = new Map(dimensions.map((dim) => [dim.key, dim]));
  const semanticColors = {
    positive: '#69a7ff',
    strongPositive: '#3de0ca',
    negative: '#ff6b8a',
    warning: '#f4b860',
    neutral: '#a9a2ff',
    muted: '#8a9ca0',
  };
  const basePalette = [
    '#3de0ca', '#f4b860', '#ff6b8a', '#7ca7ff', '#c7f464', '#ff9f66',
    '#b68cff', '#48d06d', '#ff77c8', '#76d7ff', '#e0d15c', '#ff595e',
    '#79ffa8', '#a9a2ff', '#f08a5d', '#5eead4', '#d9f99d', '#f472b6',
    '#67e8f9', '#fb7185', '#facc15', '#86efac', '#f0abfc', '#a8a29e',
  ];
  const dimensionPalettes = {
    category: basePalette,
    license: ['#3de0ca', '#7ca7ff', '#f4b860', '#ff6b8a', '#a3e635', '#c084fc', '#67e8f9', '#fb7185', '#facc15', '#86efac', '#f0abfc', '#a8a29e'],
    language: ['#7ca7ff', '#f4b860', '#3de0ca', '#ff6b8a', '#a3e635', '#c084fc', '#67e8f9', '#fb923c'],
  };
  const zoomLimits = { min: 0.72, max: 8.2 };
  const packageCoverageOrder = [
    '80 slots',
    ...Array.from({ length: 15 }, (_, index) => {
      const low = 75 - index * 5;
      return `${low}-${low + 4} slots`;
    }),
    '1-4 slots',
    '0 slots',
  ];
  const orderedValues = {
    extensionType: ['Standard Extension', 'Pure SQL Extension', 'Preload Extension', 'Headless / Metadata', 'Unknown Type'],
    catalogStatus: ['Catalog extension', 'Discovery candidate'],
    packageAvailability: ['RPM + DEB', 'RPM only', 'DEB only', 'No package'],
    pgCoverage: ['5 PG majors', '4 PG majors', '3 PG majors', '2 PG majors', '1 PG major', 'No PG major metadata'],
    starTier: ['T0: 10,000+ stars', 'T1: 1,000-9,999 stars', 'T2: 100-999 stars', 'T3: 10-99 stars', 'T4: 0-9 stars'],
    watchTier: ['T0: 100+ watchers', 'T1: 50-99 watchers', 'T2: 10-49 watchers', 'T3: 1-9 watchers', 'T4: 0 watchers'],
    forkTier: ['T0: 1,000+ forks', 'T1: 100-999 forks', 'T2: 10-99 forks', 'T3: 1-9 forks', 'T4: 0 forks'],
    lastUpdated: ['Within 1 year', '1-2 years', '2-3 years', '3-4 years', '4-5 years', '5-6 years', '6-7 years', '7-8 years', '8-9 years', '9-10 years', '10+ years', 'Unknown'],
    lastRelease: ['Within 1 year', '1-2 years', '2-3 years', '3-4 years', '4-5 years', '5-6 years', '6-7 years', '7-8 years', '8-9 years', '9-10 years', '10+ years', 'Unknown'],
    contrib: ['PostgreSQL contrib', 'External extension', 'Contrib unknown'],
    lead: ['Lead extension', 'Alias / sub-extension', 'Lead status unknown'],
    hasBin: ['Has client binary', 'No client binary', 'Binary flag unknown'],
    hasLib: ['Has shared library', 'No shared library', 'Library flag unknown'],
    needDdl: ['Needs DDL', 'No DDL needed', 'DDL flag unknown'],
    needLoad: ['Preload / LOAD required', 'No LOAD required', 'LOAD flag unknown'],
    trusted: ['Trusted', 'Superuser required', 'Trust unknown'],
    relocatable: ['Relocatable', 'Fixed or default schema', 'Relocatability unknown'],
    binaryMatrix: ['Full RPM + DEB matrix', 'Partial RPM + DEB matrix', 'RPM matrix only', 'DEB matrix only', 'No binary matrix'],
    packageCoverage: packageCoverageOrder,
    dependencyBand: ['No dependencies', '1 dependency', '2-3 dependencies', '4+ dependencies'],
    dependentBand: ['Dependency hub', 'Shared dependency', 'Lightly reused', 'Leaf extension'],
    schemaModel: ['Declared fixed schemas', 'Relocatable schema', 'Default extension schema'],
    pgxn: ['Listed on PGXN', 'Not on PGXN', 'PGXN unknown'],
    sourceArchive: ['Both', 'Rpm', 'Deb', 'None'],
  };
  const quickFilters = [
    { key: 'all', label: 'ALL', test: () => true },
    { key: 'wild', label: 'WILD', test: (ext) => ext.state !== 'available' },
    { key: 'pgext', label: 'PGEXT', test: (ext) => ext.state === 'available' },
    { key: 'pgsty', label: 'PGSTY', test: (ext) => ext.repo === 'PIGSTY' || ext.repo === 'MIXED' },
    { key: 'pgdg', label: 'PGDG', test: (ext) => ext.repo === 'PGDG' || ext.repo === 'MIXED' },
    { key: 'contrib', label: 'CONTRIB', test: (ext) => ext.contrib === true },
  ];

  const state = {
    dimension: 'category',
    entityMode: 'extension',
    filter: 'all',
    activeGroup: null,
    previewGroup: null,
    query: '',
    labelMode: 'dynamic',
    hover: null,
    linkMenuExt: null,
    zoom: 1,
    leftCollapsed: false,
    panX: 0,
    panY: 0,
    dragging: false,
    dragMoved: false,
    dragPointerId: null,
    dragStartX: 0,
    dragStartY: 0,
    dragStartPanX: 0,
    dragStartPanY: 0,
    width: 0,
    height: 0,
    dpr: 1,
    frame: 0,
  };

  const rawExtensions = payload.extensions || [];
  const allExtensions = rawExtensions.map((ext) => makeEntity(ext, 'extension'));
  const packageExtensions = buildPackageEntities(rawExtensions);
  let activeExtensions = [];

  let groups = [];
  let colorMaps = {};
  let groupCenters = new Map();
  let animationStarted = false;

  init();

  function init() {
    buildColorMaps();
    refreshActiveEntities();
    renderModeButtons();
    renderQuickFilters();
    renderDimensionList();
    bindControls();
    syncPanelCollapse();
    syncLabelButton();
    syncCanvasViewDataset();
    resize();
    requestAnimationFrame(tick);
  }

  function makeEntity(ext, entityMode, overrides = {}) {
    const isPackage = entityMode === 'package';
    const packageName = ext.pkg || ext.name || 'unknown-package';
    const extensionName = ext.name || packageName;
    const entity = {
      ...ext,
      ...overrides,
      entityMode,
      entityLabel: isPackage ? packageName : extensionName,
      extensionName,
      packageName,
    };
    entity.name = entity.entityLabel;
    entity.searchText = searchTextFor(entity);
    entity.x = 0;
    entity.y = 0;
    entity.tx = 0;
    entity.ty = 0;
    entity.radius = radiusForStars(entity.stars);
    entity.color = '#8aa0a4';
    entity.alpha = 1;
    return entity;
  }

  function buildPackageEntities(rows) {
    const byPackage = new Map();
    rows.filter((ext) => ext.lead === true).forEach((ext) => {
      const packageName = ext.pkg || ext.name;
      if (!packageName) return;
      const candidate = makeEntity(ext, 'package', { name: packageName });
      const current = byPackage.get(packageName);
      if (!current || packageCandidateScore(candidate) > packageCandidateScore(current)) {
        byPackage.set(packageName, candidate);
      }
    });
    return [...byPackage.values()].sort((a, b) => b.stars - a.stars || a.name.localeCompare(b.name));
  }

  function packageCandidateScore(ext) {
    return [
      ext.state === 'available' ? 10_000_000 : 0,
      ext.extensionName === ext.packageName ? 1_000_000 : 0,
      Number(ext.stars || 0),
      Number(ext.id || 0) / 1_000_000,
    ].reduce((sum, value) => sum + value, 0);
  }

  function searchTextFor(ext) {
    return [
      ext.name,
      ext.extensionName,
      ext.packageName,
      ext.pkg,
      ext.category,
      ext.license,
      ext.language,
      ext.repo,
      ext.state,
      ext.version,
      ext.extType,
      ext.extKernel,
      ext.extVendor,
      ext.enDesc,
      ...(ext.tags || []),
      ...(ext.requires || []),
      ...(ext.requireBy || []),
      ...Object.values(ext.dimensions || {}),
    ].filter(Boolean).join(' ').toLowerCase();
  }

  function renderDimensionList() {
    if (dimensionTitleEl) dimensionTitleEl.textContent = 'Dimensions';
    dimensionCountEl.textContent = String(dimensions.length);
    dimensionListEl.innerHTML = dimensions.map((dim) => {
      const valueCount = distinctValueCount(dim.key);
      const active = state.dimension === dim.key ? ' is-active' : '';
      const checked = state.dimension === dim.key ? 'true' : 'false';
      const tabIndex = state.dimension === dim.key ? '0' : '-1';
      return `<button class="dimension-option${active}" type="button" role="radio" aria-checked="${checked}" tabindex="${tabIndex}" data-dimension-key="${escapeAttr(dim.key)}" title="${escapeAttr(dim.description || dim.label)}">
        <span class="radio-dot"></span>
        <span class="dimension-name">${escapeHtml(dim.label)}</span>
        <span class="dimension-values">${valueCount}</span>
      </button>`;
    }).join('');

    dimensionListEl.querySelectorAll('.dimension-option').forEach((button) => {
      button.addEventListener('click', () => {
        setDimension(button.dataset.dimensionKey);
      });
      button.addEventListener('keydown', onDimensionOptionKeydown);
    });
  }

  function distinctValueCount(key) {
    const values = new Set(activeExtensions.map((ext) => valueFor(ext, key)));
    return values.size;
  }

  function renderModeButtons() {
    modeButtons.forEach((button) => {
      const active = button.dataset.entityMode === state.entityMode;
      const countEl = button.querySelector('.mode-count');
      button.classList.toggle('is-active', active);
      button.setAttribute('aria-pressed', active ? 'true' : 'false');
      if (countEl) countEl.textContent = formatNumber(countForMode(button.dataset.entityMode));
    });
    document.body.dataset.currentEntityMode = state.entityMode;
  }

  function renderQuickFilters() {
    if (!quickFiltersEl) return;
    quickFiltersEl.innerHTML = quickFilters.map((filter) => {
      const count = countForFilter(filter.key);
      const active = state.filter === filter.key ? ' is-active' : '';
      return `<button class="filter-button${active}" type="button" data-filter-key="${escapeAttr(filter.key)}" aria-pressed="${state.filter === filter.key ? 'true' : 'false'}">
        <span class="filter-label">${escapeHtml(filter.label)}</span>
        <span class="filter-count">${formatNumber(count)}</span>
      </button>`;
    }).join('');
  }

  function countForMode(mode) {
    return getBaseEntities(mode)
      .filter((ext) => passesQuickFilter(ext))
      .filter((ext) => passesQuery(ext))
      .length;
  }

  function countForFilter(filterKey) {
    return getBaseEntities()
      .filter((ext) => passesQuickFilter(ext, filterKey))
      .filter((ext) => passesQuery(ext))
      .length;
  }

  function onDimensionOptionKeydown(event) {
    const keyMoves = {
      ArrowDown: 1,
      ArrowUp: -1,
    };
    const delta = keyMoves[event.key];
    const isJumpKey = event.key === 'Home' || event.key === 'End';

    if (!delta && !isJumpKey) return;
    event.preventDefault();
    event.stopPropagation();

    if (!dimensions.length) return;

    const currentIndex = dimensions.findIndex((dim) => dim.key === state.dimension);
    let nextIndex = currentIndex < 0 ? 0 : currentIndex;

    if (event.key === 'Home') nextIndex = 0;
    else if (event.key === 'End') nextIndex = dimensions.length - 1;
    else nextIndex = (nextIndex + delta + dimensions.length) % dimensions.length;

    setDimension(dimensions[nextIndex].key, { focus: true });
  }

  function buildColorMaps() {
    colorMaps = {};
    for (const dim of dimensions) {
      const counts = new Map();
      allExtensions.forEach((ext) => counts.set(valueFor(ext, dim.key), (counts.get(valueFor(ext, dim.key)) || 0) + 1));
      const labels = [...counts.entries()].sort((a, b) => compareGroupLabels(dim.key, a, b)).map(([label]) => label);
      const palette = dimensionPalettes[dim.key] || basePalette;
      colorMaps[dim.key] = new Map(labels.map((label, index) => [label, semanticColorFor(dim.key, label) || palette[index % palette.length]]));
    }
  }

  function bindControls() {
    searchInput.addEventListener('input', () => {
      state.query = searchInput.value.trim().toLowerCase();
      state.activeGroup = null;
      state.previewGroup = null;
      hideLinkMenu();
      refreshView(true);
    });

    clearGroupButton.addEventListener('click', () => {
      setActiveGroup(null);
    });

    toggleLabelsButton.addEventListener('click', () => {
      const nextMode = { dynamic: 'all', all: 'none', none: 'dynamic' };
      state.labelMode = nextMode[state.labelMode] || 'dynamic';
      syncLabelButton();
    });

    resetViewButton?.addEventListener('click', resetZoom);

    toggleLeftPanelButton?.addEventListener('click', () => {
      state.leftCollapsed = !state.leftCollapsed;
      syncPanelCollapse();
      computeLayout(false);
    });

    modeButtons.forEach((button) => {
      button.addEventListener('click', () => {
        const mode = button.dataset.entityMode;
        if (!mode || mode === state.entityMode) return;
        state.entityMode = mode;
        state.activeGroup = null;
        state.previewGroup = null;
        state.hover = null;
        hideLinkMenu();
        resetZoom();
        refreshView(true);
      });
    });

    quickFiltersEl?.addEventListener('click', (event) => {
      const button = event.target.closest?.('[data-filter-key]');
      if (!button) return;
      const filter = button.dataset.filterKey;
      if (!filter || filter === state.filter) return;
      state.filter = filter;
      state.activeGroup = null;
      state.previewGroup = null;
      state.hover = null;
      hideLinkMenu();
      refreshView(true);
    });

    canvas.addEventListener('pointerdown', onPointerDown);
    canvas.addEventListener('pointermove', onPointerMove);
    canvas.addEventListener('pointerup', onPointerUp);
    canvas.addEventListener('pointercancel', onPointerUp);
    canvas.addEventListener('mouseleave', () => {
      if (state.dragging) onPointerUp();
      state.hover = null;
      hideTooltip();
    });
    canvas.addEventListener('click', () => {
      if (state.dragMoved) {
        state.dragMoved = false;
        return;
      }
      if (state.hover) openLinkMenu(state.hover);
      else hideLinkMenu();
    });
    canvas.addEventListener('wheel', onWheel, { passive: false });
    canvas.addEventListener('dblclick', resetZoom);
    legendEl.addEventListener('mouseleave', () => {
      state.previewGroup = null;
      syncLegendPreview();
    });
    legendEl.addEventListener('mousemove', onLegendPointer);
    legendEl.addEventListener('pointerover', onLegendPointer);
    linkMenuEl?.addEventListener('pointerdown', (event) => {
      event.stopPropagation();
    });
    linkMenuEl?.addEventListener('click', (event) => {
      if (event.target.closest?.('a[href]')) window.setTimeout(hideLinkMenu, 80);
    });
    document.addEventListener('pointerdown', (event) => {
      if (!state.linkMenuExt) return;
      if (event.target === canvas || linkMenuEl?.contains(event.target)) return;
      hideLinkMenu();
    });
    window.addEventListener('keydown', onGlobalKeydown);
    window.addEventListener('resize', resize);
  }

  function syncLabelButton() {
    if (!toggleLabelsButton) return;
    const title = {
      dynamic: 'Labels: dynamic',
      all: 'Labels: all',
      none: 'Labels: off',
    }[state.labelMode] || 'Labels: dynamic';
    toggleLabelsButton.classList.toggle('is-dynamic', state.labelMode === 'dynamic');
    toggleLabelsButton.classList.toggle('is-all', state.labelMode === 'all');
    toggleLabelsButton.classList.toggle('is-none', state.labelMode === 'none');
    toggleLabelsButton.setAttribute('aria-pressed', state.labelMode === 'dynamic' ? 'mixed' : state.labelMode === 'all' ? 'true' : 'false');
    toggleLabelsButton.setAttribute('title', title);
    toggleLabelsButton.setAttribute('aria-label', title);
    document.body.dataset.labelMode = state.labelMode;
    syncCanvasViewDataset();
  }

  function syncCanvasViewDataset() {
    canvas.dataset.zoom = state.zoom.toFixed(2);
    canvas.dataset.pan = `${Math.round(state.panX)},${Math.round(state.panY)}`;
    canvas.dataset.labelMode = state.labelMode;
  }

  function syncPanelCollapse() {
    stageEl.classList.toggle('is-left-collapsed', state.leftCollapsed);
    if (toggleLeftPanelButton) {
      toggleLeftPanelButton.textContent = state.leftCollapsed ? '›' : '‹';
      toggleLeftPanelButton.setAttribute('aria-expanded', state.leftCollapsed ? 'false' : 'true');
      toggleLeftPanelButton.setAttribute('aria-label', state.leftCollapsed ? 'Expand dimensions' : 'Collapse dimensions');
    }
    clampPan();
  }

  function onGlobalKeydown(event) {
    if (event.key === 'Escape' && state.linkMenuExt) {
      event.preventDefault();
      hideLinkMenu();
      return;
    }
    if (!['ArrowLeft', 'ArrowRight'].includes(event.key)) return;
    if (event.repeat) return;
    const target = event.target;
    if (target?.closest?.('input, textarea, select, [contenteditable="true"]')) return;
    event.preventDefault();
    cycleActiveGroup(event.key === 'ArrowRight' ? 1 : -1);
  }

  function cycleActiveGroup(delta) {
    const labels = groups.map((group) => group.label);
    if (!labels.length) {
      setActiveGroup(null);
      return;
    }
    const sequence = [null, ...labels];
    const currentIndex = sequence.findIndex((label) => label === state.activeGroup);
    const start = currentIndex < 0 ? 0 : currentIndex;
    const nextIndex = (start + delta + sequence.length) % sequence.length;
    setActiveGroup(sequence[nextIndex], { scrollLegend: true });
  }

  function setActiveGroup(nextGroup, options = {}) {
    state.activeGroup = nextGroup;
    state.previewGroup = null;
    renderLegend();
    if (options.scrollLegend && nextGroup) {
      const node = legendEl.querySelector(`[data-label="${cssEscape(nextGroup)}"]`);
      node?.scrollIntoView({ block: 'nearest', inline: 'nearest' });
    }
  }

  function refreshView(snap = false) {
    refreshActiveEntities();
    renderModeButtons();
    renderQuickFilters();
    renderDimensionList();
    computeLayout(snap);
    renderLegend();
  }

  function refreshActiveEntities() {
    activeExtensions = getBaseEntities()
      .filter((ext) => passesQuickFilter(ext, state.filter))
      .filter((ext) => passesQuery(ext));
    if (state.linkMenuExt && !activeExtensions.includes(state.linkMenuExt)) hideLinkMenu();
  }

  function getBaseEntities(mode = state.entityMode) {
    return mode === 'package' ? packageExtensions : allExtensions;
  }

  function passesQuickFilter(ext, filterKey = state.filter) {
    const filter = quickFilters.find((item) => item.key === filterKey) || quickFilters[0];
    return filter.test(ext);
  }

  function passesQuery(ext) {
    return !state.query || ext.searchText.includes(state.query);
  }

  function setDimension(nextDimension, options = {}) {
    if (!dimensionByKey.has(nextDimension)) return;
    state.dimension = nextDimension;
    state.activeGroup = null;
    state.previewGroup = null;
    hideLinkMenu();
    let activeNode = null;
    dimensionListEl.querySelectorAll('.dimension-option').forEach((node) => {
      const active = node.dataset.dimensionKey === nextDimension;
      node.classList.toggle('is-active', active);
      node.setAttribute('aria-checked', active ? 'true' : 'false');
      node.tabIndex = active ? 0 : -1;
      if (active) activeNode = node;
    });
    computeLayout(true);
    renderLegend();

    if (options.focus && activeNode) {
      activeNode.focus({ preventScroll: true });
      activeNode.scrollIntoView({ block: 'nearest', inline: 'nearest' });
    }
  }

  function resize() {
    const rect = canvas.getBoundingClientRect();
    state.width = Math.max(1, rect.width);
    state.height = Math.max(1, rect.height);
    state.dpr = Math.min(window.devicePixelRatio || 1, 2);
    canvas.width = Math.round(state.width * state.dpr);
    canvas.height = Math.round(state.height * state.dpr);
    ctx.setTransform(state.dpr, 0, 0, state.dpr, 0, 0);
    clampPan();
    refreshActiveEntities();
    computeLayout(!animationStarted);
    renderLegend();
  }

  function computeLayout(snap = false) {
    const byGroup = new Map();
    for (const ext of activeExtensions) {
      const key = valueFor(ext, state.dimension);
      if (!byGroup.has(key)) byGroup.set(key, []);
      byGroup.get(key).push(ext);
    }

    groups = [...byGroup.entries()].map(([label, items]) => {
      items.sort((a, b) => b.stars - a.stars || a.name.localeCompare(b.name));
      return {
        label,
        items,
        count: items.length,
        stars: items.reduce((sum, item) => sum + item.stars, 0),
        color: colorMaps[state.dimension].get(label) || '#8aa0a4',
      };
    }).sort((a, b) => compareGroupLabels(state.dimension, [a.label, a.count], [b.label, b.count]));

    if (state.activeGroup && !groups.some((group) => group.label === state.activeGroup)) {
      state.activeGroup = null;
    }

    const bounds = sceneBounds();
    const cx = bounds.cx;
    const cy = bounds.cy;
    const armCount = groups.length <= 6 ? 3 : 4;
    const ringCount = Math.max(1, Math.ceil(groups.length / armCount));
    const maxOrbit = Math.min(bounds.width * 0.42, bounds.height * 0.42);
    const minOrbit = Math.min(bounds.width, bounds.height) * 0.055;
    const maxCount = Math.max(...groups.map((group) => group.count), 1);
    groupCenters = new Map();

    groups.forEach((group, index) => {
      const arm = index % armCount;
      const ring = Math.floor(index / armCount);
      const progress = (ring + 0.55) / (ringCount + 0.25);
      const theta = arm * (TAU / armCount) + progress * 2.55 + 0.22;
      const orbit = minOrbit + Math.pow(progress, 0.80) * maxOrbit;
      const center = {
        x: cx + Math.cos(theta) * orbit * 1.16,
        y: cy + Math.sin(theta) * orbit * 0.73,
        theta,
        orbit,
        radius: 36 + Math.sqrt(group.count / maxCount) * Math.min(state.width, state.height) * 0.145,
      };
      groupCenters.set(group.label, center);

      group.items.forEach((ext, itemIndex) => {
        const seed = hash(ext.name);
        const localTheta = itemIndex * GOLDEN + seed * TAU;
        const localProgress = group.items.length <= 1 ? 0 : Math.sqrt(itemIndex / (group.items.length - 1));
        const feather = (hash(`${ext.name}:feather`) - 0.5) * 15;
        const spread = itemIndex === 0 ? 0 : center.radius * (0.08 + localProgress * 0.92) + feather;
        const armBias = Math.sin(localTheta * 0.5 + center.theta) * 7;
        ext.tx = center.x + Math.cos(localTheta) * spread * 1.12 + Math.cos(center.theta) * armBias;
        ext.ty = center.y + Math.sin(localTheta) * spread * 0.74 + Math.sin(center.theta) * armBias;
        ext.color = group.color;

        if (snap || !animationStarted || !ext.x || !ext.y) {
          ext.x = ext.tx;
          ext.y = ext.ty;
        }
      });
    });
  }

  function renderLegend() {
    const dim = dimensionByKey.get(state.dimension);
    legendTitleEl.textContent = dim?.label || state.dimension;
    if (legendDimensionEl) legendDimensionEl.textContent = `${formatNumber(activeExtensions.length)} ${state.entityMode === 'package' ? 'packages' : 'extensions'}`;
    const visibleGroups = groups;
    if (legendCountEl) legendCountEl.textContent = String(visibleGroups.length);
    clearGroupButton?.classList.toggle('is-active', !state.activeGroup);
    clearGroupButton?.setAttribute('aria-pressed', !state.activeGroup ? 'true' : 'false');

    legendEl.innerHTML = visibleGroups.map((group) => {
      const active = state.activeGroup === group.label ? ' is-active' : '';
      const preview = state.previewGroup === group.label ? ' is-preview' : '';
      const count = group.count;
      return `<button class="legend-item${active}${preview}" type="button" data-label="${escapeAttr(group.label)}">
        <span class="swatch" style="color:${group.color};background:${group.color}"></span>
        <span class="legend-name">${escapeHtml(displayValueFor(state.dimension, group.label))}</span>
        <span class="legend-count">${formatNumber(count)}</span>
      </button>`;
    }).join('');

    legendEl.querySelectorAll('.legend-item').forEach((button) => {
      button.addEventListener('click', () => {
        setActiveGroup(state.activeGroup === button.dataset.label ? null : button.dataset.label);
      });
      button.addEventListener('mouseenter', () => {
        state.previewGroup = button.dataset.label;
        syncLegendPreview();
      });
    });
  }

  function syncLegendPreview() {
    legendEl.querySelectorAll('.legend-item').forEach((button) => {
      button.classList.toggle('is-preview', button.dataset.label === state.previewGroup);
    });
  }

  function onLegendPointer(event) {
    const item = event.target.closest?.('.legend-item');
    if (!item || item.dataset.label === state.previewGroup) return;
    state.previewGroup = item.dataset.label;
    syncLegendPreview();
  }

  function tick() {
    animationStarted = true;
    state.frame += 1;
    updatePositions();
    draw();
    requestAnimationFrame(tick);
  }

  function updatePositions() {
    for (const ext of activeExtensions) {
      ext.x += (ext.tx - ext.x) * 0.075;
      ext.y += (ext.ty - ext.y) * 0.075;
      ext.alpha += (targetAlpha(ext) - ext.alpha) * 0.12;
    }
  }

  function draw() {
    ctx.clearRect(0, 0, state.width, state.height);
    ctx.save();
    drawGroupLabels();
    drawStars();
    drawLabels();
    ctx.restore();
    if (state.linkMenuExt) {
      hideTooltip();
      positionLinkMenu();
    } else if (state.hover) {
      showTooltip(state.hover);
    }
  }

  function drawGroupLabels() {
    ctx.save();
    ctx.font = '740 12px Inter, system-ui, sans-serif';
    ctx.textBaseline = 'middle';
    groups.forEach((group) => {
      const center = groupCenters.get(group.label);
      if (!center) return;
      const point = worldToScreen(center.x, center.y);
      const focusGroup = state.previewGroup || state.activeGroup;
      const alpha = focusGroup && focusGroup !== group.label ? 0.18 : 0.82;
      const text = displayValueFor(state.dimension, group.label);
      const x = point.x - ctx.measureText(text).width / 2;
      const y = point.y - Math.min(58, center.radius * 0.48) - 12;
      ctx.globalAlpha = alpha;
      ctx.shadowColor = 'rgba(0,0,0,0.72)';
      ctx.shadowBlur = 8;
      ctx.lineWidth = 3;
      ctx.strokeStyle = 'rgba(3, 6, 9, 0.72)';
      ctx.strokeText(text, x, y);
      ctx.shadowBlur = 0;
      ctx.fillStyle = group.color;
      ctx.fillText(text, x, y);
    });
    ctx.restore();
  }

  function drawStars() {
    const drawOrder = [...activeExtensions].sort((a, b) => a.radius - b.radius);
    for (const ext of drawOrder) {
      const isHover = state.hover === ext;
      const isSelected = state.linkMenuExt === ext;
      const isFocus = isHover || isSelected;
      const alpha = Math.max(0.05, Math.min(1, ext.alpha));
      const r = ext.radius * (isFocus ? 1.35 : 1);
      const point = worldToScreen(ext.x, ext.y);
      ctx.save();
      ctx.globalAlpha = alpha;
      ctx.shadowColor = colorWithAlpha(ext.color, isFocus ? 0.95 : 0.48);
      ctx.shadowBlur = isFocus ? 28 : Math.max(5, r * 0.9);
      ctx.fillStyle = colorWithAlpha(ext.color, isFocus ? 0.98 : 0.80);
      ctx.beginPath();
      ctx.arc(point.x, point.y, r, 0, TAU);
      ctx.fill();
      ctx.shadowBlur = 0;
      ctx.lineWidth = isFocus ? 2 : 1;
      ctx.strokeStyle = isFocus ? 'rgba(255,255,255,0.88)' : 'rgba(255,255,255,0.24)';
      ctx.stroke();

      if (ext.needLoad || ext.hasLib) {
        ctx.strokeStyle = ext.needLoad ? 'rgba(255,255,255,0.78)' : 'rgba(255,255,255,0.42)';
        ctx.lineWidth = 1;
        ctx.beginPath();
        ctx.arc(point.x, point.y, r + 2.3, 0, TAU);
        ctx.stroke();
      }

      if (ext.trusted) {
        ctx.fillStyle = 'rgba(255,255,255,0.84)';
        ctx.beginPath();
        ctx.arc(point.x + r * 0.36, point.y - r * 0.36, Math.max(1.2, r * 0.15), 0, TAU);
        ctx.fill();
      }
      ctx.restore();
    }
  }

  function drawLabels() {
    if (state.labelMode === 'none') return;

    const query = state.query;
    const allLabels = state.labelMode === 'all';
    const dynamicRankLimit = Math.min(
      activeExtensions.length,
      Math.round(44 + Math.pow(state.zoom, 1.42) * 55),
    );
    const dynamicStarFloor = state.zoom >= 6.8 ? 0
      : state.zoom >= 5.2 ? 1
        : state.zoom >= 4.0 ? 5
          : state.zoom >= 2.8 ? 10
            : state.zoom >= 1.8 ? 80
              : 700;
    const candidates = activeExtensions
      .filter((ext) => ext.alpha > (allLabels ? 0.16 : 0.4))
      .filter((ext) => allLabels || ext.stars >= dynamicStarFloor || ext.starRank <= dynamicRankLimit || ext === state.hover || ext === state.linkMenuExt || (query && ext.searchText.includes(query)))
      .sort((a, b) => {
        if (a === state.linkMenuExt) return -1;
        if (b === state.linkMenuExt) return 1;
        if (a === state.hover) return -1;
        if (b === state.hover) return 1;
        return b.stars - a.stars;
      });

    const boxes = [];
    const bounds = sceneBounds();
    ctx.save();
    ctx.font = '700 11px Inter, system-ui, sans-serif';
    ctx.textBaseline = 'middle';

    for (const ext of candidates) {
      const point = worldToScreen(ext.x, ext.y);
      const text = ext.name.length > 24 ? `${ext.name.slice(0, 22)}...` : ext.name;
      const width = Math.ceil(ctx.measureText(text).width) + 16;
      const height = 22;
      let x = point.x + ext.radius + 7;
      let y = point.y - height / 2;
      if (x + width > bounds.right - 8) x = point.x - ext.radius - width - 7;
      if (x < bounds.left + 8) x = bounds.left + 8;
      if (y < bounds.top + 6) y = bounds.top + 6;
      if (y + height > bounds.top + bounds.height - 12) y = bounds.top + bounds.height - 12 - height;
      const box = { x, y, w: width, h: height };
      const overlaps = boxes.some((used) => intersects(box, used));
      if (!allLabels && overlaps && ext !== state.hover && ext !== state.linkMenuExt && !(query && ext.searchText.includes(query))) continue;
      boxes.push(box);

      ctx.globalAlpha = ext === state.hover || ext === state.linkMenuExt ? 1 : Math.min(allLabels ? 0.76 : 0.9, ext.alpha);
      ctx.fillStyle = 'rgba(6, 9, 11, 0.82)';
      roundRect(ctx, x, y, width, height, 11);
      ctx.fill();
      ctx.strokeStyle = colorWithAlpha(ext.color, ext === state.hover || ext === state.linkMenuExt ? 0.80 : 0.42);
      ctx.stroke();
      ctx.fillStyle = '#f0f7f6';
      ctx.fillText(text, x + 8, y + height / 2);
    }
    ctx.restore();
  }

  function onPointerDown(event) {
    if (event.button !== 0) return;
    state.dragging = true;
    state.dragMoved = false;
    state.dragPointerId = event.pointerId;
    state.dragStartX = event.clientX;
    state.dragStartY = event.clientY;
    state.dragStartPanX = state.panX;
    state.dragStartPanY = state.panY;
    canvas.setPointerCapture?.(event.pointerId);
    canvas.style.cursor = 'grabbing';
  }

  function onPointerUp(event = null) {
    if (!state.dragging) return;
    if (event?.pointerId != null && event.pointerId !== state.dragPointerId) return;
    canvas.releasePointerCapture?.(state.dragPointerId);
    state.dragging = false;
    state.dragPointerId = null;
    canvas.style.cursor = state.hover ? 'pointer' : 'grab';
  }

  function onPointerMove(event) {
    if (state.dragging) {
      const dx = event.clientX - state.dragStartX;
      const dy = event.clientY - state.dragStartY;
      if (Math.abs(dx) + Math.abs(dy) > 3) state.dragMoved = true;
      state.panX = state.dragStartPanX + dx;
      state.panY = state.dragStartPanY + dy;
      clampPan();
      syncCanvasViewDataset();
      state.hover = null;
      hideTooltip();
      if (state.dragMoved) hideLinkMenu();
      canvas.style.cursor = 'grabbing';
      return;
    }

    const rect = canvas.getBoundingClientRect();
    const screenX = event.clientX - rect.left;
    const screenY = event.clientY - rect.top;
    const pointer = screenToWorld(screenX, screenY);
    let nearest = null;
    let best = Infinity;

    for (const ext of activeExtensions) {
      if (ext.alpha < 0.16) continue;
      const dx = pointer.x - ext.x;
      const dy = pointer.y - ext.y;
      const limit = (ext.radius + 6) / state.zoom;
      const dist = dx * dx + dy * dy;
      if (dist <= limit * limit && dist < best) {
        nearest = ext;
        best = dist;
      }
    }

    state.hover = nearest;
    canvas.style.cursor = nearest ? 'pointer' : 'grab';
    if (!nearest) hideTooltip();
  }

  function showTooltip(ext) {
    const stageRect = canvas.parentElement.getBoundingClientRect();
    const point = worldToScreen(ext.x, ext.y);
    const x = Math.min(stageRect.width - 386, Math.max(18, point.x + 18));
    const y = Math.min(stageRect.height - 332, Math.max(86, point.y + 18));
    tooltipEl.style.left = `${x}px`;
    tooltipEl.style.top = `${y}px`;
    tooltipEl.classList.add('is-visible');

    const dim = dimensionByKey.get(state.dimension);
    const packageLine = ext.pkg && ext.pkg !== ext.name ? `Package: ${ext.pkg}` : 'Package name matches extension';
    const desc = ext.enDesc || packageLine;
    const vendor = ext.extVendor || 'none';
    const kernel = ext.extKernel || 'Core PostgreSQL';
    const entityLabel = state.entityMode === 'package' ? 'Package' : 'Extension';
    const dates = [
      ['Updated', formatDate(ext.lastUpdateDate)],
      ['Released', formatDate(ext.lastReleaseDate)],
      ['Commit', formatDate(ext.lastCommitDate)],
    ];
    tooltipEl.innerHTML = `
      <h3>${escapeHtml(ext.name)}</h3>
      <p class="desc">${escapeHtml(desc)}</p>
      <dl>
        <dt>Mode</dt><dd>${escapeHtml(entityLabel)}</dd>
        <dt>Extension</dt><dd>${escapeHtml(ext.extensionName || ext.name)}</dd>
        <dt>Package</dt><dd>${escapeHtml(ext.pkg || 'unknown')}</dd>
        <dt>Stars</dt><dd>${formatNumber(ext.stars)} - rank ${ext.starRank}</dd>
        <dt>Watch / Fork</dt><dd>${formatNumber(ext.watchCnt)} / ${formatNumber(ext.forkCnt)}</dd>
        <dt>${escapeHtml(dim?.label || 'Dimension')}</dt><dd>${escapeHtml(displayValueFor(state.dimension, valueFor(ext, state.dimension)))}</dd>
        <dt>Category</dt><dd>${escapeHtml(ext.category)}</dd>
        <dt>Type</dt><dd>${escapeHtml(ext.extType || 'Unknown')}</dd>
        <dt>License</dt><dd>${escapeHtml(ext.license)}</dd>
        <dt>Language</dt><dd>${escapeHtml(ext.language)}</dd>
        <dt>Repo / State</dt><dd>${escapeHtml(`${ext.repo || 'n/a'} / ${ext.state || 'n/a'}`)}</dd>
        <dt>Vendor</dt><dd>${escapeHtml(vendor)}</dd>
        <dt>Kernel</dt><dd>${escapeHtml(kernel)}</dd>
        <dt>Binary</dt><dd>${escapeHtml(packageAvailability(ext))}</dd>
        <dt>Coverage</dt><dd>${escapeHtml(`${formatNumber(ext.packageSlots || 0)} / 80 slots`)}</dd>
        <dt>Build</dt><dd>${escapeHtml(displayValueFor('sourceArchive', valueFor(ext, 'sourceArchive')))}</dd>
        <dt>PG</dt><dd>${escapeHtml((ext.pgVer || []).join(', ') || 'unknown')}</dd>
        ${dates.map(([label, value]) => `<dt>${label}</dt><dd>${escapeHtml(value)}</dd>`).join('')}
        <dt>Flags</dt><dd>${escapeHtml(flagsFor(ext))}</dd>
      </dl>
    `;
  }

  function hideTooltip() {
    tooltipEl.classList.remove('is-visible');
  }

  function openLinkMenu(ext) {
    if (!linkMenuEl) return;
    state.linkMenuExt = ext;
    hideTooltip();
    renderLinkMenu(ext);
    linkMenuEl.classList.add('is-visible');
    linkMenuEl.setAttribute('aria-hidden', 'false');
    positionLinkMenu();
  }

  function hideLinkMenu() {
    if (!linkMenuEl) return;
    state.linkMenuExt = null;
    linkMenuEl.classList.remove('is-visible', 'is-above');
    linkMenuEl.setAttribute('aria-hidden', 'true');
  }

  function renderLinkMenu(ext) {
    const { primaryLinks, quickLinks } = extensionLinks(ext);
    const kind = state.entityMode === 'package' ? 'Package links' : 'Extension links';
    linkMenuEl.innerHTML = `
      <div class="link-menu-head">
        <div class="link-menu-title">${escapeHtml(ext.name)}</div>
        <div class="link-menu-kind">${escapeHtml(kind)}</div>
      </div>
      <div class="link-menu-list">
        ${primaryLinks.map((link) => `<a class="link-menu-item" role="menuitem" href="${escapeAttr(link.url)}" target="_blank" rel="noopener noreferrer" title="${escapeAttr(`${link.label}: ${link.url}`)}">
          <span class="link-menu-label">${escapeHtml(link.label)}:</span>
          <span class="link-menu-url">${escapeHtml(link.url)}</span>
        </a>`).join('')}
        ${quickLinks.length ? `<div class="link-menu-actions" aria-label="Pigsty section links">
          ${quickLinks.map((link) => `<a class="link-menu-action" role="menuitem" href="${escapeAttr(link.url)}" target="_blank" rel="noopener noreferrer" title="${escapeAttr(link.url)}">${escapeHtml(link.label)}</a>`).join('')}
        </div>` : ''}
      </div>
    `;
  }

  function extensionLinks(ext) {
    const slug = extensionSlug(ext);
    const pgextUrl = `https://pgext.cloud/e/${slug}/`;
    const pigstyIo = `https://pigsty.io/ext/e/${slug}/`;
    const url = normalizeLinkUrl(ext.url);
    const isAvailable = ext.state === 'available';
    const primaryLinks = [
      { label: 'URL', url: url || (isAvailable ? pgextUrl : '#') },
    ];
    const quickLinks = [];

    if (isAvailable) {
      primaryLinks.push(
        { label: 'PGEXT.CLOUD', url: pgextUrl },
        { label: 'Pigsty.IO', url: pigstyIo },
        { label: 'Pigsty.CC', url: `https://pigsty.cc/ext/e/${slug}/` },
      );
      quickLinks.push(
      { label: 'Info', url: `${pigstyIo}#overview` },
      { label: 'Matrix', url: `${pigstyIo}#version` },
      { label: 'Build', url: `${pigstyIo}#build` },
      { label: 'Install', url: `${pigstyIo}#install` },
      { label: 'Usage', url: `${pigstyIo}#usage` },
      );
    }

    return { primaryLinks, quickLinks };
  }

  function extensionSlug(ext) {
    const name = ext.extensionName || ext.name || ext.pkg || 'unknown';
    return encodeURIComponent(String(name).trim());
  }

  function normalizeLinkUrl(value) {
    const text = String(value || '').trim();
    return /^https?:\/\//i.test(text) ? text : '';
  }

  function positionLinkMenu() {
    if (!linkMenuEl || !state.linkMenuExt || !linkMenuEl.classList.contains('is-visible')) return;
    const stageRect = stageEl.getBoundingClientRect();
    const point = worldToScreen(state.linkMenuExt.x, state.linkMenuExt.y);
    const menuWidth = linkMenuEl.offsetWidth || 560;
    const menuHeight = linkMenuEl.offsetHeight || 320;
    const margin = 18;
    let x = point.x - 42;
    x = clamp(x, margin, stageRect.width - menuWidth - margin);
    let y = point.y + state.linkMenuExt.radius + 18;
    let isAbove = false;
    if (y + menuHeight > stageRect.height - margin) {
      y = point.y - state.linkMenuExt.radius - menuHeight - 18;
      isAbove = true;
    }
    y = clamp(y, 84, stageRect.height - menuHeight - margin);
    const arrowX = clamp(point.x - x - 6, 18, menuWidth - 26);
    linkMenuEl.style.left = `${x}px`;
    linkMenuEl.style.top = `${y}px`;
    linkMenuEl.style.setProperty('--menu-arrow-x', `${arrowX}px`);
    linkMenuEl.classList.toggle('is-above', isAbove);
  }

  function onWheel(event) {
    event.preventDefault();
    const rect = canvas.getBoundingClientRect();
    const mouseX = event.clientX - rect.left;
    const mouseY = event.clientY - rect.top;
    const before = screenToWorld(mouseX, mouseY);
    const factor = Math.exp(-event.deltaY * 0.00078);
    state.zoom = clamp(state.zoom * factor, zoomLimits.min, zoomLimits.max);
    const after = worldToScreen(before.x, before.y);
    state.panX += mouseX - after.x;
    state.panY += mouseY - after.y;
    clampPan();
    syncCanvasViewDataset();
  }

  function resetZoom() {
    state.zoom = 1;
    state.panX = 0;
    state.panY = 0;
    hideLinkMenu();
    syncCanvasViewDataset();
  }

  function viewCenter() {
    const bounds = sceneBounds();
    return { x: bounds.cx, y: bounds.cy };
  }

  function sceneBounds() {
    const compact = state.width < 1100 || state.height < 680;
    const cssVars = getComputedStyle(document.documentElement);
    const cssSideTop = parseFloat(cssVars.getPropertyValue('--side-top'));
    const sideWidth = compact
      ? Math.min(286, state.width * 0.28)
      : (document.querySelector('.controls')?.getBoundingClientRect().width || clamp(state.width * 0.192, 320, 392));
    const sideGap = compact ? 16 : 24;
    const top = compact ? 94 : (Number.isFinite(cssSideTop) ? cssSideTop : 104);
    const bottom = compact ? 14 : 18;
    const leftRail = state.leftCollapsed ? 44 : sideWidth + 26;
    const rightRail = compact ? 44 : sideWidth + 26;
    const left = sideGap + leftRail;
    const right = state.width - sideGap - rightRail;
    const width = Math.max(260, right - left);
    const height = Math.max(260, state.height - top - bottom);
    return {
      left,
      right,
      top,
      bottom,
      width,
      height,
      cx: left + width / 2,
      cy: top + height / 2,
    };
  }

  function worldToScreen(x, y) {
    const center = viewCenter();
    return {
      x: center.x + state.panX + (x - center.x) * state.zoom,
      y: center.y + state.panY + (y - center.y) * state.zoom,
    };
  }

  function screenToWorld(x, y) {
    const center = viewCenter();
    return {
      x: center.x + (x - center.x - state.panX) / state.zoom,
      y: center.y + (y - center.y - state.panY) / state.zoom,
    };
  }

  function clampPan() {
    if (state.zoom <= 1) {
      state.panX = 0;
      state.panY = 0;
      return;
    }
    const maxX = state.width * (state.zoom - 1) * 0.62;
    const maxY = state.height * (state.zoom - 1) * 0.62;
    state.panX = clamp(state.panX, -maxX, maxX);
    state.panY = clamp(state.panY, -maxY, maxY);
  }

  function targetAlpha(ext) {
    const groupValue = valueFor(ext, state.dimension);
    const focusGroup = state.previewGroup || state.activeGroup;
    let alpha = 0.92;
    if (focusGroup && focusGroup !== groupValue) alpha = 0.10;
    if (focusGroup && focusGroup === groupValue) alpha = 1;
    if (state.hover === ext) alpha = 1;
    if (state.linkMenuExt === ext) alpha = 1;
    return alpha;
  }

  function compareGroupLabels(dimKey, a, b) {
    const [labelA, countA] = a;
    const [labelB, countB] = b;
    const order = orderedValues[dimKey];
    if (order) {
      const indexA = orderIndex(order, labelA);
      const indexB = orderIndex(order, labelB);
      if (indexA !== indexB) return indexA - indexB;
    } else if (dimKey === 'pgrx') {
      const pgrxOrder = comparePgrxLabels(labelA, labelB);
      if (pgrxOrder !== 0) return pgrxOrder;
    }
    return countB - countA || labelA.localeCompare(labelB);
  }

  function orderIndex(order, label) {
    const index = order.indexOf(label);
    return index === -1 ? order.length + 100 : index;
  }

  function comparePgrxLabels(a, b) {
    if (a === 'Not PGRX' && b !== 'Not PGRX') return 1;
    if (b === 'Not PGRX' && a !== 'Not PGRX') return -1;
    if (a.includes('unknown') && !b.includes('unknown')) return 1;
    if (b.includes('unknown') && !a.includes('unknown')) return -1;
    const va = versionParts(a);
    const vb = versionParts(b);
    for (let i = 0; i < Math.max(va.length, vb.length); i += 1) {
      const diff = (vb[i] || 0) - (va[i] || 0);
      if (diff !== 0) return diff;
    }
    return a.localeCompare(b);
  }

  function versionParts(value) {
    return String(value).match(/\d+/g)?.map(Number) || [];
  }

  function semanticColorFor(dimKey, label) {
    const text = String(label);
    if (/unknown/i.test(text)) return semanticColors.muted;
    if (['starTier', 'watchTier', 'forkTier'].includes(dimKey)) {
      if (text.startsWith('T0')) return '#3de0ca';
      if (text.startsWith('T1')) return '#69a7ff';
      if (text.startsWith('T2')) return '#b68cff';
      if (text.startsWith('T3')) return '#f4b860';
      if (text.startsWith('T4')) return '#ff6b8a';
    }
    if (dimKey === 'packageCoverage') {
      const coveragePalette = [
        '#3de0ca', '#58e88f', '#84e36f', '#b6df5a', '#d6d64f', '#e7c84d',
        '#f4b860', '#f3a24f', '#ee8e49', '#e87943', '#ef6f4f', '#f06267',
        '#f45874', '#f05282', '#e84a94', '#d946a8', '#c043b8', '#ff6b8a',
      ];
      const index = packageCoverageOrder.indexOf(text);
      if (index >= 0) return coveragePalette[Math.min(index, coveragePalette.length - 1)];
    }
    if (['lastUpdated', 'lastRelease'].includes(dimKey)) {
      const agePalette = ['#3de0ca', '#6ee7b7', '#69a7ff', '#a9a2ff', '#c084fc', '#e0d15c', '#f4b860', '#fb923c', '#f97316', '#fb7185', '#ff6b8a'];
      const index = orderedValues[dimKey]?.indexOf(text) ?? -1;
      if (index >= 0) return agePalette[Math.min(index, agePalette.length - 1)];
    }
    if (dimKey === 'sourceArchive') {
      if (text === 'Both') return semanticColors.strongPositive;
      if (text === 'Rpm') return '#69a7ff';
      if (text === 'Deb') return '#f4b860';
      if (text === 'None') return semanticColors.negative;
    }
    if ([
      'No package',
      'No binary repo',
      'No binary matrix',
      'None',
      'Not on PGXN',
      'No client binary',
      'No shared library',
      'Superuser required',
      'Discovery candidate',
      'Alias / sub-extension',
    ].includes(text)) return semanticColors.negative;
    if ([
      'Catalog extension',
      'RPM + DEB',
      'Full RPM + DEB matrix',
      'Both',
      'Has client binary',
      'Has shared library',
      'Trusted',
      'Relocatable',
      'Listed on PGXN',
      'Lead extension',
      'PostgreSQL contrib',
      'No DDL needed',
      'No LOAD required',
    ].includes(text)) return semanticColors.positive;
    if (text.includes('RPM') || text.includes('DEB') || text.includes('Needs DDL') || text.includes('Preload / LOAD')) {
      return semanticColors.warning;
    }
    return null;
  }

  function displayValueFor(dimKey, label) {
    const text = String(label);
    if (dimKey === 'starTier') {
      return text
        .replace('T0: 10,000+ stars', 'T0: ≥10,000')
        .replace('T1: 1,000-9,999 stars', 'T1: ≥1,000')
        .replace('T2: 100-999 stars', 'T2: ≥100')
        .replace('T3: 10-99 stars', 'T3: ≥10')
        .replace('T4: 0-9 stars', 'T4: <10');
    }
    if (dimKey === 'watchTier') {
      return text
        .replace('T0: 100+ watchers', 'T0: ≥100')
        .replace('T1: 50-99 watchers', 'T1: ≥50')
        .replace('T2: 10-49 watchers', 'T2: ≥10')
        .replace('T3: 1-9 watchers', 'T3: ≥1')
        .replace('T4: 0 watchers', 'T4: 0');
    }
    if (dimKey === 'forkTier') {
      return text
        .replace('T0: 1,000+ forks', 'T0: ≥1,000')
        .replace('T1: 100-999 forks', 'T1: ≥100')
        .replace('T2: 10-99 forks', 'T2: ≥10')
        .replace('T3: 1-9 forks', 'T3: ≥1')
        .replace('T4: 0 forks', 'T4: 0');
    }
    if (dimKey === 'lastUpdated' || dimKey === 'lastRelease') {
      if (text === 'Within 1 year') return '< 1 year';
      if (text === '10+ years') return '> 10 years';
      const match = text.match(/^(\d+)-(\d+) years$/);
      if (match) return `< ${match[2]} years`;
    }
    if (dimKey === 'packageCoverage') {
      if (text === '80 slots') return '80 slots';
      if (text === '0 slots' || text === '1-4 slots') return text;
      const match = text.match(/^(\d+)-(\d+) slots$/);
      if (match) return `≥${match[1]} slots`;
    }
    return text;
  }

  function valueFor(ext, dim) {
    return ext.dimensions?.[dim] || ext[dim] || 'Unknown';
  }

  function colorFor(ext, dim) {
    return colorMaps[dim]?.get(valueFor(ext, dim)) || '#8aa0a4';
  }

  function packageAvailability(ext) {
    if (ext.hasRpm && ext.hasDeb) return 'RPM + DEB';
    if (ext.hasRpm) return 'RPM only';
    if (ext.hasDeb) return 'DEB only';
    return 'No package';
  }

  function flagsFor(ext) {
    return [
      ext.contrib ? 'contrib' : null,
      ext.lead ? 'lead' : null,
      ext.hasBin ? 'bin' : null,
      ext.hasLib ? 'lib' : null,
      ext.needDdl ? 'ddl' : null,
      ext.needLoad ? 'load' : null,
      ext.trusted ? 'trusted' : null,
      ext.relocatable ? 'relocatable' : null,
    ].filter(Boolean).join(', ') || 'none';
  }

  function formatDate(value) {
    return value || 'unknown';
  }

  function radiusForStars(stars) {
    const maxStars = Math.max(1, payload.summary?.maxStars || 1);
    const normalized = Math.log10((stars || 0) + 1) / Math.log10(maxStars + 1);
    return 4.0 + Math.pow(normalized, 0.72) * 10.5;
  }

  function formatMetric(value) {
    if (typeof value === 'number') return formatCompact(value);
    return value == null ? '0' : String(value);
  }

  function formatCompact(value) {
    const n = Number(value || 0);
    if (n >= 1000000) return `${(n / 1000000).toFixed(1)}M`;
    if (n >= 1000) return `${(n / 1000).toFixed(n >= 10000 ? 0 : 1)}k`;
    return String(n);
  }

  function formatNumber(value) {
    return Number(value || 0).toLocaleString('en-US');
  }

  function clamp(value, min, max) {
    return Math.min(max, Math.max(min, value));
  }

  function escapeHtml(value) {
    return String(value ?? '').replace(/[&<>"']/g, (char) => ({
      '&': '&amp;',
      '<': '&lt;',
      '>': '&gt;',
      '"': '&quot;',
      "'": '&#39;',
    }[char]));
  }

  function escapeAttr(value) {
    return escapeHtml(value).replace(/`/g, '&#96;');
  }

  function cssEscape(value) {
    if (window.CSS?.escape) return window.CSS.escape(String(value));
    return String(value).replace(/["\\\]]/g, '\\$&');
  }

  function hash(value) {
    let h = 2166136261;
    const text = String(value);
    for (let i = 0; i < text.length; i += 1) {
      h ^= text.charCodeAt(i);
      h = Math.imul(h, 16777619);
    }
    return ((h >>> 0) % 1000000) / 1000000;
  }

  function colorWithAlpha(hex, alpha) {
    const value = hex.replace('#', '');
    const r = parseInt(value.slice(0, 2), 16);
    const g = parseInt(value.slice(2, 4), 16);
    const b = parseInt(value.slice(4, 6), 16);
    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
  }

  function roundRect(context, x, y, width, height, radius) {
    const r = Math.min(radius, width / 2, height / 2);
    context.beginPath();
    context.moveTo(x + r, y);
    context.arcTo(x + width, y, x + width, y + height, r);
    context.arcTo(x + width, y + height, x, y + height, r);
    context.arcTo(x, y + height, x, y, r);
    context.arcTo(x, y, x + width, y, r);
    context.closePath();
  }

  function intersects(a, b) {
    return a.x < b.x + b.w && a.x + a.w > b.x && a.y < b.y + b.h && a.y + a.h > b.y;
  }
})();
