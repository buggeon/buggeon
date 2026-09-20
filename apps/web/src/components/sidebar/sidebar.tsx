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
    const appearanceState = useSelector((state : RootState) => state.appearance)

    return(
        <section className={`${styles.sidebar} ${className}`} style={{backgroundColor: themeDict[appearanceState.theme].foregroundColor}}>
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

            <section className={styles.account} style={{backgroundColor: themeDict[appearanceState.theme].backgroundColor}}>
                <img src={user.avatarUrl} className={styles.accountAvatar}/>
                <div className={styles.accountPersonalInfo}>
                    <p className={styles.accountName}>{user.name}</p>
                    <p className={styles.accountLogin}>{user.login}</p>
                </div>
            </section>

        </section>
    )
}

export default Sidebar;