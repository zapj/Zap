import{b as Y,d as D,a as $,u as ee,o as T,c as C,n as O,e as S,r as k,f as J,g as A,h as te,_ as ne,i as v,t as U,w as re,j as H,p as ft,k as He,l as y,m as Ie,q as M,s as dt,v as he,x as pt,y as ht,z as b,F as le,A as Ne,E as mt,B as yt,C as bt,D as wt}from"./index-13134c55.js";const Et=Y({header:{type:String,default:""},footer:{type:String,default:""},bodyStyle:{type:D([String,Object,Array]),default:""},bodyClass:String,shadow:{type:String,values:["always","hover","never"],default:"always"}}),St=$({name:"ElCard"}),Rt=$({...St,props:Et,setup(e){const t=ee("card");return(n,r)=>(T(),C("div",{class:O([S(t).b(),S(t).is(`${n.shadow}-shadow`)])},[n.$slots.header||n.header?(T(),C("div",{key:0,class:O(S(t).e("header"))},[k(n.$slots,"header",{},()=>[v(U(n.header),1)])],2)):J("v-if",!0),A("div",{class:O([S(t).e("body"),n.bodyClass]),style:te(n.bodyStyle)},[k(n.$slots,"default")],6),n.$slots.footer||n.footer?(T(),C("div",{key:1,class:O(S(t).e("footer"))},[k(n.$slots,"footer",{},()=>[v(U(n.footer),1)])],2)):J("v-if",!0)],2))}});var gt=ne(Rt,[["__file","card.vue"]]);const _t=re(gt),qe=Symbol("rowContextKey"),Ot=["start","center","end","space-around","space-between","space-evenly"],At=["top","middle","bottom"],Tt=Y({tag:{type:String,default:"div"},gutter:{type:Number,default:0},justify:{type:String,values:Ot,default:"start"},align:{type:String,values:At}}),Nt=$({name:"ElRow"}),Ct=$({...Nt,props:Tt,setup(e){const t=e,n=ee("row"),r=H(()=>t.gutter);ft(qe,{gutter:r});const i=H(()=>{const o={};return t.gutter&&(o.marginRight=o.marginLeft=`-${t.gutter/2}px`),o}),s=H(()=>[n.b(),n.is(`justify-${t.justify}`,t.justify!=="start"),n.is(`align-${t.align}`,!!t.align)]);return(o,c)=>(T(),He(Ie(o.tag),{class:O(S(s)),style:te(S(i))},{default:y(()=>[k(o.$slots,"default")]),_:3},8,["class","style"]))}});var xt=ne(Ct,[["__file","row.vue"]]);const Pt=re(xt),Ft=Y({tag:{type:String,default:"div"},span:{type:Number,default:24},offset:{type:Number,default:0},pull:{type:Number,default:0},push:{type:Number,default:0},xs:{type:D([Number,Object]),default:()=>M({})},sm:{type:D([Number,Object]),default:()=>M({})},md:{type:D([Number,Object]),default:()=>M({})},lg:{type:D([Number,Object]),default:()=>M({})},xl:{type:D([Number,Object]),default:()=>M({})}}),jt=$({name:"ElCol"}),Bt=$({...jt,props:Ft,setup(e){const t=e,{gutter:n}=dt(qe,{gutter:H(()=>0)}),r=ee("col"),i=H(()=>{const o={};return n.value&&(o.paddingLeft=o.paddingRight=`${n.value/2}px`),o}),s=H(()=>{const o=[];return["span","offset","pull","push"].forEach(f=>{const u=t[f];he(u)&&(f==="span"?o.push(r.b(`${t[f]}`)):u>0&&o.push(r.b(`${f}-${t[f]}`)))}),["xs","sm","md","lg","xl"].forEach(f=>{he(t[f])?o.push(r.b(`${f}-${t[f]}`)):pt(t[f])&&Object.entries(t[f]).forEach(([u,l])=>{o.push(u!=="span"?r.b(`${f}-${u}-${l}`):r.b(`${f}-${l}`))})}),n.value&&o.push(r.is("guttered")),[r.b(),o]});return(o,c)=>(T(),He(Ie(o.tag),{class:O(S(s)),style:te(S(i))},{default:y(()=>[k(o.$slots,"default")]),_:3},8,["class","style"]))}});var Lt=ne(Bt,[["__file","col.vue"]]);const Dt=re(Lt),Ut=Y({decimalSeparator:{type:String,default:"."},groupSeparator:{type:String,default:","},precision:{type:Number,default:0},formatter:Function,value:{type:D([Number,Object]),default:0},prefix:String,suffix:String,title:String,valueStyle:{type:D([String,Object,Array])}}),kt=$({name:"ElStatistic"}),$t=$({...kt,props:Ut,setup(e,{expose:t}){const n=e,r=ee("statistic"),i=H(()=>{const{value:s,formatter:o,precision:c,decimalSeparator:d,groupSeparator:f}=n;if(ht(o))return o(s);if(!he(s))return s;let[u,l=""]=String(s).split(".");return l=l.padEnd(c,"0").slice(0,c>0?c:0),u=u.replace(/\B(?=(\d{3})+(?!\d))/g,f),[u,l].join(l?d:"")});return t({displayValue:i}),(s,o)=>(T(),C("div",{class:O(S(r).b())},[s.$slots.title||s.title?(T(),C("div",{key:0,class:O(S(r).e("head"))},[k(s.$slots,"title",{},()=>[v(U(s.title),1)])],2)):J("v-if",!0),A("div",{class:O(S(r).e("content"))},[s.$slots.prefix||s.prefix?(T(),C("div",{key:0,class:O(S(r).e("prefix"))},[k(s.$slots,"prefix",{},()=>[A("span",null,U(s.prefix),1)])],2)):J("v-if",!0),A("span",{class:O(S(r).e("number")),style:te(s.valueStyle)},U(S(i)),7),s.$slots.suffix||s.suffix?(T(),C("div",{key:1,class:O(S(r).e("suffix"))},[k(s.$slots,"suffix",{},()=>[A("span",null,U(s.suffix),1)])],2)):J("v-if",!0)],2)],2))}});var vt=ne($t,[["__file","statistic.vue"]]);const Ht=re(vt);function Me(e,t){return function(){return e.apply(t,arguments)}}const{toString:It}=Object.prototype,{getPrototypeOf:Se}=Object,se=(e=>t=>{const n=It.call(t);return e[n]||(e[n]=n.slice(8,-1).toLowerCase())})(Object.create(null)),P=e=>(e=e.toLowerCase(),t=>se(t)===e),oe=e=>t=>typeof t===e,{isArray:q}=Array,V=oe("undefined");function qt(e){return e!==null&&!V(e)&&e.constructor!==null&&!V(e.constructor)&&N(e.constructor.isBuffer)&&e.constructor.isBuffer(e)}const ze=P("ArrayBuffer");function Mt(e){let t;return typeof ArrayBuffer<"u"&&ArrayBuffer.isView?t=ArrayBuffer.isView(e):t=e&&e.buffer&&ze(e.buffer),t}const zt=oe("string"),N=oe("function"),Je=oe("number"),ie=e=>e!==null&&typeof e=="object",Jt=e=>e===!0||e===!1,X=e=>{if(se(e)!=="object")return!1;const t=Se(e);return(t===null||t===Object.prototype||Object.getPrototypeOf(t)===null)&&!(Symbol.toStringTag in e)&&!(Symbol.iterator in e)},Vt=P("Date"),Kt=P("File"),Wt=P("Blob"),Xt=P("FileList"),Gt=e=>ie(e)&&N(e.pipe),Qt=e=>{let t;return e&&(typeof FormData=="function"&&e instanceof FormData||N(e.append)&&((t=se(e))==="formdata"||t==="object"&&N(e.toString)&&e.toString()==="[object FormData]"))},Zt=P("URLSearchParams"),Yt=e=>e.trim?e.trim():e.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g,"");function K(e,t,{allOwnKeys:n=!1}={}){if(e===null||typeof e>"u")return;let r,i;if(typeof e!="object"&&(e=[e]),q(e))for(r=0,i=e.length;r<i;r++)t.call(null,e[r],r,e);else{const s=n?Object.getOwnPropertyNames(e):Object.keys(e),o=s.length;let c;for(r=0;r<o;r++)c=s[r],t.call(null,e[c],c,e)}}function Ve(e,t){t=t.toLowerCase();const n=Object.keys(e);let r=n.length,i;for(;r-- >0;)if(i=n[r],t===i.toLowerCase())return i;return null}const Ke=(()=>typeof globalThis<"u"?globalThis:typeof self<"u"?self:typeof window<"u"?window:global)(),We=e=>!V(e)&&e!==Ke;function me(){const{caseless:e}=We(this)&&this||{},t={},n=(r,i)=>{const s=e&&Ve(t,i)||i;X(t[s])&&X(r)?t[s]=me(t[s],r):X(r)?t[s]=me({},r):q(r)?t[s]=r.slice():t[s]=r};for(let r=0,i=arguments.length;r<i;r++)arguments[r]&&K(arguments[r],n);return t}const en=(e,t,n,{allOwnKeys:r}={})=>(K(t,(i,s)=>{n&&N(i)?e[s]=Me(i,n):e[s]=i},{allOwnKeys:r}),e),tn=e=>(e.charCodeAt(0)===65279&&(e=e.slice(1)),e),nn=(e,t,n,r)=>{e.prototype=Object.create(t.prototype,r),e.prototype.constructor=e,Object.defineProperty(e,"super",{value:t.prototype}),n&&Object.assign(e.prototype,n)},rn=(e,t,n,r)=>{let i,s,o;const c={};if(t=t||{},e==null)return t;do{for(i=Object.getOwnPropertyNames(e),s=i.length;s-- >0;)o=i[s],(!r||r(o,e,t))&&!c[o]&&(t[o]=e[o],c[o]=!0);e=n!==!1&&Se(e)}while(e&&(!n||n(e,t))&&e!==Object.prototype);return t},sn=(e,t,n)=>{e=String(e),(n===void 0||n>e.length)&&(n=e.length),n-=t.length;const r=e.indexOf(t,n);return r!==-1&&r===n},on=e=>{if(!e)return null;if(q(e))return e;let t=e.length;if(!Je(t))return null;const n=new Array(t);for(;t-- >0;)n[t]=e[t];return n},an=(e=>t=>e&&t instanceof e)(typeof Uint8Array<"u"&&Se(Uint8Array)),cn=(e,t)=>{const r=(e&&e[Symbol.iterator]).call(e);let i;for(;(i=r.next())&&!i.done;){const s=i.value;t.call(e,s[0],s[1])}},ln=(e,t)=>{let n;const r=[];for(;(n=e.exec(t))!==null;)r.push(n);return r},un=P("HTMLFormElement"),fn=e=>e.toLowerCase().replace(/[-_\s]([a-z\d])(\w*)/g,function(n,r,i){return r.toUpperCase()+i}),Ce=(({hasOwnProperty:e})=>(t,n)=>e.call(t,n))(Object.prototype),dn=P("RegExp"),Xe=(e,t)=>{const n=Object.getOwnPropertyDescriptors(e),r={};K(n,(i,s)=>{let o;(o=t(i,s,e))!==!1&&(r[s]=o||i)}),Object.defineProperties(e,r)},pn=e=>{Xe(e,(t,n)=>{if(N(e)&&["arguments","caller","callee"].indexOf(n)!==-1)return!1;const r=e[n];if(N(r)){if(t.enumerable=!1,"writable"in t){t.writable=!1;return}t.set||(t.set=()=>{throw Error("Can not rewrite read-only method '"+n+"'")})}})},hn=(e,t)=>{const n={},r=i=>{i.forEach(s=>{n[s]=!0})};return q(e)?r(e):r(String(e).split(t)),n},mn=()=>{},yn=(e,t)=>(e=+e,Number.isFinite(e)?e:t),ue="abcdefghijklmnopqrstuvwxyz",xe="0123456789",Ge={DIGIT:xe,ALPHA:ue,ALPHA_DIGIT:ue+ue.toUpperCase()+xe},bn=(e=16,t=Ge.ALPHA_DIGIT)=>{let n="";const{length:r}=t;for(;e--;)n+=t[Math.random()*r|0];return n};function wn(e){return!!(e&&N(e.append)&&e[Symbol.toStringTag]==="FormData"&&e[Symbol.iterator])}const En=e=>{const t=new Array(10),n=(r,i)=>{if(ie(r)){if(t.indexOf(r)>=0)return;if(!("toJSON"in r)){t[i]=r;const s=q(r)?[]:{};return K(r,(o,c)=>{const d=n(o,i+1);!V(d)&&(s[c]=d)}),t[i]=void 0,s}}return r};return n(e,0)},Sn=P("AsyncFunction"),Rn=e=>e&&(ie(e)||N(e))&&N(e.then)&&N(e.catch),a={isArray:q,isArrayBuffer:ze,isBuffer:qt,isFormData:Qt,isArrayBufferView:Mt,isString:zt,isNumber:Je,isBoolean:Jt,isObject:ie,isPlainObject:X,isUndefined:V,isDate:Vt,isFile:Kt,isBlob:Wt,isRegExp:dn,isFunction:N,isStream:Gt,isURLSearchParams:Zt,isTypedArray:an,isFileList:Xt,forEach:K,merge:me,extend:en,trim:Yt,stripBOM:tn,inherits:nn,toFlatObject:rn,kindOf:se,kindOfTest:P,endsWith:sn,toArray:on,forEachEntry:cn,matchAll:ln,isHTMLForm:un,hasOwnProperty:Ce,hasOwnProp:Ce,reduceDescriptors:Xe,freezeMethods:pn,toObjectSet:hn,toCamelCase:fn,noop:mn,toFiniteNumber:yn,findKey:Ve,global:Ke,isContextDefined:We,ALPHABET:Ge,generateString:bn,isSpecCompliantForm:wn,toJSONObject:En,isAsyncFn:Sn,isThenable:Rn};function m(e,t,n,r,i){Error.call(this),Error.captureStackTrace?Error.captureStackTrace(this,this.constructor):this.stack=new Error().stack,this.message=e,this.name="AxiosError",t&&(this.code=t),n&&(this.config=n),r&&(this.request=r),i&&(this.response=i)}a.inherits(m,Error,{toJSON:function(){return{message:this.message,name:this.name,description:this.description,number:this.number,fileName:this.fileName,lineNumber:this.lineNumber,columnNumber:this.columnNumber,stack:this.stack,config:a.toJSONObject(this.config),code:this.code,status:this.response&&this.response.status?this.response.status:null}}});const Qe=m.prototype,Ze={};["ERR_BAD_OPTION_VALUE","ERR_BAD_OPTION","ECONNABORTED","ETIMEDOUT","ERR_NETWORK","ERR_FR_TOO_MANY_REDIRECTS","ERR_DEPRECATED","ERR_BAD_RESPONSE","ERR_BAD_REQUEST","ERR_CANCELED","ERR_NOT_SUPPORT","ERR_INVALID_URL"].forEach(e=>{Ze[e]={value:e}});Object.defineProperties(m,Ze);Object.defineProperty(Qe,"isAxiosError",{value:!0});m.from=(e,t,n,r,i,s)=>{const o=Object.create(Qe);return a.toFlatObject(e,o,function(d){return d!==Error.prototype},c=>c!=="isAxiosError"),m.call(o,e.message,t,n,r,i),o.cause=e,o.name=e.name,s&&Object.assign(o,s),o};const gn=null;function ye(e){return a.isPlainObject(e)||a.isArray(e)}function Ye(e){return a.endsWith(e,"[]")?e.slice(0,-2):e}function Pe(e,t,n){return e?e.concat(t).map(function(i,s){return i=Ye(i),!n&&s?"["+i+"]":i}).join(n?".":""):t}function _n(e){return a.isArray(e)&&!e.some(ye)}const On=a.toFlatObject(a,{},null,function(t){return/^is[A-Z]/.test(t)});function ae(e,t,n){if(!a.isObject(e))throw new TypeError("target must be an object");t=t||new FormData,n=a.toFlatObject(n,{metaTokens:!0,dots:!1,indexes:!1},!1,function(h,R){return!a.isUndefined(R[h])});const r=n.metaTokens,i=n.visitor||u,s=n.dots,o=n.indexes,d=(n.Blob||typeof Blob<"u"&&Blob)&&a.isSpecCompliantForm(t);if(!a.isFunction(i))throw new TypeError("visitor must be a function");function f(p){if(p===null)return"";if(a.isDate(p))return p.toISOString();if(!d&&a.isBlob(p))throw new m("Blob is not supported. Use a Buffer instead.");return a.isArrayBuffer(p)||a.isTypedArray(p)?d&&typeof Blob=="function"?new Blob([p]):Buffer.from(p):p}function u(p,h,R){let g=p;if(p&&!R&&typeof p=="object"){if(a.endsWith(h,"{}"))h=r?h:h.slice(0,-2),p=JSON.stringify(p);else if(a.isArray(p)&&_n(p)||(a.isFileList(p)||a.endsWith(h,"[]"))&&(g=a.toArray(p)))return h=Ye(h),g.forEach(function(B,ut){!(a.isUndefined(B)||B===null)&&t.append(o===!0?Pe([h],ut,s):o===null?h:h+"[]",f(B))}),!1}return ye(p)?!0:(t.append(Pe(R,h,s),f(p)),!1)}const l=[],E=Object.assign(On,{defaultVisitor:u,convertValue:f,isVisitable:ye});function _(p,h){if(!a.isUndefined(p)){if(l.indexOf(p)!==-1)throw Error("Circular reference detected in "+h.join("."));l.push(p),a.forEach(p,function(g,j){(!(a.isUndefined(g)||g===null)&&i.call(t,g,a.isString(j)?j.trim():j,h,E))===!0&&_(g,h?h.concat(j):[j])}),l.pop()}}if(!a.isObject(e))throw new TypeError("data must be an object");return _(e),t}function Fe(e){const t={"!":"%21","'":"%27","(":"%28",")":"%29","~":"%7E","%20":"+","%00":"\0"};return encodeURIComponent(e).replace(/[!'()~]|%20|%00/g,function(r){return t[r]})}function Re(e,t){this._pairs=[],e&&ae(e,this,t)}const et=Re.prototype;et.append=function(t,n){this._pairs.push([t,n])};et.toString=function(t){const n=t?function(r){return t.call(this,r,Fe)}:Fe;return this._pairs.map(function(i){return n(i[0])+"="+n(i[1])},"").join("&")};function An(e){return encodeURIComponent(e).replace(/%3A/gi,":").replace(/%24/g,"$").replace(/%2C/gi,",").replace(/%20/g,"+").replace(/%5B/gi,"[").replace(/%5D/gi,"]")}function tt(e,t,n){if(!t)return e;const r=n&&n.encode||An,i=n&&n.serialize;let s;if(i?s=i(t,n):s=a.isURLSearchParams(t)?t.toString():new Re(t,n).toString(r),s){const o=e.indexOf("#");o!==-1&&(e=e.slice(0,o)),e+=(e.indexOf("?")===-1?"?":"&")+s}return e}class Tn{constructor(){this.handlers=[]}use(t,n,r){return this.handlers.push({fulfilled:t,rejected:n,synchronous:r?r.synchronous:!1,runWhen:r?r.runWhen:null}),this.handlers.length-1}eject(t){this.handlers[t]&&(this.handlers[t]=null)}clear(){this.handlers&&(this.handlers=[])}forEach(t){a.forEach(this.handlers,function(r){r!==null&&t(r)})}}const je=Tn,nt={silentJSONParsing:!0,forcedJSONParsing:!0,clarifyTimeoutError:!1},Nn=typeof URLSearchParams<"u"?URLSearchParams:Re,Cn=typeof FormData<"u"?FormData:null,xn=typeof Blob<"u"?Blob:null,Pn={isBrowser:!0,classes:{URLSearchParams:Nn,FormData:Cn,Blob:xn},protocols:["http","https","file","blob","url","data"]},rt=typeof window<"u"&&typeof document<"u",Fn=(e=>rt&&["ReactNative","NativeScript","NS"].indexOf(e)<0)(typeof navigator<"u"&&navigator.product),jn=(()=>typeof WorkerGlobalScope<"u"&&self instanceof WorkerGlobalScope&&typeof self.importScripts=="function")(),Bn=Object.freeze(Object.defineProperty({__proto__:null,hasBrowserEnv:rt,hasStandardBrowserEnv:Fn,hasStandardBrowserWebWorkerEnv:jn},Symbol.toStringTag,{value:"Module"})),x={...Bn,...Pn};function Ln(e,t){return ae(e,new x.classes.URLSearchParams,Object.assign({visitor:function(n,r,i,s){return x.isNode&&a.isBuffer(n)?(this.append(r,n.toString("base64")),!1):s.defaultVisitor.apply(this,arguments)}},t))}function Dn(e){return a.matchAll(/\w+|\[(\w*)]/g,e).map(t=>t[0]==="[]"?"":t[1]||t[0])}function Un(e){const t={},n=Object.keys(e);let r;const i=n.length;let s;for(r=0;r<i;r++)s=n[r],t[s]=e[s];return t}function st(e){function t(n,r,i,s){let o=n[s++];if(o==="__proto__")return!0;const c=Number.isFinite(+o),d=s>=n.length;return o=!o&&a.isArray(i)?i.length:o,d?(a.hasOwnProp(i,o)?i[o]=[i[o],r]:i[o]=r,!c):((!i[o]||!a.isObject(i[o]))&&(i[o]=[]),t(n,r,i[o],s)&&a.isArray(i[o])&&(i[o]=Un(i[o])),!c)}if(a.isFormData(e)&&a.isFunction(e.entries)){const n={};return a.forEachEntry(e,(r,i)=>{t(Dn(r),i,n,0)}),n}return null}function kn(e,t,n){if(a.isString(e))try{return(t||JSON.parse)(e),a.trim(e)}catch(r){if(r.name!=="SyntaxError")throw r}return(n||JSON.stringify)(e)}const ge={transitional:nt,adapter:["xhr","http"],transformRequest:[function(t,n){const r=n.getContentType()||"",i=r.indexOf("application/json")>-1,s=a.isObject(t);if(s&&a.isHTMLForm(t)&&(t=new FormData(t)),a.isFormData(t))return i&&i?JSON.stringify(st(t)):t;if(a.isArrayBuffer(t)||a.isBuffer(t)||a.isStream(t)||a.isFile(t)||a.isBlob(t))return t;if(a.isArrayBufferView(t))return t.buffer;if(a.isURLSearchParams(t))return n.setContentType("application/x-www-form-urlencoded;charset=utf-8",!1),t.toString();let c;if(s){if(r.indexOf("application/x-www-form-urlencoded")>-1)return Ln(t,this.formSerializer).toString();if((c=a.isFileList(t))||r.indexOf("multipart/form-data")>-1){const d=this.env&&this.env.FormData;return ae(c?{"files[]":t}:t,d&&new d,this.formSerializer)}}return s||i?(n.setContentType("application/json",!1),kn(t)):t}],transformResponse:[function(t){const n=this.transitional||ge.transitional,r=n&&n.forcedJSONParsing,i=this.responseType==="json";if(t&&a.isString(t)&&(r&&!this.responseType||i)){const o=!(n&&n.silentJSONParsing)&&i;try{return JSON.parse(t)}catch(c){if(o)throw c.name==="SyntaxError"?m.from(c,m.ERR_BAD_RESPONSE,this,null,this.response):c}}return t}],timeout:0,xsrfCookieName:"XSRF-TOKEN",xsrfHeaderName:"X-XSRF-TOKEN",maxContentLength:-1,maxBodyLength:-1,env:{FormData:x.classes.FormData,Blob:x.classes.Blob},validateStatus:function(t){return t>=200&&t<300},headers:{common:{Accept:"application/json, text/plain, */*","Content-Type":void 0}}};a.forEach(["delete","get","head","post","put","patch"],e=>{ge.headers[e]={}});const _e=ge,$n=a.toObjectSet(["age","authorization","content-length","content-type","etag","expires","from","host","if-modified-since","if-unmodified-since","last-modified","location","max-forwards","proxy-authorization","referer","retry-after","user-agent"]),vn=e=>{const t={};let n,r,i;return e&&e.split(`
`).forEach(function(o){i=o.indexOf(":"),n=o.substring(0,i).trim().toLowerCase(),r=o.substring(i+1).trim(),!(!n||t[n]&&$n[n])&&(n==="set-cookie"?t[n]?t[n].push(r):t[n]=[r]:t[n]=t[n]?t[n]+", "+r:r)}),t},Be=Symbol("internals");function z(e){return e&&String(e).trim().toLowerCase()}function G(e){return e===!1||e==null?e:a.isArray(e)?e.map(G):String(e)}function Hn(e){const t=Object.create(null),n=/([^\s,;=]+)\s*(?:=\s*([^,;]+))?/g;let r;for(;r=n.exec(e);)t[r[1]]=r[2];return t}const In=e=>/^[-_a-zA-Z0-9^`|~,!#$%&'*+.]+$/.test(e.trim());function fe(e,t,n,r,i){if(a.isFunction(r))return r.call(this,t,n);if(i&&(t=n),!!a.isString(t)){if(a.isString(r))return t.indexOf(r)!==-1;if(a.isRegExp(r))return r.test(t)}}function qn(e){return e.trim().toLowerCase().replace(/([a-z\d])(\w*)/g,(t,n,r)=>n.toUpperCase()+r)}function Mn(e,t){const n=a.toCamelCase(" "+t);["get","set","has"].forEach(r=>{Object.defineProperty(e,r+n,{value:function(i,s,o){return this[r].call(this,t,i,s,o)},configurable:!0})})}class ce{constructor(t){t&&this.set(t)}set(t,n,r){const i=this;function s(c,d,f){const u=z(d);if(!u)throw new Error("header name must be a non-empty string");const l=a.findKey(i,u);(!l||i[l]===void 0||f===!0||f===void 0&&i[l]!==!1)&&(i[l||d]=G(c))}const o=(c,d)=>a.forEach(c,(f,u)=>s(f,u,d));return a.isPlainObject(t)||t instanceof this.constructor?o(t,n):a.isString(t)&&(t=t.trim())&&!In(t)?o(vn(t),n):t!=null&&s(n,t,r),this}get(t,n){if(t=z(t),t){const r=a.findKey(this,t);if(r){const i=this[r];if(!n)return i;if(n===!0)return Hn(i);if(a.isFunction(n))return n.call(this,i,r);if(a.isRegExp(n))return n.exec(i);throw new TypeError("parser must be boolean|regexp|function")}}}has(t,n){if(t=z(t),t){const r=a.findKey(this,t);return!!(r&&this[r]!==void 0&&(!n||fe(this,this[r],r,n)))}return!1}delete(t,n){const r=this;let i=!1;function s(o){if(o=z(o),o){const c=a.findKey(r,o);c&&(!n||fe(r,r[c],c,n))&&(delete r[c],i=!0)}}return a.isArray(t)?t.forEach(s):s(t),i}clear(t){const n=Object.keys(this);let r=n.length,i=!1;for(;r--;){const s=n[r];(!t||fe(this,this[s],s,t,!0))&&(delete this[s],i=!0)}return i}normalize(t){const n=this,r={};return a.forEach(this,(i,s)=>{const o=a.findKey(r,s);if(o){n[o]=G(i),delete n[s];return}const c=t?qn(s):String(s).trim();c!==s&&delete n[s],n[c]=G(i),r[c]=!0}),this}concat(...t){return this.constructor.concat(this,...t)}toJSON(t){const n=Object.create(null);return a.forEach(this,(r,i)=>{r!=null&&r!==!1&&(n[i]=t&&a.isArray(r)?r.join(", "):r)}),n}[Symbol.iterator](){return Object.entries(this.toJSON())[Symbol.iterator]()}toString(){return Object.entries(this.toJSON()).map(([t,n])=>t+": "+n).join(`
`)}get[Symbol.toStringTag](){return"AxiosHeaders"}static from(t){return t instanceof this?t:new this(t)}static concat(t,...n){const r=new this(t);return n.forEach(i=>r.set(i)),r}static accessor(t){const r=(this[Be]=this[Be]={accessors:{}}).accessors,i=this.prototype;function s(o){const c=z(o);r[c]||(Mn(i,o),r[c]=!0)}return a.isArray(t)?t.forEach(s):s(t),this}}ce.accessor(["Content-Type","Content-Length","Accept","Accept-Encoding","User-Agent","Authorization"]);a.reduceDescriptors(ce.prototype,({value:e},t)=>{let n=t[0].toUpperCase()+t.slice(1);return{get:()=>e,set(r){this[n]=r}}});a.freezeMethods(ce);const F=ce;function de(e,t){const n=this||_e,r=t||n,i=F.from(r.headers);let s=r.data;return a.forEach(e,function(c){s=c.call(n,s,i.normalize(),t?t.status:void 0)}),i.normalize(),s}function ot(e){return!!(e&&e.__CANCEL__)}function W(e,t,n){m.call(this,e??"canceled",m.ERR_CANCELED,t,n),this.name="CanceledError"}a.inherits(W,m,{__CANCEL__:!0});function zn(e,t,n){const r=n.config.validateStatus;!n.status||!r||r(n.status)?e(n):t(new m("Request failed with status code "+n.status,[m.ERR_BAD_REQUEST,m.ERR_BAD_RESPONSE][Math.floor(n.status/100)-4],n.config,n.request,n))}const Jn=x.hasStandardBrowserEnv?{write(e,t,n,r,i,s){const o=[e+"="+encodeURIComponent(t)];a.isNumber(n)&&o.push("expires="+new Date(n).toGMTString()),a.isString(r)&&o.push("path="+r),a.isString(i)&&o.push("domain="+i),s===!0&&o.push("secure"),document.cookie=o.join("; ")},read(e){const t=document.cookie.match(new RegExp("(^|;\\s*)("+e+")=([^;]*)"));return t?decodeURIComponent(t[3]):null},remove(e){this.write(e,"",Date.now()-864e5)}}:{write(){},read(){return null},remove(){}};function Vn(e){return/^([a-z][a-z\d+\-.]*:)?\/\//i.test(e)}function Kn(e,t){return t?e.replace(/\/?\/$/,"")+"/"+t.replace(/^\/+/,""):e}function it(e,t){return e&&!Vn(t)?Kn(e,t):t}const Wn=x.hasStandardBrowserEnv?function(){const t=/(msie|trident)/i.test(navigator.userAgent),n=document.createElement("a");let r;function i(s){let o=s;return t&&(n.setAttribute("href",o),o=n.href),n.setAttribute("href",o),{href:n.href,protocol:n.protocol?n.protocol.replace(/:$/,""):"",host:n.host,search:n.search?n.search.replace(/^\?/,""):"",hash:n.hash?n.hash.replace(/^#/,""):"",hostname:n.hostname,port:n.port,pathname:n.pathname.charAt(0)==="/"?n.pathname:"/"+n.pathname}}return r=i(window.location.href),function(o){const c=a.isString(o)?i(o):o;return c.protocol===r.protocol&&c.host===r.host}}():function(){return function(){return!0}}();function Xn(e){const t=/^([-+\w]{1,25})(:?\/\/|:)/.exec(e);return t&&t[1]||""}function Gn(e,t){e=e||10;const n=new Array(e),r=new Array(e);let i=0,s=0,o;return t=t!==void 0?t:1e3,function(d){const f=Date.now(),u=r[s];o||(o=f),n[i]=d,r[i]=f;let l=s,E=0;for(;l!==i;)E+=n[l++],l=l%e;if(i=(i+1)%e,i===s&&(s=(s+1)%e),f-o<t)return;const _=u&&f-u;return _?Math.round(E*1e3/_):void 0}}function Le(e,t){let n=0;const r=Gn(50,250);return i=>{const s=i.loaded,o=i.lengthComputable?i.total:void 0,c=s-n,d=r(c),f=s<=o;n=s;const u={loaded:s,total:o,progress:o?s/o:void 0,bytes:c,rate:d||void 0,estimated:d&&o&&f?(o-s)/d:void 0,event:i};u[t?"download":"upload"]=!0,e(u)}}const Qn=typeof XMLHttpRequest<"u",Zn=Qn&&function(e){return new Promise(function(n,r){let i=e.data;const s=F.from(e.headers).normalize();let{responseType:o,withXSRFToken:c}=e,d;function f(){e.cancelToken&&e.cancelToken.unsubscribe(d),e.signal&&e.signal.removeEventListener("abort",d)}let u;if(a.isFormData(i)){if(x.hasStandardBrowserEnv||x.hasStandardBrowserWebWorkerEnv)s.setContentType(!1);else if((u=s.getContentType())!==!1){const[h,...R]=u?u.split(";").map(g=>g.trim()).filter(Boolean):[];s.setContentType([h||"multipart/form-data",...R].join("; "))}}let l=new XMLHttpRequest;if(e.auth){const h=e.auth.username||"",R=e.auth.password?unescape(encodeURIComponent(e.auth.password)):"";s.set("Authorization","Basic "+btoa(h+":"+R))}const E=it(e.baseURL,e.url);l.open(e.method.toUpperCase(),tt(E,e.params,e.paramsSerializer),!0),l.timeout=e.timeout;function _(){if(!l)return;const h=F.from("getAllResponseHeaders"in l&&l.getAllResponseHeaders()),g={data:!o||o==="text"||o==="json"?l.responseText:l.response,status:l.status,statusText:l.statusText,headers:h,config:e,request:l};zn(function(B){n(B),f()},function(B){r(B),f()},g),l=null}if("onloadend"in l?l.onloadend=_:l.onreadystatechange=function(){!l||l.readyState!==4||l.status===0&&!(l.responseURL&&l.responseURL.indexOf("file:")===0)||setTimeout(_)},l.onabort=function(){l&&(r(new m("Request aborted",m.ECONNABORTED,e,l)),l=null)},l.onerror=function(){r(new m("Network Error",m.ERR_NETWORK,e,l)),l=null},l.ontimeout=function(){let R=e.timeout?"timeout of "+e.timeout+"ms exceeded":"timeout exceeded";const g=e.transitional||nt;e.timeoutErrorMessage&&(R=e.timeoutErrorMessage),r(new m(R,g.clarifyTimeoutError?m.ETIMEDOUT:m.ECONNABORTED,e,l)),l=null},x.hasStandardBrowserEnv&&(c&&a.isFunction(c)&&(c=c(e)),c||c!==!1&&Wn(E))){const h=e.xsrfHeaderName&&e.xsrfCookieName&&Jn.read(e.xsrfCookieName);h&&s.set(e.xsrfHeaderName,h)}i===void 0&&s.setContentType(null),"setRequestHeader"in l&&a.forEach(s.toJSON(),function(R,g){l.setRequestHeader(g,R)}),a.isUndefined(e.withCredentials)||(l.withCredentials=!!e.withCredentials),o&&o!=="json"&&(l.responseType=e.responseType),typeof e.onDownloadProgress=="function"&&l.addEventListener("progress",Le(e.onDownloadProgress,!0)),typeof e.onUploadProgress=="function"&&l.upload&&l.upload.addEventListener("progress",Le(e.onUploadProgress)),(e.cancelToken||e.signal)&&(d=h=>{l&&(r(!h||h.type?new W(null,e,l):h),l.abort(),l=null)},e.cancelToken&&e.cancelToken.subscribe(d),e.signal&&(e.signal.aborted?d():e.signal.addEventListener("abort",d)));const p=Xn(E);if(p&&x.protocols.indexOf(p)===-1){r(new m("Unsupported protocol "+p+":",m.ERR_BAD_REQUEST,e));return}l.send(i||null)})},be={http:gn,xhr:Zn};a.forEach(be,(e,t)=>{if(e){try{Object.defineProperty(e,"name",{value:t})}catch{}Object.defineProperty(e,"adapterName",{value:t})}});const De=e=>`- ${e}`,Yn=e=>a.isFunction(e)||e===null||e===!1,at={getAdapter:e=>{e=a.isArray(e)?e:[e];const{length:t}=e;let n,r;const i={};for(let s=0;s<t;s++){n=e[s];let o;if(r=n,!Yn(n)&&(r=be[(o=String(n)).toLowerCase()],r===void 0))throw new m(`Unknown adapter '${o}'`);if(r)break;i[o||"#"+s]=r}if(!r){const s=Object.entries(i).map(([c,d])=>`adapter ${c} `+(d===!1?"is not supported by the environment":"is not available in the build"));let o=t?s.length>1?`since :
`+s.map(De).join(`
`):" "+De(s[0]):"as no adapter specified";throw new m("There is no suitable adapter to dispatch the request "+o,"ERR_NOT_SUPPORT")}return r},adapters:be};function pe(e){if(e.cancelToken&&e.cancelToken.throwIfRequested(),e.signal&&e.signal.aborted)throw new W(null,e)}function Ue(e){return pe(e),e.headers=F.from(e.headers),e.data=de.call(e,e.transformRequest),["post","put","patch"].indexOf(e.method)!==-1&&e.headers.setContentType("application/x-www-form-urlencoded",!1),at.getAdapter(e.adapter||_e.adapter)(e).then(function(r){return pe(e),r.data=de.call(e,e.transformResponse,r),r.headers=F.from(r.headers),r},function(r){return ot(r)||(pe(e),r&&r.response&&(r.response.data=de.call(e,e.transformResponse,r.response),r.response.headers=F.from(r.response.headers))),Promise.reject(r)})}const ke=e=>e instanceof F?e.toJSON():e;function I(e,t){t=t||{};const n={};function r(f,u,l){return a.isPlainObject(f)&&a.isPlainObject(u)?a.merge.call({caseless:l},f,u):a.isPlainObject(u)?a.merge({},u):a.isArray(u)?u.slice():u}function i(f,u,l){if(a.isUndefined(u)){if(!a.isUndefined(f))return r(void 0,f,l)}else return r(f,u,l)}function s(f,u){if(!a.isUndefined(u))return r(void 0,u)}function o(f,u){if(a.isUndefined(u)){if(!a.isUndefined(f))return r(void 0,f)}else return r(void 0,u)}function c(f,u,l){if(l in t)return r(f,u);if(l in e)return r(void 0,f)}const d={url:s,method:s,data:s,baseURL:o,transformRequest:o,transformResponse:o,paramsSerializer:o,timeout:o,timeoutMessage:o,withCredentials:o,withXSRFToken:o,adapter:o,responseType:o,xsrfCookieName:o,xsrfHeaderName:o,onUploadProgress:o,onDownloadProgress:o,decompress:o,maxContentLength:o,maxBodyLength:o,beforeRedirect:o,transport:o,httpAgent:o,httpsAgent:o,cancelToken:o,socketPath:o,responseEncoding:o,validateStatus:c,headers:(f,u)=>i(ke(f),ke(u),!0)};return a.forEach(Object.keys(Object.assign({},e,t)),function(u){const l=d[u]||i,E=l(e[u],t[u],u);a.isUndefined(E)&&l!==c||(n[u]=E)}),n}const ct="1.6.5",Oe={};["object","boolean","number","function","string","symbol"].forEach((e,t)=>{Oe[e]=function(r){return typeof r===e||"a"+(t<1?"n ":" ")+e}});const $e={};Oe.transitional=function(t,n,r){function i(s,o){return"[Axios v"+ct+"] Transitional option '"+s+"'"+o+(r?". "+r:"")}return(s,o,c)=>{if(t===!1)throw new m(i(o," has been removed"+(n?" in "+n:"")),m.ERR_DEPRECATED);return n&&!$e[o]&&($e[o]=!0,console.warn(i(o," has been deprecated since v"+n+" and will be removed in the near future"))),t?t(s,o,c):!0}};function er(e,t,n){if(typeof e!="object")throw new m("options must be an object",m.ERR_BAD_OPTION_VALUE);const r=Object.keys(e);let i=r.length;for(;i-- >0;){const s=r[i],o=t[s];if(o){const c=e[s],d=c===void 0||o(c,s,e);if(d!==!0)throw new m("option "+s+" must be "+d,m.ERR_BAD_OPTION_VALUE);continue}if(n!==!0)throw new m("Unknown option "+s,m.ERR_BAD_OPTION)}}const we={assertOptions:er,validators:Oe},L=we.validators;class Z{constructor(t){this.defaults=t,this.interceptors={request:new je,response:new je}}request(t,n){typeof t=="string"?(n=n||{},n.url=t):n=t||{},n=I(this.defaults,n);const{transitional:r,paramsSerializer:i,headers:s}=n;r!==void 0&&we.assertOptions(r,{silentJSONParsing:L.transitional(L.boolean),forcedJSONParsing:L.transitional(L.boolean),clarifyTimeoutError:L.transitional(L.boolean)},!1),i!=null&&(a.isFunction(i)?n.paramsSerializer={serialize:i}:we.assertOptions(i,{encode:L.function,serialize:L.function},!0)),n.method=(n.method||this.defaults.method||"get").toLowerCase();let o=s&&a.merge(s.common,s[n.method]);s&&a.forEach(["delete","get","head","post","put","patch","common"],p=>{delete s[p]}),n.headers=F.concat(o,s);const c=[];let d=!0;this.interceptors.request.forEach(function(h){typeof h.runWhen=="function"&&h.runWhen(n)===!1||(d=d&&h.synchronous,c.unshift(h.fulfilled,h.rejected))});const f=[];this.interceptors.response.forEach(function(h){f.push(h.fulfilled,h.rejected)});let u,l=0,E;if(!d){const p=[Ue.bind(this),void 0];for(p.unshift.apply(p,c),p.push.apply(p,f),E=p.length,u=Promise.resolve(n);l<E;)u=u.then(p[l++],p[l++]);return u}E=c.length;let _=n;for(l=0;l<E;){const p=c[l++],h=c[l++];try{_=p(_)}catch(R){h.call(this,R);break}}try{u=Ue.call(this,_)}catch(p){return Promise.reject(p)}for(l=0,E=f.length;l<E;)u=u.then(f[l++],f[l++]);return u}getUri(t){t=I(this.defaults,t);const n=it(t.baseURL,t.url);return tt(n,t.params,t.paramsSerializer)}}a.forEach(["delete","get","head","options"],function(t){Z.prototype[t]=function(n,r){return this.request(I(r||{},{method:t,url:n,data:(r||{}).data}))}});a.forEach(["post","put","patch"],function(t){function n(r){return function(s,o,c){return this.request(I(c||{},{method:t,headers:r?{"Content-Type":"multipart/form-data"}:{},url:s,data:o}))}}Z.prototype[t]=n(),Z.prototype[t+"Form"]=n(!0)});const Q=Z;class Ae{constructor(t){if(typeof t!="function")throw new TypeError("executor must be a function.");let n;this.promise=new Promise(function(s){n=s});const r=this;this.promise.then(i=>{if(!r._listeners)return;let s=r._listeners.length;for(;s-- >0;)r._listeners[s](i);r._listeners=null}),this.promise.then=i=>{let s;const o=new Promise(c=>{r.subscribe(c),s=c}).then(i);return o.cancel=function(){r.unsubscribe(s)},o},t(function(s,o,c){r.reason||(r.reason=new W(s,o,c),n(r.reason))})}throwIfRequested(){if(this.reason)throw this.reason}subscribe(t){if(this.reason){t(this.reason);return}this._listeners?this._listeners.push(t):this._listeners=[t]}unsubscribe(t){if(!this._listeners)return;const n=this._listeners.indexOf(t);n!==-1&&this._listeners.splice(n,1)}static source(){let t;return{token:new Ae(function(i){t=i}),cancel:t}}}const tr=Ae;function nr(e){return function(n){return e.apply(null,n)}}function rr(e){return a.isObject(e)&&e.isAxiosError===!0}const Ee={Continue:100,SwitchingProtocols:101,Processing:102,EarlyHints:103,Ok:200,Created:201,Accepted:202,NonAuthoritativeInformation:203,NoContent:204,ResetContent:205,PartialContent:206,MultiStatus:207,AlreadyReported:208,ImUsed:226,MultipleChoices:300,MovedPermanently:301,Found:302,SeeOther:303,NotModified:304,UseProxy:305,Unused:306,TemporaryRedirect:307,PermanentRedirect:308,BadRequest:400,Unauthorized:401,PaymentRequired:402,Forbidden:403,NotFound:404,MethodNotAllowed:405,NotAcceptable:406,ProxyAuthenticationRequired:407,RequestTimeout:408,Conflict:409,Gone:410,LengthRequired:411,PreconditionFailed:412,PayloadTooLarge:413,UriTooLong:414,UnsupportedMediaType:415,RangeNotSatisfiable:416,ExpectationFailed:417,ImATeapot:418,MisdirectedRequest:421,UnprocessableEntity:422,Locked:423,FailedDependency:424,TooEarly:425,UpgradeRequired:426,PreconditionRequired:428,TooManyRequests:429,RequestHeaderFieldsTooLarge:431,UnavailableForLegalReasons:451,InternalServerError:500,NotImplemented:501,BadGateway:502,ServiceUnavailable:503,GatewayTimeout:504,HttpVersionNotSupported:505,VariantAlsoNegotiates:506,InsufficientStorage:507,LoopDetected:508,NotExtended:510,NetworkAuthenticationRequired:511};Object.entries(Ee).forEach(([e,t])=>{Ee[t]=e});const sr=Ee;function lt(e){const t=new Q(e),n=Me(Q.prototype.request,t);return a.extend(n,Q.prototype,t,{allOwnKeys:!0}),a.extend(n,t,null,{allOwnKeys:!0}),n.create=function(i){return lt(I(e,i))},n}const w=lt(_e);w.Axios=Q;w.CanceledError=W;w.CancelToken=tr;w.isCancel=ot;w.VERSION=ct;w.toFormData=ae;w.AxiosError=m;w.Cancel=w.CanceledError;w.all=function(t){return Promise.all(t)};w.spread=nr;w.isAxiosError=rr;w.mergeConfig=I;w.AxiosHeaders=F;w.formToJSON=e=>st(a.isHTMLForm(e)?new FormData(e):e);w.getAdapter=at.getAdapter;w.HttpStatusCode=sr;w.default=w;const or=w,ir={baseURL:"/api",useTokenAuthorization:!0},Te=or.create({baseURL:ir.baseURL,timeout:2e5,withCredentials:!0,headers:{"Content-Type":"application/json","X-Requested-With":"XMLHttpRequest"}});let ve;Te.interceptors.request.use(function(e){return e},function(e){return Promise.reject(e)});Te.interceptors.response.use(function(e){ve.close();const t=e.data;return t.reset,t},function(e){return ve.close(),e.config.url.indexOf("/api/file/list")==-1,Promise.reject(e)});const ar={style:{display:"inline-flex","align-items":"center"}},cr={class:"card-header"},lr=A("span",null,"Card name",-1),ur={class:"card-header"},fr=A("span",null,"Card name",-1),pr={__name:"DashboardView",setup(e){function t(){Te({url:"/v1/loadavg"})}return(n,r)=>{const i=Ht,s=Dt,o=mt,c=Pt,d=yt,f=_t;return T(),C(le,null,[b(c,{class:"mb-3",gutter:10},{default:y(()=>[b(s,{md:6,sm:12},{default:y(()=>[b(i,{title:"Daily active users",value:268500})]),_:1}),b(s,{md:6,sm:12},{default:y(()=>[b(i,{value:138},{title:y(()=>[A("div",ar,[v(" Ratio of men to women "),b(o,{style:{"margin-left":"4px"},size:12},{default:y(()=>[b(S(bt))]),_:1})])]),suffix:y(()=>[v("/100")]),_:1})]),_:1}),b(s,{md:6,sm:12},{default:y(()=>[b(i,{title:"Total Transactions",value:172e3})]),_:1}),b(s,{md:6,sm:12},{default:y(()=>[b(i,{title:"Feedback number",value:562},{suffix:y(()=>[b(o,{style:{"vertical-align":"-0.125em"}},{default:y(()=>[b(S(wt))]),_:1})]),_:1})]),_:1})]),_:1}),b(c,{gutter:10},{default:y(()=>[b(s,{md:12,sm:24},{default:y(()=>[b(f,{class:"box-card"},{header:y(()=>[A("div",cr,[lr,b(d,{class:"button",onClick:t},{default:y(()=>[v("Operation button")]),_:1})])]),default:y(()=>[(T(),C(le,null,Ne(4,u=>A("div",{key:u,class:"text item"},U("List item "+u),1)),64))]),_:1})]),_:1}),b(s,{md:12,sm:24},{default:y(()=>[b(f,{class:"box-card"},{header:y(()=>[A("div",ur,[fr,b(d,{class:"button",text:""},{default:y(()=>[v("Operation button")]),_:1})])]),default:y(()=>[(T(),C(le,null,Ne(4,u=>A("div",{key:u,class:"text item"},U("List item "+u),1)),64))]),_:1})]),_:1})]),_:1})],64)}}};export{pr as default};
