import styles from './infoCard.module.scss'

export interface IInfoCard{
    title : string
    value : string
    icon : React.ReactNode
    additionalInfo : string
    additionalInfoTextColor : string
}

function InfoCard(infoCardData : IInfoCard) {
    return(
        <div className={styles.main}>
            <div>
                <p>{infoCardData.title}</p>
                <h1>{infoCardData.value}</h1>
                <p style={{color: infoCardData.additionalInfoTextColor}}>{infoCardData.additionalInfo}</p>
            </div>
            {infoCardData.icon}
        </div>
    )
}

export default InfoCard