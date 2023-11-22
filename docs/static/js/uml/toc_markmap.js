"use strict";
class TocItem {
    content;
    depth;
    id;
    children;
    parent;
    constructor(content, depth, parent = undefined) {
        this.content = content;
        this.depth = depth;
        this.id = undefined;
        this.parent = parent;
        this.children = [];
    }
}
function parseHeading(headingSet) {
    const tocData = [];
    let curLevel = 0;
    let preTocItem = undefined;
    headingSet.forEach(heading => {
        const hLevel = Number(heading.outerHTML.match(/<h([\d]).*>/)[1]);
        const titleText = heading.textContent;
        if (heading.id === "") {
            heading.id = titleText.replace(/ /g, "-").toLowerCase();
        }
        const titleHTML = `<a href="#${heading.id}">${titleText}</a>`;
        switch (hLevel >= curLevel) {
            case true:
                if (preTocItem === undefined) {
                    preTocItem = new TocItem(titleHTML, hLevel);
                    tocData.push(preTocItem);
                }
                else {
                    const curTocItem = new TocItem(titleHTML, hLevel);
                    const parent = curTocItem.depth > preTocItem.depth ? preTocItem : preTocItem.parent;
                    curTocItem.parent = parent;
                    if (parent !== undefined) {
                        parent.children.push(curTocItem);
                    }
                    preTocItem = curTocItem;
                }
                break;
            case false:
                const curTocItem = new TocItem(titleHTML, hLevel);
                while (1) {
                    if (preTocItem.depth < curTocItem.depth) {
                        preTocItem.children.push(curTocItem);
                        curTocItem.parent = preTocItem;
                        preTocItem = curTocItem;
                        break;
                    }
                    preTocItem = preTocItem.parent;
                    if (preTocItem === undefined) {
                        tocData.push(curTocItem);
                        preTocItem = curTocItem;
                        break;
                    }
                }
                break;
        }
        preTocItem.id = heading.id;
        curLevel = hLevel;
    });
    return tocData;
}
class Toc {
    static createMarkmap(svgID, data) {
        window.WebFontConfig = {
            custom: { families: ["KaTeX_AMS", "KaTeX_Caligraphic:n4,n7", "KaTeX_Fraktur:n4,n7", "KaTeX_Main:n4,n7,i4,i7", "KaTeX_Math:i4,i7", "KaTeX_Script", "KaTeX_SansSerif:n4,n7,i4", "KaTeX_Size1", "KaTeX_Size2", "KaTeX_Size3", "KaTeX_Size4", "KaTeX_Typewriter"] },
            active: () => {
                window.markmap.refreshHook.call();
            }
        };
        window.mm = window.markmap.Markmap.create("svg#" + svgID, window.markmap.deriveOptions({ "colorFreezeLevel": 4 }), data);
    }
}
const initSVGHoverAttr = (svg) => {
    const clientRectBody = document.body.getBoundingClientRect();
    const clientRectSVG = svg.getBoundingClientRect();
    const new_x = (clientRectBody.width - clientRectSVG.width) / 2;
    const left = -(clientRectSVG.x - new_x) * 0.8;
    const sheetName = "styles";
    setStyleRule(sheetName, "#markmap-toc:hover", "left:" + left + "px");
};
const setStyleRule = (sheetName, selector, rule) => {
    let linkElem = document.querySelector('link[href*=' + sheetName + ']');
    if (linkElem) {
        const stylesheet = linkElem.sheet;
        stylesheet.insertRule(selector + '{ ' + rule + '}', stylesheet.cssRules.length);
    }
};
function hideBtnCopyPre() {
    let copy_btn_list = document.getElementsByClassName("btn-copy-pre");
    for (let btn of copy_btn_list) {
        btn.style.display = "none";
    }
}
function showBtnCopyPre() {
    let copy_btn_list = document.getElementsByClassName("btn-copy-pre");
    for (let btn of copy_btn_list) {
        btn.style.display = "";
    }
}
(() => {
    window.addEventListener('DOMContentLoaded', () => {
        const headingSet = [...document.querySelectorAll("h1, h2, h3, h4, h5, h6")];
        const tocItems = parseHeading(headingSet);
        if (tocItems.length === 0) {
            return;
        }
        const idName = 'markmap-toc';
        const frag = document.createRange().createContextualFragment(`
        <div class="markmapWrapper navbar-right">
        <svg id="${idName}" class="markmap" />
        </div>
        `);
        const svgElem = frag.querySelector(`svg`);
        const tocContainer = frag.querySelector('div');
        document.body.append(frag);
        const dictData = tocItems[0];
        Toc.createMarkmap(idName, dictData);
        initSVGHoverAttr(svgElem);
    });
})();
