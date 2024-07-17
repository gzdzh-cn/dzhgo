import{i as t}from"./vue-demi-71ba0ef2.js";import{w as e,r as n,m as s,h as o,i as c,b as a,t as r,l as i,o as u,v as f}from"./@vue.reactivity-54707199.js";import{i as p,y as l,p as h,n as y,q as d}from"./@vue.runtime-core-6173334e.js";
/*!
  * pinia v2.0.28
  * (c) 2022 Eduardo San Martin Morote
  * @license MIT
  */let v;const b=t=>v=t,m=Symbol();function _(t){return t&&"object"==typeof t&&"[object Object]"===Object.prototype.toString.call(t)&&"function"!=typeof t.toJSON}var j,$;function O(){const o=e(!0),c=o.run((()=>n({})));let a=[],r=[];const i=s({install(t){b(i),i._a=t,t.provide(m,i),t.config.globalProperties.$pinia=i,r.forEach((t=>a.push(t))),r=[]},use(e){return this._a||t?a.push(e):r.push(e),this},_p:a,_a:null,_e:o,_s:new Map,state:c});return i}($=j||(j={})).direct="direct",$.patchObject="patch object",$.patchFunction="patch function";const g=()=>{};function P(t,e,n,s=g){t.push(e);const o=()=>{const n=t.indexOf(e);n>-1&&(t.splice(n,1),s())};return!n&&i()&&u(o),o}function S(t,...e){t.slice().forEach((t=>{t(...e)}))}function w(t,e){t instanceof Map&&e instanceof Map&&e.forEach(((e,n)=>t.set(n,e))),t instanceof Set&&e instanceof Set&&e.forEach(t.add,t);for(const n in e){if(!e.hasOwnProperty(n))continue;const s=e[n],o=t[n];_(o)&&_(s)&&t.hasOwnProperty(n)&&!c(s)&&!a(s)?t[n]=w(o,s):t[n]=s}return t}const E=Symbol();const{assign:I}=Object;function M(t,i,u={},f,p,l){let d;const v=I({actions:{}},u),m={deep:!0};let $,O,M,x=s([]),A=s([]);const F=f.state.value[t];let k;function q(e){let n;$=O=!1,"function"==typeof e?(e(f.state.value[t]),n={type:j.patchFunction,storeId:t,events:M}):(w(f.state.value[t],e),n={type:j.patchObject,payload:e,storeId:t,events:M});const s=k=Symbol();y().then((()=>{k===s&&($=!0)})),O=!0,S(x,n,f.state.value[t])}l||F||(f.state.value[t]={}),n({});const J=g;function N(e,n){return function(){b(f);const s=Array.from(arguments),o=[],c=[];function a(t){o.push(t)}function r(t){c.push(t)}let i;S(A,{args:s,name:e,store:B,after:a,onError:r});try{i=n.apply(this&&this.$id===t?this:B,s)}catch(u){throw S(c,u),u}return i instanceof Promise?i.then((t=>(S(o,t),t))).catch((t=>(S(c,t),Promise.reject(t)))):(S(o,i),i)}}const z={_p:f,$id:t,$onAction:P.bind(null,A),$patch:q,$reset:J,$subscribe(e,n={}){const s=P(x,e,n.detached,(()=>o())),o=d.run((()=>h((()=>f.state.value[t]),(s=>{("sync"===n.flush?O:$)&&e({storeId:t,type:j.direct,events:M},s)}),I({},m,n))));return s},$dispose:function(){d.stop(),x=[],A=[],f._s.delete(t)}},B=o(z);f._s.set(t,B);const C=f._e.run((()=>(d=e(),d.run((()=>i())))));for(const e in C){const n=C[e];if(c(n)&&(!c(G=n)||!G.effect)||a(n))l||(!F||_(D=n)&&D.hasOwnProperty(E)||(c(n)?n.value=F[e]:w(n,F[e])),f.state.value[t][e]=n);else if("function"==typeof n){const t=N(e,n);C[e]=t,v.actions[e]=n}}var D,G;return I(B,C),I(r(B),C),Object.defineProperty(B,"$state",{get:()=>f.state.value[t],set:t=>{q((e=>{I(e,t)}))}}),f._p.forEach((t=>{I(B,d.run((()=>t({store:B,app:f._a,pinia:f,options:v}))))})),F&&l&&u.hydrate&&u.hydrate(B.$state,F),$=!0,O=!0,B}function x(t,e,n){let o,c;const a="function"==typeof e;function r(t,n){const r=p();(t=t||r&&l(m,null))&&b(t),(t=v)._s.has(o)||(a?M(o,e,c,t):function(t,e,n,o){const{state:c,actions:a,getters:r}=e,i=n.state.value[t];let u;u=M(t,(function(){i||(n.state.value[t]=c?c():{});const e=f(n.state.value[t]);return I(e,a,Object.keys(r||{}).reduce(((e,o)=>(e[o]=s(d((()=>{b(n);const e=n._s.get(t);return r[o].call(e,e)}))),e)),{}))}),e,n,0,!0),u.$reset=function(){const t=c?c():{};this.$patch((e=>{I(e,t)}))}}(o,c,t));return t._s.get(o)}return"string"==typeof t?(o=t,c=a?n:e):(c=t,o=t.id),r.$id=o,r}export{O as c,x as d};
