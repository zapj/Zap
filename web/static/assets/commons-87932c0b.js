const n=function(e,o){if(e===0)return"0 Bytes";var a=1024,t=o||2,B=["Bytes","KB","MB","GB","TB","PB","EB","ZB","YB"],r=Math.floor(Math.log(e)/Math.log(a));return parseFloat((e/Math.pow(a,r)).toFixed(t))+" "+B[r]},s=function(e,o){if(e===0)return"0 Bytes";var a=1024,t=o||2,B=["Bytes","KB","MB","GB","TB","PB","EB","ZB","YB"],r=Math.floor(Math.log(e)/Math.log(a));return[parseFloat((e/Math.pow(a,r)).toFixed(t)),B[r]]},l=(e,o)=>{let a=e[o.property];if(a===null||a===0)return"";let t=new Date(a*1e3);return t.getFullYear()+"-"+(t.getMonth()+1)+"-"+t.getDate()+" "+t.getHours()+":"+t.getMinutes()+":"+t.getSeconds()};export{s as a,l as b,n as f};
