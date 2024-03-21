import{E,m as Be,a as Ae}from"./typescript-e92ae92a.js";import{v as U,w as ie,x as ze,d as He,r as Le,E as Te,i as q,y as Ce,z as De}from"./index-f7883cef.js";import{t as J}from"./aria-bc8e8b0f.js";import{d as H,u as w,a as z,b as ve,w as Q,i as O,R as Se,U as We,e as Re,T as Ee,_ as oe,c as f,L as Y,r as x,J as fe,Z as X,M as ce,o as he,a1 as we,ab as g,O as ke,V as je,F as be,$ as ge,a6 as D,a4 as Fe,X as Ve,a2 as qe,ad as Ge,G as Ue,f as ee,m as pe,g as te,p as Ze,t as Je,l as Xe,aa as Me}from"./index-5f1d5677.js";import{a as $e,C as Ke}from"./el-tooltip-3d978724.js";import{T as Qe,u as Ye}from"./index-fcde9027.js";import{b as ye,d as me}from"./runtime-89a56c03.js";import{t as ne}from"./error-78e43d3e.js";import{f as et}from"./vnode-98fc9e80.js";import{i as xe}from"./isNil-c75b1b34.js";const tt=H({name:"ElCollapseTransition"}),nt=H({...tt,setup(e){const o=w("collapse-transition"),n=t=>{t.style.maxHeight="",t.style.overflow=t.dataset.oldOverflow,t.style.paddingTop=t.dataset.oldPaddingTop,t.style.paddingBottom=t.dataset.oldPaddingBottom},i={beforeEnter(t){t.dataset||(t.dataset={}),t.dataset.oldPaddingTop=t.style.paddingTop,t.dataset.oldPaddingBottom=t.style.paddingBottom,t.style.height&&(t.dataset.elExistsHeight=t.style.height),t.style.maxHeight=0,t.style.paddingTop=0,t.style.paddingBottom=0},enter(t){requestAnimationFrame(()=>{t.dataset.oldOverflow=t.style.overflow,t.dataset.elExistsHeight?t.style.maxHeight=t.dataset.elExistsHeight:t.scrollHeight!==0?t.style.maxHeight=`${t.scrollHeight}px`:t.style.maxHeight=0,t.style.paddingTop=t.dataset.oldPaddingTop,t.style.paddingBottom=t.dataset.oldPaddingBottom,t.style.overflow="hidden"})},afterEnter(t){t.style.maxHeight="",t.style.overflow=t.dataset.oldOverflow},enterCancelled(t){n(t)},beforeLeave(t){t.dataset||(t.dataset={}),t.dataset.oldPaddingTop=t.style.paddingTop,t.dataset.oldPaddingBottom=t.style.paddingBottom,t.dataset.oldOverflow=t.style.overflow,t.style.maxHeight=`${t.scrollHeight}px`,t.style.overflow="hidden"},leave(t){t.scrollHeight!==0&&(t.style.maxHeight=0,t.style.paddingTop=0,t.style.paddingBottom=0)},afterLeave(t){n(t)},leaveCancelled(t){n(t)}};return(t,I)=>(z(),ve(Ee,Se({name:Re(o).b()},We(i)),{default:Q(()=>[O(t.$slots,"default")]),_:3},16,["name"]))}});var K=oe(nt,[["__file","collapse-transition.vue"]]);K.install=e=>{e.component(K.name,K)};const ot=K;let at=class{constructor(o,n){this.parent=o,this.domNode=n,this.subIndex=0,this.subIndex=0,this.init()}init(){this.subMenuItems=this.domNode.querySelectorAll("li"),this.addListeners()}gotoSubIndex(o){o===this.subMenuItems.length?o=0:o<0&&(o=this.subMenuItems.length-1),this.subMenuItems[o].focus(),this.subIndex=o}addListeners(){const o=this.parent.domNode;Array.prototype.forEach.call(this.subMenuItems,n=>{n.addEventListener("keydown",i=>{let t=!1;switch(i.code){case E.down:{this.gotoSubIndex(this.subIndex+1),t=!0;break}case E.up:{this.gotoSubIndex(this.subIndex-1),t=!0;break}case E.tab:{J(o,"mouseleave");break}case E.enter:case E.space:{t=!0,i.currentTarget.click();break}}return t&&(i.preventDefault(),i.stopPropagation()),!1})})}},st=class{constructor(o,n){this.domNode=o,this.submenu=null,this.submenu=null,this.init(n)}init(o){this.domNode.setAttribute("tabindex","0");const n=this.domNode.querySelector(`.${o}-menu`);n&&(this.submenu=new at(this,n)),this.addListeners()}addListeners(){this.domNode.addEventListener("keydown",o=>{let n=!1;switch(o.code){case E.down:{J(o.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(0),n=!0;break}case E.up:{J(o.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(this.submenu.subMenuItems.length-1),n=!0;break}case E.tab:{J(o.currentTarget,"mouseleave");break}case E.enter:case E.space:{n=!0,o.currentTarget.click();break}}n&&o.preventDefault()})}},lt=class{constructor(o,n){this.domNode=o,this.init(n)}init(o){const n=this.domNode.childNodes;Array.from(n).forEach(i=>{i.nodeType===1&&new st(i,o)})}};const ut=H({name:"ElMenuCollapseTransition",setup(){const e=w("menu");return{listeners:{onBeforeEnter:n=>n.style.opacity="0.2",onEnter(n,i){U(n,`${e.namespace.value}-opacity-transition`),n.style.opacity="1",i()},onAfterEnter(n){ie(n,`${e.namespace.value}-opacity-transition`),n.style.opacity=""},onBeforeLeave(n){n.dataset||(n.dataset={}),ze(n,e.m("collapse"))?(ie(n,e.m("collapse")),n.dataset.oldOverflow=n.style.overflow,n.dataset.scrollWidth=n.clientWidth.toString(),U(n,e.m("collapse"))):(U(n,e.m("collapse")),n.dataset.oldOverflow=n.style.overflow,n.dataset.scrollWidth=n.clientWidth.toString(),ie(n,e.m("collapse"))),n.style.width=`${n.scrollWidth}px`,n.style.overflow="hidden"},onLeave(n){U(n,"horizontal-collapse-transition"),n.style.width=`${n.dataset.scrollWidth}px`}}}}});function it(e,o,n,i,t,I){return z(),ve(Ee,Se({mode:"out-in"},e.listeners),{default:Q(()=>[O(e.$slots,"default")]),_:3},16)}var rt=oe(ut,[["render",it],["__file","menu-collapse-transition.vue"]]);function Oe(e,o){const n=f(()=>{let t=e.parent;const I=[o.value];for(;t.type.name!=="ElMenu";)t.props.index&&I.unshift(t.props.index),t=t.parent;return I});return{parentMenu:f(()=>{let t=e.parent;for(;t&&!["ElMenu","ElSubMenu"].includes(t.type.name);)t=t.parent;return t}),indexPath:n}}function dt(e){return f(()=>{const n=e.backgroundColor;return n?new Qe(n).shade(20).toString():""})}const _e=(e,o)=>{const n=w("menu");return f(()=>n.cssVarBlock({"text-color":e.textColor||"","hover-text-color":e.textColor||"","bg-color":e.backgroundColor||"","hover-bg-color":dt(e).value||"","active-color":e.activeTextColor||"",level:`${o}`}))},ct=ye({index:{type:String,required:!0},showTimeout:Number,hideTimeout:Number,popperClass:String,disabled:Boolean,popperAppendToBody:{type:Boolean,default:void 0},teleported:{type:Boolean,default:void 0},popperOffset:Number,expandCloseIcon:{type:q},expandOpenIcon:{type:q},collapseCloseIcon:{type:q},collapseOpenIcon:{type:q}}),Z="ElSubMenu";var Ie=H({name:Z,props:ct,setup(e,{slots:o,expose:n}){Ye({from:"popper-append-to-body",replacement:"teleported",scope:Z,version:"2.3.0",ref:"https://element-plus.org/en-US/component/menu.html#submenu-attributes"},f(()=>e.popperAppendToBody!==void 0));const i=ge(),{indexPath:t,parentMenu:I}=Oe(i,f(()=>e.index)),c=w("menu"),b=w("sub-menu"),u=Y("rootMenu");u||ne(Z,"can not inject root menu");const p=Y(`subMenu:${I.value.uid}`);p||ne(Z,"can not inject sub menu");const v=x({}),M=x({});let C;const k=x(!1),ae=x(),G=x(null),_=f(()=>N.value==="horizontal"&&P.value?"bottom-start":"right-start"),L=f(()=>N.value==="horizontal"&&P.value||N.value==="vertical"&&!u.props.collapse?e.expandCloseIcon&&e.expandOpenIcon?S.value?e.expandOpenIcon:e.expandCloseIcon:He:e.collapseCloseIcon&&e.collapseOpenIcon?S.value?e.collapseOpenIcon:e.collapseCloseIcon:Le),P=f(()=>p.level===0),W=f(()=>{var s;const d=(s=e.teleported)!=null?s:e.popperAppendToBody;return d===void 0?P.value:d}),se=f(()=>u.props.collapse?`${c.namespace.value}-zoom-in-left`:`${c.namespace.value}-zoom-in-top`),le=f(()=>N.value==="horizontal"&&P.value?["bottom-start","bottom-end","top-start","top-end","right-start","left-start"]:["right-start","right","right-end","left-start","bottom-start","bottom-end","top-start","top-end"]),S=f(()=>u.openedMenus.includes(e.index)),R=f(()=>{let s=!1;return Object.values(v.value).forEach(d=>{d.active&&(s=!0)}),Object.values(M.value).forEach(d=>{d.active&&(s=!0)}),s}),N=f(()=>u.props.mode),B=fe({index:e.index,indexPath:t,active:R}),j=_e(u.props,p.level+1),F=f(()=>{var s;return(s=e.popperOffset)!=null?s:u.props.popperOffset}),V=f(()=>{var s;return(s=e.popperClass)!=null?s:u.props.popperClass}),a=f(()=>{var s;return(s=e.showTimeout)!=null?s:u.props.showTimeout}),l=f(()=>{var s;return(s=e.hideTimeout)!=null?s:u.props.hideTimeout}),r=()=>{var s,d,m;return(m=(d=(s=G.value)==null?void 0:s.popperRef)==null?void 0:d.popperInstanceRef)==null?void 0:m.destroy()},y=s=>{s||r()},h=()=>{u.props.menuTrigger==="hover"&&u.props.mode==="horizontal"||u.props.collapse&&u.props.mode==="vertical"||e.disabled||u.handleSubMenuClick({index:e.index,indexPath:t.value,active:R.value})},T=(s,d=a.value)=>{var m;if(s.type!=="focus"){if(u.props.menuTrigger==="click"&&u.props.mode==="horizontal"||!u.props.collapse&&u.props.mode==="vertical"||e.disabled){p.mouseInChild.value=!0;return}p.mouseInChild.value=!0,C==null||C(),{stop:C}=Ce(()=>{u.openMenu(e.index,t.value)},d),W.value&&((m=I.value.vnode.el)==null||m.dispatchEvent(new MouseEvent("mouseenter")))}},$=(s=!1)=>{var d;if(u.props.menuTrigger==="click"&&u.props.mode==="horizontal"||!u.props.collapse&&u.props.mode==="vertical"){p.mouseInChild.value=!1;return}C==null||C(),p.mouseInChild.value=!1,{stop:C}=Ce(()=>!k.value&&u.closeMenu(e.index,t.value),l.value),W.value&&s&&((d=p.handleMouseleave)==null||d.call(p,!0))};X(()=>u.props.collapse,s=>y(!!s));{const s=m=>{M.value[m.index]=m},d=m=>{delete M.value[m.index]};ce(`subMenu:${i.uid}`,{addSubMenu:s,removeSubMenu:d,handleMouseleave:$,mouseInChild:k,level:p.level+1})}return n({opened:S}),he(()=>{u.addSubMenu(B),p.addSubMenu(B)}),we(()=>{p.removeSubMenu(B),u.removeSubMenu(B)}),()=>{var s;const d=[(s=o.title)==null?void 0:s.call(o),g(Te,{class:b.e("icon-arrow"),style:{transform:S.value?e.expandCloseIcon&&e.expandOpenIcon||e.collapseCloseIcon&&e.collapseOpenIcon&&u.props.collapse?"none":"rotateZ(180deg)":"none"}},{default:()=>D(L.value)?g(i.appContext.components[L.value]):g(L.value)})],m=u.isMenuPopup?g($e,{ref:G,visible:S.value,effect:"light",pure:!0,offset:F.value,showArrow:!1,persistent:!0,popperClass:V.value,placement:_.value,teleported:W.value,fallbackPlacements:le.value,transition:se.value,gpuAcceleration:!1},{content:()=>{var A;return g("div",{class:[c.m(N.value),c.m("popup-container"),V.value],onMouseenter:ue=>T(ue,100),onMouseleave:()=>$(!0),onFocus:ue=>T(ue,100)},[g("ul",{class:[c.b(),c.m("popup"),c.m(`popup-${_.value}`)],style:j.value},[(A=o.default)==null?void 0:A.call(o)])])},default:()=>g("div",{class:b.e("title"),onClick:h},d)}):g(be,{},[g("div",{class:b.e("title"),ref:ae,onClick:h},d),g(ot,{},{default:()=>{var A;return ke(g("ul",{role:"menu",class:[c.b(),c.m("inline")],style:j.value},[(A=o.default)==null?void 0:A.call(o)]),[[je,S.value]])}})]);return g("li",{class:[b.b(),b.is("active",R.value),b.is("opened",S.value),b.is("disabled",e.disabled)],role:"menuitem",ariaHaspopup:!0,ariaExpanded:S.value,onMouseenter:T,onMouseleave:()=>$(),onFocus:T},[m])}}});const pt=ye({mode:{type:String,values:["horizontal","vertical"],default:"vertical"},defaultActive:{type:String,default:""},defaultOpeneds:{type:me(Array),default:()=>Be([])},uniqueOpened:Boolean,router:Boolean,menuTrigger:{type:String,values:["hover","click"],default:"hover"},collapse:Boolean,backgroundColor:String,textColor:String,activeTextColor:String,closeOnClickOutside:Boolean,collapseTransition:{type:Boolean,default:!0},ellipsis:{type:Boolean,default:!0},popperOffset:{type:Number,default:6},ellipsisIcon:{type:q,default:()=>De},popperEffect:{type:String,values:["dark","light"],default:"dark"},popperClass:String,showTimeout:{type:Number,default:300},hideTimeout:{type:Number,default:300}}),re=e=>Array.isArray(e)&&e.every(o=>D(o)),mt={close:(e,o)=>D(e)&&re(o),open:(e,o)=>D(e)&&re(o),select:(e,o,n,i)=>D(e)&&re(o)&&Ve(n)&&(i===void 0||i instanceof Promise)};var vt=H({name:"ElMenu",props:pt,emits:mt,setup(e,{emit:o,slots:n,expose:i}){const t=ge(),I=t.appContext.config.globalProperties.$router,c=x(),b=w("menu"),u=w("sub-menu"),p=x(-1),v=x(e.defaultOpeneds&&!e.collapse?e.defaultOpeneds.slice(0):[]),M=x(e.defaultActive),C=x({}),k=x({}),ae=f(()=>e.mode==="horizontal"||e.mode==="vertical"&&e.collapse),G=()=>{const a=M.value&&C.value[M.value];if(!a||e.mode==="horizontal"||e.collapse)return;a.indexPath.forEach(r=>{const y=k.value[r];y&&_(r,y.indexPath)})},_=(a,l)=>{v.value.includes(a)||(e.uniqueOpened&&(v.value=v.value.filter(r=>l.includes(r))),v.value.push(a),o("open",a,l))},L=a=>{const l=v.value.indexOf(a);l!==-1&&v.value.splice(l,1)},P=(a,l)=>{L(a),o("close",a,l)},W=({index:a,indexPath:l})=>{v.value.includes(a)?P(a,l):_(a,l)},se=a=>{(e.mode==="horizontal"||e.collapse)&&(v.value=[]);const{index:l,indexPath:r}=a;if(!(xe(l)||xe(r)))if(e.router&&I){const y=a.route||l,h=I.push(y).then(T=>(T||(M.value=l),T));o("select",l,r,{index:l,indexPath:r,route:y},h)}else M.value=l,o("select",l,r,{index:l,indexPath:r})},le=a=>{const l=C.value,r=l[a]||M.value&&l[M.value]||l[e.defaultActive];r?M.value=r.index:M.value=a},S=()=>{var a,l;if(!c.value)return-1;const r=Array.from((l=(a=c.value)==null?void 0:a.childNodes)!=null?l:[]).filter(m=>m.nodeName!=="#comment"&&(m.nodeName!=="#text"||m.nodeValue)),y=64,h=Number.parseInt(getComputedStyle(c.value).paddingLeft,10),T=Number.parseInt(getComputedStyle(c.value).paddingRight,10),$=c.value.clientWidth-h-T;let s=0,d=0;return r.forEach((m,A)=>{s+=m.offsetWidth||0,s<=$-y&&(d=A+1)}),d===r.length?-1:d},R=a=>k.value[a].indexPath,N=(a,l=33.34)=>{let r;return()=>{r&&clearTimeout(r),r=setTimeout(()=>{a()},l)}};let B=!0;const j=()=>{const a=()=>{p.value=-1,qe(()=>{p.value=S()})};B?a():N(a)(),B=!1};X(()=>e.defaultActive,a=>{C.value[a]||(M.value=""),le(a)}),X(()=>e.collapse,a=>{a&&(v.value=[])}),X(C.value,G);let F;Fe(()=>{e.mode==="horizontal"&&e.ellipsis?F=Ae(c,j).stop:F==null||F()});const V=x(!1);{const a=h=>{k.value[h.index]=h},l=h=>{delete k.value[h.index]};ce("rootMenu",fe({props:e,openedMenus:v,items:C,subMenus:k,activeIndex:M,isMenuPopup:ae,addMenuItem:h=>{C.value[h.index]=h},removeMenuItem:h=>{delete C.value[h.index]},addSubMenu:a,removeSubMenu:l,openMenu:_,closeMenu:P,handleMenuItemClick:se,handleSubMenuClick:W})),ce(`subMenu:${t.uid}`,{addSubMenu:a,removeSubMenu:l,mouseInChild:V,level:0})}return he(()=>{e.mode==="horizontal"&&new lt(t.vnode.el,b.namespace.value)}),i({open:l=>{const{indexPath:r}=k.value[l];r.forEach(y=>_(y,r))},close:L,handleResize:j}),()=>{var a,l;let r=(l=(a=n.default)==null?void 0:a.call(n))!=null?l:[];const y=[];if(e.mode==="horizontal"&&c.value){const s=et(r),d=p.value===-1?s:s.slice(0,p.value),m=p.value===-1?[]:s.slice(p.value);m!=null&&m.length&&e.ellipsis&&(r=d,y.push(g(Ie,{index:"sub-menu-more",class:u.e("hide-arrow"),popperOffset:e.popperOffset},{title:()=>g(Te,{class:u.e("icon-more")},{default:()=>g(e.ellipsisIcon)}),default:()=>m})))}const h=_e(e,0),T=e.closeOnClickOutside?[[Ke,()=>{v.value.length&&(V.value||(v.value.forEach(s=>o("close",s,R(s))),v.value=[]))}]]:[],$=ke(g("ul",{key:String(e.collapse),role:"menubar",ref:c,style:h.value,class:{[b.b()]:!0,[b.m(e.mode)]:!0,[b.m("collapse")]:e.collapse}},[...r,...y]),T);return e.collapseTransition&&e.mode==="vertical"?g(rt,()=>$):$}}});const ft=ye({index:{type:me([String,null]),default:null},route:{type:me([String,Object])},disabled:Boolean}),ht={click:e=>D(e.index)&&Array.isArray(e.indexPath)},de="ElMenuItem",bt=H({name:de,components:{ElTooltip:$e},props:ft,emits:ht,setup(e,{emit:o}){const n=ge(),i=Y("rootMenu"),t=w("menu"),I=w("menu-item");i||ne(de,"can not inject root menu");const{parentMenu:c,indexPath:b}=Oe(n,Ge(e,"index")),u=Y(`subMenu:${c.value.uid}`);u||ne(de,"can not inject sub menu");const p=f(()=>e.index===i.activeIndex),v=fe({index:e.index,indexPath:b,active:p}),M=()=>{e.disabled||(i.handleMenuItemClick({index:e.index,indexPath:b.value,route:e.route}),o("click",v))};return he(()=>{u.addSubMenu(v),i.addMenuItem(v)}),we(()=>{u.removeSubMenu(v),i.removeMenuItem(v)}),{parentMenu:c,rootMenu:i,active:p,nsMenu:t,nsMenuItem:I,handleClick:M}}});function gt(e,o,n,i,t,I){const c=Ue("el-tooltip");return z(),ee("li",{class:te([e.nsMenuItem.b(),e.nsMenuItem.is("active",e.active),e.nsMenuItem.is("disabled",e.disabled)]),role:"menuitem",tabindex:"-1",onClick:o[0]||(o[0]=(...b)=>e.handleClick&&e.handleClick(...b))},[e.parentMenu.type.name==="ElMenu"&&e.rootMenu.props.collapse&&e.$slots.title?(z(),ve(c,{key:0,effect:e.rootMenu.props.popperEffect,placement:"right","fallback-placements":["left"],persistent:""},{content:Q(()=>[O(e.$slots,"title")]),default:Q(()=>[pe("div",{class:te(e.nsMenu.be("tooltip","trigger"))},[O(e.$slots,"default")],2)]),_:3},8,["effect"])):(z(),ee(be,{key:1},[O(e.$slots,"default"),O(e.$slots,"title")],64))],2)}var Pe=oe(bt,[["render",gt],["__file","menu-item.vue"]]);const Mt={title:String},yt="ElMenuItemGroup",It=H({name:yt,props:Mt,setup(){return{ns:w("menu-item-group")}}});function Ct(e,o,n,i,t,I){return z(),ee("li",{class:te(e.ns.b())},[pe("div",{class:te(e.ns.e("title"))},[e.$slots.title?O(e.$slots,"title",{key:1}):(z(),ee(be,{key:0},[Ze(Je(e.title),1)],64))],2),pe("ul",null,[O(e.$slots,"default")])],2)}var Ne=oe(It,[["render",Ct],["__file","menu-item-group.vue"]]);const zt=Xe(vt,{MenuItem:Pe,MenuItemGroup:Ne,SubMenu:Ie}),Ht=Me(Pe);Me(Ne);const Lt=Me(Ie);export{Ht as E,Lt as a,zt as b};