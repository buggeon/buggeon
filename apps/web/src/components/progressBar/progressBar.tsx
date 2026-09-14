import styles from './progressBar.module.scss'

function ProgressBar({progress, color, className} : {progress : number, color : string, className? : string}) {

    return(

        <div className={`${styles.progress} ${className}`}>
            <div className={`${styles.emptyProgressBar} ${className}`}>
                <div style={{width: `${progress}%`, backgroundColor: color}}>
                </div>
            </div>
        </div>

    )

}

export default ProgressBar