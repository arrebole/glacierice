import { openDataApi, thanks, crashCourse, papers } from "./static_data"
import { Get } from "./fetch"
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
  async contributor() {
    return {
      code: 0,
      message: "success",
      data: await Get("http://glaciericer.gkdark.xyz/api/data?type=contributor")
    }
  }
  async dataResource() {
    return {
      code: 0,
      message: "success",
      data: await Get("http://glaciericer.gkdark.xyz/api/data?type=resource_3rd")
    }
  }
  async crashCourse() {
    return {
      code: 0,
      message: "success",
      data: {
        link: "https://github.com/wuhan2020/Covid-19-data-science/tree/master/Wiki",
        html: crashCourse
      }
    }
  }
  async papers() {
    return {
      code: 0,
      message: "success",
      data: {
        link: "https://github.com/wuhan2020/Covid-19-data-science/blob/master/Reference/paper.md",
        html: papers
      }
    }
  }
  async modelsOfDay() {
    return {
      code: 0,
      message: "success",
      data: await await Get("http://glaciericer.gkdark.xyz/api/data?type=resource_github")
    }
  }
}

export default function install(Vue) {
  Vue.prototype.$api = new Api()
}