import{V as f,aj as u,ak as g}from"./index-e26e72e9.js";function h(e){for(var t=-1,a=e==null?0:e.length,n={};++t<a;){var i=e[t];n[i[0]]=i[1]}return n}const c="__epPropKey",O=e=>e,m=e=>f(e)&&!!e[c],w=(e,t)=>{if(!f(e)||m(e))return e;const{values:a,required:n,default:i,type:p,validator:o}=e,d={type:p,required:!!n,validator:a||o?s=>{let r=!1,l=[];if(a&&(l=Array.from(a),u(e,"default")&&l.push(i),r||(r=l.includes(s))),o&&(r||(r=o(s))),!r&&l.length>0){const v=[...new Set(l)].map(P=>JSON.stringify(P)).join(", ");g(`Invalid prop: validation failed${t?` for prop "${t}"`:""}. Expected one of [${v}], got value ${JSON.stringify(s)}.`)}return r}:void 0,[c]:!0};return u(e,"default")&&(d.default=i),d},j=e=>h(Object.entries(e).map(([t,a])=>[t,w(a,t)]));export{w as a,j as b,O as d,h as f};
