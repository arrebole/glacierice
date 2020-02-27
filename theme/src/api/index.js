import thank from "./thank"
import openDataApi from "./open_data_api"
import contributor from "./contributor"
export class Api {
    async thank() {
        return {
            code: 0,
            message: "success",
            data: thank
        }
    }
    async openDataApi() {
      return {
          code: 0,
          message: "success",
          data: openDataApi
      }
    }
    async contributor(){
      return {
        code:0,
        message: "success",
        data: contributor
      }
    }
}

export default function install(Vue) {
  Vue.prototype.$api = new Api()
}