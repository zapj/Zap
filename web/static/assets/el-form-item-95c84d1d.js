import{u as Ze,f as xe,a as ge}from"./use-form-common-props-bf13ae6f.js";import{aC as $e,aD as He,aE as xt,aF as qt,aG as jt,b as qe,c as Qe,d as me,aH as Et,g as ye,aI as Xe,r as V,f as E,e as H,u as je,w as ne,P as ke,O as et,aJ as tt,o as ve,j as rt,q as ee,s as K,k as q,_ as nt,aK as it,L as ie,y as at,Q as st,aL as Tt,a1 as St,B as ae,V as _t,a2 as ot,aM as Pt,aN as $t,h as Ie,aO as de,m as ce,l as It,p as Mt,n as Me,a6 as Nt,a7 as Ne,D as Le,a5 as Re,aP as Lt,t as Rt,a8 as Bt}from"./index-b976bc21.js";import{c as he}from"./castArray-aa1129c0.js";import{d as Vt,t as Wt}from"./event-2c073216.js";import{u as Ct}from"./use-form-item-014ad19a.js";import{c as oe,k as ft,a as Ee,g as lt,s as Dt,b as Ut,d as Gt,e as zt,f as ut,h as Kt,i as Te,n as se,j as dt,l as Jt,m as Yt,o as Zt,p as Ht,S as Qt,q as Xt}from"./_initCloneObject-a305d225.js";function kt(r,e){for(var t=-1,n=r==null?0:r.length;++t<n&&e(r[t],t,r)!==!1;);return r}function er(r,e){return r&&oe(e,ft(e),r)}function tr(r,e){return r&&oe(e,Ee(e),r)}function rr(r,e){return oe(r,lt(r),e)}var nr=Object.getOwnPropertySymbols,ir=nr?function(r){for(var e=[];r;)Ut(e,lt(r)),r=Gt(r);return e}:Dt;const ct=ir;function ar(r,e){return oe(r,ct(r),e)}function sr(r){return zt(r,Ee,ct)}var or=Object.prototype,fr=or.hasOwnProperty;function lr(r){var e=r.length,t=new r.constructor(e);return e&&typeof r[0]=="string"&&fr.call(r,"index")&&(t.index=r.index,t.input=r.input),t}function ur(r,e){var t=e?ut(r.buffer):r.buffer;return new r.constructor(t,r.byteOffset,r.byteLength)}var dr=/\w*$/;function cr(r){var e=new r.constructor(r.source,dr.exec(r));return e.lastIndex=r.lastIndex,e}var Be=$e?$e.prototype:void 0,Ve=Be?Be.valueOf:void 0;function pr(r){return Ve?Object(Ve.call(r)):{}}var gr="[object Boolean]",mr="[object Date]",yr="[object Map]",vr="[object Number]",hr="[object RegExp]",br="[object Set]",wr="[object String]",Fr="[object Symbol]",Ar="[object ArrayBuffer]",Or="[object DataView]",xr="[object Float32Array]",qr="[object Float64Array]",jr="[object Int8Array]",Er="[object Int16Array]",Tr="[object Int32Array]",Sr="[object Uint8Array]",_r="[object Uint8ClampedArray]",Pr="[object Uint16Array]",$r="[object Uint32Array]";function Ir(r,e,t){var n=r.constructor;switch(e){case Ar:return ut(r);case gr:case mr:return new n(+r);case Or:return ur(r,t);case xr:case qr:case jr:case Er:case Tr:case Sr:case _r:case Pr:case $r:return Kt(r,t);case yr:return new n;case vr:case wr:return new n(r);case hr:return cr(r);case br:return new n;case Fr:return pr(r)}}var Mr="[object Map]";function Nr(r){return He(r)&&Te(r)==Mr}var We=se&&se.isMap,Lr=We?dt(We):Nr;const Rr=Lr;var Br="[object Set]";function Vr(r){return He(r)&&Te(r)==Br}var Ce=se&&se.isSet,Wr=Ce?dt(Ce):Vr;const Cr=Wr;var Dr=1,Ur=2,Gr=4,pt="[object Arguments]",zr="[object Array]",Kr="[object Boolean]",Jr="[object Date]",Yr="[object Error]",gt="[object Function]",Zr="[object GeneratorFunction]",Hr="[object Map]",Qr="[object Number]",mt="[object Object]",Xr="[object RegExp]",kr="[object Set]",en="[object String]",tn="[object Symbol]",rn="[object WeakMap]",nn="[object ArrayBuffer]",an="[object DataView]",sn="[object Float32Array]",on="[object Float64Array]",fn="[object Int8Array]",ln="[object Int16Array]",un="[object Int32Array]",dn="[object Uint8Array]",cn="[object Uint8ClampedArray]",pn="[object Uint16Array]",gn="[object Uint32Array]",O={};O[pt]=O[zr]=O[nn]=O[an]=O[Kr]=O[Jr]=O[sn]=O[on]=O[fn]=O[ln]=O[un]=O[Hr]=O[Qr]=O[mt]=O[Xr]=O[kr]=O[en]=O[tn]=O[dn]=O[cn]=O[pn]=O[gn]=!0;O[Yr]=O[gt]=O[rn]=!1;function te(r,e,t,n,i,s){var a,o=e&Dr,u=e&Ur,b=e&Gr;if(t&&(a=i?t(r,n,i,s):t(r)),a!==void 0)return a;if(!xt(r))return r;var g=qt(r);if(g){if(a=lr(r),!o)return Jt(r,a)}else{var m=Te(r),h=m==gt||m==Zr;if(Yt(r))return Zt(r,o);if(m==mt||m==pt||h&&!i){if(a=u||h?{}:Ht(r),!o)return u?ar(r,tr(a,r)):rr(r,er(a,r))}else{if(!O[m])return i?r:{};a=Ir(r,m,o)}}s||(s=new Qt);var x=s.get(r);if(x)return x;s.set(r,a),Cr(r)?r.forEach(function(y){a.add(te(y,e,t,y,r,s))}):Rr(r)&&r.forEach(function(y,f){a.set(f,te(y,e,t,f,r,s))});var j=b?u?sr:Xt:u?Ee:ft,d=g?void 0:j(r);return kt(d||r,function(y,f){d&&(f=y,y=r[f]),jt(a,f,te(y,e,t,f,r,s))}),a}var mn=4;function De(r){return te(r,mn)}const yn=qe({size:{type:String,values:Qe},disabled:Boolean}),vn=qe({...yn,model:Object,rules:{type:me(Object)},labelPosition:{type:String,values:["left","right","top"],default:"right"},requireAsteriskPosition:{type:String,values:["left","right"],default:"left"},labelWidth:{type:[String,Number],default:""},labelSuffix:{type:String,default:""},inline:Boolean,inlineMessage:Boolean,statusIcon:Boolean,showMessage:{type:Boolean,default:!0},validateOnRuleChange:{type:Boolean,default:!0},hideRequiredAsterisk:Boolean,scrollToError:Boolean,scrollIntoViewOptions:{type:[Object,Boolean]}}),hn={validate:(r,e,t)=>(Et(r)||ye(r))&&Xe(e)&&ye(t)};function bn(){const r=V([]),e=E(()=>{if(!r.value.length)return"0";const s=Math.max(...r.value);return s?`${s}px`:""});function t(s){const a=r.value.indexOf(s);return a===-1&&e.value,a}function n(s,a){if(s&&a){const o=t(a);r.value.splice(o,1,s)}else s&&r.value.push(s)}function i(s){const a=t(s);a>-1&&r.value.splice(a,1)}return{autoLabelWidth:e,registerLabelWidth:n,deregisterLabelWidth:i}}const X=(r,e)=>{const t=he(e);return t.length>0?r.filter(n=>n.prop&&t.includes(n.prop)):r},wn="ElForm",Fn=H({name:wn}),An=H({...Fn,props:vn,emits:hn,setup(r,{expose:e,emit:t}){const n=r,i=[],s=Ze(),a=je("form"),o=E(()=>{const{labelPosition:l,inline:c}=n;return[a.b(),a.m(s.value||"default"),{[a.m(`label-${l}`)]:l,[a.m("inline")]:c}]}),u=l=>i.find(c=>c.prop===l),b=l=>{i.push(l)},g=l=>{l.prop&&i.splice(i.indexOf(l),1)},m=(l=[])=>{n.model&&X(i,l).forEach(c=>c.resetField())},h=(l=[])=>{X(i,l).forEach(c=>c.clearValidate())},x=E(()=>!!n.model),j=l=>{if(i.length===0)return[];const c=X(i,l);return c.length?c:[]},d=async l=>f(void 0,l),y=async(l=[])=>{if(!x.value)return!1;const c=j(l);if(c.length===0)return!0;let A={};for(const w of c)try{await w.validate("")}catch(_){A={...A,..._}}return Object.keys(A).length===0?!0:Promise.reject(A)},f=async(l=[],c)=>{const A=!it(c);try{const w=await y(l);return w===!0&&(c==null||c(w)),w}catch(w){if(w instanceof Error)throw w;const _=w;return n.scrollToError&&$(Object.keys(_)[0]),c==null||c(!1,_),A&&Promise.reject(_)}},$=l=>{var c;const A=X(i,l)[0];A&&((c=A.$el)==null||c.scrollIntoView(n.scrollIntoViewOptions))};return ne(()=>n.rules,()=>{n.validateOnRuleChange&&d().catch(l=>Vt())},{deep:!0}),ke(xe,et({...tt(n),emit:t,resetFields:m,clearValidate:h,validateField:f,getField:u,addField:b,removeField:g,...bn()})),e({validate:d,validateField:f,resetFields:m,clearValidate:h,scrollToField:$}),(l,c)=>(ve(),rt("form",{class:K(q(o))},[ee(l.$slots,"default")],2))}});var On=nt(An,[["__file","form.vue"]]);function W(){return W=Object.assign?Object.assign.bind():function(r){for(var e=1;e<arguments.length;e++){var t=arguments[e];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(r[n]=t[n])}return r},W.apply(this,arguments)}function xn(r,e){r.prototype=Object.create(e.prototype),r.prototype.constructor=r,Z(r,e)}function be(r){return be=Object.setPrototypeOf?Object.getPrototypeOf.bind():function(t){return t.__proto__||Object.getPrototypeOf(t)},be(r)}function Z(r,e){return Z=Object.setPrototypeOf?Object.setPrototypeOf.bind():function(n,i){return n.__proto__=i,n},Z(r,e)}function qn(){if(typeof Reflect>"u"||!Reflect.construct||Reflect.construct.sham)return!1;if(typeof Proxy=="function")return!0;try{return Boolean.prototype.valueOf.call(Reflect.construct(Boolean,[],function(){})),!0}catch{return!1}}function re(r,e,t){return qn()?re=Reflect.construct.bind():re=function(i,s,a){var o=[null];o.push.apply(o,s);var u=Function.bind.apply(i,o),b=new u;return a&&Z(b,a.prototype),b},re.apply(null,arguments)}function jn(r){return Function.toString.call(r).indexOf("[native code]")!==-1}function we(r){var e=typeof Map=="function"?new Map:void 0;return we=function(n){if(n===null||!jn(n))return n;if(typeof n!="function")throw new TypeError("Super expression must either be null or a function");if(typeof e<"u"){if(e.has(n))return e.get(n);e.set(n,i)}function i(){return re(n,arguments,be(this).constructor)}return i.prototype=Object.create(n.prototype,{constructor:{value:i,enumerable:!1,writable:!0,configurable:!0}}),Z(i,n)},we(r)}var En=/%[sdj%]/g,Tn=function(){};typeof process<"u"&&process.env;function Fe(r){if(!r||!r.length)return null;var e={};return r.forEach(function(t){var n=t.field;e[n]=e[n]||[],e[n].push(t)}),e}function M(r){for(var e=arguments.length,t=new Array(e>1?e-1:0),n=1;n<e;n++)t[n-1]=arguments[n];var i=0,s=t.length;if(typeof r=="function")return r.apply(null,t);if(typeof r=="string"){var a=r.replace(En,function(o){if(o==="%%")return"%";if(i>=s)return o;switch(o){case"%s":return String(t[i++]);case"%d":return Number(t[i++]);case"%j":try{return JSON.stringify(t[i++])}catch{return"[Circular]"}break;default:return o}});return a}return r}function Sn(r){return r==="string"||r==="url"||r==="hex"||r==="email"||r==="date"||r==="pattern"}function T(r,e){return!!(r==null||e==="array"&&Array.isArray(r)&&!r.length||Sn(e)&&typeof r=="string"&&!r)}function _n(r,e,t){var n=[],i=0,s=r.length;function a(o){n.push.apply(n,o||[]),i++,i===s&&t(n)}r.forEach(function(o){e(o,a)})}function Ue(r,e,t){var n=0,i=r.length;function s(a){if(a&&a.length){t(a);return}var o=n;n=n+1,o<i?e(r[o],s):t([])}s([])}function Pn(r){var e=[];return Object.keys(r).forEach(function(t){e.push.apply(e,r[t]||[])}),e}var Ge=function(r){xn(e,r);function e(t,n){var i;return i=r.call(this,"Async Validation Error")||this,i.errors=t,i.fields=n,i}return e}(we(Error));function $n(r,e,t,n,i){if(e.first){var s=new Promise(function(h,x){var j=function(f){return n(f),f.length?x(new Ge(f,Fe(f))):h(i)},d=Pn(r);Ue(d,t,j)});return s.catch(function(h){return h}),s}var a=e.firstFields===!0?Object.keys(r):e.firstFields||[],o=Object.keys(r),u=o.length,b=0,g=[],m=new Promise(function(h,x){var j=function(y){if(g.push.apply(g,y),b++,b===u)return n(g),g.length?x(new Ge(g,Fe(g))):h(i)};o.length||(n(g),h(i)),o.forEach(function(d){var y=r[d];a.indexOf(d)!==-1?Ue(y,t,j):_n(y,t,j)})});return m.catch(function(h){return h}),m}function In(r){return!!(r&&r.message!==void 0)}function Mn(r,e){for(var t=r,n=0;n<e.length;n++){if(t==null)return t;t=t[e[n]]}return t}function ze(r,e){return function(t){var n;return r.fullFields?n=Mn(e,r.fullFields):n=e[t.field||r.fullField],In(t)?(t.field=t.field||r.fullField,t.fieldValue=n,t):{message:typeof t=="function"?t():t,fieldValue:n,field:t.field||r.fullField}}}function Ke(r,e){if(e){for(var t in e)if(e.hasOwnProperty(t)){var n=e[t];typeof n=="object"&&typeof r[t]=="object"?r[t]=W({},r[t],n):r[t]=n}}return r}var yt=function(e,t,n,i,s,a){e.required&&(!n.hasOwnProperty(e.field)||T(t,a||e.type))&&i.push(M(s.messages.required,e.fullField))},Nn=function(e,t,n,i,s){(/^\s+$/.test(t)||t==="")&&i.push(M(s.messages.whitespace,e.fullField))},k,Ln=function(){if(k)return k;var r="[a-fA-F\\d:]",e=function(c){return c&&c.includeBoundaries?"(?:(?<=\\s|^)(?="+r+")|(?<="+r+")(?=\\s|$))":""},t="(?:25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]\\d|\\d)(?:\\.(?:25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]\\d|\\d)){3}",n="[a-fA-F\\d]{1,4}",i=(`
(?:
(?:`+n+":){7}(?:"+n+`|:)|                                    // 1:2:3:4:5:6:7::  1:2:3:4:5:6:7:8
(?:`+n+":){6}(?:"+t+"|:"+n+`|:)|                             // 1:2:3:4:5:6::    1:2:3:4:5:6::8   1:2:3:4:5:6::8  1:2:3:4:5:6::1.2.3.4
(?:`+n+":){5}(?::"+t+"|(?::"+n+`){1,2}|:)|                   // 1:2:3:4:5::      1:2:3:4:5::7:8   1:2:3:4:5::8    1:2:3:4:5::7:1.2.3.4
(?:`+n+":){4}(?:(?::"+n+"){0,1}:"+t+"|(?::"+n+`){1,3}|:)| // 1:2:3:4::        1:2:3:4::6:7:8   1:2:3:4::8      1:2:3:4::6:7:1.2.3.4
(?:`+n+":){3}(?:(?::"+n+"){0,2}:"+t+"|(?::"+n+`){1,4}|:)| // 1:2:3::          1:2:3::5:6:7:8   1:2:3::8        1:2:3::5:6:7:1.2.3.4
(?:`+n+":){2}(?:(?::"+n+"){0,3}:"+t+"|(?::"+n+`){1,5}|:)| // 1:2::            1:2::4:5:6:7:8   1:2::8          1:2::4:5:6:7:1.2.3.4
(?:`+n+":){1}(?:(?::"+n+"){0,4}:"+t+"|(?::"+n+`){1,6}|:)| // 1::              1::3:4:5:6:7:8   1::8            1::3:4:5:6:7:1.2.3.4
(?::(?:(?::`+n+"){0,5}:"+t+"|(?::"+n+`){1,7}|:))             // ::2:3:4:5:6:7:8  ::2:3:4:5:6:7:8  ::8             ::1.2.3.4
)(?:%[0-9a-zA-Z]{1,})?                                             // %eth0            %1
`).replace(/\s*\/\/.*$/gm,"").replace(/\n/g,"").trim(),s=new RegExp("(?:^"+t+"$)|(?:^"+i+"$)"),a=new RegExp("^"+t+"$"),o=new RegExp("^"+i+"$"),u=function(c){return c&&c.exact?s:new RegExp("(?:"+e(c)+t+e(c)+")|(?:"+e(c)+i+e(c)+")","g")};u.v4=function(l){return l&&l.exact?a:new RegExp(""+e(l)+t+e(l),"g")},u.v6=function(l){return l&&l.exact?o:new RegExp(""+e(l)+i+e(l),"g")};var b="(?:(?:[a-z]+:)?//)",g="(?:\\S+(?::\\S*)?@)?",m=u.v4().source,h=u.v6().source,x="(?:(?:[a-z\\u00a1-\\uffff0-9][-_]*)*[a-z\\u00a1-\\uffff0-9]+)",j="(?:\\.(?:[a-z\\u00a1-\\uffff0-9]-*)*[a-z\\u00a1-\\uffff0-9]+)*",d="(?:\\.(?:[a-z\\u00a1-\\uffff]{2,}))",y="(?::\\d{2,5})?",f='(?:[/?#][^\\s"]*)?',$="(?:"+b+"|www\\.)"+g+"(?:localhost|"+m+"|"+h+"|"+x+j+d+")"+y+f;return k=new RegExp("(?:^"+$+"$)","i"),k},Je={email:/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+\.)+[a-zA-Z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]{2,}))$/,hex:/^#?([a-f0-9]{6}|[a-f0-9]{3})$/i},J={integer:function(e){return J.number(e)&&parseInt(e,10)===e},float:function(e){return J.number(e)&&!J.integer(e)},array:function(e){return Array.isArray(e)},regexp:function(e){if(e instanceof RegExp)return!0;try{return!!new RegExp(e)}catch{return!1}},date:function(e){return typeof e.getTime=="function"&&typeof e.getMonth=="function"&&typeof e.getYear=="function"&&!isNaN(e.getTime())},number:function(e){return isNaN(e)?!1:typeof e=="number"},object:function(e){return typeof e=="object"&&!J.array(e)},method:function(e){return typeof e=="function"},email:function(e){return typeof e=="string"&&e.length<=320&&!!e.match(Je.email)},url:function(e){return typeof e=="string"&&e.length<=2048&&!!e.match(Ln())},hex:function(e){return typeof e=="string"&&!!e.match(Je.hex)}},Rn=function(e,t,n,i,s){if(e.required&&t===void 0){yt(e,t,n,i,s);return}var a=["integer","float","array","regexp","object","method","email","number","date","url","hex"],o=e.type;a.indexOf(o)>-1?J[o](t)||i.push(M(s.messages.types[o],e.fullField,e.type)):o&&typeof t!==e.type&&i.push(M(s.messages.types[o],e.fullField,e.type))},Bn=function(e,t,n,i,s){var a=typeof e.len=="number",o=typeof e.min=="number",u=typeof e.max=="number",b=/[\uD800-\uDBFF][\uDC00-\uDFFF]/g,g=t,m=null,h=typeof t=="number",x=typeof t=="string",j=Array.isArray(t);if(h?m="number":x?m="string":j&&(m="array"),!m)return!1;j&&(g=t.length),x&&(g=t.replace(b,"_").length),a?g!==e.len&&i.push(M(s.messages[m].len,e.fullField,e.len)):o&&!u&&g<e.min?i.push(M(s.messages[m].min,e.fullField,e.min)):u&&!o&&g>e.max?i.push(M(s.messages[m].max,e.fullField,e.max)):o&&u&&(g<e.min||g>e.max)&&i.push(M(s.messages[m].range,e.fullField,e.min,e.max))},z="enum",Vn=function(e,t,n,i,s){e[z]=Array.isArray(e[z])?e[z]:[],e[z].indexOf(t)===-1&&i.push(M(s.messages[z],e.fullField,e[z].join(", ")))},Wn=function(e,t,n,i,s){if(e.pattern){if(e.pattern instanceof RegExp)e.pattern.lastIndex=0,e.pattern.test(t)||i.push(M(s.messages.pattern.mismatch,e.fullField,t,e.pattern));else if(typeof e.pattern=="string"){var a=new RegExp(e.pattern);a.test(t)||i.push(M(s.messages.pattern.mismatch,e.fullField,t,e.pattern))}}},v={required:yt,whitespace:Nn,type:Rn,range:Bn,enum:Vn,pattern:Wn},Cn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t,"string")&&!e.required)return n();v.required(e,t,i,a,s,"string"),T(t,"string")||(v.type(e,t,i,a,s),v.range(e,t,i,a,s),v.pattern(e,t,i,a,s),e.whitespace===!0&&v.whitespace(e,t,i,a,s))}n(a)},Dn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&v.type(e,t,i,a,s)}n(a)},Un=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(t===""&&(t=void 0),T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&(v.type(e,t,i,a,s),v.range(e,t,i,a,s))}n(a)},Gn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&v.type(e,t,i,a,s)}n(a)},zn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),T(t)||v.type(e,t,i,a,s)}n(a)},Kn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&(v.type(e,t,i,a,s),v.range(e,t,i,a,s))}n(a)},Jn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&(v.type(e,t,i,a,s),v.range(e,t,i,a,s))}n(a)},Yn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(t==null&&!e.required)return n();v.required(e,t,i,a,s,"array"),t!=null&&(v.type(e,t,i,a,s),v.range(e,t,i,a,s))}n(a)},Zn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&v.type(e,t,i,a,s)}n(a)},Hn="enum",Qn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s),t!==void 0&&v[Hn](e,t,i,a,s)}n(a)},Xn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t,"string")&&!e.required)return n();v.required(e,t,i,a,s),T(t,"string")||v.pattern(e,t,i,a,s)}n(a)},kn=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t,"date")&&!e.required)return n();if(v.required(e,t,i,a,s),!T(t,"date")){var u;t instanceof Date?u=t:u=new Date(t),v.type(e,u,i,a,s),u&&v.range(e,u.getTime(),i,a,s)}}n(a)},ei=function(e,t,n,i,s){var a=[],o=Array.isArray(t)?"array":typeof t;v.required(e,t,i,a,s,o),n(a)},pe=function(e,t,n,i,s){var a=e.type,o=[],u=e.required||!e.required&&i.hasOwnProperty(e.field);if(u){if(T(t,a)&&!e.required)return n();v.required(e,t,i,o,s,a),T(t,a)||v.type(e,t,i,o,s)}n(o)},ti=function(e,t,n,i,s){var a=[],o=e.required||!e.required&&i.hasOwnProperty(e.field);if(o){if(T(t)&&!e.required)return n();v.required(e,t,i,a,s)}n(a)},Y={string:Cn,method:Dn,number:Un,boolean:Gn,regexp:zn,integer:Kn,float:Jn,array:Yn,object:Zn,enum:Qn,pattern:Xn,date:kn,url:pe,hex:pe,email:pe,required:ei,any:ti};function Ae(){return{default:"Validation error on field %s",required:"%s is required",enum:"%s must be one of %s",whitespace:"%s cannot be empty",date:{format:"%s date %s is invalid for format %s",parse:"%s date could not be parsed, %s is invalid ",invalid:"%s date %s is invalid"},types:{string:"%s is not a %s",method:"%s is not a %s (function)",array:"%s is not an %s",object:"%s is not an %s",number:"%s is not a %s",date:"%s is not a %s",boolean:"%s is not a %s",integer:"%s is not an %s",float:"%s is not a %s",regexp:"%s is not a valid %s",email:"%s is not a valid %s",url:"%s is not a valid %s",hex:"%s is not a valid %s"},string:{len:"%s must be exactly %s characters",min:"%s must be at least %s characters",max:"%s cannot be longer than %s characters",range:"%s must be between %s and %s characters"},number:{len:"%s must equal %s",min:"%s cannot be less than %s",max:"%s cannot be greater than %s",range:"%s must be between %s and %s"},array:{len:"%s must be exactly %s in length",min:"%s cannot be less than %s in length",max:"%s cannot be greater than %s in length",range:"%s must be between %s and %s in length"},pattern:{mismatch:"%s value %s does not match pattern %s"},clone:function(){var e=JSON.parse(JSON.stringify(this));return e.clone=this.clone,e}}}var Oe=Ae(),Q=function(){function r(t){this.rules=null,this._messages=Oe,this.define(t)}var e=r.prototype;return e.define=function(n){var i=this;if(!n)throw new Error("Cannot configure a schema with no rules");if(typeof n!="object"||Array.isArray(n))throw new Error("Rules must be an object");this.rules={},Object.keys(n).forEach(function(s){var a=n[s];i.rules[s]=Array.isArray(a)?a:[a]})},e.messages=function(n){return n&&(this._messages=Ke(Ae(),n)),this._messages},e.validate=function(n,i,s){var a=this;i===void 0&&(i={}),s===void 0&&(s=function(){});var o=n,u=i,b=s;if(typeof u=="function"&&(b=u,u={}),!this.rules||Object.keys(this.rules).length===0)return b&&b(null,o),Promise.resolve(o);function g(d){var y=[],f={};function $(c){if(Array.isArray(c)){var A;y=(A=y).concat.apply(A,c)}else y.push(c)}for(var l=0;l<d.length;l++)$(d[l]);y.length?(f=Fe(y),b(y,f)):b(null,o)}if(u.messages){var m=this.messages();m===Oe&&(m=Ae()),Ke(m,u.messages),u.messages=m}else u.messages=this.messages();var h={},x=u.keys||Object.keys(this.rules);x.forEach(function(d){var y=a.rules[d],f=o[d];y.forEach(function($){var l=$;typeof l.transform=="function"&&(o===n&&(o=W({},o)),f=o[d]=l.transform(f)),typeof l=="function"?l={validator:l}:l=W({},l),l.validator=a.getValidationMethod(l),l.validator&&(l.field=d,l.fullField=l.fullField||d,l.type=a.getType(l),h[d]=h[d]||[],h[d].push({rule:l,value:f,source:o,field:d}))})});var j={};return $n(h,u,function(d,y){var f=d.rule,$=(f.type==="object"||f.type==="array")&&(typeof f.fields=="object"||typeof f.defaultField=="object");$=$&&(f.required||!f.required&&d.value),f.field=d.field;function l(w,_){return W({},_,{fullField:f.fullField+"."+w,fullFields:f.fullFields?[].concat(f.fullFields,[w]):[w]})}function c(w){w===void 0&&(w=[]);var _=Array.isArray(w)?w:[w];!u.suppressWarning&&_.length&&r.warning("async-validator:",_),_.length&&f.message!==void 0&&(_=[].concat(f.message));var I=_.map(ze(f,o));if(u.first&&I.length)return j[f.field]=1,y(I);if(!$)y(I);else{if(f.required&&!d.value)return f.message!==void 0?I=[].concat(f.message).map(ze(f,o)):u.error&&(I=[u.error(f,M(u.messages.required,f.field))]),y(I);var B={};f.defaultField&&Object.keys(d.value).map(function(L){B[L]=f.defaultField}),B=W({},B,d.rule.fields);var C={};Object.keys(B).forEach(function(L){var N=B[L],fe=Array.isArray(N)?N:[N];C[L]=fe.map(l.bind(null,L))});var D=new r(C);D.messages(u.messages),d.rule.options&&(d.rule.options.messages=u.messages,d.rule.options.error=u.error),D.validate(d.value,d.rule.options||u,function(L){var N=[];I&&I.length&&N.push.apply(N,I),L&&L.length&&N.push.apply(N,L),y(N.length?N:null)})}}var A;if(f.asyncValidator)A=f.asyncValidator(f,d.value,c,d.source,u);else if(f.validator){try{A=f.validator(f,d.value,c,d.source,u)}catch(w){console.error==null||console.error(w),u.suppressValidatorError||setTimeout(function(){throw w},0),c(w.message)}A===!0?c():A===!1?c(typeof f.message=="function"?f.message(f.fullField||f.field):f.message||(f.fullField||f.field)+" fails"):A instanceof Array?c(A):A instanceof Error&&c(A.message)}A&&A.then&&A.then(function(){return c()},function(w){return c(w)})},function(d){g(d)},o)},e.getType=function(n){if(n.type===void 0&&n.pattern instanceof RegExp&&(n.type="pattern"),typeof n.validator!="function"&&n.type&&!Y.hasOwnProperty(n.type))throw new Error(M("Unknown rule type %s",n.type));return n.type||"string"},e.getValidationMethod=function(n){if(typeof n.validator=="function")return n.validator;var i=Object.keys(n),s=i.indexOf("message");return s!==-1&&i.splice(s,1),i.length===1&&i[0]==="required"?Y.required:Y[this.getType(n)]||void 0},r}();Q.register=function(e,t){if(typeof t!="function")throw new Error("Cannot register a validator by type, validator is not a function");Y[e]=t};Q.warning=Tn;Q.messages=Oe;Q.validators=Y;const ri=["","error","validating","success"],ni=qe({label:String,labelWidth:{type:[String,Number],default:""},prop:{type:me([String,Array])},required:{type:Boolean,default:void 0},rules:{type:me([Object,Array])},error:String,validateStatus:{type:String,values:ri},for:String,inlineMessage:{type:[String,Boolean],default:""},showMessage:{type:Boolean,default:!0},size:{type:String,values:Qe}}),Ye="ElLabelWrap";var ii=H({name:Ye,props:{isAutoWidth:Boolean,updateAll:Boolean},setup(r,{slots:e}){const t=ie(xe,void 0),n=ie(ge);n||Wt(Ye,"usage: <el-form-item><label-wrap /></el-form-item>");const i=je("form"),s=V(),a=V(0),o=()=>{var g;if((g=s.value)!=null&&g.firstElementChild){const m=window.getComputedStyle(s.value.firstElementChild).width;return Math.ceil(Number.parseFloat(m))}else return 0},u=(g="update")=>{ot(()=>{e.default&&r.isAutoWidth&&(g==="update"?a.value=o():g==="remove"&&(t==null||t.deregisterLabelWidth(a.value)))})},b=()=>u("update");return at(()=>{b()}),st(()=>{u("remove")}),Tt(()=>b()),ne(a,(g,m)=>{r.updateAll&&(t==null||t.registerLabelWidth(g,m))}),St(E(()=>{var g,m;return(m=(g=s.value)==null?void 0:g.firstElementChild)!=null?m:null}),b),()=>{var g,m;if(!e)return null;const{isAutoWidth:h}=r;if(h){const x=t==null?void 0:t.autoLabelWidth,j=n==null?void 0:n.hasLabel,d={};if(j&&x&&x!=="auto"){const y=Math.max(0,Number.parseInt(x,10)-a.value),f=t.labelPosition==="left"?"marginRight":"marginLeft";y&&(d[f]=`${y}px`)}return ae("div",{ref:s,class:[i.be("item","label-wrap")],style:d},[(g=e.default)==null?void 0:g.call(e)])}else return ae(_t,{ref:s},[(m=e.default)==null?void 0:m.call(e)])}}});const ai=["role","aria-labelledby"],si=H({name:"ElFormItem"}),oi=H({...si,props:ni,setup(r,{expose:e}){const t=r,n=Pt(),i=ie(xe,void 0),s=ie(ge,void 0),a=Ze(void 0,{formItem:!1}),o=je("form-item"),u=Ct().value,b=V([]),g=V(""),m=$t(g,100),h=V(""),x=V();let j,d=!1;const y=E(()=>{if((i==null?void 0:i.labelPosition)==="top")return{};const p=Ie(t.labelWidth||(i==null?void 0:i.labelWidth)||"");return p?{width:p}:{}}),f=E(()=>{if((i==null?void 0:i.labelPosition)==="top"||i!=null&&i.inline)return{};if(!t.label&&!t.labelWidth&&B)return{};const p=Ie(t.labelWidth||(i==null?void 0:i.labelWidth)||"");return!t.label&&!n.label?{marginLeft:p}:{}}),$=E(()=>[o.b(),o.m(a.value),o.is("error",g.value==="error"),o.is("validating",g.value==="validating"),o.is("success",g.value==="success"),o.is("required",fe.value||t.required),o.is("no-asterisk",i==null?void 0:i.hideRequiredAsterisk),(i==null?void 0:i.requireAsteriskPosition)==="right"?"asterisk-right":"asterisk-left",{[o.m("feedback")]:i==null?void 0:i.statusIcon}]),l=E(()=>Xe(t.inlineMessage)?t.inlineMessage:(i==null?void 0:i.inlineMessage)||!1),c=E(()=>[o.e("error"),{[o.em("error","inline")]:l.value}]),A=E(()=>t.prop?ye(t.prop)?t.prop:t.prop.join("."):""),w=E(()=>!!(t.label||n.label)),_=E(()=>t.for||(b.value.length===1?b.value[0]:void 0)),I=E(()=>!_.value&&w.value),B=!!s,C=E(()=>{const p=i==null?void 0:i.model;if(!(!p||!t.prop))return de(p,t.prop).value}),D=E(()=>{const{required:p}=t,F=[];t.rules&&F.push(...he(t.rules));const P=i==null?void 0:i.rules;if(P&&t.prop){const S=de(P,t.prop).value;S&&F.push(...he(S))}if(p!==void 0){const S=F.map((R,G)=>[R,G]).filter(([R])=>Object.keys(R).includes("required"));if(S.length>0)for(const[R,G]of S)R.required!==p&&(F[G]={...R,required:p});else F.push({required:p})}return F}),L=E(()=>D.value.length>0),N=p=>D.value.filter(P=>!P.trigger||!p?!0:Array.isArray(P.trigger)?P.trigger.includes(p):P.trigger===p).map(({trigger:P,...S})=>S),fe=E(()=>D.value.some(p=>p.required)),ht=E(()=>{var p;return m.value==="error"&&t.showMessage&&((p=i==null?void 0:i.showMessage)!=null?p:!0)}),Se=E(()=>`${t.label||""}${(i==null?void 0:i.labelSuffix)||""}`),U=p=>{g.value=p},bt=p=>{var F,P;const{errors:S,fields:R}=p;(!S||!R)&&console.error(p),U("error"),h.value=S?(P=(F=S==null?void 0:S[0])==null?void 0:F.message)!=null?P:`${t.prop} is required`:"",i==null||i.emit("validate",t.prop,!1,h.value)},wt=()=>{U("success"),i==null||i.emit("validate",t.prop,!0,"")},Ft=async p=>{const F=A.value;return new Q({[F]:p}).validate({[F]:C.value},{firstFields:!0}).then(()=>(wt(),!0)).catch(S=>(bt(S),Promise.reject(S)))},_e=async(p,F)=>{if(d||!t.prop)return!1;const P=it(F);if(!L.value)return F==null||F(!1),!1;const S=N(p);return S.length===0?(F==null||F(!0),!0):(U("validating"),Ft(S).then(()=>(F==null||F(!0),!0)).catch(R=>{const{fields:G}=R;return F==null||F(!1,G),P?!1:Promise.reject(G)}))},le=()=>{U(""),h.value="",d=!1},Pe=async()=>{const p=i==null?void 0:i.model;if(!p||!t.prop)return;const F=de(p,t.prop);d=!0,F.value=De(j),await ot(),le(),d=!1},At=p=>{b.value.includes(p)||b.value.push(p)},Ot=p=>{b.value=b.value.filter(F=>F!==p)};ne(()=>t.error,p=>{h.value=p||"",U(p?"error":"")},{immediate:!0}),ne(()=>t.validateStatus,p=>U(p||""));const ue=et({...tt(t),$el:x,size:a,validateState:g,labelId:u,inputIds:b,isGroup:I,hasLabel:w,fieldValue:C,addInputId:At,removeInputId:Ot,resetField:Pe,clearValidate:le,validate:_e});return ke(ge,ue),at(()=>{t.prop&&(i==null||i.addField(ue),j=De(C.value))}),st(()=>{i==null||i.removeField(ue)}),e({size:a,validateMessage:h,validateState:g,validate:_e,clearValidate:le,resetField:Pe}),(p,F)=>{var P;return ve(),rt("div",{ref_key:"formItemRef",ref:x,class:K(q($)),role:q(I)?"group":void 0,"aria-labelledby":q(I)?q(u):void 0},[ae(q(ii),{"is-auto-width":q(y).width==="auto","update-all":((P=q(i))==null?void 0:P.labelWidth)==="auto"},{default:ce(()=>[q(w)?(ve(),It(Mt(q(_)?"label":"div"),{key:0,id:q(u),for:q(_),class:K(q(o).e("label")),style:Me(q(y))},{default:ce(()=>[ee(p.$slots,"label",{label:q(Se)},()=>[Nt(Ne(q(Se)),1)])]),_:3},8,["id","for","class","style"])):Le("v-if",!0)]),_:3},8,["is-auto-width","update-all"]),Re("div",{class:K(q(o).e("content")),style:Me(q(f))},[ee(p.$slots,"default"),ae(Lt,{name:`${q(o).namespace.value}-zoom-in-top`},{default:ce(()=>[q(ht)?ee(p.$slots,"error",{key:0,error:h.value},()=>[Re("div",{class:K(q(c))},Ne(h.value),3)]):Le("v-if",!0)]),_:3},8,["name"])],6)],10,ai)}}});var vt=nt(oi,[["__file","form-item.vue"]]);const gi=Rt(On,{FormItem:vt}),mi=Bt(vt);export{mi as E,gi as a};
