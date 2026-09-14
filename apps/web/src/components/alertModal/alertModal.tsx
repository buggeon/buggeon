import { createPortal } from 'react-dom'
import styles from "./alertModal.module.scss"
import { useEffect } from 'react'
import Icons from '../../assets/icons'

interface AlertModalProps {
    isOpen : boolean
    title : string
    setState : (state : boolean) => void
    onFooterActiveButtonClick : () => void
    footerActiveButtonLabel : string
    content : string
}

function AlertModal({isOpen, title, setState, onFooterActiveButtonClick, footerActiveButtonLabel, content} : AlertModalProps) {

    useEffect(() => {

        const handleEsc = (e : KeyboardEvent) => {
            if(e.key == "Escape") setState(false)
        }

        if(isOpen) {
            document.addEventListener('keydown', handleEsc)
            document.body.style.overflow = 'hidden'
        }

        return () => {
            document.removeEventListener('keydown', handleEsc)
            document.body.style.overflow = 'unset'
        }

    }, [isOpen, setState])

    if(!isOpen) return null

    return createPortal(
        <div className={styles.overlay} onClick={() => setState(false)}>
            <div className={styles.form} onClick={e => e.stopPropagation()}>
                <header>
                    <p>{title}</p>
                    <Icons.Cross width={20} height={20} color='white' onClick={() => {setState(false)}} className={styles.cross}/>
                </header>
                <p>{content}</p>
                <footer>
                    <button onClick={() => setState(false)}>Cancel</button>
                    <button onClick={async () => {
                        onFooterActiveButtonClick()
                        setState(false)
                    }}>{footerActiveButtonLabel}</button>
                </footer>
            </div>
        </div>,
        document.body
    )

}

export default AlertModal