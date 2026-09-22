const API_URL = "/api"
const FILES_URL = `${location.protocol}//${location.host}/files`
const WS_URL = `${location.protocol == "https:" ? "wss:" : "ws:"}//${location.host}/ws`

export { API_URL, FILES_URL, WS_URL }