import { useLocation, useNavigate, useNavigation } from "react-router-dom";
import type { ISidebarMenuItem } from "../sidebarMenuItem/sidebarMenuItem";
import SidebarMenuItem from "../sidebarMenuItem/sidebarMenuItem";
import styles from "./sidebar.module.scss"

function Sidebar({tabs, className, header, footer} : {tabs : ISidebarMenuItem[], className? : string, header? : React.ReactNode, footer? : React.ReactNode}) {

    const location = useLocation().pathname
    const navigate = useNavigate()

    return(
        <section className={`${styles.sidebar} ${className}`}>
            {header}
            <ul>
                {
                    tabs.map((item) => (
                        <li key={item.label}>
                            <SidebarMenuItem itemData={item} selected={location.includes(item.path)} onClick={() => navigate(item.path)}/>
                        </li>
                    ))
                }
            </ul>

            {footer}

        </section>
    )
}

export default Sidebar;