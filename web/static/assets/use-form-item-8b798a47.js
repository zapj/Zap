import{f,b as p}from"./use-form-common-props-75f0c225.js";import{au as m,h as v,m as h,K as j,F as i,r as c,o as y,e as F,W as K,av as w}from"./index-163cc183.js";const l={prefix:Math.floor(Math.random()*1e4),current:0},N=Symbol("elIdInjection"),R=()=>j()?i(N,l):l,_=n=>{const u=R(),s=m();return v(()=>h(n)||`${s.value}-id-${u.prefix}-${u.current++}`)},E=()=>{const n=i(f,void 0),u=i(p,void 0);return{form:n,formItem:u}},U=(n,{formItemContext:u,disableIdGeneration:s,disableIdManagement:o})=>{s||(s=c(!1)),o||(o=c(!1));const e=c();let r;const I=v(()=>{var a;return!!(!n.label&&u&&u.inputIds&&((a=u.inputIds)==null?void 0:a.length)<=1)});return y(()=>{r=F([K(n,"id"),s],([a,d])=>{const t=a??(d?void 0:_().value);t!==e.value&&(u!=null&&u.removeInputId&&(e.value&&u.removeInputId(e.value),!(o!=null&&o.value)&&!d&&t&&u.addInputId(t)),e.value=t)},{immediate:!0})}),w(()=>{r&&r(),u!=null&&u.removeInputId&&e.value&&u.removeInputId(e.value)}),{isLabeledByFormItem:I,inputId:e}};export{E as a,U as b,R as c,_ as u};
