import fs from 'node:fs';import vm from 'node:vm';import path from 'node:path';import {spawnSync}from'node:child_process';
const root=path.resolve('.example-check');fs.mkdirSync(root,{recursive:true});
const ctx=vm.createContext({window:{}});vm.runInContext(fs.readFileSync('dist/curriculum.js','utf8'),ctx);
const results=[];
for(const m of ctx.window.CURRICULUM)for(const l of m.lessons)for(const key of ['code','solution']){
 const source=l[key];if(!/^package (main|\w+)\s/m.test(source)||source.includes('```')||source.includes('// File:')||source.includes('// file:'))continue;
 const dir=path.join(root,l.id+'-'+key);fs.mkdirSync(dir,{recursive:true});fs.writeFileSync(path.join(dir,'go.mod'),'module example.com/check\n\ngo 1.22\n');
 if(l.id==='modules'&&key==='solution'){
  fs.writeFileSync(path.join(dir,'go.mod'),'module example.com/taskboard\n\ngo 1.22\n');
  fs.mkdirSync(path.join(dir,'internal/task'),{recursive:true});fs.mkdirSync(path.join(dir,'cmd/taskboard'),{recursive:true});
  const parts=source.split('// cmd/taskboard/main.go (separate file)');fs.writeFileSync(path.join(dir,'internal/task/task.go'),parts[0]);fs.writeFileSync(path.join(dir,'cmd/taskboard/main.go'),parts[1]);
  if(fs.existsSync(path.join(dir,'main.go')))fs.unlinkSync(path.join(dir,'main.go'));
 }else{
  fs.writeFileSync(path.join(dir,source.includes('func Test')||source.includes('func Fuzz')?'example_test.go':'main.go'),source);
  if(l.id==='unit-tests')fs.writeFileSync(path.join(dir,'clamp.go'),'package main\n'+ctx.window.CURRICULUM.flatMap(m=>m.lessons).find(x=>x.id==='functions').solution);
 }
 const check=spawnSync('go',['test','./...'],{cwd:dir,encoding:'utf8',timeout:120000});results.push({lesson:l.id,kind:key,status:check.status,output:(check.stdout+check.stderr).trim()});
}
fs.writeFileSync('.example-results.json',JSON.stringify(results,null,2));console.log(JSON.stringify({checked:results.length,failed:results.filter(r=>r.status!==0)},null,2));

