import{r as i,c as h,o as v,a as m,f as g,O as w,b as x,w as p,j as t,p as d,t as n,m as D,F as E}from"./index-5f1d5677.js";/* empty css                   */import{E as M,a as y}from"./el-table-column-62994bae.js";import"./el-tooltip-3d978724.js";import"./el-scrollbar-f1666b53.js";/* empty css               */import{a as B,v as C}from"./client-784b0ed6.js";import"./error-78e43d3e.js";import"./runtime-89a56c03.js";import"./_initCloneObject-1dcc91f6.js";import"./use-global-config-71ca7aac.js";import"./index-f7883cef.js";import"./index-fce2b169.js";import"./event-8e91c63d.js";import"./use-form-common-props-4758c941.js";import"./constants-84857360.js";import"./use-form-item-7da9f99e.js";import"./index-110ef5f1.js";import"./isNil-c75b1b34.js";import"./typescript-e92ae92a.js";import"./focus-trap-c933673f.js";import"./index-90345743.js";const N={class:"text-center font-size-3"},$={__name:"PSView",setup(P){const r=i([]),a=i(!1),l=i(300),c=()=>{let o=document.body.clientHeight-215-140;o<300?l.value=300:l.value=o,console.log(o)},u=h(()=>r.value.length);v(()=>{_(),window.onresize=c});const _=()=>{a.value=!0,B({url:"/v1/server/processlist",loading:!1}).then(o=>{r.value=o.data,c(),a.value=!1}).finally(()=>a.value=!1)};return(o,T)=>{const e=M,b=y,f=C;return m(),g(E,null,[w((m(),x(b,{data:r.value,border:"",style:{width:"100%"},"default-sort":{prop:"pid",order:"descending"},height:l.value},{default:p(()=>[t(e,{prop:"pid",label:"PID",width:"120",sortable:"",fixed:""}),t(e,{prop:"name",label:"进程名称",sortable:""}),t(e,{prop:"username",label:"用户",width:"150",sortable:""}),t(e,{prop:"cpu_percent",label:"CPU%",width:"100",sortable:""},{default:p(s=>[d(n(s.row.cpu_percent)+" % ",1)]),_:1}),t(e,{prop:"mem_percent",label:"MEM%",width:"100",sortable:""},{default:p(s=>[d(n(s.row.mem_percent)+" % ",1)]),_:1}),t(e,{prop:"create_time",label:"启动时间",width:"180"}),t(e,{prop:"cmdline",label:"COMMAND"})]),_:1},8,["data","height"])),[[f,a.value]]),D("p",N,"总进程数: "+n(u.value),1)],64)}}};export{$ as default};