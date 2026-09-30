import type { ISidebarMenuItem } from "../sidebarMenuItem/sidebarMenuItem";
import SidebarMenuItem from "../sidebarMenuItem/sidebarMenuItem";
import styles from './inlineSidebar.module.scss'

function InlineSidebar({tabs, currentTabLabel, setCurrentTabLabel} : {tabs : ISidebarMenuItem[], currentTabLabel : string, setCurrentTabLabel : any}) {

    return(
        <div className={styles.sidebar}>
            <ul>
                {
                    tabs.map((item) => (
                        <li key={item.label}>
                            <SidebarMenuItem itemData={item} selected={item.label == currentTabLabel} onClick={() => setCurrentTabLabel(item.label)}/>
                        </li>
                    ))
                }
            </ul>
        </div>
    )
    
}

export default InlineSidebar