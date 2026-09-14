import { createPortal } from 'react-dom'
import styles from "./modal.module.scss"
import { useEffect } from 'react'
import Icons from '../../assets/icons'

interface ModalProps {
    isOpen : boolean
    title : string
    setState : (state : boolean) => void
    children : React.ReactNode
    className? : string
    footerActiveButtonLabel : string
    onFooterActiveButtonClick : () => void
}

function Modal({isOpen, title, setState, children, className, onFooterActiveButtonClick, footerActiveButtonLabel} : ModalProps) {

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
            <div className={`${styles.form} ${className}`} onClick={e => e.stopPropagation()}>
                <header>
                    <p>{title}</p>
                    <Icons.Cross width={20} height={20} color='white' onClick={() => {setState(false)}} className={styles.cross}/>
                </header>
                {children}
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

export default Modal