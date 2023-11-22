"use strict";
!function (t, e) {
    "use strict";
    function n(t) { if (t && t.__esModule)
        return t; var e = Object.create(null); if (t)
        for (var n in t)
            e[n] = t[n]; return e.default = t, Object.freeze(e); }
    var r = n(e);
    class i {
        constructor() { this.listeners = []; }
        tap(t) { return this.listeners.push(t), () => this.revoke(t); }
        revoke(t) { const e = this.listeners.indexOf(t); e >= 0 && this.listeners.splice(e, 1); }
        revokeAll() { this.listeners.splice(0); }
        call(...t) { for (const e of this.listeners)
            e(...t); }
    }
    function a() { return a = Object.assign || function (t) { for (var e = 1; e < arguments.length; e++) {
        var n = arguments[e];
        for (var r in n)
            Object.prototype.hasOwnProperty.call(n, r) && (t[r] = n[r]);
    } return t; }, a.apply(this, arguments); }
    const o = Math.random().toString(36).slice(2, 8);
    let s = 0;
    function l() { }
    function c(t, e, n = "children") { const r = (t, i) => e(t, (() => { var e; null == (e = t[n]) || e.forEach((e => { r(e, t); })); }), i); r(t); }
    function h(t) { if (Array.from)
        return Array.from(t); const e = []; for (let n = 0; n < t.length; n += 1)
        e.push(t[n]); return e; }
    function d(t) { if ("string" == typeof t) {
        const e = t;
        t = t => t.tagName === e;
    } const e = t; return function () { let t = h(this.childNodes); return e && (t = t.filter((t => e(t)))), t; }; }
    function u(t, e, n) { const r = document.createElement(t); return e && Object.entries(e).forEach((([t, e]) => { r[t] = e; })), n && Object.entries(n).forEach((([t, e]) => { r.setAttribute(t, e); })), r; }
    const p = function (t) { const e = {}; return function (...n) { const r = `${n[0]}`; let i = e[r]; return i || (i = { value: t(...n) }, e[r] = i), i.value; }; }((t => { document.head.append(u("link", { rel: "preload", as: "script", href: t })); }));
    async function f(t, e) { if (!t.loaded && ("script" === t.type && (t.loaded = new Promise(((e, n) => { var r; document.head.append(u("script", a({}, t.data, { onload: e, onerror: n }))), null != (r = t.data) && r.src || e(void 0); })).then((() => { t.loaded = !0; }))), "iife" === t.type)) {
        const { fn: n, getParams: r } = t.data;
        n(...(null == r ? void 0 : r(e)) || []), t.loaded = !0;
    } await t.loaded; }
    function m(t) { t.loaded || (t.loaded = !0, "style" === t.type ? document.head.append(u("style", { textContent: t.data })) : "stylesheet" === t.type && document.head.append(u("link", a({ rel: "stylesheet" }, t.data)))); }
    function g() { return g = Object.assign || function (t) { for (var e = 1; e < arguments.length; e++) {
        var n = arguments[e];
        for (var r in n)
            Object.prototype.hasOwnProperty.call(n, r) && (t[r] = n[r]);
    } return t; }, g.apply(this, arguments); }
    function y(t) { var e = 0, n = t.children, r = n && n.length; if (r)
        for (; --r >= 0;)
            e += n[r].value;
    else
        e = 1; t.value = e; }
    function v(t, e) { var n, r, i, a, o, s = new b(t), l = +t.value && (s.value = t.value), c = [s]; for (null == e && (e = x); n = c.pop();)
        if (l && (n.value = +n.data.value), (i = e(n.data)) && (o = i.length))
            for (n.children = new Array(o), a = o - 1; a >= 0; --a)
                c.push(r = n.children[a] = new b(i[a])), r.parent = n, r.depth = n.depth + 1; return s.eachBefore(z); }
    function x(t) { return t.children; }
    function k(t) { t.data = t.data.data; }
    function z(t) { var e = 0; do {
        t.height = e;
    } while ((t = t.parent) && t.height < ++e); }
    function b(t) { this.data = t, this.depth = this.height = 0, this.parent = null; }
    b.prototype = v.prototype = { constructor: b, count: function () { return this.eachAfter(y); }, each: function (t) { var e, n, r, i, a = this, o = [a]; do {
            for (e = o.reverse(), o = []; a = e.pop();)
                if (t(a), n = a.children)
                    for (r = 0, i = n.length; r < i; ++r)
                        o.push(n[r]);
        } while (o.length); return this; }, eachAfter: function (t) { for (var e, n, r, i = this, a = [i], o = []; i = a.pop();)
            if (o.push(i), e = i.children)
                for (n = 0, r = e.length; n < r; ++n)
                    a.push(e[n]); for (; i = o.pop();)
            t(i); return this; }, eachBefore: function (t) { for (var e, n, r = this, i = [r]; r = i.pop();)
            if (t(r), e = r.children)
                for (n = e.length - 1; n >= 0; --n)
                    i.push(e[n]); return this; }, sum: function (t) { return this.eachAfter((function (e) { for (var n = +t(e.data) || 0, r = e.children, i = r && r.length; --i >= 0;)
            n += r[i].value; e.value = n; })); }, sort: function (t) { return this.eachBefore((function (e) { e.children && e.children.sort(t); })); }, path: function (t) { for (var e = this, n = function (t, e) { if (t === e)
            return t; var n = t.ancestors(), r = e.ancestors(), i = null; t = n.pop(), e = r.pop(); for (; t === e;)
            i = t, t = n.pop(), e = r.pop(); return i; }(e, t), r = [e]; e !== n;)
            e = e.parent, r.push(e); for (var i = r.length; t !== n;)
            r.splice(i, 0, t), t = t.parent; return r; }, ancestors: function () { for (var t = this, e = [t]; t = t.parent;)
            e.push(t); return e; }, descendants: function () { var t = []; return this.each((function (e) { t.push(e); })), t; }, leaves: function () { var t = []; return this.eachBefore((function (e) { e.children || t.push(e); })), t; }, links: function () { var t = this, e = []; return t.each((function (n) { n !== t && e.push({ source: n.parent, target: n }); })), e; }, copy: function () { return v(this).eachBefore(k); } };
    var S = { name: "d3-flextree", version: "2.1.2", main: "build/d3-flextree.js", module: "index", "jsnext:main": "index", author: { name: "Chris Maloney", url: "http://chrismaloney.org" }, description: "Flexible tree layout algorithm that allows for variable node sizes.", keywords: ["d3", "d3-module", "layout", "tree", "hierarchy", "d3-hierarchy", "plugin", "d3-plugin", "infovis", "visualization", "2d"], homepage: "https://github.com/klortho/d3-flextree", license: "WTFPL", repository: { type: "git", url: "https://github.com/klortho/d3-flextree.git" }, scripts: { clean: "rm -rf build demo test", "build:demo": "rollup -c --environment BUILD:demo", "build:dev": "rollup -c --environment BUILD:dev", "build:prod": "rollup -c --environment BUILD:prod", "build:test": "rollup -c --environment BUILD:test", build: "rollup -c", lint: "eslint index.js src", "test:main": "node test/bundle.js", "test:browser": "node test/browser-tests.js", test: "npm-run-all test:*", prepare: "npm-run-all clean build lint test" }, dependencies: { "d3-hierarchy": "^1.1.5" }, devDependencies: { "babel-plugin-external-helpers": "^6.22.0", "babel-preset-es2015-rollup": "^3.0.0", d3: "^4.13.0", "d3-selection-multi": "^1.0.1", eslint: "^4.19.1", jsdom: "^11.6.2", "npm-run-all": "^4.1.2", rollup: "^0.55.3", "rollup-plugin-babel": "^2.7.1", "rollup-plugin-commonjs": "^8.0.2", "rollup-plugin-copy": "^0.2.3", "rollup-plugin-json": "^2.3.0", "rollup-plugin-node-resolve": "^3.0.2", "rollup-plugin-uglify": "^3.0.0", "uglify-es": "^3.3.9" } };
    const { version: w } = S, E = Object.freeze({ children: t => t.children, nodeSize: t => t.data.size, spacing: 0 });
    function j(t) { const e = Object.assign({}, E, t); function n(t) { const n = e[t]; return "function" == typeof n ? n : () => n; } function r(t) { const e = a(function () { const t = i(), e = n("nodeSize"), r = n("spacing"); return class extends t {
        constructor(t) { super(t), Object.assign(this, { x: 0, y: 0, relX: 0, prelim: 0, shift: 0, change: 0, lExt: this, lExtRelX: 0, lThr: null, rExt: this, rExtRelX: 0, rThr: null }); }
        get size() { return e(this.data); }
        spacing(t) { return r(this.data, t.data); }
        get x() { return this.data.x; }
        set x(t) { this.data.x = t; }
        get y() { return this.data.y; }
        set y(t) { this.data.y = t; }
        update() { return C(this), X(this), this; }
    }; }(), t, (t => t.children)); return e.update(), e.data; } function i() { const t = n("nodeSize"), e = n("spacing"); return class n extends v.prototype.constructor {
        constructor(t) { super(t); }
        copy() { const t = a(this.constructor, this, (t => t.children)); return t.each((t => t.data = t.data.data)), t; }
        get size() { return t(this); }
        spacing(t) { return e(this, t); }
        get nodes() { return this.descendants(); }
        get xSize() { return this.size[0]; }
        get ySize() { return this.size[1]; }
        get top() { return this.y; }
        get bottom() { return this.y + this.ySize; }
        get left() { return this.x - this.xSize / 2; }
        get right() { return this.x + this.xSize / 2; }
        get root() { const t = this.ancestors(); return t[t.length - 1]; }
        get numChildren() { return this.hasChildren ? this.children.length : 0; }
        get hasChildren() { return !this.noChildren; }
        get noChildren() { return null === this.children; }
        get firstChild() { return this.hasChildren ? this.children[0] : null; }
        get lastChild() { return this.hasChildren ? this.children[this.numChildren - 1] : null; }
        get extents() { return (this.children || []).reduce(((t, e) => n.maxExtents(t, e.extents)), this.nodeExtents); }
        get nodeExtents() { return { top: this.top, bottom: this.bottom, left: this.left, right: this.right }; }
        static maxExtents(t, e) { return { top: Math.min(t.top, e.top), bottom: Math.max(t.bottom, e.bottom), left: Math.min(t.left, e.left), right: Math.max(t.right, e.right) }; }
    }; } function a(t, e, n) { const r = (e, i) => { const a = new t(e); Object.assign(a, { parent: i, depth: null === i ? 0 : i.depth + 1, height: 0, length: 1 }); const o = n(e) || []; return a.children = 0 === o.length ? null : o.map((t => r(t, a))), a.children && Object.assign(a, a.children.reduce(((t, e) => ({ height: Math.max(t.height, e.height + 1), length: t.length + e.length })), a)), a; }; return r(e, null); } return Object.assign(r, { nodeSize(t) { return arguments.length ? (e.nodeSize = t, r) : e.nodeSize; }, spacing(t) { return arguments.length ? (e.spacing = t, r) : e.spacing; }, children(t) { return arguments.length ? (e.children = t, r) : e.children; }, hierarchy(t, n) { const r = void 0 === n ? e.children : n; return a(i(), t, r); }, dump(t) { const e = n("nodeSize"), r = t => n => { const i = t + "  ", a = t + "    ", { x: o, y: s } = n, l = e(n), c = n.children || [], h = 0 === c.length ? " " : `,${i}children: [${a}${c.map(r(a)).join(a)}${i}],${t}`; return `{ size: [${l.join(", ")}],${i}x: ${o}, y: ${s}${h}},`; }; return r("\n")(t); } }), r; }
    j.version = w;
    const C = (t, e = 0) => (t.y = e, (t.children || []).reduce(((e, n) => { const [r, i] = e; C(n, t.y + t.ySize); const a = (0 === r ? n.lExt : n.rExt).bottom; 0 !== r && I(t, r, i); return [r + 1, D(a, r, i)]; }), [0, null]), O(t), T(t), t), X = (t, e, n) => { void 0 === e && (e = -t.relX - t.prelim, n = 0); const r = e + t.relX; return t.relX = r + t.prelim - n, t.prelim = 0, t.x = n + t.relX, (t.children || []).forEach((e => X(e, r, t.x))), t; }, O = t => { (t.children || []).reduce(((t, e) => { const [n, r] = t, i = n + e.shift, a = r + i + e.change; return e.relX += a, [i, a]; }), [0, 0]); }, I = (t, e, n) => { const r = t.children[e - 1], i = t.children[e]; let a = r, o = r.relX, s = i, l = i.relX, c = !0; for (; a && s;) {
        a.bottom > n.lowY && (n = n.next);
        const r = o + a.prelim - (l + s.prelim) + a.xSize / 2 + s.xSize / 2 + a.spacing(s);
        (r > 0 || r < 0 && c) && (l += r, M(i, r), A(t, e, n.index, r)), c = !1;
        const h = a.bottom, d = s.bottom;
        h <= d && (a = R(a), a && (o += a.relX)), h >= d && (s = $(s), s && (l += s.relX));
    } !a && s ? B(t, e, s, l) : a && !s && N(t, e, a, o); }, M = (t, e) => { t.relX += e, t.lExtRelX += e, t.rExtRelX += e; }, A = (t, e, n, r) => { const i = t.children[e], a = e - n; if (a > 1) {
        const e = r / a;
        t.children[n + 1].shift += e, i.shift -= e, i.change -= r - e;
    } }, $ = t => t.hasChildren ? t.firstChild : t.lThr, R = t => t.hasChildren ? t.lastChild : t.rThr, B = (t, e, n, r) => { const i = t.firstChild, a = i.lExt, o = t.children[e]; a.lThr = n; const s = r - n.relX - i.lExtRelX; a.relX += s, a.prelim -= s, i.lExt = o.lExt, i.lExtRelX = o.lExtRelX; }, N = (t, e, n, r) => { const i = t.children[e], a = i.rExt, o = t.children[e - 1]; a.rThr = n; const s = r - n.relX - i.rExtRelX; a.relX += s, a.prelim -= s, i.rExt = o.rExt, i.rExtRelX = o.rExtRelX; }, T = t => { if (t.hasChildren) {
        const e = t.firstChild, n = t.lastChild, r = (e.prelim + e.relX - e.xSize / 2 + n.relX + n.prelim + n.xSize / 2) / 2;
        Object.assign(t, { prelim: r, lExt: e.lExt, lExtRelX: e.lExtRelX, rExt: n.rExt, rExtRelX: n.rExtRelX });
    } }, D = (t, e, n) => { for (; null !== n && t >= n.lowY;)
        n = n.next; return { lowY: t, index: e, next: n }; };
    var H = "http://www.w3.org/1999/xlink", L = { show: H, actuate: H, href: H };
    function P(t, e) { var n; if ("string" == typeof t)
        n = 1;
    else {
        if ("function" != typeof t)
            throw new Error("Invalid VNode type");
        n = 2;
    } return { vtype: n, type: t, props: e }; }
    function F(t) { return t.children; }
    var Y = { isSvg: !1 };
    function W(t, e) { if (1 === e.type)
        null != e.node && t.append(e.node);
    else {
        if (4 !== e.type)
            throw new Error("Unkown ref type " + JSON.stringify(e));
        e.children.forEach((function (e) { W(t, e); }));
    } }
    var _ = { className: "class", labelFor: "for" };
    function U(t, e, n, r) { if (e = _[e] || e, !0 === n)
        t.setAttribute(e, "");
    else if (!1 === n)
        t.removeAttribute(e);
    else {
        var i = r ? L[e] : void 0;
        void 0 !== i ? t.setAttributeNS(i, e, n) : t.setAttribute(e, n);
    } }
    function V(t, e) { if (void 0 === e && (e = Y), null == t || "boolean" == typeof t)
        return { type: 1, node: null }; if (t instanceof Node)
        return { type: 1, node: t }; if (2 === (null == (o = t) ? void 0 : o.vtype)) {
        var n = t, r = n.type, i = n.props;
        if (r === F) {
            var a = document.createDocumentFragment();
            if (i.children)
                W(a, V(i.children, e));
            return { type: 1, node: a };
        }
        return V(r(i), e);
    } var o; if (function (t) { return "string" == typeof t || "number" == typeof t; }(t))
        return { type: 1, node: document.createTextNode("" + t) }; if (function (t) { return 1 === (null == t ? void 0 : t.vtype); }(t)) {
        var s, l, c = t, h = c.type, d = c.props;
        if (e.isSvg || "svg" !== h || (e = Object.assign({}, e, { isSvg: !0 })), function (t, e, n) { for (var r in e)
            "key" !== r && "children" !== r && "ref" !== r && ("dangerouslySetInnerHTML" === r ? t.innerHTML = e[r].__html : "innerHTML" === r || "textContent" === r || "innerText" === r ? t[r] = e[r] : r.startsWith("on") ? t[r.toLowerCase()] = e[r] : U(t, r, e[r], n.isSvg)); }(s = e.isSvg ? document.createElementNS("http://www.w3.org/2000/svg", h) : document.createElement(h), d, e), d.children) {
            var u = e;
            e.isSvg && "foreignObject" === h && (u = Object.assign({}, u, { isSvg: !1 })), l = V(d.children, u);
        }
        null != l && W(s, l);
        var p = d.ref;
        return "function" == typeof p && p(s), { type: 1, node: s };
    } if (Array.isArray(t))
        return { type: 4, children: t.map((function (t) { return V(t, e); })) }; throw new Error("mount: Invalid Vnode!"); }
    function Z(t) { for (var e = [], n = 0; n < t.length; n += 1) {
        var r = t[n];
        Array.isArray(r) ? e = e.concat(Z(r)) : null != r && e.push(r);
    } return e; }
    function G(t) { return 1 === t.type ? t.node : t.children.map(G); }
    function J(t) { return Array.isArray(t) ? Z(t.map(J)) : G(V(t)); }
    var K = ".markmap{font:300 16px/20px sans-serif}.markmap-link{fill:none}.markmap-node>circle{cursor:pointer}.markmap-foreign{display:inline-block}.markmap-foreign a{color:#0097e6}.markmap-foreign a:hover{color:#00a8ff}.markmap-foreign code{background-color:#f0f0f0;border-radius:2px;color:#555;font-size:calc(1em - 2px)}.markmap-foreign :not(pre)>code{padding:.2em .4em}.markmap-foreign del{text-decoration:line-through}.markmap-foreign em{font-style:italic}.markmap-foreign strong{font-weight:bolder}.markmap-foreign mark{background:#ffeaa7}.markmap-foreign pre,.markmap-foreign pre[class*=language-]{margin:0;padding:.2em .4em}", q = ".markmap-container{height:0;left:-100px;overflow:hidden;position:absolute;top:-100px;width:0}.markmap-container>.markmap-foreign{display:inline-block}.markmap-container>.markmap-foreign>div:last-child{white-space:nowrap}";
    function Q(t) { const e = t.data; return Math.max(4 - 2 * e.depth, 1.5); }
    function tt(t, e) { return t[r.minIndex(t, e)]; }
    function et(t) { t.stopPropagation(); }
    const nt = new i, rt = r.scaleOrdinal(r.schemeCategory10), it = "undefined" != typeof navigator && navigator.userAgent.includes("Macintosh");
    class at {
        constructor(t, e) { this.revokers = [], ["handleZoom", "handleClick", "handlePan"].forEach((t => { this[t] = this[t].bind(this); })), this.viewHooks = { transformHtml: new i }, this.svg = t.datum ? t : r.select(t), this.styleNode = this.svg.append("style"), this.zoom = r.zoom().filter((t => this.options.scrollForPan && "wheel" === t.type ? t.ctrlKey && !t.button : !(t.ctrlKey && "wheel" !== t.type || t.button))).on("zoom", this.handleZoom), this.setOptions(e), this.state = { id: this.options.id || this.svg.attr("id") || (s += 1, `mm-${o}-${s}`) }, this.g = this.svg.append("g"), this.updateStyle(), this.revokers.push(nt.tap((() => { this.setData(); }))); }
        getStyleContent() { const { style: t } = this.options, { id: e } = this.state, n = "function" == typeof t ? t(e) : ""; return [this.options.embedGlobalCSS && K, n].filter(Boolean).join("\n"); }
        updateStyle() { this.svg.attr("class", function (t, ...e) { const n = (t || "").split(" ").filter(Boolean); return e.forEach((t => { t && n.indexOf(t) < 0 && n.push(t); })), n.join(" "); }(this.svg.attr("class"), "markmap", this.state.id)); const t = this.getStyleContent(); this.styleNode.text(t); }
        handleZoom(t) { const { transform: e } = t; this.g.attr("transform", e); }
        handlePan(t) { t.preventDefault(); const e = r.zoomTransform(this.svg.node()), n = e.translate(-t.deltaX / e.k, -t.deltaY / e.k); this.svg.call(this.zoom.transform, n); }
        handleClick(t, e) { var n; const { data: r } = e; r.payload = g({}, r.payload, { fold: null != (n = r.payload) && n.fold ? 0 : 1 }), this.renderData(e.data); }
        initializeData(t) { let e = 0; const { color: n, nodeMinHeight: r, maxWidth: i, initialExpandLevel: a } = this.options, { id: o } = this.state, s = J(P("div", { className: `markmap-container markmap ${o}-g` })), l = J(P("style", { children: [this.getStyleContent(), q].join("\n") })); document.body.append(s, l); const d = i ? `max-width: ${i}px` : ""; let u = 0; c(t, ((t, r, i) => { var o, l, c; t.children = null == (o = t.children) ? void 0 : o.map((t => g({}, t))), e += 1; const h = J(P("div", { className: "markmap-foreign", style: d, children: P("div", { dangerouslySetInnerHTML: { __html: t.content } }) })); s.append(h), t.state = g({}, t.state, { id: e, el: h.firstChild }), t.state.path = [null == i || null == (l = i.state) ? void 0 : l.path, t.state.id].filter(Boolean).join("."), n(t); const p = 2 === (null == (c = t.payload) ? void 0 : c.fold); p ? u += 1 : (u || a >= 0 && t.depth >= a) && (t.payload = g({}, t.payload, { fold: 1 })), r(), p && (u -= 1); })); const p = h(s.childNodes).map((t => t.firstChild)); this.viewHooks.transformHtml.call(this, p), p.forEach((t => { t.parentNode.append(t.cloneNode(!0)); })), c(t, ((t, e, n) => { var i; const a = t.state.el.getBoundingClientRect(); t.content = t.state.el.innerHTML, t.state.size = [Math.ceil(a.width) + 1, Math.max(Math.ceil(a.height), r)], t.state.key = [null == n || null == (i = n.state) ? void 0 : i.id, t.state.id].filter(Boolean).join(".") + t.content, e(); })), s.remove(), l.remove(); }
        setOptions(t) { this.options = g({}, at.defaultOptions, t), this.options.zoom ? this.svg.call(this.zoom) : this.svg.on(".zoom", null), this.svg.on("wheel", this.options.pan ? this.handlePan : null); }
        setData(t, e) { t && (this.state.data = t), e && this.setOptions(e), this.initializeData(this.state.data), this.renderData(); }
        renderData(t) { var e, n; if (!this.state.data)
            return; const { spacingHorizontal: i, paddingX: a, spacingVertical: o, autoFit: s, color: l } = this.options, h = j().children((t => { var e; return !(null != (e = t.payload) && e.fold) && t.children; })).nodeSize((t => { const [e, n] = t.data.state.size; return [n, e + (e ? 2 * a : 0) + i]; })).spacing(((t, e) => t.parent === e.parent ? o : 2 * o)), u = h.hierarchy(this.state.data); h(u), function (t, e) { c(t, ((t, n) => { t.ySizeInner = t.ySize - e, t.y += e, n(); }), "children"); }(u, i); const p = u.descendants().reverse(), f = u.links(), m = r.linkHorizontal(), g = r.min(p, (t => t.x - t.xSize / 2)), y = r.max(p, (t => t.x + t.xSize / 2)), v = r.min(p, (t => t.y)), x = r.max(p, (t => t.y + t.ySizeInner)); Object.assign(this.state, { minX: g, maxX: y, minY: v, maxY: x }), s && this.fit(); const k = t && p.find((e => e.data === t)) || u, z = null != (e = k.data.state.x0) ? e : k.x, b = null != (n = k.data.state.y0) ? n : k.y, S = this.g.selectAll(d("g")).data(p, (t => t.data.state.key)), w = S.enter().append("g").attr("data-depth", (t => t.data.depth)).attr("data-path", (t => t.data.state.path)).attr("transform", (t => `translate(${b + k.ySizeInner - t.ySizeInner},${z + k.xSize / 2 - t.xSize})`)), E = this.transition(S.exit()); E.select("line").attr("x1", (t => t.ySizeInner)).attr("x2", (t => t.ySizeInner)), E.select("foreignObject").style("opacity", 0), E.attr("transform", (t => `translate(${k.y + k.ySizeInner - t.ySizeInner},${k.x + k.xSize / 2 - t.xSize})`)).remove(); const C = S.merge(w).attr("class", (t => { var e; return ["markmap-node", (null == (e = t.data.payload) ? void 0 : e.fold) && "markmap-fold"].filter(Boolean).join(" "); })); this.transition(C).attr("transform", (t => `translate(${t.y},${t.x - t.xSize / 2})`)); const X = C.selectAll(d("line")).data((t => [t]), (t => t.data.state.key)).join((t => t.append("line").attr("x1", (t => t.ySizeInner)).attr("x2", (t => t.ySizeInner))), (t => t), (t => t.remove())); this.transition(X).attr("x1", -1).attr("x2", (t => t.ySizeInner + 2)).attr("y1", (t => t.xSize)).attr("y2", (t => t.xSize)).attr("stroke", (t => l(t.data))).attr("stroke-width", Q); const O = C.selectAll(d("circle")).data((t => t.data.children ? [t] : []), (t => t.data.state.key)).join((t => t.append("circle").attr("stroke-width", "1.5").attr("cx", (t => t.ySizeInner)).attr("cy", (t => t.xSize)).attr("r", 0).on("click", ((t, e) => this.handleClick(t, e)))), (t => t), (t => t.remove())); this.transition(O).attr("r", 6).attr("cx", (t => t.ySizeInner)).attr("cy", (t => t.xSize)).attr("stroke", (t => l(t.data))).attr("fill", (t => { var e; return null != (e = t.data.payload) && e.fold && t.data.children ? l(t.data) : "#fff"; })); const I = C.selectAll(d("foreignObject")).data((t => [t]), (t => t.data.state.key)).join((t => { const e = t.append("foreignObject").attr("class", "markmap-foreign").attr("x", a).attr("y", 0).style("opacity", 0).on("mousedown", et).on("dblclick", et); return e.append("xhtml:div").select((function (t) { const e = t.data.state.el.cloneNode(!0); return this.replaceWith(e), e; })).attr("xmlns", "http://www.w3.org/1999/xhtml"), e; }), (t => t), (t => t.remove())).attr("width", (t => Math.max(0, t.ySizeInner - 2 * a))).attr("height", (t => t.xSize)); this.transition(I).style("opacity", 1); const M = this.g.selectAll(d("path")).data(f, (t => t.target.data.state.key)).join((t => { const e = [b + k.ySizeInner, z + k.xSize / 2]; return t.insert("path", "g").attr("class", "markmap-link").attr("data-depth", (t => t.target.data.depth)).attr("data-path", (t => t.target.data.state.path)).attr("d", m({ source: e, target: e })); }), (t => t), (t => { const e = [k.y + k.ySizeInner, k.x + k.xSize / 2]; return this.transition(t).attr("d", m({ source: e, target: e })).remove(); })); this.transition(M).attr("stroke", (t => l(t.target.data))).attr("stroke-width", (t => Q(t.target))).attr("d", (t => { const e = [t.source.y + t.source.ySizeInner, t.source.x + t.source.xSize / 2], n = [t.target.y, t.target.x + t.target.xSize / 2]; return m({ source: e, target: n }); })), p.forEach((t => { t.data.state.x0 = t.x, t.data.state.y0 = t.y; })); }
        transition(t) { const { duration: e } = this.options; return t.transition().duration(e); }
        async fit() { const t = this.svg.node(), { width: e, height: n } = t.getBoundingClientRect(), { fitRatio: i } = this.options, { minX: a, maxX: o, minY: s, maxY: c } = this.state, h = c - s, d = o - a, u = Math.min(e / h * i, n / d * i, 2), p = r.zoomIdentity.translate((e - h * u) / 2 - s * u, (n - d * u) / 2 - a * u).scale(u); return this.transition(this.svg).call(this.zoom.transform, p).end().catch(l); }
        async ensureView(t, e) { let n, i; if (this.g.selectAll(d("g")).each((function (e) { e.data === t && (n = this, i = e); })), !n || !i)
            return; const a = this.svg.node(), o = a.getBoundingClientRect(), s = r.zoomTransform(a), [c, h] = [i.y, i.y + i.ySizeInner + 2].map((t => t * s.k + s.x)), [u, p] = [i.x - i.xSize / 2, i.x + i.xSize / 2].map((t => t * s.k + s.y)), f = g({ left: 0, right: 0, top: 0, bottom: 0 }, e), m = [f.left - c, o.width - f.right - h], y = [f.top - u, o.height - f.bottom - p], v = m[0] * m[1] > 0 ? tt(m, Math.abs) / s.k : 0, x = y[0] * y[1] > 0 ? tt(y, Math.abs) / s.k : 0; if (v || x) {
            const t = s.translate(v, x);
            return this.transition(this.svg).call(this.zoom.transform, t).end().catch(l);
        } }
        async rescale(t) { const e = this.svg.node(), { width: n, height: i } = e.getBoundingClientRect(), a = n / 2, o = i / 2, s = r.zoomTransform(e), c = s.translate((a - s.x) * (1 - t) / s.k, (o - s.y) * (1 - t) / s.k).scale(t); return this.transition(this.svg).call(this.zoom.transform, c).end().catch(l); }
        destroy() { this.svg.on(".zoom", null), this.svg.html(null), this.revokers.forEach((t => { t(); })); }
        static create(t, e, n) { const r = new at(t, e); return n && (r.setData(n), r.fit()), r; }
    }
    at.defaultOptions = { autoFit: !1, color: t => rt(`${t.state.path}`), duration: 500, embedGlobalCSS: !0, fitRatio: .95, maxWidth: 0, nodeMinHeight: 16, paddingX: 8, scrollForPan: it, spacingHorizontal: 80, spacingVertical: 5, initialExpandLevel: -1, zoom: !0, pan: !0 }, t.Markmap = at, t.defaultColorFn = rt, t.deriveOptions = function (t) { const e = {}; t || (t = {}); const { color: n, colorFreezeLevel: i } = t; if (1 === (null == n ? void 0 : n.length)) {
        const t = n[0];
        e.color = () => t;
    }
    else if (null != n && n.length) {
        const t = r.scaleOrdinal(n);
        e.color = e => t(`${e.state.path}`);
    } if (i) {
        const t = e.color || at.defaultOptions.color;
        e.color = e => (e = g({}, e, { state: g({}, e.state, { path: e.state.path.split(".").slice(0, i).join(".") }) }), t(e));
    } return ["duration", "maxWidth", "initialExpandLevel"].forEach((n => { const r = t[n]; "number" == typeof r && (e[n] = r); })), ["zoom", "pan"].forEach((n => { const r = t[n]; null != r && (e[n] = !!r); })), e; }, t.globalCSS = ".markmap{font:300 16px/20px sans-serif}.markmap-link{fill:none}.markmap-node>circle{cursor:pointer}.markmap-foreign{display:inline-block}.markmap-foreign a{color:#0097e6}.markmap-foreign a:hover{color:#00a8ff}.markmap-foreign code{background-color:#f0f0f0;border-radius:2px;color:#555;font-size:calc(1em - 2px)}.markmap-foreign :not(pre)>code{padding:.2em .4em}.markmap-foreign del{text-decoration:line-through}.markmap-foreign em{font-style:italic}.markmap-foreign strong{font-weight:bolder}.markmap-foreign mark{background:#ffeaa7}.markmap-foreign pre,.markmap-foreign pre[class*=language-]{margin:0;padding:.2em .4em}", t.loadCSS = function (t) { for (const e of t)
        m(e); }, t.loadJS = async function (t, e) { const n = t.filter((t => { var e; return "script" === t.type && (null == (e = t.data) ? void 0 : e.src); })); n.length > 1 && n.forEach((t => p(t.data.src))), e = a({ getMarkmap: () => window.markmap }, e); for (const n of t)
        await f(n, e); }, t.refreshHook = nt;
}(this.markmap = this.markmap || {}, d3);
