import styles from './toastModal.module.scss'

export interface ToastModalProps {
    label : string
    onClick? : () => void
    type : "warning" | "error" | "notification"
}

function ToastModal({label, onClick, type} : ToastModalProps) {

    return(
        <div className={styles.toast} onClick={onClick} style={{backgroundColor: ""}}>
            <p>{label}</p>
        </div>
    )

}

export default ToastModal