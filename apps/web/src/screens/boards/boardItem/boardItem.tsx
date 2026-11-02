import { SortableContext, verticalListSortingStrategy } from '@dnd-kit/sortable';
import CardItem from '../cardItem/cardItem';
import styles from './boardItem.module.scss'
import { useDroppable } from '@dnd-kit/core';
import Icons from '../../../assets/icons';
import { useState } from 'react';
import CreateCard from '../createCard/createCard';
import { useMutation } from '@apollo/client/react';
import { DeleteBoardDocument } from '../../../graphql/generated/graphql';
import { useGetBoards } from '../../../hooks/useGetBoards';
import Modal from '../../../components/modal/modal';
import EditBoard from '../editBoard/editBoard';
import type { IBoardData } from '../editBoard/editBoard'
import useVisibleCards from '../../../hooks/useVisibleCards';

export interface BoardItemProps {

    id : string
    name : string
    direction : string
    cardsStatus : string
    cards : any
    projectId : string

}

function BoardItem({id, name, direction, cardsStatus, projectId, cards} : BoardItemProps) {

    const [isNewCardModalVisible, setNewCardModalVisibility] = useState(false)
    const [isEditBoardModalVisible, setEditBoardModalVisibility] = useState(false)
    const [isDeleteBoardAlertModalVisible, seteleteBoardAlertModalVisibility] = useState(false)
    const { refetch } = useGetBoards(projectId)

    const { setNodeRef } = useDroppable({
        id: `board-${id}`,
        data: {
            boardId: id
        }
    })

    const cardIds = cards.map(card => card.id)

    const [deleteBoard, { loading, error }] = useMutation(DeleteBoardDocument)
    
    const handleCreateProjectButtonClick = async () => {

        console.log("===CREATE PROJECT===")

        try {
            await deleteBoard({
                variables: {
                    projectId: projectId,
                    boardId: id
                }
            })
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }

    return(
        <section className={styles.main} ref={setNodeRef} key={id}>
            <Modal
                isOpen={isDeleteBoardAlertModalVisible}
                setState={seteleteBoardAlertModalVisibility}
                onFooterActiveButtonClick={async () => {
                    await handleCreateProjectButtonClick()
                    await refetch()
                }}
                footerActiveButtonLabel='Удалить'
                footerActiveButtonColor='var(--red)'
                title='Удаление доски'
            >
                <p>Вы действительно хотите удалить эту доску?</p>
            </Modal>
            <EditBoard
                isVisible={isEditBoardModalVisible}
                setVisibility={setEditBoardModalVisibility}
                boardData={{
                    name: name,
                    direction: direction,
                    cardsStatus: cardsStatus,
                    id: id
                }}
            />
            <CreateCard
                isVisible={isNewCardModalVisible}
                setVisibility={setNewCardModalVisibility}
                boardId={id}
            />
            <div className={styles.header}>
                <p>{name}</p>
                <div className={styles.issuesAmount}>
                    <p>{cards.length}</p>
                </div>
            </div>
            <SortableContext items={cardIds} strategy={verticalListSortingStrategy}>
                {
                    cards.map((card, index) => (
                        <CardItem
                            key={card.id}
                            boardId={id}
                            id={card.id}
                            title={card.title}
                            priority={card.priority}
                            status={card.status}
                            //customStyle={{display: index + 1 > visibleCards ? "none" : ""}}
                        />
                    ))
                }
                <button className={styles.addBoard} onClick={() => {
                    setNewCardModalVisibility(true)
                }}>
                    <Icons.Plus width={20} height={20} color='var(--light-grey)'/>
                </button>
                <div className={styles.footer}>
                    <Icons.Edit
                        width={20}
                        height={20}
                        color='white'
                        onClick={() => setEditBoardModalVisibility(true)}
                    />
                    <Icons.Trash width={20} height={20} color='white' onClick={() => seteleteBoardAlertModalVisibility(true)}/>
                </div>
            </SortableContext>
        </section>
    )
}

export default BoardItem;