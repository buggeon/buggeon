import { useEffect, useState } from 'react'
import Icons from '../../../assets/icons'
import StandartInput from '../../../components/inputs/standartInput/standartInput'
import { useGetUsers } from '../../../hooks/useGetUsers'
import styles from './addMember.module.scss'
import { useSelector } from 'react-redux'
import type { RootState } from '../../../store/slices'
import AvatarImage from '../../../components/avatarImage/avatarImage'
import { useMutation } from '@apollo/client/react'
import { CreateMemberDocument } from '../../../graphql/generated/graphql'
import { useParams } from 'react-router-dom'

function AddMember() {

    const { projectId } = useParams()

    const userState = useSelector((state : RootState) => state.user)
    const { usersData, loading } = useGetUsers()
    const [filter, setFilter] = useState("")
    const [isCopied, setCopied] = useState(false)
    const [isInvited, setInvited] = useState<Record<string, boolean>>({})
    const [createMember, { error }] = useMutation(CreateMemberDocument)

    const handleCreateMemberButtonClick = async (userId : string) => {

        try {
            const result = await createMember({
                variables: {
                    projectId: projectId,
                    input: {
                        userId,
                        role: "member",
                        rights: ["project.close.card"]
                    }
                }
            })

            return result.data?.createMember
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }

    useEffect(() => {

        if(!isCopied) return
        const id = setTimeout(() => {
            setCopied(false)
        }, 1000)
        return () => clearTimeout(id)

    }, [isCopied])

    useEffect(() => {

        if(loading) return

        const usersInitialState = {}

        usersData.users.forEach(user => {usersInitialState[user.id] = false})

        setInvited(usersInitialState)

    }, [loading])

    return(
        <main className={styles.main}>
            <section className={styles.inviteLinkBlock}>
                <div className={styles.inviteLinkDescriptionBlock}>
                    <Icons.Chain
                        width={20}
                        height={20}
                        color="var(--color-text)"
                    />
                    <div>
                        <p>Invite Link</p>
                        <p>Anybody has this link can invite to project</p>
                    </div>
                </div>
                <div className={styles.inviteLinkCopyBlock} onClick={async () => {
                    await navigator.clipboard.writeText("https://my.buggeon.com/invite/7f3a9e2c")
                    setCopied(true)
                }}>
                    <p>{"https://my.buggeon.com/invite/7f3a9e2c"}</p>
                    {
                        isCopied
                            ?   <Icons.Check
                                    width={20}
                                    height={20}
                                    color="var(--color-text)"
                                />
                            :   <Icons.Copy
                                    width={20}
                                    height={20}
                                    color="var(--color-text)"
                                />
                    }
                </div>
            </section>
            <StandartInput
                placeholder="Search by email or name..."
                prefixIcon={<Icons.Search width={20} height={20} color="var(--color-text)"/>}
                onChange={(text) => {setFilter(text)}}
            />
            <section className={styles.usersBlock}>
                <p>Found Users</p>
                <ul>
                    {
                        usersData?.users.filter(user => filter == "" || user.name.toLowerCase().includes(filter.toLowerCase()) || user.email.toLowerCase().includes(filter.toLowerCase())).filter(u => u.id != userState.id).map(user => (
                            <li key={user.id}>
                                <div>
                                    <AvatarImage imageKey={user.avatarKey} alt={user.name}/>
                                    <div>
                                        <p>{user.name}</p>
                                        <p>{user.email}</p>
                                    </div>
                                </div>
                                {
                                    !isInvited[user.id]
                                        ?   <Icons.Plus
                                                width={20}
                                                height={20}
                                                color="var(--color-text)"
                                                onClick={async () => {
                                                    handleCreateMemberButtonClick(user.id)
                                                    setInvited(prev => ({...prev, [user.id]: true}))}
                                                }
                                            />
                                        :   <Icons.Check
                                                width={20}
                                                height={20}
                                                color="var(--color-text)"
                                            />
                                }
                            </li>
                        ))
                    }
                </ul>
            </section>
        </main>
    )

}

export default AddMember