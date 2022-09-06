import {time} from "../time/init.js"

export function log2body(msg: string) {
  document.body.innerHTML += `<i>${time.Now()}</i>` + msg
}
