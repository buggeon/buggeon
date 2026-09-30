import { useParams } from "react-router-dom";
import { useGetMember } from "../../../hooks/useGetMember";
import MainLayout from "../../../layouts/mainLayout/mainLayout";
import styles from './memberAbout.module.scss'
import Mark from "../../../components/mark/mark";
import { useState, type ReactNode } from "react";
import Icons from "../../../assets/icons";
import Switch from "../../../components/switch/switch";
import AvatarImage from "../../../components/avatarImage/avatarImage";

type BoardPermission = "Create boards" | "Delete boards" | "Edit boards"
type IssuePermission = "Create issues" | "Edit issues" | "Delete issues"
type MemberPermission = "Add members" | "Delete members" | "Edit members"
type ProjectPermission = "Edit project"

type PermissionContent = {
    descritpion : string
    icon : ReactNode
    state : boolean
}

type Permissions = {
    "Boards": Record<BoardPermission, PermissionContent>
    "Issues": Record<IssuePermission, PermissionContent>
    "Members": Record<MemberPermission, PermissionContent>
    "Project Managment": Record<ProjectPermission, PermissionContent>
}

function MemberAboutScreen() {

    const { memberId } = useParams()
    const { memberData, loading } = useGetMember(memberId, ["role", "user{name, email, avatarKey}"])
    const [permissions, setPermissions] = useState<Permissions>({
        "Boards": {
            "Create boards": {
                descritpion: "",
                icon: <Icons.Plus/>,
                state: false
            },
            "Delete boards": {
                descritpion: "",
                icon: <Icons.Trash/>,
                state: false
            },
            "Edit boards": {
                descritpion: "",
                icon: <Icons.SettingsOutline/>,
                state: false
            }
        },
        "Issues": {
            "Create issues": {
                descritpion: "",
                icon: <Icons.Plus/>,
                state: false
            },
            "Delete issues": {
                descritpion: "",
                icon: <Icons.Trash/>,
                state: false
            },
            "Edit issues": {
                descritpion: "",
                icon: <Icons.SettingsOutline/>,
                state: false
            },
        },
        "Members": {
            "Add members": {
                descritpion: "",
                icon: <Icons.Plus/>,
                state: false
            },
            "Delete members": {
                descritpion: "",
                icon: <Icons.Trash/>,
                state: false
            },
            "Edit members": {
                descritpion: "",
                icon: <Icons.SettingsOutline/>,
                state: false
            },
        },
        "Project Managment": {
            "Edit project": {
                descritpion: "",
                icon: <Icons.SettingsOutline/>,
                state: false
            },
        },
    })

    return(
        <MainLayout
            isModal
        >
            <main className={styles.main}>
                <header className={styles.header}> 
                    <button>Save changes</button>
                </header>
                    <section className={styles.app}>
                        <section className={styles.memberInfo}>
                        <AvatarImage imageKey={memberData.member.user.avatarKey} title={memberData.member.user.name}/>
                        <div>
                            <p>{memberData?.member.user.name}</p>
                            <p>{memberData?.member.user.email}</p>
                        </div>
                        <Mark
                            title={memberData?.member.role}
                            backgroundColor="var(--background-blue)"
                            foregroundColor="var(--foreground-blue)"
                        />
                    </section>
                    <section className={styles.permissionsInfo}>
                        <header>
                            <h2>Permissions</h2>
                            <p>Set what member can do across the platform.</p>
                        </header>
                        <ul className={styles.permissions}>
                            {
                                Object.entries(permissions).map(([key, value]) => (
                                    <li>
                                        <h3>{key}</h3>
                                        <ul className={styles.permissionBlock}>
                                            {
                                                Object.entries(value).map(([permissionName, permissionContent]) => (
                                                    <li>
                                                        <div>
                                                            {permissionContent.icon}
                                                            <p>{permissionName}</p>
                                                        </div>
                                                        <p>{permissionContent.descritpion}</p>
                                                        <Switch
                                                            enabled={permissionContent.state}
                                                            onChange={() => setPermissions(prev => ({
                                                                ...prev,
                                                                [key]: {
                                                                    ...prev[key],
                                                                    [permissionName]: {
                                                                        ...prev[key][permissionName],
                                                                        state: !prev[key][permissionName].state
                                                                    }
                                                                }
                                                            }))}
                                                        />
                                                    </li>
                                                ))
                                            }
                                        </ul>
                                    </li>
                                ))
                            }
                        </ul>
                    </section>
                </section>
            </main>
        </MainLayout>
    )

}

export default MemberAboutScreen