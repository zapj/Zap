import{r as u,o as B,a as g,f as D,j as t,w as e,p as n,e as s,h as _,m as R,b as z,S,ah as A,F as M,q as N,z as q}from"./index-e26e72e9.js";import{E as F}from"./el-dialog-6ca727bc.js";import"./el-overlay-797d8413.js";import{E as I,a as P}from"./el-col-fa218f38.js";import{E as j}from"./el-card-e1aad0b1.js";import{E as O,a as U}from"./el-table-column-95096c36.js";import"./el-tooltip-95af38b0.js";import"./el-scrollbar-eafc889c.js";import{E as G}from"./el-button-91f41336.js";/* empty css               */import{a as m}from"./client-00b0bb33.js";import{b as w}from"./commons-87932c0b.js";import{E as H}from"./index-5dd34bf7.js";import{E as f}from"./index-b3b7ff7e.js";import"./index-6fc089e4.js";import"./error-78e43d3e.js";import"./index-3bac334e.js";import"./index-0e6c4a62.js";import"./runtime-909e2ee3.js";import"./constants-b6c0c234.js";import"./vnode-4474326f.js";import"./use-global-config-e0c846e4.js";import"./use-form-common-props-0f5bc46c.js";import"./focus-trap-f40e746e.js";import"./typescript-d2ff6dd6.js";import"./isNil-c75b1b34.js";import"./refs-5d6d1997.js";import"./index-7cdd4eea.js";import"./index-6838fe58.js";import"./event-8e91c63d.js";import"./index-9f57bd93.js";import"./_initCloneObject-9228ffc4.js";import"./use-form-item-59e5f47e.js";const J={class:"dialog-footer"},Lt={__name:"TaskListView",setup(K){const d=u([]),p=u(!1),k=u(""),y=N(()=>q(()=>import("./LogViewer-f46a82ef.js"),["assets/LogViewer-f46a82ef.js","assets/xterm-db03418d.js","assets/xterm-ff373be6.css","assets/index-e26e72e9.js","assets/index-dd61018d.css"]));B(()=>{c()});const c=()=>{m({url:"/v1/task/appinstall/tasklist"}).then(r=>{d.value=r.data})},E=(r,o)=>{m({url:"/v1/task/appinstall/removetask",data:{id:o},dataType:"form",method:"post"}).then(a=>{d.value.splice(r,1),f({message:"删除成功",type:"success"})})},b=(r,o)=>{m({url:"/v1/task/appinstall/canceltask",data:{id:o},dataType:"form",method:"post"}).then(a=>{c(),f({message:"取消成功",type:"success"})})},h=(r,o)=>{k.value="install_"+o+".log",p.value=!0},C=()=>{m({url:"/v1/task/appinstall/gentask",dataType:"form",method:"post"}).then(r=>{f({message:"生成成功",type:"success"})})};return(r,o)=>{const a=O,i=G,V=U,T=j,x=I,L=P,$=F;return g(),D(M,null,[t(L,{gutter:20},{default:e(()=>[t(x,{span:24,class:"mb-2"},{default:e(()=>[t(T,{"body-style":{padding:"0px"},shadow:"hover"},{footer:e(()=>[t(i,{type:"primary",onClick:C},{default:e(()=>[n("安装")]),_:1}),t(i,{type:"primary",onClick:c},{default:e(()=>[n("刷新")]),_:1})]),default:e(()=>[t(V,{data:s(d),stripe:"",style:{width:"100%"}},{default:e(()=>[t(a,{prop:"id",label:"ID",width:"180"}),t(a,{prop:"title",label:"任务名称"}),t(a,{prop:"start_time",label:"开始时间",width:"180",formatter:s(w)},null,8,["formatter"]),t(a,{prop:"end_time",label:"完成时间",width:"180",formatter:s(w)},null,8,["formatter"]),t(a,{prop:"Status",label:"状态",width:"120"}),t(a,{prop:"error",label:"Error"}),t(a,{label:"操作",width:"120"},{default:e(l=>[t(i,{link:"",type:"primary",size:"small",onClick:_(v=>E(l.$index,l.row.id),["prevent"])},{default:e(()=>[n(" 删除 ")]),_:2},1032,["onClick"]),t(i,{link:"",type:"primary",size:"small",onClick:_(v=>b(l.$index,l.row.id),["prevent"])},{default:e(()=>[n(" 取消 ")]),_:2},1032,["onClick"]),t(i,{link:"",type:"primary",size:"small",onClick:_(v=>h(l.$index,l.row.id),["prevent"])},{default:e(()=>[n(" View Log ")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])]),_:1})]),_:1})]),_:1}),t($,{modelValue:s(p),"onUpdate:modelValue":o[1]||(o[1]=l=>A(p)?p.value=l:null),width:"80%",title:"LogViewer"},{footer:e(()=>[R("div",J,[t(i,{onClick:o[0]||(o[0]=l=>p.value=!1)},{default:e(()=>[n("关闭")]),_:1})])]),default:e(()=>[(g(),z(S,null,{default:e(()=>[t(s(y),{logfile:s(k)},null,8,["logfile"])]),fallback:e(()=>[t(s(H),{rows:4})]),_:1}))]),_:1},8,["modelValue"])],64)}}};export{Lt as default};
