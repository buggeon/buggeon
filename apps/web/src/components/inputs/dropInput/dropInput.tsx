import { useEffect, useRef, useState } from 'react'
import styles from './dropInput.module.scss'
import { ArrowIcon } from '../../../assets/arrow/arrow'

export interface DropInputProps {
    title?: string
    placeholder: string
    onChange: (item: any) => void
    isNecessary?: boolean
    content: any[]
    renderFunc: (item: any) => React.ReactNode
    style?: React.CSSProperties
    value?: any
    subtitle? : string
}

function DropInput({ 
    title, 
    placeholder, 
    onChange, 
    isNecessary = false, 
    content = [], 
    renderFunc, 
    style,
    value,
    subtitle
}: DropInputProps) {
    const [isNecessaryPromptVisible, setNecessaryPromptVisible] = useState(false)
    const [isListVisible, setIsListVisible] = useState(false)
    const [searchValue, setSearchValue] = useState("")
    const [selectedItem, setSelectedItem] = useState<any>(value || null)
    
    const wrapperRef = useRef<HTMLDivElement>(null)
    const inputRef = useRef<HTMLInputElement>(null)

    const displayValue = selectedItem 
        ? (typeof selectedItem === 'object' ? selectedItem.name : selectedItem)
        : searchValue

    useEffect(() => {
        const handleClickOutside = (e: MouseEvent) => {
            if (wrapperRef.current && !wrapperRef.current.contains(e.target as Node)) {
                setIsListVisible(false)
            }
        }
        document.addEventListener('mousedown', handleClickOutside)
        return () => document.removeEventListener('mousedown', handleClickOutside)
    }, [])

    const filteredItems = content.filter(item => {
        const itemName = typeof item === 'object' ? item.name?.toLowerCase() || '' : String(item).toLowerCase()
        const search = searchValue.toLowerCase()
        return itemName.includes(search)
    })

    const handleFocus = () => {
        setIsListVisible(true)
    }

    const handleItemSelect = (item: any) => {
        console.log(item)

        setSelectedItem(item)
        setSearchValue("")
        setIsListVisible(false)
        onChange(item)
        inputRef.current?.blur()
    }

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value
        setSearchValue(value)
        onChange(value)
        setSelectedItem(null)
        setIsListVisible(true)
        
        if (value === "") {
            onChange(null)
        }
    }

    const handleToggleList = (e : React.MouseEvent) => {
        e.preventDefault();
        e.stopPropagation();

        setIsListVisible(!isListVisible)
        if (!isListVisible) {
            inputRef.current?.focus()
        }
    }

    return (
        <div className={styles.main} ref={wrapperRef}>
            {title && (
                <div className={styles.title}>
                    <h5>{title}</h5>
                    {isNecessary && (
                        <p 
                            className={styles.star} 
                            id="star"
                            onMouseOver={() => setNecessaryPromptVisible(true)}
                            onMouseLeave={() => setNecessaryPromptVisible(false)}
                        >*</p>
                    )}
                    {isNecessary && isNecessaryPromptVisible && (
                        <p className={styles.necessaryPrompt}>Это поле обязательно</p>
                    )}
                </div>
            )}
            
            <div className={styles.inputBlock} style={style}>
                <input
                    ref={inputRef}
                    placeholder={placeholder}
                    value={displayValue}
                    onChange={handleInputChange}
                    onFocus={handleFocus}
                    autoComplete="off"
                />
                <ArrowIcon
                    width={20}
                    height={20}
                    color='var(--dark-grey)'
                    className={`${styles.suffixIcon} ${isListVisible ? styles.rotated : ''}`}
                    onClick={handleToggleList}
                />
            </div>

            {isListVisible && filteredItems.length > 0 && (
                <div className={styles.list}>
                    {filteredItems.map((item, index) => (
                        <div
                            key={index}
                            className={styles.listItem}
                            onMouseDown={(e) => {
                                e.preventDefault()
                                handleItemSelect(item)
                            }}
                        >
                            {renderFunc ? renderFunc(item) : (
                                <span>{typeof item === 'object' ? item.name : item}</span>
                            )}
                        </div>
                    ))}
                </div>
            )}

            {isListVisible && filteredItems.length === 0 && (
                <div className={styles.list}>
                    <div className={styles.emptyState}>
                        <span>Ничего не найдено</span>
                    </div>
                </div>
            )} 

            {
                subtitle && <p className={styles.subtitle}>{subtitle}</p>
            }
        </div>
    )
}

export default DropInput