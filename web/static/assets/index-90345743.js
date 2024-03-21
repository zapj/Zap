import{d as S,u as ae,c as d,a as p,f as M,i as U,j as x,w as N,O as j,m as _,g as m,e as n,t as q,V as A,T as Z,_ as G,l as re,aH as le,r as B,o as ie,Z as ue,b as h,n as ce,k as T,Q as de,F as pe,h as me,a6 as D,aq as K,K as $,ar as F,aI as fe}from"./index-5f1d5677.js";import{m as ge,u as ye,a as ve,E as Ce}from"./typescript-e92ae92a.js";import{b as Q,i as k,d as H,f as he}from"./runtime-89a56c03.js";import{i as be,j as J,G as V,E as P,T as Te,y as Ne}from"./index-f7883cef.js";import{e as Se}from"./use-global-config-71ca7aac.js";const R={},we=Q({value:{type:[String,Number],default:""},max:{type:Number,default:99},isDot:Boolean,hidden:Boolean,type:{type:String,values:["primary","success","warning","info","danger"],default:"danger"}}),Ee=["textContent"],Be=S({name:"ElBadge"}),Me=S({...Be,props:we,setup(s,{expose:t}){const e=s,o=ae("badge"),a=d(()=>e.isDot?"":k(e.value)&&k(e.max)?e.max<e.value?`${e.max}+`:`${e.value}`:`${e.value}`);return t({content:a}),(l,u)=>(p(),M("div",{class:m(n(o).b())},[U(l.$slots,"default"),x(Z,{name:`${n(o).namespace.value}-zoom-in-center`,persisted:""},{default:N(()=>[j(_("sup",{class:m([n(o).e("content"),n(o).em("content",l.type),n(o).is("fixed",!!l.$slots.default),n(o).is("dot",l.isDot)]),textContent:q(n(a))},null,10,Ee),[[A,!l.hidden&&(n(a)||l.isDot)]])]),_:1},8,["name"])],2))}});var _e=G(Me,[["__file","badge.vue"]]);const $e=re(_e),W=["success","info","warning","error"],i=ge({customClass:"",center:!1,dangerouslyUseHTMLString:!1,duration:3e3,icon:void 0,id:"",message:"",onClose:void 0,showClose:!1,type:"info",offset:16,zIndex:0,grouping:!1,repeatNum:1,appendTo:J?document.body:void 0}),ke=Q({customClass:{type:String,default:i.customClass},center:{type:Boolean,default:i.center},dangerouslyUseHTMLString:{type:Boolean,default:i.dangerouslyUseHTMLString},duration:{type:Number,default:i.duration},icon:{type:be,default:i.icon},id:{type:String,default:i.id},message:{type:H([String,Object,Function]),default:i.message},onClose:{type:H(Function),required:!1},showClose:{type:Boolean,default:i.showClose},type:{type:String,values:W,default:i.type},offset:{type:Number,default:i.offset},zIndex:{type:Number,default:i.zIndex},grouping:{type:Boolean,default:i.grouping},repeatNum:{type:Number,default:i.repeatNum}}),xe={destroy:()=>!0},c=le([]),Ie=s=>{const t=c.findIndex(a=>a.id===s),e=c[t];let o;return t>0&&(o=c[t-1]),{current:e,prev:o}},ze=s=>{const{prev:t}=Ie(s);return t?t.vm.exposed.bottom.value:0},Le=(s,t)=>c.findIndex(o=>o.id===s)>0?20:t,Oe=["id"],De=["innerHTML"],Fe=S({name:"ElMessage"}),He=S({...Fe,props:ke,emits:xe,setup(s,{expose:t}){const e=s,{Close:o}=Te,{ns:a,zIndex:l}=Se("message"),{currentZIndex:u,nextZIndex:f}=l,g=B(),v=B(!1),C=B(0);let w;const Y=d(()=>e.type?e.type==="error"?"danger":e.type:"info"),ee=d(()=>{const r=e.type;return{[a.bm("icon",r)]:r&&V[r]}}),I=d(()=>e.icon||V[e.type]||""),se=d(()=>ze(e.id)),z=d(()=>Le(e.id,e.offset)+se.value),te=d(()=>C.value+z.value),ne=d(()=>({top:`${z.value}px`,zIndex:u.value}));function E(){e.duration!==0&&({stop:w}=Ne(()=>{b()},e.duration))}function L(){w==null||w()}function b(){v.value=!1}function oe({code:r}){r===Ce.esc&&b()}return ie(()=>{E(),f(),v.value=!0}),ue(()=>e.repeatNum,()=>{L(),E()}),ye(document,"keydown",oe),ve(g,()=>{C.value=g.value.getBoundingClientRect().height}),t({visible:v,bottom:te,close:b}),(r,O)=>(p(),h(Z,{name:n(a).b("fade"),onBeforeLeave:r.onClose,onAfterLeave:O[0]||(O[0]=qe=>r.$emit("destroy")),persisted:""},{default:N(()=>[j(_("div",{id:r.id,ref_key:"messageRef",ref:g,class:m([n(a).b(),{[n(a).m(r.type)]:r.type},n(a).is("center",r.center),n(a).is("closable",r.showClose),r.customClass]),style:ce(n(ne)),role:"alert",onMouseenter:L,onMouseleave:E},[r.repeatNum>1?(p(),h(n($e),{key:0,value:r.repeatNum,type:n(Y),class:m(n(a).e("badge"))},null,8,["value","type","class"])):T("v-if",!0),n(I)?(p(),h(n(P),{key:1,class:m([n(a).e("icon"),n(ee)])},{default:N(()=>[(p(),h(de(n(I))))]),_:1},8,["class"])):T("v-if",!0),U(r.$slots,"default",{},()=>[r.dangerouslyUseHTMLString?(p(),M(pe,{key:1},[T(" Caution here, message could've been compromised, never use user's input as message "),_("p",{class:m(n(a).e("content")),innerHTML:r.message},null,10,De)],2112)):(p(),M("p",{key:0,class:m(n(a).e("content"))},q(r.message),3))]),r.showClose?(p(),h(n(P),{key:2,class:m(n(a).e("closeBtn")),onClick:me(b,["stop"])},{default:N(()=>[x(n(o))]),_:1},8,["class","onClick"])):T("v-if",!0)],46,Oe),[[A,v.value]])]),_:3},8,["name","onBeforeLeave"]))}});var Ve=G(He,[["__file","message.vue"]]);let Pe=1;const X=s=>{const t=!s||D(s)||K(s)||$(s)?{message:s}:s,e={...i,...t};if(!e.appendTo)e.appendTo=document.body;else if(D(e.appendTo)){let o=document.querySelector(e.appendTo);he(o)||(o=document.body),e.appendTo=o}return e},Re=s=>{const t=c.indexOf(s);if(t===-1)return;c.splice(t,1);const{handler:e}=s;e.close()},Ue=({appendTo:s,...t},e)=>{const o=`message_${Pe++}`,a=t.onClose,l=document.createElement("div"),u={...t,id:o,onClose:()=>{a==null||a(),Re(C)},onDestroy:()=>{F(null,l)}},f=x(Ve,u,$(u.message)||K(u.message)?{default:$(u.message)?u.message:()=>u.message}:null);f.appContext=e||y._context,F(f,l),s.appendChild(l.firstElementChild);const g=f.component,C={id:o,vnode:f,vm:g,handler:{close:()=>{g.exposed.visible.value=!1}},props:f.component.props};return C},y=(s={},t)=>{if(!J)return{close:()=>{}};if(k(R.max)&&c.length>=R.max)return{close:()=>{}};const e=X(s);if(e.grouping&&c.length){const a=c.find(({vnode:l})=>{var u;return((u=l.props)==null?void 0:u.message)===e.message});if(a)return a.props.repeatNum+=1,a.props.type=e.type,a.handler}const o=Ue(e,t);return c.push(o),o.handler};W.forEach(s=>{y[s]=(t={},e)=>{const o=X(t);return y({...o,type:s},e)}});function je(s){for(const t of c)(!s||s===t.props.type)&&t.handler.close()}y.closeAll=je;y._context=null;const Je=fe(y,"$message");export{Je as E};