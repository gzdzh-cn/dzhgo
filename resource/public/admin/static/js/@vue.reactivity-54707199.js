import{i as t,h as e,a as s,e as n,t as i,d as r,b as c,N as o,c as a,f as u,g as h,j as l,m as f,k as _}from"./@vue.shared-0cc4c744.js";let d;class p{constructor(t=!1){this.detached=t,this.active=!0,this.effects=[],this.cleanups=[],this.parent=d,!t&&d&&(this.index=(d.scopes||(d.scopes=[])).push(this)-1)}run(t){if(this.active){const e=d;try{return d=this,t()}finally{d=e}}}on(){d=this}off(){d=this.parent}stop(t){if(this.active){let e,s;for(e=0,s=this.effects.length;e<s;e++)this.effects[e].stop();for(e=0,s=this.cleanups.length;e<s;e++)this.cleanups[e]();if(this.scopes)for(e=0,s=this.scopes.length;e<s;e++)this.scopes[e].stop(!0);if(!this.detached&&this.parent&&!t){const t=this.parent.scopes.pop();t&&t!==this&&(this.parent.scopes[this.index]=t,t.index=this.index)}this.parent=void 0,this.active=!1}}}function v(t){return new p(t)}function g(t,e=d){e&&e.active&&e.effects.push(t)}function w(){return d}function y(t){d&&d.cleanups.push(t)}const R=t=>{const e=new Set(t);return e.w=0,e.n=0,e},b=t=>(t.w&x)>0,m=t=>(t.n&x)>0,S=new WeakMap;let k=0,x=1;let E;const j=Symbol(""),z=Symbol("");class M{constructor(t,e=null,s){this.fn=t,this.scheduler=e,this.active=!0,this.deps=[],this.parent=void 0,g(this,s)}run(){if(!this.active)return this.fn();let t=E,e=V;for(;t;){if(t===this)return;t=t.parent}try{return this.parent=E,E=this,V=!0,x=1<<++k,k<=30?(({deps:t})=>{if(t.length)for(let e=0;e<t.length;e++)t[e].w|=x})(this):O(this),this.fn()}finally{k<=30&&(t=>{const{deps:e}=t;if(e.length){let s=0;for(let n=0;n<e.length;n++){const i=e[n];b(i)&&!m(i)?i.delete(t):e[s++]=i,i.w&=~x,i.n&=~x}e.length=s}})(this),x=1<<--k,E=this.parent,V=e,this.parent=void 0,this.deferStop&&this.stop()}}stop(){E===this?this.deferStop=!0:this.active&&(O(this),this.onStop&&this.onStop(),this.active=!1)}}function O(t){const{deps:e}=t;if(e.length){for(let s=0;s<e.length;s++)e[s].delete(t);e.length=0}}function P(t,e){t.effect&&(t=t.effect.fn);const s=new M(t);e&&(n(s,e),e.scope&&g(s,e.scope)),e&&e.lazy||s.run();const i=s.run.bind(s);return i.effect=s,i}function W(t){t.effect.stop()}let V=!0;const A=[];function N(){A.push(V),V=!1}function K(){const t=A.pop();V=void 0===t||t}function q(t,e,s){if(V&&E){let e=S.get(t);e||S.set(t,e=new Map);let n=e.get(s);n||e.set(s,n=R()),B(n)}}function B(t,e){let s=!1;k<=30?m(t)||(t.n|=x,s=!b(t)):s=!t.has(E),s&&(t.add(E),E.deps.push(t))}function C(t,e,n,i,r,c){const o=S.get(t);if(!o)return;let a=[];if("clear"===e)a=[...o.values()];else if("length"===n&&s(t)){const t=_(i);o.forEach(((e,s)=>{("length"===s||s>=t)&&a.push(e)}))}else switch(void 0!==n&&a.push(o.get(n)),e){case"add":s(t)?h(n)&&a.push(o.get("length")):(a.push(o.get(j)),l(t)&&a.push(o.get(z)));break;case"delete":s(t)||(a.push(o.get(j)),l(t)&&a.push(o.get(z)));break;case"set":l(t)&&a.push(o.get(j))}if(1===a.length)a[0]&&I(a[0]);else{const t=[];for(const e of a)e&&t.push(...e);I(R(t))}}function I(t,e){const n=s(t)?t:[...t];for(const s of n)s.computed&&D(s);for(const s of n)s.computed||D(s)}function D(t,e){(t!==E||t.allowRecurse)&&(t.scheduler?t.scheduler():t.run())}const F=f("__proto__,__v_isRef,__isVue"),G=new Set(Object.getOwnPropertyNames(Symbol).filter((t=>"arguments"!==t&&"caller"!==t)).map((t=>Symbol[t])).filter(u)),H=X(),J=X(!1,!0),L=X(!0),Q=X(!0,!0),T=U();function U(){const t={};return["includes","indexOf","lastIndexOf"].forEach((e=>{t[e]=function(...t){const s=qt(this);for(let e=0,i=this.length;e<i;e++)q(s,0,e+"");const n=s[e](...t);return-1===n||!1===n?s[e](...t.map(qt)):n}})),["push","pop","shift","unshift","splice"].forEach((e=>{t[e]=function(...t){N();const s=qt(this)[e].apply(this,t);return K(),s}})),t}function X(e=!1,n=!1){return function(i,r,c){if("__v_isReactive"===r)return!e;if("__v_isReadonly"===r)return e;if("__v_isShallow"===r)return n;if("__v_raw"===r&&c===(e?n?jt:Et:n?xt:kt).get(i))return i;const o=s(i);if(!e&&o&&a(T,r))return Reflect.get(T,r,c);const l=Reflect.get(i,r,c);return(u(r)?G.has(r):F(r))?l:(e||q(i,0,r),n?l:Gt(l)?o&&h(r)?l:l.value:t(l)?e?Ot(l):zt(l):l)}}function Y(t=!1){return function(n,i,r,c){let o=n[i];if(At(o)&&Gt(o)&&!Gt(r))return!1;if(!t&&(Nt(r)||At(r)||(o=qt(o),r=qt(r)),!s(n)&&Gt(o)&&!Gt(r)))return o.value=r,!0;const u=s(n)&&h(i)?Number(i)<n.length:a(n,i),l=Reflect.set(n,i,r,c);return n===qt(c)&&(u?e(r,o)&&C(n,"set",i,r):C(n,"add",i,r)),l}}const Z={get:H,set:Y(),deleteProperty:function(t,e){const s=a(t,e);t[e];const n=Reflect.deleteProperty(t,e);return n&&s&&C(t,"delete",e,void 0),n},has:function(t,e){const s=Reflect.has(t,e);return u(e)&&G.has(e)||q(t,0,e),s},ownKeys:function(t){return q(t,0,s(t)?"length":j),Reflect.ownKeys(t)}},$={get:L,set:(t,e)=>!0,deleteProperty:(t,e)=>!0},tt=n({},Z,{get:J,set:Y(!0)}),et=n({},$,{get:Q}),st=t=>t,nt=t=>Reflect.getPrototypeOf(t);function it(t,e,s=!1,n=!1){const i=qt(t=t.__v_raw),r=qt(e);s||(e!==r&&q(i,0,e),q(i,0,r));const{has:c}=nt(i),o=n?st:s?It:Ct;return c.call(i,e)?o(t.get(e)):c.call(i,r)?o(t.get(r)):void(t!==i&&t.get(e))}function rt(t,e=!1){const s=this.__v_raw,n=qt(s),i=qt(t);return e||(t!==i&&q(n,0,t),q(n,0,i)),t===i?s.has(t):s.has(t)||s.has(i)}function ct(t,e=!1){return t=t.__v_raw,!e&&q(qt(t),0,j),Reflect.get(t,"size",t)}function ot(t){t=qt(t);const e=qt(this);return nt(e).has.call(e,t)||(e.add(t),C(e,"add",t,t)),this}function at(t,s){s=qt(s);const n=qt(this),{has:i,get:r}=nt(n);let c=i.call(n,t);c||(t=qt(t),c=i.call(n,t));const o=r.call(n,t);return n.set(t,s),c?e(s,o)&&C(n,"set",t,s):C(n,"add",t,s),this}function ut(t){const e=qt(this),{has:s,get:n}=nt(e);let i=s.call(e,t);i||(t=qt(t),i=s.call(e,t)),n&&n.call(e,t);const r=e.delete(t);return i&&C(e,"delete",t,void 0),r}function ht(){const t=qt(this),e=0!==t.size,s=t.clear();return e&&C(t,"clear",void 0,void 0),s}function lt(t,e){return function(s,n){const i=this,r=i.__v_raw,c=qt(r),o=e?st:t?It:Ct;return!t&&q(c,0,j),r.forEach(((t,e)=>s.call(n,o(t),o(e),i)))}}function ft(t,e,s){return function(...n){const i=this.__v_raw,r=qt(i),c=l(r),o="entries"===t||t===Symbol.iterator&&c,a="keys"===t&&c,u=i[t](...n),h=s?st:e?It:Ct;return!e&&q(r,0,a?z:j),{next(){const{value:t,done:e}=u.next();return e?{value:t,done:e}:{value:o?[h(t[0]),h(t[1])]:h(t),done:e}},[Symbol.iterator](){return this}}}}function _t(t){return function(...e){return"delete"!==t&&this}}function dt(){const t={get(t){return it(this,t)},get size(){return ct(this)},has:rt,add:ot,set:at,delete:ut,clear:ht,forEach:lt(!1,!1)},e={get(t){return it(this,t,!1,!0)},get size(){return ct(this)},has:rt,add:ot,set:at,delete:ut,clear:ht,forEach:lt(!1,!0)},s={get(t){return it(this,t,!0)},get size(){return ct(this,!0)},has(t){return rt.call(this,t,!0)},add:_t("add"),set:_t("set"),delete:_t("delete"),clear:_t("clear"),forEach:lt(!0,!1)},n={get(t){return it(this,t,!0,!0)},get size(){return ct(this,!0)},has(t){return rt.call(this,t,!0)},add:_t("add"),set:_t("set"),delete:_t("delete"),clear:_t("clear"),forEach:lt(!0,!0)};return["keys","values","entries",Symbol.iterator].forEach((i=>{t[i]=ft(i,!1,!1),s[i]=ft(i,!0,!1),e[i]=ft(i,!1,!0),n[i]=ft(i,!0,!0)})),[t,s,e,n]}const[pt,vt,gt,wt]=dt();function yt(t,e){const s=e?t?wt:gt:t?vt:pt;return(e,n,i)=>"__v_isReactive"===n?!t:"__v_isReadonly"===n?t:"__v_raw"===n?e:Reflect.get(a(s,n)&&n in e?s:e,n,i)}const Rt={get:yt(!1,!1)},bt={get:yt(!1,!0)},mt={get:yt(!0,!1)},St={get:yt(!0,!0)},kt=new WeakMap,xt=new WeakMap,Et=new WeakMap,jt=new WeakMap;function zt(t){return At(t)?t:Wt(t,!1,Z,Rt,kt)}function Mt(t){return Wt(t,!1,tt,bt,xt)}function Ot(t){return Wt(t,!0,$,mt,Et)}function Pt(t){return Wt(t,!0,et,St,jt)}function Wt(e,s,n,r,c){if(!t(e))return e;if(e.__v_raw&&(!s||!e.__v_isReactive))return e;const o=c.get(e);if(o)return o;const a=(u=e).__v_skip||!Object.isExtensible(u)?0:function(t){switch(t){case"Object":case"Array":return 1;case"Map":case"Set":case"WeakMap":case"WeakSet":return 2;default:return 0}}(i(u));var u;if(0===a)return e;const h=new Proxy(e,2===a?r:n);return c.set(e,h),h}function Vt(t){return At(t)?Vt(t.__v_raw):!(!t||!t.__v_isReactive)}function At(t){return!(!t||!t.__v_isReadonly)}function Nt(t){return!(!t||!t.__v_isShallow)}function Kt(t){return Vt(t)||At(t)}function qt(t){const e=t&&t.__v_raw;return e?qt(e):t}function Bt(t){return r(t,"__v_skip",!0),t}const Ct=e=>t(e)?zt(e):e,It=e=>t(e)?Ot(e):e;function Dt(t){V&&E&&B((t=qt(t)).dep||(t.dep=R()))}function Ft(t,e){(t=qt(t)).dep&&I(t.dep)}function Gt(t){return!(!t||!0!==t.__v_isRef)}function Ht(t){return Lt(t,!1)}function Jt(t){return Lt(t,!0)}function Lt(t,e){return Gt(t)?t:new Qt(t,e)}class Qt{constructor(t,e){this.__v_isShallow=e,this.dep=void 0,this.__v_isRef=!0,this._rawValue=e?t:qt(t),this._value=e?t:Ct(t)}get value(){return Dt(this),this._value}set value(t){const s=this.__v_isShallow||Nt(t)||At(t);t=s?t:qt(t),e(t,this._rawValue)&&(this._rawValue=t,this._value=s?t:Ct(t),Ft(this))}}function Tt(t){Ft(t)}function Ut(t){return Gt(t)?t.value:t}const Xt={get:(t,e,s)=>Ut(Reflect.get(t,e,s)),set:(t,e,s,n)=>{const i=t[e];return Gt(i)&&!Gt(s)?(i.value=s,!0):Reflect.set(t,e,s,n)}};function Yt(t){return Vt(t)?t:new Proxy(t,Xt)}class Zt{constructor(t){this.dep=void 0,this.__v_isRef=!0;const{get:e,set:s}=t((()=>Dt(this)),(()=>Ft(this)));this._get=e,this._set=s}get value(){return this._get()}set value(t){this._set(t)}}function $t(t){return new Zt(t)}function te(t){const e=s(t)?new Array(t.length):{};for(const s in t)e[s]=se(t,s);return e}class ee{constructor(t,e,s){this._object=t,this._key=e,this._defaultValue=s,this.__v_isRef=!0}get value(){const t=this._object[this._key];return void 0===t?this._defaultValue:t}set value(t){this._object[this._key]=t}}function se(t,e,s){const n=t[e];return Gt(n)?n:new ee(t,e,s)}var ne;class ie{constructor(t,e,s,n){this._setter=e,this.dep=void 0,this.__v_isRef=!0,this[ne]=!1,this._dirty=!0,this.effect=new M(t,(()=>{this._dirty||(this._dirty=!0,Ft(this))})),this.effect.computed=this,this.effect.active=this._cacheable=!n,this.__v_isReadonly=s}get value(){const t=qt(this);return Dt(t),!t._dirty&&t._cacheable||(t._dirty=!1,t._value=t.effect.run()),t._value}set value(t){this._setter(t)}}function re(t,e,s=!1){let n,i;const r=c(t);r?(n=t,i=o):(n=t.get,i=t.set);return new ie(n,i,r||!i,s)}ne="__v_isReadonly";export{At as A,Pt as B,W as C,p as E,M as R,Nt as a,Vt as b,Kt as c,re as d,K as e,Yt as f,q as g,zt as h,Gt as i,C as j,Jt as k,w as l,Bt as m,Ot as n,y as o,N as p,se as q,Ht as r,Mt as s,qt as t,Ut as u,te as v,v as w,Tt as x,$t as y,P as z};
