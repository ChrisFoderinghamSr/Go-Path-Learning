// Amplify hosts the public course catalog. Student sessions belong to the HTTPS Go service.
import fs from 'node:fs';
import vm from 'node:vm';

const context = vm.createContext({window: {}});
vm.runInContext(fs.readFileSync('dist/curriculum.js', 'utf8'), context);
const modules = context.window.CURRICULUM;
const escape = value => String(value).replace(/[&<>"']/g, c =>
  ({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]));
const target = process.env.STUDENT_SERVICE_ORIGIN;
let enrollment = '<p>Verified enrollment opens when the student service is connected.</p>';
if (target) {
  const url = new URL(target);
  if (url.protocol !== 'https:' || url.origin !== target) throw Error('Use a bare HTTPS student service origin.');
  enrollment = '<p><a class="btn" href="'+escape(url.origin)+'/">Enter the student service →</a></p>';
}
const out = 'amplify-site';
fs.mkdirSync(out, {recursive:true});
for (const name of ['style.css','path-emblem.png','course-guide.md']) {
  fs.copyFileSync('dist/'+name, out+'/'+name);
}
const count = modules.reduce((n,m)=>n+m.lessons.length,0);
const list = modules.map((m,i)=>'<details><summary>'+String(i+1).padStart(2,'0')+' · '+escape(m.title)+'</summary><p>'+escape(m.subtitle)+'</p><ul>'+m.lessons.map(l=>'<li>'+escape(l.title)+'</li>').join('')+'</ul><p>Project · '+escape(m.project.title)+'</p></details>').join('');
fs.writeFileSync(out+'/index.html','<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Go Path · PATH Learning Center</title><link rel="stylesheet" href="style.css"><main><article class="lesson-body"><img src="path-emblem.png" width="96" height="96" alt="PATH Learning Center of excellence seal"><div class="eyebrow">PATH LEARNING CENTER of excellence</div><h1>Go Path: Beginner to Full-Stack</h1><p>Explore the '+count+' lesson checkpoints and '+modules.length+' projects, from your first Go program to a full-stack task board.</p><a class="btn secondary" href="course-guide.md" download>Download the complete course guide</a><h2>Student enrollment</h2>'+enrollment+'<h2>Curriculum</h2>'+list+'</article></main></html>');
