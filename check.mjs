import fs from 'node:fs';
import vm from 'node:vm';
import assert from 'node:assert/strict';
const elements=new Map();
const element=key=>{if(!elements.has(key))elements.set(key,{innerHTML:'',textContent:'',value:'',style:{},addEventListener(){},focus(){},classList:{toggle(){}}});return elements.get(key)};
const listeners={};
const context=vm.createContext({window:{scrollTo(){},addEventListener(){}},location:{hash:''},document:{querySelector:element,querySelectorAll:()=>[],addEventListener:(name,fn)=>listeners[name]=fn},localStorage:{getItem:()=>null,setItem(){}},setTimeout:()=>0,clearTimeout(){},URL,Blob,console,confirm:()=>false});
vm.runInContext(fs.readFileSync('dist/curriculum.js','utf8'),context);
vm.runInContext(fs.readFileSync('dist/app.js','utf8'),context);
const modules=context.window.CURRICULUM;
assert.equal(modules.length,12);
assert.equal(modules.flatMap(m=>m.lessons).length,40);
const ids=new Set();
for(const m of modules){assert(m.project.rubric.length>=3);assert(m.project.requirements.length>=3);for(const l of m.lessons){assert(!ids.has(l.id));ids.add(l.id);assert(l.sections.length>=3,`${l.id} needs depth`);assert(l.code.length>40);assert(l.exercise.length>40);assert(l.solution.length>20);assert(l.quiz.options.length>=3);assert(Number.isInteger(l.quiz.answer));assert(l.quiz.answer>=0&&l.quiz.answer<l.quiz.options.length);assert(l.sources.length>0);for(const s of l.sources)assert(new URL(s.url).protocol==='https:')}}
assert(element('main').innerHTML.includes('Big ideas. Small, runnable steps.'),'root route must show home');
const routes=['home','path','projects','notes','resources',...modules.flatMap(m=>['module/'+m.id,'project/'+m.id,'assessment/'+m.id,...m.lessons.map(l=>'lesson/'+l.id)])];
for(const route of routes){context.location.hash='#'+route;vm.runInContext('render()',context);const html=element('main').innerHTML;assert(html.length>500,route);assert(!html.includes('PAGE NOT FOUND'),route);assert(!html.includes('undefined'),route)}
context.location.hash='#lesson/not-real';vm.runInContext('render()',context);assert(element('main').innerHTML.includes('PAGE NOT FOUND'));
assert.throws(()=>vm.runInContext('validate({version:1,completed:[]})',context));
assert.throws(()=>vm.runInContext('validate({...blank(),completed:["invalid"]})',context));
assert.throws(()=>vm.runInContext('validate({...blank(),notes:{abc:42}})',context));
vm.runInContext('state=blank();state.notes[lessons[0].id]="<img src=x onerror=alert(1)>";location.hash="#notes";render()',context);
assert(!element('main').innerHTML.includes('<img'),'imported note text must be escaped');
console.log(`Passed: ${routes.length} routes, ${ids.size} lesson schemas, backup validation and note escaping.`);
