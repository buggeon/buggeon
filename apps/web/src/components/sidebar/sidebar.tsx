import { useLocation } from "react-router-dom";
import type { ISidebarMenuItem } from "./menuItem/menuItem";
import SidebarMenuItem from "./menuItem/menuItem";
import styles from "./sidebar.module.scss"
import { LogoIcon } from "../../assets/logo/logo";

function Sidebar({tabs, className} : {tabs : ISidebarMenuItem[], className? : string}) {

    const location = useLocation().pathname

    return(
        <section className={`${styles.sidebar} ${className}`}>
            <header>
                <LogoIcon width={30} height={30}/>
                <h2>Buggeon</h2>
            </header>
            <ul>
                {
                    tabs.map((item) => (
                        <li key={item.label}>
                            <SidebarMenuItem itemData={item} selected={location.includes(item.path)}/>
                        </li>
                    ))
                }
            </ul>

            <section className={styles.account}>
                <img src="/assets/rpp.png" className={styles.accountAvatar}/>
                <div className={styles.accountPersonalInfo}>
                    <p className={styles.accountName}>Jane Cooper</p>
                    <p className={styles.accountRole}>Project Manager</p>
                </div>
                <img src="/assets/arrow-down.svg" className={styles.accountArrow}/>
            </section>

        </section>
    )
}

export default Sidebar;