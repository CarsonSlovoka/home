// https://www.dannyguo.com/blog/how-to-add-copy-to-clipboard-buttons-to-code-blocks-in-hugo/

function addCopyButtons(clipboard: Clipboard) {
    (document.querySelectorAll('pre > code') as any).forEach((codeElem: HTMLSpanElement) => {

        const frag = document.createRange().createContextualFragment(`
<div class="btn-copy-pre">
<span></span>
<img src="/static/img/svg/fas/fa-copy.svg" title="Copy"/>
</div>
`)
        const divElem = frag.querySelector("div") as HTMLDivElement
        const btn = frag.querySelector("img") as HTMLImageElement
        const spanElem = frag.querySelector(`span`) as HTMLSpanElement

        btn.addEventListener('click', () => {
            clipboard.writeText(codeElem.innerText).then(() => {
                /* Chrome doesn't seem to blur automatically,
                   leaving the button in a focused state. */
                btn.blur()

                divElem.style.width = "fit-content" // 改變寬度，讓其寬度符合Copied的文字寬
                spanElem.innerText = 'Copied!'
                btn.style.display = "none" // 讓按鈕先看不見


                setTimeout(() => {
                    spanElem.innerText = "" // 注意，如果您改變了innerText，該Elem「內」的所有其他子元素會消失
                    // spanElem.appendChild(mySubElem)  // 如果要讓子元素再出來，可能要考慮重新append
                    btn.style.display = ""
                    divElem.style.width = ""
                }, 1000)
            }, (err) => {
                btn.innerText = '[Error]' + err
            })
        })

        const preElem = codeElem.parentNode
        const preParentNode = preElem!.parentNode as HTMLElement
        preParentNode.insertBefore(frag, preElem) // 這會讓元素放在preElem等前面(和preElem同層)
    })
}


window.addEventListener("load", () => {
    if (navigator && navigator.clipboard) {
        // https://developer.mozilla.org/en-US/docs/Web/API/Navigator/clipboard
        addCopyButtons(navigator.clipboard)
    }
})
