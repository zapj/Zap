import{r as v,y as C,ao as b,o as t,l as n,m as o,j as _,V as d,aT as p,k as w,B as A,a5 as a,a7 as r,A as B,a6 as u}from"./index-c1703e63.js";import{E as V}from"./el-card-df563978.js";import{E as L}from"./index-de19a356.js";import{E as N,a as R}from"./index-4db19fab.js";import"./use-form-item-4c026880.js";import"./use-form-common-props-bc166b1a.js";const T={style:{padding:"14px"}},M=["textContent"],j={class:"bottom"},q={class:"time font-size-3"},J={__name:"AppListView",setup(z){const l=v([]);C(()=>{m()});const m=()=>{b({url:"/v1/app/appstore"}).then(s=>{l.value=s.data})};function f(s,c){console.log(s,c)}return(s,c)=>{const i=L,h=V,k=N,x=R;return t(),n(x,{gutter:20},{default:o(()=>[(t(!0),_(d,null,p(w(l),(e,y)=>(t(),n(k,{span:6,class:"mb-2",key:y},{default:o(()=>[A(h,{"body-style":{padding:"0px"},shadow:"hover"},{default:o(()=>[a("div",T,[a("span",{textContent:r(e.Title)},null,8,M),a("div",j,[a("time",q,r(e.Version),1),e.Actions?(t(!0),_(d,{key:0},p(e.Actions,(E,g)=>(t(),n(i,{class:"button",onClick:B(D=>f(e.Id,g),["prevent"])},{default:o(()=>[u(r(E),1)]),_:2},1032,["onClick"]))),256)):(t(),n(i,{key:1,class:"button"},{default:o(()=>[u("安装")]),_:1}))])])]),_:2},1024)]),_:2},1024))),128))]),_:1})}}};export{J as default};