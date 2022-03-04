import {time} from "../time/init.js"

/**
 * @param {string} msg innerHTML
 **/
export function log2body(msg) {
  document.body.innerHTML += `<i>${time.Now()}</i>` + msg
}
