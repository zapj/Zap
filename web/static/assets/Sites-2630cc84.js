import{b as pe,_ as ue,e as ce,E as G,at as me,f as D,r as k,u as fe,au as be,h as ge,a4 as L,o as p,l as H,B as t,m as o,S as J,a5 as w,F as we,A as U,s as g,j as h,q as I,a7 as S,D as E,U as ve,T as he,av as _e,t as ye,O as W,y as Ee,aw as Ce,ao as O,V as q,al as R,a6 as c,k as V,ax as ke,ay as ze,az as Ve,aa as M,aA as Se,aB as $e}from"./index-b976bc21.js";/* empty css                   */import{E as De,a as Le}from"./el-form-item-95c84d1d.js";/* empty css                 *//* empty css               */import{E as Ie,a as Re,b as Ue}from"./el-pagination-96d55ca8.js";import{E as Ae}from"./el-card-2216a0ff.js";import{E as Be,a as Te}from"./el-table-column-4403e15e.js";import{E as Fe}from"./el-empty-132ed155.js";import{E as We}from"./index-acc30c8b.js";import{d as Oe,a as He,u as Pe}from"./use-dialog-abcdb732.js";import{b as Ne}from"./index-3766283a.js";import{u as j,E as qe}from"./index-0d2c5633.js";import{E as Me,a as je,b as Ge}from"./index-53517561.js";import{E as Je}from"./index-a6754de2.js";import"./use-form-common-props-bf13ae6f.js";import"./castArray-aa1129c0.js";import"./event-2c073216.js";import"./use-form-item-014ad19a.js";import"./_initCloneObject-a305d225.js";import"./index-c52b2a81.js";import"./isNil-c75b1b34.js";const Ke=pe({...Oe,direction:{type:String,default:"rtl",values:["ltr","rtl","ttb","btt"]},size:{type:[String,Number],default:"30%"},withHeader:{type:Boolean,default:!0},modalFade:{type:Boolean,default:!0},headerAriaLevel:{type:String,default:"2"}}),Qe=He,Xe=ce({name:"ElDrawer",components:{ElOverlay:We,ElFocusTrap:Ne,ElIcon:G,Close:me},inheritAttrs:!1,props:Ke,emits:Qe,setup(e,{slots:d}){j({scope:"el-drawer",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/drawer.html#slots"},D(()=>!!d.title)),j({scope:"el-drawer",from:"custom-class",replacement:"class",version:"2.3.0",ref:"https://element-plus.org/en-US/component/drawer.html#attributes",type:"Attribute"},D(()=>!!e.customClass));const m=k(),f=k(),_=fe("drawer"),{t:n}=be(),u=D(()=>e.direction==="rtl"||e.direction==="ltr"),i=D(()=>ge(e.size));return{...Pe(e,m),drawerRef:m,focusStartRef:f,isHorizontal:u,drawerSize:i,ns:_,t:n}}}),Ye=["aria-label","aria-labelledby","aria-describedby"],Ze=["id","aria-level"],xe=["aria-label"],et=["id"];function tt(e,d,m,f,_,n){const u=L("close"),i=L("el-icon"),A=L("el-focus-trap"),B=L("el-overlay");return p(),H(_e,{to:"body",disabled:!e.appendToBody},[t(he,{name:e.ns.b("fade"),onAfterEnter:e.afterEnter,onAfterLeave:e.afterLeave,onBeforeLeave:e.beforeLeave,persisted:""},{default:o(()=>[J(t(B,{mask:e.modal,"overlay-class":e.modalClass,"z-index":e.zIndex,onClick:e.onModalClick},{default:o(()=>[t(A,{loop:"",trapped:e.visible,"focus-trap-el":e.drawerRef,"focus-start-el":e.focusStartRef,onReleaseRequested:e.onCloseRequested},{default:o(()=>[w("div",we({ref:"drawerRef","aria-modal":"true","aria-label":e.title||void 0,"aria-labelledby":e.title?void 0:e.titleId,"aria-describedby":e.bodyId},e.$attrs,{class:[e.ns.b(),e.direction,e.visible&&"open",e.customClass],style:e.isHorizontal?"width: "+e.drawerSize:"height: "+e.drawerSize,role:"dialog",onClick:d[1]||(d[1]=U(()=>{},["stop"]))}),[w("span",{ref:"focusStartRef",class:g(e.ns.e("sr-focus")),tabindex:"-1"},null,2),e.withHeader?(p(),h("header",{key:0,class:g(e.ns.e("header"))},[e.$slots.title?I(e.$slots,"title",{key:1},()=>[E(" DEPRECATED SLOT ")]):I(e.$slots,"header",{key:0,close:e.handleClose,titleId:e.titleId,titleClass:e.ns.e("title")},()=>[e.$slots.title?E("v-if",!0):(p(),h("span",{key:0,id:e.titleId,role:"heading","aria-level":e.headerAriaLevel,class:g(e.ns.e("title"))},S(e.title),11,Ze))]),e.showClose?(p(),h("button",{key:2,"aria-label":e.t("el.drawer.close"),class:g(e.ns.e("close-btn")),type:"button",onClick:d[0]||(d[0]=(...T)=>e.handleClose&&e.handleClose(...T))},[t(i,{class:g(e.ns.e("close"))},{default:o(()=>[t(u)]),_:1},8,["class"])],10,xe)):E("v-if",!0)],2)):E("v-if",!0),e.rendered?(p(),h("div",{key:1,id:e.bodyId,class:g(e.ns.e("body"))},[I(e.$slots,"default")],10,et)):E("v-if",!0),e.$slots.footer?(p(),h("div",{key:2,class:g(e.ns.e("footer"))},[I(e.$slots,"footer")],2)):E("v-if",!0)],16,Ye)]),_:3},8,["trapped","focus-trap-el","focus-start-el","onReleaseRequested"])]),_:3},8,["mask","overlay-class","z-index","onClick"]),[[ve,e.visible]])]),_:3},8,["name","onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["disabled"])}var ot=ue(Xe,[["render",tt],["__file","drawer.vue"]]);const at=ye(ot);const lt={class:"card-header font-size-3"},st={class:"el-dropdown-link"},nt=["id"],it={style:{float:"left"}},rt={style:{float:"right",color:"var(--el-text-color-secondary)"}},dt={style:{flex:"auto"}},C="140px",Rt={__name:"Sites",setup(e){const d=k(!1),m=k("50%"),f=k([]),_=k(400),n=W({domain:"",title:"",description:"",domain_names:"",www_root:"",run_directory:"",application:""}),u=W({config:{}}),i=W({total:0,pagesize:50,currentpage:1,tbLoading:!1}),A=()=>{n.www_root=n.domain.replaceAll("*.","")},B=()=>{O({url:"/v1/website/config",method:"get"}).then(s=>{if(s.code===0)u.config=s.data;else{R.error(s.msg),d.value=!1;return}console.log(u.config)})},T=()=>{console.log(n),console.log(u.config.applications[n.application])},K=s=>{R.success("打开网站")},Q=()=>{$({page:i.currentpage,pagesize:i.pagesize})},X=(s,l)=>{i.currentpage=s,i.pagesize=l,$({page:s,pagesize:l})},F=()=>{document.body.clientWidth<1200&&document.body.clientWidth>768?m.value="60%":document.body.clientWidth<768?m.value="100%":m.value="50%";let s=document.body.clientHeight-215-140;s<300?_.value=300:_.value=s};Ee(()=>{$({page:1,pagesize:i.pagesize}),window.addEventListener("resize",F)}),Ce(()=>{window.removeEventListener("resize",F)});function $(s){s=Object.assign({page:1,pagesize:50},s),i.tbLoading=!0,O({url:"/v1/website/list",method:"get",params:{path:s.path,page:s.page,limit:s.pagesize},dataType:"form",loading:!1}).then(r=>{f.value=[],f.value.push(...r.data),i.total=r.total,i.currentpage=r.page,F()}).catch(r=>{console.log(r)}).finally(()=>i.tbLoading=!1)}const Y=s=>{console.log(f.value[s]),O({url:"/v1/website/delete",method:"post",data:{id:f.value[s].ID},dataType:"form"}).then(l=>{l.code===0&&(R.success("网站删除成功"),$({page:1,pagesize:i.pagesize}))}).catch(l=>{R.error("网站删除失败")})};return(s,l)=>{const r=qe,v=Be,P=Me,Z=je,x=Ge,ee=Fe,te=Te,oe=Ie,ae=Ae,le=G,z=Je,y=De,se=Ue,ne=Re,ie=Le,re=at,de=ze;return p(),h(q,null,[t(ae,{class:"box-card filemanager"},{header:o(()=>[w("div",lt,[c(" 网站管理 "),t(r,{type:"primary",icon:V(Ve),onClick:l[0]||(l[0]=a=>d.value=!0)},null,8,["icon"])])]),footer:o(()=>[t(oe,{background:"",layout:"total,prev, pager, next",total:i.total,"page-size":i.pagesize,"current-page":i.currentpage,"onUpdate:currentPage":l[1]||(l[1]=a=>i.currentpage=a),"default-current-page":1,onChange:X},null,8,["total","page-size","current-page"])]),default:o(()=>[J((p(),H(te,{data:f.value,"min-height":"200",style:{width:"100%"},height:_.value,fit:""},{empty:o(()=>[t(ee,{description:"暂无数据"},{default:o(()=>[t(r,{type:"primary",onClick:Q},{default:o(()=>[c("刷新")]),_:1})]),_:1})]),default:o(()=>[t(v,{type:"selection",width:"30"}),t(v,{prop:"domain",label:"域名","show-overflow-tooltip":""},{default:o(({row:a,$index:b})=>[t(r,{link:"",onClick:U(N=>K(b),["prevent"])},{icon:o(()=>[t(V(M),{icon:"dashicons:admin-site"})]),default:o(()=>[c(" "+S(a.domain),1)]),_:2},1032,["onClick"])]),_:1}),t(v,{prop:"title",label:"网站名称"}),t(v,{prop:"description",label:"描述"}),t(v,{prop:"home",label:"站点目录"}),t(v,{prop:"status",label:"网站状态",width:"120"}),t(v,{fixed:"right",label:"操作",width:"100"},{default:o(a=>[t(r,{icon:V(Se),circle:"",class:"mr-2 hidden-md-and-down",title:"设置网站",onClick:U(b=>s.openFile(a.$index),["prevent"])},null,8,["icon","onClick"]),t(x,null,{dropdown:o(()=>[t(Z,null,{default:o(()=>[t(P,{onClick:U(b=>Y(a.$index),["prevent"])},{default:o(()=>[c("删除")]),_:2},1032,["onClick"]),t(P,null,{default:o(()=>[c("停止")]),_:1})]),_:2},1024)]),default:o(()=>[w("span",st,[t(V(M),{icon:"ri:more-fill",width:"24",height:"24"})])]),_:2},1024)]),_:1})]),_:1},8,["data","height"])),[[de,i.tbLoading]])]),_:1}),t(re,{modelValue:d.value,"onUpdate:modelValue":l[8]||(l[8]=a=>d.value=a),direction:"rtl",size:m.value,"show-close":!1,onOpen:B},{header:o(({close:a,titleId:b,titleClass:N})=>[w("h4",{id:b,class:g(N)},"添加网站",10,nt),t(r,{onClick:a},{default:o(()=>[t(le,{class:"el-icon--left"},{default:o(()=>[t(V($e))]),_:1}),c(" 关闭 ")]),_:2},1032,["onClick"])]),default:o(()=>[t(ie,{model:n,"label-position":"left"},{default:o(()=>[t(y,{label:"域名","label-width":C},{default:o(()=>[t(z,{modelValue:n.domain,"onUpdate:modelValue":l[2]||(l[2]=a=>n.domain=a),autocomplete:"off",onChange:A},null,8,["modelValue"])]),_:1}),t(y,{label:"附加域名","label-width":C},{default:o(()=>[t(z,{modelValue:n.domain_names,"onUpdate:modelValue":l[3]||(l[3]=a=>n.domain_names=a),autocomplete:"off",type:"textarea"},null,8,["modelValue"])]),_:1}),t(y,{label:"网站名称","label-width":C},{default:o(()=>[t(z,{modelValue:n.title,"onUpdate:modelValue":l[4]||(l[4]=a=>n.title=a),autocomplete:"off"},null,8,["modelValue"])]),_:1}),t(y,{label:"站点目录","label-width":C},{default:o(()=>[t(z,{modelValue:n.www_root,"onUpdate:modelValue":l[5]||(l[5]=a=>n.www_root=a),autocomplete:"off"},{prepend:o(()=>[c(S(u.config.home_dir)+"/",1)]),_:1},8,["modelValue"])]),_:1}),t(y,{label:"中间件","label-width":C},{default:o(()=>[t(ne,{modelValue:n.application,"onUpdate:modelValue":l[6]||(l[6]=a=>n.application=a),class:"m-2",placeholder:"Select",style:{width:"240px"}},{default:o(()=>[(p(!0),h(q,null,ke(u.config.applications,(a,b)=>(p(),H(se,{key:b,label:a.title,value:b},{default:o(()=>[w("span",it,S(a.title),1),w("span",rt,S(a.expose_port),1)]),_:2},1032,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),t(y,{label:"备注","label-width":C},{default:o(()=>[t(z,{modelValue:n.description,"onUpdate:modelValue":l[7]||(l[7]=a=>n.description=a),autocomplete:"off",type:"textarea"},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),footer:o(()=>[w("div",dt,[t(r,{onClick:s.cancelClick},{default:o(()=>[c("取消")]),_:1},8,["onClick"]),t(r,{type:"primary",onClick:T},{default:o(()=>[c("确认")]),_:1})])]),_:1},8,["modelValue","size"])],64)}}};export{Rt as default};
