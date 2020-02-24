
import marked from "marked"

export default function install(Vue) {
    Vue.prototype.$marked = marked
}