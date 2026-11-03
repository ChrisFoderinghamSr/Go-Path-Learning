// Load the course only after the backend confirms an enrolled GitHub identity.
(async function () {
 const main = document.querySelector('main');
 function notice(message) { main.textContent = message; }
 try {
   const response = await fetch('/student/state', {cache:'no-store',redirect:'error'});
   if (!response.ok) throw Error('Sign in at the student portal.');
   const account = await response.json();
   if (!Number.isSafeInteger(account.id) || account.id <= 0) throw Error('Enrollment required.');
   let version=account.version, pending, running=false, stopped=false, timer;
   const status=document.createElement('span');
   status.setAttribute('role','status');
   const portal=document.createElement('a');portal.href='/student';portal.textContent='Student record';
   document.querySelector('.top-actions').prepend(portal,status);
   async function flush(){
     if(running || stopped || !pending)return;
     running=true;
     const progress=pending;pending=null;
     try{
       const result=await fetch('/student/progress',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({version,progress})});
       if(result.status===409){stopped=true;throw Error('Another tab saved changes. Export your draft, then reload.');}
       if(!result.ok)throw Error('Save failed. Keep this tab open and retry or export your work.');
       const saved=await result.json();version=saved.version;status.textContent=saved.warning||'Saved to your account';
     }catch(error){pending=pending||progress;status.textContent=error.message;running=false;return}
     running=false;if(pending)flush();
   }
   window.addEventListener('online',flush);
   const retry=document.createElement('button');retry.className='quiet';retry.textContent='Retry save';retry.onclick=flush;status.after(retry);
   window.addEventListener('beforeunload',e=>{if(running||pending){e.preventDefault();e.returnValue='';}});
   window.studentAccount={id:account.id,progress:account.progress,save(progress){pending=JSON.parse(JSON.stringify(progress));status.textContent='Saving…';clearTimeout(timer);timer=setTimeout(flush,450);}};
   const script=document.createElement('script');script.src='/app.js';document.body.append(script);
 }catch{
   notice('This course requires the GitHub student backend. ');
   const a=document.createElement('a');a.href='/student/login';a.textContent='Sign in with GitHub';main.append(a);
 }
})();
