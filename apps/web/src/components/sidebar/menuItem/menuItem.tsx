import { useNavigate } from "react-router-dom"
import styles from "./menuItem.module.scss"

export interface ISidebarMenuItem {
    icon : React.ReactNode
    label : string
    path : string
}

function SidebarMenuItem({itemData, selected = false} : {itemData : ISidebarMenuItem, selected? : boolean}) {

    const navigate = useNavigate()

    return(
        <div className={`${styles.item} ${selected && styles.selected}`} onClick={() => navigate(itemData.path)}>
            {itemData.icon}
            <p>{itemData.label}</p>
        </div>
    )
}

export default SidebarMenuItem;