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
    const [cardColorSchemeType, setCardColorSchemeType] = useState<"priority" | "fixed">("priority")
    const [boardThemeColor, setBoardThemeColor] = useState("#715bc5")
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
                        themeColor: cardColorSchemeType == "priority" ? "" : boardThemeColor
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
                <div>
                    <p className={styles.chooseCardColorSchemeTitle}>Card Color Scheme</p>
                    <p className={styles.chooseCardColorSchemeDescription}>Choose how cards are color-coded on this board</p>
                </div>
                <div className={styles.cardColorSchemeVariantsBlock}>
                    <section onClick={() => setCardColorSchemeType('priority')} style={cardColorSchemeType == "priority" ? {borderStyle: "solid", borderColor: "var(--foreground-purple)", borderWidth: 0.5} : {}}>
                        <div className={styles.chooseCardColorSchemeVariantHeader}>
                            <div className={styles.chooseCardColorSchemeVariantTitleBlock}>
                                <RadioButton value={cardColorSchemeType == "priority"} onChange={() => {
                                    setCardColorSchemeType("priority")
                                }}/>
                                <p>By Priority</p>
                            </div>
                            <p>Cards are colored based on priority lavel</p>
                        </div>
                        <div style={{display: "flex", justifyContent: "space-between", width: "100%", gap: 10}}>
                            {
                                Object.entries(priorityDict).map(([title, colors]) => (
                                    <div style={{display: "flex", gap: 5, alignItems: "center"}}>
                                        <div style={{width: 10, height: 10, borderRadius: 50, backgroundColor: colors.foregroundColor}}></div>
                                        <p>{title}</p>
                                    </div>
                                ))
                            }
                        </div>
                    </section>
                    <section onClick={() => setCardColorSchemeType('fixed')} style={cardColorSchemeType == "fixed" ? {borderStyle: "solid", borderColor: "var(--foreground-purple)", borderWidth: 0.5} : {}}>
                        <div className={styles.chooseCardColorSchemeVariantHeader}>
                            <div className={styles.chooseCardColorSchemeVariantTitleBlock}>
                                <RadioButton value={cardColorSchemeType == "fixed"} onChange={() => {
                                    setCardColorSchemeType("fixed")
                                }}/>
                                <p>Fixed Color</p>
                            </div>
                            <p>All cards use the same color</p>
                        </div>
                        <div style={{borderRadius: 10, width: "100%", backgroundColor: boardThemeColor, height: 10}}></div>
                    </section>
                </div>
                {
                    cardColorSchemeType == "fixed" &&
                        <div>
                            <p className={styles.chooseCardColorSchemeTitle}>Board Color</p>
                            <p className={styles.chooseCardColorSchemeDescription}>Choose the theme color for this board</p>
                        </div>
                }
                {
                    cardColorSchemeType == "fixed" &&
                        <div style={{display: "flex", gap: 15}}>
                            {
                                ["#715bc5", "#4da592", "#4774fa", "#eb5940", "#ef6f1d", "#a0abbd"].map(color => (
                                    <div style={{
                                        backgroundColor: color,
                                        borderRadius: 5,
                                        width: 30,
                                        height: 30,
                                        display: "flex",
                                        alignItems: "center",
                                        justifyContent: "center"
                                    }} onClick={() => setBoardThemeColor(color)}>
                                        {
                                            boardThemeColor == color && <Icons.CompletedOutline width={20} height={20} color='white'/>
                                        }
                                    </div>
                                ))
                            }
                        </div>
                }
            </section>
        </Modal>
    )
}

export default CreateBoard