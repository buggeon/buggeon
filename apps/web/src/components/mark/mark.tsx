import styles from './mark.module.scss'

function Mark({title, backgroundColor, foregroundColor} : {title : string, backgroundColor : string, foregroundColor : string}) {
    return(
        <div style={{backgroundColor: backgroundColor}} className={styles.main}>
            <p style={{color: foregroundColor}}>{title}</p>
        </div>
    )
}

export default Mark;