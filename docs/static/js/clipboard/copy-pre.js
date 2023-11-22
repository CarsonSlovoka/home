"use strict";
function addCopyButtons(clipboard) {
    document.querySelectorAll('pre > code').forEach((codeElem) => {
        const frag = document.createRange().createContextualFragment(`
<div class="btn-copy-pre">
<span></span>
<img src="/static/img/svg/fas/fa-copy.svg" title="Copy"/>
</div>
`);
        const divElem = frag.querySelector("div");
        const btn = frag.querySelector("img");
        const spanElem = frag.querySelector(`span`);
        btn.addEventListener('click', () => {
            clipboard.writeText(codeElem.innerText).then(() => {
                btn.blur();
                divElem.style.width = "fit-content";
                spanElem.innerText = 'Copied!';
                btn.style.display = "none";
                setTimeout(() => {
                    spanElem.innerText = "";
                    btn.style.display = "";
                    divElem.style.width = "";
                }, 1000);
            }, (err) => {
                btn.innerText = '[Error]' + err;
            });
        });
        const preElem = codeElem.parentNode;
        const preParentNode = preElem.parentNode;
        preParentNode.insertBefore(frag, preElem);
    });
}
window.addEventListener("load", () => {
    if (navigator && navigator.clipboard) {
        addCopyButtons(navigator.clipboard);
    }
});
