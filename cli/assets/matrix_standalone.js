const tabs = document.querySelectorAll(".mx-tab");
const legendCounts = document.querySelectorAll(".mx-legend-item em[data-full-count]");
function applyMatrixMode(mode) {
  document.body.dataset.view = mode;
  for (const count of legendCounts) {
    count.textContent = mode === "pgdg" ? count.dataset.pgdgCount : count.dataset.fullCount;
  }
}
for (const tab of tabs) {
  tab.addEventListener("click", () => {
    applyMatrixMode(tab.dataset.mode);
    for (const item of tabs) {
      const active = item === tab;
      item.classList.toggle("is-active", active);
      item.setAttribute("aria-selected", active ? "true" : "false");
    }
  });
}
applyMatrixMode(document.body.dataset.view || "pgdg");
