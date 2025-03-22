var collect=function(){"use strict";var e=window,t=navigator,n=document,r={win:e,doc:n,nav:t,loc:location,ref:n.referrer,domain:n.domain,ua:t.userAgent},i=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e};var o=function e(t){var n=this;!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.set=function(e,t){n._col[e]=t},this.get=function(e){return!!e&&n._col[e]},this.toJSON=function(){return i({},n._col)},this.id=t,this._col={}},s="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},a=Object.hasOwnProperty,c=document,u=function(e){return e=e.toLowerCase(),function(t){var n=void 0===t?"undefined":s(t);return"object"===n?(r=t,Object.prototype.toString.call(r).toLowerCase()==="[object "+e+"]"):n===e;var r}},f=u("Object"),d=u("String"),h=Array.isArray||u("Array"),l=u("Function"),v=u("Undefined"),p=u("Number"),g=function e(t){for(var n=0,r=t.length,i=[];n<r;n++){var o=t[n];o=h(o)?e(o):f(o)?_(o):b(o),i.push(o)}return i.join(",")},_=function e(t){var n=[];return Object.keys(t).forEach(function(r){var i=t[r];i=h(i)?g(i):f(i)?e(i):b(i),n.push(r+"="+i)}),n.join("&")},m=function(e,t){if(e){var n=void 0,r=void 0,i=e.length,o=(e.splice||e.callee)instanceof Function;if(!isNaN(i)&&o)for(n=0;n<i&&(r=e[n],!1!==t.call(null,r,n,e));n++);else for(var s in e)if(a.call(e,s)&&(r=e[s],!1===t.call(null,r,s,e)))break}},y=function(e){var t="EventListener",n="add"+t;return e[n]?{on:function(e,t,r){var i=t.split(" ");m(i,function(t){e[n](t,r,!1)})},off:function(e,n,r){var i=n.split(" ");m(i,function(n){e["remove"+t](n,r,!1)})}}:{on:function(e,t,n){var r=t.split(" ");m(r,function(t){e.attachEvent("on"+t,n)})},off:function(e,t,n){var r=t.split(" ");m(r,function(t){e.detachEvent("on"+t,n)})}}}(window),b=encodeURIComponent,w=decodeURIComponent,S=function(e,t,n,r){var i=void 0,o=void 0,s=void 0,a=e+"="+b(t)+";path=/;domain="+r;void 0!==n&&p(n)&&(i=new Date,o=60*parseInt(n,10)*1e3,s=i.getTime()+o,i.setTime(s),a+=";expires="+i.toUTCString());try{return c.cookie=a,!0}catch(e){return!1}},C=function(e){var t=c.cookie.match(new RegExp("(?:^|;)\\s*"+e+"=([^;]+)"));return w(t?t[1]:"")},O=function(e){return"number"!=typeof e&&(void 0===e||(null===e||(""===e||void 0)))},E={parseParams:function(e){var t=[];return h(e)?t.push(g(e)):f(e)?t.push(_(e)):t.push(b(e)),t.join("&")},isNumber:p,isString:d,isFunction:l,isArray:h,isObject:f,keys:Object.keys,undef:v,hasOwn:function(e,t){return a.call(e,t)},merge:function e(t,n){if(!n)return t;var r=void 0,i=void 0,o=void 0;return t=t||{},n instanceof Array?[]:(Object.keys(n).forEach(function(a){o=n[a],r=s(t[a]),i=s(n[a]),"object"===r?t[a]=e(t[a],n[a]):!t[a]&&o&&(t[a]="object"===i?e({},o):o)}),t)},now:function(){return parseInt(+new Date/1e3,10)},on:y.on,off:y.off,each:m,getCookie:C,topCookie:function(e,t,n){var r=c.domain,i=r.split("."),o=i.length,s=o-1,a=void 0,u=!1;for(t=""+t;!u&&s>=0&&(r=i.slice(s,o).join("."),S(e,t,n,r),void 0!==(a=C(e))&&(u=a===t),s--,!u););},parseURL:function(e){var t=document.createElement("a");return t.href=e,t.hostname},isNil:O,convertObj:function(e,t){var n=Object.create(null);return e&&"string"==typeof e&&!O(t)?n[e]=String(t):"object"===(void 0===e?"undefined":s(e))&&(n=e),n}},k=E.parseURL,x=screen.width||0,D=screen.height||0,P=x+" x "+D,T=navigator.appVersion,A=navigator.userAgent,j=navigator.language,I=document.referrer,R=k(I),q="",N="",M="",W=""+parseFloat(T),H=void 0,F=void 0;-1!==(H=A.indexOf("Opera"))&&(M="Opera",W=A.substring(H+6),-1!==(H=A.indexOf("Version"))&&(W=A.substring(H+8))),-1!==(H=A.indexOf("Edge"))?(M="Microsoft Edge",W=A.substring(H+5)):-1!==(H=A.indexOf("MSIE"))?(M="Microsoft Internet Explorer",W=A.substring(H+5)):-1!==(H=A.indexOf("Chrome"))?(M="Chrome",W=A.substring(H+7)):-1!==(H=A.indexOf("Safari"))?(M="Safari",W=A.substring(H+7),-1!==(H=A.indexOf("Version"))&&(W=A.substring(H+8))):-1!==(H=A.indexOf("Firefox"))&&(M="Firefox",W=A.substring(H+8)),-1!==(F=W.indexOf(";"))&&(W=W.substring(0,F)),-1!==(F=W.indexOf(" "))&&(W=W.substring(0,F)),-1!==(F=W.indexOf(")"))&&(W=W.substring(0,F));for(var L=/Mobile|htc|mini|Android|iP(ad|od|hone)/.test(T)?"wap":"web",X=[{s:"Windows 10",r:/(Windows 10.0|Windows NT 10.0)/},{s:"Windows 8.1",r:/(Windows 8.1|Windows NT 6.3)/},{s:"Windows 8",r:/(Windows 8|Windows NT 6.2)/},{s:"Windows 7",r:/(Windows 7|Windows NT 6.1)/},{s:"Android",r:/Android/},{s:"Sun OS",r:/SunOS/},{s:"Linux",r:/(Linux|X11)/},{s:"iOS",r:/(iPhone|iPad|iPod)/},{s:"Mac OS X",r:/Mac OS X/},{s:"Mac OS",r:/(MacPPC|MacIntel|Mac_PowerPC|Macintosh)/}],U=0;U<X.length;U++){var z=X[U];if(z.r.test(A)){q=z.s;break}}switch(/Windows/.test(q)&&(N=/Windows (.*)/.exec(q)[1],q="windows"),q){case"Mac OS X":N=/Mac OS X (10[\.\_\d]+)/.exec(A)[1],q="mac";break;case"Android":N=/Android ([\.\_\d]+)/.exec(A)[1],q="android";break;case"iOS":N=(N=/OS (\d+)_(\d+)_?(\d+)?/.exec(T))[1]+"."+N[2]+"."+(0|N[3]),q="ios"}var J={screen_size:P,browser:M,browser_version:W,platform:L,os_name:q,os_version:N,userAgent:A,screen_width:x,screen_height:D,device_model:q,language:j,referrer:I,referrer_host:R};var K=(new Date).getTimezoneOffset(),B=-K/60,V=60*K,G=void 0,Q=void 0;try{Q="2.2.11"}catch(e){Q="2.x"}var $=new function e(){!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this._init=function(){var e={app_id:0,app_name:"",os_name:J.os_name,os_version:J.os_version+"",device_model:J.device_model,headers:{platform:J.platform,sdk_version:Q,browser:J.browser,browser_version:J.browser_version,region:"",province:"",city:"",language:J.language,timezone:B,tz_offset:V,screen_width:J.screen_width,screen_height:J.screen_height,referrer:J.referrer,referrer_host:J.referrer_host,custom:{}}};(G=new o("_default")).set("env",e)},this.toJSON=function(){return G.get("env")},this._init()},Y="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},Z=!1,ee=function(e){if(Z)throw new Error(e);console.error(e)},te=function(){if(Z){var e;(e=console).warn.apply(e,arguments)}},ne=function(){if(Z){var e;(e=console).log.apply(e,arguments)}},re=function(e){void 0===e||Y(e)},ie=function(){var e=!(arguments.length>0&&void 0!==arguments[0])||arguments[0];Z=e},oe="__tea_sdk__",se=function(e){return oe+e},ae=function(){var e={},t=void 0,n=window,r=document,i=void 0,o="localStorage"in n&&n.localStorage;if(e.set=function(e,t){},e.get=function(e,t){},e.has=function(n){return e.get(n)!==t},e.remove=function(e){},e.clear=function(){},e.forEach=function(){},e.serialize=function(e){return JSON.stringify(e)},e.deserialize=function(e){if("string"!=typeof e)return t;try{return JSON.parse(e)}catch(n){return e||t}},o)i=n.localStorage,e.set=function(n,r){return r===t?e.remove(n):(i.setItem(n,e.serialize(r)),r)},e.get=function(n,r){var o=e.deserialize(i.getItem(n));return o===t?r:o},e.remove=function(e){i.removeItem(e)},e.clear=function(){i.clear()},e.forEach=function(t){for(var n=0;n<i.length;n++){var r=i.key(n);t(r,e.get(r))}};else if(r&&r.documentElement.addBehavior){var s=void 0,a=void 0;try{(a=new ActiveXObject("htmlfile")).open(),a.write('<script>document.w=window<\/script><iframe src="/favicon.ico"></iframe>'),a.close(),s=a.w.frames[0].document,i=s.createElement("div")}catch(e){i=r.createElement("div"),s=r.body}var c=function(t){return function(){var n=Array.prototype.slice.call(arguments,0);n.unshift(i),s.appendChild(i),i.addBehavior("#default#userData"),i.load("localStorage");var r=t.apply(e,n);return s.removeChild(i),r}},u=new RegExp("[!\"#$%&'()*+,/\\\\:;<=>?@[\\]^`{|}~]","g"),f=function(e){return e.replace(/^d/,"___$&").replace(u,"___")};e.set=c(function(n,r,i){return r=f(r),i===t?e.remove(r):(n.setAttribute(r,e.serialize(i)),n.save("localStorage"),i)}),e.get=c(function(n,r,i){r=f(r);var o=e.deserialize(n.getAttribute(r));return o===t?i:o}),e.remove=c(function(e,t){t=f(t),e.removeAttribute(t),e.save("localStorage")}),e.clear=c(function(e){var t=e.XMLDocument.documentElement.attributes;e.load("localStorage");for(var n=t.length-1;n>=0;n--)e.removeAttribute(t[n].name);e.save("localStorage")}),e.forEach=c(function(t,n){for(var r,i=t.XMLDocument.documentElement.attributes,o=0;r=i[o];++o)n(r.name,e.deserialize(t.getAttribute(r.name)))})}return e}(),ce=function(e){return ae.get(se(e))},ue=function(e,t){return ae.set(se(e),t)};var fe=new function e(){!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.get=function(e){var t="";try{t=ce(e)||{}}catch(e){t={}}return t},this.set=function(e,t){try{ue(e,t)}catch(e){}}},de=+new Date+Math.floor(1e5*Math.random()),he=function(){},le=function(e,t,n){try{n("success",ve,e,t,n);var r=e.split("v1")[0];r||console.error("上报 url 不对");t.payload.forEach(function(e){var t=function(e){var t="";for(var n in e)e.hasOwnProperty(n)&&(t+="&"+n+"="+encodeURIComponent(JSON.stringify(e[n])));return t="&"===t[0]?t.slice(1):t}(e),n=new Image(1,1);n.onload=function(){n=null},n.onerror=function(e){n=null,console.error("tea-sdk 上报数据失败："+e.message)},n.src=r+"/v1/gif?"+t})}catch(e){console.error(e.message)}},ve=function(e,t,n){if(!t||t.payload&&!t.payload.length)return!1;if(window.XDomainRequest)return le(e,t,n);-1===e.indexOf("tea_sdk_random")&&(e=e+"?tea_sdk_random="+(de++).toString(16));try{!function(e,t,n){if(n=n||he,window.navigator&&window.navigator.sendBeacon&&window.navigator.sendBeacon(e,JSON.stringify(t.payload)))return n("success",ve,e,t,n);var r=new XMLHttpRequest;r.open("POST",e,!0),r.timeout=15e3,r.onreadystatechange=function(){if(4===r.readyState){if(r.status>=200&&r.status<300||304===r.status){n("success",ve,e,t,n);try{var i=JSON.parse(r.responseText);0!==i.e&&(te("------------------------------TEA SDK 上报数据失败------------------------------"),te("当前上报数据 token：",t.token),te("错误状态码: ",i.e),-2===i.e&&te("字段格式错误，很有可能是把 int 写成了 string"),te("                                                                       "))}catch(e){console.error(e.message)}}else n("fail",ve,e,t,n);r.onreadystatechange=null}},r.onerror=function(){r.abort(),n("fail",ve,e,t,n)},r.send(JSON.stringify(t.payload))}(e,t,n)}catch(e){console.error("统计数据发送失败",e.message)}return!1},pe={post:ve,get:function(e){var t=new XMLHttpRequest;t.open("GET",e,!0),t.onreadystatechange=function(){4===t.readyState&&(t.status>=200&&window.alert("sasas"),t.onreadystatechange=null)},t.onerror=function(){t.abort()},t.send(null)}};var ge=new function e(){var t=this;!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.postFailed=function(e,n,r,i){if(r.repeat)return t.REPORT_TOKENS[r.token]=!1,!1;r.repeat=!0,setTimeout(function(){te("------------------------------再次尝试发送刚刚上报失败了的数据------------------------------"),e(n,r,i)},1e3)},this.postSuccess=function(e,t,n){ne("当前上报数据 token：",n.token),ne("------------------------------TEA SDK 上报数据成功------------------------------"),ne("                                                                       "),n=n||{};var r=fe.get("ls_cache__")||{};delete r[n.token],fe.set("ls_cache__",r)},this.sortOut=function(e,n,r,i,o){switch(e){case"success":return t.postSuccess(n,r,i);case"fail":default:return t.postFailed(n,r,i,o)}},this.report=function(e){var n=e.url;delete e.url,t.REPORT_TOKENS[e.token]=!0,ne(""),ne("------------------------------TEA SDK 上报数据开始------------------------------"),ne("上报路径：",n),ne("当前上报数据 token：",e.token),ne("上报数据实体：",e.payload),ne("                                                                       "),pe.post(n,e,t.sortOut)},this.REPORT_TOKENS={}},_e="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},me=function(e){var t,n;t=Array.prototype.forEach,n=Array.prototype.map,this.each=function(e,n,r){if(null!==e)if(t&&e.forEach===t)e.forEach(n,r);else if(e.length===+e.length){for(var i=0,o=e.length;i<o;i++)if(n.call(r,e[i],i,e)==={})return}else for(var s in e)if(e.hasOwnProperty(s)&&n.call(r,e[s],s,e)==={})return},this.map=function(e,t,r){var i=[];return null==e?i:n&&e.map===n?e.map(t,r):(this.each(e,function(e,n,o){i[i.length]=t.call(r,e,n,o)}),i)},"object"==(void 0===e?"undefined":_e(e))?(this.hasher=e.hasher,this.screen_resolution=e.screen_resolution,this.screen_orientation=e.screen_orientation,this.canvas=e.canvas,this.ie_activex=e.ie_activex):"function"==typeof e&&(this.hasher=e)};me.prototype={get:function(){var e=[];if(e.push(navigator.userAgent),e.push(navigator.language),e.push(screen.colorDepth),this.screen_resolution){var t=this.getScreenResolution();void 0!==t&&e.push(t.join("x"))}return e.push((new Date).getTimezoneOffset()),e.push(this.hasSessionStorage()),e.push(this.hasLocalStorage()),e.push(this.hasIndexDb()),document.body?e.push(_e(document.body.addBehavior)):e.push("undefined"),e.push(_e(window.openDatabase)),e.push(navigator.cpuClass),e.push(navigator.platform),e.push(navigator.doNotTrack),e.push(this.getPluginsString()),this.canvas&&this.isCanvasSupported()&&e.push(this.getCanvasFingerprint()),this.hasher?this.hasher(e.join("###"),31):this.murmurhash3_32_gc(e.join("###"),31)},murmurhash3_32_gc:function(e,t){var n,r,i,o,s,a,c,u;for(n=3&e.length,r=e.length-n,i=t,s=3432918353,a=461845907,u=0;u<r;)c=255&e.charCodeAt(u)|(255&e.charCodeAt(++u))<<8|(255&e.charCodeAt(++u))<<16|(255&e.charCodeAt(++u))<<24,++u,i=27492+(65535&(o=5*(65535&(i=(i^=c=(65535&(c=(c=(65535&c)*s+(((c>>>16)*s&65535)<<16)&4294967295)<<15|c>>>17))*a+(((c>>>16)*a&65535)<<16)&4294967295)<<13|i>>>19))+((5*(i>>>16)&65535)<<16)&4294967295))+((58964+(o>>>16)&65535)<<16);switch(c=0,n){case 3:c^=(255&e.charCodeAt(u+2))<<16;case 2:c^=(255&e.charCodeAt(u+1))<<8;case 1:i^=c=(65535&(c=(c=(65535&(c^=255&e.charCodeAt(u)))*s+(((c>>>16)*s&65535)<<16)&4294967295)<<15|c>>>17))*a+(((c>>>16)*a&65535)<<16)&4294967295}return i^=e.length,i=2246822507*(65535&(i^=i>>>16))+((2246822507*(i>>>16)&65535)<<16)&4294967295,i=3266489909*(65535&(i^=i>>>13))+((3266489909*(i>>>16)&65535)<<16)&4294967295,(i^=i>>>16)>>>0},hasLocalStorage:function(){try{return!!window.localStorage}catch(e){return!0}},hasSessionStorage:function(){try{return!!window.sessionStorage}catch(e){return!0}},hasIndexDb:function(){try{return!!window.indexedDB}catch(e){return!0}},isCanvasSupported:function(){var e=document.createElement("canvas");return!(!e.getContext||!e.getContext("2d"))},isIE:function(){return"Microsoft Internet Explorer"===navigator.appName||!("Netscape"!==navigator.appName||!/Trident/.test(navigator.userAgent))},getPluginsString:function(){return this.isIE()&&this.ie_activex?this.getIEPluginsString():this.getRegularPluginsString()},getRegularPluginsString:function(){return this.map(navigator.plugins,function(e){var t=this.map(e,function(e){return[e.type,e.suffixes].join("~")}).join(",");return[e.name,e.description,t].join("::")},this).join(";")},getIEPluginsString:function(){if(window.ActiveXObject){return this.map(["ShockwaveFlash.ShockwaveFlash","AcroPDF.PDF","PDF.PdfCtrl","QuickTime.QuickTime","rmocx.RealPlayer G2 Control","rmocx.RealPlayer G2 Control.1","RealPlayer.RealPlayer(tm) ActiveX Control (32-bit)","RealVideo.RealVideo(tm) ActiveX Control (32-bit)","RealPlayer","SWCtl.SWCtl","WMPlayer.OCX","AgControl.AgControl","Skype.Detection"],function(e){try{return new ActiveXObject(e),e}catch(e){return null}}).join(";")}return""},getScreenResolution:function(){return this.screen_orientation?screen.height>screen.width?[screen.height,screen.width]:[screen.width,screen.height]:[screen.height,screen.width]},getCanvasFingerprint:function(){var e=document.createElement("canvas"),t=e.getContext("2d"),n="http://valve.github.io";return t.textBaseline="top",t.font="14px 'Arial'",t.textBaseline="alphabetic",t.fillStyle="#f60",t.fillRect(125,1,62,20),t.fillStyle="#069",t.fillText(n,2,15),t.fillStyle="rgba(102, 204, 0, 0.7)",t.fillText(n,4,17),e.toDataURL()}};var ye=function(e){return e.indexOf("web")>-1?e:"__tea_sdk__"+e},be=function(e,t,n,r){if(window.XDomainRequest)return r&&r(),!1;var i=new XMLHttpRequest;i.open("POST",e,!0),i.timeout=5e3,i.setRequestHeader("Content-Type","application/json; charset=utf-8"),i.onreadystatechange=function(){if(4===i.readyState){try{var e=JSON.parse(i.responseText);n&&n(e)}catch(e){r&&r()}i.onreadystatechange=null}},i.ontimeout=function(){r&&r()},i.onerror=function(){i.abort(),r&&r()},i.send(JSON.stringify(t))},we=new function e(){var t=this;!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.verify=function(e,n,r){return"function"==typeof r&&(t.cb=r),t.app_id=e,t.urlPrefix=n,t.web_id.length<=11?t.fetchWebId():-1===t.ssid.indexOf("-")?t.fetchSSID():void 0},this.forceReport=function(){t.cb&&t.cb()},this.fetchWebId=function(){ne("开始获取 ssid 和 webid"),t.isFetchingWebId=!0,be(t.urlPrefix+"/v1/user/webid",{app_id:t.app_id,url:location.href,user_agent:J.userAgent,referer:J.referrer,user_unique_id:""},function(e){ne("ssid 和webid获取成功，按照约定流程复制和设置 cookie"),ne("注：ssid 值 "+e.ssid+"，可能还会因为您设置 user_unique_id 而变更"),ne("注：webid 值"+e.web_id+"，是与浏览器特征挂钩的，除非你主动清除 cookie，否则这个值不再变化~");var n=e.ssid,r=e.web_id;t.ssid=n,t.setCookie("ssid",t.ssid),t.web_id=r,t.setCookie("tt_webid",t.web_id),t.user_unique_id=r,t.setCookie("user_unique_id",t.user_unique_id),t.isFetchingWebId=!1,t.forceReport(),t.real_uuid&&t.verifySSID(t.real_uuid)},function(){te("webid获取失败，本地临时生成 web_id，作为上报凭证；但是绝对不允许写 cookie");var e=new me({canvas:!0,screen_resolution:!0}).get()+"";t.web_id=e,t.ssid=e,t.user_unique_id=e})},this.verifySSID=function(e){return t.real_uuid=e,!t.isFetchingWebId&&(t.real_uuid===t.user_unique_id&&t.web_id&&t.ssid?(t.isSSIDUpdated=!0,ne("ssid 无需更新，直接发送事件"),t.forceReport(),!1):(t.user_unique_id=t.real_uuid,t.setCookie("user_unique_id",t.user_unique_id),void t.fetchSSID()))},this.fetchSSID=function(){be(t.urlPrefix+"/v1/user/ssid",{app_id:t.app_id,web_id:t.web_id,user_unique_id:""+t.user_unique_id},function(e){t.isSSIDUpdated=!0,ne("ssid 获取成功,您user_unique_id值"+t.user_unique_id+"对应的真实 ssid 为"+e.ssid),t.ssid=e.ssid,t.setCookie("ssid",t.ssid),t.forceReport()},function(){t.isSSIDUpdated=!0,te("单独获取 ssid 失败~"),t.forceReport()})},this.setCookie=function(e,t){return E.topCookie(ye(e),t,15768e3)},this.readCookies=function(e){var t=E.getCookie(ye(e));return"undefined"!==t&&t||(t=""),t},this.getCookieData=function(){return{isSSIDUpdated:t.isSSIDUpdated,cookies:{user_unique_id:t.user_unique_id,web_id:t.web_id,ssid:t.ssid}}},this.user_unique_id=this.readCookies("user_unique_id"),this.web_id=this.readCookies("tt_webid")||this.readCookies("tt_web_id"),this.ssid=this.readCookies("ssid"),this.isFetchingWebId=!1,this.isSSIDUpdated=!1};var Se="https://mcs.byted.org",Ce="https://sgali-mcs.byteoversea.com",Oe="https://mcs.snssdk.com",Ee={va_out:"https://vaali-mcs.byteoversea.com",sg_out:Ce,cn_out:Oe,cn_inner:Se,web_id_cn_inner:Se,web_id_cn_out:Oe,web_id_sg_out:Ce},ke=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e};var xe=new function e(t){var n=this;!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.setEvtCache=function(e){n.evtCache.push(e)},this.getEvtCache=function(){var e=n.evtCache;return n.evtCache=[],e},this.hasEvtCache=function(){return n.evtCache.length},this.getLen=function(){return n.evtCache.length},this.evtCacheMax=t,this.evtCache=[]}(5),De=1e3*(+new Date+Math.ceil(1e4*Math.random())),Pe=0,Te=E.isNil,Ae=function e(t){var n=this;!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e),this.set=function(e,t){switch(n._col.set(e,t),e){case"user":case"head":case"eventCommonParams":n.forceEvent()}},this.get=function(e){return n._col.get(e)},this.mergeCookie=function(){var e=n._col.toJSON(),t=we.getCookieData().cookies;return Object.keys(t).forEach(function(n){if("user_unique_id"===n&&e.user[n])return!1;e.user[n]=t[n]}),e},this.isChannelReady=function(e){var t=e.isDisableTeaTrack,n=[],r=e.required;E.keys(r).forEach(function(i){var o="";switch(i){case"customHeaders":o=e.header.headers.custom;break;default:o=e[i]}r[i].forEach(function(e){var r=o[e];if(Te(r)&&n.push(e),"user_unique_id"===e&&!t){we.getCookieData().isSSIDUpdated||(n.push("根据user_unique_id 更新的 ssid"),ne("因为您 主动设置了 user_unique_id ，ssid 正在进行验证、更新，所以暂时不能发送事件，必须等 ssid 回调触发"))}})});return n.length?(te("------------------------------TEA SDK 缺失以下预定的必传字段------------------------------"),te("本次事件上报失败，因为缺失上报必备参数，如下：",n),ne("                                                                       ")):(ne("------------------------------字段检测全部ok，相关字段如下------------------------------"),re(r)),0===n.length},this.event=function(e){xe.setEvtCache(e),xe.getLen()>xe.evtCacheMax?n.report():setTimeout(n.report,1)},this.forceEvent=function(){n.report(!0)},this.report=function(){if(!n.get("reportPrefix"))return te("找不到上报地址前缀，应该是你还没有调用setIntranetMode"),!1;var e=n.composeData();e&&ge.report(e)},this.composeData=function(){var e=[],t=fe.get("ls_cache__");Object.keys(t).forEach(function(r){if(r.indexOf(n.id+"_")>-1&&!ge.REPORT_TOKENS[r]){var i=t[r];e.push.apply(e,i),delete t[r]}});var r=n.composeCacheEvtData();if(e.push.apply(e,r),!e.length)return!1;var i=n.id+"_"+(De+Pe++);t[i]=e,fe.set("ls_cache__",t);var o=n.get("reportPrefix")||Ee[window.__report_zone__];return{payload:e,url:o+"/v1/list",token:i}},this.composeCacheEvtData=function(){if(!xe.hasEvtCache())return[];var e={},t=xe.getEvtCache(),r=n.getChannelInfo();if(!r)return t.forEach(function(e){xe.setEvtCache(e)}),[];var i=r.eventCommonParams||{};t=n.mergeEvtCommonParams(t,i),e.events=t;var o=r.header,s=r.user;return Object.keys(s).forEach(function(e){s[e]||(ne(e+"的值为空，将会从上报数据中移除"),delete s[e])}),e.user=s,e.header=o,[e]},this.mergeEvtCommonParams=function(e,t){var n=E.keys(t);return e.forEach(function(e){n.forEach(function(n){Te(e.params[n])&&(e.params[n]=t[n])}),e.params=JSON.stringify(e.params)}),e},this.getChannelInfo=function(){var e=n.mergeCookie();if(!n.isChannelReady(e))return!1;var t=e.header,r=$.toJSON();r=ke({},r);var i=ke({},r.headers);t=ke({},t);var o=ke({},t.headers),s=ke({},i.custom),a=ke({},o.custom),c=ke({},s,a),u=ke({},i,o);u.custom=c,u=JSON.stringify(u);var f=ke({},r,t);f.headers=u,e.header=f;try{e.header.os_version=e.header.os_version+""}catch(e){}return ke({},e)},this.id=t,this._col=new o(t),this.set("isDisableHttps",!1),this.set("isDisableTeaTrack",!1),this.set("user",{}),this.set("eventCommonParams",{}),this.set("header",{app_id:Number(t),headers:{custom:{}}}),this.set("required",{user:["ssid","web_id"],header:[],eventCommonParams:[],customHeaders:[]})},je={},Ie=function(e){var t=je[e];return t||(t=new Ae(e),je[e]=t),t},Re={user_id:1,web_id:1,user_unique_id:1,ssid:1,user_type:1},qe={app_id:1,app_name:1,os_name:1,os_version:"1",device_model:1,ab_client:1,ab_version:1,traffic_type:1,utm_source:1,utm_medium:1,campaign:1,headers:{platform:1,sdk_version:1,browser:1,browser_version:1,region:1,province:1,city:1,language:1,timezone:1,tz_offset:1,screen_height:1,screen_width:1,referrer:1,referrer_host:1}},Ne=Date.now(),Me="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},We=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},He=E.isNil,Fe=r.win,Le=r.doc,Xe=r.loc,Ue=!1,ze=Fe[Fe.TeaAnalyticsObject];ze||((ze=function(){for(var e=arguments.length,t=Array(e),n=0;n<e;n++)t[n]=arguments[n];return ze.q.push(t),ze}).q=ze.q||[],ze.l=+new Date);return function(){var e=ze.l,t=!1,n=void 0,i=ze;i.setAppId=function(e){"number"!=typeof e&&ee("app_id 是产品线的识别id，这里必须设置 & 类型是一个数字"),i.use(e)},i.use=function(e){if("number"!=typeof e)return ee("请输入有效 app_id"),!1;n=Ie(e)},i.setUser=function(e,t){var r=E.convertObj(e,t);delete r.web_id,delete r.ssid;var i=n.get("user"),o=E.keys(r),s=r.user_unique_id||r.user_id||"";r.user_unique_id=s,o.forEach(function(t){var n=e[t];He(n)&&ee("setUser 方法对应的 value 不允许为空"),Re[t]&&(i[t]="user_type"===t?Number(n):String(n))}),n.set("user",i),s&&we.verifySSID(String(s))},i.setHeader=function(e,t){var r=n.get("header");E.isObject(e)?E.keys(e).forEach(function(t){var n=e[t];He(n)&&ee("setHeader 方法对应的 value 不允许为空"),qe[t]&&(r[t]=n)}):(He(t)&&ee("setHeader 方法对应的 value 不允许为空"),qe[e]&&(r[e]=t)),n.set("header",r)},i.setHeaderHeaders=function(e,t){var r=n.get("header");r.headers||(r.headers={}),E.isObject(e)?E.keys(e).forEach(function(t){var n=e[t];He(n)&&ee("setHeaderHeaders 方法对应的 value 不允许为空"),qe.headers[t]&&(r.headers[t]=n)}):(He(t)&&ee("setHeaderHeaders 方法对应的 value 不允许为空"),qe.headers[e]&&(r.headers[e]=t)),n.set("header",r)},i.setCustomHeader=function(e,t){var r=n.get("header");r.headers||(r.headers={}),r.headers.custom||(r.headers.custom={}),E.isObject(e)?E.keys(e).forEach(function(t){var n=e[t];He(n)&&ee("setCustomHeader 方法对应的 value 不允许为空"),r.headers.custom[t]=n}):(He(t)&&ee("setCustomHeader 方法对应的 value 不允许为空"),r.headers.custom[e]=t),n.set("header",r)},i.setDebug=function(){var e=arguments.length>0&&void 0!==arguments[0]&&arguments[0]?1:0;i.setEventCommonParams("_staging_flag",e)},i.setTeaTrack=function(e){if(!e){n.set("isDisableTeaTrack",!0);var t=n.get("required"),r=t.user.filter(function(e){return"ssid"!==e&&"web_id"!==e&&e});t.user=r,n.set("required",t)}},i.setRequiredKeys=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{user:[],header:[],eventCommonParams:[],customHeaders:[]},t=n.get("required");E.keys(e).forEach(function(n){var r=t[n]||[],i=e[n]||[];E.isArray(i)||(i=[i]);var o=function(e){for(var t=[],n={},r=0,i=e.length;r<i;r++){var o=e[r],s=(void 0===o?"undefined":Me(o))+o;1!==n[s]&&(t.push(o),n[s]=1)}return t}(r.concat(i));t[n]=o}),n.set("required",t)},i.setEventCommonParams=function(e,t){var r=n.get("eventCommonParams");E.isObject(e)?E.keys(e).forEach(function(t){var n=e[t];He(n)&&ee("setEventCommonParams 方法对应的 value 不允许为空"),r[t]=n}):(He(t)&&ee("setEventCommonParams 方法对应的 value 不允许为空"),r[e]=t),n.set("eventCommonParams",r)},i.disableHttps=function(e){n.set("isDisableHttps",e)},i.getWebIdArea=function(e){var t="web_id_"+e;return t.indexOf("va")>-1&&(t="web_id_sg_out"),t},i.setIntranetMode=function(){var e=arguments.length>0&&void 0!==arguments[0]&&arguments[0],t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"cn";"true"===e&&(e=!0),"false"===e&&(e=!1);var r="";r="cn"===t?Ee.cn_out:Ee.sg_out;var o=Ee[t+"_"+(e?"inner":"out")];o||ee("setIntranetMode 方法的参数设置错误，导致无法设置上报路径，请参考文档重新设置"),r||ee("setIntranetMode 方法的参数设置错误，导致无法设置 webid 获取路径，请参考文档重新设置"),we.verify(n.id,r,i._forceReport),n.set("reportPrefix",o)},i.predefine_pageview=function(){var n=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},r={title:Le.title||Xe.pathname,url:Xe.href,url_path:Xe.pathname},o={event:"predefine_pageview",params:We({},r,n),local_time_ms:t?+new Date:e};t=!0,i._event(o)},i._event=function(e){n&&(e.params.event_index=++Ne,n.event(e))},i._forceReport=function(){return n.forceEvent(),i},i.setLog=function(e){ie(e)};var o=function(){i._forceReport.call(i)},s=function(){if(Ue)return!1;Ue=!0,function(e,t){var r=function(e){var t=(e=Array.prototype.slice.call(e))[0],r=t.split("."),o=e.slice(1),s=void 0,a=void 0;r.length>1&&(a=n,s=r[0],t=r[1]),s&&i.use(Number(s));var c=i[t];if(E.isFunction(c))c.apply(i,o);else{var u=o[0];E.isObject(u)||(u={}),i._event({event:t,params:u,local_time_ms:+new Date})}s&&i.use(a.id)};if(t.q.push=function(e){return r(e),t},!e.length)return!1;for(var o=0,s=[],a=[],c=[];o<e.length;o++){var u=e[o][0];u.indexOf("set")>-1?"setAppId"===u?s.unshift(e[o]):"setTeaTrack"===u?a.unshift(e[o]):"setRequiredKeys"===u?s.push(e[o]):a.push(e[o]):c.push(e[o])}(e=s.concat(a,c)).forEach(function(e){if(!e.length)return!1;r(e)}),i.predefine_pageview()}(ze.q,ze),E.on(Fe,"unload",o),E.on(Fe,"beforeunload",o)};return"complete"!==r.doc.readyState?(E.on(r.doc,"DOMContentLoaded",s),E.on(r.win,"load",s)):s(),i}()}();
