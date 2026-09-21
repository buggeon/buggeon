
import styles from './boards.module.scss'
import MainLayout from '../../layouts/mainLayout/mainLayout';
import { useParams } from 'react-router-dom';
import { useGetBoards } from '../../hooks/useGetBoards';
import { useEffect, useState } from 'react';
import BoardItem from './boardItem/boardItem';
import Icons from '../../assets/icons';
import { DndContext, DragOverlay, PointerSensor, useSensor, useSensors, type DragEndEvent, type DragStartEvent } from '@dnd-kit/core';
import CardPreview from './cardItem/preview/cardPreview';
import projectApi from '../../api/project.api';
import CreateBoard from './createBoard/createBoard';
import { priorityDict } from '../../dicts/priority';
import AlertModal from '../../components/alertModal/alertModal';

function BoardsScreen() {

    const { projectId } = useParams()
    
    const [isNewBoardModalVisible, setNewBoardModalVisibility] = useState(false)
    const { data } = useGetBoards(projectId)
    const [focusedDirection, setFocusedDirection] = useState("")
    const [groupedBoards, setGroupBoards] = useState<Record<string, typeof data.boards>>({})
    const [activeCard, setActiveCard] = useState<{
        id: string;
        title: string;
        priority: string;
        boardId: string;
        boardThemeColor: string;
        status : string;
    } | null>(null);
    
    useEffect(() => {

        if(data) {
            setGroupBoards(data.boards.reduce((acc, board) => {
                const direction = board.direction
                if (!acc[direction]) acc[direction] = []
                acc[direction].push(board)
                return acc
            }, {} as Record<string, typeof data.boards>))
        }

    }, [data])

    const sensors = useSensors(
        useSensor(PointerSensor, {
            activationConstraint: { distance: 5 }
        })
    )

    const handleDragStart = (event : DragStartEvent) => {

        const data = event.active.data.current as {
            id : string;
            title : string;
            priority : string;
            status : string;
            boardId : string;
            boardThemeColor : string;
        }

        setActiveCard(data)

    }

    const handleDragEnd = async (event : DragEndEvent) => {

        try{
            const newGroupedBoards = JSON.parse(JSON.stringify(groupedBoards)) as Record<string, typeof data.boards>;

            const { active, over } = event;
            if(!over) return

            const activeData = active.data.current as { cardId : string, boardId : string } | null
            const overData = over.data.current as { cardId : string, boardId : string } | null

            if(!activeData || !overData?.boardId) return

            const allBoards = Object.values(newGroupedBoards).flat()
            const sourceBoardId = activeData.boardId
            const targetBoardId = overData.boardId

            if(sourceBoardId == targetBoardId) return

            const sourceBoard = allBoards.find(b => b.id == sourceBoardId)
            const targetBoard = allBoards.find(b => b.id == targetBoardId)

            if(!sourceBoard || !targetBoard) return

            const cardIndex = sourceBoard.cards.findIndex(c => c.id == active.id)
            if(cardIndex == -1) return

            const [movedCard] = sourceBoard.cards.splice(cardIndex, 1)
            targetBoard.cards.push(movedCard)

            setGroupBoards(newGroupedBoards)

            try{
                await projectApi.updateCardLocation(projectId, activeCard.id, sourceBoardId, targetBoardId)
            }

            catch(e) {
                console.error(e)
            }
        }
        finally{
            setActiveCard(null)
        }

    }

    return(
        <MainLayout
            title="Project boards"
            description="Here'are all project's boards."
        >
            <DndContext sensors={sensors} onDragEnd={handleDragEnd} onDragStart={handleDragStart}>
                <section className={styles.main}>
                    <CreateBoard isVisible={isNewBoardModalVisible} setVisibility={setNewBoardModalVisibility} direction={focusedDirection}/>
                    {
                        data && data.boards.length > 0 && Object.entries(groupedBoards).map(([direction, boards]) => (
                            <section key={direction}>
                                <h3>{direction}</h3>
                                <section className={styles.boardBlock}>
                                    {
                                        boards.map(board => (
                                            <BoardItem
                                                key={board.id}
                                                id={board.id}
                                                name={board.name}
                                                cards={board.cards}
                                                projectId={projectId}
                                            />
                                        ))
                                    }
                                    <button className={styles.addBoard} onClick={() => {setFocusedDirection(direction); setNewBoardModalVisibility(true)}}>
                                        <Icons.Plus width={40} height={40} color='var(--light-grey)'/>
                                    </button>
                                </section>
                            </section>
                        ))
                    }
                    <button className={styles.addBoard} onClick={() => {setNewBoardModalVisibility(true)}}>
                        <Icons.Plus width={40} height={40} color='var(--light-grey)'/>
                    </button>
                </section>
                <DragOverlay>
                    {
                        activeCard && 
                        <CardPreview
                            title={activeCard.title}
                            priority={activeCard.priority}
                            boardId={activeCard.boardId}
                            cardId={activeCard.id}
                            status={activeCard.status}
                        />
                    }
                </DragOverlay>
            </DndContext>
        </MainLayout>
    )
}

export default BoardsScreen;