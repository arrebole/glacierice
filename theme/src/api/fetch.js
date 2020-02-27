export function Get(url){
    return fetch(url).then(res=>res.json())
}