import{E as me}from"./index-06fc50a2.js";import{c as L,d as Le,u as ue,a as Pe,E as Ke,O as De,w as ce,F as Me}from"./index-c3ba0162.js";import{_ as M,e as P,q as S,r as C,P as J,L as h,y as Ae,k as I,Q as ve,b as X,d as D,H as g,f as $,b3 as Ge,a3 as W,w as ge,x as ze,a4 as E,o as T,l as z,m as w,B as R,b4 as Ye,b5 as Ue,a2 as He,a as be,E as ne,M as Je,u as x,aJ as Ve,h as je,j as V,b6 as We,F as q,s as Q,D as H,W as we,a5 as qe,p as Ie,A as oe,V as Qe,n as Xe,t as Ee,a8 as _e}from"./index-a62a7a28.js";import{u as ye}from"./use-form-item-453acac3.js";import{u as Ze}from"./use-form-common-props-9c2de981.js";import{c as xe}from"./castArray-c6dfbd2f.js";import{c as Ce}from"./index-818b1577.js";const eo=P({inheritAttrs:!1});function oo(e,o,s,t,d,a){return S(e.$slots,"default")}var no=M(eo,[["render",oo],["__file","collection.vue"]]);const to=P({name:"ElCollectionItem",inheritAttrs:!1});function lo(e,o,s,t,d,a){return S(e.$slots,"default")}var ro=M(to,[["render",lo],["__file","collection-item.vue"]]);const he="data-el-collection-item",$e=e=>{const o=`El${e}Collection`,s=`${o}Item`,t=Symbol(o),d=Symbol(s),a={...no,name:o,setup(){const r=C(null),p=new Map;J(t,{itemMap:p,getItems:()=>{const m=I(r);if(!m)return[];const f=Array.from(m.querySelectorAll(`[${he}]`));return[...p.values()].sort((l,c)=>f.indexOf(l.ref)-f.indexOf(c.ref))},collectionRef:r})}},n={...ro,name:s,setup(r,{attrs:p}){const v=C(null),m=h(t,void 0);J(d,{collectionItemRef:v}),Ae(()=>{const f=I(v);f&&m.itemMap.set(f,{ref:f,...p})}),ve(()=>{const f=I(v);m.itemMap.delete(f)})}};return{COLLECTION_INJECTION_KEY:t,COLLECTION_ITEM_INJECTION_KEY:d,ElCollection:a,ElCollectionItem:n}},so=X({style:{type:D([String,Array,Object])},currentTabId:{type:D(String)},defaultCurrentTabId:String,loop:Boolean,dir:{type:String,values:["ltr","rtl"],default:"ltr"},orientation:{type:D(String)},onBlur:Function,onFocus:Function,onMousedown:Function}),{ElCollection:io,ElCollectionItem:ao,COLLECTION_INJECTION_KEY:te,COLLECTION_ITEM_INJECTION_KEY:uo}=$e("RovingFocusGroup"),le=Symbol("elRovingFocusGroup"),Te=Symbol("elRovingFocusGroupItem"),co={ArrowLeft:"prev",ArrowUp:"prev",ArrowRight:"next",ArrowDown:"next",PageUp:"first",Home:"first",PageDown:"last",End:"last"},po=(e,o)=>{if(o!=="rtl")return e;switch(e){case g.right:return g.left;case g.left:return g.right;default:return e}},fo=(e,o,s)=>{const t=po(e.key,s);if(!(o==="vertical"&&[g.left,g.right].includes(t))&&!(o==="horizontal"&&[g.up,g.down].includes(t)))return co[t]},mo=(e,o)=>e.map((s,t)=>e[(t+o)%e.length]),re=e=>{const{activeElement:o}=document;for(const s of e)if(s===o||(s.focus(),o!==document.activeElement))return},pe="currentTabIdChange",fe="rovingFocusGroup.entryFocus",vo={bubbles:!1,cancelable:!0},go=P({name:"ElRovingFocusGroupImpl",inheritAttrs:!1,props:so,emits:[pe,"entryFocus"],setup(e,{emit:o}){var s;const t=C((s=e.currentTabId||e.defaultCurrentTabId)!=null?s:null),d=C(!1),a=C(!1),n=C(null),{getItems:r}=h(te,void 0),p=$(()=>[{outline:"none"},e.style]),v=u=>{o(pe,u)},m=()=>{d.value=!0},f=L(u=>{var b;(b=e.onMousedown)==null||b.call(e,u)},()=>{a.value=!0}),_=L(u=>{var b;(b=e.onFocus)==null||b.call(e,u)},u=>{const b=!I(a),{target:A,currentTarget:N}=u;if(A===N&&b&&!I(d)){const G=new Event(fe,vo);if(N==null||N.dispatchEvent(G),!G.defaultPrevented){const y=r().filter(k=>k.focusable),O=y.find(k=>k.active),F=y.find(k=>k.id===I(t)),Y=[O,F,...y].filter(Boolean).map(k=>k.ref);re(Y)}}a.value=!1}),l=L(u=>{var b;(b=e.onBlur)==null||b.call(e,u)},()=>{d.value=!1}),c=(...u)=>{o("entryFocus",...u)};J(le,{currentTabbedId:Ge(t),loop:W(e,"loop"),tabIndex:$(()=>I(d)?-1:0),rovingFocusGroupRef:n,rovingFocusGroupRootStyle:p,orientation:W(e,"orientation"),dir:W(e,"dir"),onItemFocus:v,onItemShiftTab:m,onBlur:l,onFocus:_,onMousedown:f}),ge(()=>e.currentTabId,u=>{t.value=u??null}),ze(n,fe,c)}});function bo(e,o,s,t,d,a){return S(e.$slots,"default")}var wo=M(go,[["render",bo],["__file","roving-focus-group-impl.vue"]]);const Io=P({name:"ElRovingFocusGroup",components:{ElFocusGroupCollection:io,ElRovingFocusGroupImpl:wo}});function Eo(e,o,s,t,d,a){const n=E("el-roving-focus-group-impl"),r=E("el-focus-group-collection");return T(),z(r,null,{default:w(()=>[R(n,Ye(Ue(e.$attrs)),{default:w(()=>[S(e.$slots,"default")]),_:3},16)]),_:3})}var _o=M(Io,[["render",Eo],["__file","roving-focus-group.vue"]]);const yo=P({components:{ElRovingFocusCollectionItem:ao},props:{focusable:{type:Boolean,default:!0},active:{type:Boolean,default:!1}},emits:["mousedown","focus","keydown"],setup(e,{emit:o}){const{currentTabbedId:s,loop:t,onItemFocus:d,onItemShiftTab:a}=h(le,void 0),{getItems:n}=h(te,void 0),r=ye(),p=C(null),v=L(l=>{o("mousedown",l)},l=>{e.focusable?d(I(r)):l.preventDefault()}),m=L(l=>{o("focus",l)},()=>{d(I(r))}),f=L(l=>{o("keydown",l)},l=>{const{key:c,shiftKey:u,target:b,currentTarget:A}=l;if(c===g.tab&&u){a();return}if(b!==A)return;const N=fo(l);if(N){l.preventDefault();let y=n().filter(O=>O.focusable).map(O=>O.ref);switch(N){case"last":{y.reverse();break}case"prev":case"next":{N==="prev"&&y.reverse();const O=y.indexOf(A);y=t.value?mo(y,O+1):y.slice(O+1);break}}He(()=>{re(y)})}}),_=$(()=>s.value===I(r));return J(Te,{rovingFocusGroupItemRef:p,tabIndex:$(()=>I(_)?0:-1),handleMousedown:v,handleFocus:m,handleKeydown:f}),{id:r,handleKeydown:f,handleFocus:m,handleMousedown:v}}});function Co(e,o,s,t,d,a){const n=E("el-roving-focus-collection-item");return T(),z(n,{id:e.id,focusable:e.focusable,active:e.active},{default:w(()=>[S(e.$slots,"default")]),_:3},8,["id","focusable","active"])}var ho=M(yo,[["render",Co],["__file","roving-focus-item.vue"]]);const $o=X({trigger:Le.trigger,effect:{...ue.effect,default:"light"},type:{type:D(String)},placement:{type:D(String),default:"bottom"},popperOptions:{type:D(Object),default:()=>({})},id:String,size:{type:String,default:""},splitButton:Boolean,hideOnClick:{type:Boolean,default:!0},loop:{type:Boolean,default:!0},showTimeout:{type:Number,default:150},hideTimeout:{type:Number,default:150},tabindex:{type:D([Number,String]),default:0},maxHeight:{type:D([Number,String]),default:""},popperClass:{type:String,default:""},disabled:{type:Boolean,default:!1},role:{type:String,default:"menu"},buttonProps:{type:D(Object)},teleported:ue.teleported}),Oe=X({command:{type:[Object,String,Number],default:()=>({})},disabled:Boolean,divided:Boolean,textValue:String,icon:{type:be}}),To=X({onKeydown:{type:D(Function)}}),Oo=[g.down,g.pageDown,g.home],Se=[g.up,g.pageUp,g.end],So=[...Oo,...Se],{ElCollection:Fo,ElCollectionItem:ko,COLLECTION_INJECTION_KEY:Ro,COLLECTION_ITEM_INJECTION_KEY:No}=$e("Dropdown"),ee=Symbol("elDropdown"),{ButtonGroup:Bo}=me,Lo=P({name:"ElDropdown",components:{ElButton:me,ElButtonGroup:Bo,ElScrollbar:Pe,ElDropdownCollection:Fo,ElTooltip:Ke,ElRovingFocusGroup:_o,ElOnlyChild:De,ElIcon:ne,ArrowDown:Je},props:$o,emits:["visible-change","click","command"],setup(e,{emit:o}){const s=we(),t=x("dropdown"),{t:d}=Ve(),a=C(),n=C(),r=C(null),p=C(null),v=C(null),m=C(null),f=C(!1),_=[g.enter,g.space,g.down],l=$(()=>({maxHeight:je(e.maxHeight)})),c=$(()=>[t.m(O.value)]),u=$(()=>xe(e.trigger)),b=ye().value,A=$(()=>e.id||b);ge([a,u],([i,B],[U])=>{var ie,ae,de;(ie=U==null?void 0:U.$el)!=null&&ie.removeEventListener&&U.$el.removeEventListener("pointerenter",K),(ae=i==null?void 0:i.$el)!=null&&ae.removeEventListener&&i.$el.removeEventListener("pointerenter",K),(de=i==null?void 0:i.$el)!=null&&de.addEventListener&&B.includes("hover")&&i.$el.addEventListener("pointerenter",K)},{immediate:!0}),ve(()=>{var i,B;(B=(i=a.value)==null?void 0:i.$el)!=null&&B.removeEventListener&&a.value.$el.removeEventListener("pointerenter",K)});function N(){G()}function G(){var i;(i=r.value)==null||i.onClose()}function y(){var i;(i=r.value)==null||i.onOpen()}const O=Ze();function F(...i){o("command",...i)}function K(){var i,B;(B=(i=a.value)==null?void 0:i.$el)==null||B.focus()}function Y(){}function k(){const i=I(p);u.value.includes("hover")&&(i==null||i.focus()),m.value=null}function se(i){m.value=i}function Z(i){f.value||(i.preventDefault(),i.stopImmediatePropagation())}function j(){o("visible-change",!0)}function Ne(i){(i==null?void 0:i.type)==="keydown"&&p.value.focus()}function Be(){o("visible-change",!1)}return J(ee,{contentRef:p,role:$(()=>e.role),triggerId:A,isUsingKeyboard:f,onItemEnter:Y,onItemLeave:k}),J("elDropdown",{instance:s,dropdownSize:O,handleClick:N,commandHandler:F,trigger:W(e,"trigger"),hideOnClick:W(e,"hideOnClick")}),{t:d,ns:t,scrollbar:v,wrapStyle:l,dropdownTriggerKls:c,dropdownSize:O,triggerId:A,triggerKeys:_,currentTabId:m,handleCurrentTabIdChange:se,handlerMainButtonClick:i=>{o("click",i)},handleEntryFocus:Z,handleClose:G,handleOpen:y,handleBeforeShowTooltip:j,handleShowTooltip:Ne,handleBeforeHideTooltip:Be,onFocusAfterTrapped:i=>{var B,U;i.preventDefault(),(U=(B=p.value)==null?void 0:B.focus)==null||U.call(B,{preventScroll:!0})},popperRef:r,contentRef:p,triggeringElementRef:a,referenceElementRef:n}}});function Po(e,o,s,t,d,a){var n;const r=E("el-dropdown-collection"),p=E("el-roving-focus-group"),v=E("el-scrollbar"),m=E("el-only-child"),f=E("el-tooltip"),_=E("el-button"),l=E("arrow-down"),c=E("el-icon"),u=E("el-button-group");return T(),V("div",{class:Q([e.ns.b(),e.ns.is("disabled",e.disabled)])},[R(f,{ref:"popperRef",role:e.role,effect:e.effect,"fallback-placements":["bottom","top"],"popper-options":e.popperOptions,"gpu-acceleration":!1,"hide-after":e.trigger==="hover"?e.hideTimeout:0,"manual-mode":!0,placement:e.placement,"popper-class":[e.ns.e("popper"),e.popperClass],"reference-element":(n=e.referenceElementRef)==null?void 0:n.$el,trigger:e.trigger,"trigger-keys":e.triggerKeys,"trigger-target-el":e.contentRef,"show-after":e.trigger==="hover"?e.showTimeout:0,"stop-popper-mouse-event":!1,"virtual-ref":e.triggeringElementRef,"virtual-triggering":e.splitButton,disabled:e.disabled,transition:`${e.ns.namespace.value}-zoom-in-top`,teleported:e.teleported,pure:"",persistent:"",onBeforeShow:e.handleBeforeShowTooltip,onShow:e.handleShowTooltip,onBeforeHide:e.handleBeforeHideTooltip},We({content:w(()=>[R(v,{ref:"scrollbar","wrap-style":e.wrapStyle,tag:"div","view-class":e.ns.e("list")},{default:w(()=>[R(p,{loop:e.loop,"current-tab-id":e.currentTabId,orientation:"horizontal",onCurrentTabIdChange:e.handleCurrentTabIdChange,onEntryFocus:e.handleEntryFocus},{default:w(()=>[R(r,null,{default:w(()=>[S(e.$slots,"dropdown")]),_:3})]),_:3},8,["loop","current-tab-id","onCurrentTabIdChange","onEntryFocus"])]),_:3},8,["wrap-style","view-class"])]),_:2},[e.splitButton?void 0:{name:"default",fn:w(()=>[R(m,{id:e.triggerId,ref:"triggeringElementRef",role:"button",tabindex:e.tabindex},{default:w(()=>[S(e.$slots,"default")]),_:3},8,["id","tabindex"])])}]),1032,["role","effect","popper-options","hide-after","placement","popper-class","reference-element","trigger","trigger-keys","trigger-target-el","show-after","virtual-ref","virtual-triggering","disabled","transition","teleported","onBeforeShow","onShow","onBeforeHide"]),e.splitButton?(T(),z(u,{key:0},{default:w(()=>[R(_,q({ref:"referenceElementRef"},e.buttonProps,{size:e.dropdownSize,type:e.type,disabled:e.disabled,tabindex:e.tabindex,onClick:e.handlerMainButtonClick}),{default:w(()=>[S(e.$slots,"default")]),_:3},16,["size","type","disabled","tabindex","onClick"]),R(_,q({id:e.triggerId,ref:"triggeringElementRef"},e.buttonProps,{role:"button",size:e.dropdownSize,type:e.type,class:e.ns.e("caret-button"),disabled:e.disabled,tabindex:e.tabindex,"aria-label":e.t("el.dropdown.toggleDropdown")}),{default:w(()=>[R(c,{class:Q(e.ns.e("icon"))},{default:w(()=>[R(l)]),_:1},8,["class"])]),_:1},16,["id","size","type","class","disabled","tabindex","aria-label"])]),_:3})):H("v-if",!0)],2)}var Ko=M(Lo,[["render",Po],["__file","dropdown.vue"]]);const Do=P({name:"DropdownItemImpl",components:{ElIcon:ne},props:Oe,emits:["pointermove","pointerleave","click","clickimpl"],setup(e,{emit:o}){const s=x("dropdown"),{role:t}=h(ee,void 0),{collectionItemRef:d}=h(No,void 0),{collectionItemRef:a}=h(uo,void 0),{rovingFocusGroupItemRef:n,tabIndex:r,handleFocus:p,handleKeydown:v,handleMousedown:m}=h(Te,void 0),f=Ce(d,a,n),_=$(()=>t.value==="menu"?"menuitem":t.value==="navigation"?"link":"button"),l=L(c=>{const{code:u}=c;if(u===g.enter||u===g.space)return c.preventDefault(),c.stopImmediatePropagation(),o("clickimpl",c),!0},v);return{ns:s,itemRef:f,dataset:{[he]:""},role:_,tabIndex:r,handleFocus:p,handleKeydown:l,handleMousedown:m}}}),Mo=["aria-disabled","tabindex","role"];function Ao(e,o,s,t,d,a){const n=E("el-icon");return T(),V(Qe,null,[e.divided?(T(),V("li",q({key:0,role:"separator",class:e.ns.bem("menu","item","divided")},e.$attrs),null,16)):H("v-if",!0),qe("li",q({ref:e.itemRef},{...e.dataset,...e.$attrs},{"aria-disabled":e.disabled,class:[e.ns.be("menu","item"),e.ns.is("disabled",e.disabled)],tabindex:e.tabIndex,role:e.role,onClick:o[0]||(o[0]=r=>e.$emit("clickimpl",r)),onFocus:o[1]||(o[1]=(...r)=>e.handleFocus&&e.handleFocus(...r)),onKeydown:o[2]||(o[2]=oe((...r)=>e.handleKeydown&&e.handleKeydown(...r),["self"])),onMousedown:o[3]||(o[3]=(...r)=>e.handleMousedown&&e.handleMousedown(...r)),onPointermove:o[4]||(o[4]=r=>e.$emit("pointermove",r)),onPointerleave:o[5]||(o[5]=r=>e.$emit("pointerleave",r))}),[e.icon?(T(),z(n,{key:0},{default:w(()=>[(T(),z(Ie(e.icon)))]),_:1})):H("v-if",!0),S(e.$slots,"default")],16,Mo)],64)}var Go=M(Do,[["render",Ao],["__file","dropdown-item-impl.vue"]]);const Fe=()=>{const e=h("elDropdown",{}),o=$(()=>e==null?void 0:e.dropdownSize);return{elDropdown:e,_elDropdownSize:o}},zo=P({name:"ElDropdownItem",components:{ElDropdownCollectionItem:ko,ElRovingFocusItem:ho,ElDropdownItemImpl:Go},inheritAttrs:!1,props:Oe,emits:["pointermove","pointerleave","click"],setup(e,{emit:o,attrs:s}){const{elDropdown:t}=Fe(),d=we(),a=C(null),n=$(()=>{var l,c;return(c=(l=I(a))==null?void 0:l.textContent)!=null?c:""}),{onItemEnter:r,onItemLeave:p}=h(ee,void 0),v=L(l=>(o("pointermove",l),l.defaultPrevented),ce(l=>{if(e.disabled){p(l);return}const c=l.currentTarget;c===document.activeElement||c.contains(document.activeElement)||(r(l),l.defaultPrevented||c==null||c.focus())})),m=L(l=>(o("pointerleave",l),l.defaultPrevented),ce(l=>{p(l)})),f=L(l=>{if(!e.disabled)return o("click",l),l.type!=="keydown"&&l.defaultPrevented},l=>{var c,u,b;if(e.disabled){l.stopImmediatePropagation();return}(c=t==null?void 0:t.hideOnClick)!=null&&c.value&&((u=t.handleClick)==null||u.call(t)),(b=t.commandHandler)==null||b.call(t,e.command,d,l)}),_=$(()=>({...e,...s}));return{handleClick:f,handlePointerMove:v,handlePointerLeave:m,textContent:n,propsAndAttrs:_}}});function Yo(e,o,s,t,d,a){var n;const r=E("el-dropdown-item-impl"),p=E("el-roving-focus-item"),v=E("el-dropdown-collection-item");return T(),z(v,{disabled:e.disabled,"text-value":(n=e.textValue)!=null?n:e.textContent},{default:w(()=>[R(p,{focusable:!e.disabled},{default:w(()=>[R(r,q(e.propsAndAttrs,{onPointerleave:e.handlePointerLeave,onPointermove:e.handlePointerMove,onClickimpl:e.handleClick}),{default:w(()=>[S(e.$slots,"default")]),_:3},16,["onPointerleave","onPointermove","onClickimpl"])]),_:3},8,["focusable"])]),_:3},8,["disabled","text-value"])}var ke=M(zo,[["render",Yo],["__file","dropdown-item.vue"]]);const Uo=P({name:"ElDropdownMenu",props:To,setup(e){const o=x("dropdown"),{_elDropdownSize:s}=Fe(),t=s.value,{focusTrapRef:d,onKeydown:a}=h(Me,void 0),{contentRef:n,role:r,triggerId:p}=h(ee,void 0),{collectionRef:v,getItems:m}=h(Ro,void 0),{rovingFocusGroupRef:f,rovingFocusGroupRootStyle:_,tabIndex:l,onBlur:c,onFocus:u,onMousedown:b}=h(le,void 0),{collectionRef:A}=h(te,void 0),N=$(()=>[o.b("menu"),o.bm("menu",t==null?void 0:t.value)]),G=Ce(n,v,d,f,A),y=L(F=>{var K;(K=e.onKeydown)==null||K.call(e,F)},F=>{const{currentTarget:K,code:Y,target:k}=F;if(K.contains(k),g.tab===Y&&F.stopImmediatePropagation(),F.preventDefault(),k!==I(n)||!So.includes(Y))return;const Z=m().filter(j=>!j.disabled).map(j=>j.ref);Se.includes(Y)&&Z.reverse(),re(Z)});return{size:t,rovingFocusGroupRootStyle:_,tabIndex:l,dropdownKls:N,role:r,triggerId:p,dropdownListWrapperRef:G,handleKeydown:F=>{y(F),a(F)},onBlur:c,onFocus:u,onMousedown:b}}}),Ho=["role","aria-labelledby"];function Jo(e,o,s,t,d,a){return T(),V("ul",{ref:e.dropdownListWrapperRef,class:Q(e.dropdownKls),style:Xe(e.rovingFocusGroupRootStyle),tabindex:-1,role:e.role,"aria-labelledby":e.triggerId,onBlur:o[0]||(o[0]=(...n)=>e.onBlur&&e.onBlur(...n)),onFocus:o[1]||(o[1]=(...n)=>e.onFocus&&e.onFocus(...n)),onKeydown:o[2]||(o[2]=oe((...n)=>e.handleKeydown&&e.handleKeydown(...n),["self"])),onMousedown:o[3]||(o[3]=oe((...n)=>e.onMousedown&&e.onMousedown(...n),["self"]))},[S(e.$slots,"default")],46,Ho)}var Re=M(Uo,[["render",Jo],["__file","dropdown-menu.vue"]]);const an=Ee(Ko,{DropdownItem:ke,DropdownMenu:Re}),dn=_e(ke),un=_e(Re),Vo=X({type:{type:String,values:["primary","success","warning","info","danger","default"],default:"default"},underline:{type:Boolean,default:!0},disabled:{type:Boolean,default:!1},href:{type:String,default:""},target:{type:String,default:"_self"},icon:{type:be}}),jo={click:e=>e instanceof MouseEvent},Wo=["href","target"],qo=P({name:"ElLink"}),Qo=P({...qo,props:Vo,emits:jo,setup(e,{emit:o}){const s=e,t=x("link"),d=$(()=>[t.b(),t.m(s.type),t.is("disabled",s.disabled),t.is("underline",s.underline&&!s.disabled)]);function a(n){s.disabled||o("click",n)}return(n,r)=>(T(),V("a",{class:Q(I(d)),href:n.disabled||!n.href?void 0:n.href,target:n.disabled||!n.href?void 0:n.target,onClick:a},[n.icon?(T(),z(I(ne),{key:0},{default:w(()=>[(T(),z(Ie(n.icon)))]),_:1})):H("v-if",!0),n.$slots.default?(T(),V("span",{key:1,class:Q(I(t).e("inner"))},[S(n.$slots,"default")],2)):H("v-if",!0),n.$slots.icon?S(n.$slots,"icon",{key:2}):H("v-if",!0)],10,Wo))}});var Xo=M(Qo,[["__file","link.vue"]]);const cn=Ee(Xo);export{cn as E,dn as a,un as b,an as c};
