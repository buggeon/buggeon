import { useNavigate } from "react-router-dom"
import styles from "./sidebarMenuItem.module.scss"

export interface ISidebarMenuItem {
    icon : React.ReactNode
    label : string
    path? : string
}

function SidebarMenuItem({itemData, onClick, selected = false} : {itemData : ISidebarMenuItem, selected? : boolean, onClick : () => void}) {

    const navigate = useNavigate()

    return(
        <div className={`${styles.item} ${selected && styles.selected}`} onClick={onClick}>
            {itemData.icon}
            <p>{itemData.label}</p>
        </div>
    )
}

export default SidebarMenuItem;