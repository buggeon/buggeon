import { useEffect, useRef, useState } from 'react'
import styles from './standartInput.module.scss'

export interface StandartInputProps {
    title? : string
    prefixIcon? : React.ReactNode
    suffixIcon? : React.ReactNode
    placeholder : string
    onSuffixClick? : () => void
    type? : string
    onChange : (text : string) => void
    subtitle? : string
    isNecessary? : boolean
    maxChars? : number
    style? : React.CSSProperties
    value? : string
}

function StandartInput(props : StandartInputProps) {

    const [isNecessaryPromptVisible, changeNecessaryPromptVisibility] = useState(false)
    const [charsCount, changeCharsCount] = useState(0)
    const starRef = useRef<HTMLParagraphElement>(null)

    useEffect(() => {
        const star = starRef.current
        if(!star) return
        
        const handleMouseOver = () => changeNecessaryPromptVisibility(true)
        const handleMouseLeave = () => changeNecessaryPromptVisibility(false)

        star.addEventListener('mouseover', handleMouseOver)
        star.addEventListener('mouseleave', handleMouseLeave)

        return () => {
            star.removeEventListener("mouseover", handleMouseOver)
            star.removeEventListener("mouseleave", handleMouseLeave)
        }
    }, [])

    const isMultiline = props.style?.height && parseInt(props.style.height.toString()) >= 60;

    return (
        <div className={styles.main}>
            {props.title && 
                <div className={styles.title}>
                    <h5>{props.title}</h5>
                    {props.isNecessary && 
                        <p className={styles.star} ref={starRef}>*</p>
                    }
                    {isNecessaryPromptVisible &&
                        <p style={{color: "var(--dark-grey)"}}>Это поле обязательно</p>
                    }
                </div>
            }
            <div className={`${styles.inputBlock} ${isMultiline ? styles.multiline : ''}`} style={props.style}>
                {props.prefixIcon && <div className={styles.prefixIcon}>{props.prefixIcon}</div>}
                
                {isMultiline ? (
                    <textarea
                        style={props.style}
                        placeholder={props.placeholder} 
                        onChange={(e) => {
                            props.onChange(e.target.value)
                            changeCharsCount(e.target.value.length)
                        }} 
                        maxLength={props.maxChars}
                        className={styles.textareaElement}
                        value={props.value}
                    />
                ) : (
                    <input 
                        placeholder={props.placeholder} 
                        type={props.type} 
                        onChange={(e) => {
                            props.onChange(e.target.value)
                            changeCharsCount(e.target.value.length)
                        }} 
                        maxLength={props.maxChars}
                        className={styles.inputElement}
                        value={props.value}
                    />
                )}

                {props.maxChars && 
                    <p className={styles.counter}>
                        {`${charsCount}/${props.maxChars}`}
                    </p>
                }
                {props.suffixIcon && <div className={styles.suffixIcon} onClick={props.onSuffixClick}>{props.suffixIcon}</div>}
            </div>
            {
                props.subtitle && <p className={styles.subtitle}>{props.subtitle}</p>
            }
        </div>
    )
}

export default StandartInput;