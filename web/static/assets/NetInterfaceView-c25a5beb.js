import{r as d,f as E,y as T,ao as B,o as a,j as u,S as V,l as p,m as o,B as l,V as m,aH as b,a6 as h,a7 as _,a5 as D,b1 as H}from"./index-fbdd90b6.js";/* empty css                   */import{E as L,a as N}from"./el-table-column-97e915f5.js";/* empty css               */import{E as k}from"./index-11674a5e.js";import"./index-73fe64b8.js";import"./event-2c073216.js";import"./isNil-c75b1b34.js";import"./use-form-item-61c38390.js";import"./use-form-common-props-82140e6e.js";import"./_initCloneObject-d6f68058.js";const z={class:"text-center font-size-3"},K={__name:"NetInterfaceView",setup(I){const r=d([]),s=d(!1),n=d(300),f=()=>{let e=document.body.clientHeight-215-140;e<300?n.value=300:n.value=e,console.log(e)},v=E(()=>r.value.length);T(()=>{w(),window.onresize=f});const w=()=>{s.value=!0,B({url:"/v1/server/netinterface_list",loading:!1}).then(e=>{r.value=e.data,f(),s.value=!1}).finally(()=>s.value=!1)};return(e,R)=>{const t=L,g=k,y=N,x=H;return a(),u(m,null,[V((a(),p(y,{data:r.value,border:"",style:{width:"100%"},height:n.value},{default:o(()=>[l(t,{prop:"index",label:"ID",width:"50"}),l(t,{prop:"name",label:"设备名称"}),l(t,{prop:"hardwareAddr",label:"硬件地址",width:"150"}),l(t,{prop:"mtu",label:"MTU",width:"150"}),l(t,{prop:"flags",label:"Flags"},{default:o(i=>[(a(!0),u(m,null,b(i.row.flags,c=>(a(),p(g,{class:"ml-2",type:"success"},{default:o(()=>[h(_(c),1)]),_:2},1024))),256))]),_:1}),l(t,{prop:"addrs",label:"IP 地址"},{default:o(i=>[(a(!0),u(m,null,b(i.row.addrs,c=>(a(),p(g,{class:"ml-2",type:"success"},{default:o(()=>[h(_(c.addr),1)]),_:2},1024))),256))]),_:1})]),_:1},8,["data","height"])),[[x,s.value]]),D("p",z,"网络设备数: "+_(v.value),1)],64)}}};export{K as default};
