import { config } from "../../config.js";
export { Now, };
function Now() {
    return new Date().toLocaleString("zh-TW", config.datetimeFormat);
}
