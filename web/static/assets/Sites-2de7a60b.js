import{_ as _e,d as ve,c as U,r as v,u as he,J as R,a as f,b as P,j as o,w as i,B as K,K as h,m as ye,h as A,g as _,f as z,i as B,M as I,k as L,C as Ee,T as Ce,a1 as ke,l as ze,q as O,o as Ve,a2 as $e,F as N,L as g,e as F,a3 as Se}from"./index-ef788864.js";/* empty css                   */import{d as De,a as Le,u as Ie}from"./el-overlay-5305a6db.js";import{E as Te,a as Fe}from"./el-form-item-632ca7fb.js";/* empty css                 *//* empty css               */import{E as Ae,a as Ue,b as Re}from"./el-pagination-0f3fe061.js";import{b as Be}from"./el-tooltip-6781d3c2.js";/* empty css                */import{E as We}from"./el-card-8494bd30.js";import{E as qe,a as Oe}from"./el-table-column-ecd9c7ec.js";import{E as Pe}from"./el-empty-7bcdd05f.js";import{u as j,E as He}from"./el-button-470da2a3.js";/* empty css                        */import{E as Me,a as Ne,b as je}from"./el-dropdown-item-b6895a84.js";import{a as C,v as Je}from"./client-16215366.js";import{f as Ke,p as Ge,g as Qe,h as Xe}from"./index-17346e61.js";import{I as J}from"./iconify-4b1e5988.js";import{_ as Ye}from"./_plugin-vue_export-helper-c27b6911.js";import{E as Ze}from"./index-b5f3f95a.js";import{E as G,c as xe,d as et}from"./index-499c28ec.js";import{b as tt}from"./runtime-6c5e7d4b.js";import{b as m}from"./index-faf3eb96.js";import{E as ot}from"./index-0c2eda1a.js";import"./event-8e91c63d.js";import"./use-form-common-props-f4525b0b.js";import"./use-form-item-26fbbe24.js";import"./castArray-8ff7c110.js";import"./error-78e43d3e.js";import"./_initCloneObject-ea0126de.js";import"./index-55ed1832.js";import"./isNil-c75b1b34.js";const at=tt({...De,direction:{type:String,default:"rtl",values:["ltr","rtl","ttb","btt"]},size:{type:[String,Number],default:"30%"},withHeader:{type:Boolean,default:!0},modalFade:{type:Boolean,default:!0},headerAriaLevel:{type:String,default:"2"}}),lt=Le,it=ve({name:"ElDrawer",components:{ElOverlay:Ze,ElFocusTrap:Be,ElIcon:G,Close:Ke},inheritAttrs:!1,props:at,emits:lt,setup(e,{slots:p}){j({scope:"el-drawer",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/drawer.html#slots"},U(()=>!!p.title)),j({scope:"el-drawer",from:"custom-class",replacement:"class",version:"2.3.0",ref:"https://element-plus.org/en-US/component/drawer.html#attributes",type:"Attribute"},U(()=>!!e.customClass));const c=v(),V=v(),y=he("drawer"),{t:u}=xe(),b=U(()=>e.direction==="rtl"||e.direction==="ltr"),$=U(()=>et(e.size));return{...Ie(e,c),drawerRef:c,focusStartRef:V,isHorizontal:b,drawerSize:$,ns:y,t:u}}}),st=["aria-label","aria-labelledby","aria-describedby"],rt=["id","aria-level"],nt=["aria-label"],dt=["id"];function pt(e,p,c,V,y,u){const b=R("close"),$=R("el-icon"),a=R("el-focus-trap"),w=R("el-overlay");return f(),P(ke,{to:"body",disabled:!e.appendToBody},[o(Ce,{name:e.ns.b("fade"),onAfterEnter:e.afterEnter,onAfterLeave:e.afterLeave,onBeforeLeave:e.beforeLeave,persisted:""},{default:i(()=>[K(o(w,{mask:e.modal,"overlay-class":e.modalClass,"z-index":e.zIndex,onClick:e.onModalClick},{default:i(()=>[o(a,{loop:"",trapped:e.visible,"focus-trap-el":e.drawerRef,"focus-start-el":e.focusStartRef,onReleaseRequested:e.onCloseRequested},{default:i(()=>[h("div",ye({ref:"drawerRef","aria-modal":"true","aria-label":e.title||void 0,"aria-labelledby":e.title?void 0:e.titleId,"aria-describedby":e.bodyId},e.$attrs,{class:[e.ns.b(),e.direction,e.visible&&"open",e.customClass],style:e.isHorizontal?"width: "+e.drawerSize:"height: "+e.drawerSize,role:"dialog",onClick:p[1]||(p[1]=A(()=>{},["stop"]))}),[h("span",{ref:"focusStartRef",class:_(e.ns.e("sr-focus")),tabindex:"-1"},null,2),e.withHeader?(f(),z("header",{key:0,class:_(e.ns.e("header"))},[e.$slots.title?B(e.$slots,"title",{key:1},()=>[L(" DEPRECATED SLOT ")]):B(e.$slots,"header",{key:0,close:e.handleClose,titleId:e.titleId,titleClass:e.ns.e("title")},()=>[e.$slots.title?L("v-if",!0):(f(),z("span",{key:0,id:e.titleId,role:"heading","aria-level":e.headerAriaLevel,class:_(e.ns.e("title"))},I(e.title),11,rt))]),e.showClose?(f(),z("button",{key:2,"aria-label":e.t("el.drawer.close"),class:_(e.ns.e("close-btn")),type:"button",onClick:p[0]||(p[0]=(...n)=>e.handleClose&&e.handleClose(...n))},[o($,{class:_(e.ns.e("close"))},{default:i(()=>[o(b)]),_:1},8,["class"])],10,nt)):L("v-if",!0)],2)):L("v-if",!0),e.rendered?(f(),z("div",{key:1,id:e.bodyId,class:_(e.ns.e("body"))},[B(e.$slots,"default")],10,dt)):L("v-if",!0),e.$slots.footer?(f(),z("div",{key:2,class:_(e.ns.e("footer"))},[B(e.$slots,"footer")],2)):L("v-if",!0)],16,st)]),_:3},8,["trapped","focus-trap-el","focus-start-el","onReleaseRequested"])]),_:3},8,["mask","overlay-class","z-index","onClick"]),[[Ee,e.visible]])]),_:3},8,["name","onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["disabled"])}var ut=_e(it,[["render",pt],["__file","drawer.vue"]]);const mt=ze(ut);const ft={class:"card-header font-size-3"},ct={class:"el-dropdown-link"},wt=["id"],gt={style:{float:"left"}},bt={style:{float:"right",color:"var(--el-text-color-secondary)"}},_t={style:{flex:"auto"}},k="140px",vt={__name:"Sites",setup(e){let p=!1;const c=v([]),V=v(400),y=v(),u=v(!1),b=v("50%"),$=v("添加网站"),a=O({domain:"",title:"",description:"",domain_names:"",www_root:"",run_directory:"",application:0,website_id:0,index_files:"index.html index.htm index.php"}),w=O({config:{}}),n=O({total:0,pagesize:50,currentpage:1,tbLoading:!1}),Q={domain:[{required:!0,trigger:"blur",validator:(l,t,r)=>{t==""&&t===void 0&&r(new Error("域名不能为空")),C({url:"v1/website/check_domain",dataType:"form",data:{domain:t,website_id:a.website_id},method:"POST",loading:!1}).then(d=>{d.code===0?r():r(new Error(d.msg))}).catch(d=>{r(new Error("域名验证失败"))})}},{min:3,max:20,message:"长度在 3 到 20 个字符",trigger:"blur"}],title:[{required:!0,message:"请输入网站名称",trigger:"blur"}],www_root:[{required:!0,message:"请输入站点目录",trigger:"blur"}],application:[{required:!0,message:"请选择应用类型",trigger:"change"}]};Ve(()=>{T({page:1,pagesize:n.pagesize}),window.addEventListener("resize",W)}),$e(()=>{window.removeEventListener("resize",W)});const X=()=>{a.www_root===""&&(a.www_root=a.domain.replaceAll("*.","")),a.title===""&&(a.title=a.domain.replaceAll("www.",""))},Y=()=>C({url:"/v1/website/config",method:"get"}).then(l=>{if(l.code===0)w.config=l.data;else{m.error(l.msg),u.value=!1;return}}),Z=()=>{C({url:"/v1/website/stop",method:"post"}).then(l=>{l.code===0?m.success("停止成功"):m.error(l.msg)})},H=l=>{p=!0,u.value=!0,$.value="网站设置",a.website_id=l,C({url:"/v1/website/settings",method:"get",params:{id:l}}).then(t=>{if(t.code===0)w.config=t.data.config,a.domain=t.data.website.domain,a.domain_names=t.data.website.domain_names,a.www_root=t.data.website.www_root,a.run_directory=t.data.website.run_directory,a.application=t.data.website.application_id,a.title=t.data.website.title,a.description=t.data.website.description,a.website_id=t.data.website.ID,a.www_root===""?a.www_root=a.domain.replaceAll("*.",""):a.www_root=t.data.website.www_root.replace(w.config.home_dir+"/",""),a.title===""&&(a.title=a.domain.replaceAll("www.",""));else{m.error(t.msg),u.value=!1;return}})},x=l=>{l&&(l.resetFields(),u.value=!1)},ee=l=>{Y().finally(()=>{u.value=!0}),p=!1,a.website_id=0,l&&l.resetFields()},te=l=>{if(!l)return;l.validate(d=>{if(!d)return!1});let t=Object.assign({},a);if(w.config.applications.length===0){m.error("请先添加应用");return}let r=null;if(w.config.applications.forEach((d,q)=>{d.id===a.application&&(r=d)}),!r){m.error("应用不存在");return}t.expose_port=r.expose_port,t.expose_proto=r.expose_proto,t.application=r.id,C({url:p?"/v1/website/update":"/v1/website/create",method:"post",data:t}).then(d=>{d.code===0?(m.success(p?"更新成功":"添加成功"),T({page:1,pagesize:n.pagesize})):m.error(d.msg)}).finally(()=>{u.value=!1,l.resetFields()})},oe=l=>{H(l)},ae=()=>{T({page:n.currentpage,pagesize:n.pagesize})},le=(l,t)=>{n.currentpage=l,n.pagesize=t,T({page:l,pagesize:t})},W=()=>{document.body.clientWidth<1200&&document.body.clientWidth>768?b.value="60%":document.body.clientWidth<768?b.value="100%":b.value="50%";let l=document.body.clientHeight-215-140;l<300?V.value=300:V.value=l};function T(l){l=Object.assign({page:1,pagesize:50},l),n.tbLoading=!0,C({url:"/v1/website/list",method:"get",params:{path:l.path,page:l.page,limit:l.pagesize},dataType:"form",loading:!1}).then(r=>{c.value=[],c.value.push(...r.data),n.total=r.total,n.currentpage=r.page,W()}).catch(r=>{console.log(r)}).finally(()=>n.tbLoading=!1)}const ie=l=>{console.log(c.value[l]),C({url:"/v1/website/delete",method:"post",data:{id:c.value[l].ID},dataType:"form"}).then(t=>{t.code===0&&(m.success("网站删除成功"),T({page:1,pagesize:n.pagesize}))}).catch(t=>{m.error("网站删除失败")})};return(l,t)=>{const r=He,d=qe,q=Me,se=Ne,re=je,ne=Pe,de=Oe,pe=Ae,ue=We,me=G,S=ot,E=Te,fe=Re,ce=Ue,we=Fe,ge=mt,be=Je;return f(),z(N,null,[o(ue,{class:"box-card filemanager","body-style":"padding:0"},{header:i(()=>[h("div",ft,[g(" 网站管理 "),o(r,{type:"success",icon:F(Ge),onClick:t[0]||(t[0]=s=>ee(y.value))},null,8,["icon"])])]),footer:i(()=>[o(pe,{background:"",layout:"total,prev, pager, next",total:n.total,"page-size":n.pagesize,"current-page":n.currentpage,"onUpdate:currentPage":t[1]||(t[1]=s=>n.currentpage=s),"default-current-page":1,onChange:le},null,8,["total","page-size","current-page"])]),default:i(()=>[K((f(),P(de,{data:c.value,"min-height":"300",style:{width:"100%"},height:V.value,fit:""},{empty:i(()=>[o(ne,{description:"暂无数据"},{default:i(()=>[o(r,{type:"primary",onClick:ae},{default:i(()=>[g("刷新")]),_:1})]),_:1})]),default:i(()=>[o(d,{type:"selection",width:"30"}),o(d,{prop:"domain",label:"域名","show-overflow-tooltip":""},{default:i(({row:s,$index:D})=>[o(r,{link:"",onClick:A(M=>oe(s.ID),["prevent"])},{icon:i(()=>[o(F(J),{icon:"dashicons:admin-site"})]),default:i(()=>[g(" "+I(s.domain),1)]),_:2},1032,["onClick"])]),_:1}),o(d,{prop:"title",label:"网站名称"}),o(d,{prop:"www_root",label:"站点目录"}),o(d,{prop:"status",label:"网站状态",width:"120"}),o(d,{fixed:"right",label:"操作",width:"100"},{default:i(s=>[o(r,{icon:F(Qe),circle:"",class:"mr-2 hidden-md-and-down",title:"设置网站",onClick:A(D=>H(s.row.ID),["prevent"])},null,8,["icon","onClick"]),o(re,null,{dropdown:i(()=>[o(se,null,{default:i(()=>[o(q,{onClick:A(D=>ie(s.$index),["prevent"])},{default:i(()=>[g("删除")]),_:2},1032,["onClick"]),o(q,{onClick:A(D=>Z(s.$index),["prevent"])},{default:i(()=>[g("停止")]),_:2},1032,["onClick"])]),_:2},1024)]),default:i(()=>[h("span",ct,[o(F(J),{icon:"ri:more-fill",width:"24",height:"24"})])]),_:2},1024)]),_:1})]),_:1},8,["data","height"])),[[be,n.tbLoading]])]),_:1}),o(ge,{modelValue:u.value,"onUpdate:modelValue":t[11]||(t[11]=s=>u.value=s),direction:"rtl",size:b.value,"show-close":!1},{header:i(({close:s,titleId:D,titleClass:M})=>[h("h4",{id:D,class:_(M)},I($.value),11,wt),o(r,{onClick:s},{default:i(()=>[o(me,{class:"el-icon--left"},{default:i(()=>[o(F(Xe))]),_:1}),g(" 关闭 ")]),_:2},1032,["onClick"])]),default:i(()=>[o(we,{model:a,"label-position":"left",ref_key:"websiteFormDrawerRef",ref:y,rules:Q},{default:i(()=>[o(E,{label:"域名","label-width":k,prop:"domain"},{default:i(()=>[o(S,{modelValue:a.domain,"onUpdate:modelValue":t[2]||(t[2]=s=>a.domain=s),autocomplete:"off",onChange:X},null,8,["modelValue"])]),_:1}),o(E,{label:"附加域名","label-width":k,prop:"domain_names"},{default:i(()=>[o(S,{modelValue:a.domain_names,"onUpdate:modelValue":t[3]||(t[3]=s=>a.domain_names=s),autocomplete:"off",type:"textarea"},null,8,["modelValue"])]),_:1}),o(E,{label:"网站名称","label-width":k,prop:"title"},{default:i(()=>[o(S,{modelValue:a.title,"onUpdate:modelValue":t[4]||(t[4]=s=>a.title=s),autocomplete:"off"},null,8,["modelValue"])]),_:1}),o(E,{label:"站点目录","label-width":k,prop:"www_root"},{default:i(()=>[o(S,{modelValue:a.www_root,"onUpdate:modelValue":t[5]||(t[5]=s=>a.www_root=s),autocomplete:"off"},{prepend:i(()=>[g(I(w.config.home_dir)+"/",1)]),_:1},8,["modelValue"])]),_:1}),o(E,{label:"默认主页","label-width":k,prop:"index_files"},{default:i(()=>[o(S,{modelValue:a.index_files,"onUpdate:modelValue":t[6]||(t[6]=s=>a.index_files=s),autocomplete:"off"},null,8,["modelValue"])]),_:1}),o(E,{label:"应用类型","label-width":k,prop:"application"},{default:i(()=>[o(ce,{modelValue:a.application,"onUpdate:modelValue":t[7]||(t[7]=s=>a.application=s),class:"m-2",placeholder:"中间件服务器",style:{width:"240px"}},{default:i(()=>[(f(!0),z(N,null,Se(w.config.applications,(s,D)=>(f(),P(fe,{key:s.id,label:s.title,value:s.id},{default:i(()=>[h("span",gt,I(s.title),1),h("span",bt,I(s.expose_port),1)]),_:2},1032,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),o(E,{label:"备注","label-width":k,prop:"description"},{default:i(()=>[o(S,{modelValue:a.description,"onUpdate:modelValue":t[8]||(t[8]=s=>a.description=s),autocomplete:"off",type:"textarea"},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),footer:i(()=>[h("div",_t,[o(r,{onClick:t[9]||(t[9]=s=>x(y.value))},{default:i(()=>[g("取消")]),_:1}),o(r,{type:"primary",onClick:t[10]||(t[10]=s=>te(y.value))},{default:i(()=>[g("确认")]),_:1})])]),_:1},8,["modelValue","size"])],64)}}},Zt=Ye(vt,[["__scopeId","data-v-8fa14623"]]);export{Zt as default};
