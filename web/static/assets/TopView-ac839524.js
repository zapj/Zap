import{r as l,o as b,a as m,f as c,K as d,M as u,j as e,w as _,F as h}from"./index-57ba148c.js";import{E as f,a as g}from"./el-table-column-28499b9c.js";import"./el-tooltip-54d9aabc.js";/* empty css               */import{a as v}from"./client-157b369f.js";import"./error-78e43d3e.js";import"./use-form-common-props-7f97df7e.js";import"./runtime-14a2c257.js";import"./_initCloneObject-0341bd43.js";import"./index-74d71456.js";import"./index-def65ea4.js";import"./event-8e91c63d.js";import"./use-form-item-c8f82fc4.js";import"./isNil-c75b1b34.js";import"./index-407d4441.js";const w={class:"border-1 border-rd color-gray-500 font-size-3"},F={__name:"TopView",setup(E){const r=l(""),p=l([]),a=l(300),s=()=>{let t=document.body.clientHeight-260-140;t<300?a.value=300:a.value=t,console.log(t)};b(()=>{i(),window.onresize=s});const i=()=>{v({url:"/v1/server/top_info"}).then(t=>{r.value=t.data,p.value=t.rows,s()})};return(t,M)=>{const o=f,n=g;return m(),c(h,null,[d("pre",w,u(r.value),1),e(n,{data:p.value,style:{width:"100%"},height:a.value},{default:_(()=>[e(o,{prop:"0",label:"PID",width:"180",sortable:""}),e(o,{prop:"1",label:"USER",width:"180",sortable:""}),e(o,{prop:"2",label:"PR"}),e(o,{prop:"3",label:"NI"}),e(o,{prop:"4",label:"VIRT"}),e(o,{prop:"5",label:"RES",sortable:""}),e(o,{prop:"6",label:"SHR",sortable:""}),e(o,{prop:"7",label:"S"}),e(o,{prop:"8",label:"%CPU",sortable:""}),e(o,{prop:"9",label:"%MEM",sortable:""}),e(o,{prop:"10",label:"TIME+"}),e(o,{prop:"11",label:"COMMAND"})]),_:1},8,["data","height"])],64)}}};export{F as default};