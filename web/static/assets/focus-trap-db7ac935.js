import{m as g,K as R,r as y,_ as te,d as ne,F as se,w as L,e as T,j as oe,C as B,a as re}from"./index-7b2c4294.js";import{E as X}from"./aria-a6bf21a8.js";import{j}from"./use-form-common-props-a3884077.js";function ce(e){return e==null}class ae extends Error{constructor(n){super(n),this.name="ElementPlusError"}}function he(e,n){throw new ae(`[${e}] ${n}`)}function Le(e,n){}const Ce="update:modelValue";let E=[];const M=e=>{const n=e;n.key===X.esc&&E.forEach(s=>s(n))},ue=e=>{g(()=>{E.length===0&&document.addEventListener("keydown",M),j&&E.push(e)}),R(()=>{E=E.filter(n=>n!==e),E.length===0&&j&&document.removeEventListener("keydown",M)})},C="focus-trap.focus-after-trapped",O="focus-trap.focus-after-released",ie="focus-trap.focusout-prevented",W={cancelable:!0,bubbles:!1},de={cancelable:!0,bubbles:!1},q="focusAfterTrapped",J="focusAfterReleased",fe=Symbol("elFocusTrap"),N=y(),P=y(0),k=y(0);let F=0;const Z=e=>{const n=[],s=document.createTreeWalker(e,NodeFilter.SHOW_ELEMENT,{acceptNode:o=>{const c=o.tagName==="INPUT"&&o.type==="hidden";return o.disabled||o.hidden||c?NodeFilter.FILTER_SKIP:o.tabIndex>=0||o===document.activeElement?NodeFilter.FILTER_ACCEPT:NodeFilter.FILTER_SKIP}});for(;s.nextNode();)n.push(s.currentNode);return n},Y=(e,n)=>{for(const s of e)if(!le(s,n))return s},le=(e,n)=>{if(getComputedStyle(e).visibility==="hidden")return!0;for(;e;){if(n&&e===n)return!1;if(getComputedStyle(e).display==="none")return!0;e=e.parentElement}return!1},ve=e=>{const n=Z(e),s=Y(n,e),o=Y(n.reverse(),e);return[s,o]},Ee=e=>e instanceof HTMLInputElement&&"select"in e,l=(e,n)=>{if(e&&e.focus){const s=document.activeElement;e.focus({preventScroll:!0}),k.value=window.performance.now(),e!==s&&Ee(e)&&n&&e.select()}};function z(e,n){const s=[...e],o=e.indexOf(n);return o!==-1&&s.splice(o,1),s}const pe=()=>{let e=[];return{push:o=>{const c=e[0];c&&o!==c&&c.pause(),e=z(e,o),e.unshift(o)},remove:o=>{var c,d;e=z(e,o),(d=(c=e[0])==null?void 0:c.resume)==null||d.call(c)}}},me=(e,n=!1)=>{const s=document.activeElement;for(const o of e)if(l(o,n),document.activeElement!==s)return},G=pe(),Te=()=>P.value>k.value,_=()=>{N.value="pointer",P.value=window.performance.now()},Q=()=>{N.value="keyboard",P.value=window.performance.now()},Fe=()=>(g(()=>{F===0&&(document.addEventListener("mousedown",_),document.addEventListener("touchstart",_),document.addEventListener("keydown",Q)),F++}),R(()=>{F--,F<=0&&(document.removeEventListener("mousedown",_),document.removeEventListener("touchstart",_),document.removeEventListener("keydown",Q))}),{focusReason:N,lastUserFocusTimestamp:P,lastAutomatedFocusTimestamp:k}),b=e=>new CustomEvent(ie,{...de,detail:e}),_e=ne({name:"ElFocusTrap",inheritAttrs:!1,props:{loop:Boolean,trapped:Boolean,focusTrapEl:Object,focusStartEl:{type:[Object,String],default:"first"}},emits:[q,J,"focusin","focusout","focusout-prevented","release-requested"],setup(e,{emit:n}){const s=y();let o,c;const{focusReason:d}=Fe();ue(t=>{e.trapped&&!v.paused&&n("release-requested",t)});const v={paused:!1,pause(){this.paused=!0},resume(){this.paused=!1}},p=t=>{if(!e.loop&&!e.trapped||v.paused)return;const{key:r,altKey:a,ctrlKey:u,metaKey:i,currentTarget:x,shiftKey:$}=t,{loop:H}=e,ee=r===X.tab&&!a&&!u&&!i,m=document.activeElement;if(ee&&m){const w=x,[S,h]=ve(w);if(S&&h){if(!$&&m===h){const f=b({focusReason:d.value});n("focusout-prevented",f),f.defaultPrevented||(t.preventDefault(),H&&l(S,!0))}else if($&&[S,w].includes(m)){const f=b({focusReason:d.value});n("focusout-prevented",f),f.defaultPrevented||(t.preventDefault(),H&&l(h,!0))}}else if(m===w){const f=b({focusReason:d.value});n("focusout-prevented",f),f.defaultPrevented||t.preventDefault()}}};se(fe,{focusTrapRef:s,onKeydown:p}),L(()=>e.focusTrapEl,t=>{t&&(s.value=t)},{immediate:!0}),L([s],([t],[r])=>{t&&(t.addEventListener("keydown",p),t.addEventListener("focusin",U),t.addEventListener("focusout",K)),r&&(r.removeEventListener("keydown",p),r.removeEventListener("focusin",U),r.removeEventListener("focusout",K))});const A=t=>{n(q,t)},I=t=>n(J,t),U=t=>{const r=T(s);if(!r)return;const a=t.target,u=t.relatedTarget,i=a&&r.contains(a);e.trapped||u&&r.contains(u)||(o=u),i&&n("focusin",t),!v.paused&&e.trapped&&(i?c=a:l(c,!0))},K=t=>{const r=T(s);if(!(v.paused||!r))if(e.trapped){const a=t.relatedTarget;!ce(a)&&!r.contains(a)&&setTimeout(()=>{if(!v.paused&&e.trapped){const u=b({focusReason:d.value});n("focusout-prevented",u),u.defaultPrevented||l(c,!0)}},0)}else{const a=t.target;a&&r.contains(a)||n("focusout",t)}};async function D(){await B();const t=T(s);if(t){G.push(v);const r=t.contains(document.activeElement)?o:document.activeElement;if(o=r,!t.contains(r)){const u=new Event(C,W);t.addEventListener(C,A),t.dispatchEvent(u),u.defaultPrevented||B(()=>{let i=e.focusStartEl;re(i)||(l(i),document.activeElement!==i&&(i="first")),i==="first"&&me(Z(t),!0),(document.activeElement===r||i==="container")&&l(t)})}}}function V(){const t=T(s);if(t){t.removeEventListener(C,A);const r=new CustomEvent(O,{...W,detail:{focusReason:d.value}});t.addEventListener(O,I),t.dispatchEvent(r),!r.defaultPrevented&&(d.value=="keyboard"||!Te()||t.contains(document.activeElement))&&l(o??document.body),t.removeEventListener(O,I),G.remove(v)}}return g(()=>{e.trapped&&D(),L(()=>e.trapped,t=>{t?D():V()})}),R(()=>{e.trapped&&V()}),{onKeydown:p}}});function be(e,n,s,o,c,d){return oe(e.$slots,"default",{handleKeydown:e.onKeydown})}var Oe=te(_e,[["render",be],["__file","focus-trap.vue"]]);export{Oe as E,fe as F,Ce as U,Le as d,ce as i,he as t};
