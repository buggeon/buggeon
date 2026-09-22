import { RUSTFS_URL } from "../../config";

function buildLink(key : string) {

    return `${RUSTFS_URL}/${key}`

}

export default buildLink