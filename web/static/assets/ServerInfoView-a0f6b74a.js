import{S as g,o as y,k as c,l as p,G as t,H as n,A as r,q as i,Y as u,U as w,ae as S,m as f,ac as x,ad as P}from"./index-163cc183.js";import{E as V}from"./el-progress-9fba6ff7.js";/* empty css               */import{a as b}from"./client-cc321679.js";import{f as m}from"./commons-890e5e57.js";import{_ as B}from"./_plugin-vue_export-helper-c27b6911.js";import{E as C}from"./index-dc84e4c6.js";import"./index-5c76aaa6.js";import"./index-ffe902e3.js";import"./runtime-bf08026f.js";import"./use-form-common-props-75f0c225.js";import"./index-d0b2b934.js";import"./aria-4df7e579.js";const l=d=>(x("data-v-a7411e4b"),d=d(),P(),d),E={class:"font-size-3 serverinfo"},U=l(()=>t("thead",null,[t("tr",null,[t("td"),t("td"),t("td")])],-1)),z=l(()=>t("td",null,"操作系统版本",-1)),F=l(()=>t("td",null,null,-1)),H=l(()=>t("td",null,"系统内核",-1)),N=l(()=>t("td",null,null,-1)),T=l(()=>t("td",null,"主机名",-1)),q=l(()=>t("td",null,null,-1)),A=l(()=>t("td",null,"服务器启动时间",-1)),G=l(()=>t("td",null,null,-1)),K=l(()=>t("td",null,"运行时间",-1)),D=l(()=>t("td",null,null,-1)),L=l(()=>t("td",null,"系统负载",-1)),M=l(()=>t("td",null,null,-1)),O=l(()=>t("td",null,"CPU 名称",-1)),R=l(()=>t("td",null,null,-1)),Y=l(()=>t("td",null,"CPU 频率",-1)),j=l(()=>t("td",null,null,-1)),k=l(()=>t("td",null,"CPU 核心数",-1)),J=l(()=>t("td",null,null,-1)),Q=l(()=>t("td",null,"内存总大小",-1)),W=l(()=>t("td",null,null,-1)),X=l(()=>t("td",null,"已用内存",-1)),Z=l(()=>t("td",null,null,-1)),$=l(()=>t("td",null,"剩余可以内存",-1)),tt=l(()=>t("td",null,null,-1)),lt=l(()=>t("td",null,"Swap内存总大小",-1)),et=l(()=>t("td",null,null,-1)),nt=l(()=>t("td",null,"已用内存",-1)),ot=l(()=>t("td",null,null,-1)),st=l(()=>t("td",null,"剩余可以内存",-1)),rt=l(()=>t("td",null,null,-1)),o=3,ut={__name:"ServerInfoView",setup(d){const e=g({serverInfo:{}});y(()=>{v()});const v=()=>{b({url:"/v1/server/info"}).then(h=>{e.serverInfo=h.data,console.log(e.serverInfo)})};return(h,dt)=>{const _=C,a=V;return c(),p("table",E,[U,t("tbody",null,[t("tr",null,[t("td",{width:"150px",colspan:o},"服务器信息")]),t("tr",null,[z,t("td",null,n(e.serverInfo.OS)+" / "+n(e.serverInfo.Platform)+"-"+n(e.serverInfo.PlatformVersion),1),F]),t("tr",null,[H,t("td",null,n(e.serverInfo.KernelVersion)+" "+n(e.serverInfo.KernelArch),1),N]),t("tr",null,[T,t("td",null,n(e.serverInfo.Hostname),1),q]),t("tr",null,[A,t("td",null,n(e.serverInfo.BootTime),1),G]),t("tr",null,[K,t("td",null,n(e.serverInfo.uptime),1),D]),t("tr",null,[L,t("td",null,[r(_,{class:"ml-2",type:"success"},{default:i(()=>[u(n(e.serverInfo.load1),1)]),_:1}),u(),r(_,{class:"ml-2",type:"success"},{default:i(()=>[u(n(e.serverInfo.load5),1)]),_:1}),r(_,{class:"ml-2",type:"success"},{default:i(()=>[u(n(e.serverInfo.load15),1)]),_:1}),r(a,{percentage:e.serverInfo.system_load1},null,8,["percentage"])]),M]),t("tr",null,[t("td",{colspan:o}," ")]),t("tr",null,[t("td",{colspan:o},"CPU")]),t("tr",null,[O,t("td",null,n(e.serverInfo.cpu_name),1),R]),t("tr",null,[Y,t("td",null,n(e.serverInfo.cpu_ghz)+"GHz",1),j]),t("tr",null,[k,t("td",null,n(e.serverInfo.cpu_cores)+"核心",1),J]),t("tr",null,[t("td",{colspan:o}," ")]),t("tr",null,[t("td",{colspan:o},"内存")]),t("tr",null,[Q,t("td",null,n(e.serverInfo.memory_total),1),W]),t("tr",null,[X,t("td",null,[u(n(e.serverInfo.memory_used)+" ",1),r(a,{percentage:e.serverInfo.memory_used_percent},null,8,["percentage"])]),Z]),t("tr",null,[$,t("td",null,n(e.serverInfo.memory_available),1),tt]),t("tr",null,[t("td",{colspan:o}," ")]),t("tr",null,[t("td",{colspan:o},"Swap内存")]),t("tr",null,[lt,t("td",null,n(e.serverInfo.swapmem_total),1),et]),t("tr",null,[nt,t("td",null,[u(n(e.serverInfo.swapmem_used)+" ",1),r(a,{percentage:e.serverInfo.swapmem_used_percent},null,8,["percentage"])]),ot]),t("tr",null,[st,t("td",null,n(e.serverInfo.swapmem_free),1),rt]),t("tr",null,[t("td",{colspan:o}," ")]),t("tr",null,[t("td",{colspan:o},"硬盘信息")]),t("tr",null,[t("td",{colspan:o},[t("table",null,[(c(!0),p(w,null,S(e.serverInfo.disk,(s,I)=>(c(),p("tr",{key:I},[t("td",null,n(s.path),1),t("td",null,n(s.fstype),1),t("td",null,n(f(m)(s.total)),1),t("td",null,n(f(m)(s.free)),1),t("td",null,n(f(m)(s.used)),1),t("td",null,n(parseFloat(s.usedPercent).toFixed(2))+"%",1)]))),128))])])])])])}}},St=B(ut,[["__scopeId","data-v-a7411e4b"]]);export{St as default};