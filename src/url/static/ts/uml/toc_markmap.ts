// https://stackoverflow.com/a/69260236/9935654

type markMapData = {
    // 以下的type, payload，都是為了還原資料所用，所以可以省略
    // type: string // heading表示hx, node-list表示為項目符號
    // payload: {lines: [number]} // 可以知道是第幾列
    depth: number
    content: string
    children: markMapData[] | undefined | null
}

class TocItem {
    public id: string | undefined
    public children: TocItem[] | undefined
    public parent: TocItem | undefined
    constructor(public content: any, public depth: number, parent = undefined) {
        this.id = undefined
        this.parent = parent
        this.children = []
    }
}

// 爬取自定義的{h1,..., hn}系列轉換成TocItem[]
function parseHeading(headingSet: HTMLHeadingElement[]|Element[]): TocItem[] {
    const tocData = [] as TocItem[]
    let curLevel = 0
    let preTocItem: TocItem|undefined = undefined

    headingSet.forEach(heading => {
        const hLevel = Number(heading.outerHTML.match(/<h([\d]).*>/)![1])
        const titleText = heading.textContent
        if (heading.id === "") {
            heading.id = titleText!.replace(/ /g, "-").toLowerCase()
        }
        const titleHTML = `<a href="#${heading.id}">${titleText}</a>` // 我們用的是markmap，它的內容可以直接放<a>

        switch (hLevel >= curLevel) {
            case true:
                if (preTocItem === undefined) {
                    preTocItem = new TocItem(titleHTML, hLevel)
                    tocData.push(preTocItem)
                } else {
                    const curTocItem = new TocItem(titleHTML, hLevel)
                    const parent = curTocItem.depth > preTocItem.depth ? preTocItem : preTocItem.parent
                    curTocItem.parent = parent
                    if (parent !== undefined) { // 如果沒有h1，然後從h2開始就有可能會發生
                        parent.children!.push(curTocItem)
                    }
                    preTocItem = curTocItem
                }
                break
            case false:
                // 從preTocItem中找到適合的父項
                // We need to find the appropriate parent node from the preTocItem
                const curTocItem = new TocItem(titleHTML, hLevel)
                while (1) {
                    if (preTocItem!.depth < curTocItem.depth) {
                        // 父項已被找到
                        preTocItem!.children!.push(curTocItem)
                        curTocItem.parent = preTocItem
                        preTocItem = curTocItem
                        break
                    }

                    // 表示前項不為本項目的父項，所以再找"前項.父項"
                    preTocItem = preTocItem!.parent

                    if (preTocItem === undefined) {
                        // 如果到最後這個父項還是找不到，就直接將此節點添加
                        tocData.push(curTocItem)
                        preTocItem = curTocItem
                        break
                    }
                }
                break
        }
        preTocItem!.id = heading.id
        curLevel = hLevel
    })

    return tocData
}

class Toc {
    static createMarkmap(svgID: string, data: markMapData) {
        (window as any).WebFontConfig = { // 數學符號Katex可以正常顯示用
            custom: {families: ["KaTeX_AMS", "KaTeX_Caligraphic:n4,n7", "KaTeX_Fraktur:n4,n7", "KaTeX_Main:n4,n7,i4,i7", "KaTeX_Math:i4,i7", "KaTeX_Script", "KaTeX_SansSerif:n4,n7,i4", "KaTeX_Size1", "KaTeX_Size2", "KaTeX_Size3", "KaTeX_Size4", "KaTeX_Typewriter"]},
            active: () => {
                (window as any).markmap.refreshHook.call()
            }
        };
        (window as any).mm = (window as any).markmap.Markmap.create(
            "svg#" + svgID,
            (window as any).markmap.deriveOptions(
                {"colorFreezeLevel": 4}, // 分支條的顏色數量，用太多會太花
            ),
            data
        )
    }
}

// 調整svg出現的位置，讓其出現在畫面中間
const initSVGHoverAttr = (svg: SVGElement) => {
    const clientRectBody = document.body.getBoundingClientRect()
    const clientRectSVG = svg.getBoundingClientRect();
    // const clientHeight = document.documentElement.clientHeight
    const new_x = (clientRectBody.width - clientRectSVG.width) / 2  // 計算出兩邊應該留白多少
    // const new_y = (clientHeight - clientRectSVG.height) / 2
    const left = -(clientRectSVG.x - new_x) *0.8  // 從目前的位置移置到應留白的起始位置 (因為我們已知道svg是在右邊，要往左移動所以用-號)
    // const top = -(clientHeight - new_y)
    /*
    svg.style["background-color"] = "rgb(0, 0, 0)"
    svg.style.transform = "scale(5)"
    svg.style.position = "relative"
    */
    //document.styleSheets
    const sheetName = "styles"  // .css
    setStyleRule(sheetName, "#markmap-toc:hover", "left:" + left + "px") // 直接對此css做異動
    // setStyleRule(sheetName, "#markmap-toc:hover", "top:" + top + "px")
    // setStyleRule(sheetName, "#markmap-toc:hover", "background-color: rgb(255, 0, 0)")
}

const setStyleRule = (sheetName: string, selector: string, rule: string) => {
    let linkElem = document.querySelector('link[href*=' + sheetName + ']') as HTMLLinkElement

    if (linkElem) {
        const stylesheet = linkElem.sheet as CSSStyleSheet
        stylesheet.insertRule(selector + '{ ' + rule + '}', stylesheet.cssRules.length)
    }
}

// ~~因為code-block的按鈕會與svg的呈現相衝，所以一者出現另一者就隱藏~~ 使用z-index取代
function hideBtnCopyPre() {
    let copy_btn_list = document.getElementsByClassName("btn-copy-pre") as HTMLCollection
    for (let btn of copy_btn_list as any) {
        btn.style.display = "none"
    }
}

function showBtnCopyPre() {
    let copy_btn_list = document.getElementsByClassName("btn-copy-pre")
    for (let btn of copy_btn_list as any) {
        btn.style.display = ""
    }
}

(
    () => {
        // 測試資料
        /*
        const dictDataOld: markMapData = {
            // "type": "heading", // 也不重要，可能是回原完本時會用到，表示這個# hx
            "depth": 0,
            // "payload":{"lines":[1,2]}, // 推測它可以透過這個在還原成原本的文本，但在svg中，這個可以忽略
            "content": "markmap",
            "children": [
                {
                    "depth": 1,
                    "content": "Links",
                    "children": [
                        {
                            "depth": 2,
                            "content": "<a href=\"https://markmap.js.org/\">https://markmap.js.org/</a>",
                            "children": null
                        }, {
                            "depth": 2,
                            "content": "<a href=\"https://github.com/gera2ld/markmap\">GitHub</a>",
                            "children": null
                        }
                    ]
                },
                {
                    // "type": "heading",
                    "depth": 1,
                    "content": "Related Projects",
                    "children": [
                        {
                            "depth": 2,
                            "content": "<a href=\"https://github.com/gera2ld/coc-markmap\">coc-markmap</a>",
                            "children": null
                        }, {
                            "depth": 2,
                            "content": "<a href=\"https://github.com/gera2ld/gatsby-remark-markmap\">gatsby-remark-markmap</a>",
                            "children": null
                        }
                    ]
                },
                {
                    "depth": 1,
                    "content": "Features",
                    "children": [
                        {
                            "depth": 2,
                            "content": "links",
                            "children": null
                        }, {

                            "depth": 2,
                            "content": "<strong>strong</strong> <del>del</del> <em>italic</em> <mark>highlight</mark>",
                            "children": null
                        }, {

                            "depth": 2,
                            "content": "multiline<br>\ntext",
                            "children": null
                        }, {
                            "depth": 2,
                            "content": "<code>inline code</code>",
                            "children": null
                        }, {
                            "depth": 2,
                            "content": "<pre class=\"language-js\"><code class=\"language-js\">console<span class=\"token punctuation\">.</span><span class=\"token function\">log</span><span class=\"token punctuation\">(</span><span class=\"token string\">'code block'</span><span class=\"token punctuation\">)</span><span class=\"token punctuation\">;</span>\n</code></pre>\n",
                            "children": null
                        }, {
                            "depth": 2,
                            "content": "Katex",
                            "children": [{
                                "depth": 3,
                                "content": "<span class=\"katex\"><span class=\"katex-mathml\"><math xmlns=\"http://www.w3.org/1998/Math/MathML\"><semantics><mrow><mi>x</mi><mo>=</mo><mfrac><mrow><mo>−</mo><mi>b</mi><mo>±</mo><msqrt><mrow><msup><mi>b</mi><mn>2</mn></msup><mo>−</mo><mn>4</mn><mi>a</mi><mi>c</mi></mrow></msqrt></mrow><mrow><mn>2</mn><mi>a</mi></mrow></mfrac></mrow><annotation encoding=\"application/x-tex\">x = {-b \\pm \\sqrt{b^2-4ac} \\over 2a}</annotation></semantics></math></span><span class=\"katex-html\" aria-hidden=\"true\"><span class=\"base\"><span class=\"strut\" style=\"height:0.4306em;\"></span><span class=\"mord mathnormal\">x</span><span class=\"mspace\" style=\"margin-right:0.2778em;\"></span><span class=\"mrel\">=</span><span class=\"mspace\" style=\"margin-right:0.2778em;\"></span></span><span class=\"base\"><span class=\"strut\" style=\"height:1.3845em;vertical-align:-0.345em;\"></span><span class=\"mord\"><span class=\"mord\"><span class=\"mopen nulldelimiter\"></span><span class=\"mfrac\"><span class=\"vlist-t vlist-t2\"><span class=\"vlist-r\"><span class=\"vlist\" style=\"height:1.0395em;\"><span style=\"top:-2.655em;\"><span class=\"pstrut\" style=\"height:3em;\"></span><span class=\"sizing reset-size6 size3 mtight\"><span class=\"mord mtight\"><span class=\"mord mtight\">2</span><span class=\"mord mathnormal mtight\">a</span></span></span></span><span style=\"top:-3.23em;\"><span class=\"pstrut\" style=\"height:3em;\"></span><span class=\"frac-line\" style=\"border-bottom-width:0.04em;\"></span></span><span style=\"top:-3.394em;\"><span class=\"pstrut\" style=\"height:3em;\"></span><span class=\"sizing reset-size6 size3 mtight\"><span class=\"mord mtight\"><span class=\"mord mtight\">−</span><span class=\"mord mathnormal mtight\">b</span><span class=\"mbin mtight\">±</span><span class=\"mord sqrt mtight\"><span class=\"vlist-t vlist-t2\"><span class=\"vlist-r\"><span class=\"vlist\" style=\"height:0.9221em;\"><span class=\"svg-align\" style=\"top:-3em;\"><span class=\"pstrut\" style=\"height:3em;\"></span><span class=\"mord mtight\" style=\"padding-left:0.833em;\"><span class=\"mord mtight\"><span class=\"mord mathnormal mtight\">b</span><span class=\"msupsub\"><span class=\"vlist-t\"><span class=\"vlist-r\"><span class=\"vlist\" style=\"height:0.7463em;\"><span style=\"top:-2.786em;margin-right:0.0714em;\"><span class=\"pstrut\" style=\"height:2.5em;\"></span><span class=\"sizing reset-size3 size1 mtight\"><span class=\"mord mtight\">2</span></span></span></span></span></span></span></span><span class=\"mbin mtight\">−</span><span class=\"mord mtight\">4</span><span class=\"mord mathnormal mtight\">a</span><span class=\"mord mathnormal mtight\">c</span></span></span><span style=\"top:-2.8821em;\"><span class=\"pstrut\" style=\"height:3em;\"></span><span class=\"hide-tail mtight\" style=\"min-width:0.853em;height:1.08em;\"><svg xmlns=\"http://www.w3.org/2000/svg\" width='400em' height='1.08em' viewBox='0 0 400000 1080' preserveAspectRatio='xMinYMin slice'><path d='M95,702\nc-2.7,0,-7.17,-2.7,-13.5,-8c-5.8,-5.3,-9.5,-10,-9.5,-14\nc0,-2,0.3,-3.3,1,-4c1.3,-2.7,23.83,-20.7,67.5,-54\nc44.2,-33.3,65.8,-50.3,66.5,-51c1.3,-1.3,3,-2,5,-2c4.7,0,8.7,3.3,12,10\ns173,378,173,378c0.7,0,35.3,-71,104,-213c68.7,-142,137.5,-285,206.5,-429\nc69,-144,104.5,-217.7,106.5,-221\nl0 -0\nc5.3,-9.3,12,-14,20,-14\nH400000v40H845.2724\ns-225.272,467,-225.272,467s-235,486,-235,486c-2.7,4.7,-9,7,-19,7\nc-6,0,-10,-1,-12,-3s-194,-422,-194,-422s-65,47,-65,47z\nM834 80h400000v40h-400000z'/></svg></span></span></span><span class=\"vlist-s\">​</span></span><span class=\"vlist-r\"><span class=\"vlist\" style=\"height:0.1179em;\"><span></span></span></span></span></span></span></span></span></span><span class=\"vlist-s\">​</span></span><span class=\"vlist-r\"><span class=\"vlist\" style=\"height:0.345em;\"><span></span></span></span></span></span><span class=\"mclose nulldelimiter\"></span></span></span></span></span></span>",
                                "children": null
                            }, {
                                "depth": 3,
                                "content": "<a href=\"#?d=gist:af76a4c245b302206b16aec503dbe07b:katex.md\">More Katex Examples</a>",
                                "children": null
                            }]
                        }, {
                            "depth": 2,
                            "content": "Now we can wrap very very very very long text based on <code>maxWidth</code> option",
                            "children": null
                        }]
                }]
        }
         */
        window.addEventListener('DOMContentLoaded', () => {
            const headingSet = [...document.querySelectorAll("h1, h2, h3, h4, h5, h6")]
            const tocItems = parseHeading(headingSet)
            if (tocItems.length === 0) {
                return
            }

            const idName = 'markmap-toc'
            const frag = document.createRange().createContextualFragment(`
        <div class="markmapWrapper navbar-right">
        <svg id="${idName}" class="markmap" />
        </div>
        `)
            const svgElem = frag.querySelector(`svg`) as SVGElement

            // 使用z-index來輔助，不需要倚靠js來隱藏
            // svgElem.onmouseover = hideBtnCopyPre // 這是我們自己創建用來copy code block的按鈕，在顯示markmap的時候，這個按鈕會和markmap所提供的SVG相衝，所以把它隱藏
            // svgElem.onmouseout = showBtnCopyPre

            // navElem.replaceWith(svgElem)
            const tocContainer = frag.querySelector('div')
            document.body.append(frag)

            const dictData: markMapData = tocItems[0] as markMapData // 理論上都是h1, h2, ... 不支持h2, h2, ...的這種模式，也就是第一個一定是採用h1
            Toc.createMarkmap(idName, dictData)
            initSVGHoverAttr(svgElem)  // 要放在最後面，因為計算hover的寬度會需要用到svg的位置資訊

            /* 不需要用js調整，直接用js把top設定成20%就行了
            window.addEventListener('scroll', () => {
                const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
                tocContainer!.style.top = `${Math.max(120, scrollTop + 20)}px` // 預設為120
            })
             */
        })
    }
)()
