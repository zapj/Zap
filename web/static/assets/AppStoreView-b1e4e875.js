import{e as f,r as V,v as E,y as w,o as _,j as A,B as o,m as i,a5 as B,l as C,p as R,V as b,ad as s,ai as l}from"./index-a62a7a28.js";import{E as h}from"./el-card-85a77e75.js";import{E as I,a as L}from"./el-radio-button-606dd4b4.js";import"./event-2c073216.js";import"./use-form-common-props-9c2de981.js";import"./use-form-item-453acac3.js";const k={class:"mt-3"},S=f({__name:"AppStoreView",setup(x){const n=s(()=>l(()=>import("./AppListView-f939ae89.js"),["assets/AppListView-f939ae89.js","assets/index-a62a7a28.js","assets/index-8e7b09c6.css","assets/el-card-85a77e75.js","assets/el-card-7b6ad268.css","assets/index-aee6e995.js","assets/use-form-common-props-9c2de981.js","assets/index-06fc50a2.js","assets/use-form-item-453acac3.js","assets/index-35c6f9af.js","assets/el-tag-263a511e.css"])),c=s(()=>l(()=>import("./AppInstalledView-eab17e53.js"),["assets/AppInstalledView-eab17e53.js","assets/index-a62a7a28.js","assets/index-8e7b09c6.css","assets/el-table-column-8104c127.js","assets/index-c3ba0162.js","assets/event-2c073216.js","assets/isNil-c75b1b34.js","assets/use-form-item-453acac3.js","assets/use-form-common-props-9c2de981.js","assets/_initCloneObject-7450ca73.js","assets/el-table-column-f582ce8b.css","assets/index-06fc50a2.js","assets/index-6ab50f87.js","assets/index-07ca6cb1.js","assets/aria-bc8e8b0f.js","assets/el-tag-263a511e.css"])),p=s(()=>l(()=>import("./TaskListView-fa43f7b8.js"),["assets/TaskListView-fa43f7b8.js","assets/index-a62a7a28.js","assets/index-8e7b09c6.css","assets/el-card-85a77e75.js","assets/el-card-7b6ad268.css","assets/el-table-column-8104c127.js","assets/index-c3ba0162.js","assets/event-2c073216.js","assets/isNil-c75b1b34.js","assets/use-form-item-453acac3.js","assets/use-form-common-props-9c2de981.js","assets/_initCloneObject-7450ca73.js","assets/el-table-column-f582ce8b.css","assets/commons-d455d4a5.js","assets/index-1a3cb46c.js","assets/index-06fc50a2.js","assets/index-35c6f9af.js","assets/index-818b1577.js","assets/index-07ca6cb1.js","assets/el-tag-263a511e.css"])),e=V("应用市场"),a=E(n),m=()=>{e.value==="应用市场"?a.value=n:e.value==="已安装"?a.value=c:e.value==="任务列表"&&(a.value=p)};return w(()=>{}),(y,r)=>{const t=I,u=L,d=h;return _(),A(b,null,[o(d,{class:"box-card"},{default:i(()=>[o(u,{modelValue:e.value,"onUpdate:modelValue":r[0]||(r[0]=v=>e.value=v),type:"success",onChange:m},{default:i(()=>[o(t,{label:"应用市场"}),o(t,{label:"已安装"}),o(t,{label:"任务列表"})]),_:1},8,["modelValue"])]),_:1}),B("div",k,[(_(),C(R(a.value)))])],64)}}});export{S as default};