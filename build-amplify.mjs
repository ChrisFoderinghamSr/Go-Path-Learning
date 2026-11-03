import fs from 'node:fs';

const origin = process.env.STUDENT_SERVICE_ORIGIN || '';
if (origin) {
  const url = new URL(origin);
  if (url.protocol !== 'https:' || url.origin !== origin) {
    throw new Error('STUDENT_SERVICE_ORIGIN must be a bare HTTPS origin.');
  }
}

const out = 'amplify-site';
fs.mkdirSync(out, {recursive:true});
for (const name of ['app.js','style.css','curriculum.js','path-emblem.png','course-guide.md','sources.json','devdocs-index.json','devdocs-topics.json']) {
  fs.copyFileSync('dist/'+name, out+'/'+name);
}
const index = fs.readFileSync('dist/index.html','utf8')
  .replace('src="student.js"', 'src="catalog-config.js"></script><script src="app.js"')
  .replace('id="export"', 'hidden id="export"')
  .replace('GitHub student account','Public course preview');
fs.writeFileSync(out+'/index.html',index);
fs.writeFileSync(out+'/catalog-config.js','window.publicCatalog = '+JSON.stringify({origin})+';\n');
