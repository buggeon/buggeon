import { useEffect, useState } from "react"
import projectApi from "../../../api/project.api"
import DropInput from "../../../components/inputs/dropInput/dropInput"
import StandartInput from "../../../components/inputs/standartInput/standartInput"
import Modal from "../../../components/modal/modal"
import RadioButton from "../../../components/radioButton/radioButton"
import { useGetBoards } from "../../../hooks/useGetBoards"
import styles from './editBoard.module.scss'
import { useParams } from "react-router-dom"
import { priorityDict } from "../../../dicts/priority"
import Icons from "../../../assets/icons"
import { useMutation } from "@apollo/client/react"
import { CreateBoardDocument, UpdateBoardDocument } from "../../../graphql/generated/graphql"

export interface IBoardData {
    id : string
    name : string
    direction : string
    cardsStatus : string
}

interface EditBoardProps {
    isVisible : boolean
    setVisibility
    boardData : IBoardData
}

function EditBoard({isVisible, setVisibility, boardData} : EditBoardProps) {

    console.log(boardData)

    const { projectId } = useParams()
    //const [findName, setFindName] = useState("")
    const { refetch } = useGetBoards(projectId)
    const [newBoardData, setNewBoardData] = useState<IBoardData>(boardData)
    const [updateBoard, { loading, error }] = useMutation(UpdateBoardDocument)

    useEffect(() => {

        setNewBoardData({name: boardData.name, direction: boardData.direction, cardsStatus: boardData.cardsStatus, id: boardData.id});

    }, [isVisible])

    const handleCreateBoardButtonClick = async () => {

        const {id, ...inputData} = newBoardData

        try {
            const result = await updateBoard({
                variables: {
                    boardId: boardData.id,
                    input: inputData
                }
            })

            return result.data?.updateBoard
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }

    return(
        <Modal
            isOpen={isVisible}
            setState={setVisibility}
            title='Edit Board'
            footerActiveButtonLabel='Confirm Changes'
            onFooterActiveButtonClick={async () => {
                await handleCreateBoardButtonClick()
                await refetch()
            }}

        >
            <section className={styles.main}>
                <StandartInput
                    value={newBoardData.name}
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
                    value={newBoardData.direction}
                    subtitle='Boards help organize work within a specific area'
                    renderFunc={() => (
                        <div>
                            
                        </div>
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

export default EditBoard