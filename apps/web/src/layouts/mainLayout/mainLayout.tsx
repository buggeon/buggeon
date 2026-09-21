import { useLocation, useNavigate, useParams } from "react-router-dom";
import type { ISidebarMenuItem } from "../../components/sidebar/menuItem/menuItem";
import Sidebar from "../../components/sidebar/sidebar";
import styles from './mainLayout.module.scss'
import Icons from "../../assets/icons";
import { useMemo } from "react";
import themeDict from "../../dicts/theme";
import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/slices";
import { setAppearance } from "../../store/slices/themeSlice";

function MainLayout({children, title, description, isModal = false}) {
    
    const globalLocation = useLocation().pathname.split("/")[1]
    const locationParts = useLocation().pathname.split("/")
    const appearanceState = useSelector((state : RootState) => state.appearance)
    const navigator = useNavigate()
    const { projectId } = useParams()
    const dispatch = useDispatch()

    const projectControlPanelMenuItems = useMemo<ISidebarMenuItem[]>(() => [
        {
            icon: <Icons.HomeOutline width={20} height={20} color="var(--color-text"/>,
            label: "Overview",
            path: `/project/${projectId}/overview`
        },
        {
            icon: <Icons.BoardOutline width={20} height={20} color="var(--color-text"/>,
            label: "Boards",
            path: `/project/${projectId}/boards`
        },
        {
            icon: <Icons.SettingsOutline width={20} height={20} color="var(--color-text"/>,
            label: "Settings",
            path: `/project/${projectId}/settings`
        },
    ], [projectId]);

    const dashboardMenuItems = useMemo<ISidebarMenuItem[]>(() => [
            {
                icon: <Icons.HomeOutline width={20} height={20} color="var(--color-text"/>,
                label: "Profile",
                path: "/dashboard/profile"
            },
            {
                icon: <Icons.FoldersOutline width={20} height={20} color="var(--color-text"/>,
                label: "Projects",
                path: "/dashboard/projects"
            },
            {
                icon: <Icons.SettingsOutline width={20} height={20} color="var(--color-text"/>,
                label: "Settings",
                path: "/dashboard/settings"
            },
    ], [projectId])

    return (
        
        <section className={styles.app}>
            <Sidebar tabs={globalLocation == "dashboard" ? dashboardMenuItems : projectControlPanelMenuItems} className={styles.sidebar}/>
            <section className={styles.screen}>
                <div className={styles.header}>
                    <div className={styles.highSection}>
                        {
                            isModal
                            && <div onClick={() => {
                                navigator(-1)
                            }} className={styles.backBlock}>
                                <Icons.Arrow2 width={20} height={20} color="var(--light-grey)"/>
                                <p style={{color: "var(--light-grey)"}}>Back to {locationParts.at(-2)[0].toUpperCase() + locationParts.at(-2).slice(1, locationParts.at(-2).length)}</p>
                            </div>
                        }
                        <div className={styles.buttonBlock}>
                            <button>
                                <Icons.BellOutline width={20} height={20} color="#ffffff"/>
                            </button>
                            <button onClick={() => dispatch(setAppearance({theme: appearanceState.theme == "dark" ? "light" : "dark", accentColor: appearanceState.accentColor}))}>
                                {
                                    appearanceState.theme == "light"
                                        ?   <Icons.MoonOutline
                                                width={20}
                                                height={20}
                                                color="#ffffff"
                                            />
                                        :   <Icons.Sun
                                                width={20}
                                                height={20}
                                                color="#ffffff"
                                            />
                                }
                            </button>
                        </div>
                    </div>
                    <h1>{title}</h1>
                    <p style={{color: "var(--color-text)"}}>{description}</p>
                </div>
                {children}
            </section>
        </section>
    )
}

export default MainLayout;