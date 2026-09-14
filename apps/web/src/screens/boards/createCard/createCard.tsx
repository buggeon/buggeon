import { useState } from "react"
import projectApi from "../../../api/project.api"
import DropInput from "../../../components/inputs/dropInput/dropInput"
import StandartInput from "../../../components/inputs/standartInput/standartInput"
import Modal from "../../../components/modal/modal"
import { useGetBoards } from "../../../hooks/useGetBoards"
import styles from './createCard.module.scss'
import { useParams } from "react-router-dom"
import type { CardData } from "../../../api/types/createCard.type"
import { useGetMembers } from "../../../hooks/useGetMembers"
import { useMutation } from "@apollo/client/react"
import { CreateCardDocument } from "../../../graphql/generated/graphql"

interface CreateCardProps {
    isVisible : boolean
    setVisibility,
    boardId
}

function CreateCard({isVisible, setVisibility, boardId} : CreateCardProps) {

    const { projectId } = useParams()
    const { data } = useGetMembers(projectId)
    const { refetch } = useGetBoards(projectId)
    const [newCardData, setNewCardData] = useState<CardData | null>({
        assignees: [""],
        title: "",
        content: "",
        boardId: "",
        projectId: "",
        priority: ""
    })
    const [createCard, { loading, error }] = useMutation(CreateCardDocument)

    const handleCreateCardButtonClick = async () => {

        console.log("===CREATE PROJECT===")

        try {
            const result = await createCard({
                variables: {
                    boardId,
                    input: {
                        content: newCardData.content,
                        dueDate: newCardData.dueDate,
                        priority: newCardData.priority,
                        title: newCardData.title
                    }
                }
            })

            return result.data?.createCard
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }

    return(
        <Modal
            isOpen={isVisible}
            setState={setVisibility}
            title='Create New Card'
            footerActiveButtonLabel='Create Card'
            onFooterActiveButtonClick={async () => {
                await handleCreateCardButtonClick()
                await refetch()
            }}
        >
            <section className={styles.main}>
                <StandartInput
                    title='Title'
                    placeholder='Enter card title...'
                    subtitle='A clear, descriptive title for your task' 
                    onChange={(text) => {setNewCardData(prev => ({...prev, title: text}))}}
                />
                <DropInput
                    value={""}
                    title='Assignees'
                    placeholder='Select assignees...'
                    onChange={(text) => { console.log(text); setNewCardData(prev => ({...prev, assignees: [...prev.assignees, text]}))}}
                    content={data ? data.members : []}
                    subtitle='Choose team members responsible for this task'
                    renderFunc={(member) => (
                        <div>
                            <p>{member.user.name}</p>
                        </div>
                    )}
                />
                <StandartInput
                    style={{height: 100}}
                    title='Content'
                    placeholder='Describe the task, add details, requirements...'
                    subtitle='Add context, requirements, or any important details'
                    onChange={(text) => {setNewCardData(prev => ({...prev, content: text}))}}
                />
                <div>
                    <StandartInput
                        title='Due Date'
                        placeholder='Select due date...'
                        subtitle='When should this be completed?'
                        onChange={(text) => {setNewCardData(prev => ({...prev, dueDate: text}))}}
                    />
                    <DropInput
                        title='Priority'
                        placeholder='Select priority...'
                        onChange={(text) => {setNewCardData(prev => ({...prev, priority: text}))}}
                        content={["Low", "Medium", "High", "None"]}
                        subtitle='Choose priority level'
                        renderFunc={(priority) => (
                            <div style={{flexDirection: "column"}}>
                                <div></div>
                                <p>{priority}</p>
                            </div>
                        )}
                    />
                </div>
            </section>
        </Modal>
    )
}

export default CreateCard