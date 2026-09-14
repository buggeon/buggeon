// import { useEffect, useState, type HTMLAttributes } from 'react'
// import styles from './standartInput.module.scss'

// function BaseInput({children}) {

//     const [isNecessaryPromptVisible, changeNecessaryPromptVisibility] = useState(false)
//     const [charsCount, changeCharsCount] = useState(0)

//     useEffect(() => {
//         const star = document.getElementById("star")

//         if(star) {
//             star.addEventListener("mouseover", () => {
//                 changeNecessaryPromptVisibility(true)
//             })
//             star.addEventListener("mouseleave", () => {
//                 changeNecessaryPromptVisibility(false)
//             })

//             return () => {
//                 star.removeEventListener("mouseover", () => {})
//                 star.removeEventListener("mouseleave", () => {})
//             }
//         }

//     }, [])

//     return (
//         <div className={styles.main} style={props.style}>
//             {props.title && 
//                 <div className={styles.title}>
//                     <h5>{props.title}</h5>
//                     {props.isNecessary && 
//                         <p className={styles.star} id={props.isNecessary ? "star" : ""}>*</p>
//                     }
//                     {
//                         props.isNecessary && isNecessaryPromptVisible &&
//                             <p style={{color: "var(--dark-grey)"}}>Это поле обязательно</p>
//                     }
//                 </div>
//             }
//             <div className={styles.inputBlock}>
//                 {props.prefixIcon && <div className={styles.prefixIcon}>{props.prefixIcon}</div>}
//                 <input placeholder={props.placeholder} prefix="" type={props.type} onChange={(e) => {
//                     props.onChange(e.target.value)
//                     changeCharsCount(e.target.value.length)
//                 }} maxLength={props.maxChars} aria-multiline/>
//                 {props.maxChars && 
//                     <p>
//                         {`${charsCount}/${props.maxChars}`}
//                     </p>
//                 }
//                 {props.suffixIcon && <div className={styles.suffixIcon} onClick={props.onSuffixClick}>{props.suffixIcon}</div>}
//             </div>
//         </div>
//     )
// }

// export default BaseInput;