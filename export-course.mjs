import fs from 'node:fs';
import vm from 'node:vm';
const context=vm.createContext({window:{}});vm.runInContext(fs.readFileSync('dist/curriculum.js','utf8'),context);
const modules=context.window.CURRICULUM;
let text='# Go Path — From first line to full stack\n\nA self-paced curriculum. Read, experiment, test, and build. Guided reading estimates exclude independent project work.\n\n';
text+='Recommended order: '+modules.map(m=>m.title).join(' → ')+'.\n\nComplete each exercise, pass its checkpoint, and review the module project before moving on. Revisit earlier modules whenever a concept feels uncertain.\n\n';
for(const [i,m] of modules.entries()){
 text+=`# ${i+1}. ${m.title}\n\n${m.subtitle}\n\nLevel: ${m.level}. Prerequisite: ${i?modules[i-1].title:'None'}.\n\n`;
 for(const l of m.lessons){text+=`## ${l.title}\n\n${l.minutes} minutes of guided study, plus practice.\n\nLearning objectives:\n\n${l.objectives.map(o=>'- '+o).join('\n')}\n\n`;
 for(const s of l.sections)text+=`### ${s.title}\n\n${s.body}\n\n`;
 text+=`### Example\n\n\`\`\`\n${l.code}\n\`\`\`\n\n### Exercise\n\n${l.exercise}\n\nHint: ${l.hint}\n\nReview criteria:\n\n${l.checks.map(c=>'- [ ] '+c).join('\n')}\n\n### Worked solution\n\n\`\`\`\n${l.solution}\n\`\`\`\n\n### Checkpoint\n\n${l.quiz.question}\n\n${l.quiz.options.map((o,i)=>`${i+1}. ${o}`).join('\n')}\n\nAnswer: ${l.quiz.answer+1}. ${l.quiz.explanation}\n\nSources:\n\n${l.sources.map(s=>`- [${s.title}](${s.url})`).join('\n')}\n\n`;
 }
 text+=`## Project: ${m.project.title}\n\n${m.project.brief}\n\nRequirements:\n\n${m.project.requirements.map(r=>'- '+r).join('\n')}\n\nSelf-review rubric:\n\n${m.project.rubric.map(r=>'- [ ] '+r).join('\n')}\n\n`;
}
text+='## Coverage and learning records\n\nThe core path teaches the language and a practical web stack; specialist standard-library packages remain further reference tracks. See sources.json and devdocs-topics.json for the audit. The app saves progress on your device. Export a backup to move or preserve learning records. Code exercises require local execution or the official Go Playground; project rubrics are self-assessments.\n';
fs.writeFileSync('CURRICULUM.md',text);fs.writeFileSync('dist/course-guide.md',text);
console.log(`Exported ${modules.length} modules and ${modules.flatMap(m=>m.lessons).length} lessons.`);
