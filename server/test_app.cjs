const fs = require('node:fs');
const path = require('node:path');
const vm = require('node:vm');

// Use the actual application, without attaching DOM events or starting boot.
const source = fs.readFileSync(path.join(__dirname, 'web/app.js'), 'utf8');
const app = source.slice(0, source.lastIndexOf('\nattachEvents();'));

exports.createAppContext = (globals = {}) => {
  const context = vm.createContext({ window: {}, location: { hostname: 'localhost', search: '' }, ...globals });
  vm.runInContext(app, context);
  return context;
};
