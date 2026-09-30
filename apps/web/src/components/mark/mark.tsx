import styles from './mark.module.scss'

function Mark({title, backgroundColor, foregroundColor, className} : {title : string, backgroundColor : string, foregroundColor : string, className? : string}) {
    return(
        <div style={{backgroundColor: backgroundColor}} className={`${styles.main} ${className}`}>
            <p style={{color: foregroundColor}}>{title}</p>
        </div>
    )
}

export default Mark;