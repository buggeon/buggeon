import styles from './standartButton.module.scss'

interface StandartButtonProps{
    title : string
    onClick : () => void
    className? : string
}

function StandartButton({title, onClick, className} : StandartButtonProps) {

    return(
        <button className={`${styles.main} ${className}`} onClick={onClick}>
            {title}
        </button>
    )
    
}

export default StandartButton