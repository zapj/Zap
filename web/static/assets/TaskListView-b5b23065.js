import{r as u,o as B,a as g,f as D,j as e,w as t,L as n,e as r,h as _,K as R,b as z,a0 as A,ai as M,F as N,O as S,U as F}from"./index-ef788864.js";import{E as I}from"./el-dialog-d8980eac.js";import"./el-overlay-5305a6db.js";import{E as O,a as P}from"./el-col-e9d5cd19.js";import{E as U}from"./el-card-8494bd30.js";import{E as j,a as q}from"./el-table-column-ecd9c7ec.js";import"./el-tooltip-6781d3c2.js";import{E as K}from"./el-button-470da2a3.js";/* empty css               */import{a as m}from"./client-16215366.js";import{b as w}from"./commons-87932c0b.js";import{E as G}from"./index-dd988199.js";import{b as f}from"./index-faf3eb96.js";import"./index-b5f3f95a.js";import"./error-78e43d3e.js";import"./index-499c28ec.js";import"./index-17346e61.js";import"./runtime-6c5e7d4b.js";import"./use-form-common-props-f4525b0b.js";import"./index-fa7350a0.js";import"./event-8e91c63d.js";import"./use-form-item-26fbbe24.js";import"./_initCloneObject-ea0126de.js";import"./isNil-c75b1b34.js";const H={class:"dialog-footer"},we={__name:"TaskListView",setup(J){const d=u([]),p=u(!1),k=u(""),b=S(()=>F(()=>import("./LogViewer-c44f040f.js"),["assets/LogViewer-c44f040f.js","assets/xterm-db03418d.js","assets/xterm-ff373be6.css","assets/index-ef788864.js","assets/index-d7f79bda.css"]));B(()=>{c()});const c=()=>{m({url:"/v1/task/appinstall/tasklist"}).then(s=>{d.value=s.data})},y=(s,a)=>{m({url:"/v1/task/appinstall/removetask",data:{id:a},dataType:"form",method:"post"}).then(o=>{d.value.splice(s,1),f({message:"删除成功",type:"success"})})},E=(s,a)=>{m({url:"/v1/task/appinstall/canceltask",data:{id:a},dataType:"form",method:"post"}).then(o=>{c(),f({message:"取消成功",type:"success"})})},h=(s,a)=>{k.value="install_"+a+".log",p.value=!0},C=()=>{m({url:"/v1/task/appinstall/gentask",dataType:"form",method:"post"}).then(s=>{f({message:"生成成功",type:"success"})})};return(s,a)=>{const o=j,i=K,V=q,T=U,x=O,L=P,$=I;return g(),D(N,null,[e(L,{gutter:20},{default:t(()=>[e(x,{span:24,class:"mb-2"},{default:t(()=>[e(T,{"body-style":{padding:"0px"},shadow:"hover"},{footer:t(()=>[e(i,{type:"primary",onClick:C},{default:t(()=>[n("安装")]),_:1}),e(i,{type:"primary",onClick:c},{default:t(()=>[n("刷新")]),_:1})]),default:t(()=>[e(V,{data:r(d),stripe:"",style:{width:"100%"}},{default:t(()=>[e(o,{prop:"id",label:"ID",width:"180"}),e(o,{prop:"title",label:"任务名称"}),e(o,{prop:"start_time",label:"开始时间",width:"180",formatter:r(w)},null,8,["formatter"]),e(o,{prop:"end_time",label:"完成时间",width:"180",formatter:r(w)},null,8,["formatter"]),e(o,{prop:"Status",label:"状态",width:"120"}),e(o,{prop:"error",label:"Error"}),e(o,{label:"操作",width:"120"},{default:t(l=>[e(i,{link:"",type:"primary",size:"small",onClick:_(v=>y(l.$index,l.row.id),["prevent"])},{default:t(()=>[n(" 删除 ")]),_:2},1032,["onClick"]),e(i,{link:"",type:"primary",size:"small",onClick:_(v=>E(l.$index,l.row.id),["prevent"])},{default:t(()=>[n(" 取消 ")]),_:2},1032,["onClick"]),e(i,{link:"",type:"primary",size:"small",onClick:_(v=>h(l.$index,l.row.id),["prevent"])},{default:t(()=>[n(" View Log ")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])]),_:1})]),_:1})]),_:1}),e($,{modelValue:r(p),"onUpdate:modelValue":a[1]||(a[1]=l=>M(p)?p.value=l:null),width:"80%",title:"LogViewer"},{footer:t(()=>[R("div",H,[e(i,{onClick:a[0]||(a[0]=l=>p.value=!1)},{default:t(()=>[n("关闭")]),_:1})])]),default:t(()=>[(g(),z(A,null,{default:t(()=>[e(r(b),{logfile:r(k)},null,8,["logfile"])]),fallback:t(()=>[e(r(G),{rows:4})]),_:1}))]),_:1},8,["modelValue"])],64)}}};export{we as default};