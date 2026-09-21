import { CSS } from '@dnd-kit/utilities';
import { useSortable } from '@dnd-kit/sortable'
import CardPreview from './preview/cardPreview';

interface CardItemProps {
    id : string;
    title : string;
    priority : string;
    boardId : string;
    status : string
}

function CardItem({id, title, priority, boardId, status} : CardItemProps) {

    console.log(status)

    const {
        attributes,
        listeners,
        setNodeRef,
        transform,
        transition,
        isDragging,
    } = useSortable({
        id: id,
        data: {
            title: title,
            priority: priority,
            id: id,
            boardId: boardId,
            status: status
        },
    });

    const style = {
        transform: CSS.Transform.toString(transform),
        transition,
        opacity: isDragging ? 0 : 1,
        cursor: isDragging ? "grab" : "pointer"
    }

    return(
        <section 
            style={style}
            ref={setNodeRef}
            {...listeners}
            {...attributes}
        >
            <CardPreview
                title={title}
                priority={priority}
                cardId={id}
                boardId={boardId}
                status={status}
            />        
        </section>
    )
}

export default CardItem;