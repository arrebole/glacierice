import {openDataApi, thanks} from "./static_data"
import {Get} from "./fetch"
export class Api {
    async thanks() {
        return {
            code: 0,
            message: "success",
            data: thanks
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
        data: await Get("/api/data?type=contributor")
      }
    }
    async dataResource(){
      return {
        code:0,
        message: "success",
        data: await Get("/api/data?type=resource_3rd")
      }
    }
}

export default function install(Vue) {
  Vue.prototype.$api = new Api()
}