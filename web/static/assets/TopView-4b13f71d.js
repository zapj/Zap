import{v as l,o as b,aC as c,z as m,A as d,Y as u,Z as _,P as e,E as h,ah as v}from"./index-c5a08050.js";import{E as f,a as g}from"./el-table-column-f533cbfe.js";/* empty css               */import"./index-28183222.js";import"./event-2c073216.js";import"./isNil-c75b1b34.js";import"./use-form-item-46d381f0.js";import"./use-form-common-props-2b0f30e3.js";import"./_initCloneObject-5f8c88a6.js";const E={class:"border-1 border-rd color-gray-500 font-size-3"},N={__name:"TopView",setup(w){const r=l(""),p=l([]),t=l(300),s=()=>{let a=document.body.clientHeight-260-140;a<300?t.value=300:t.value=a,console.log(a)};b(()=>{n(),window.onresize=s});const n=()=>{c({url:"/v1/server/top_info"}).then(a=>{r.value=a.data,p.value=a.rows,s()})};return(a,R)=>{const o=f,i=g;return m(),d(v,null,[u("pre",E,_(r.value),1),e(i,{data:p.value,style:{width:"100%"},height:t.value},{default:h(()=>[e(o,{prop:"0",label:"PID",width:"180",sortable:""}),e(o,{prop:"1",label:"USER",width:"180",sortable:""}),e(o,{prop:"2",label:"PR"}),e(o,{prop:"3",label:"NI"}),e(o,{prop:"4",label:"VIRT"}),e(o,{prop:"5",label:"RES",sortable:""}),e(o,{prop:"6",label:"SHR",sortable:""}),e(o,{prop:"7",label:"S"}),e(o,{prop:"8",label:"%CPU",sortable:""}),e(o,{prop:"9",label:"%MEM",sortable:""}),e(o,{prop:"10",label:"TIME+"}),e(o,{prop:"11",label:"COMMAND"})]),_:1},8,["data","height"])],64)}}};export{N as default};