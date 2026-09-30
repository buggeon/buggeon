import { FILES_URL } from "../../config"

function buildLink(key : string) {

    return `${FILES_URL}/${key}`

}

export default buildLink