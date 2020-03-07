export function Get(url) {
    return fetch(url, {
        method: 'GET',
        mode: 'cors',
    }).then(res => res.json())
}