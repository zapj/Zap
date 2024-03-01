import{b as g,e as f,u as _,f as n,P as R,o as $,l as h,m as v,q as w,s as j,k as c,n as N,p as x,_ as C,t as E,d as u,Y as r,L as S,i as b,a0 as K}from"./index-cc2a4353.js";const O=Symbol("rowContextKey"),P=["start","center","end","space-around","space-between","space-evenly"],k=["top","middle","bottom"],L=g({tag:{type:String,default:"div"},gutter:{type:Number,default:0},justify:{type:String,values:P,default:"start"},align:{type:String,values:k}}),B=f({name:"ElRow"}),q=f({...B,props:L,setup(p){const e=p,l=_("row"),a=n(()=>e.gutter);R(O,{gutter:a});const i=n(()=>{const t={};return e.gutter&&(t.marginRight=t.marginLeft=`-${e.gutter/2}px`),t}),d=n(()=>[l.b(),l.is(`justify-${e.justify}`,e.justify!=="start"),l.is(`align-${e.align}`,!!e.align)]);return(t,m)=>($(),h(x(t.tag),{class:j(c(d)),style:N(c(i))},{default:v(()=>[w(t.$slots,"default")]),_:3},8,["class","style"]))}});var A=C(q,[["__file","row.vue"]]);const G=E(A),D=g({tag:{type:String,default:"div"},span:{type:Number,default:24},offset:{type:Number,default:0},pull:{type:Number,default:0},push:{type:Number,default:0},xs:{type:u([Number,Object]),default:()=>r({})},sm:{type:u([Number,Object]),default:()=>r({})},md:{type:u([Number,Object]),default:()=>r({})},lg:{type:u([Number,Object]),default:()=>r({})},xl:{type:u([Number,Object]),default:()=>r({})}}),I=f({name:"ElCol"}),J=f({...I,props:D,setup(p){const e=p,{gutter:l}=S(O,{gutter:n(()=>0)}),a=_("col"),i=n(()=>{const t={};return l.value&&(t.paddingLeft=t.paddingRight=`${l.value/2}px`),t}),d=n(()=>{const t=[];return["span","offset","pull","push"].forEach(s=>{const o=e[s];b(o)&&(s==="span"?t.push(a.b(`${e[s]}`)):o>0&&t.push(a.b(`${s}-${e[s]}`)))}),["xs","sm","md","lg","xl"].forEach(s=>{b(e[s])?t.push(a.b(`${s}-${e[s]}`)):K(e[s])&&Object.entries(e[s]).forEach(([o,y])=>{t.push(o!=="span"?a.b(`${s}-${o}-${y}`):a.b(`${s}-${y}`))})}),l.value&&t.push(a.is("guttered")),[a.b(),t]});return(t,m)=>($(),h(x(t.tag),{class:j(c(d)),style:N(c(i))},{default:v(()=>[w(t.$slots,"default")]),_:3},8,["class","style"]))}});var T=C(J,[["__file","col.vue"]]);const H=E(T);export{H as E,G as a};