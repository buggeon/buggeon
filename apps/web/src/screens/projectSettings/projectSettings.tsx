import { useState } from "react"
import Icons from "../../assets/icons"
import InlineSidebar from "../../components/inlineSidebar/inlineSidebar"
import MainLayout from "../../layouts/mainLayout/mainLayout"
import styles from './projectSettings.module.scss'

function ProjectSettingsScreen() {

    const [currentTabLabel, setCurrentTabLabel] = useState("Global")

    return(
        <MainLayout
            title="Settings"
            description="Manage your project settings, team, integrations and more."
        >
            <section className={styles.screen}>
                <InlineSidebar
                    currentTabLabel={currentTabLabel}
                    setCurrentTabLabel={setCurrentTabLabel}
                    tabs={[
                        {
                            icon: <Icons.PersonOutline width={20} height={20} color="var(--color-text"/>,
                            label: "Global"
                        },
                        {
                            icon: <Icons.FolderOutline width={20} height={20} color="var(--color-text"/>,
                            label: "Project"
                        },
                        {
                            icon: <Icons.TeamOutline width={20} height={20} color="var(--color-text"/>,
                            label: "Members"
                        },
                        {
                            icon: <Icons.Palette width={20} height={20} color="var(--color-text"/>,
                            label: "Appearance"
                        },
                        {
                            icon: <Icons.Shield width={20} height={20} color="var(--color-text"/>,
                            label: "Security"
                        },
                    ]}
                />
            </section>
        </MainLayout>
    )

}

export default ProjectSettingsScreen