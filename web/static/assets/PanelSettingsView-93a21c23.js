import{r as l,q as m,a as d,f as F,j as r,w as s,F as b,K as n,L as v}from"./index-57ba148c.js";import{a as w,E as x}from"./el-form-item-1909e3b0.js";import{E}from"./el-button-7b2c59ce.js";/* empty css                   */import{b as y}from"./index-407d4441.js";import"./use-form-common-props-7f97df7e.js";import"./runtime-14a2c257.js";import"./castArray-7af8d8fb.js";import"./index-74d71456.js";import"./index-def65ea4.js";import"./error-78e43d3e.js";import"./use-form-item-c8f82fc4.js";import"./_initCloneObject-0341bd43.js";const g=n("div",{class:"about"},[n("h5",null,"面板设置")],-1),P={__name:"PanelSettingsView",setup(h){const i=l("default"),a=l(),u=m({}),c=m({}),p=async t=>{t&&await t.validate((e,o)=>{e?y.success("保存成功！"):console.log("error submit!",o)})};return(t,e)=>{const o=E,_=x,f=w;return d(),F(b,null,[g,r(f,{ref_key:"ruleFormRef",ref:a,style:{"max-width":"600px"},model:u,rules:c,"label-width":"auto",class:"demo-ruleForm",size:i.value,"status-icon":""},{default:s(()=>[r(_,null,{default:s(()=>[r(o,{type:"primary",onClick:e[0]||(e[0]=k=>p(a.value))},{default:s(()=>[v(" 保存 ")]),_:1})]),_:1})]),_:1},8,["model","rules","size"])],64)}}};export{P as default};