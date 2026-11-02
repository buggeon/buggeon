import { FILES_URL } from "../../../config"

function buildLink(imageKey : string) {
    
    return `${FILES_URL}/${imageKey}`

}

function AvatarImage({imageKey, alt, size = 50} : {imageKey : string, alt? : string, size? : number}) {

    if(!imageKey) return

    console.log(imageKey.slice(0, 4))

    if(imageKey[0] == "#")

        return(
            <div style={{
                backgroundColor: imageKey,
                width: size,
                height: size,
                borderRadius: "100%",
                flexShrink: 0,
                display: "flex",
                alignItems: "center",
                justifyContent: "center"
            }}>
                <p style={{fontSize: 20, fontWeight: 700, color: "white"}}>{alt[0]}</p>
            </div>
    )

    if(imageKey.slice(0, 4) == "data")
        return <img src={imageKey} style={{
            width: size,
            height: size,
            borderRadius: "100%",
            flexShrink: 0
        }}/>

    return(
        <img src={buildLink(imageKey)} style={{
            width: size,
            height: size,
            borderRadius: "100%",
            flexShrink: 0
        }}/>
    )
}

export default AvatarImage