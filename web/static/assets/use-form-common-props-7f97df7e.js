import{A as d,c as r,D as f,p as n,e as i,r as b}from"./index-57ba148c.js";import{a as p}from"./runtime-14a2c257.js";const C=e=>e===void 0,g=e=>typeof e=="boolean",x=e=>typeof e=="number",E=e=>typeof Element>"u"?!1:e instanceof Element,K=e=>d(e)?!Number.isNaN(Number(e)):!1,v=["","default","small","large"],u=e=>{const o=f();return r(()=>{var s,t;return(t=(s=o==null?void 0:o.proxy)==null?void 0:s.$props)==null?void 0:t[e]})},_=p({type:String,values:v,required:!1}),y=Symbol("size"),z=()=>{const e=n(y,{});return r(()=>i(e.size)||"")},m=Symbol("formContextKey"),S=Symbol("formItemContextKey"),j=(e,o={})=>{const s=b(void 0),t=o.prop?s:u("size"),c=o.global?s:z(),a=o.form?{size:void 0}:n(m,void 0),l=o.formItem?{size:void 0}:n(S,void 0);return r(()=>t.value||i(e)||(l==null?void 0:l.size)||(a==null?void 0:a.size)||c.value||"")},P=e=>{const o=u("disabled"),s=n(m,void 0);return r(()=>o.value||i(e)||(s==null?void 0:s.disabled)||!1)};export{y as S,x as a,_ as b,v as c,g as d,P as e,m as f,S as g,E as h,C as i,K as j,j as u};