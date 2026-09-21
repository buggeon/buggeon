import { useLocation } from "react-router-dom";
import type { ISidebarMenuItem } from "./menuItem/menuItem";
import SidebarMenuItem from "./menuItem/menuItem";
import styles from "./sidebar.module.scss"
import { LogoIcon } from "../../assets/logo/logo";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/slices";
import themeDict from "../../dicts/theme";

function Sidebar({tabs, className} : {tabs : ISidebarMenuItem[], className? : string}) {

    const location = useLocation().pathname
    const user = useSelector((state : RootState) => state.user)
    
    return(
        <section className={`${styles.sidebar} ${className}`}>
            <header>
                <LogoIcon width={30} height={30} color="var(--color-accent)"/>
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
                <img src={user.avatarUrl} className={styles.accountAvatar}/>
                <div className={styles.accountPersonalInfo}>
                    <p>{user.name}</p>
                    <p className={styles.accountLogin}>{user.login}</p>
                </div>
            </section>

        </section>
    )
}

export default Sidebar;