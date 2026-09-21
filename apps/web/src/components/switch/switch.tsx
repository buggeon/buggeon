import { useSelector } from 'react-redux'
import styles from './switch.module.scss'
import type { RootState } from '../../store/slices'
import themeDict from '../../dicts/theme'

export interface SwitchProps {
    enabled : boolean
    onChange : () => void
}

function Switch({enabled, onChange} : SwitchProps) {

    return(

        <div
            className={styles.main}
            style={{backgroundColor: enabled ? "var(--color-accent)" : "var(--color-bg-elevated)"}}
            onClick={onChange}
        >
            <div    
                className={styles.circle}
                style={{transform: enabled ? "translateX(20px)" : "translateX(0)"}}
            />
        </div>
        
    )

}

export default Switch