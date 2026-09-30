import { useNavigate, useParams } from "react-router-dom"
import { useGetMembers } from "../../hooks/useGetMembers"
import MainLayout from "../../layouts/mainLayout/mainLayout"
import styles from './membersScreen.module.scss'
import Mark from "../../components/mark/mark"
import Icons from "../../assets/icons"
import { use, useState } from "react"
import Modal from "../../components/modal/modal"
import StandartInput from "../../components/inputs/standartInput/standartInput"
import { useGetUsers } from "../../hooks/useGetUsers"
import AddMember from "./addMember/addMember"
import AvatarImage from "../../components/avatarImage/avatarImage"

function MembersScreen() {

    type RoleType = "All Members" | "Leads" | "Developers" | "Designers"

    const { projectId } = useParams()
    const { membersData, loading } = useGetMembers(projectId)
    const navigate = useNavigate()
    const [roleFilter, setRoleFilter] = useState<RoleType>("All Members")
    const [isInviteMemberModalVisible, setInviteMemberModalVisibility] = useState(false)



    return(
        <MainLayout
            title="Members"
            description="Manage your team DmembersData, their roles and permissions."
        >
            <main className={styles.main}>
                <Modal
                    isOpen={isInviteMemberModalVisible}
                    setState={setInviteMemberModalVisibility}
                    title="Add Member"
                    className={styles.modal}
                >
                    <AddMember/>
                </Modal>
                <header>
                    <div className={styles.rolesList}>
                        {
                            ["All Members", "Leads", "Developers", "Designers"].map(role => (
                                <div
                                    key={role}
                                    className={role == roleFilter ? styles.active : ""}
                                    onClick={() => setRoleFilter(role as RoleType)}
                                >
                                    <p>{role}</p>
                                </div>
                            ))
                        }
                    </div>
                    <button onClick={() => setInviteMemberModalVisibility(true)}>
                        <Icons.Plus
                            width={20}
                            height={20}
                            color="var(--color-text)"
                        />
                        <p>Add Member</p>
                    </button>
                </header>
                <table className={styles.table}>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Role</th>
                            <th>Permissions</th>
                            <th>Status</th>
                            <th>Joined</th>
                            <th/>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            membersData && membersData.members.filter(member => roleFilter.includes(member.role) || roleFilter == "All Members").map(member => (
                                <tr onClick={() => navigate(`/projects/${projectId}/members/${member.id}`)} key={member.id}>
                                    <td>
                                        <div className={styles.personalInfo}>
                                            <AvatarImage imageKey={member.user.avatarKey} title={member.user.name}/>
                                            <div>
                                                <p>{member.user.name}</p>
                                                <p>{member.user.email}</p>
                                            </div>
                                        </div>
                                    </td>
                                    <td>
                                        <Mark
                                            title={member.role}
                                            backgroundColor="var(--background-blue)"
                                            foregroundColor="var(--foreground-blue)"
                                            className={styles.role}
                                        />
                                    </td>
                                    <td>
                                        <div>
                                            <p></p>
                                            <p></p>
                                        </div>
                                    </td>
                                    <td>
                                        <div className={styles.status}>
                                            <div style={{backgroundColor: "var(--green)"}}/>
                                            <p>Online</p>
                                        </div>
                                    </td>
                                    <td>
                                        <p>{member.createdAt}</p>
                                    </td>
                                    <td>
                                        <div className={styles.options}>
                                            <Icons.Options
                                                width={20}
                                                height={20}
                                                color="var(--color-text)"
                                                onClick={(e) => {
                                                    e.stopPropagation()
                                                    alert(1)
                                                }}
                                            />
                                        </div>
                                    </td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </main>
        </MainLayout>
    )

}

export default MembersScreen