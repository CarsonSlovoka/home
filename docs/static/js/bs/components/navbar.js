function initNavbar() {
  window.onscroll = function () {
    if (initNavbar.prevScrollPos > window.pageYOffset) {
      document.querySelector(`nav`).style.top = "0"
    } else {
      document.querySelector(`nav`).style.top = "-5rem"
    }
    initNavbar.prevScrollPos = window.pageYOffset
  }
}

initNavbar.prevScrollPos = window.pageYOffset

window.addEventListener("load", () => {
  initNavbar()
})
