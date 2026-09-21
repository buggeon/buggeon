import { useEffect, useState } from "react"
import projectApi from "../../../api/project.api"
import DropInput from "../../../components/inputs/dropInput/dropInput"
import StandartInput from "../../../components/inputs/standartInput/standartInput"
import Modal from "../../../components/modal/modal"
import RadioButton from "../../../components/radioButton/radioButton"
import { useGetBoards } from "../../../hooks/useGetBoards"
import styles from './createBoard.module.scss'
import { useParams } from "react-router-dom"
import { priorityDict } from "../../../dicts/priority"
import Icons from "../../../assets/icons"
import { useMutation } from "@apollo/client/react"
import { CreateBoardDocument } from "../../../graphql/generated/graphql"

interface CreateBoardProps {
    isVisible : boolean
    setVisibility
    direction : string
}

interface INewBoardData {
    name : string
    direction : string
    cardsStatus : string
}

function CreateBoard({isVisible, setVisibility, direction} : CreateBoardProps) {

    const { projectId } = useParams()
    //const [findName, setFindName] = useState("")
    const { refetch } = useGetBoards(projectId)
    const [newBoardData, setNewBoardData] = useState<INewBoardData>({name: "", direction: "", cardsStatus: "In progress"})
    const [createBoard, { loading, error }] = useMutation(CreateBoardDocument)

    useEffect(() => {

        setNewBoardData({name: "", direction: "", cardsStatus: "In progress"});

        setNewBoardData(prev => ({
            ...prev,
            direction
        }))

    }, [isVisible, direction])

    const handleCreateBoardButtonClick = async () => {

        console.log("===CREATE PROJECT===")

        try {
            const result = await createBoard({
                variables: {
                    projectId: projectId,
                    input: {
                        ...newBoardData,
                    }
                }
            })

            return result.data?.createBoard
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }

    return(
        <Modal
            isOpen={isVisible}
            setState={setVisibility}
            title='Create New Board'
            footerActiveButtonLabel='Create Board'
            onFooterActiveButtonClick={async () => {
                await handleCreateBoardButtonClick()
                await refetch()
            }}

        >
            <section className={styles.main}>
                <StandartInput
                    title='Board Name'
                    placeholder='Enter board name...'
                    subtitle='Choose a clear, descriptive name for your board'
                    onChange={(text) => {setNewBoardData(prev => ({...prev, name: text}))}}
                />
                <DropInput
                    title='Workstream/Direction'
                    placeholder='Select workstream or direction...'
                    onChange={(text) => {setNewBoardData(prev => ({...prev, direction: text}))}}
                    content={[]}
                    value={newBoardData.direction == "" ? direction : newBoardData.direction}
                    subtitle='Boards help organize work within a specific area'
                    renderFunc={() => (
                        <div></div>
                    )}
                />
                <DropInput
                    title='Card statuses'
                    placeholder="Select board status"
                    onChange={(text) => {setNewBoardData(prev => ({...prev, cardsStatus: text}))}}
                    content={["In progress", "Done"]}
                    value={newBoardData.cardsStatus}
                    subtitle='All cards on this board will have selected status'
                    renderFunc={(status : string) => (
                        <div>
                            <p>{status}</p>
                        </div>
                    )}
                />
            </section>
        </Modal>
    )
}

export default CreateBoard