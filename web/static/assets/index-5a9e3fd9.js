import{e as E,au as G,L as M,f as v,o as f,j as I,a5 as h,q as r,s as t,k as e,a7 as H,B as k,m as n,l as T,p as Q,E as W,D as w,n as N,_ as z,bu as X,aM as Z,u as x,r as S,P as ee,S as oe,F as se,ba as ae,U as le,T as te,av as ne,t as re}from"./index-b976bc21.js";import{E as ie,a as de}from"./index-acc30c8b.js";import{b as ce,e as ue,c as fe,d as pe,a as me,u as ge}from"./use-dialog-abcdb732.js";import{F as ve,b as be}from"./index-3766283a.js";import{u as ye}from"./index-107051b1.js";import{u as P}from"./index-0d2c5633.js";const j=Symbol("dialogInjectionKey"),Ce=["aria-level"],he=["aria-label"],ke=["id"],Ee=E({name:"ElDialogContent"}),Re=E({...Ee,props:ce,emits:ue,setup(R){const l=R,{t:i}=G(),{Close:$}=X,{dialogRef:d,headerRef:p,bodyId:A,ns:s,style:m}=M(j),{focusTrapRef:b}=M(ve),y=v(()=>[s.b(),s.is("fullscreen",l.fullscreen),s.is("draggable",l.draggable),s.is("align-center",l.alignCenter),{[s.m("center")]:l.center},l.customClass]),D=fe(b,d),F=v(()=>l.draggable);return ye(d,p,F),(a,C)=>(f(),I("div",{ref:e(D),class:t(e(y)),style:N(e(m)),tabindex:"-1"},[h("header",{ref_key:"headerRef",ref:p,class:t(e(s).e("header"))},[r(a.$slots,"header",{},()=>[h("span",{role:"heading","aria-level":a.ariaLevel,class:t(e(s).e("title"))},H(a.title),11,Ce)]),a.showClose?(f(),I("button",{key:0,"aria-label":e(i)("el.dialog.close"),class:t(e(s).e("headerbtn")),type:"button",onClick:C[0]||(C[0]=L=>a.$emit("close"))},[k(e(W),{class:t(e(s).e("close"))},{default:n(()=>[(f(),T(Q(a.closeIcon||e($))))]),_:1},8,["class"])],10,he)):w("v-if",!0)],2),h("div",{id:e(A),class:t(e(s).e("body"))},[r(a.$slots,"default")],10,ke),a.$slots.footer?(f(),I("footer",{key:0,class:t(e(s).e("footer"))},[r(a.$slots,"footer")],2)):w("v-if",!0)],6))}});var $e=z(Re,[["__file","dialog-content.vue"]]);const Ae=["aria-label","aria-labelledby","aria-describedby"],De=E({name:"ElDialog",inheritAttrs:!1}),Fe=E({...De,props:pe,emits:me,setup(R,{expose:l}){const i=R,$=Z();P({scope:"el-dialog",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/dialog.html#slots"},v(()=>!!$.title)),P({scope:"el-dialog",from:"custom-class",replacement:"class",version:"2.3.0",ref:"https://element-plus.org/en-US/component/dialog.html#attributes",type:"Attribute"},v(()=>!!i.customClass));const d=x("dialog"),p=S(),A=S(),s=S(),{visible:m,titleId:b,bodyId:y,style:D,overlayDialogStyle:F,rendered:a,zIndex:C,afterEnter:L,afterLeave:q,beforeLeave:K,handleClose:B,onModalClick:O,onOpenAutoFocus:U,onCloseAutoFocus:V,onCloseRequested:_,onFocusoutPrevented:J}=ge(i,p);ee(j,{dialogRef:p,headerRef:A,bodyId:y,ns:d,rendered:a,style:D});const c=de(O),Y=v(()=>i.draggable&&!i.fullscreen);return l({visible:m,dialogContentRef:s}),(o,u)=>(f(),T(ne,{to:o.appendTo,disabled:o.appendTo!=="body"?!1:!o.appendToBody},[k(te,{name:"dialog-fade",onAfterEnter:e(L),onAfterLeave:e(q),onBeforeLeave:e(K),persisted:""},{default:n(()=>[oe(k(e(ie),{"custom-mask-event":"",mask:o.modal,"overlay-class":o.modalClass,"z-index":e(C)},{default:n(()=>[h("div",{role:"dialog","aria-modal":"true","aria-label":o.title||void 0,"aria-labelledby":o.title?void 0:e(b),"aria-describedby":e(y),class:t(`${e(d).namespace.value}-overlay-dialog`),style:N(e(F)),onClick:u[0]||(u[0]=(...g)=>e(c).onClick&&e(c).onClick(...g)),onMousedown:u[1]||(u[1]=(...g)=>e(c).onMousedown&&e(c).onMousedown(...g)),onMouseup:u[2]||(u[2]=(...g)=>e(c).onMouseup&&e(c).onMouseup(...g))},[k(e(be),{loop:"",trapped:e(m),"focus-start-el":"container",onFocusAfterTrapped:e(U),onFocusAfterReleased:e(V),onFocusoutPrevented:e(J),onReleaseRequested:e(_)},{default:n(()=>[e(a)?(f(),T($e,se({key:0,ref_key:"dialogContentRef",ref:s},o.$attrs,{"custom-class":o.customClass,center:o.center,"align-center":o.alignCenter,"close-icon":o.closeIcon,draggable:e(Y),fullscreen:o.fullscreen,"show-close":o.showClose,title:o.title,"aria-level":o.headerAriaLevel,onClose:e(B)}),ae({header:n(()=>[o.$slots.title?r(o.$slots,"title",{key:1}):r(o.$slots,"header",{key:0,close:e(B),titleId:e(b),titleClass:e(d).e("title")})]),default:n(()=>[r(o.$slots,"default")]),_:2},[o.$slots.footer?{name:"footer",fn:n(()=>[r(o.$slots,"footer")])}:void 0]),1040,["custom-class","center","align-center","close-icon","draggable","fullscreen","show-close","title","aria-level","onClose"])):w("v-if",!0)]),_:3},8,["trapped","onFocusAfterTrapped","onFocusAfterReleased","onFocusoutPrevented","onReleaseRequested"])],46,Ae)]),_:3},8,["mask","overlay-class","z-index"]),[[le,e(m)]])]),_:3},8,["onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["to","disabled"]))}});var Ie=z(Fe,[["__file","dialog.vue"]]);const Pe=re(Ie);export{Pe as E};
